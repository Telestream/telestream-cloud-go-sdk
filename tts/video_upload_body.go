/*
 * Tts API
 *
 * Description
 *
 * API version: 3.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tts

type VideoUploadBody struct {

	// Size of the file that will be uploaded in `bytes`.
	FileSize int64 `json:"file_size"`

	// Name of the file that will be uploaded.
	FileName string `json:"file_name"`

	Profiles string `json:"profiles,omitempty"`

	MultiChunk bool `json:"multi_chunk,omitempty"`

	// A list of names of additional files that will be uploaded.
	ExtraFiles []ExtraFile `json:"extra_files,omitempty"`
}