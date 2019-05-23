/*
 * Qc API
 *
 * QC API
 *
 * API version: 2.0.3
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package qc

type Data struct {

	// Human-readable identifier of a Project.
	Name string `json:"name,omitempty"`

	// Name of QC template.
	Template string `json:"template,omitempty"`

	Options *Options `json:"options,omitempty"`
}
