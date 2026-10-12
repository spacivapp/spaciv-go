# ScriptRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int64** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**ExitCode** | **int32** |  | 
**Id** | **string** |  | 
**Kind** | **string** | \&quot;build\&quot; for a dependency build, \&quot;run\&quot; for an execution. | 
**Message** | Pointer to **string** |  | [optional] 
**QueuedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**Trigger** | **string** |  | 

## Methods

### NewScriptRun

`func NewScriptRun(exitCode int32, id string, kind string, status string, trigger string, ) *ScriptRun`

NewScriptRun instantiates a new ScriptRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptRunWithDefaults

`func NewScriptRunWithDefaults() *ScriptRun`

NewScriptRunWithDefaults instantiates a new ScriptRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *ScriptRun) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ScriptRun) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ScriptRun) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ScriptRun) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndedAt

`func (o *ScriptRun) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ScriptRun) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ScriptRun) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ScriptRun) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExitCode

`func (o *ScriptRun) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ScriptRun) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ScriptRun) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetId

`func (o *ScriptRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptRun) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ScriptRun) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScriptRun) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScriptRun) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMessage

`func (o *ScriptRun) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScriptRun) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScriptRun) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScriptRun) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQueuedAt

`func (o *ScriptRun) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *ScriptRun) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *ScriptRun) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *ScriptRun) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ScriptRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ScriptRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ScriptRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ScriptRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ScriptRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScriptRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScriptRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTrigger

`func (o *ScriptRun) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ScriptRun) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ScriptRun) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


