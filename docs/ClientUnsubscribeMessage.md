# ClientUnsubscribeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topics** | **[]string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewClientUnsubscribeMessage

`func NewClientUnsubscribeMessage(topics []string, type_ string, v int32, ) *ClientUnsubscribeMessage`

NewClientUnsubscribeMessage instantiates a new ClientUnsubscribeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUnsubscribeMessageWithDefaults

`func NewClientUnsubscribeMessageWithDefaults() *ClientUnsubscribeMessage`

NewClientUnsubscribeMessageWithDefaults instantiates a new ClientUnsubscribeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopics

`func (o *ClientUnsubscribeMessage) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ClientUnsubscribeMessage) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ClientUnsubscribeMessage) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetType

`func (o *ClientUnsubscribeMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientUnsubscribeMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientUnsubscribeMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ClientUnsubscribeMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientUnsubscribeMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientUnsubscribeMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


