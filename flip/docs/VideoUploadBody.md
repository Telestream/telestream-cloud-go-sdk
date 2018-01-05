# VideoUploadBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSize** | **int64** | Size of the file that will be uploaded in &#x60;bytes&#x60;. | [default to null]
**FileName** | **string** | Name of the file that will be uploaded. | [default to null]
**ExtraFiles** | [**[]ExtraFile**](ExtraFile.md) | A list of names of additional files that will be uploaded. | [optional] [default to null]
**Profiles** | **string** | A comma-separated list of profile names or IDs to be used during encoding. Alternatively, specify none so no encodings will created right away. | [optional] [default to null]
**PathFormat** | **string** |  | [optional] [default to null]
**Payload** | **string** | Payload is an arbitrary text of length 256 or shorter that you can store along the Video. It is typically used to retain an association with one of your own DB record ID. | [optional] [default to null]
**ExtraVariables** | **map[string]string** |  | [optional] [default to null]
**WatermarkUrl** | **string** | URL pointing to an image that will be used asa watermark. | [optional] [default to null]
**WatermarkLeft** | **string** | Determines distance between the left edge of a video and the left edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_right is not. | [optional] [default to null]
**WatermarkTop** | **string** | Determines distance between the top edge of a video and the top edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_bottom is not. | [optional] [default to null]
**WatermarkRight** | **string** | Determines distance between the right edge of a video and the right edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_left is not. | [optional] [default to null]
**WatermarkBottom** | **string** | Determines distance between the bottom edge of a video and the bottom edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_top is not. | [optional] [default to null]
**WatermarkWidth** | **string** | Determines width of the watermark image. Should be specified in pixels. | [optional] [default to null]
**WatermarkHeight** | **string** | Determines width of the watermark image. Should be specified in pixels. | [optional] [default to null]
**ClipLength** | **string** | Length of the uploaded video. Should be formatted as follows: HH:MM:SS | [optional] [default to null]
**ClipOffset** | **string** | Clip starts at a specific offset. | [optional] [default to null]
**MultiChunk** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


