package flip

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/telestream/telestream-cloud-go-sdk/client"
)

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

func TestFlipFactories(t *testing.T) {
	testFactories := testData.Factories
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath(factoriesPath) && r.Method == "GET" {
				w.Write(mustMarshal(testFactories))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		}))
	defer ts.Close()
	cl := newClient(stripHTTP(ts.URL))
	flip := Flip{cl}
	fs, err := flip.Factories()
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	if len(fs) != len(testFactories) {
		t.Fatalf("expected to get %d factories, got %d",
			len(testFactories), len(fs))
	}
	for i, factory := range fs {
		testFactory := testFactories[i]
		testFactory.client = cl
		factory.CreatedAt.Time = factory.CreatedAt.UTC()
		factory.UpdatedAt.Time = factory.UpdatedAt.UTC()
		if !reflect.DeepEqual(&factory, &testFactory) {
			t.Errorf("expected factory=%v; got %v", testFactory, factory)
		}
	}
}

func TestFlipFactory(t *testing.T) {
	testFactory := testData.Factories[1]
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("factories", testFactory.ID+".json") &&
				r.Method == "GET" {
				w.Write(mustMarshal(&testFactory))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		}))
	defer ts.Close()
	cl := newClient(stripHTTP(ts.URL))
	flip := Flip{cl}
	factory, err := flip.Factory(testFactory.ID)
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	testFactory.client = cl
	factory.CreatedAt.Time = factory.CreatedAt.UTC()
	factory.UpdatedAt.Time = factory.UpdatedAt.UTC()
	if !reflect.DeepEqual(factory, &testFactory) {
		t.Errorf("expected factory=%v; got %v", testFactory, factory)
	}
}
