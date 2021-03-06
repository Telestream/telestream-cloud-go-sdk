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

type AdvancedGopLengthTest struct {
	FirstGopEnabled bool `json:"first_gop_enabled,omitempty"`
	FirstGopI string `json:"first_gop_i,omitempty"`
	FirstGopP string `json:"first_gop_p,omitempty"`
	FirstGopClosed OpenOrClosed `json:"first_gop_closed,omitempty"`
	OtherGopEnabled bool `json:"other_gop_enabled,omitempty"`
	OtherGopI string `json:"other_gop_i,omitempty"`
	OtherGopP string `json:"other_gop_p,omitempty"`
	OtherGopClosed OpenOrClosed `json:"other_gop_closed,omitempty"`
	LastGopEnabled bool `json:"last_gop_enabled,omitempty"`
	LastGopI string `json:"last_gop_i,omitempty"`
	LastGopP string `json:"last_gop_p,omitempty"`
	LastGopClosed OpenOrClosed `json:"last_gop_closed,omitempty"`
	Order GopOrder `json:"order,omitempty"`
	Report GopReport `json:"report,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
