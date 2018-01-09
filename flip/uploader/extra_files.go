package uploader

import (
	sdk "github.com/Telestream/telestream-cloud-go-sdk/flip"
	"fmt"
	"io"
)

type ExtraFilesInfo []ExtraFileInfo

type ExtraFileInfo struct {
	Tag   string
	Files []ExtraFileItem
}

type ExtraFileItem struct {
	Name string
	File io.ReaderAt
	Size int64
}

func (efi ExtraFilesInfo) ConvertToExtraFiles() []sdk.ExtraFile {
	var data []sdk.ExtraFile

	for _, tag := range efi {
		for i, file := range tag.Files {
			key := fmt.Sprintf("%s.index-%d", tag.Tag, i)
			data = append(data, sdk.ExtraFile{
				Tag:      key,
				FileName: file.Name,
				FileSize: file.Size,
			})
		}
	}
	return data
}
