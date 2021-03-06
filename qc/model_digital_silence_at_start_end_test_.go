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

type DigitalSilenceAtStartEndTest struct {
	DurationAtStart float64 `json:"duration_at_start,omitempty"`
	DurationAtEnd float64 `json:"duration_at_end,omitempty"`
	DurationSecsOrFrames SecsOrFramesType `json:"duration_secs_or_frames,omitempty"`
	MustOrMustNotBeSilent MustOrMustNot `json:"must_or_must_not_be_silent,omitempty"`
	Channels Channels `json:"channels,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	DoCorrection bool `json:"do_correction,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
