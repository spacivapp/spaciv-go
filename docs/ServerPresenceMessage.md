# ServerPresenceMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Member** | [**PresenceMember**](PresenceMember.md) |  | 
**Topic** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewServerPresenceMessage

`func NewServerPresenceMessage(event string, member PresenceMember, topic string, type_ string, v int32, ) *ServerPresenceMessage`

NewServerPresenceMessage instantiates a new ServerPresenceMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPresenceMessageWithDefaults

`func NewServerPresenceMessageWithDefaults() *ServerPresenceMessage`

NewServerPresenceMessageWithDefaults instantiates a new ServerPresenceMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ServerPresenceMessage) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ServerPresenceMessage) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ServerPresenceMessage) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMember

`func (o *ServerPresenceMessage) GetMember() PresenceMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServerPresenceMessage) GetMemberOk() (*PresenceMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServerPresenceMessage) SetMember(v PresenceMember)`

SetMember sets Member field to given value.


### GetTopic

`func (o *ServerPresenceMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerPresenceMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerPresenceMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ServerPresenceMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerPresenceMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerPresenceMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerPresenceMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerPresenceMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerPresenceMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


