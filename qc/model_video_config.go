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

type VideoConfig struct {
	TrackSelectTest TrackSelectTest `json:"track_select_test,omitempty"`
	TrackIdTest TrackIdTest `json:"track_id_test,omitempty"`
	IgnoreVbiTest IgnoreVbiTest `json:"ignore_vbi_test,omitempty"`
	ForceColorSpaceTest ForceColorSpaceTest `json:"force_color_space_test,omitempty"`
	VideoSegmentDetectionTest VideoSegmentDetectionTest `json:"video_segment_detection_test,omitempty"`
	VideoLayoutTest LayoutTest `json:"video_layout_test,omitempty"`
	LetterboxingTest LetterboxingTest `json:"letterboxing_test,omitempty"`
	BlankingTest BlankingTest `json:"blanking_test,omitempty"`
	LossOfChromaTest LossOfChromaTest `json:"loss_of_chroma_test,omitempty"`
	ChromaLevelTest ChromaLevelTest `json:"chroma_level_test,omitempty"`
	BlackLevelTest BlackLevelTest `json:"black_level_test,omitempty"`
	RgbGamutTest RgbGamutTest `json:"rgb_gamut_test,omitempty"`
	HdrTest HdrTest `json:"hdr_test,omitempty"`
	ColourBarsTest ColourBarsTest `json:"colour_bars_test,omitempty"`
	BlackFrameTest BlackFrameTest `json:"black_frame_test,omitempty"`
	SingleColorTest SingleColorTest `json:"single_color_test,omitempty"`
	FreezeFrameTest FreezeFrameTest `json:"freeze_frame_test,omitempty"`
	BlockinessTest BlockinessTest `json:"blockiness_test,omitempty"`
	FieldOrderTest FieldOrderTest `json:"field_order_test,omitempty"`
	CadenceTest CadenceTest `json:"cadence_test,omitempty"`
	DropoutTest DropoutTest `json:"dropout_test,omitempty"`
	DigitalDropoutTest DigitalDropoutTest `json:"digital_dropout_test,omitempty"`
	StripeTest StripeTest `json:"stripe_test,omitempty"`
	CorruptFrameTest CorruptFrameTest `json:"corrupt_frame_test,omitempty"`
	FlashTest FlashTest `json:"flash_test,omitempty"`
	MediaOfflineTest MediaOfflineTest `json:"media_offline_test,omitempty"`
}