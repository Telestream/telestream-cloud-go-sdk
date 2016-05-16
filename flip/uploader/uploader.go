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

type Status struct {
	MissingParts []int `json:"missing_parts"`
}

type delivery struct {
	Offset int64
	Len    int64
	Part   int
}

type Uploader struct {
	cl        *client.Client
	wg        sync.WaitGroup
	factoryID string
	log       *log.Logger
}

func New(cl *client.Client, factoryID string) (*Uploader, error) {
	if cl == nil {
		return nil, errors.New("uploader: client cannot be nil")
	}
	return &Uploader{
		cl:        cl,
		factoryID: factoryID,
		log:       log.New(ioutil.Discard, "", 0),
	}, nil
}

func (u *Uploader) UploadSession(s *Session, r io.ReaderAt, size int64) error {
	return u.upload(s, r, size)
}

func (u *Uploader) Upload(r io.ReaderAt, filename string, size int64,
	profiles []string) error {
	s, err := u.NewSession(filename, size, profiles)
	if err != nil {
		return err
	}
	return u.upload(s, r, size)
}

func (u *Uploader) upload(s *Session, r io.ReaderAt, size int64) (err error) {
	u.log.Printf("parts=%d, part_size=%d bytes, max_connections=%d\n", s.Parts,
		s.PartSize, s.MaxConnections)

	n := min(s.MaxConnections, s.Parts)
	deliverych := make(chan delivery)
	stopch := make(chan struct{})
	errch := make(chan error, n)
	partCounter := new(int32)

	u.startWorkers(s, r, n, deliverych, stopch, errch, partCounter)
	if err := u.deliverData(s, size, deliverych, stopch); err != nil {
		return err
	}

	done := make(chan struct{})
	go func() { u.wg.Wait(); close(done) }()

	var once sync.Once
loop:
	for {
		select {
		case <-done:
			break loop
		case e := <-errch:
			if e != nil {
				once.Do(func() {
					close(stopch)
					err = e
				})
			}
		}
	}

	defer u.log.Printf("uploaded %d/%d parts\n", atomic.LoadInt32(partCounter),
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

func (u *Uploader) SetLogger(l *log.Logger) {
	u.log = l
}

func (u *Uploader) startWorkers(s *Session, r io.ReaderAt, n int,
	deliverych chan delivery, stopch chan struct{}, errch chan<- error,
	partCounter *int32) {
	u.log.Printf("starting %d workers\n", n)
	for i := 0; i < n; i++ {
		u.wg.Add(1)
		go func(i int) {
			defer u.wg.Done()
			w := &worker{
				id:          i,
				r:           r,
				location:    s.Location,
				deliverych:  deliverych,
				errch:       errch,
				retryLimit:  5,
				partCounter: partCounter,
				log:         u.log,
			}
			w.start()
		}(i)
	}
}

func (u *Uploader) deliverData(s *Session, fsize int64,
	deliverych chan<- delivery, stopch <-chan struct{}) error {
	status, err := u.uploadStatus(s.Location)
	if err != nil {
		return err
	}
	mp := make(map[int]struct{})
	for _, p := range status.MissingParts {
		mp[p] = struct{}{}
	}

	u.wg.Add(1)
	go func() {
		defer u.wg.Done()
		defer u.log.Println("done sending data from file")
		defer close(deliverych)
		for i := 0; i < s.Parts; i, fsize = i+1, fsize-s.PartSize {
			select {
			case <-stopch:
				return
			default:
				if _, ok := mp[i]; !ok {
					continue
				}
				deliverych <- delivery{
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

func (u *Uploader) NewSession(fname string, fsize int64,
	profiles []string) (*Session, error) {
	params := url.Values{
		"cloud_id":  {u.factoryID},
		"file_size": {strconv.FormatInt(fsize, 10)},
		"file_name": {fname},
		"profiles":  profiles,
	}
	u.cl.SignParams("POST", "/videos/upload.json", params)
	req, err := http.NewRequest("POST",
		"http://"+u.cl.Host+"/v3.0/videos/upload.json?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Multi-Chunk", "1")
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
