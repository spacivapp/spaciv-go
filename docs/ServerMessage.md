# ServerMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EntityChangedEvent**](EntityChangedEvent.md) |  | 
**Topic** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 
**Code** | **string** |  | 
**Message** | **string** |  | 
**Event** | **string** |  | 
**Member** | [**PresenceMember**](PresenceMember.md) |  | 
**Members** | [**[]PresenceMember**](PresenceMember.md) | Everyone else present on the topic. It replaces what the connection knew, and presence events follow it. | 

## Methods

### NewServerMessage

`func NewServerMessage(events []EntityChangedEvent, topic string, type_ string, v int32, code string, message string, event string, member PresenceMember, members []PresenceMember, ) *ServerMessage`

NewServerMessage instantiates a new ServerMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMessageWithDefaults

`func NewServerMessageWithDefaults() *ServerMessage`

NewServerMessageWithDefaults instantiates a new ServerMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ServerMessage) GetEvents() []EntityChangedEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ServerMessage) GetEventsOk() (*[]EntityChangedEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ServerMessage) SetEvents(v []EntityChangedEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *ServerMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ServerMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerMessage) SetV(v int32)`

SetV sets V field to given value.


### GetCode

`func (o *ServerMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServerMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServerMessage) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ServerMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEvent

`func (o *ServerMessage) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ServerMessage) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ServerMessage) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMember

`func (o *ServerMessage) GetMember() PresenceMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServerMessage) GetMemberOk() (*PresenceMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServerMessage) SetMember(v PresenceMember)`

SetMember sets Member field to given value.


### GetMembers

`func (o *ServerMessage) GetMembers() []PresenceMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ServerMessage) GetMembersOk() (*[]PresenceMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ServerMessage) SetMembers(v []PresenceMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


