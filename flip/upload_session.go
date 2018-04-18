/*
 * Flip API
 *
 * Description
 *
 * API version: 2.0.1
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flip

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
	ExtraFiles *interface{} `json:"extra_files,omitempty"`
}
