# \IamApi

All URIs are relative to *https://api.cloud.telestream.net/iam/v1.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCredential**](IamApi.md#DeleteCredential) | **Delete** /credentials/{id} | 
[**GetCredential**](IamApi.md#GetCredential) | **Get** /credentials/{id} | 
[**GetCredentials**](IamApi.md#GetCredentials) | **Get** /credentials | 
[**GetPolicy**](IamApi.md#GetPolicy) | **Post** /credentials/policy | 
[**PostCredentials**](IamApi.md#PostCredentials) | **Post** /credentials | 
[**UpdateCredential**](IamApi.md#UpdateCredential) | **Patch** /credentials/{id} | 


# **DeleteCredential**
> DeleteCredential(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredential**
> Credential GetCredential(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredentials**
> CredentialsResponse GetCredentials(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> map[string]bool GetPolicy(ctx, policy)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policy** | [**StatementsList**](StatementsList.md)|  | 

### Return type

**map[string]bool**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCredentials**
> Credential PostCredentials(ctx, createCredential)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createCredential** | [**Credential**](Credential.md)|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredential**
> Credential UpdateCredential(ctx, id, updateCredential)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **updateCredential** | [**Credential**](Credential.md)|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

