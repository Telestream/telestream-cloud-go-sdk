package flip

import (
	"errors"
	"path"
	"time"
)

// UpdatableResource is the interface that wraps UpdatePath() method used by
// the objects which can be updated.
//
// UpdatePath returns a full path on which a PUT request must be sent in order
// to update the resource.
type UpdatableResource interface {
	UpdatePath() (string, error)
}

// DeletableResource is the interface that wraps DeletePath() method used by
// the objects which can be deleted.
//
// DeletePath returns a full path on which a DELETE request must be sent in
// order to update the resource.
type DeletableResource interface {
	DeletePath() (string, error)
}

// Status describes the current state of the video or the encoding.
type Status string

const (
	// StatusSuccess indicates that the job has been done successfully.
	StatusSuccess = Status("success")
	// StatusFail indicates that the job has failed. Seek ErrorClass and
	// ErrorMessage fields for more information.
	StatusFail = Status("fail")
	// StatusProcessing indicates that the job is still being processed.
	StatusProcessing = Status("processing")
)

// ClosedCaptionsType describes a mode in which closed captions should be
// encoded.
type ClosedCaptionsType string

const (
	// ClosedCaptionsBurn will cause the encoder to burn the closed captions
	// in an output video.
	ClosedCaptionsBurn = ClosedCaptionsType("burn")
	// ClosedCaptionsAdd will add captions as seperate streams.
	ClosedCaptionsAdd = ClosedCaptionsType("add")
)

// DashProfile describes a type of available media formats.
type DashProfile string

const (
	// DashProfileLive describes dash live profile
	// urn:mpeg:dash:profile:isoff-live:2011
	DashProfileLive = DashProfile("live")
	// DashProfileOnDemand describes dash on-demand profile
	// urn:mpeg:dash:profile:isoff-on-demand:2011
	DashProfileOnDemand = DashProfile("on-demand")
)

// InterlaceMode describes a type of available temporal field interlacing modes.
type InterlaceMode string

const (
	// InterlaceTop interleaves the upper field from odd frames with the lower
	// field from even frames, generating a frame with unchanged height at half
	// frame rate.
	InterlaceTop = InterlaceMode("top")
	// InterlaceBottom interleaves the lower field from odd frames with the upper
	// field from even frames, generating a frame with unchanged height at half
	// frame rate.
	InterlaceBottom = InterlaceMode("bottom")
)

// AspectMode describes a type of available aspect ratio modes.
type AspectMode string

const (
	// ModeLetterBox - Aspect ratio is maintained. Adds black bars to your output
	// to match your profile frame height (above and below only).
	ModeLetterBox = AspectMode("letterbox")
	// ModePreserve - Original size and ratio is preserved.
	ModePreserve = AspectMode("preserve")
	// ModeConstrain - Aspect ratio is maintained. No black bars is added to your
	// output.
	ModeConstrain = AspectMode("constrain")
	// ModePad - Aspect ratio is maintained. Adds black bars to your output to
	// match your profile frame size.
	ModePad = AspectMode("pad")
	// ModeStretch - Stretches the video to the frame size.
	ModeStretch = AspectMode("stretch")
	// ModeCenter - Aspect ratio of the actual video is maintained. However, black
	// bars might be added to your output.
	ModeCenter = AspectMode("center")
	// ModeFill - Aspect ratio is maintained. Fills your profile frame size and
	// crops the rest.
	ModeFill = AspectMode("fill")
)

// MetaData holds all the video's metadata
type MetaData map[string]interface{}

const timeFormat = "2006/01/02 15:04:05 -0700"

// Time holds time.Time and is capable of marshalling and unmarshalling Panda's
// timestamp in the correct way
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface. The time is a quoted
// string in 2006/01/02 15:04:05 -0700 format.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"` + timeFormat + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. The time is expected
// to be a quoted string in 2006/01/02 15:04:05 -0700 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	time, err := time.Parse(`"`+timeFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*t = Time{time}
	return nil
}

