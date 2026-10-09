# AssistantChatPollBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]AssistantChatPollAction**](AssistantChatPollAction.md) |  | 
**Failure** | Pointer to [**AssistantConversationFailureBody**](AssistantConversationFailureBody.md) |  | [optional] 
**Messages** | [**[]AssistantChatPollMessage**](AssistantChatPollMessage.md) |  | 
**Status** | **string** |  | 

## Methods

### NewAssistantChatPollBody

`func NewAssistantChatPollBody(actions []AssistantChatPollAction, messages []AssistantChatPollMessage, status string, ) *AssistantChatPollBody`

NewAssistantChatPollBody instantiates a new AssistantChatPollBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantChatPollBodyWithDefaults

`func NewAssistantChatPollBodyWithDefaults() *AssistantChatPollBody`

NewAssistantChatPollBodyWithDefaults instantiates a new AssistantChatPollBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AssistantChatPollBody) GetActions() []AssistantChatPollAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AssistantChatPollBody) GetActionsOk() (*[]AssistantChatPollAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AssistantChatPollBody) SetActions(v []AssistantChatPollAction)`

SetActions sets Actions field to given value.


### GetFailure

`func (o *AssistantChatPollBody) GetFailure() AssistantConversationFailureBody`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *AssistantChatPollBody) GetFailureOk() (*AssistantConversationFailureBody, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *AssistantChatPollBody) SetFailure(v AssistantConversationFailureBody)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *AssistantChatPollBody) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetMessages

`func (o *AssistantChatPollBody) GetMessages() []AssistantChatPollMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AssistantChatPollBody) GetMessagesOk() (*[]AssistantChatPollMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AssistantChatPollBody) SetMessages(v []AssistantChatPollMessage)`

SetMessages sets Messages field to given value.


### GetStatus

`func (o *AssistantChatPollBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssistantChatPollBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssistantChatPollBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


