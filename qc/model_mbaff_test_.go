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

type MbaffTest struct {
	MbaffRequired bool `json:"mbaff_required,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