// Encoding represents a version of a video encoded with the settings defined in
// the profile used. When an encoding is created (either automatically when a
// video is uploaded, or using the POST method below) it is immediately placed
// on your encoding queue.
//
// If the original video has been successfully uploaded to your s3 bucket, the
// video status will be set to success. If the video has failed, the encoding
// will fail as well.
//
// If the encoding was successful, you can download it or play it using the
// following convention Path+Extname.
//
// If the encoding fails, this is usually because there was either an invalid
// profile being used, or the origin video is corrupted or not recognised.
//
// It implements the DeletableResource interface.
type Encoding struct {
	AudioBitrate      int      `json:"audio_bitrate,omitempty" url:"audio_bitrate,omitempty"`
	AudioChannels     int      `json:"audio_channels,omitempty" url:"audio_channels,omitempty"`
	AudioCodec        string   `json:"audio_codec,omitempty" url:"audio_codec,omitempty"`
	AudioSampleRate   int      `json:"audio_sample_rate,omitempty" url:"audio_sample_rate,omitempty"`
	CreatedAt         Time     `json:"created_at,omitempty" url:"created_at,omitempty"`
	Duration          int      `json:"duration,omitempty" url:"duration,omitempty"`
	EncodingProgress  float64  `json:"encoding_progress,omitempty" url:"encoding_progress,omitempty"`
	EncodingTime      float64  `json:"encoding_time,omitempty" url:"encoding_time,omitempty"`
	ErrorClass        string   `json:"error_class,omitempty" url:"error_class,omitempty"`
	ErrorMessage      string   `json:"error_message,omitempty" url:"error_message,omitempty"`
	ExternalID        string   `json:"external_id,omitempty" url:"external_id,omitempty"`
	Extname           string   `json:"extname,omitempty" url:"extname,omitempty"`
	FileSize          int64    `json:"file_size,omitempty" url:"file_size,omitempty"`
	Files             []string `json:"files,omitempty" url:"files,omitempty"`
	Fps               float64  `json:"fps,omitempty" url:"fps,omitempty"`
	Height            int      `json:"height,omitempty" url:"height,omitempty"`
	ID                string   `json:"id,omitempty" url:"id,omitempty"`
	MimeType          string   `json:"mime_type,omitempty" url:"mime_type,omitempty"`
	Path              string   `json:"path,omitempty" url:"path,omitempty"`
	ProfileID         string   `json:"profile_id,omitempty" url:"profile_id,omitempty"`
	ProfileName       string   `json:"profile_name,omitempty" url:"profile_name,omitempty"`
	StartedEncodingAt Time     `json:"started_encoding_at,omitempty" url:"started_encoding_at,omitempty"`
	Status            Status   `json:"status,omitempty" url:"status,omitempty"`
	UpdatedAt         Time     `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	VideoBitrate      int      `json:"video_bitrate,omitempty" url:"video_bitrate,omitempty"`
	VideoCodec        string   `json:"video_codec,omitempty" url:"video_codec,omitempty"`
	VideoID           string   `json:"video_id,omitempty" url:"video_id,omitempty"`
	Width             int      `json:"width,omitempty" url:"width,omitempty"`
	Screenshots       []string `json:"screenshots,omitempty" url:"screenshots,omitempty"`
}

// DeletePath returns an api path on which delete encoding request must be send.
// It returns error if ID is empty.
func (e *Encoding) DeletePath() (string, error) {
	if e.ID == "" {
		return "", errors.New("there must be an ID available to delete this object")
	}
	return path.Join("/encodings", e.ID+".json"), nil
}

// Video is a resource representing the original media. It describes all major
// metadata of a media file, it’s size and it’s original filename. It contains a
// status attribute representing the uploading process of media to
// S3. Status is set to success when the media has been successfully uploaded to
// S3. Careful: It does not mean that it’s encodings have been encoded. It
// implements the DeletableResource interface.
type Video struct {
	AudioBitrate     int     `json:"audio_bitrate,omitempty" url:"audio_bitrate,omitempty"`
	AudioChannels    int     `json:"audio_channels,omitempty" url:"audio_channels,omitempty"`
	AudioCodec       string  `json:"audio_codec,omitempty" url:"audio_codec,omitempty"`
	AudioSampleRate  int     `json:"audio_sample_rate,omitempty" url:"audio_sample_rate,omitempty"`
	CreatedAt        Time    `json:"created_at,omitempty" url:"created_at,omitempty"`
	Duration         int     `json:"duration,omitempty" url:"duration,omitempty"`
	ErrorClass       string  `json:"error_class,omitempty" url:"error_class,omitempty"`
	ErrorMessage     string  `json:"error_message,omitempty" url:"error_message,omitempty"`
	Extname          string  `json:"extname,omitempty" url:"extname,omitempty"`
	FileSize         int64   `json:"file_size,omitempty" url:"file_size,omitempty"`
	Fps              float64 `json:"fps,omitempty" url:"fps,omitempty"`
	Height           int     `json:"height,omitempty" url:"height,omitempty"`
	ID               string  `json:"id,omitempty" url:"id,omitempty"`
	MimeType         string  `json:"mime_type,omitempty" url:"mime_type,omitempty"`
	OriginalFilename string  `json:"original_filename,omitempty" url:"original_filename,omitempty"`
	Path             string  `json:"path,omitempty" url:"path,omitempty"`
	Payload          string  `json:"payload,omitempty" url:"payload,omitempty"`
	SourceURL        string  `json:"source_url,omitempty" url:"source_url,omitempty"`
	Status           Status  `json:"status,omitempty" url:"status,omitempty"`
	UpdatedAt        Time    `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	VideoBitrate     int     `json:"video_bitrate,omitempty" url:"video_bitrate,omitempty"`
	VideoCodec       string  `json:"video_codec,omitempty" url:"video_codec,omitempty"`
	Width            int     `json:"width,omitempty" url:"width,omitempty"`
}

