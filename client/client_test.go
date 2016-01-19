package client

import (
	"net/url"
	"testing"
	"time"
)

func TestBuildSignature(t *testing.T) {
	expected := "WH0ogZNRhxwvdxg2ptRGcgPi9UAgJw05IA6F/PlbFbU="
	cl := Client{
		Host: "api.cloud.telestream.net:443",
		Options: &Options{
			AccessKey: "abcdefgh",
			SecretKey: "ijklmnop",
		},
	}
	d := time.Date(2011, 3, 1, 15, 39, 10, 260762000, time.UTC)
	v := url.Values{"factory_id": []string{"123456789"}}
	cl.addAuthParams(v, d)
	sign, err := cl.buildSignature(v, "GET", "/videos.json")
	if err != nil {
		t.Fatalf("want err=nil; got %v", err)
	}
	if sign != expected {
		t.Errorf("want signature=%s; got %s", expected, sign)
	}
}

func TestBuildUrl(t *testing.T) {
	cases := []struct {
		cl  Client
		v   url.Values
		url string
		exp string
	}{
		{
			Client{
				Host:    "localhost:9999",
				Options: &Options{},
			},
			url.Values{},
			"/videos.json",
			"http://localhost:9999/v3.0/videos.json",
		},
		{
			Client{
				Host:    "localhost:80",
				Options: &Options{},
			},
			url.Values{"factory_id": []string{"12345"}},
			"/videos.json",
			"http://localhost:80/v3.0/videos.json?factory_id=12345",
		},
		{
			Client{
				Host:    "localhost",
				Options: &Options{},
			},
			url.Values{"factory_id": []string{"12345"}},
			"/videos.json",
			"http://localhost/v3.0/videos.json?factory_id=12345",
		},
		{
			Client{
				Host:    "localhost:443",
				Options: &Options{},
			},
			url.Values{},
			"/videos.json",
			"https://localhost:443/v3.0/videos.json",
		},
	}
	for i, cas := range cases {
		if res := cas.cl.buildURL(cas.v, cas.url).String(); res != cas.exp {
			t.Errorf("expected uri=%s; got %s (i=%d)", cas.exp, res, i)
		}
	}
}
