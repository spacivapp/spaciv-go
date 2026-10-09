# AssistantChatPollMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AssistantChatPollAttachment**](AssistantChatPollAttachment.md) | Files attached to a user message, as they were when it was sent. Empty on every other message. | 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**Kind** | **int64** |  | 
**Payload** | **string** |  | 
**RequestId** | **string** |  | 
**Role** | **string** |  | 
**Text** | **string** |  | 

## Methods

### NewAssistantChatPollMessage

`func NewAssistantChatPollMessage(attachments []AssistantChatPollAttachment, createdAt string, id string, kind int64, payload string, requestId string, role string, text string, ) *AssistantChatPollMessage`

NewAssistantChatPollMessage instantiates a new AssistantChatPollMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantChatPollMessageWithDefaults

`func NewAssistantChatPollMessageWithDefaults() *AssistantChatPollMessage`

NewAssistantChatPollMessageWithDefaults instantiates a new AssistantChatPollMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *AssistantChatPollMessage) GetAttachments() []AssistantChatPollAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AssistantChatPollMessage) GetAttachmentsOk() (*[]AssistantChatPollAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AssistantChatPollMessage) SetAttachments(v []AssistantChatPollAttachment)`

SetAttachments sets Attachments field to given value.


### GetCreatedAt

`func (o *AssistantChatPollMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssistantChatPollMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssistantChatPollMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *AssistantChatPollMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssistantChatPollMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssistantChatPollMessage) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *AssistantChatPollMessage) GetKind() int64`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AssistantChatPollMessage) GetKindOk() (*int64, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AssistantChatPollMessage) SetKind(v int64)`

SetKind sets Kind field to given value.


### GetPayload

`func (o *AssistantChatPollMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AssistantChatPollMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AssistantChatPollMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetRequestId

`func (o *AssistantChatPollMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssistantChatPollMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssistantChatPollMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRole

`func (o *AssistantChatPollMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssistantChatPollMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssistantChatPollMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetText

`func (o *AssistantChatPollMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AssistantChatPollMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AssistantChatPollMessage) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


