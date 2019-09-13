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

type Project struct {
	// A unique identifier of a Project.
	Id string `json:"id,omitempty"`
	// Human-readable identifier of a Project.
	Name string `json:"name,omitempty"`
	// Project status.
	Status string `json:"status,omitempty"`
	// Name of QC template.
	Template string `json:"template,omitempty"`
	// JSON with specific options
	Options map[string]interface{} `json:"options,omitempty"`
}