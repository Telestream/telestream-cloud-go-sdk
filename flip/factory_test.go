package flip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
)

const namespace = "v3.0"

var testData = struct {
	Factories     []Factory
	Videos        []Video
	Encodings     []Encoding
	Profiles      []Profile
	Notifications []Notification
}{
	Factories: []Factory{
		{
			CreatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			ID:              "123",
			Name:            "FactoryOne",
			S3PrivateAccess: true,
			S3VideosBucket:  "FlipBucket",
			URL:             "",
			UpdatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
		{
			CreatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			ID:              "456",
			Name:            "FactoryTwo",
			S3PrivateAccess: true,
			S3VideosBucket:  "FlipBucket",
			URL:             "",
			UpdatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
	},
	Videos: []Video{
		{
			ID:              "456",
			Fps:             29.75,
			Width:           300,
			Height:          240,
			Duration:        14000,
			Status:          StatusSuccess,
			VideoBitrate:    344,
			AudioSampleRate: 44100,
			CreatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
		{
			ID:              "789",
			Fps:             29.75,
			Width:           1920,
			Height:          1080,
			Duration:        14000,
			Status:          StatusSuccess,
			VideoBitrate:    344,
			AudioSampleRate: 44100,
			CreatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:       Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
	},
	Encodings: []Encoding{
		{
			ID:                "123",
			Fps:               29.75,
			Width:             300,
			Height:            240,
			Duration:          14000,
			Status:            StatusSuccess,
			VideoBitrate:      344,
			AudioSampleRate:   44100,
			CreatedAt:         Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:         Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			StartedEncodingAt: Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
		{
			ID:                "789",
			Fps:               29.75,
			Width:             1920,
			Height:            1080,
			Duration:          14000,
			Status:            StatusSuccess,
			VideoBitrate:      344,
			AudioSampleRate:   44100,
			CreatedAt:         Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:         Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			StartedEncodingAt: Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
	},
	Profiles: []Profile{
		{
			ID:         "789",
			PresetName: "h264",
			Name:       "Profile1",
			AspectMode: ModeConstrain,
			CreatedAt:  Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:  Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
		{
			ID:         "456",
			PresetName: "h264",
			Name:       "Profile2",
			AspectMode: ModePreserve,
			CreatedAt:  Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
			UpdatedAt:  Time{time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)},
		},
	},
	Notifications: []Notification{
		{
			Delay: 10.0,
			Events: Events{
				EncodingCompleted: true,
				EncodingProgress:  false,
				VideoCreated:      false,
				VideoEncoded:      true,
			},
			URL: "www.example.com",
		},
	},
}

var errWrongPath = func(path string) *client.Error {
	return &client.Error{
		Code:    404,
		Message: fmt.Sprintf("got wrong path: %s", path),
	}
}

var factoryHandler = func(factoryID, method string,
	handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if id := r.FormValue("factory_id"); id != factoryID {
			err := client.Error{
				Code:    404,
				Message: fmt.Sprintf("invalid factory %s", id),
			}
			http.Error(w, string(mustMarshal(err)), err.Code)
			return
		}
		if r.Method != method {
			err := client.Error{
				Code:    400,
				Message: fmt.Sprintf("invalid method %s", r.Method),
			}
			http.Error(w, string(mustMarshal(err)), err.Code)
			return
		}
		handle(w, r)
	}
}

var listHandler = func(v interface{}, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slice := reflect.ValueOf(v)
		if slice.Kind() != reflect.Slice {
			err := &client.Error{
				Code:    500,
				Err:     "listHandler",
				Message: "v is not slice",
			}
			http.Error(w, string(mustMarshal(err)), err.Code)
			return
		}
		if r.URL.Path == path {
			n, err := strconv.Atoi(r.FormValue("per_page"))
			if err != nil {
				n = slice.Len()
			}
			w.Write(mustMarshal(slice.Slice(0, n).Interface()))
			return
		}
		err := errWrongPath(r.URL.Path)
		http.Error(w, string(mustMarshal(err)), err.Code)
	}
}

var videoReaderHandler = func(content []byte,
	filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpError := func(w http.ResponseWriter, code int, msg string) {
			e := client.Error{Code: code, Message: msg}
			http.Error(w, string(mustMarshal(e)), e.Code)
		}
		if r.URL.Path == urlPath("videos.json") {
			if err := r.ParseMultipartForm(1024); err != nil {
				httpError(w, 500, err.Error())
				return
			}
			if len(r.MultipartForm.File["file"]) != 1 {
				httpError(w, 500, "len should be 1")
				return
			}
			if name := r.MultipartForm.File["file"][0].
				Filename; name != filename {
				httpError(w, 500, "filename mismatch: "+name)
				return
			}
			f, err := r.MultipartForm.File["file"][0].Open()
			if err != nil {
				httpError(w, 500, err.Error())
				return
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				httpError(w, 500, err.Error())
				return
			}
			if !bytes.Equal(b, content) {
				httpError(w, 500, err.Error())
				return
			}
			w.Write([]byte("{}"))
			return
		}
		err := errWrongPath(r.URL.Path)
		http.Error(w, string(mustMarshal(err)), err.Code)
	}
}

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func urlPath(s ...string) string {
	return path.Join(append([]string{"/" + namespace}, s...)...)
}

func newFactory(host, id string) *Factory {
	return &Factory{
		client: newClient(stripHTTP(host)),
		ID:     id,
	}
}

func stripHTTP(url string) string {
	return strings.Replace(url, "http://", "", 1)
}

func TestFactoryVideo(t *testing.T) {
	const factoryID = "123"
	testVideo := testData.Videos[0]
	ts := httptest.NewServer(factoryHandler(factoryID, "GET", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("videos", testVideo.ID+".json") {
				w.Write(mustMarshal(&testVideo))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	v, err := newFactory(ts.URL, factoryID).Video(testVideo.ID)
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	v.CreatedAt.Time = v.CreatedAt.UTC()
	v.UpdatedAt.Time = v.UpdatedAt.UTC()
	if !reflect.DeepEqual(&testVideo, v) {
		t.Errorf("expected v=%#v; got %#v", testVideo, v)
	}
}

func TestFactoryVideos(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		listHandler(testData.Videos, urlPath("videos.json"))))
	defer ts.Close()
	fixtures := []struct {
		VideoRequest *VideoRequest
		Expected     []Video
	}{
		{
			&VideoRequest{PerPage: 1},
			testData.Videos[:1],
		},
		{
			nil,
			testData.Videos,
		},
	}
	for i, fixture := range fixtures {
		videos, err := newFactory(ts.URL, factoryID).Videos(fixture.VideoRequest)
		if err != nil {
			t.Fatalf("expected err=nil; got %v (i=%d)", err, i)
		}
		if len(videos) != len(fixture.Expected) {
			t.Fatalf("expected to get %d videos, got %d (i=%d)", i,
				len(fixture.Expected), len(videos))
		}
		for j, video := range videos {
			testVideo := fixture.Expected[j]
			video.CreatedAt.Time = video.CreatedAt.UTC()
			video.UpdatedAt.Time = video.UpdatedAt.UTC()
			if !reflect.DeepEqual(&video, &testVideo) {
				t.Errorf("expected factory=%v; got %v (i=%d)", testVideo, video, i)
			}
		}
	}
}

func TestFactoryVideoDeleteSource(t *testing.T) {
	const factoryID = "123"
	const videoID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "DELETE",
		http.HandlerFunc(func(w http.ResponseWriter,
			r *http.Request) {
			if r.URL.Path == urlPath("videos", videoID, "source.json") {
				w.Write([]byte("{\"deleted\":true}"))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	if err := newFactory(ts.URL, factoryID).
		VideoDeleteSource(videoID); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryVideoEncodings(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		listHandler(testData.Encodings,
			urlPath("videos", "123", "encodings.json"))))
	defer ts.Close()
	fixtures := []struct {
		EncodingRequest *EncodingRequest
		Expected        []Encoding
	}{
		{
			&EncodingRequest{PerPage: 1},
			testData.Encodings[:1],
		},
		{
			nil,
			testData.Encodings,
		},
	}
	for i, fixture := range fixtures {
		encodings, err := newFactory(ts.URL, factoryID).VideoEncodings("123",
			fixture.EncodingRequest)
		if err != nil {
			t.Fatalf("expected err=nil; got %v (i=%d)", err, i)
		}
		if len(encodings) != len(fixture.Expected) {
			t.Fatalf("expected to get %d encodings, got %d (i=%d)",
				len(fixture.Expected), len(encodings), i)
		}
		for j, encoding := range encodings {
			testEncoding := fixture.Expected[j]
			encoding.CreatedAt.Time = encoding.CreatedAt.UTC()
			encoding.UpdatedAt.Time = encoding.UpdatedAt.UTC()
			encoding.StartedEncodingAt.Time = encoding.StartedEncodingAt.UTC()
			if !reflect.DeepEqual(&encoding, &testEncoding) {
				t.Errorf("expected factory=%v; got %v (i=%d)",
					testEncoding, encoding, i)
			}
		}
	}
}

func TestFactoryEncoding(t *testing.T) {
	testEncoding := testData.Encodings[0]
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "GET", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("encodings", testEncoding.ID+".json") {
				w.Write(mustMarshal(&testEncoding))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	factory := &Factory{
		client: newClient(stripHTTP(ts.URL)),
		ID:     factoryID,
	}
	e, err := factory.Encoding(testEncoding.ID)
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
	e.CreatedAt.Time = e.CreatedAt.UTC()
	e.UpdatedAt.Time = e.UpdatedAt.UTC()
	e.StartedEncodingAt.Time = e.StartedEncodingAt.UTC()
	if !reflect.DeepEqual(&testEncoding, e) {
		t.Errorf("expected v=%#v; got %#v", testEncoding, e)
	}
}

func TestFactoryEncodings(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		listHandler(testData.Encodings, urlPath("encodings.json"))))
	defer ts.Close()
	fixtures := []struct {
		EncodingRequest *EncodingRequest
		Expected        []Encoding
	}{
		{
			&EncodingRequest{PerPage: 1},
			testData.Encodings[:1],
		},
		{
			nil,
			testData.Encodings,
		},
	}
	for i, fixture := range fixtures {
		encodings, err := newFactory(ts.URL, factoryID).
			Encodings(fixture.EncodingRequest)
		if err != nil {
			t.Fatalf("expected err=nil; got %v (i=%d)", err, i)
		}
		if len(encodings) != len(fixture.Expected) {
			t.Fatalf("expected to get %d encodings, got %d (i=%d)", i,
				len(fixture.Expected), len(encodings))
		}
		for j, encoding := range encodings {
			testEncoding := fixture.Expected[j]
			encoding.CreatedAt.Time = encoding.CreatedAt.UTC()
			encoding.UpdatedAt.Time = encoding.UpdatedAt.UTC()
			encoding.StartedEncodingAt.Time = encoding.StartedEncodingAt.UTC()
			if !reflect.DeepEqual(&encoding, &testEncoding) {
				t.Errorf("expected factory=%v; got %v (i=%d)",
					testEncoding, encoding, i)
			}
		}
	}
}

func TestNewEncoding(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("profile_name") == "Profile1" &&
				r.URL.Path == urlPath("encodings.json") {
				w.Write([]byte("{}"))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	_, err := newFactory(ts.URL, factoryID).NewEncoding(&NewEncodingRequest{
		ProfileName: "Profile1",
	})
	if err != nil {
		t.Fatalf("expected err=nil; got %v", err)
	}
}

func TestFactoryCancelRetryEncoding(t *testing.T) {
	const factoryID = "123"
	const encodingID = "456"
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case urlPath("encodings", encodingID, "cancel.json"):
				w.Write([]byte("{\"canceled\":true}"))
				return
			case urlPath("encodings", encodingID, "retry.json"):
				w.Write([]byte("{\"retried\":true}"))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	if err := newFactory(ts.URL, factoryID).
		CancelEncoding(encodingID); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
	if err := newFactory(ts.URL, factoryID).
		RetryEncoding(encodingID); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryProfile(t *testing.T) {
	const factoryID = "123"
	testProfile := testData.Profiles[1]
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("profiles", testProfile.ID+".json") {
				w.Write(mustMarshal(testProfile))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	p, err := newFactory(ts.URL, factoryID).Profile(testProfile.ID)
	if err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
	p.CreatedAt.Time = p.CreatedAt.UTC()
	p.UpdatedAt.Time = p.UpdatedAt.UTC()
	if !reflect.DeepEqual(p, &testProfile) {
		t.Errorf("expected %v; got %v", &testProfile, p)
	}
}

func TestFactoryProfiles(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		listHandler(testData.Profiles, urlPath("profiles.json"))))
	defer ts.Close()
	fixtures := []struct {
		ProfileRequest *ProfileRequest
		Expected       []Profile
	}{
		{
			&ProfileRequest{PerPage: 1},
			testData.Profiles[:1],
		},
		{
			nil,
			testData.Profiles,
		},
	}
	for i, fixture := range fixtures {
		profiles, err := newFactory(ts.URL, factoryID).
			Profiles(fixture.ProfileRequest)
		if err != nil {
			t.Fatalf("expected err=nil; got %v (i=%d)", err, i)
		}
		if len(profiles) != len(fixture.Expected) {
			t.Fatalf("expected to get %d profiles, got %d (i=%d)", i,
				len(fixture.Expected), len(profiles))
		}
		for j, profile := range profiles {
			testProfile := fixture.Expected[j]
			profile.CreatedAt.Time = profile.CreatedAt.UTC()
			profile.UpdatedAt.Time = profile.UpdatedAt.UTC()
			if !reflect.DeepEqual(&profile, &testProfile) {
				t.Errorf("expected factory=%v; got %v (i=%d)",
					testProfile, profile, i)
			}
		}
	}
}

func TestFactoryUpdate(t *testing.T) {
	const factoryID = "123"
	testProfile := testData.Profiles[0]
	expectedTime := testProfile.UpdatedAt.Add(time.Minute)
	testNotification := testData.Notifications[0]
	ts := httptest.NewServer(factoryHandler(factoryID, "PUT",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case urlPath("profiles", testProfile.ID+".json"):
				p := testProfile
				p.UpdatedAt.Time = expectedTime
				w.Write(mustMarshal(p))
				return
			case urlPath("notifications.json"):
				// check default values
				if r.FormValue("events[encoding_progress]") != "false" ||
					r.FormValue("events[video_created]") != "false" {
					err := client.Error{
						Code:    500,
						Message: "default values were not provided",
					}
					http.Error(w, string(mustMarshal(err)), err.Code)
					return
				}
				w.Write(mustMarshal(testNotification))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	for i, updater := range []UpdatableResource{
		&testProfile,
		&testNotification,
	} {
		if err := newFactory(ts.URL, factoryID).Update(updater); err != nil {
			t.Errorf("expected err=nil; got %v (i=%d)", err, i)
		}
	}
	if !testProfile.UpdatedAt.Equal(expectedTime) {
		t.Errorf("expected %v; got %v", expectedTime, testProfile)
	}
}

func TestFactoryDelete(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "DELETE",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case urlPath("profiles", "789.json"),
				urlPath("videos", "456.json"),
				urlPath("encodings", "123.json"):
				w.Write([]byte(`{"deleted":true}`))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	for i, deleter := range []DeletableResource{
		&testData.Profiles[0],
		&testData.Encodings[0],
		&testData.Videos[0],
	} {
		if err := newFactory(ts.URL, factoryID).Delete(deleter); err != nil {
			t.Errorf("expected err=nil; got %v (i=%d)", err, i)
		}
	}
}

func TestFactoryNotifications(t *testing.T) {
	const factoryID = "123"
	testNotification := testData.Notifications[0]
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("notifications.json") {
				w.Write(mustMarshal(testNotification))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	notification, err := newFactory(ts.URL, factoryID).Notifications()
	if err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
	if !reflect.DeepEqual(notification, &testNotification) {
		t.Errorf("expectec notification=%v; got %v", testNotification, notification)
	}
}

func TestFactoryNewVideoURL(t *testing.T) {
	const factoryID = "123"
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("videos.json") {
				w.Write([]byte("{}"))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	if _, err := newFactory(ts.URL, factoryID).NewVideoURL(
		"www.example.com", nil); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryNewVideoURLArg(t *testing.T) {
	const factoryID = "123"
	testRequest := &NewVideoRequest{
		PathFormat: ":id",
		Payload:    "payload",
		Profiles:   []string{"Profile1", "Profile2"},
	}
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("videos.json") {
				if r.FormValue("path_format") != testRequest.PathFormat ||
					r.FormValue("payload") != testRequest.Payload ||
					r.FormValue("profiles") != strings.Join(testRequest.Profiles, ",") {
					err := client.Error{
						Code:    500,
						Message: "params were not succesfully provided",
					}
					http.Error(w, string(mustMarshal(err)), err.Code)
					return
				}
				w.Write([]byte("{}"))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	if _, err := newFactory(ts.URL, factoryID).
		NewVideoURL("www.example.com", testRequest); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryVideoReader(t *testing.T) {
	const factoryID = "123"
	testContent := []byte("test")
	testFilename := "name"
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		videoReaderHandler(testContent, testFilename)))
	defer ts.Close()
	r := bytes.NewReader(testContent)
	if _, err := newFactory(ts.URL, factoryID).
		NewVideoReader(r, testFilename, nil); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryVideoReaderArg(t *testing.T) {
	const factoryID = "123"
	testContent := []byte("test")
	testFilename := "name"
	testRequest := &NewVideoRequest{
		PathFormat: ":id",
		Payload:    "payload",
		Profiles:   []string{"Profile1", "Profile2"},
	}
	ts := httptest.NewServer(factoryHandler(factoryID, "POST",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("path_format") != testRequest.PathFormat ||
				r.FormValue("payload") != testRequest.Payload ||
				r.FormValue("profiles") != strings.Join(testRequest.Profiles, ",") {
				err := client.Error{
					Code:    500,
					Message: "params were not succesfully provided",
				}
				http.Error(w, string(mustMarshal(err)), err.Code)
				return
			}
			videoReaderHandler(testContent, testFilename)(w, r)
		})))
	defer ts.Close()
	r := bytes.NewReader(testContent)
	if _, err := newFactory(ts.URL, factoryID).
		NewVideoReader(r, testFilename, testRequest); err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
}

func TestFactoryVideoMetaData(t *testing.T) {
	const factoryID = "123"
	expected := map[string]interface{}{"test": true}
	ts := httptest.NewServer(factoryHandler(factoryID, "GET",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == urlPath("videos", "123", "metadata.json") {
				w.Write(mustMarshal(expected))
				return
			}
			err := errWrongPath(r.URL.Path)
			http.Error(w, string(mustMarshal(err)), err.Code)
		})))
	defer ts.Close()
	meta, err := newFactory(ts.URL, factoryID).VideoMetaData("123")
	if err != nil {
		t.Errorf("expected err=nil; got %v", err)
	}
	if v, ok := meta["test"]; !ok || !reflect.DeepEqual(v, true) {
		t.Errorf("expected %v; got %v", expected, meta)
	}
}
