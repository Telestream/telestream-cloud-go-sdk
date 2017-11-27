# UploadSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An unique identifier of the UploadSession. | [default to null]
**Location** | **string** | An URL to which chunks of the uploaded file should be sent | [default to null]
**Parts** | **int32** | A number of chunks that are expected by the upstream. | [optional] [default to null]
**PartSize** | **int32** | An expected size of uploaded chunks. | [optional] [default to null]
**MaxConnections** | **int32** | A maximum number of concurrent connections. | [optional] [default to null]
**ExtraFiles** | [**interface{}**](interface{}.md) | An object containing additional files uploaded using the session. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


