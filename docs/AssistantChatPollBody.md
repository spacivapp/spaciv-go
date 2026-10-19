# AssistantChatPollBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]AssistantChatPollAction**](AssistantChatPollAction.md) |  | 
**Failure** | Pointer to [**AssistantConversationFailureBody**](AssistantConversationFailureBody.md) |  | [optional] 
**FinishedToolCallIds** | **[]string** | Tool calls whose results, failed or not, were recorded after the since cursor, by the toolCallId of their action message. A call that succeeded has no message of its own, so this is how a client tells a finished action from one still running. A later poll can list a call again. | 
**Messages** | [**[]AssistantChatPollMessage**](AssistantChatPollMessage.md) |  | 
**Status** | **string** |  | 
**SubsessionSteps** | **int64** | While a subsession is running, how many messages it has recorded so far, which grows with each of its steps even when it calls the same tool again. Zero when no subsession is running, including once it has ended and before the turn has recorded its result. | 
**SubsessionTool** | **string** | While a subsession is running, the tool it is running now. Empty while it is thinking, and when no subsession is running. The running subsession is the run_subsession action whose toolCallId is not yet among the finished tool calls. | 

## Methods

### NewAssistantChatPollBody

`func NewAssistantChatPollBody(actions []AssistantChatPollAction, finishedToolCallIds []string, messages []AssistantChatPollMessage, status string, subsessionSteps int64, subsessionTool string, ) *AssistantChatPollBody`

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

### GetFinishedToolCallIds

`func (o *AssistantChatPollBody) GetFinishedToolCallIds() []string`

GetFinishedToolCallIds returns the FinishedToolCallIds field if non-nil, zero value otherwise.

### GetFinishedToolCallIdsOk

`func (o *AssistantChatPollBody) GetFinishedToolCallIdsOk() (*[]string, bool)`

GetFinishedToolCallIdsOk returns a tuple with the FinishedToolCallIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedToolCallIds

`func (o *AssistantChatPollBody) SetFinishedToolCallIds(v []string)`

SetFinishedToolCallIds sets FinishedToolCallIds field to given value.


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


### GetSubsessionSteps

`func (o *AssistantChatPollBody) GetSubsessionSteps() int64`

GetSubsessionSteps returns the SubsessionSteps field if non-nil, zero value otherwise.

### GetSubsessionStepsOk

`func (o *AssistantChatPollBody) GetSubsessionStepsOk() (*int64, bool)`

GetSubsessionStepsOk returns a tuple with the SubsessionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsessionSteps

`func (o *AssistantChatPollBody) SetSubsessionSteps(v int64)`

SetSubsessionSteps sets SubsessionSteps field to given value.


### GetSubsessionTool

`func (o *AssistantChatPollBody) GetSubsessionTool() string`

GetSubsessionTool returns the SubsessionTool field if non-nil, zero value otherwise.

### GetSubsessionToolOk

`func (o *AssistantChatPollBody) GetSubsessionToolOk() (*string, bool)`

GetSubsessionToolOk returns a tuple with the SubsessionTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsessionTool

`func (o *AssistantChatPollBody) SetSubsessionTool(v string)`

SetSubsessionTool sets SubsessionTool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


