# \StoresApi

All URIs are relative to *https://api.cloud.telestream.net/stores/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](StoresApi.md#CreateStore) | **Post** /stores | 
[**CreateStoreClientLink**](StoresApi.md#CreateStoreClientLink) | **Post** /stores/{store_id}/service/{service_name}/id/{service_id} | 
[**CreateWatchRule**](StoresApi.md#CreateWatchRule) | **Post** /watch_rules | 
[**DeleteStore**](StoresApi.md#DeleteStore) | **Delete** /stores/{id} | 
[**DeleteStoreClientLink**](StoresApi.md#DeleteStoreClientLink) | **Delete** /stores/{store_id}/service/{service_name}/id/{service_id} | 
[**DeleteWatchRule**](StoresApi.md#DeleteWatchRule) | **Delete** /watch_rules/{id} | 
[**GetObjectUrl**](StoresApi.md#GetObjectUrl) | **Get** /stores/{id}/object_url | 
[**GetStore**](StoresApi.md#GetStore) | **Get** /stores/{id} | 
[**GetStoreIdsForClient**](StoresApi.md#GetStoreIdsForClient) | **Get** /stores/service/{service_name}/id/{service_id} | 
[**GetStores**](StoresApi.md#GetStores) | **Get** /stores | 
[**GetWatchRule**](StoresApi.md#GetWatchRule) | **Get** /watch_rules/{id} | 
[**GetWatchRules**](StoresApi.md#GetWatchRules) | **Get** /watch_rules | 
[**SyncWatchRule**](StoresApi.md#SyncWatchRule) | **Post** /watch_rules/{id}/sync | 
[**UpdateStore**](StoresApi.md#UpdateStore) | **Patch** /stores/{id} | 
[**UpdateWatchRule**](StoresApi.md#UpdateWatchRule) | **Patch** /watch_rules/{id} | 
[**ValidateBucket**](StoresApi.md#ValidateBucket) | **Post** /validate_bucket | 


# **CreateStore**
> Store CreateStore(ctx, storeBody)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeBody** | [**StoreBody**](StoreBody.md)|  | 

### Return type

[**Store**](Store.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoreClientLink**
> CreateStoreClientLink(ctx, storeId, serviceName, serviceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeId** | **string**| an id of a store | 
  **serviceName** | **string**| a name of a client service | 
  **serviceId** | **string**| a service id of a resource that queries for its stores | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWatchRule**
> WatchRule CreateWatchRule(ctx, watchRuleBody)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **watchRuleBody** | [**WatchRule**](WatchRule.md)|  | 

### Return type

[**WatchRule**](WatchRule.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStore**
> DeleteStore(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a store | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoreClientLink**
> DeleteStoreClientLink(ctx, storeId, serviceName, serviceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeId** | **string**| an id of a store | 
  **serviceName** | **string**| a name of a client service | 
  **serviceId** | **string**| a service id of a resource that queries for its stores | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWatchRule**
> DeleteWatchRule(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a resource | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectUrl**
> ObjectUrl GetObjectUrl(ctx, id, path, expiresIn)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a store | 
  **path** | **string**| an path to a file | 
  **expiresIn** | **string**| expiration time in seconds | 

### Return type

[**ObjectUrl**](ObjectURL.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStore**
> Store GetStore(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a store | 

### Return type

[**Store**](Store.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoreIdsForClient**
> []string GetStoreIdsForClient(ctx, serviceName, serviceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceName** | **string**| a name of a client service | 
  **serviceId** | **string**| a service id of a resource that queries for its stores | 

### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStores**
> []Store GetStores(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketName** | **string**| filter stores by bucket_name | 
 **provider** | **string**| filter stores by storage provider | 

### Return type

[**[]Store**](Store.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWatchRule**
> WatchRule GetWatchRule(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a watch rule | 

### Return type

[**WatchRule**](WatchRule.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWatchRules**
> []WatchRule GetWatchRules(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string**| filter watch rules by service_id | 
 **storeId** | **string**| filter watch rules by store_id | 

### Return type

[**[]WatchRule**](WatchRule.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncWatchRule**
> SyncWatchRule(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a watch rule | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStore**
> Store UpdateStore(ctx, id, storeBody)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a store | 
  **storeBody** | [**StoreBody**](StoreBody.md)|  | 

### Return type

[**Store**](Store.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWatchRule**
> WatchRule UpdateWatchRule(ctx, id, watchRuleBody)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| an id of a watch rule | 
  **watchRuleBody** | [**WatchRule**](WatchRule.md)|  | 

### Return type

[**WatchRule**](WatchRule.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateBucket**
> ValidateBucketResponse ValidateBucket(ctx, validateBucketBody)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **validateBucketBody** | [**ValidateBucketBody**](ValidateBucketBody.md)|  | 

### Return type

[**ValidateBucketResponse**](ValidateBucketResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

