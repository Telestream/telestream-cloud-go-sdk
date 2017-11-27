# FlipApi

All URIs are relative to *http://localhost:5555/api/flip/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelEncoding**](FlipApi.md#CancelEncoding) | **Post** /encodings/{id}/cancel.json | Cancels an Encoding.
[**CopyProfile**](FlipApi.md#CopyProfile) | **Post** /profiles/{id}/copy.json | Copies a given Profile
[**CreateEncoding**](FlipApi.md#CreateEncoding) | **Post** /encodings.json | Creates an Encoding
[**CreateFactory**](FlipApi.md#CreateFactory) | **Post** /factories.json | Creates a new factory
[**CreateProfile**](FlipApi.md#CreateProfile) | **Post** /profiles.json | Creates a Profile
[**CreateWorkorder**](FlipApi.md#CreateWorkorder) | **Post** /workorders.json | Creates a Workorder.
[**DeleteEncoding**](FlipApi.md#DeleteEncoding) | **Delete** /encodings/{id}.json | Deletes an Encoding from both Telestream Cloud and your storage. Returns an information whether the operation was successful.
[**DeleteProfile**](FlipApi.md#DeleteProfile) | **Delete** /profiles/{id}.json | Deletes a given Profile
[**DeleteVideo**](FlipApi.md#DeleteVideo) | **Delete** /videos/{id}.json | Deletes a Video object.
[**DeleteVideoSource**](FlipApi.md#DeleteVideoSource) | **Delete** /videos/{id}/source.json | Delete a video&#39;s source file.
[**Encoding**](FlipApi.md#Encoding) | **Get** /encodings/{id}.json | Returns an Encoding object.
[**Encodings**](FlipApi.md#Encodings) | **Get** /encodings.json | Returns a list of Encoding objects
[**EncodingsCount**](FlipApi.md#EncodingsCount) | **Get** /encodings/count.json | Returns a number of Encoding objects created using a given factory.
[**Factories**](FlipApi.md#Factories) | **Get** /factories.json | Returns a collection of Factory objects.
[**Factory**](FlipApi.md#Factory) | **Get** /factories/{id}.json | Returns a Factory object.
[**Notifications**](FlipApi.md#Notifications) | **Get** /notifications.json | Returns a Factory&#39;s notification settings.
[**Profile**](FlipApi.md#Profile) | **Get** /profiles/{id_or_name}.json | Returns a Profile object.
[**ProfileEncodings**](FlipApi.md#ProfileEncodings) | **Get** /profiles/{id_or_name}/encodings.json | Returns a list of Encodings that belong to a Profile.
[**Profiles**](FlipApi.md#Profiles) | **Get** /profiles.json | Returns a collection of Profile objects.
[**QueuedVideos**](FlipApi.md#QueuedVideos) | **Get** /videos/queued.json | Returns a collection of Video objects queued for encoding.
[**ResubmitVideo**](FlipApi.md#ResubmitVideo) | **Post** /videos/resubmit.json | Resubmits a video to encode.
[**RetryEncoding**](FlipApi.md#RetryEncoding) | **Post** /encodings/{id}/retry.json | Retries a failed encoding.
[**SignedEncodingUrl**](FlipApi.md#SignedEncodingUrl) | **Get** /encodings/{id}/signed-url.json | Returns a signed url pointing to an Encoding.
[**SignedEncodingUrls**](FlipApi.md#SignedEncodingUrls) | **Get** /encodings/{id}/signed-urls.json | Returns a list of signed urls pointing to an Encoding&#39;s outputs.
[**SignedVideoUrl**](FlipApi.md#SignedVideoUrl) | **Get** /videos/{id}/signed-url.json | Returns a signed url pointing to a Video.
[**ToggleFactorySync**](FlipApi.md#ToggleFactorySync) | **Post** /factories/{id}/sync.json | Toggles synchronisation settings.
[**UpdateEncoding**](FlipApi.md#UpdateEncoding) | **Put** /encodings/{id}.json | Updates an Encoding
[**UpdateFactory**](FlipApi.md#UpdateFactory) | **Patch** /factories/{id}.json | Updates a Factory&#39;s settings. Returns a Factory object.
[**UpdateNotifications**](FlipApi.md#UpdateNotifications) | **Put** /notifications.json | Updates a Factory&#39;s notification settings.
[**UpdateProfile**](FlipApi.md#UpdateProfile) | **Put** /profiles/{id}.json | Updates a given Profile
[**UploadVideo**](FlipApi.md#UploadVideo) | **Post** /videos/upload.json | Creates an upload session.
[**Video**](FlipApi.md#Video) | **Get** /videos/{id}.json | Returns a Video object.
[**VideoEncodings**](FlipApi.md#VideoEncodings) | **Get** /videos/{id}/encodings.json | Returns a list of Encodings that belong to a Video.
[**VideoMetadata**](FlipApi.md#VideoMetadata) | **Get** /videos/{id}/metadata.json | Returns a Video&#39;s metadata
[**Videos**](FlipApi.md#Videos) | **Get** /videos.json | Returns a collection of Video objects.
[**Workflows**](FlipApi.md#Workflows) | **Get** /workflows.json | Returns a collection of Workflows that belong to a Factory.


# **CancelEncoding**
> CanceledResponse CancelEncoding($id, $factoryId)

Cancels an Encoding.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**CanceledResponse**](CanceledResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyProfile**
> Profile CopyProfile($id, $factoryId, $copyProfileBody, $expand)

Copies a given Profile


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Profile. |
 **factoryId** | **string**| Id of a Factory. |
 **copyProfileBody** | [**CopyProfileBody**](CopyProfileBody.md)|  |
 **expand** | **bool**| If expand option is set Profile objects will contain all command parameters, even if their value is default. By default this is not set. | [optional]

### Return type

[**Profile**](Profile.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEncoding**
> Encoding CreateEncoding($factoryId, $createEncodingBody, $screenshots, $preciseStatus)

Creates an Encoding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **createEncodingBody** | [**CreateEncodingBody**](CreateEncodingBody.md)|  |
 **screenshots** | **bool**| Determines whether the response will include screenshots. By default this is not set. | [optional]
 **preciseStatus** | **bool**| Determines whether the response will include a precise status. By default this is not set. | [optional]

### Return type

[**Encoding**](Encoding.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFactory**
> Factory CreateFactory($createFactoryBody, $withStorageProvider)

Creates a new factory


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFactoryBody** | [**FactoryBody**](FactoryBody.md)|  |
 **withStorageProvider** | **bool**| if set to &#x60;true&#x60;, results will include a storage provider&#39;s id | [optional]

### Return type

[**Factory**](Factory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProfile**
> Profile CreateProfile($factoryId, $createProfileBody, $excludeAdvancedServices, $expand)

Creates a Profile


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **createProfileBody** | [**ProfileBody**](ProfileBody.md)|  |
 **excludeAdvancedServices** | **bool**|  | [optional]
 **expand** | **bool**| If expand option is set Profile objects will contain all command parameters, even if their value is default. By default it is not set. | [optional]

### Return type

[**Profile**](Profile.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWorkorder**
> CreateWorkorder($factoryId, $profileId, $file, $sourceUrl)

Creates a Workorder.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **profileId** | **string**| Id of a Profile. | [optional]
 **file** | ***os.File**| Input file. | [optional]
 **sourceUrl** | **string**| URL pointing to an input file. | [optional]

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEncoding**
> DeletedResponse DeleteEncoding($id, $factoryId)

Deletes an Encoding from both Telestream Cloud and your storage. Returns an information whether the operation was successful.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**DeletedResponse**](DeletedResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProfile**
> DeletedResponse DeleteProfile($id, $factoryId)

Deletes a given Profile


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Profile |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**DeletedResponse**](DeletedResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideo**
> DeletedResponse DeleteVideo($id, $factoryId)

Deletes a Video object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**DeletedResponse**](DeletedResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoSource**
> DeletedResponse DeleteVideoSource($id, $factoryId)

Delete a video's source file.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**DeletedResponse**](DeletedResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Encoding**
> Encoding Encoding($id, $factoryId, $screenshots, $preciseStatus)

Returns an Encoding object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |
 **screenshots** | **bool**| Determines whether the response will include screenshots. By default this is not set. | [optional]
 **preciseStatus** | **bool**| Determines whether the response will include a precise status. By default this is not set. | [optional]

### Return type

[**Encoding**](Encoding.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Encodings**
> PaginatedEncodingsCollection Encodings($factoryId, $videoId, $status, $profileId, $profileName, $page, $perPage, $screenshots, $preciseStatus)

Returns a list of Encoding objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **videoId** | **string**| Id of a Video. When specified, the resulting list will contain videos that belong to the Video. | [optional]
 **status** | **string**| One of &#x60;success&#x60;, &#x60;fail&#x60;, &#x60;processing&#x60;. When specified, the resulting list will contain ecodings filtered by status. | [optional]
 **profileId** | **string**| Filter by profile_id. | [optional]
 **profileName** | **string**| Filter by profile_name. | [optional]
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]
 **screenshots** | **bool**| Determines whether the response will include screenshots. By default this is not set. | [optional]
 **preciseStatus** | **bool**| Determines whether the response will include a precise status. By default this is not set. | [optional]

### Return type

[**PaginatedEncodingsCollection**](PaginatedEncodingsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EncodingsCount**
> CountResponse EncodingsCount($factoryId)

Returns a number of Encoding objects created using a given factory.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**CountResponse**](CountResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Factories**
> PaginatedFactoryCollection Factories($page, $perPage, $withStorageProvider)

Returns a collection of Factory objects.

Returns a collection of Factory objects.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]
 **withStorageProvider** | **bool**| if set to &#x60;true&#x60;, results will include a storage provider&#39;s id | [optional]

### Return type

[**PaginatedFactoryCollection**](PaginatedFactoryCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Factory**
> Factory Factory($id, $withStorageProvider)

Returns a Factory object.

Returns a Factory object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id of a factory |
 **withStorageProvider** | **bool**| if set to &#x60;true&#x60;, results will include a storage provider&#39;s id | [optional]

### Return type

[**Factory**](Factory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Notifications**
> CloudNotificationSettings Notifications($factoryId)

Returns a Factory's notification settings.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**CloudNotificationSettings**](CloudNotificationSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Profile**
> Profile Profile($idOrName, $factoryId, $expand)

Returns a Profile object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idOrName** | **string**| A name or an id of a Profile. |
 **factoryId** | **string**| Id of a Factory. |
 **expand** | **bool**| If expand option is set Profile objects will contain all command parameters, even if their value is default. By default this is not set. | [optional]

### Return type

[**Profile**](Profile.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileEncodings**
> PaginatedEncodingsCollection ProfileEncodings($idOrName, $factoryId)

Returns a list of Encodings that belong to a Profile.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idOrName** | **string**| Id or name of a Profile. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**PaginatedEncodingsCollection**](PaginatedEncodingsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Profiles**
> PaginatedProfilesCollection Profiles($factoryId, $excludeAdvancedServices, $expand, $page, $perPage)

Returns a collection of Profile objects.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **excludeAdvancedServices** | **bool**| Determine whether exclude Advanced Services profiles from the results. By default this is not set. | [optional]
 **expand** | **bool**| If expand option is set Profile objects will contain all command parameters, even if their value is default. By default this is not set. | [optional]
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]

### Return type

[**PaginatedProfilesCollection**](PaginatedProfilesCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueuedVideos**
> PaginatedVideoCollection QueuedVideos($factoryId, $page, $perPage)

Returns a collection of Video objects queued for encoding.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]

### Return type

[**PaginatedVideoCollection**](PaginatedVideoCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResubmitVideo**
> ResubmitVideo($factoryId, $resubmitVideoBody)

Resubmits a video to encode.

Resubmits the video to encode. Please note that this option will work only for videos in `success` status.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **resubmitVideoBody** | [**ResubmitVideoBody**](ResubmitVideoBody.md)|  |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryEncoding**
> RetriedResponse RetryEncoding($id, $factoryId)

Retries a failed encoding.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**RetriedResponse**](RetriedResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignedEncodingUrl**
> EncodingSignedUrl SignedEncodingUrl($id, $factoryId)

Returns a signed url pointing to an Encoding.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**EncodingSignedUrl**](EncodingSignedUrl.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignedEncodingUrls**
> EncodingSignedUrls SignedEncodingUrls($id, $factoryId)

Returns a list of signed urls pointing to an Encoding's outputs.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**EncodingSignedUrls**](EncodingSignedUrls.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignedVideoUrl**
> SignedVideoUrl SignedVideoUrl($id, $factoryId)

Returns a signed url pointing to a Video.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**SignedVideoUrl**](SignedVideoUrl.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleFactorySync**
> FactorySync ToggleFactorySync($id, $factorySyncBody)

Toggles synchronisation settings.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id of the factory |
 **factorySyncBody** | [**FactorySyncBody**](FactorySyncBody.md)|  |

### Return type

[**FactorySync**](FactorySync.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEncoding**
> Encoding UpdateEncoding($id, $factoryId, $updateEncodingBody, $screenshots, $preciseStatus)

Updates an Encoding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of an Encoding. |
 **factoryId** | **string**| Id of a Factory. |
 **updateEncodingBody** | [**UpdateEncodingBody**](UpdateEncodingBody.md)|  |
 **screenshots** | **bool**| Determines whether the response will include screenshots. By default this is not set. | [optional]
 **preciseStatus** | **bool**| Determines whether the response will include a precise status. By default this is not set. | [optional]

### Return type

[**Encoding**](Encoding.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFactory**
> Factory UpdateFactory($id, $updateFactoryBody, $withStorageProvider)

Updates a Factory's settings. Returns a Factory object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id of the factory |
 **updateFactoryBody** | [**FactoryBody**](FactoryBody.md)|  |
 **withStorageProvider** | **bool**| if set to &#x60;true&#x60;, results will include a storage provider&#39;s id | [optional]

### Return type

[**Factory**](Factory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotifications**
> CloudNotificationSettings UpdateNotifications($factoryId, $cloudNotificationSettingsBody)

Updates a Factory's notification settings.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **cloudNotificationSettingsBody** | [**CloudNotificationSettings**](CloudNotificationSettings.md)|  |

### Return type

[**CloudNotificationSettings**](CloudNotificationSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfile**
> Profile UpdateProfile($id, $factoryId, $updateProfileBody, $excludeAdvancedServices, $expand)

Updates a given Profile


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  |
 **factoryId** | **string**| Id of a Factory. |
 **updateProfileBody** | [**ProfileBody**](ProfileBody.md)|  |
 **excludeAdvancedServices** | **bool**|  | [optional]
 **expand** | **bool**| If expand option is set Profile objects will contain all command parameters, even if their value is default. By default this is not set. | [optional]

### Return type

[**Profile**](Profile.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadVideo**
> UploadSession UploadVideo($factoryId, $videoUploadBody)

Creates an upload session.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **videoUploadBody** | [**VideoUploadBody**](VideoUploadBody.md)|  |

### Return type

[**UploadSession**](UploadSession.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Video**
> Video Video($id, $factoryId)

Returns a Video object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**Video**](Video.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoEncodings**
> PaginatedEncodingsCollection VideoEncodings($id, $factoryId, $page, $perPage, $screenshots, $preciseStatus)

Returns a list of Encodings that belong to a Video.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]
 **screenshots** | **bool**| Determines whether the response will include screenshots. By default this is not set. | [optional]
 **preciseStatus** | **bool**| Determines whether the response will include a precise status. By default this is not set. | [optional]

### Return type

[**PaginatedEncodingsCollection**](PaginatedEncodingsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoMetadata**
> VideoMetadata VideoMetadata($id, $factoryId)

Returns a Video's metadata


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of a Video. |
 **factoryId** | **string**| Id of a Factory. |

### Return type

[**VideoMetadata**](VideoMetadata.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Videos**
> PaginatedVideoCollection Videos($factoryId, $page, $perPage)

Returns a collection of Video objects.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]

### Return type

[**PaginatedVideoCollection**](PaginatedVideoCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Workflows**
> PaginatedWorkflowsCollection Workflows($factoryId, $page, $perPage)

Returns a collection of Workflows that belong to a Factory.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factoryId** | **string**| Id of a Factory. |
 **page** | **int32**| A page to be fetched. Default is &#x60;1&#x60;. | [optional]
 **perPage** | **int32**| A number of results per page. Default is &#x60;100&#x60;. | [optional]

### Return type

[**PaginatedWorkflowsCollection**](PaginatedWorkflowsCollection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
