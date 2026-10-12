# ScriptAuditResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]ScriptAuditEntry**](ScriptAuditEntry.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewScriptAuditResBody

`func NewScriptAuditResBody(entries []ScriptAuditEntry, ) *ScriptAuditResBody`

NewScriptAuditResBody instantiates a new ScriptAuditResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptAuditResBodyWithDefaults

`func NewScriptAuditResBodyWithDefaults() *ScriptAuditResBody`

NewScriptAuditResBodyWithDefaults instantiates a new ScriptAuditResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ScriptAuditResBody) GetEntries() []ScriptAuditEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ScriptAuditResBody) GetEntriesOk() (*[]ScriptAuditEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ScriptAuditResBody) SetEntries(v []ScriptAuditEntry)`

SetEntries sets Entries field to given value.


### GetNextCursor

`func (o *ScriptAuditResBody) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ScriptAuditResBody) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ScriptAuditResBody) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ScriptAuditResBody) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


