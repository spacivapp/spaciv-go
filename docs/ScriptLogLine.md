# ScriptLogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewScriptLogLine

`func NewScriptLogLine(message string, timestamp time.Time, ) *ScriptLogLine`

NewScriptLogLine instantiates a new ScriptLogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptLogLineWithDefaults

`func NewScriptLogLineWithDefaults() *ScriptLogLine`

NewScriptLogLineWithDefaults instantiates a new ScriptLogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ScriptLogLine) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScriptLogLine) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScriptLogLine) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestamp

`func (o *ScriptLogLine) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ScriptLogLine) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ScriptLogLine) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


