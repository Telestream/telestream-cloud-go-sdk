# Params

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | (for email subscription type);  E-mail addresses  | [optional] [default to null]
**Url** | **string** | (for webhook subscription type);  Webhook URL  | [optional] [default to null]
**Method** | **string** | (for webhook subscription type);  HTTP method; default: POST (GET, POST)  | [optional] [default to null]
**Retries** | **int32** | (for webhook subscription type);  Number of retries before forgetting the notification; default: 0  | [optional] [default to null]
**ContentType** | **string** | (for webhook subscription type); default: application/json (application/json, application/x-www-form-urlencoded)  | [optional] [default to null]
**TopicArn** | **string** | (for sns subscription type) identifier of an AWS SNS Topic that events will be posted to.  | [optional] [default to null]
**RoleArn** | **string** | (for sns subscription type) identifier of an AWS IAM Role that will be used for authorization.  | [optional] [default to null]
**TopicEndpoint** | **string** | (for aeg subscription type) address of an Azure Event Grid Topic that events will be posted to.  | [optional] [default to null]
**AccessKey** | **string** | (for aeg subscription type) secret access key that authorizes Telestream Cloud to write to an Azure Event Grid Topic.  | [optional] [default to null]
**ProjectId** | **string** | (for pubsub subscription type) id of a Google Cloud project that hosts the topic.  | [optional] [default to null]
**TopicName** | **string** | (for pubsub subscription type) name of a Google Cloud Pub/Sub topic to which notifications will be published.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


