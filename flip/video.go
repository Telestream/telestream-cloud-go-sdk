/*
 * Flip API
 *
 * Description
 *
 * API version: 2.0.1
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flip

type Video struct {

	// A unique identifier of the Video.
	Id string `json:"id,omitempty"`

	// audio bitrate (in bits/s)
	AudioBitrate int32 `json:"audio_bitrate,omitempty"`

	// A number of audio channels.
	AudioChannels int32 `json:"audio_channels,omitempty"`

	// A codec that has been used to encode audio streams.
	AudioCodec string `json:"audio_codec,omitempty"`

	// A number of samples of audio carried per second.
	AudioSampleRate int32 `json:"audio_sample_rate,omitempty"`

	// A date and time when the Video has been created.
	CreatedAt string `json:"created_at,omitempty"`

	// A duration of the video in seconds.
	Duration int32 `json:"duration,omitempty"`

	// A number of related Encoding objects.
	EncodingsCount int32 `json:"encodings_count,omitempty"`

	// A class of an error that has occurred during the encoding process. It is present only if the encoding status is equal to `fail`.
	ErrorClass string `json:"error_class,omitempty"`

	// A message that explains why the encoding process has resulted in an error. It is present only if the encoding status is equal to `fail`.
	ErrorMessage string `json:"error_message,omitempty"`

	// Extension of the source file.
	Extname string `json:"extname,omitempty"`

	// A size of the source file.
	FileSize int64 `json:"file_size,omitempty"`

	// Number of frames per second.
	Fps float32 `json:"fps,omitempty"`

	// Height of the output video.
	Height int32 `json:"height,omitempty"`

	// Width of the output video.
	Width int32 `json:"width,omitempty"`

	// A mime type of the source file.
	MimeType string `json:"mime_type,omitempty"`

	// A name of the source file.
	OriginalFilename string `json:"original_filename,omitempty"`

	Path string `json:"path,omitempty"`

	// Payload is an arbitrary text of length 256 or shorter that you can store along the Video. It is typically used to retain an association with one of your own DB record ID.
	Payload string `json:"payload,omitempty"`

	// An URL pointing to the source file.
	SourceUrl string `json:"source_url,omitempty"`

	// Determines at what stage of importing process the Video is at the moment.
	Status string `json:"status,omitempty"`

	// A date and time when a Video has been updated last time.
	UpdatedAt string `json:"updated_at,omitempty"`

	// video bitrate (in bits/s)
	VideoBitrate int32 `json:"video_bitrate,omitempty"`

	// A codec that has been used to encode the input file's video streams.
	VideoCodec string `json:"video_codec,omitempty"`
}
