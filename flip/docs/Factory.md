# Factory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of a Factory. | [optional] [default to null]
**Name** | **string** | Human-readable identifier of a Factory. | [optional] [default to null]
**FactoryRegion** | **string** | A region where the factory is located. | [optional] [default to null]
**OutputBucketName** | **string** | A bucket where processed files will be stored. | [optional] [default to null]
**Acl** | **string** | Specify if your files are public or private (private files need authorization url to access). By default this is not set. | [optional] [default to null]
**CreatedAt** | **string** | A date and time when a Factory has been created. | [optional] [default to null]
**UpdatedAt** | **string** | A date and time when a Factory has been updated last time. | [optional] [default to null]
**Url** | **string** | An URL pointing to the output_bucket_name. | [optional] [default to null]
**ServerSideEncryption** | **bool** | Specify if you want to use multi-factor server-side 256-bit AES-256 data encryption with Amazon S3-managed encryption keys (SSE-S3). Each object is encrypted using a unique key which as an additional safeguard is encrypted itself with a master key that S3 regularly rotates. By default this is not set. | [optional] [default to null]
**InputBucketName** | **string** | A name of an input bucket. | [optional] [default to null]
**InputBucketWatch** | **bool** | Determines whether the Factory should be notified about new files added to the input bucket. | [optional] [default to null]
**InputBucketFilesMap** | **string** |  | [optional] [default to null]
**InputBucketSyncEveryNMin** | **string** | Determines how often the input bucket is synchronised. | [optional] [default to null]
**InputBucketRecursive** | **string** |  | [optional] [default to null]
**InputBucketFilePattern** | **string** | A pattern that will be used to locate files in the input bucket. Valid wildcards might be used. | [optional] [default to null]
**OutputsPathFormat** | **string** |  | [optional] [default to null]
**StorageProvider** | **int32** | Specifies which storage provider the factory should use. Available options: S3: 0, Google Cloud Storage: 1, FTP storage: 2, Google Cloud Interoperability Storage: 5, Flip storage: 7, FASP storage: 8, Azure Blob Storage: 9 | [optional] [default to null]
**ProviderSpecificSettings** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


