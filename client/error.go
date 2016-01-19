package client

import "fmt"

// Error represents an error which is returned from the Telestream Cloud API.
type Error struct {
	Code    int
	Err     string `json:"error"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("telestream cloud: %d %s: %s", e.Code, e.Err, e.Message)
}
