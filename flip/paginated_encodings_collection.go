/* 
 * Flip API
 *
 * Description
 *
 * OpenAPI spec version: 3.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package flip

type PaginatedEncodingsCollection struct {

	Encodings []Encoding `json:"encodings"`

	// A number of the fetched page.
	Page int32 `json:"page"`

	// A number of encodings per page.
	PerPage int32 `json:"per_page"`

	// A number of all encodings stored in the db.
	Total int32 `json:"total"`
}
