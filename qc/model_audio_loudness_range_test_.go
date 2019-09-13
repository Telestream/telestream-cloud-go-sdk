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

type AudioLoudnessRangeTest struct {
	DoMin bool `json:"do_min,omitempty"`
	RangeMin float32 `json:"range_min,omitempty"`
	DoMax bool `json:"do_max,omitempty"`
	RangeMax float32 `json:"range_max,omitempty"`
	Channels Channels `json:"channels,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}