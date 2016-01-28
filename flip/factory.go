package flip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"reflect"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
	"github.com/ernesto-jimenez/go-querystring/query"
)

// Factory holds the information about the factory. It is capable of calling
// multiple api actions. The Factory is supposed to be created by a Flip, not
// by a user.
type Factory struct {
	client          *client.Client
	CreatedAt       Time   `json:"created_at,omitempty" url:"created_at,omitempty"`
	ID              string `json:"id,omitempty" url:"id,omitempty"`
	Name            string `json:"name,omitempty" url:"name,omitempty"`
	S3PrivateAccess bool   `json:"s3_private_access,omitempty" url:"s3_private_access,omitempty"`
	S3VideosBucket  string `json:"s3_videos_bucket,omitempty" url:"s3_videos_bucket,omitempty"`
	URL             string `json:"url,omitempty" url:"url,omitempty"`
	UpdatedAt       Time   `json:"updated_at,omitempty" url:"updated_at,omitempty"`
}

func (f *Factory) buildParams(v interface{}) (params url.Values, err error) {
	params = url.Values{}
	if v != nil && !reflect.ValueOf(v).IsNil() {
		if params, err = query.Values(v); err != nil {
			return nil, err
		}
	}
	params.Add("factory_id", f.ID)
	return params, nil
}

