# ScriptRunsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** |  | [optional] 
**Runs** | [**[]ScriptRun**](ScriptRun.md) |  | 

## Methods

### NewScriptRunsResBody

`func NewScriptRunsResBody(runs []ScriptRun, ) *ScriptRunsResBody`

NewScriptRunsResBody instantiates a new ScriptRunsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptRunsResBodyWithDefaults

`func NewScriptRunsResBodyWithDefaults() *ScriptRunsResBody`

NewScriptRunsResBodyWithDefaults instantiates a new ScriptRunsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *ScriptRunsResBody) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ScriptRunsResBody) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ScriptRunsResBody) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ScriptRunsResBody) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetRuns

`func (o *ScriptRunsResBody) GetRuns() []ScriptRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ScriptRunsResBody) GetRunsOk() (*[]ScriptRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ScriptRunsResBody) SetRuns(v []ScriptRun)`

SetRuns sets Runs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


