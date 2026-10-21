# ServerSubscribedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewServerSubscribedMessage

`func NewServerSubscribedMessage(topic string, type_ string, v int32, ) *ServerSubscribedMessage`

NewServerSubscribedMessage instantiates a new ServerSubscribedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSubscribedMessageWithDefaults

`func NewServerSubscribedMessageWithDefaults() *ServerSubscribedMessage`

NewServerSubscribedMessageWithDefaults instantiates a new ServerSubscribedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *ServerSubscribedMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerSubscribedMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerSubscribedMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ServerSubscribedMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerSubscribedMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerSubscribedMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerSubscribedMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerSubscribedMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerSubscribedMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


