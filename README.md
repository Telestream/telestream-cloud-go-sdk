# telestream-cloud-go-sdk

[![API Reference](http://img.shields.io/badge/api-reference-blue.svg)](http://cloud.telestream.net/docs#api)
[![GoDoc](https://godoc.org/github.com/Telestream/telestream-cloud-go-sdk?status.svg)](http://godoc.org/github.com/Telestream/telestream-cloud-go-sdk)

telestream-cloud-go-sdk is the official Telestream Cloud SDK for the Go programming language.

## Installation


    go get github.com/Telestream/telestream-cloud-go-sdk/...


## Using flip

Package flip provides a client for Telestream Cloud transcoding service.

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/Telestream/telestream-cloud-go-sdk/flip"
)

func main() {
  fp := flip.NewFlip(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY"), nil)
  factory, err := fp.Factory(os.Getenv("FACTORY_ID"))
  if err != nil {
    log.Fatal(err)
  }
  profile, err := factory.NewProfile(&flip.NewProfileRequest{
    PresetName: "h264",
    Name:       "h264-640x480",
    Width:      640,
    Height:     480,
    AspectMode: flip.ModeLetterBox,
  })
  if err != nil {
    log.Fatal(err)
  }
  video, err := factory.NewVideo("path/to/file.mp4", &flip.NewVideoRequest{
    PathFormat: ":profile/:id",
    Profiles:   []string{profile.Name},
  })
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("videoID:", video.ID)
  encodings, err := factory.VideoEncodings(video.ID, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, encoding := range encodings {
    fmt.Println("encodingID:", encoding.ID)
  }
}
```

For more information and examples look at [GoDoc](http://godoc.org/github.com/Telestream/telestream-cloud-go-sdk/flip).
