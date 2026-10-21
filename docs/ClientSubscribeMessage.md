# ClientSubscribeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topics** | **[]string** | A message must fit within 128 KiB as sent, so topics whose JSON encoding is long go in several messages. | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewClientSubscribeMessage

`func NewClientSubscribeMessage(topics []string, type_ string, v int32, ) *ClientSubscribeMessage`

NewClientSubscribeMessage instantiates a new ClientSubscribeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSubscribeMessageWithDefaults

`func NewClientSubscribeMessageWithDefaults() *ClientSubscribeMessage`

NewClientSubscribeMessageWithDefaults instantiates a new ClientSubscribeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopics

`func (o *ClientSubscribeMessage) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ClientSubscribeMessage) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ClientSubscribeMessage) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetType

`func (o *ClientSubscribeMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientSubscribeMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientSubscribeMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ClientSubscribeMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientSubscribeMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientSubscribeMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


