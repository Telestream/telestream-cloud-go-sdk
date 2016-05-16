package uploader

import (
	"encoding/json"
	"strconv"
)

type Session struct {
	ID             string `json:"id"`
	Location       string `json:"location"`
	Parts          int    `json:"parts"`
	PartSize       int64  `json:"part_size"`
	MaxConnections int    `json:"max_connections"`
}

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
