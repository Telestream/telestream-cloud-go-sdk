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

type AudioPeakLevelTest struct {
	MinPeakLevelEnabled bool `json:"min_peak_level_enabled,omitempty"`
	MinPeakLevel float32 `json:"min_peak_level,omitempty"`
	MaxPeakLevelEnabled bool `json:"max_peak_level_enabled,omitempty"`
	MaxPeakLevel float32 `json:"max_peak_level,omitempty"`
	CorrectionThreshold float32 `json:"correction_threshold,omitempty"`
	Channels Channels `json:"channels,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	DoCorrection bool `json:"do_correction,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
