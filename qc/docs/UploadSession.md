# UploadSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An unique identifier of the UploadSession. | 
**Location** | **string** | An URL to which chunks of the uploaded file should be sent | 
**Parts** | **int32** | A number of chunks that are expected by the upstream. | [optional] 
**PartSize** | **int32** | An expected size of uploaded chunks. | [optional] 
**MaxConnections** | **int32** | A maximum number of concurrent connections. | [optional] 
**ExtraFiles** | [**map[string]interface{}**](.md) | An object containing additional files uploaded using the session. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


