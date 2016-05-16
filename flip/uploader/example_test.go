package uploader_test

import (
	"fmt"
	"os"

	"github.com/Telestream/telestream-cloud-go-sdk/client"
	"github.com/Telestream/telestream-cloud-go-sdk/flip/uploader"
)

func ExampleUploader_Upload() {
	f, err := os.Open("video.mp4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cl := client.Client{
		Host: "api.cloud.telestream.net:443",
		Options: &client.Options{
			AccessKey: "access_key",
			SecretKey: "secret_key",
		},
	}

	u, err := uploader.New(&cl, "factoryID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	profiles := []string{"h264"}

	// To setup a debug Log:
	//
	// u.DebugLog = log.New(os.Stdout, "", log.LstdFlags)
	//

	if err := u.Upload(f, fi.Name(), fi.Size(), profiles); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExampleUploader_UploadSession() {
	f, err := os.Open("video.mp4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cl := client.Client{
		Host: "api.cloud.telestream.net:443",
		Options: &client.Options{
			AccessKey: "access_key",
			SecretKey: "secret_key",
		},
	}

	u, err := uploader.New(&cl, "factoryID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	profiles := []string{"h264"}

	s, err := u.NewSession(fi.Name(), fi.Size(), profiles)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := u.UploadSession(s, f, fi.Size()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
