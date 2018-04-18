/*
 * Qc API
 *
 * QC API
 *
 * API version: 2.0.1
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package qc

type Alert struct {

	Level string `json:"level,omitempty"`

	// Extra information for an alert.
	Info string `json:"info,omitempty"`

	// Start time of alert.
	Begin float32 `json:"begin,omitempty"`

	// End time of alert.
	End float32 `json:"end,omitempty"`

	Stream int32 `json:"stream,omitempty"`

	Detail string `json:"detail,omitempty"`
}
