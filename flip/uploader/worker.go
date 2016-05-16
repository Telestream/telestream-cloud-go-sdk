package uploader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

type worker struct {
	id          int
	r           io.ReaderAt
	location    string
	deliverych  <-chan delivery
	errch       chan<- error
	retryLimit  int
	partCounter *int32
	log         *log.Logger
}

func (w *worker) start() error {
	w.log.Printf("worker %d: started\n", w.id)
	defer w.log.Printf("worker %d: bye\n", w.id)
	for del := range w.deliverych {
		w.log.Printf("worker %d: uploading part %d\n", w.id, del.Part)
		if err := w.tryHandle(&del, w.retryLimit); err != nil {
			w.errch <- err
		}
	}
	return nil
}

func (w *worker) handle(d *delivery) error {
	req, err := http.NewRequest("PUT", w.location,
		io.NewSectionReader(w.r, d.Offset, d.Len))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("X-Part", strconv.Itoa(d.Part))
	req.ContentLength = d.Len

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return err
	}
	atomic.AddInt32(w.partCounter, 1)
	return nil
}

func (w *worker) tryHandle(del *delivery, retryLimit int) (err error) {
	for retry := 1; ; retry++ {
		if err = w.handle(del); err != nil {
			if retry < w.retryLimit {
				w.log.Printf("worker %d: error while uploading part %d; %v; "+
					"retrying %d/%d\n", w.id, del.Part, err, retry, w.retryLimit)
				time.Sleep(time.Second)
				continue
			} else {
				w.log.Printf("worker %d: failed to upload part %d\n", w.id,
					del.Part)
				return fmt.Errorf("failed to upload part %d: %v", del.Part, err)
			}
		}
		return nil
	}
}
