# Telestream Cloud Quality Control Go SDK

This library provides a low-level interface to the REST API of Telestream Cloud, the online video encoding service.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 3.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Documentation for API Endpoints

All URIs are relative to *https://api.cloud.telestream.net/qc/v1.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*QcApi* | [**CancelJob**](docs/QcApi.md#canceljob) | **Put** /projects/{project_id}/jobs/{job_id}/cancel.json | Cancel QC job
*QcApi* | [**CreateJob**](docs/QcApi.md#createjob) | **Post** /projects/{project_id}/jobs.json | Create a new job
*QcApi* | [**CreateProject**](docs/QcApi.md#createproject) | **Post** /projects.json | Create a new project
*QcApi* | [**GetJob**](docs/QcApi.md#getjob) | **Get** /projects/{project_id}/jobs/{job_id}.json | Get QC job
*QcApi* | [**GetProject**](docs/QcApi.md#getproject) | **Get** /projects/{project_id}.json | Get project by Id
*QcApi* | [**ImportTemplate**](docs/QcApi.md#importtemplate) | **Post** /projects/import.json | Import Vidchecker template
*QcApi* | [**ListJobs**](docs/QcApi.md#listjobs) | **Get** /projects/{project_id}/jobs.json | Get jobs form projects
*QcApi* | [**ListProjects**](docs/QcApi.md#listprojects) | **Get** /projects.json | List all projects for an account
*QcApi* | [**ModifyProject**](docs/QcApi.md#modifyproject) | **Put** /projects/{project_id}.json | Modify project
*QcApi* | [**Proxy**](docs/QcApi.md#proxy) | **Get** /projects/{project_id}/jobs/{job_id}/proxy.json | 
*QcApi* | [**RemoveJob**](docs/QcApi.md#removejob) | **Delete** /projects/{project_id}/jobs/{job_id}.json | Remove QC job
*QcApi* | [**RemoveProject**](docs/QcApi.md#removeproject) | **Delete** /projects/{project_id}.json | Remove project
*QcApi* | [**SignedUrls**](docs/QcApi.md#signedurls) | **Get** /projects/{project_id}/jobs/{job_id}/signed-urls.json | Get QC job signed urls
*QcApi* | [**Templates**](docs/QcApi.md#templates) | **Get** /templates.json | List all templates
*QcApi* | [**UploadVideo**](docs/QcApi.md#uploadvideo) | **Post** /projects/{project_id}/upload.json | Creates an upload session


## Documentation For Models

 - [ActiveFormatDescriptorTest](docs/ActiveFormatDescriptorTest.md)
 - [ActiveFormatTest](docs/ActiveFormatTest.md)
 - [AdvancedGopLengthTest](docs/AdvancedGopLengthTest.md)
 - [Alert](docs/Alert.md)
 - [ArrayOfPictureEssenceCoding](docs/ArrayOfPictureEssenceCoding.md)
 - [ArrayOfSoundEssenceCoding](docs/ArrayOfSoundEssenceCoding.md)
 - [As11Rules](docs/As11Rules.md)
 - [As11UkDppMetadataTest](docs/As11UkDppMetadataTest.md)
 - [As11XprofileTest](docs/As11XprofileTest.md)
 - [AudioBitDepthTest](docs/AudioBitDepthTest.md)
 - [AudioBitrateTest](docs/AudioBitrateTest.md)
 - [AudioChannelPositionsTest](docs/AudioChannelPositionsTest.md)
 - [AudioChannelsTest](docs/AudioChannelsTest.md)
 - [AudioClippingTest](docs/AudioClippingTest.md)
 - [AudioCodecTest](docs/AudioCodecTest.md)
 - [AudioCodecType](docs/AudioCodecType.md)
 - [AudioConfig](docs/AudioConfig.md)
 - [AudioConfigs](docs/AudioConfigs.md)
 - [AudioDialnormTest](docs/AudioDialnormTest.md)
 - [AudioFrequencyTest](docs/AudioFrequencyTest.md)
 - [AudioLengthTest](docs/AudioLengthTest.md)
 - [AudioLoudnessItest](docs/AudioLoudnessItest.md)
 - [AudioLoudnessMtest](docs/AudioLoudnessMtest.md)
 - [AudioLoudnessRangeTest](docs/AudioLoudnessRangeTest.md)
 - [AudioLoudnessStest](docs/AudioLoudnessStest.md)
 - [AudioMinLevelDurationTest](docs/AudioMinLevelDurationTest.md)
 - [AudioPeakLevelTest](docs/AudioPeakLevelTest.md)
 - [AudioPhaseTest](docs/AudioPhaseTest.md)
 - [AudioPpmLevelTest](docs/AudioPpmLevelTest.md)
 - [AudioSampleRateTest](docs/AudioSampleRateTest.md)
 - [AudioStream](docs/AudioStream.md)
 - [AudioTest](docs/AudioTest.md)
 - [AudioTracksTest](docs/AudioTracksTest.md)
 - [AudioTransientTest](docs/AudioTransientTest.md)
 - [AvcCodedContentKindTest](docs/AvcCodedContentKindTest.md)
 - [AvcContentKind](docs/AvcContentKind.md)
 - [AvcSpsPpsTest](docs/AvcSpsPpsTest.md)
 - [BitRateMode](docs/BitRateMode.md)
 - [BitRateModeTest](docs/BitRateModeTest.md)
 - [BlackFrameTest](docs/BlackFrameTest.md)
 - [BlackLevelTest](docs/BlackLevelTest.md)
 - [BlankingTest](docs/BlankingTest.md)
 - [BlockinessTest](docs/BlockinessTest.md)
 - [BoolValueTest](docs/BoolValueTest.md)
 - [BufferSizeTest](docs/BufferSizeTest.md)
 - [CabacTest](docs/CabacTest.md)
 - [CadenceTest](docs/CadenceTest.md)
 - [CadenceType](docs/CadenceType.md)
 - [CaptionsTest](docs/CaptionsTest.md)
 - [ChanPos](docs/ChanPos.md)
 - [ChanPositions](docs/ChanPositions.md)
 - [Channels](docs/Channels.md)
 - [ChromaLevelTest](docs/ChromaLevelTest.md)
 - [ChromaSubsampling](docs/ChromaSubsampling.md)
 - [ChromaSubsamplingTest](docs/ChromaSubsamplingTest.md)
 - [CleanApertureTest](docs/CleanApertureTest.md)
 - [ColorBarStandardType](docs/ColorBarStandardType.md)
 - [ColorRangeTest](docs/ColorRangeTest.md)
 - [ColorSitingTest](docs/ColorSitingTest.md)
 - [ColorSpaceType](docs/ColorSpaceType.md)
 - [ColourBarsTest](docs/ColourBarsTest.md)
 - [ComponentDepthTest](docs/ComponentDepthTest.md)
 - [Container](docs/Container.md)
 - [ContainerEssenceConsistencyTest](docs/ContainerEssenceConsistencyTest.md)
 - [ContainerTest](docs/ContainerTest.md)
 - [ContainerType](docs/ContainerType.md)
 - [CorruptFrameTest](docs/CorruptFrameTest.md)
 - [DefaultOrCustomType](docs/DefaultOrCustomType.md)
 - [DigitalDropoutTest](docs/DigitalDropoutTest.md)
 - [DigitalSilenceAtStartEndTest](docs/DigitalSilenceAtStartEndTest.md)
 - [DigitalSilenceWholeTrackTest](docs/DigitalSilenceWholeTrackTest.md)
 - [DontCopyAvDelayTest](docs/DontCopyAvDelayTest.md)
 - [DropFrameTest](docs/DropFrameTest.md)
 - [DropFrameType](docs/DropFrameType.md)
 - [DropoutTest](docs/DropoutTest.md)
 - [EmbeddedXmlDocuments](docs/EmbeddedXmlDocuments.md)
 - [EnhancedSyntaxTest](docs/EnhancedSyntaxTest.md)
 - [ExtendedBool](docs/ExtendedBool.md)
 - [ExtendedBoolValueTest](docs/ExtendedBoolValueTest.md)
 - [ExtraAudioLayoutModes](docs/ExtraAudioLayoutModes.md)
 - [ExtraFile](docs/ExtraFile.md)
 - [FieldDominanceTest](docs/FieldDominanceTest.md)
 - [FieldOrderTest](docs/FieldOrderTest.md)
 - [FieldOrderType](docs/FieldOrderType.md)
 - [FileBitrateTest](docs/FileBitrateTest.md)
 - [FileConfig](docs/FileConfig.md)
 - [FileDurationTest](docs/FileDurationTest.md)
 - [FileFormatSpecificationIdentificationLabel](docs/FileFormatSpecificationIdentificationLabel.md)
 - [FlashTest](docs/FlashTest.md)
 - [ForceColorSpaceTest](docs/ForceColorSpaceTest.md)
 - [FrameAspectRatioTest](docs/FrameAspectRatioTest.md)
 - [FramerateTest](docs/FramerateTest.md)
 - [FramesizeTest](docs/FramesizeTest.md)
 - [FreezeFrameTest](docs/FreezeFrameTest.md)
 - [GopLengthTest](docs/GopLengthTest.md)
 - [GopOrder](docs/GopOrder.md)
 - [GopReport](docs/GopReport.md)
 - [HdrStandardType](docs/HdrStandardType.md)
 - [HdrTest](docs/HdrTest.md)
 - [HeaderByteCountTest](docs/HeaderByteCountTest.md)
 - [HeaderFillTest](docs/HeaderFillTest.md)
 - [ITunesCompatibilityTest](docs/ITunesCompatibilityTest.md)
 - [IgnoreVbiTest](docs/IgnoreVbiTest.md)
 - [ImfConformanceTest](docs/ImfConformanceTest.md)
 - [IndexTableTest](docs/IndexTableTest.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Job](docs/Job.md)
 - [JobData](docs/JobData.md)
 - [JobsCollection](docs/JobsCollection.md)
 - [KagSizeTest](docs/KagSizeTest.md)
 - [LayoutTest](docs/LayoutTest.md)
 - [LayoutType](docs/LayoutType.md)
 - [LetterboxingTest](docs/LetterboxingTest.md)
 - [LocationTest](docs/LocationTest.md)
 - [LongMinMaxTest](docs/LongMinMaxTest.md)
 - [LongValueTest](docs/LongValueTest.md)
 - [LossOfChromaTest](docs/LossOfChromaTest.md)
 - [LoudnessMode](docs/LoudnessMode.md)
 - [LowPassFilterType](docs/LowPassFilterType.md)
 - [MbaffTest](docs/MbaffTest.md)
 - [Media](docs/Media.md)
 - [MediaOfflineTest](docs/MediaOfflineTest.md)
 - [ModifyProjectBody](docs/ModifyProjectBody.md)
 - [ModifyVidChecker8Body](docs/ModifyVidChecker8Body.md)
 - [MustOrMustNot](docs/MustOrMustNot.md)
 - [MxfColorSiting](docs/MxfColorSiting.md)
 - [MxfFieldDominance](docs/MxfFieldDominance.md)
 - [MxfKeyTest](docs/MxfKeyTest.md)
 - [MxfOpTest](docs/MxfOpTest.md)
 - [MxfTest](docs/MxfTest.md)
 - [MxfVersion](docs/MxfVersion.md)
 - [NetflixPhotonTest](docs/NetflixPhotonTest.md)
 - [NielsenWatermarkDetectionTest](docs/NielsenWatermarkDetectionTest.md)
 - [OpenOrClosed](docs/OpenOrClosed.md)
 - [OperationalPattern](docs/OperationalPattern.md)
 - [OperationalPatternTest](docs/OperationalPatternTest.md)
 - [OptionalFlag](docs/OptionalFlag.md)
 - [PSeParameterType](docs/PSeParameterType.md)
 - [PaddingBitsTest](docs/PaddingBitsTest.md)
 - [PartitionStatusTest](docs/PartitionStatusTest.md)
 - [PicFrameSizeTest](docs/PicFrameSizeTest.md)
 - [PictureEssenceCoding](docs/PictureEssenceCoding.md)
 - [PictureEssenceCodingTest](docs/PictureEssenceCodingTest.md)
 - [PictureEssenceConstraints](docs/PictureEssenceConstraints.md)
 - [PictureOffsetsTest](docs/PictureOffsetsTest.md)
 - [PixelAspectRatioTest](docs/PixelAspectRatioTest.md)
 - [PpmMode](docs/PpmMode.md)
 - [Project](docs/Project.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [Proxy](docs/Proxy.md)
 - [RatioOrLinesType](docs/RatioOrLinesType.md)
 - [RatioTest](docs/RatioTest.md)
 - [ReferenceLevelsTest](docs/ReferenceLevelsTest.md)
 - [RequireOrDisallow](docs/RequireOrDisallow.md)
 - [RgbGamutTest](docs/RgbGamutTest.md)
 - [RipPresentTest](docs/RipPresentTest.md)
 - [RunInTest](docs/RunInTest.md)
 - [SdtiTimecodeContinuityTest](docs/SdtiTimecodeContinuityTest.md)
 - [SecsOrFramesType](docs/SecsOrFramesType.md)
 - [SensitivityType](docs/SensitivityType.md)
 - [SidsAnyOrSpecific](docs/SidsAnyOrSpecific.md)
 - [SignalStandardTest](docs/SignalStandardTest.md)
 - [SingleColorTest](docs/SingleColorTest.md)
 - [SingleSampleDescriptionTest](docs/SingleSampleDescriptionTest.md)
 - [SoundEssenceCoding](docs/SoundEssenceCoding.md)
 - [SoundEssenceCodingTest](docs/SoundEssenceCodingTest.md)
 - [SpsPpsTest](docs/SpsPpsTest.md)
 - [StartTcRangeMethod](docs/StartTcRangeMethod.md)
 - [StartTimecodeTest](docs/StartTimecodeTest.md)
 - [StripeTest](docs/StripeTest.md)
 - [SubsamplingTest](docs/SubsamplingTest.md)
 - [Summary](docs/Summary.md)
 - [SynchronizationEvent](docs/SynchronizationEvent.md)
 - [TeletextType](docs/TeletextType.md)
 - [Template](docs/Template.md)
 - [TextStreamTest](docs/TextStreamTest.md)
 - [TimeCodeSource](docs/TimeCodeSource.md)
 - [TimecodeContinuityTest](docs/TimecodeContinuityTest.md)
 - [TimecodeTrackTest](docs/TimecodeTrackTest.md)
 - [ToneType](docs/ToneType.md)
 - [TrackIdTest](docs/TrackIdTest.md)
 - [TrackSelectTest](docs/TrackSelectTest.md)
 - [TrackSelectorType](docs/TrackSelectorType.md)
 - [UkDppShim](docs/UkDppShim.md)
 - [UploadSession](docs/UploadSession.md)
 - [UseStartTimecodeTest](docs/UseStartTimecodeTest.md)
 - [VersionTest](docs/VersionTest.md)
 - [VidChecker8Body](docs/VidChecker8Body.md)
 - [VidChecker8JobData](docs/VidChecker8JobData.md)
 - [Vidchecker8](docs/Vidchecker8.md)
 - [VideoBitDepthTest](docs/VideoBitDepthTest.md)
 - [VideoBitrateTest](docs/VideoBitrateTest.md)
 - [VideoCodecTest](docs/VideoCodecTest.md)
 - [VideoCodecType](docs/VideoCodecType.md)
 - [VideoConfig](docs/VideoConfig.md)
 - [VideoConfigs](docs/VideoConfigs.md)
 - [VideoDescriptorTest](docs/VideoDescriptorTest.md)
 - [VideoDescriptorType](docs/VideoDescriptorType.md)
 - [VideoLevelType](docs/VideoLevelType.md)
 - [VideoLineMapTest](docs/VideoLineMapTest.md)
 - [VideoProfileType](docs/VideoProfileType.md)
 - [VideoSegmentDetectionTest](docs/VideoSegmentDetectionTest.md)
 - [VideoStream](docs/VideoStream.md)
 - [VideoSubDescriptorTest](docs/VideoSubDescriptorTest.md)
 - [VideoSubDescriptorType](docs/VideoSubDescriptorType.md)
 - [VideoTest](docs/VideoTest.md)
 - [VideoUploadBody](docs/VideoUploadBody.md)
 - [WrappingType](docs/WrappingType.md)
 - [WrappingTypeTest](docs/WrappingTypeTest.md)


## Documentation For Authorization



## apiKey

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## Author

cloudsupport@telestream.net

