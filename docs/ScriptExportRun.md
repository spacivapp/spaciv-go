# ScriptExportRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int64** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**ExitCode** | **int32** |  | 
**Id** | **string** |  | 
**Kind** | **string** | \&quot;build\&quot; for a dependency build, \&quot;run\&quot; for an execution. | 
**LogsPath** | **string** | API path that returns the run&#39;s logs while the run is kept. | 
**Message** | Pointer to **string** |  | [optional] 
**QueuedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**Trigger** | **string** |  | 

## Methods

### NewScriptExportRun

`func NewScriptExportRun(exitCode int32, id string, kind string, logsPath string, status string, trigger string, ) *ScriptExportRun`

NewScriptExportRun instantiates a new ScriptExportRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExportRunWithDefaults

`func NewScriptExportRunWithDefaults() *ScriptExportRun`

NewScriptExportRunWithDefaults instantiates a new ScriptExportRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *ScriptExportRun) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ScriptExportRun) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ScriptExportRun) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ScriptExportRun) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndedAt

`func (o *ScriptExportRun) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ScriptExportRun) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ScriptExportRun) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ScriptExportRun) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExitCode

`func (o *ScriptExportRun) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ScriptExportRun) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ScriptExportRun) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetId

`func (o *ScriptExportRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptExportRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptExportRun) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ScriptExportRun) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScriptExportRun) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScriptExportRun) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLogsPath

`func (o *ScriptExportRun) GetLogsPath() string`

GetLogsPath returns the LogsPath field if non-nil, zero value otherwise.

### GetLogsPathOk

`func (o *ScriptExportRun) GetLogsPathOk() (*string, bool)`

GetLogsPathOk returns a tuple with the LogsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsPath

`func (o *ScriptExportRun) SetLogsPath(v string)`

SetLogsPath sets LogsPath field to given value.


### GetMessage

`func (o *ScriptExportRun) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScriptExportRun) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScriptExportRun) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScriptExportRun) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQueuedAt

`func (o *ScriptExportRun) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *ScriptExportRun) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *ScriptExportRun) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *ScriptExportRun) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ScriptExportRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ScriptExportRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ScriptExportRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ScriptExportRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ScriptExportRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScriptExportRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScriptExportRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTrigger

`func (o *ScriptExportRun) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ScriptExportRun) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ScriptExportRun) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


