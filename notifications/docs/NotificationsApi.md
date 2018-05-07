# \NotificationsApi

All URIs are relative to *https://api.cloud.telestream.net/notifications/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](NotificationsApi.md#CreateSubscription) | **Post** /subscriptions | Create a new subscription
[**DeleteSubscription**](NotificationsApi.md#DeleteSubscription) | **Delete** /subscriptions/{id} | 
[**ListSubscriptions**](NotificationsApi.md#ListSubscriptions) | **Get** /subscriptions | 


# **CreateSubscription**
> CreateSubscription(ctx, optional)
Create a new subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription** | [**Subscription**](Subscription.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscription**
> DeleteSubscription(ctx, id)


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

# **ListSubscriptions**
> []Subscription ListSubscriptions(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscription**](subscription.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

