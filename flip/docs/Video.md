# Video

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the Video. | [optional] [default to null]
**AudioBitrate** | **int32** | audio bitrate (in bits/s) | [optional] [default to null]
**AudioChannels** | **int32** | A number of audio channels. | [optional] [default to null]
**AudioCodec** | **string** | A codec that has been used to encode audio streams. | [optional] [default to null]
**AudioSampleRate** | **int32** | A number of samples of audio carried per second. | [optional] [default to null]
**CreatedAt** | **string** | A date and time when the Video has been created. | [optional] [default to null]
**Duration** | **int32** | A duration of the video in seconds. | [optional] [default to null]
**EncodingsCount** | **int32** | A number of related Encoding objects. | [optional] [default to null]
**ErrorClass** | **string** | A class of an error that has occurred during the encoding process. It is present only if the encoding status is equal to &#x60;fail&#x60;. | [optional] [default to null]
**ErrorMessage** | **string** | A message that explains why the encoding process has resulted in an error. It is present only if the encoding status is equal to &#x60;fail&#x60;. | [optional] [default to null]
**Extname** | **string** | Extension of the source file. | [optional] [default to null]
**FileSize** | **int64** | A size of the source file. | [optional] [default to null]
**Fps** | **float32** | Number of frames per second. | [optional] [default to null]
**Height** | **int32** | Height of the output video. | [optional] [default to null]
**Width** | **int32** | Width of the output video. | [optional] [default to null]
**MimeType** | **string** | A mime type of the source file. | [optional] [default to null]
**OriginalFilename** | **string** | A name of the source file. | [optional] [default to null]
**Path** | **string** |  | [optional] [default to null]
**Payload** | **string** | Payload is an arbitrary text of length 256 or shorter that you can store along the Video. It is typically used to retain an association with one of your own DB record ID. | [optional] [default to null]
**SourceUrl** | **string** | An URL pointing to the source file. | [optional] [default to null]
**Status** | **string** | Determines at what stage of importing process the Video is at the moment. | [optional] [default to null]
**UpdatedAt** | **string** | A date and time when a Video has been updated last time. | [optional] [default to null]
**VideoBitrate** | **string** | video bitrate (in bits/s) | [optional] [default to null]
**VideoCodec** | **string** | A codec that has been used to encode the input file&#39;s video streams. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


