# CreateVideoBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUrl** | **string** | An URL pointing to a source file. | [optional] [default to null]
**Profiles** | **string** | Comma-separated list of profile names or IDs to be used during encoding. Alternatively, specify none so no encodings are created yet. | [optional] [default to null]
**Payload** | **string** | Payload is an arbitrary text of length 256 or shorter that you can store along the Video. It is typically used to retain an association with one of your own DB record ID. | [optional] [default to null]
**Pipeline** | **string** | String-encoded JSON describing profiles pipeline. | [optional] [default to null]
**SubtitleFiles** | **[]string** | A list of urls pointing to remote subtitle files. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


