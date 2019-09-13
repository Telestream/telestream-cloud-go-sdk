/*
 * Qc API
 *
 * Qc API
 *
 * API version: 3.0.0
 * Contact: cloudsupport@telestream.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package qc

type UploadSession struct {
	// An unique identifier of the UploadSession.
	Id string `json:"id"`
	// An URL to which chunks of the uploaded file should be sent
	Location string `json:"location"`
	// A number of chunks that are expected by the upstream.
	Parts int32 `json:"parts,string,omitempty"`
	// An expected size of uploaded chunks.
	PartSize int32 `json:"part_size,string,omitempty"`
	// A maximum number of concurrent connections.
	MaxConnections int32 `json:"max_connections,string,omitempty"`
	// An object containing additional files uploaded using the session.
	ExtraFiles map[string]interface{} `json:"extra_files,omitempty"`
}