func (f *Factory) get(URL string, req, dst interface{}) (err error) {
	params, err := f.buildParams(req)
	if err != nil {
		return err
	}
	b, err := f.client.Get(URL, params)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func (f *Factory) post(URL string, req, v interface{}) (err error) {
	params, err := f.buildParams(req)
	if err != nil {
		return err
	}
	b, err := f.client.Post(URL, "", params, nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (f *Factory) put(URL string, v interface{}) error {
	params, err := f.buildParams(v)
	if err != nil {
		return err
	}
	b, err := f.client.Put(URL, "", params, nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (f *Factory) delete(URL string) error {
	_, err := f.client.Delete(URL, url.Values{"factory_id": []string{f.ID}})
	return err
}

// NewEncoding creates a new encoding for the existing video.
func (f *Factory) NewEncoding(er *NewEncodingRequest) (*Encoding, error) {
	if er == nil {
		return nil, errors.New("NewEncodingRequest cannot be nil")
	}
	e := new(Encoding)
	if err := f.post(encodingsPath, er, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// Encoding gets encoding object with the given id.
func (f *Factory) Encoding(id string) (*Encoding, error) {
	e := new(Encoding)
	if err := f.get(fmt.Sprintf(encodingsIDPath, id), nil, e); err != nil {
		return nil, err
	}
	return e, nil
}

// Encodings gets all encodings from the current factory. Encodings can be
// filtered by options set in the EncodingRequest struct. If EncodingRequest is
// nil defaults are going to be used.
func (f *Factory) Encodings(er *EncodingRequest) (es []Encoding, err error) {
	if err := f.get(encodingsPath, er, &es); err != nil {
		return nil, err
	}
	return es, nil
}

// CancelEncoding cancels the encoding with the given id.
func (f *Factory) CancelEncoding(id string) error {
	return f.post(fmt.Sprintf(encodingsIDCancelPath, id), nil, &struct{}{})
}

// RetryEncoding retries the encoding with the given id.
func (f *Factory) RetryEncoding(id string) error {
	return f.post(fmt.Sprintf(encodingsIDRetryPath, id), nil, &struct{}{})
}

// NewProfile creates new profile based on profile request object.
func (f *Factory) NewProfile(pr *NewProfileRequest) (*Profile, error) {
	if pr == nil {
		return nil, errors.New("NewProfileRequest cannot be nil")
	}
	p := new(Profile)
	if err := f.post(profilesPath, pr, p); err != nil {
		return nil, err
	}
	return p, nil
}

// Profile gets profile with the given ID.
func (f *Factory) Profile(id string) (*Profile, error) {
	p := new(Profile)
	if err := f.get(fmt.Sprintf(profilesIDPath, id), nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

// Profiles gets all profiles from the current factory. if ProfileRequest is nil
// defaults are going to be used.
func (f *Factory) Profiles(pr *ProfileRequest) (ps []Profile, err error) {
	err = f.get(profilesPath, pr, &ps)
	return
}

// Update sends an update api call. It takes the path from the Updater
// interface.
func (f *Factory) Update(v UpdatableResource) error {
	path, err := v.UpdatePath()
	if err != nil {
		return err
	}
	return f.put(path, v)
}

// Delete sends a delete api call. It takes the path from the Deleter interface.
func (f *Factory) Delete(v DeletableResource) error {
	path, err := v.DeletePath()
	if err != nil {
		return err
	}
	return f.delete(path)
}

// NewVideo creates a new video in Telestream cloud. If NewVideoRequest is nil
// defaults for the new video will be used.
func (f *Factory) NewVideo(filename string,
	vr *NewVideoRequest) (*Video, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f.NewVideoReader(file, filename, vr)
}

// NewVideoURL creates a new video from the given source url. If NewVideoRequest
// is nil defaults for the new video will be used.
func (f *Factory) NewVideoURL(URL string, vr *NewVideoRequest) (*Video, error) {
	params, err := f.buildParams(vr)
	if err != nil {
		return nil, err
	}
	params.Set("source_url", URL)
	b, err := f.client.Post(videosPath, "", params, nil)
	if err != nil {
		return nil, err
	}
	v := new(Video)
	if err = json.Unmarshal(b, v); err != nil {
		return nil, err
	}
	return v, nil
}

// NewVideoReader creates a new video in Telestream cloud reading the data from
// the given reader. If NewVideoRequest is nil defaults for the new video will
// be used.
func (f *Factory) NewVideoReader(r io.Reader, name string,
	vr *NewVideoRequest) (*Video, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	if err := w.SetBoundary("--video--"); err != nil {
		return nil, err
	}
	p, err := w.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(p, r); err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	params, err := f.buildParams(vr)
	if err != nil {
		return nil, err
	}
	b, err := f.client.Post(videosPath, w.FormDataContentType(), params, buf)
	if err != nil {
		return nil, err
	}
	v := new(Video)
	if err = json.Unmarshal(b, v); err != nil {
		return nil, err
	}
	return v, nil
}

// Video gets video with the given id.
func (f *Factory) Video(id string) (*Video, error) {
	v := new(Video)
	if err := f.get(fmt.Sprintf(videosIDPath, id), nil, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Videos gets all the videos from the current Factory. Videos can be filtered
// by options set in the VideoRequest struct. If VideoRequest is nil then
// defaults are going to be used.
func (f *Factory) Videos(vr *VideoRequest) (v []Video, err error) {
	err = f.get(videosPath, vr, &v)
	return
}

// VideoEncodings gets all encodings related to the video with the given id.
// Encodings can be filtered by options set in the EncodingRequest struct.
// If EncodingRequest is nil defaults are going to be used.
func (f *Factory) VideoEncodings(id string,
	er *EncodingRequest) (es []Encoding, err error) {
	err = f.get(fmt.Sprintf(videosIDEncodingPath, id), er, &es)
	return
}

// VideoMetaData gets meta data for the video with the given id.
func (f *Factory) VideoMetaData(id string) (MetaData, error) {
	md := make(MetaData)
	if err := f.get(fmt.Sprintf(videosIDMetaDataPath, id), nil, &md); err != nil {
		return nil, err
	}
	return md, nil
}

// VideoDeleteSource deletes the source video for the given video id.
func (f *Factory) VideoDeleteSource(id string) error {
	return f.delete(fmt.Sprintf(videosIDDeleteSourcePath, id))
}

// Notifications gets notifications for the current factory.
func (f *Factory) Notifications() (*Notification, error) {
	n := new(Notification)
	if err := f.get(notificationsPath, nil, n); err != nil {
		return nil, err
	}
	return n, nil
}
