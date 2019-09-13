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

type PictureOffsetsTest struct {
	StoredFtwo int32 `json:"stored_ftwo,omitempty"`
	DisplayFtwo int32 `json:"display_ftwo,omitempty"`
	SampledX int32 `json:"sampled_x,omitempty"`
	SampledY int32 `json:"sampled_y,omitempty"`
	DisplayX int32 `json:"display_x,omitempty"`
	DisplayY int32 `json:"display_y,omitempty"`
	ImageStart int32 `json:"image_start,omitempty"`
	ImageEnd int32 `json:"image_end,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}