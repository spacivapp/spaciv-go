# ServerPresenceSnapshotMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]PresenceMember**](PresenceMember.md) | Everyone else present on the topic. It replaces what the connection knew, and presence events follow it. | 
**Topic** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewServerPresenceSnapshotMessage

`func NewServerPresenceSnapshotMessage(members []PresenceMember, topic string, type_ string, v int32, ) *ServerPresenceSnapshotMessage`

NewServerPresenceSnapshotMessage instantiates a new ServerPresenceSnapshotMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPresenceSnapshotMessageWithDefaults

`func NewServerPresenceSnapshotMessageWithDefaults() *ServerPresenceSnapshotMessage`

NewServerPresenceSnapshotMessageWithDefaults instantiates a new ServerPresenceSnapshotMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ServerPresenceSnapshotMessage) GetMembers() []PresenceMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ServerPresenceSnapshotMessage) GetMembersOk() (*[]PresenceMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ServerPresenceSnapshotMessage) SetMembers(v []PresenceMember)`

SetMembers sets Members field to given value.


### GetTopic

`func (o *ServerPresenceSnapshotMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerPresenceSnapshotMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerPresenceSnapshotMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ServerPresenceSnapshotMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerPresenceSnapshotMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerPresenceSnapshotMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerPresenceSnapshotMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerPresenceSnapshotMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerPresenceSnapshotMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


