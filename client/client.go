// Package client provides a client capable of communicating with Telestream
// Cloud service.
package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// HostGCE is the default api host in Telestream Cloud.
const HostGCE = "api.cloud.telestream.net:443"

// ClientOptions hold credentials required for authenticating requests to the
// Telestream Cloud.
type Options struct {
	AccessKey string
	SecretKey string
	Namespace string
}

var queryFixer = strings.NewReplacer("+", "%20", "%5B", "[", "%5D", "]",
	"%7E", "~")

// Client is capable of sending signed requests to the Telestream Cloud.
type Client struct {
	Host       string
	Options    *Options
	HTTPClient *http.Client
}

func (c *Client) hostPort() string {
	if c.Host != "" {
		return c.Host
	}
	return HostGCE
}

func (c *Client) host() string {
	hp := c.hostPort()
	i := strings.Index(hp, ":")
	if i < 0 {
		return hp
	}
	return hp[:i]
}

func (c *Client) namespace() string {
	if c.Options.Namespace == "" {
		return "v3.0"
	}
	return c.Options.Namespace
}

func (c *Client) httpclient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}

func (c *Client) addAuthParams(v url.Values, t time.Time) {
	v.Set("access_key", c.Options.AccessKey)
	v.Set("timestamp", t.Format(time.RFC3339Nano))
}

func (c *Client) fixQuery(s string) string {
	return queryFixer.Replace(s)
}

func (c *Client) buildSignature(v url.Values, method,
	u string) (sign string, err error) {
	toSign := fmt.Sprintf("%s\n%s\n%s\n%s", method, c.host(), u,
		c.fixQuery(v.Encode()))
	mac := hmac.New(sha256.New, []byte(c.Options.SecretKey))
	if _, err = mac.Write([]byte(toSign)); err == nil {
		sign = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	}
	return
}

func (c *Client) buildURL(v url.Values, urlPath string) *url.URL {
	scheme := "http"
	if strings.HasSuffix(c.hostPort(), ":443") {
		scheme = "https"
	}
	return &url.URL{
		Scheme:   scheme,
		Host:     c.hostPort(),
		Path:     path.Join(c.namespace(), urlPath),
		RawQuery: v.Encode(),
	}
}

func (c *Client) do(method, path, cntType string,
	params url.Values, r io.Reader) (b []byte, err error) {
	if params == nil {
		params = url.Values{}
	}
	if err = c.SignParams(method, path, params); err != nil {
		return
	}
	req, err := http.NewRequest(method, c.buildURL(params, path).String(), r)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", cntType)
	resp, err := c.httpclient().Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		e := &Error{Code: resp.StatusCode}
		if err = json.Unmarshal(b, e); err != nil {
			return nil, e
		}
		err = e
	}
	return
}

// SignParams signs given parameters by adding required authorization
// fields and values. Params cannot be nil.
func (c *Client) SignParams(method, path string, params url.Values) error {
	if params == nil {
		panic("params cannot be nil!")
	}
	c.addAuthParams(params, time.Now().UTC())
	s, err := c.buildSignature(params, method, path)
	if err != nil {
		return err
	}
	params.Set("signature", s)
	return nil
}

// Get issues a signed GET request to the Panda Cloud
func (c *Client) Get(url string, params url.Values) ([]byte, error) {
	return c.do("GET", url, "", params, nil)
}

// Post issues a signed POST request to the Panda Cloud and creates content
// based on the given params
func (c *Client) Post(url, cntType string, params url.Values,
	r io.Reader) ([]byte, error) {
	return c.do("POST", url, cntType, params, r)
}

// Put issues a signed PUT request to the Panda Cloud and updates object
// according to given params
func (c *Client) Put(url, cntType string, params url.Values,
	r io.Reader) ([]byte, error) {
	return c.do("PUT", url, cntType, params, r)
}

// Delete issues a signed DELETE request to the Panda Cloud and deletes content
// under the given url
func (c *Client) Delete(url string, params url.Values) ([]byte, error) {
	return c.do("DELETE", url, "", params, nil)
}
