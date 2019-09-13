# \QcApi

All URIs are relative to *https://api.cloud.telestream.net/qc/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](QcApi.md#CancelJob) | **Put** /projects/{project_id}/jobs/{job_id}/cancel.json | Cancel QC job
[**CreateJob**](QcApi.md#CreateJob) | **Post** /projects/{project_id}/jobs.json | Create a new job
[**CreateProject**](QcApi.md#CreateProject) | **Post** /projects.json | Create a new project
[**GetJob**](QcApi.md#GetJob) | **Get** /projects/{project_id}/jobs/{job_id}.json | Get QC job
[**GetProject**](QcApi.md#GetProject) | **Get** /projects/{project_id}.json | Get project by Id
[**ImportTemplate**](QcApi.md#ImportTemplate) | **Post** /projects/import.json | Import Vidchecker template
[**ListJobs**](QcApi.md#ListJobs) | **Get** /projects/{project_id}/jobs.json | Get jobs form projects
[**ListProjects**](QcApi.md#ListProjects) | **Get** /projects.json | List all projects for an account
[**ModifyProject**](QcApi.md#ModifyProject) | **Put** /projects/{project_id}.json | Modify project
[**Proxy**](QcApi.md#Proxy) | **Get** /projects/{project_id}/jobs/{job_id}/proxy.json | 
[**RemoveJob**](QcApi.md#RemoveJob) | **Delete** /projects/{project_id}/jobs/{job_id}.json | Remove QC job
[**RemoveProject**](QcApi.md#RemoveProject) | **Delete** /projects/{project_id}.json | Remove project
[**SignedUrls**](QcApi.md#SignedUrls) | **Get** /projects/{project_id}/jobs/{job_id}/signed-urls.json | Get QC job signed urls
[**Templates**](QcApi.md#Templates) | **Get** /templates.json | List all templates
[**UploadVideo**](QcApi.md#UploadVideo) | **Post** /projects/{project_id}/upload.json | Creates an upload session



## CancelJob

> CancelJob(ctx, projectId, jobId)
Cancel QC job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**jobId** | **string**| A unique identifier of a Job. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> Job CreateJob(ctx, projectId, vidChecker8JobData)
Create a new job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**vidChecker8JobData** | [**VidChecker8JobData**](VidChecker8JobData.md)|  | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> Project CreateProject(ctx, vidChecker8Body)
Create a new project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vidChecker8Body** | [**VidChecker8Body**](VidChecker8Body.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> Job GetJob(ctx, projectId, jobId)
Get QC job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**jobId** | **string**| A unique identifier of a Job. | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, projectId)
Get project by Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 

### Return type

[**Project**](Project.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTemplate

> []InlineResponse200 ImportTemplate(ctx, optional)
Import Vidchecker template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **body** | **optional.String**|  | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobs

> JobsCollection ListJobs(ctx, projectId, optional)
Get jobs form projects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
 **optional** | ***ListJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListJobsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.Bool**| Expand details of job | 
 **status** | **optional.String**| Filter jobs by status | 
 **perPage** | **optional.Int32**| Limit number of listed jobs | [default to 30]
 **page** | **optional.Int32**| Index of jobs page to be listed | 

### Return type

[**JobsCollection**](JobsCollection.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> []Project ListProjects(ctx, )
List all projects for an account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Project**](Project.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProject

> Project ModifyProject(ctx, projectId, modifyVidChecker8Body)
Modify project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**modifyVidChecker8Body** | [**ModifyVidChecker8Body**](ModifyVidChecker8Body.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Proxy

> Proxy Proxy(ctx, projectId, jobId)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**jobId** | **string**| A unique identifier of a Job. | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveJob

> RemoveJob(ctx, projectId, jobId)
Remove QC job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**jobId** | **string**| A unique identifier of a Job. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProject

> RemoveProject(ctx, projectId)
Remove project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignedUrls

> map[string]string SignedUrls(ctx, projectId, jobId)
Get QC job signed urls

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**jobId** | **string**| A unique identifier of a Job. | 

### Return type

**map[string]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Templates

> []Template Templates(ctx, )
List all templates

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Template**](Template.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVideo

> UploadSession UploadVideo(ctx, projectId, videoUploadBody)
Creates an upload session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| A unique identifier of a Project. | 
**videoUploadBody** | [**VideoUploadBody**](VideoUploadBody.md)|  | 

### Return type

[**UploadSession**](UploadSession.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

