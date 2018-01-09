/*
 * Tts API
 *
 * Description
 *
 * API version: 3.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tts

type FragmentVariant struct {

	// An alternative hypothesis for a fragment from the input audio.
	Fragment string `json:"fragment,omitempty"`

	// The confidence score of the fragment variant hypothesis in the range of 0 to 1.
	Confidence float32 `json:"confidence,omitempty"`
}