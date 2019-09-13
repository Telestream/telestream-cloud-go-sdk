# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSyncEvents** | **int32** | Total number of synchronization detection events considered for summary. | [optional] 
**AvsyncAnalysis** | **string** |  | [optional] 
**AvsyncInference** | **string** |  | [optional] 
**Confidence** | **float32** | Confidence of avsync_inference and avsync_analysis. &#39;confidence&#39; is a value between 0 and 100. confidence &gt;&#x3D; 90  - Very high confidence. 80 &lt;&#x3D; confidence &lt; 90   - High confidence. 70 &lt;&#x3D; confidence &lt; 80   - Should be reviewed by human.  | [optional] 
**Skew** | **int32** | Skew estimate in #frames; Skew is the number of frames that audio is delayed to the video | [optional] 
**TotalLength** | **float32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


