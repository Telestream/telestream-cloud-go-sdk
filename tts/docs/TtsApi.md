# \TtsApi

All URIs are relative to *https://api.cloud.telestream.net/tts/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Corpora**](TtsApi.md#Corpora) | **Get** /projects/{projectID}/corpora | Returns a collection of Corpora
[**Corpus**](TtsApi.md#Corpus) | **Get** /projects/{projectID}/corpora/{name} | Returns the Corpus
[**CreateCorpus**](TtsApi.md#CreateCorpus) | **Post** /projects/{projectID}/corpora/{name} | Creates a new Corpus
[**CreateJob**](TtsApi.md#CreateJob) | **Post** /projects/{projectID}/jobs | Creates a new Job
[**CreateProject**](TtsApi.md#CreateProject) | **Post** /projects | Creates a new Project
[**DeleteCorpus**](TtsApi.md#DeleteCorpus) | **Delete** /projects/{projectID}/corpora/{name} | Creates a new Corpus
[**DeleteJob**](TtsApi.md#DeleteJob) | **Delete** /projects/{projectID}/jobs/{jobID} | Deletes the Job
[**DeleteProject**](TtsApi.md#DeleteProject) | **Delete** /projects/{projectID} | Deletes the Project
[**Job**](TtsApi.md#Job) | **Get** /projects/{projectID}/jobs/{jobID} | Returns the Job
[**JobOutputs**](TtsApi.md#JobOutputs) | **Get** /projects/{projectID}/jobs/{jobID}/outputs | Returns the Job Outputs
[**JobResult**](TtsApi.md#JobResult) | **Get** /projects/{projectID}/jobs/{jobID}/result | Returns the Job Result
[**Jobs**](TtsApi.md#Jobs) | **Get** /projects/{projectID}/jobs | Returns a collection of Jobs
[**Project**](TtsApi.md#Project) | **Get** /projects/{projectID} | Returns the Project
[**Projects**](TtsApi.md#Projects) | **Get** /projects | Returns a collection of Projects
[**TrainProject**](TtsApi.md#TrainProject) | **Post** /projects/{projectID}/train | Queues training
[**UpdateProject**](TtsApi.md#UpdateProject) | **Put** /projects/{projectID} | Updates an existing Project
[**UploadVideo**](TtsApi.md#UploadVideo) | **Post** /projects/{projectID}/jobs/upload | Creates an upload session


# **Corpora**
> CorporaCollection Corpora(ctx, projectID)
Returns a collection of Corpora

Returns a collection of Corpora

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 

### Return type

[**CorporaCollection**](CorporaCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Corpus**
> Corpus Corpus(ctx, projectID, name)
Returns the Corpus

Returns the Corpus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **name** | **string**| Corpus name | 

### Return type

[**Corpus**](Corpus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCorpus**
> CreateCorpus(ctx, projectID, name, body)
Creates a new Corpus

Creates a new Corpus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **name** | **string**| Corpus name | 
  **body** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJob**
> Job CreateJob(ctx, projectID, job)
Creates a new Job

Creates a new Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **job** | [**Job**](Job.md)|  | 

### Return type

[**Job**](Job.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> Project CreateProject(ctx, project)
Creates a new Project

Creates a new Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **project** | [**Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCorpus**
> DeleteCorpus(ctx, projectID, name)
Creates a new Corpus

Creates a new Corpus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **name** | **string**| Corpus name | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJob(ctx, projectID, jobID)
Deletes the Job

Deletes the Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **jobID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, )
Deletes the Project

Deletes the Project

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Job**
> Job Job(ctx, projectID, jobID)
Returns the Job

Returns the Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **jobID** | **string**|  | 

### Return type

[**Job**](Job.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobOutputs**
> []JobOutput JobOutputs(ctx, projectID, jobID)
Returns the Job Outputs

Returns the Job Outputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **jobID** | **string**|  | 

### Return type

[**[]JobOutput**](JobOutput.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobResult**
> JobResult JobResult(ctx, projectID, jobID)
Returns the Job Result

Returns the Job Result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **jobID** | **string**|  | 

### Return type

[**JobResult**](JobResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Jobs**
> JobsCollection Jobs(ctx, projectID, optional)
Returns a collection of Jobs

Returns a collection of Jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectID** | **string**| ID of the Project | 
 **page** | **int32**| page number | 
 **perPage** | **int32**| number of records per page | 

### Return type

[**JobsCollection**](JobsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Project**
> Project Project(ctx, projectID)
Returns the Project

Returns the Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Projects**
> ProjectsCollection Projects(ctx, )
Returns a collection of Projects

Returns a collection of Projects

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProjectsCollection**](ProjectsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrainProject**
> TrainProject(ctx, projectID)
Queues training

Queues training

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> Project UpdateProject(ctx, project)
Updates an existing Project

Updates an existing Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **project** | [**Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadVideo**
> UploadSession UploadVideo(ctx, projectID, videoUploadBody)
Creates an upload session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectID** | **string**| ID of the Project | 
  **videoUploadBody** | [**VideoUploadBody**](VideoUploadBody.md)|  | 

### Return type

[**UploadSession**](UploadSession.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

