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

type ProjectBody struct {
	// Human-readable identifier of a Project.
	Name string `json:"name,omitempty"`
	// Name of QC template.
	Template string `json:"template,omitempty"`
	// JSON with specific options
	Options map[string]interface{} `json:"options,omitempty"`
}
