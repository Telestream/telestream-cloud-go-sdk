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

type JobData struct {
	// JSON with specific options
	Options map[string]interface{} `json:"options,omitempty"`
	Url string `json:"url,omitempty"`
	// Payload is an arbitrary text of length 256 or shorter that you can store along the Media. It is typically used to retain an association with one of your own DB record ID.
	Payload string `json:"payload,omitempty"`
}