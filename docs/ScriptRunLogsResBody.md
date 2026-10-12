# ScriptRunLogsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | [**[]ScriptLogLine**](ScriptLogLine.md) |  | 
**Truncated** | **bool** | Set when the run wrote more output than is kept, so some of it is missing. | 

## Methods

### NewScriptRunLogsResBody

`func NewScriptRunLogsResBody(lines []ScriptLogLine, truncated bool, ) *ScriptRunLogsResBody`

NewScriptRunLogsResBody instantiates a new ScriptRunLogsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptRunLogsResBodyWithDefaults

`func NewScriptRunLogsResBodyWithDefaults() *ScriptRunLogsResBody`

NewScriptRunLogsResBodyWithDefaults instantiates a new ScriptRunLogsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *ScriptRunLogsResBody) GetLines() []ScriptLogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ScriptRunLogsResBody) GetLinesOk() (*[]ScriptLogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ScriptRunLogsResBody) SetLines(v []ScriptLogLine)`

SetLines sets Lines field to given value.


### GetTruncated

`func (o *ScriptRunLogsResBody) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ScriptRunLogsResBody) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ScriptRunLogsResBody) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


