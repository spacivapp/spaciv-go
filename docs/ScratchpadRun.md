# ScratchpadRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int64** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**ExitCode** | **int32** |  | 
**ExpiresAt** | Pointer to **time.Time** | When the run and its output are deleted. | [optional] 
**Files** | **map[string]string** | The code that was run, keyed by path. | 
**Id** | **string** |  | 
**Language** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Output** | Pointer to [**ScratchpadOutput**](ScratchpadOutput.md) | What the run produced. Absent while it is in progress. | [optional] 
**QueuedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**StartupMs** | Pointer to **int64** | How long the sandbox took to start. | [optional] 
**Status** | **string** | queued or running while the run is in progress, then success, error, timeout, killed or quota_exceeded. | 
**TimeoutSeconds** | **int32** |  | 
**Warm** | **bool** | Whether the run used a sandbox that had already started. | 

## Methods

### NewScratchpadRun

`func NewScratchpadRun(exitCode int32, files map[string]string, id string, language string, status string, timeoutSeconds int32, warm bool, ) *ScratchpadRun`

NewScratchpadRun instantiates a new ScratchpadRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadRunWithDefaults

`func NewScratchpadRunWithDefaults() *ScratchpadRun`

NewScratchpadRunWithDefaults instantiates a new ScratchpadRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *ScratchpadRun) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ScratchpadRun) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ScratchpadRun) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ScratchpadRun) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndedAt

`func (o *ScratchpadRun) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ScratchpadRun) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ScratchpadRun) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ScratchpadRun) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExitCode

`func (o *ScratchpadRun) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ScratchpadRun) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ScratchpadRun) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetExpiresAt

`func (o *ScratchpadRun) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ScratchpadRun) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ScratchpadRun) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ScratchpadRun) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFiles

`func (o *ScratchpadRun) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScratchpadRun) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScratchpadRun) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.


### GetId

`func (o *ScratchpadRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScratchpadRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScratchpadRun) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *ScratchpadRun) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScratchpadRun) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScratchpadRun) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetMessage

`func (o *ScratchpadRun) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScratchpadRun) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScratchpadRun) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScratchpadRun) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *ScratchpadRun) GetOutput() ScratchpadOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ScratchpadRun) GetOutputOk() (*ScratchpadOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ScratchpadRun) SetOutput(v ScratchpadOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ScratchpadRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetQueuedAt

`func (o *ScratchpadRun) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *ScratchpadRun) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *ScratchpadRun) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *ScratchpadRun) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ScratchpadRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ScratchpadRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ScratchpadRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ScratchpadRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStartupMs

`func (o *ScratchpadRun) GetStartupMs() int64`

GetStartupMs returns the StartupMs field if non-nil, zero value otherwise.

### GetStartupMsOk

`func (o *ScratchpadRun) GetStartupMsOk() (*int64, bool)`

GetStartupMsOk returns a tuple with the StartupMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupMs

`func (o *ScratchpadRun) SetStartupMs(v int64)`

SetStartupMs sets StartupMs field to given value.

### HasStartupMs

`func (o *ScratchpadRun) HasStartupMs() bool`

HasStartupMs returns a boolean if a field has been set.

### GetStatus

`func (o *ScratchpadRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScratchpadRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScratchpadRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeoutSeconds

`func (o *ScratchpadRun) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ScratchpadRun) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ScratchpadRun) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetWarm

`func (o *ScratchpadRun) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *ScratchpadRun) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *ScratchpadRun) SetWarm(v bool)`

SetWarm sets Warm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


