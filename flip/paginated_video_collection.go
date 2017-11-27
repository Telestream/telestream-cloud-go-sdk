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

type PaginatedVideoCollection struct {

	Videos []Video `json:"videos,omitempty"`

	// A number of the fetched page.
	Page int32 `json:"page,omitempty"`

	// A number of videos per page.
	PerPage int32 `json:"per_page,omitempty"`

	// A number of all videos stored in the db.
	Total int32 `json:"total,omitempty"`
}
