package uploader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
)

// Status describes the object returned from the upload API. Upload Status
// contains a slice of part numbers which are still missing.
type Status struct {
	MissingParts []int `json:"missing_parts"`
}

// delivery is what a worker gets from the Uploader. It describes a part of a
// file which worker should read and send to the upload API.
type delivery struct {
	Offset int64
	Len    int64
	Part   int
}

// env provides the entire environment needed for one upload.
type env struct {
	// partCounter is being incremented with every successful part upload.
	partCounter *int32

	// workers specifies the number of concurrent workers to create.
	workers int

	// deliverych is the channel on which the file data will be sent. Workers
	// are the only consumers.
	deliverych chan delivery

	// stopch is used in case of an upload error. It guarantees to close
	// delivery channel.
	stopch chan struct{}

	// errch is used by workers to inform the uploader of any upload errors.
	errch chan error

	// wg is used to make sure all goroutines will end their work before Upload
	// returns.
	wg sync.WaitGroup
}

// newEnv returns env with every value initialized.
func newEnv(s *Session) *env {
	n := min(s.MaxConnections, s.Parts)
	return &env{
		partCounter: new(int32),
		workers:     n,
		deliverych:  make(chan delivery),
		stopch:      make(chan struct{}),
		errch:       make(chan error, n),
	}
}

// Uploader takes care of multi-chunk file upload to Telestream Cloud. It can
// be reused after every upload.
type Uploader struct {
	cl        *client.Client
	factoryID string

	// DebugLog specifies an optional logger for any events which take place
	// during the upload.
	DebugLog *log.Logger
}

// New returns a new Uploader which will be uploading files to the given
// factory.
func New(cl *client.Client, factoryID string) (*Uploader, error) {
	if cl == nil {
		return nil, errors.New("uploader: client cannot be nil")
	}
	return &Uploader{
		cl:        cl,
		factoryID: factoryID,
	}, nil
}

// UploadSession uploads files based on the given session. If the upload has
// failed or has been stopped, it is possible to resume it using this method.
// Uploading stops on the first worker error and returns it.
func (u *Uploader) UploadSession(s *Session, r io.ReaderAt, size int64) error {
	return u.upload(s, r, size)
}

// Upload is a short version of UploadSession. It takes care of creating
// sessions and immediately starts the file upload.
func (u *Uploader) Upload(r io.ReaderAt, filename string, size int64,
	profiles []string) error {
	s, err := u.NewSession(filename, size, profiles)
	if err != nil {
		return err
	}
	return u.upload(s, r, size)
}

func (u *Uploader) upload(s *Session, r io.ReaderAt, size int64) (err error) {
	u.logf("parts=%d, part_size=%d bytes, max_connections=%d\n", s.Parts,
		s.PartSize, s.MaxConnections)

	env := newEnv(s)
	u.startWorkers(s, r, env)
	if err := u.deliverData(s, size, env); err != nil {
		return err
	}

	done := make(chan struct{})
	go func() { env.wg.Wait(); close(done) }()

	var once sync.Once
loop:
	for {
		select {
		case <-done:
			break loop
		case e := <-env.errch:
			if e != nil {
				once.Do(func() {
					close(env.stopch)
					err = e
				})
			}
		}
	}

	defer u.logf("uploaded %d/%d parts\n", atomic.LoadInt32(env.partCounter),
		s.Parts)

	if err != nil {
		return err
	}

	status, err := u.uploadStatus(s.Location)
	if err != nil {
		return err
	}
	if len(status.MissingParts) > 0 {
		return fmt.Errorf("some parts failed to be uploaded %v",
			status.MissingParts)
	}
	return nil
}

func (u *Uploader) startWorkers(s *Session, r io.ReaderAt, env *env) {
	u.logf("starting %d workers\n", env.workers)

	for i := 0; i < env.workers; i++ {
		env.wg.Add(1)
		go func(i int) {
			defer env.wg.Done()
			w := &worker{
				id:          i,
				r:           r,
				location:    s.Location,
				deliverych:  env.deliverych,
				errch:       env.errch,
				retryLimit:  5,
				partCounter: env.partCounter,
				log:         u.DebugLog,
			}
			w.start()
		}(i)
	}
}

func (u *Uploader) deliverData(s *Session, fsize int64, env *env) error {
	status, err := u.uploadStatus(s.Location)
	if err != nil {
		return err
	}
	mp := make(map[int]struct{})
	for _, p := range status.MissingParts {
		mp[p] = struct{}{}
	}

	env.wg.Add(1)
	go func() {
		defer env.wg.Done()
		defer u.logf("done sending data from file")
		defer close(env.deliverych)
		for i := 0; i < s.Parts; i, fsize = i+1, fsize-s.PartSize {
			select {
			case <-env.stopch:
				return
			default:
				if _, ok := mp[i]; !ok {
					continue
				}
				env.deliverych <- delivery{
					Offset: int64(i) * s.PartSize,
					Len:    minInt64(s.PartSize, fsize),
					Part:   i,
				}
			}
		}
	}()
	return nil
}

func (u *Uploader) uploadStatus(location string) (*Status, error) {
	resp, err := http.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var status Status
	if err := json.Unmarshal(b, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

// NewSession creates new upload session which gives a unique resource location
// to upload the file. These sessions are being created with multi-chunk option.
func (u *Uploader) NewSession(fname string, fsize int64,
	profiles []string) (*Session, error) {
	params := url.Values{
		"cloud_id":    {u.factoryID},
		"file_size":   {strconv.FormatInt(fsize, 10)},
		"file_name":   {fname},
		"profiles":    profiles,
		"multi_chunk": {"1"},
	}
	u.cl.SignParams("POST", "/videos/upload.json", params)
	req, err := http.NewRequest("POST",
		"http://"+u.cl.Host+"/v3.0/videos/upload.json?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var session Session
	return &session, json.Unmarshal(b, &session)
}

func (u *Uploader) logf(format string, args ...interface{}) {
	if u.DebugLog != nil {
		u.DebugLog.Printf("uploader: "+format, args...)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
