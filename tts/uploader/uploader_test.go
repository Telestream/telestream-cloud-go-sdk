package uploader

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func stripHTTP(url string) string {
	return strings.Replace(url, "http://", "", 1)
}

func newClient(host string) *client.Client {
	return &client.Client{
		Host: host,
		Options: &client.Options{
			AccessKey: "123",
			Namespace: "v3.0",
		},
		HTTPClient: nil,
	}
}

func runUploader(t *testing.T, input []byte, host string) error {
	uploader, err := New(newClient(host), "123")
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	uploader.retryDelay = time.Millisecond * 50
	// uploader.DebugLog = log.New(os.Stdout, "", 0)

	errch := make(chan error)
	go func() {
		errch <- uploader.Upload(bytes.NewReader(input), "f", 10, nil, nil)
	}()

	select {
	case err := <-errch:
		return err
	case <-time.After(time.Second * 5):
		t.Fatalf("timeout!")
	}
	return nil
}

func TestUploaderSession(t *testing.T) {
	expected := &Session{
		Parts:          10,
		MaxConnections: 2,
		PartSize:       2,
	}
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/v3.1/videos/upload.json" && r.Method == "POST" {
				w.Write(mustMarshal(expected))
				return
			}
			http.Error(w, "not found", 404)
		}))
	defer ts.Close()

	uploader, err := New(newClient(stripHTTP(ts.URL)), "123")
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	session, err := uploader.NewSession("fname", 10, []string{"h264"}, nil)
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	if !reflect.DeepEqual(expected, session) {
		t.Errorf("expected %v; got %v", expected, session)
	}
}

func TestUploader(t *testing.T) {
	session := &Session{
		Parts:          10,
		MaxConnections: 4,
		PartSize:       1,
	}
	serverData := struct {
		sync.Mutex
		data         []string
		missingParts map[int]struct{}
	}{missingParts: make(map[int]struct{})}
	// at first all parts are missing.
	for i := 0; i < session.Parts; i++ {
		serverData.missingParts[i] = struct{}{}
	}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			switch {
			case r.URL.Path == "/api/v3.1/videos/upload.json" &&
				r.Method == "POST":
				w.Write(mustMarshal(session))
				return
			case r.URL.Path == "/upload":
				switch r.Method {
				case "GET":
					serverData.Lock()
					var parts []int
					for k := range serverData.missingParts {
						parts = append(parts, k)
					}
					serverData.Unlock()
					sort.Ints(parts)
					w.Write(mustMarshal(&Status{parts}))
					return
				case "PUT":
					b, err := ioutil.ReadAll(r.Body)
					if err != nil {
						http.Error(w, err.Error(), 500)
						return
					}
					part, err := strconv.Atoi(r.Header.Get("X-Part"))
					if err != nil {
						http.Error(w, err.Error(), 500)
						return
					}
					serverData.Lock()
					delete(serverData.missingParts, part)
					serverData.data = append(serverData.data, string(b))
					serverData.Unlock()
					return
				}
			}
			http.Error(w, "not found", 404)
		}))
	session.Location = ts.URL + "/upload"
	defer ts.Close()

	input := "0123456789"
	if err := runUploader(t, []byte(input), stripHTTP(ts.URL)); err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}

	if l := len(serverData.data); l != session.Parts {
		t.Errorf("expected len=%d; got %d", session.Parts, l)
	}
	sort.Strings(serverData.data)
	if out := strings.Join(serverData.data, ""); !reflect.DeepEqual(input, out) {
		t.Errorf("expected %v; got %v", input, out)
	}
}

// uploader should retry every request few times in case of fail.
func TestUploaderFail(t *testing.T) {
	session := &Session{
		Parts:          4,
		MaxConnections: 2,
		PartSize:       1,
	}
	serverData := struct {
		sync.Mutex
		partRequestCounter map[int]int
		partsStatusMap     map[int]bool
	}{
		partRequestCounter: make(map[int]int),
		partsStatusMap: map[int]bool{
			0: true,
			1: true,
			2: false,
			3: true,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			switch {
			case r.URL.Path == "/v3.0/videos/upload.json" &&
				r.Method == "POST":
				w.Write(mustMarshal(session))
				return
			case r.URL.Path == "/upload":
				switch r.Method {
				case "GET":
					var parts = []int{0, 1, 2, 3}
					w.Write(mustMarshal(&Status{parts}))
					return
				case "PUT":
					part, err := strconv.Atoi(r.Header.Get("X-Part"))
					if err != nil {
						http.Error(w, err.Error(), 500)
						return
					}
					serverData.Lock()
					serverData.partRequestCounter[part]++
					serverData.Unlock()
					if !serverData.partsStatusMap[part] {
						http.Error(w, "error", 500)
					}
					return
				}
			}
			http.Error(w, "not found", 404)
		}))
	session.Location = ts.URL + "/upload"
	defer ts.Close()

	if err := runUploader(t, []byte("0123"), stripHTTP(ts.URL)); err == nil {
		t.Fatal("expected err!=nil")
	}

	expected := map[int]int{
		0: 1,
		1: 1,
		2: 5,
		3: 1,
	}
	if !reflect.DeepEqual(expected, serverData.partRequestCounter) {
		t.Errorf("expected %v, got %v", expected, serverData.partRequestCounter)
	}
}
