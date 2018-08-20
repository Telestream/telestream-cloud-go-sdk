# Topic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | [read-only] Account identifier connected to subscription notification  | [optional] [default to null]
**Service** | **string** | Name of service (qc, flip, tts)  | [default to null]
**Event** | **string** | Name of the event;  Quality Control (media-passed, media-error, media-warning, media-rejected, media-canceled, job-created, job-progress), Flip (video-created, video-encoded, encoding-complete, encoding-progress), Transcription (job-created, job-completed, job-error, job-progress)  | [default to null]
**Project** | **string** | (for tts, qc service); Project ID  | [optional] [default to null]
**Factory** | **string** | (for flip service); Factory ID  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


