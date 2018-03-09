# Encoding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of an Encoding. | [optional] [default to null]
**AudioBitrate** | **int32** | Audio bitrate (in bits/s). | [optional] [default to null]
**AudioChannels** | **int32** | A number of audio channels. | [optional] [default to null]
**AudioCodec** | **string** | A codec that is used to encode audio streams. | [optional] [default to null]
**AudioSampleRate** | **int32** | A number of samples of audio carried per second. | [optional] [default to null]
**CreatedAt** | **string** | A date and time when the Encoding has been created. | [optional] [default to null]
**Duration** | **int32** |  | [optional] [default to null]
**EncodingProgress** | **int32** |  | [optional] [default to null]
**EncodingTime** | **int32** |  | [optional] [default to null]
**ErrorClass** | **string** | A class of an error that has occurred during the encoding process. It is present only if the encoding status is equal to &#x60;fail&#x60;. | [optional] [default to null]
**ErrorMessage** | **string** | A message that explains why the encoding process has resulted in an error. It is present only if the encoding status is equal to &#x60;fail&#x60;. | [optional] [default to null]
**ExternalId** | **string** |  | [optional] [default to null]
**Extname** | **string** | Extension of the output file. | [optional] [default to null]
**FileSize** | **int32** | A size of the output file. | [optional] [default to null]
**Files** | **[]string** | An array of output file names. | [optional] [default to null]
**Fps** | **float32** | Number of frames per second. | [optional] [default to null]
**Height** | **int32** | Height of the output video. | [optional] [default to null]
**Width** | **int32** | Width of the output video. | [optional] [default to null]
**LogfileUrl** | **string** | An URL pointing to a logfile. | [optional] [default to null]
**MimeType** | **string** | A mime type of the encoded file. | [optional] [default to null]
**ParentEncodingId** | **string** |  | [optional] [default to null]
**Path** | **string** |  | [optional] [default to null]
**ProfileId** | **string** | An id of a related Profile. | [optional] [default to null]
**ProfileName** | **string** | A name of the used Profile. | [optional] [default to null]
**Screenshots** | **[]string** |  | [optional] [default to null]
**StartedEncodingAt** | **string** | A date and time when the encoding process has been started | [optional] [default to null]
**Status** | **string** | Determines at what stage the encoding process is at the moment. | [optional] [default to null]
**UpdatedAt** | **string** | A date and time when a Encoding has been updated last time. | [optional] [default to null]
**VideoBitrate** | **int32** | video bitrate (in bits/s) | [optional] [default to null]
**VideoCodec** | **string** | A codec that is used to encode video streams. | [optional] [default to null]
**VideoId** | **string** | An id of a related Video object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


