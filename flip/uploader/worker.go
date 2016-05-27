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
	retryDelay  time.Duration
	partCounter *int32
	log         *log.Logger
}

func (w *worker) start() {
	w.logf("started\n")
	defer w.logf("bye\n")
	for del := range w.deliverych {
		w.logf("uploading part %d\n", del.Part)
		if err := w.tryHandle(&del, w.retryLimit); err != nil {
			w.errch <- err
		}
	}
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
		return fmt.Errorf("error: status code: %d", resp.StatusCode)
	}
	atomic.AddInt32(w.partCounter, 1)
	return nil
}

func (w *worker) tryHandle(del *delivery, retryLimit int) (err error) {
	for retry := 1; ; retry++ {
		if err = w.handle(del); err != nil {
			if retry < w.retryLimit {
				w.logf("error while uploading part %d; %v; "+
					"retrying %d/%d\n", del.Part, err, retry, w.retryLimit)
				time.Sleep(w.retryDelay)
				continue
			} else {
				w.logf("failed to upload part %d\n", del.Part)
				return fmt.Errorf("failed to upload part %d: %v", del.Part, err)
			}
		}
		return nil
	}
}

func (w *worker) logf(format string, args ...interface{}) {
	if w.log != nil {
		w.log.Printf(fmt.Sprintf("worker %d: %s", w.id, format), args...)
	}
}
