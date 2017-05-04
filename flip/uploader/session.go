package uploader

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Session describes upload API's upload session.
type Session struct {
	// ID specifies a unique upload session ID.
	ID string `json:"id"`

	// Location specifies a unique resource location for the file upload.
	Location string `json:"location"`

	// Parts specifies the number of parts the file is supposed to be split into
	// for uploading purposes.
	Parts int `json:"parts,string"`

	// PartSize specifies the size of each part. The last part can be smaller
	// than this value.
	PartSize int64 `json:"part_size,string"`

	// MaxConnections specifies the maximum number of possible concurrent
	// uploads.
	MaxConnections int `json:"max_connections,string"`

	ExtraFiles map[string]extra_files `json:"extra_files"`
}

type extra_files struct {
	// Parts specifies the number of parts the file is supposed to be split into
	// for uploading purposes.
	Parts int `json:"parts,string"`

	// PartSize specifies the size of each part. The last part can be smaller
	// than this value.
	PartSize int64 `json:"part_size,string"`
}

type ExtraFileInfo []struct {
	Tag   string
	Files []string
}

func (self *ExtraFileInfo) MarshalJSON() ([]byte, error) {
	data := make([]map[string]interface{}, 0)

	for _, tag := range *self {
		for i, name := range tag.Files {
			key := fmt.Sprintf("%s.index-%d", tag.Tag, i)
			st, err := os.Stat(name)

			if err != nil {
				return nil, err
			}

			data = append(data, map[string]interface{}{
				"tag":       key,
				"file_name": name,
				"file_size": st.Size(),
			})
		}
	}

	return json.Marshal(data)
}

// MarshalJSON implements the json.Marshaler interface. Telestream Cloud's API
// returns every value as string.
func (s *Session) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"id":              s.ID,
		"location":        s.Location,
		"parts":           strconv.Itoa(s.Parts),
		"part_size":       strconv.FormatInt(s.PartSize, 10),
		"max_connections": strconv.Itoa(s.MaxConnections),
	})
}
