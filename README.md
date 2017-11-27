# Go API client for flip

Description

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 3.1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./flip"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:5555/api/flip/3.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FlipApi* | [**CancelEncoding**](docs/FlipApi.md#cancelencoding) | **Post** /encodings/{id}/cancel.json | Cancels an Encoding.
*FlipApi* | [**CopyProfile**](docs/FlipApi.md#copyprofile) | **Post** /profiles/{id}/copy.json | Copies a given Profile
*FlipApi* | [**CreateEncoding**](docs/FlipApi.md#createencoding) | **Post** /encodings.json | Creates an Encoding
*FlipApi* | [**CreateFactory**](docs/FlipApi.md#createfactory) | **Post** /factories.json | Creates a new factory
*FlipApi* | [**CreateProfile**](docs/FlipApi.md#createprofile) | **Post** /profiles.json | Creates a Profile
*FlipApi* | [**CreateWorkorder**](docs/FlipApi.md#createworkorder) | **Post** /workorders.json | Creates a Workorder.
*FlipApi* | [**DeleteEncoding**](docs/FlipApi.md#deleteencoding) | **Delete** /encodings/{id}.json | Deletes an Encoding from both Telestream Cloud and your storage. Returns an information whether the operation was successful.
*FlipApi* | [**DeleteProfile**](docs/FlipApi.md#deleteprofile) | **Delete** /profiles/{id}.json | Deletes a given Profile
*FlipApi* | [**DeleteVideo**](docs/FlipApi.md#deletevideo) | **Delete** /videos/{id}.json | Deletes a Video object.
*FlipApi* | [**DeleteVideoSource**](docs/FlipApi.md#deletevideosource) | **Delete** /videos/{id}/source.json | Delete a video&#39;s source file.
*FlipApi* | [**Encoding**](docs/FlipApi.md#encoding) | **Get** /encodings/{id}.json | Returns an Encoding object.
*FlipApi* | [**Encodings**](docs/FlipApi.md#encodings) | **Get** /encodings.json | Returns a list of Encoding objects
*FlipApi* | [**EncodingsCount**](docs/FlipApi.md#encodingscount) | **Get** /encodings/count.json | Returns a number of Encoding objects created using a given factory.
*FlipApi* | [**Factories**](docs/FlipApi.md#factories) | **Get** /factories.json | Returns a collection of Factory objects.
*FlipApi* | [**Factory**](docs/FlipApi.md#factory) | **Get** /factories/{id}.json | Returns a Factory object.
*FlipApi* | [**Notifications**](docs/FlipApi.md#notifications) | **Get** /notifications.json | Returns a Factory&#39;s notification settings.
*FlipApi* | [**Profile**](docs/FlipApi.md#profile) | **Get** /profiles/{id_or_name}.json | Returns a Profile object.
*FlipApi* | [**ProfileEncodings**](docs/FlipApi.md#profileencodings) | **Get** /profiles/{id_or_name}/encodings.json | Returns a list of Encodings that belong to a Profile.
*FlipApi* | [**Profiles**](docs/FlipApi.md#profiles) | **Get** /profiles.json | Returns a collection of Profile objects.
*FlipApi* | [**QueuedVideos**](docs/FlipApi.md#queuedvideos) | **Get** /videos/queued.json | Returns a collection of Video objects queued for encoding.
*FlipApi* | [**ResubmitVideo**](docs/FlipApi.md#resubmitvideo) | **Post** /videos/resubmit.json | Resubmits a video to encode.
*FlipApi* | [**RetryEncoding**](docs/FlipApi.md#retryencoding) | **Post** /encodings/{id}/retry.json | Retries a failed encoding.
*FlipApi* | [**SignedEncodingUrl**](docs/FlipApi.md#signedencodingurl) | **Get** /encodings/{id}/signed-url.json | Returns a signed url pointing to an Encoding.
*FlipApi* | [**SignedEncodingUrls**](docs/FlipApi.md#signedencodingurls) | **Get** /encodings/{id}/signed-urls.json | Returns a list of signed urls pointing to an Encoding&#39;s outputs.
*FlipApi* | [**SignedVideoUrl**](docs/FlipApi.md#signedvideourl) | **Get** /videos/{id}/signed-url.json | Returns a signed url pointing to a Video.
*FlipApi* | [**ToggleFactorySync**](docs/FlipApi.md#togglefactorysync) | **Post** /factories/{id}/sync.json | Toggles synchronisation settings.
*FlipApi* | [**UpdateEncoding**](docs/FlipApi.md#updateencoding) | **Put** /encodings/{id}.json | Updates an Encoding
*FlipApi* | [**UpdateFactory**](docs/FlipApi.md#updatefactory) | **Patch** /factories/{id}.json | Updates a Factory&#39;s settings. Returns a Factory object.
*FlipApi* | [**UpdateNotifications**](docs/FlipApi.md#updatenotifications) | **Put** /notifications.json | Updates a Factory&#39;s notification settings.
*FlipApi* | [**UpdateProfile**](docs/FlipApi.md#updateprofile) | **Put** /profiles/{id}.json | Updates a given Profile
*FlipApi* | [**UploadVideo**](docs/FlipApi.md#uploadvideo) | **Post** /videos/upload.json | Creates an upload session.
*FlipApi* | [**Video**](docs/FlipApi.md#video) | **Get** /videos/{id}.json | Returns a Video object.
*FlipApi* | [**VideoEncodings**](docs/FlipApi.md#videoencodings) | **Get** /videos/{id}/encodings.json | Returns a list of Encodings that belong to a Video.
*FlipApi* | [**VideoMetadata**](docs/FlipApi.md#videometadata) | **Get** /videos/{id}/metadata.json | Returns a Video&#39;s metadata
*FlipApi* | [**Videos**](docs/FlipApi.md#videos) | **Get** /videos.json | Returns a collection of Video objects.
*FlipApi* | [**Workflows**](docs/FlipApi.md#workflows) | **Get** /workflows.json | Returns a collection of Workflows that belong to a Factory.


## Documentation For Models

 - [CanceledResponse](docs/CanceledResponse.md)
 - [CloudNotificationSettings](docs/CloudNotificationSettings.md)
 - [CloudNotificationSettingsEvents](docs/CloudNotificationSettingsEvents.md)
 - [CopyProfileBody](docs/CopyProfileBody.md)
 - [CountResponse](docs/CountResponse.md)
 - [CreateEncodingBody](docs/CreateEncodingBody.md)
 - [DeletedResponse](docs/DeletedResponse.md)
 - [Encoding](docs/Encoding.md)
 - [EncodingSignedUrl](docs/EncodingSignedUrl.md)
 - [EncodingSignedUrls](docs/EncodingSignedUrls.md)
 - [ExtraFile](docs/ExtraFile.md)
 - [Factory](docs/Factory.md)
 - [FactoryBody](docs/FactoryBody.md)
 - [FactoryBodyStorageCredentialAttributes](docs/FactoryBodyStorageCredentialAttributes.md)
 - [FactorySync](docs/FactorySync.md)
 - [FactorySyncBody](docs/FactorySyncBody.md)
 - [ModelError](docs/ModelError.md)
 - [PaginatedEncodingsCollection](docs/PaginatedEncodingsCollection.md)
 - [PaginatedFactoryCollection](docs/PaginatedFactoryCollection.md)
 - [PaginatedProfilesCollection](docs/PaginatedProfilesCollection.md)
 - [PaginatedVideoCollection](docs/PaginatedVideoCollection.md)
 - [PaginatedWorkflowsCollection](docs/PaginatedWorkflowsCollection.md)
 - [Profile](docs/Profile.md)
 - [ProfileBody](docs/ProfileBody.md)
 - [ResubmitVideoBody](docs/ResubmitVideoBody.md)
 - [RetriedResponse](docs/RetriedResponse.md)
 - [SignedVideoUrl](docs/SignedVideoUrl.md)
 - [UpdateEncodingBody](docs/UpdateEncodingBody.md)
 - [UploadSession](docs/UploadSession.md)
 - [Video](docs/Video.md)
 - [VideoMetadata](docs/VideoMetadata.md)
 - [VideoUploadBody](docs/VideoUploadBody.md)


## Documentation For Authorization


## api_key

- **Type**: API key 
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header


## Author

cloudsupport@telestream.net

