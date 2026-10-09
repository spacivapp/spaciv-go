# RetryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Event** | [**SerializedEvent**](SerializedEvent.md) |  | 
**RetryTimestamp** | **int64** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewRetryBody

`func NewRetryBody(error_ string, event SerializedEvent, retryTimestamp int64, timestamp int64, ) *RetryBody`

NewRetryBody instantiates a new RetryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryBodyWithDefaults

`func NewRetryBodyWithDefaults() *RetryBody`

NewRetryBodyWithDefaults instantiates a new RetryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RetryBody) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RetryBody) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RetryBody) SetError(v string)`

SetError sets Error field to given value.


### GetEvent

`func (o *RetryBody) GetEvent() SerializedEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RetryBody) GetEventOk() (*SerializedEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RetryBody) SetEvent(v SerializedEvent)`

SetEvent sets Event field to given value.


### GetRetryTimestamp

`func (o *RetryBody) GetRetryTimestamp() int64`

GetRetryTimestamp returns the RetryTimestamp field if non-nil, zero value otherwise.

### GetRetryTimestampOk

`func (o *RetryBody) GetRetryTimestampOk() (*int64, bool)`

GetRetryTimestampOk returns a tuple with the RetryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimestamp

`func (o *RetryBody) SetRetryTimestamp(v int64)`

SetRetryTimestamp sets RetryTimestamp field to given value.


### GetTimestamp

`func (o *RetryBody) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RetryBody) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RetryBody) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


