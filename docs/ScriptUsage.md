# ScriptUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedRuns** | **int64** | Runs that ended in error, timeout or killed. | 
**QuotaExceededRuns** | **int64** | Runs that did not start because a limit was reached. | 
**Runs** | **int64** | Runs that were not refused by a limit, whatever their outcome, including runs that failed before their sandbox started. | 
**ScriptId** | **string** |  | 

## Methods

### NewScriptUsage

`func NewScriptUsage(failedRuns int64, quotaExceededRuns int64, runs int64, scriptId string, ) *ScriptUsage`

NewScriptUsage instantiates a new ScriptUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptUsageWithDefaults

`func NewScriptUsageWithDefaults() *ScriptUsage`

NewScriptUsageWithDefaults instantiates a new ScriptUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedRuns

`func (o *ScriptUsage) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *ScriptUsage) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *ScriptUsage) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetQuotaExceededRuns

`func (o *ScriptUsage) GetQuotaExceededRuns() int64`

GetQuotaExceededRuns returns the QuotaExceededRuns field if non-nil, zero value otherwise.

### GetQuotaExceededRunsOk

`func (o *ScriptUsage) GetQuotaExceededRunsOk() (*int64, bool)`

GetQuotaExceededRunsOk returns a tuple with the QuotaExceededRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExceededRuns

`func (o *ScriptUsage) SetQuotaExceededRuns(v int64)`

SetQuotaExceededRuns sets QuotaExceededRuns field to given value.


### GetRuns

`func (o *ScriptUsage) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ScriptUsage) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ScriptUsage) SetRuns(v int64)`

SetRuns sets Runs field to given value.


### GetScriptId

`func (o *ScriptUsage) GetScriptId() string`

GetScriptId returns the ScriptId field if non-nil, zero value otherwise.

### GetScriptIdOk

`func (o *ScriptUsage) GetScriptIdOk() (*string, bool)`

GetScriptIdOk returns a tuple with the ScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptId

`func (o *ScriptUsage) SetScriptId(v string)`

SetScriptId sets ScriptId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


