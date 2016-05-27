package uploader

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
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

func TestUploaderSession(t *testing.T) {
	expected := &Session{
		Parts:          10,
		MaxConnections: 2,
		PartSize:       2,
	}
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v3.0/videos/upload.json" && r.Method == "POST" {
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
	session, err := uploader.NewSession("fname", 10, []string{"h264"})
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
			case r.URL.Path == "/v3.0/videos/upload.json" &&
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

	uploader, err := New(newClient(stripHTTP(ts.URL)), "123")
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	uploader.DebugLog = log.New(os.Stdout, "", 0)

	input := "0123456789"

	errch := make(chan error)
	go func() {
		errch <- uploader.Upload(bytes.NewReader([]byte(input)), "f", 10, nil)
	}()
	select {
	case err := <-errch:
		if err != nil {
			t.Fatalf("expected err=nil; got %v", err)
		}
	case <-time.After(time.Second * 5):
		t.Fatalf("timeout!")
	}

	if l := len(serverData.data); l != session.Parts {
		t.Errorf("expected len=%d; got %d", session.Parts, l)
	}
	sort.Strings(serverData.data)
	if out := strings.Join(serverData.data, ""); !reflect.DeepEqual(input, out) {
		t.Errorf("expected %v; got %v", input, out)
	}
}
