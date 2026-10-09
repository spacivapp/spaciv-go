# AssistantTranscriptMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AssistantChatPollAttachment**](AssistantChatPollAttachment.md) | Files attached to a user message. Empty on every other message. | 
**CreatedAt** | **string** |  | 
**Failed** | **bool** | Whether the tool call a tool message answers failed: the tool refused it, returned an error, or was not available. | 
**Id** | **string** | The message&#39;s index in the conversation. A poll&#39;s message ids start with it. | 
**Reason** | **string** | Why a notice was given, as the id a poll shows it with. | 
**RequestId** | **string** |  | 
**Role** | **string** |  | 
**Text** | **string** | For a tool message, the result the assistant was given. | 
**ToolCallId** | **string** | The id of the tool call a tool message is the result of. | 
**ToolCalls** | [**[]AssistantTranscriptToolCall**](AssistantTranscriptToolCall.md) | The tools an assistant message called. Empty on every other message. | 

## Methods

### NewAssistantTranscriptMessage

`func NewAssistantTranscriptMessage(attachments []AssistantChatPollAttachment, createdAt string, failed bool, id string, reason string, requestId string, role string, text string, toolCallId string, toolCalls []AssistantTranscriptToolCall, ) *AssistantTranscriptMessage`

NewAssistantTranscriptMessage instantiates a new AssistantTranscriptMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantTranscriptMessageWithDefaults

`func NewAssistantTranscriptMessageWithDefaults() *AssistantTranscriptMessage`

NewAssistantTranscriptMessageWithDefaults instantiates a new AssistantTranscriptMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *AssistantTranscriptMessage) GetAttachments() []AssistantChatPollAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AssistantTranscriptMessage) GetAttachmentsOk() (*[]AssistantChatPollAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AssistantTranscriptMessage) SetAttachments(v []AssistantChatPollAttachment)`

SetAttachments sets Attachments field to given value.


### GetCreatedAt

`func (o *AssistantTranscriptMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssistantTranscriptMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssistantTranscriptMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailed

`func (o *AssistantTranscriptMessage) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *AssistantTranscriptMessage) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *AssistantTranscriptMessage) SetFailed(v bool)`

SetFailed sets Failed field to given value.


### GetId

`func (o *AssistantTranscriptMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssistantTranscriptMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssistantTranscriptMessage) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *AssistantTranscriptMessage) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssistantTranscriptMessage) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssistantTranscriptMessage) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRequestId

`func (o *AssistantTranscriptMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssistantTranscriptMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssistantTranscriptMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRole

`func (o *AssistantTranscriptMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssistantTranscriptMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssistantTranscriptMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetText

`func (o *AssistantTranscriptMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AssistantTranscriptMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AssistantTranscriptMessage) SetText(v string)`

SetText sets Text field to given value.


### GetToolCallId

`func (o *AssistantTranscriptMessage) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *AssistantTranscriptMessage) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *AssistantTranscriptMessage) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetToolCalls

`func (o *AssistantTranscriptMessage) GetToolCalls() []AssistantTranscriptToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AssistantTranscriptMessage) GetToolCallsOk() (*[]AssistantTranscriptToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AssistantTranscriptMessage) SetToolCalls(v []AssistantTranscriptToolCall)`

SetToolCalls sets ToolCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


