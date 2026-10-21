# ServerChangedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EntityChangedEvent**](EntityChangedEvent.md) |  | 
**Topic** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewServerChangedMessage

`func NewServerChangedMessage(events []EntityChangedEvent, topic string, type_ string, v int32, ) *ServerChangedMessage`

NewServerChangedMessage instantiates a new ServerChangedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerChangedMessageWithDefaults

`func NewServerChangedMessageWithDefaults() *ServerChangedMessage`

NewServerChangedMessageWithDefaults instantiates a new ServerChangedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ServerChangedMessage) GetEvents() []EntityChangedEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ServerChangedMessage) GetEventsOk() (*[]EntityChangedEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ServerChangedMessage) SetEvents(v []EntityChangedEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *ServerChangedMessage) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ServerChangedMessage) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ServerChangedMessage) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetType

`func (o *ServerChangedMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerChangedMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerChangedMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ServerChangedMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ServerChangedMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ServerChangedMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