// DeletePath returns an api path on which delete video request must be send.
// It returns error if ID is empty.
func (v *Video) DeletePath() (string, error) {
	if v.ID == "" {
		return "", errors.New("there must be an ID available to delete this object")
	}
	return path.Join("/videos", v.ID+".json"), nil
}

// Profile holds the informations about the created Profile. Some values are
// reserved for specific presets. This struct gathers all possible combinations.
// It implements DeletableResource and UpdatableResource interface.
type Profile struct {
	AspectMode          AspectMode         `json:"aspect_mode,omitempty" url:"aspect_mode,omitempty"`
	AspectRatio         string             `json:"aspect_ratio,omitempty" url:"aspect_ratio,omitempty"`
	AudioBitrate        int                `json:"audio_bitrate,omitempty" url:"audio_bitrate,omitempty"`
	AudioChannels       int                `json:"audio_channels,omitempty" url:"audio_channels,omitempty"`
	AudioSampleRate     int                `json:"audio_sample_rate,omitempty" url:"audio_sample_rate,omitempty"`
	BufferSize          int                `json:"buffer_size,omitempty" url:"buffer_size,omitempty"`
	Cbr                 bool               `json:"cbr,omitempty" url:"cbr,omitempty"`
	ClipLength          string             `json:"clip_length,omitempty" url:"clip_length,omitempty"`
	ClipOffset          string             `json:"clip_offset,omitempty" url:"clip_offset,omitempty"`
	ClosedCaptions      ClosedCaptionsType `json:"closed_captions,omitempty" url:"closed_captions,omitempty"`
	Command             string             `json:"command,omitempty" url:"command,omitempty"`
	CreatedAt           Time               `json:"created_at,omitempty" url:"created_at,omitempty"`
	DashProfile         DashProfile        `json:"dash_profile" url:"dash_profile,omitempty"`
	Deinterlace         string             `json:"deinterlace,omitempty" url:"deinterlace,omitempty"`
	Encryption          bool               `json:"encryption,omitempty" url:"encryption,omitempty"`
	EncryptionIV        string             `json:"encryption_iv,omitempty" url:"encryption_iv,omitempty"`
	EncryptionKey       string             `json:"encryption_key,omitempty" url:"encryption_key,omitempty"`
	EncryptionKeyURL    string             `json:"encryption_key_url,omitempty" url:"encryption_key_url,omitempty"`
	Extname             string             `json:"extname,omitempty" url:"extname,omitempty"`
	Fps                 float64            `json:"fps,omitempty" url:"fps,omitempty"`
	FrameCount          int                `json:"frame_count,omitempty" url:"frame_count,omitempty"`
	FrameInterval       string             `json:"frame_interval,omitempty" url:"frame_interval,omitempty"`
	FrameOffsets        string             `json:"frame_offsets,omitempty" url:"frame_offsets,omitempty"`
	H264Crf             int                `json:"h264_crf,omitempty" url:"h264_crf,omitempty"`
	H264Level           string             `json:"h264_level,omitempty" url:"h264_level,omitempty"`
	H264Profile         string             `json:"h264_profile,omitempty" url:"h264_profile,omitempty"`
	H264Tune            string             `json:"h264_tune,omitempty" url:"h264_tune,omitempty"`
	X264Options         string             `json:"x264_options,omitempty" url:"x264_options,omitempty"`
	X265Options         string             `json:"x265_options,omitempty" url:"x265_options,omitempty"`
	Height              int                `json:"height,omitempty" url:"height,omitempty"`
	ID                  string             `json:"id,omitempty" url:"id,omitempty"`
	Interlace           InterlaceMode      `json:"interlace,omitempty" url:"interlace,omitempty"`
	KeyframeInterval    int                `json:"keyframe_interval,omitempty" url:"keyframe_interval,omitempty"`
	KeyframeRate        float64            `json:"Keyframe_rate,omitempty" url:"Keyframe_rate,omitempty"`
	MaxRate             int                `json:"max_rate,omitempty" url:"max_rate,omitempty"`
	MergeAudioStreams   bool               `json:"merge_audio_streams,omitempty" url:"merge_audio_streams,omitempty"`
	MotionCompensateFps bool               `json:"motion_compensate_fps,omitempty" url:"motion_compensate_fps,omitempty"`
	Name                string             `json:"name,omitempty" url:"name,omitempty"`
	PresetName          string             `json:"preset_name,omitempty" url:"preset_name,omitempty"`
	Priority            int                `json:"priority,omitempty" url:"priority,omitempty"`
	Stack               string             `json:"stack,omitempty" url:"stack,omitempty"`
	Title               string             `json:"title,omitempty" url:"title,omitempty"`
	TwoPass             string             `json:"two_pass,omitempty" url:"two_pass,omitempty"`
	UpdatedAt           Time               `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	Upscale             bool               `json:"upscale,omitempty" url:"upscale,omitempty"`
	Variants            string             `json:"variants,omitempty" url:"variants,omitempty"`
	VideoBitrate        int                `json:"video_bitrate,omitempty" url:"video_bitrate,omitempty"`
	WatermarkBottom     string             `json:"watermark_bottom,omitempty" url:"watermark_bottom,omitempty"`
	WatermarkHeight     string             `json:"watermark_height,omitempty" url:"watermark_height,omitempty"`
	WatermarkLeft       string             `json:"watermark_left,omitempty" url:"watermark_left,omitempty"`
	WatermarkRight      string             `json:"watermark_right,omitempty" url:"watermark_right,omitempty"`
	WatermarkTop        string             `json:"watermark_top,omitempty" url:"watermark_top,omitempty"`
	WatermarkURL        string             `json:"watermark_url,omitempty" url:"watermark_url,omitempty"`
	WatermarkWidth      string             `json:"watermark_width,omitempty" url:"watermark_width,omitempty"`
	Width               int                `json:"width,omitempty" url:"width,omitempty"`
	TimeCode            bool               `json:"time_code,omitempty" url:"time_code,omitempty"`
}

func firstNoEmpty(strs ...string) string {
	for _, str := range strs {
		if str != "" {
			return str
		}
	}
	return ""
}

// UpdatePath returns an api path on which an update profile request must be
// send. It fails if both ID and Name are empty.
func (p *Profile) UpdatePath() (string, error) {
	id := firstNoEmpty(p.ID, p.Name)
	if id == "" {
		return "", errors.New(
			"there must be an ID or a Name set to change this object")
	}
	return path.Join("/profiles", id+".json"), nil
}

// DeletePath returns an api path on which a delete profile request must be
// send. It fails if both ID and Name are empty.
func (p *Profile) DeletePath() (string, error) {
	return p.UpdatePath()
}

// Notification holds the information about the Factory's notification setings.
// It implements UpdatableResource interface.
type Notification struct {
	// Delay - notification delay.
	Delay float64 `json:"delay,omitempty" url:"delay,omitempty"`
	// Events when to send certain notifications.
	Events Events `json:"events,omitempty" url:"events,omitempty"`
	// URL where to send notifications.
	URL string `json:"url,omitempty" url:"url,omitempty"`
}

// UpdatePath returns an api path on which an update notification request must
// be send. The error is always nil.
func (p *Notification) UpdatePath() (string, error) {
	return "/notifications.json", nil
}

// Events hold the information on which events the notification should be sent.
type Events struct {
	// EncodingCompleted - notify when encoding has been completed.
	EncodingCompleted bool `json:"encoding_completed" url:"encoding_completed"`
	// EncodingProgress - notify when encoding started.
	EncodingProgress bool `json:"encoding_progress" url:"encoding_progress"`
	// VideoCreated - notify when video has been created.
	VideoCreated bool `json:"video_created" url:"video_created"`
	// VideoEncoded - notify when all encodings have been completed.
	VideoEncoded bool `json:"video_encoded" url:"video_encoded"`
}

// EncodingRequest defines desired filters while getting a list of encodings.
type EncodingRequest struct {
	// Page describes which page of record to list.
	Page int `json:"page,omitempty" url:"page,omitempty"`
	// PerPage describes how many records should be listed per page.
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
	// ProfileID can be used as a filter by profile's id.
	ProfileID string `json:"profile_id,omitempty" url:"profile_id,omitempty"`
	// ProfileName can be used as a filter by profile's name.
	ProfileName string `json:"profile_name,omitempty" url:"profile_name,omitempty"`
	// Status can be used as a filter.
	Status Status `json:"status,omitempty" url:"status,omitempty"`
	// VideoID can be used as a filter by videos's name.
	VideoID string `json:"video_id,omitempty" url:"video_id,omitempty"`
	// Screenshots can be used to show screenshots
	Screenshots bool `json:"screenshots,omitempty" url:"screenshots,omitempty"`
}

// NewEncodingRequest defines desired settings while creating a new encoding.
type NewEncodingRequest struct {
	// ProfileID can be used as a filter by profile's id.
	ProfileID string `json:"profile_id,omitempty" url:"profile_id,omitempty"`
	// ProfileName can be used as a filter by profile's name.
	ProfileName string `json:"profile_name,omitempty" url:"profile_name,omitempty"`
	// VideoID can be used as a filter by videos's name.
	VideoID string `json:"video_id,omitempty" url:"video_id,omitempty"`
}

// NewProfileRequest defines desired settings while creating a new profile.
// Some fields are reserved for the specific presets, and cannot be used at
// once.
type NewProfileRequest struct {
	AspectMode          AspectMode         `json:"aspect_mode,omitempty" url:"aspect_mode,omitempty"`
	AspectRatio         string             `json:"aspect_ratio,omitempty" url:"aspect_ratio,omitempty"`
	AudioBitrate        int                `json:"audio_bitrate,omitempty" url:"audio_bitrate,omitempty"`
	AudioChannels       int                `json:"audio_channels,omitempty" url:"audio_channels,omitempty"`
	AudioSampleRate     int                `json:"audio_sample_rate,omitempty" url:"audio_sample_rate,omitempty"`
	BufferSize          int                `json:"buffer_size,omitempty" url:"buffer_size,omitempty"`
	Cbr                 bool               `json:"cbr,omitempty" url:"cbr,omitempty"`
	ClipLength          string             `json:"clip_length,omitempty" url:"clip_length,omitempty"`
	ClipOffset          string             `json:"clip_offset,omitempty" url:"clip_offset,omitempty"`
	ClosedCaptions      ClosedCaptionsType `json:"closed_captions,omitempty" url:"closed_captions,omitempty"`
	Command             string             `json:"command,omitempty" url:"command,omitempty"`
	DashProfile         DashProfile        `json:"dash_profile" url:"dash_profile,omitempty"`
	Deinterlace         string             `json:"deinterlace,omitempty" url:"deinterlace,omitempty"`
	Encryption          bool               `json:"encryption,omitempty" url:"encryption,omitempty"`
	EncryptionIV        string             `json:"encryption_iv,omitempty" url:"encryption_iv,omitempty"`
	EncryptionKey       string             `json:"encryption_key,omitempty" url:"encryption_key,omitempty"`
	EncryptionKeyURL    string             `json:"encryption_key_url,omitempty" url:"encryption_key_url,omitempty"`
	Extname             string             `json:"extname,omitempty" url:"extname,omitempty"`
	Fps                 float64            `json:"fps,omitempty" url:"fps,omitempty"`
	FrameCount          int                `json:"frame_count,omitempty" url:"frame_count,omitempty"`
	FrameInterval       string             `json:"frame_interval,omitempty" url:"frame_interval,omitempty"`
	FrameOffsets        string             `json:"frame_offsets,omitempty" url:"frame_offsets,omitempty"`
	H264Crf             int                `json:"h264_crf,omitempty" url:"h264_crf,omitempty"`
	H264Level           string             `json:"h264_level,omitempty" url:"h264_level,omitempty"`
	H264Profile         string             `json:"h264_profile,omitempty" url:"h264_profile,omitempty"`
	H264Tune            string             `json:"h264_tune,omitempty" url:"h264_tune,omitempty"`
	X264Options         string             `json:"x264_options,omitempty" url:"x264_options,omitempty"`
	X265Options         string             `json:"x265_options,omitempty" url:"x265_options,omitempty"`
	Height              int                `json:"height,omitempty" url:"height,omitempty"`
	Interlace           InterlaceMode      `json:"interlace,omitempty" url:"interlace,omitempty"`
	KeyframeInterval    int                `json:"keyframe_interval,omitempty" url:"keyframe_interval,omitempty"`
	KeyframeRate        float64            `json:"Keyframe_rate,omitempty" url:"Keyframe_rate,omitempty"`
	MaxRate             int                `json:"max_rate,omitempty" url:"max_rate,omitempty"`
	MergeAudioStreams   bool               `json:"merge_audio_streams,omitempty" url:"merge_audio_streams,omitempty"`
	MotionCompensateFps bool               `json:"motion_compensate_fps,omitempty" url:"motion_compensate_fps,omitempty"`
	Name                string             `json:"name,omitempty" url:"name,omitempty"`
	PresetName          string             `json:"preset_name,omitempty" url:"preset_name,omitempty"`
	Priority            int                `json:"priority,omitempty" url:"priority,omitempty"`
	Stack               string             `json:"stack,omitempty" url:"stack,omitempty"`
	Title               string             `json:"title,omitempty" url:"title,omitempty"`
	TwoPass             string             `json:"two_pass,omitempty" url:"two_pass,omitempty"`
	Upscale             bool               `json:"upscale,omitempty" url:"upscale,omitempty"`
	Variants            string             `json:"variants,omitempty" url:"variants,omitempty"`
	VideoBitrate        int                `json:"video_bitrate,omitempty" url:"video_bitrate,omitempty"`
	WatermarkBottom     string             `json:"watermark_bottom,omitempty" url:"watermark_bottom,omitempty"`
	WatermarkHeight     string             `json:"watermark_height,omitempty" url:"watermark_height,omitempty"`
	WatermarkLeft       string             `json:"watermark_left,omitempty" url:"watermark_left,omitempty"`
	WatermarkRight      string             `json:"watermark_right,omitempty" url:"watermark_right,omitempty"`
	WatermarkTop        string             `json:"watermark_top,omitempty" url:"watermark_top,omitempty"`
	WatermarkURL        string             `json:"watermark_url,omitempty" url:"watermark_url,omitempty"`
	WatermarkWidth      string             `json:"watermark_width,omitempty" url:"watermark_width,omitempty"`
	Width               int                `json:"width,omitempty" url:"width,omitempty"`
	TimeCode            bool               `json:"time_code,omitempty" url:"time_code,omitempty"`
}

// NewVideoRequest defines desired settings while creating a new video.
type NewVideoRequest struct {
	// PathFormat represents the complete video path without the extension name.
	// It can be constructed using some provided keywords.
	//
	// :id          Id of the encoding. Required.
	//
	// :video_id    Id of the video.
	//
	// :original    Original filename of the uploaded video
	//
	// :date        Date the video was uploaded (2009-10-09).
	//
	// :profile     Profile name used by the encoding (original is used when the
	//              file is the original video).
	//
	// :type        Video type, original or encodings.
	//
	// :resolution  Resolution of the video or the encoding (480x340).
	PathFormat string `json:"path_format,omitempty" url:"path_format,omitempty"`
	// Payload is an arbitrary text of length 256 or shorter that you can store
	// along the Video. It is typically used to retain an association with one of
	// your own DB record ID.
	Payload string `json:"payload,omitempty" url:"payload,omitempty"`
	// Profiles is the list of profile names or IDs to be used during encoding.
	// Alternatively, specify none so no encodings are created yet.
	Profiles []string `json:"profiles,omitempty" url:"profiles,omitempty,comma"`
}

// ProfileRequest defines desired filters while getting a list of profiles.
type ProfileRequest struct {
	// If expand option is set Profile objects will contain all command
	// parameters, even if their value is default. By default is not set.
	Expand bool `json:"expand,omitempty" url:"expand,omitempty"`
	// Page describes which page of record to list.
	Page int `json:"page,omitempty" url:"page,omitempty"`
	// PerPage describes how many records should be listed per page.
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
}

// VideoRequest defines desired filters while getting a list of videos.
type VideoRequest struct {
	// Page describes which page of record to list.
	Page int `json:"page,omitempty" url:"page,omitempty"`
	// PerPage describes how many records should be listed per page.
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
	// Status can be used as a filter.
	Status Status `json:"status,omitempty" url:"status,omitempty"`
}
