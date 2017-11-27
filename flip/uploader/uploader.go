package uploader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Telestream/telestream-cloud-go-sdk/flip"
	"strconv"
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
	Tag    string
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
func newEnv(s *flip.UploadSession) *env {
	n := min(int(s.MaxConnections), int(s.Parts))
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
	cl         *flip.FlipApi
	factoryID  string
	retryDelay time.Duration

	// DebugLog specifies an optional logger for any events which take place
	// during the upload.
	DebugLog *log.Logger
}

// New returns a new Uploader which will be uploading files to the given
// factory.
func New(cl *flip.FlipApi, factoryID string) (*Uploader, error) {
	if cl == nil {
		return nil, errors.New("uploader: client cannot be nil")
	}
	return &Uploader{
		cl:         cl,
		factoryID:  factoryID,
		retryDelay: time.Second,
	}, nil
}

// flip.UploadSession uploads files based on the given flip.UploadSession. If the upload has
// failed or has been stopped, it is possible to resume it using this method.
// Uploading stops on the first worker error and returns it.
func (u *Uploader) UploadSession(s *flip.UploadSession, r io.ReaderAt, size int64, extraFiles *ExtraFilesInfo) error {
	err := u.upload(s, r, size, "")
	if err != nil {
		return err
	}

	err = u.uploadExtraFiles(s, extraFiles)
	if err != nil {
		return err
	}

	return nil
}

// Upload is a short version of flip.UploadSession. It takes care of creating
// flip.UploadSessions and immediately starts the file upload.
func (u *Uploader) Upload(r io.ReaderAt, filename string, size int64,
	profiles string, extraFilesInfo *ExtraFilesInfo) error {
	var extraFiles []flip.ExtraFile
	if extraFilesInfo != nil {
		extraFiles = extraFilesInfo.ConvertToExtraFiles()
	}
	s, err := u.NewUploadSession(filename, size, profiles, extraFiles)
	if err != nil {
		fmt.Println("Failed to create a new session.")
		return err
	}

	err = u.upload(s, r, size, "")
	if err != nil {
		fmt.Println("Failed to upload the input file.")
		return err
	}

	err = u.uploadExtraFiles(s, extraFilesInfo)
	if err != nil {
		fmt.Println("Failed to upload extra files.")
		return err
	}

	return nil
}

func (u *Uploader) upload(s *flip.UploadSession, r io.ReaderAt, size int64, tag string) (err error) {
	u.logf("parts=%d, part_size=%d bytes, max_connections=%d, tag=%s\n", s.Parts,
		s.PartSize, s.MaxConnections, tag)

	env := newEnv(s)
	u.startWorkers(s, r, env)
	if err := u.deliverData(s, size, env, tag); err != nil {
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

	status, err := u.uploadStatus(s.Location, tag)
	if err != nil {
		return err
	}
	if len(status.MissingParts) > 0 {
		return fmt.Errorf("some parts failed to be uploaded %v",
			status.MissingParts)
	}
	return nil
}

func (u *Uploader) uploadExtraFiles(s *flip.UploadSession, extraFiles *ExtraFilesInfo) error {
	if extraFiles == nil {
		return nil
	}
	extraFilesResponse := s.ExtraFiles.(map[string]interface{})

	for _, tag := range *extraFiles {
		for i, file := range tag.Files {
			key := fmt.Sprintf("%s.index-%d", tag.Tag, i)
			efr := extraFilesResponse[key].(map[string]interface{})
			parts, err := strconv.ParseInt(efr["parts"].(string), 10, 32)
			if err != nil {
				return err
			}
			partSize, err := strconv.ParseInt(efr["part_size"].(string), 10, 32)
			if err != nil {
				return err
			}
			session := flip.UploadSession{
				Location:       s.Location,
				Parts: int32(parts),
				PartSize: int32(partSize),
				MaxConnections: s.MaxConnections,
			}
			u.upload(&session, file.File, file.Size, key)
		}
	}

	return nil
}

func (u *Uploader) startWorkers(s *flip.UploadSession, r io.ReaderAt, env *env) {
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
				retryDelay:  u.retryDelay,
				partCounter: env.partCounter,
				log:         u.DebugLog,
			}
			w.start()
		}(i)
	}
}

func (u *Uploader) deliverData(s *flip.UploadSession, fsize int64, env *env, tag string) error {
	status, err := u.uploadStatus(s.Location, tag)
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
		for i := 0; i < int(s.Parts); i, fsize = i+1, fsize-int64(s.PartSize) {
			select {
			case <-env.stopch:
				return
			default:
				if _, ok := mp[i]; !ok {
					continue
				}
				env.deliverych <- delivery{
					Offset: int64(i) * int64(s.PartSize),
					Len:    minInt64(int64(s.PartSize), fsize),
					Part:   i,
					Tag:    tag,
				}
			}
		}
	}()
	return nil
}

func (u *Uploader) uploadStatus(location, tag string) (*Status, error) {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Extra-File-Tag", tag)
	resp, err := (http.DefaultClient).Do(req)
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

// Newflip.UploadSession creates new upload flip.UploadSession which gives a unique resource location
// to upload the file. These flip.UploadSessions are being created with multi-chunk option.
func (u *Uploader) NewUploadSession(fname string, fsize int64,
	profiles string, extraFiles []flip.ExtraFile) (*flip.UploadSession, error) {
	body := flip.VideoUploadBody{
		FileSize:   fsize,
		FileName:   fname,
		Profiles:   profiles,
		MultiChunk: true,
		ExtraFiles: extraFiles,
	}

	session, _, err := u.cl.UploadVideo(u.factoryID, body)

	if err != nil {
		return nil, err
	}

	return session, nil
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
