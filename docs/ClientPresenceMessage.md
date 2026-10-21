# ClientPresenceMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | none leaves the topic&#39;s presence, as unsubscribing does. | 
**Target** | Pointer to **string** | What on the topic&#39;s resource is being viewed or edited, such as a card. | [optional] 
**Topic** | **string** | A topic this connection has subscribed to; presence on any other topic is ignored. | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewClientPresenceMessage

`func NewClientPresenceMessage(state string, topic string, type_ string, v int32, ) *ClientPresenceMessage`

NewClientPresenceMessage instantiates a new ClientPresenceMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientPresenceMessageWithDefaults

`func NewClientPresenceMessageWithDefaults() *ClientPresenceMessage`

NewClientPresenceMessageWithDefaults instantiates a new ClientPresenceMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ClientPresenceMessage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientPresenceMessage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClientPresenceMessage) SetState(v string)`

SetState sets State field to given value.


### GetTarget

`func (o *ClientPresenceMessage) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ClientPresenceMessage) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ClientPresenceMessage) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ClientPresenceMessage) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTopic

`func (o *ClientPresenceMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ClientPresenceMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ClientPresenceMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ClientPresenceMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientPresenceMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientPresenceMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ClientPresenceMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientPresenceMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientPresenceMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


