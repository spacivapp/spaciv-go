# ScratchpadRunPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | **map[string]string** | Source files keyed by path. Python needs script.py defining handle(event, ctx) and JavaScript needs script.js exporting handle(event, ctx), as an integration script does, so the code can be saved as one unchanged. | 
**Input** | Pointer to **interface{}** |  | [optional] 
**Language** | **string** | python or javascript. Go code has to be saved as a script, which builds it. | 
**TimeoutSeconds** | Pointer to **int32** | Longest the code may run. Defaults to 30. | [optional] 
**WaitSeconds** | Pointer to **int32** | Longest to wait for the run to finish before responding. Defaults to 20. | [optional] 

## Methods

### NewScratchpadRunPostReqBody

`func NewScratchpadRunPostReqBody(files map[string]string, language string, ) *ScratchpadRunPostReqBody`

NewScratchpadRunPostReqBody instantiates a new ScratchpadRunPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadRunPostReqBodyWithDefaults

`func NewScratchpadRunPostReqBodyWithDefaults() *ScratchpadRunPostReqBody`

NewScratchpadRunPostReqBodyWithDefaults instantiates a new ScratchpadRunPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


