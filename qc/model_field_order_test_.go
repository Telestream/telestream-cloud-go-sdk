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

type FieldOrderTest struct {
	FlaggedFieldOrder FieldOrderType `json:"flagged_field_order,omitempty"`
	BasebandEnabled bool `json:"baseband_enabled,omitempty"`
	Simple bool `json:"simple,omitempty"`
	BasebandFieldOrder FieldOrderType `json:"baseband_field_order,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}