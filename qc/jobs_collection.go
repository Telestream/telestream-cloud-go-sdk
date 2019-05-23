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

type JobsCollection struct {

	Jobs []Job `json:"jobs,omitempty"`

	// A number of the fetched page.
	Page int32 `json:"page,omitempty"`

	// A number of jobs per page.
	PerPage int32 `json:"per_page,omitempty"`

	// A number of all pages.
	PageCount int32 `json:"page_count,omitempty"`

	// A number of all jobs stored in the db.
	TotalCount int32 `json:"total_count,omitempty"`
}
