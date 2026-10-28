# ScratchpadBuildPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | **map[string]string** | Source files keyed by path, as for a run, with the language&#39;s dependency manifest (requirements.txt, go.mod or package.json) among them. | 
**Language** | **string** | python, go or javascript. | 
**WaitSeconds** | Pointer to **int32** | Longest to wait for the build to finish before responding. Defaults to 20. | [optional] 

## Methods

### NewScratchpadBuildPostReqBody

`func NewScratchpadBuildPostReqBody(files map[string]string, language string, ) *ScratchpadBuildPostReqBody`

NewScratchpadBuildPostReqBody instantiates a new ScratchpadBuildPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadBuildPostReqBodyWithDefaults

`func NewScratchpadBuildPostReqBodyWithDefaults() *ScratchpadBuildPostReqBody`

NewScratchpadBuildPostReqBodyWithDefaults instantiates a new ScratchpadBuildPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ScratchpadBuildPostReqBody) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScratchpadBuildPostReqBody) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScratchpadBuildPostReqBody) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.


### GetLanguage

`func (o *ScratchpadBuildPostReqBody) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScratchpadBuildPostReqBody) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScratchpadBuildPostReqBody) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetWaitSeconds

`func (o *ScratchpadBuildPostReqBody) GetWaitSeconds() int32`

GetWaitSeconds returns the WaitSeconds field if non-nil, zero value otherwise.

### GetWaitSecondsOk

`func (o *ScratchpadBuildPostReqBody) GetWaitSecondsOk() (*int32, bool)`

GetWaitSecondsOk returns a tuple with the WaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitSeconds

`func (o *ScratchpadBuildPostReqBody) SetWaitSeconds(v int32)`

SetWaitSeconds sets WaitSeconds field to given value.

### HasWaitSeconds

`func (o *ScratchpadBuildPostReqBody) HasWaitSeconds() bool`

HasWaitSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


