# Telestream Cloud API client

## Getting Started
### Obtain address for TCS api
In order to use TCS api client first you need to get `ApiKey`. Login to [website](https://cloud.telestream.net/console), go to *Flip* service and open *API Access* tab.
You account will be identified by unique *Api Key*, if it is unavailable click *Reset* button.
### Installation

    go get github.com/Telestream/telestream-cloud-go-sdk/...

### Usage
This example show uploading media file to flip service. If you want to use other service refer to [services](#services).

    package main

    import (
        "fmt"
        "os"

        "github.com/Telestream/telestream-cloud-go-sdk/flip"
        "github.com/Telestream/telestream-cloud-go-sdk/flip/uploader"
    )

    func main() {
        api_key := "tcs_78fecxxf034"
        factory := "ta0180de394aec03412x840ca8304"
        profiles := "h264"

        client := flip.NewFlipApi()
        client.Configuration.APIKey["X-Api-Key"] = api_key

        upload, err := uploader.New(client, factory)
        if err != nil {
            panic(err)
        }

        file, err := os.Open("video.mp4")
        if err != nil {
            panic(err)
        }

        fileInfo, err := file.Stat()
        if err != nil {
            panic(err)
        }

        err = upload.Upload(file, fileInfo.Name(), fileInfo.Size(), profiles, nil)
        if err != nil {
            panic(err)
        }

        fmt.Printf("Created new video %s\n", upload.MediaID)
    }


## Services
Api client is divided into parts corresponding to services provided. Currently supported services include:
- [Flip](flip/README.md) - high-volume media transcoding to multiple formats
- [Timed Text Speech](tts/README.md) - automated captions and subtitles creation
- [Quality Control](qc/README.md) - automated quality control for file base media
