// Package flip provides a client for Telestream Cloud transcoding service.
package flip

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
)

// Flip wraps Telestream Client. Flip client is capable of listing all
// available Factories and passing them it's own telestream client.
type Flip struct {
	client *client.Client
}

// NewFlip builds flip client using the default api host. If HttpClient is nil,
// then the http.DefaultClient will be used.
func NewFlip(accessKey, secretKey string, httpClient *http.Client) *Flip {
	return newFlip(client.HostGCE, accessKey, secretKey, httpClient)
}

// NewFlipHost builds flip client based on the given parameters. If HttpClient
// is nil, then the http.DefaultClient will be used.
func NewFlipHost(host, accessKey, secretKey string,
	httpClient *http.Client) *Flip {
	return newFlip(host, accessKey, secretKey, httpClient)
}

func newFlip(host, accessKey, secretKey string, httpClient *http.Client) *Flip {
	return &Flip{
		client: &client.Client{
			Host:       host,
			HTTPClient: httpClient,
			Options: &client.Options{
				AccessKey: accessKey,
				SecretKey: secretKey,
				Namespace: "v3.0",
			},
		},
	}
}

// Factory gets the factory with the given ID.
func (f *Flip) Factory(id string) (*Factory, error) {
	b, err := f.client.Get(fmt.Sprintf(factoriesIDPath, id), nil)
	if err != nil {
		return nil, err
	}
	fac := new(Factory)
	if err := json.Unmarshal(b, fac); err != nil {
		return nil, err
	}
	fac.client = f.client
	return fac, nil
}

// Factories gets all factories on the given account.
func (f *Flip) Factories() (factories []Factory, err error) {
	b, err := f.client.Get(factoriesPath, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &factories); err != nil {
		return nil, err
	}
	for i := range factories {
		factories[i].client = f.client
	}
	return factories, nil
}
