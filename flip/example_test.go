package flip_test

import (
	"fmt"

	"github.com/Telestream/telestream-cloud-go-sdk/flip"
)

func ExampleFlip_Factories() {
	factories, err := flip.NewFlip("accessKey", "secretKey", nil).Factories()
	if err != nil {
		panic(err)
	}
	fmt.Println(factories)
}

func ExampleFactory_NewVideoURL() {
	factory, err := flip.NewFlip("accessKey", "secretKey", nil).
		Factory("factoryID")
	if err != nil {
		panic(err)
	}
	video, err := factory.NewVideoURL("www.example.com", &flip.NewVideoRequest{
		Profiles: []string{"MyProfile"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(video)
}

func ExampleFactory_NewProfile() {
	factory, err := flip.NewFlip("accessKey", "secretKey", nil).
		Factory("factoryID")
	if err != nil {
		panic(err)
	}
	profile, err := factory.NewProfile(&flip.NewProfileRequest{
		PresetName: "h264",
		AspectMode: flip.ModeLetterBox,
		Fps:        29.7,
		Name:       "MyNewProfile",
		Width:      640,
		Height:     480,
		TimeCode:   true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)
}

func ExampleFactory_NewEncoding() {
	factory, err := flip.NewFlip("accessKey", "secretKey", nil).
		Factory("factoryID")
	encoding, err := factory.NewEncoding(&flip.NewEncodingRequest{
		VideoID:     "videoID",
		ProfileName: "MyProfile",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(encoding)
}

func ExampleFactory_Update() {
	factory, err := flip.NewFlip("accessKey", "secretKey", nil).
		Factory("factoryID")
	if err != nil {
		panic(err)
	}
	profile, err := factory.Profile("profileID")
	if err != nil {
		panic(err)
	}
	profile.AspectMode = flip.ModeConstrain
	if err := factory.Update(profile); err != nil {
		panic(err)
	}
	fmt.Println(profile)
}

func ExampleFactory_Delete() {
	factory, err := flip.NewFlip("accessKey", "secretKey", nil).
		Factory("factoryID")
	if err != nil {
		panic(err)
	}
	profile, err := factory.Profile("profileID")
	if err != nil {
		panic(err)
	}
	if err := factory.Delete(profile); err != nil {
		panic(err)
	}
}
