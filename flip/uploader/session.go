package uploader

import (
	"encoding/json"
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
	Parts int `json:"parts"`

	// PartSize specifies the size of each part. The last part can be smaller
	// than this value.
	PartSize int64 `json:"part_size"`

	// MaxConnections specifies the maximum number of possible concurrent
	// uploads.
	MaxConnections int `json:"max_connections"`
}

// UnmarshalJSON implements the json.Unmarshaler interface. Telestream Cloud's
// API returns every value as string.
func (s *Session) UnmarshalJSON(b []byte) error {
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	partSize, err := strconv.ParseInt(m["part_size"], 10, 64)
	if err != nil {
		return err
	}
	parts, err := strconv.Atoi(m["parts"])
	if err != nil {
		return err
	}
	maxConnections, err := strconv.Atoi(m["max_connections"])
	if err != nil {
		return err
	}
	s.ID = m["id"]
	s.Location = m["location"]
	s.Parts = parts
	s.PartSize = partSize
	s.MaxConnections = maxConnections
	return nil
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
