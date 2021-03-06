/*
 * Notifications API
 *
 * Notifications API V2
 *
 * API version: 2.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package notifications

type Topic struct {

	// [read-only] Account identifier connected to subscription notification 
	Account string `json:"account,omitempty"`

	// Name of service (qc, flip, tts) 
	Service string `json:"service"`

	// Name of the event;  Quality Control (media-passed, media-error, media-warning, media-rejected, media-canceled, job-created, job-progress), Flip (video-created, video-encoded, encoding-complete, encoding-progress), Transcription (job-created, job-completed, job-error, job-progress) 
	Event string `json:"event"`

	// (for tts, qc service); Project ID 
	Project string `json:"project,omitempty"`

	// (for flip service); Factory ID 
	Factory string `json:"factory,omitempty"`
}
