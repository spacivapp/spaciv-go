# ClientMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 
**Topics** | **[]string** |  | 
**State** | **string** | none leaves the topic&#39;s presence, as unsubscribing does. | 
**Target** | Pointer to **string** | What on the topic&#39;s resource is being viewed or edited, such as a card. | [optional] 
**Topic** | **string** | A topic this connection has subscribed to; presence on any other topic is ignored. | 

## Methods

### NewClientMessage

`func NewClientMessage(token string, type_ string, v int32, topics []string, state string, topic string, ) *ClientMessage`

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

`func (o *ClientMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientMessage) SetV(v int32)`

SetV sets V field to given value.


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


### GetState

`func (o *ClientMessage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientMessage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClientMessage) SetState(v string)`

SetState sets State field to given value.


### GetTarget

`func (o *ClientMessage) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ClientMessage) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ClientMessage) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ClientMessage) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTopic

`func (o *ClientMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ClientMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ClientMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


