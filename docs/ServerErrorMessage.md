# ServerErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**Topic** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewServerErrorMessage

`func NewServerErrorMessage(code string, message string, type_ string, v int32, ) *ServerErrorMessage`

NewServerErrorMessage instantiates a new ServerErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerErrorMessageWithDefaults

`func NewServerErrorMessageWithDefaults() *ServerErrorMessage`

NewServerErrorMessageWithDefaults instantiates a new ServerErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ServerErrorMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServerErrorMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServerErrorMessage) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ServerErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTopic

`func (o *ServerErrorMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerErrorMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerErrorMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ServerErrorMessage) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *ServerErrorMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerErrorMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerErrorMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerErrorMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerErrorMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerErrorMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


