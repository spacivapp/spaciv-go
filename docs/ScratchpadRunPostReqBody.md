# ScratchpadRunPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildId** | Pointer to **string** | The id of a build from POST /v1/scripts/scratchpad/builds, to run the files and dependencies that build holds instead of files. | [optional] 
**Files** | Pointer to **map[string]string** | Source files keyed by path. Python needs script.py defining handle(event, ctx), JavaScript needs script.js exporting handle(event, ctx), and Go needs script.go in a non-main package exporting Handle(ctx context.Context, event handler.Event, hc handler.Context) (any, error) from spaciv/handler, as an integration script does, so the code can be saved as one unchanged. Required unless buildId is given, and not allowed with it. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Language** | **string** | python, go or javascript. | 
**TimeoutSeconds** | Pointer to **int32** | Longest the code may run. Defaults to 30. | [optional] 
**WaitSeconds** | Pointer to **int32** | Longest to wait for the run to finish before responding. Defaults to 20. | [optional] 
**WriteFiles** | Pointer to **bool** | Gives the run&#39;s access token the *Edit file storage* permission as well, so the code can upload files to the account&#39;s file storage. Requires the caller to have it. | [optional] 

## Methods

### NewScratchpadRunPostReqBody

`func NewScratchpadRunPostReqBody(language string, ) *ScratchpadRunPostReqBody`

NewScratchpadRunPostReqBody instantiates a new ScratchpadRunPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadRunPostReqBodyWithDefaults

`func NewScratchpadRunPostReqBodyWithDefaults() *ScratchpadRunPostReqBody`

NewScratchpadRunPostReqBodyWithDefaults instantiates a new ScratchpadRunPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildId

`func (o *ScratchpadRunPostReqBody) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *ScratchpadRunPostReqBody) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *ScratchpadRunPostReqBody) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *ScratchpadRunPostReqBody) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetFiles

`func (o *ScratchpadRunPostReqBody) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScratchpadRunPostReqBody) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScratchpadRunPostReqBody) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ScratchpadRunPostReqBody) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetInput

`func (o *ScratchpadRunPostReqBody) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScratchpadRunPostReqBody) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScratchpadRunPostReqBody) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ScratchpadRunPostReqBody) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ScratchpadRunPostReqBody) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ScratchpadRunPostReqBody) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLanguage

`func (o *ScratchpadRunPostReqBody) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScratchpadRunPostReqBody) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScratchpadRunPostReqBody) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTimeoutSeconds

`func (o *ScratchpadRunPostReqBody) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ScratchpadRunPostReqBody) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ScratchpadRunPostReqBody) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ScratchpadRunPostReqBody) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetWaitSeconds

`func (o *ScratchpadRunPostReqBody) GetWaitSeconds() int32`

GetWaitSeconds returns the WaitSeconds field if non-nil, zero value otherwise.

### GetWaitSecondsOk

`func (o *ScratchpadRunPostReqBody) GetWaitSecondsOk() (*int32, bool)`

GetWaitSecondsOk returns a tuple with the WaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitSeconds

`func (o *ScratchpadRunPostReqBody) SetWaitSeconds(v int32)`

SetWaitSeconds sets WaitSeconds field to given value.

### HasWaitSeconds

`func (o *ScratchpadRunPostReqBody) HasWaitSeconds() bool`

HasWaitSeconds returns a boolean if a field has been set.

### GetWriteFiles

`func (o *ScratchpadRunPostReqBody) GetWriteFiles() bool`

GetWriteFiles returns the WriteFiles field if non-nil, zero value otherwise.

### GetWriteFilesOk

`func (o *ScratchpadRunPostReqBody) GetWriteFilesOk() (*bool, bool)`

GetWriteFilesOk returns a tuple with the WriteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteFiles

`func (o *ScratchpadRunPostReqBody) SetWriteFiles(v bool)`

SetWriteFiles sets WriteFiles field to given value.

### HasWriteFiles

`func (o *ScratchpadRunPostReqBody) HasWriteFiles() bool`

HasWriteFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


