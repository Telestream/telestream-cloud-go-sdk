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