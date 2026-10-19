# AssistantTranscriptBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** |  | 
**Messages** | [**[]AssistantTranscriptMessage**](AssistantTranscriptMessage.md) |  | 
**ParentToolCallId** | **string** | For a subsession, the id of the run_subsession tool call that started it. Empty for a conversation the user started. | 
**Subsessions** | [**[]AssistantTranscriptBody**](AssistantTranscriptBody.md) | The transcript of each subsession the assistant started, in the order it started them. A subsession&#39;s messages are not part of messages: the tool message answering its run_subsession call holds only its reply. A subsession is never compacted and starts none of its own, so its summary fields and its own subsessions are empty. | 
**Summary** | **string** | What the assistant is told of the messages it no longer replays, once the conversation has been compacted. Empty until then. | 
**SummaryUpTo** | **int64** | How many of the leading messages the summary stands in for. | 

## Methods

### NewAssistantTranscriptBody

`func NewAssistantTranscriptBody(conversationId string, messages []AssistantTranscriptMessage, parentToolCallId string, subsessions []AssistantTranscriptBody, summary string, summaryUpTo int64, ) *AssistantTranscriptBody`

NewAssistantTranscriptBody instantiates a new AssistantTranscriptBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantTranscriptBodyWithDefaults

`func NewAssistantTranscriptBodyWithDefaults() *AssistantTranscriptBody`

NewAssistantTranscriptBodyWithDefaults instantiates a new AssistantTranscriptBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *AssistantTranscriptBody) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AssistantTranscriptBody) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AssistantTranscriptBody) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetMessages

`func (o *AssistantTranscriptBody) GetMessages() []AssistantTranscriptMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AssistantTranscriptBody) GetMessagesOk() (*[]AssistantTranscriptMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AssistantTranscriptBody) SetMessages(v []AssistantTranscriptMessage)`

SetMessages sets Messages field to given value.


### GetParentToolCallId

`func (o *AssistantTranscriptBody) GetParentToolCallId() string`

GetParentToolCallId returns the ParentToolCallId field if non-nil, zero value otherwise.

### GetParentToolCallIdOk

`func (o *AssistantTranscriptBody) GetParentToolCallIdOk() (*string, bool)`

GetParentToolCallIdOk returns a tuple with the ParentToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentToolCallId

`func (o *AssistantTranscriptBody) SetParentToolCallId(v string)`

SetParentToolCallId sets ParentToolCallId field to given value.


### GetSubsessions

`func (o *AssistantTranscriptBody) GetSubsessions() []AssistantTranscriptBody`

GetSubsessions returns the Subsessions field if non-nil, zero value otherwise.

### GetSubsessionsOk

`func (o *AssistantTranscriptBody) GetSubsessionsOk() (*[]AssistantTranscriptBody, bool)`

GetSubsessionsOk returns a tuple with the Subsessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsessions

`func (o *AssistantTranscriptBody) SetSubsessions(v []AssistantTranscriptBody)`

SetSubsessions sets Subsessions field to given value.


### GetSummary

`func (o *AssistantTranscriptBody) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AssistantTranscriptBody) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AssistantTranscriptBody) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetSummaryUpTo

`func (o *AssistantTranscriptBody) GetSummaryUpTo() int64`

GetSummaryUpTo returns the SummaryUpTo field if non-nil, zero value otherwise.

### GetSummaryUpToOk

`func (o *AssistantTranscriptBody) GetSummaryUpToOk() (*int64, bool)`

GetSummaryUpToOk returns a tuple with the SummaryUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryUpTo

`func (o *AssistantTranscriptBody) SetSummaryUpTo(v int64)`

SetSummaryUpTo sets SummaryUpTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


