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

type FileDurationTest struct {
	FileDurationLowerHours int32 `json:"file_duration_lower_hours,omitempty"`
	FileDurationLowerMinutes int32 `json:"file_duration_lower_minutes,omitempty"`
	FileDurationLowerSeconds int32 `json:"file_duration_lower_seconds,omitempty"`
	FileDurationLowerFrames int32 `json:"file_duration_lower_frames,omitempty"`
	FileDurationUpperHours int32 `json:"file_duration_upper_hours,omitempty"`
	FileDurationUpperMinutes int32 `json:"file_duration_upper_minutes,omitempty"`
	FileDurationUpperSeconds int32 `json:"file_duration_upper_seconds,omitempty"`
	FileDurationUpperFrames int32 `json:"file_duration_upper_frames,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
