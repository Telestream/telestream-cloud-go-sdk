# CreateVideoBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUrl** | **string** | An URL pointing to a source file. | [optional] [default to null]
**Profiles** | **string** | Comma-separated list of profile names or IDs to be used during encoding. Alternatively, specify none so no encodings are created yet. | [optional] [default to null]
**Payload** | **string** | Arbitrary string stored along the Video object. | [optional] [default to null]
**Pipeline** | **string** | String-encoded JSON describing profiles pipeline. | [optional] [default to null]
**SubtitleFiles** | **[]string** | A list of urls pointing to remote subtitle files. | [optional] [default to null]
**ExtraFiles** | [**map[string][]string**](array.md) |  | [optional] [default to null]
**ExtraVariables** | **map[string]string** |  | [optional] [default to null]
**PathFormat** | **string** |  | [optional] [default to null]
**ClipEnd** | **string** | Clip ends at a specific time. | [optional] [default to null]
**ClipLength** | **string** | A clip’s duration. | [optional] [default to null]
**ClipOffset** | **string** | Clip starts at a specific offset. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


