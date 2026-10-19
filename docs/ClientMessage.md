# ClientMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **[]string** | Between 1 and 100 topics for subscribe and unsubscribe. A message must fit within 128 KiB as sent, so topics whose JSON encoding is long go in several messages. | [optional] 
**Type** | **string** |  | 
**V** | **int64** |  | 

## Methods

### NewClientMessage

`func NewClientMessage(type_ string, v int64, ) *ClientMessage`

NewClientMessage instantiates a new ClientMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMessageWithDefaults

`func NewClientMessageWithDefaults() *ClientMessage`

NewClientMessageWithDefaults instantiates a new ClientMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ClientMessage) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientMessage) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientMessage) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientMessage) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTopics

`func (o *ClientMessage) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ClientMessage) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ClientMessage) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ClientMessage) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetType

`func (o *ClientMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ClientMessage) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientMessage) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientMessage) SetV(v int64)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


