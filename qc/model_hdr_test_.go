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

type HdrTest struct {
	HdrStandard HdrStandardType `json:"hdr_standard,omitempty"`
	MaxFallMaxEnabled bool `json:"max_fall_max_enabled,omitempty"`
	MaxFallMax int32 `json:"max_fall_max,omitempty"`
	MaxFallErrorEnabled bool `json:"max_fall_error_enabled,omitempty"`
	MaxFallError int32 `json:"max_fall_error,omitempty"`
	MaxCllMaxEnabled bool `json:"max_cll_max_enabled,omitempty"`
	MaxCllMax int32 `json:"max_cll_max,omitempty"`
	MaxCllErrorEnabled bool `json:"max_cll_error_enabled,omitempty"`
	MaxCllError int32 `json:"max_cll_error,omitempty"`
	AlwaysCalculate bool `json:"always_calculate,omitempty"`
	AlwaysReport bool `json:"always_report,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}