# AssistantChatPollBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]AssistantChatPollAction**](AssistantChatPollAction.md) | What the turns that finished after the since cursor asked the app to do, each given once, in order. A read from the start gives those of every turn it reads, which a client restoring the conversation has already acted on: it should drop them. | 
**ActiveRequestId** | Pointer to **string** | The id of the chat request whose turn is running, or is next to, while the conversation is pending. A client that sent a request knows its turn has ended once this names another. | [optional] 
**BackgroundSubsessionIds** | **[]string** | The subsessions and script runs the conversation started in the background that are still running, which carry on after the turn that started them has finished: a script run is listed by its run id, as scriptRuns gives it for its run_script call. A client keeps polling while there are any: the message delivering a finished one&#39;s result starts a new turn. The call that started one is not listed among the finished tool calls until then. | 
**Confirmations** | [**[]AssistantChatPollConfirmation**](AssistantChatPollConfirmation.md) | What tool calls whose results were recorded after the since cursor asked the user to confirm. A tool never carries it out itself: the app shows it, and carries it out with the user&#39;s own request if they confirm. A later poll can list a call again. | 
**Failure** | Pointer to [**AssistantConversationFailureBody**](AssistantConversationFailureBody.md) |  | [optional] 
**Failures** | Pointer to [**[]AssistantConversationFailureBody**](AssistantConversationFailureBody.md) | The failures of every turn that ended after the since cursor, in order, each with the request it was of, when more than one turn failed between two polls. failure is the last of them. | [optional] 
**FinishedToolCallIds** | **[]string** | Tool calls whose results, failed or not, were recorded after the since cursor, by the toolCallId of their action message. A call that succeeded has no message of its own, so this is how a client tells a finished action from one still running. A run_subsession call that started a background subsession, or a run_script call that started a background script run, is listed once it has ended and its result has been delivered to the conversation. A later poll can list a call again. | 
**Messages** | [**[]AssistantChatPollMessage**](AssistantChatPollMessage.md) |  | 
**QueuedTurns** | **int64** | Messages queued behind the turn that has just finished, such as a background subsession&#39;s result, which each run as a turn of their own. A client keeps polling while there are any, even though the conversation is finished until the next one starts. | 
**ScriptRuns** | [**[]AssistantChatPollScriptRun**](AssistantChatPollScriptRun.md) | The scratchpad run each run_script call returned, for calls whose results were recorded after the since cursor. Only a call that started its run is listed, including one whose wait for the run then failed: one that only waited on a run by its ID is not. The run, its code and its output are read from GET /v1/scripts/scratchpad/runs/{runId} by the user who started it, while it is kept. A later poll can list a call again. | 
**Status** | **string** |  | 
**SubsessionSteps** | **int64** | While a subsession is running, how many messages it has recorded so far, which grows with each of its steps even when it calls the same tool again. Zero when no subsession is running, including once it has ended and before the turn has recorded its result. | 
**SubsessionTool** | **string** | While a subsession is running, the tool it is running now. Empty while it is thinking, and when no subsession is running. The running subsession is the run_subsession action whose toolCallId is not yet among the finished tool calls. | 

## Methods

### NewAssistantChatPollBody

`func NewAssistantChatPollBody(actions []AssistantChatPollAction, backgroundSubsessionIds []string, confirmations []AssistantChatPollConfirmation, finishedToolCallIds []string, messages []AssistantChatPollMessage, queuedTurns int64, scriptRuns []AssistantChatPollScriptRun, status string, subsessionSteps int64, subsessionTool string, ) *AssistantChatPollBody`

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


### GetActiveRequestId

`func (o *AssistantChatPollBody) GetActiveRequestId() string`

GetActiveRequestId returns the ActiveRequestId field if non-nil, zero value otherwise.

### GetActiveRequestIdOk

`func (o *AssistantChatPollBody) GetActiveRequestIdOk() (*string, bool)`

GetActiveRequestIdOk returns a tuple with the ActiveRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRequestId

`func (o *AssistantChatPollBody) SetActiveRequestId(v string)`

SetActiveRequestId sets ActiveRequestId field to given value.

### HasActiveRequestId

`func (o *AssistantChatPollBody) HasActiveRequestId() bool`

HasActiveRequestId returns a boolean if a field has been set.

### GetBackgroundSubsessionIds

`func (o *AssistantChatPollBody) GetBackgroundSubsessionIds() []string`

GetBackgroundSubsessionIds returns the BackgroundSubsessionIds field if non-nil, zero value otherwise.

### GetBackgroundSubsessionIdsOk

`func (o *AssistantChatPollBody) GetBackgroundSubsessionIdsOk() (*[]string, bool)`

GetBackgroundSubsessionIdsOk returns a tuple with the BackgroundSubsessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundSubsessionIds

`func (o *AssistantChatPollBody) SetBackgroundSubsessionIds(v []string)`

SetBackgroundSubsessionIds sets BackgroundSubsessionIds field to given value.


### GetConfirmations

`func (o *AssistantChatPollBody) GetConfirmations() []AssistantChatPollConfirmation`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *AssistantChatPollBody) GetConfirmationsOk() (*[]AssistantChatPollConfirmation, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *AssistantChatPollBody) SetConfirmations(v []AssistantChatPollConfirmation)`

SetConfirmations sets Confirmations field to given value.


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

### GetFailures

`func (o *AssistantChatPollBody) GetFailures() []AssistantConversationFailureBody`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *AssistantChatPollBody) GetFailuresOk() (*[]AssistantConversationFailureBody, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *AssistantChatPollBody) SetFailures(v []AssistantConversationFailureBody)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *AssistantChatPollBody) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

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


### GetQueuedTurns

`func (o *AssistantChatPollBody) GetQueuedTurns() int64`

GetQueuedTurns returns the QueuedTurns field if non-nil, zero value otherwise.

### GetQueuedTurnsOk

`func (o *AssistantChatPollBody) GetQueuedTurnsOk() (*int64, bool)`

GetQueuedTurnsOk returns a tuple with the QueuedTurns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedTurns

`func (o *AssistantChatPollBody) SetQueuedTurns(v int64)`

SetQueuedTurns sets QueuedTurns field to given value.


### GetScriptRuns

`func (o *AssistantChatPollBody) GetScriptRuns() []AssistantChatPollScriptRun`

GetScriptRuns returns the ScriptRuns field if non-nil, zero value otherwise.

### GetScriptRunsOk

`func (o *AssistantChatPollBody) GetScriptRunsOk() (*[]AssistantChatPollScriptRun, bool)`

GetScriptRunsOk returns a tuple with the ScriptRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptRuns

`func (o *AssistantChatPollBody) SetScriptRuns(v []AssistantChatPollScriptRun)`

SetScriptRuns sets ScriptRuns field to given value.


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


