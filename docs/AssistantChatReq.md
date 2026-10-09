# AssistantChatReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentFileIds** | Pointer to **[]string** |  | [optional] 
**Context** | Pointer to [**AssistantChatContext**](AssistantChatContext.md) |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**LlmModel** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewAssistantChatReq

`func NewAssistantChatReq(message string, ) *AssistantChatReq`

NewAssistantChatReq instantiates a new AssistantChatReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantChatReqWithDefaults

`func NewAssistantChatReqWithDefaults() *AssistantChatReq`

NewAssistantChatReqWithDefaults instantiates a new AssistantChatReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentFileIds

`func (o *AssistantChatReq) GetAttachmentFileIds() []string`

GetAttachmentFileIds returns the AttachmentFileIds field if non-nil, zero value otherwise.

### GetAttachmentFileIdsOk

`func (o *AssistantChatReq) GetAttachmentFileIdsOk() (*[]string, bool)`

GetAttachmentFileIdsOk returns a tuple with the AttachmentFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentFileIds

`func (o *AssistantChatReq) SetAttachmentFileIds(v []string)`

SetAttachmentFileIds sets AttachmentFileIds field to given value.

### HasAttachmentFileIds

`func (o *AssistantChatReq) HasAttachmentFileIds() bool`

HasAttachmentFileIds returns a boolean if a field has been set.

### GetContext

`func (o *AssistantChatReq) GetContext() AssistantChatContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AssistantChatReq) GetContextOk() (*AssistantChatContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AssistantChatReq) SetContext(v AssistantChatContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AssistantChatReq) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetConversationId

`func (o *AssistantChatReq) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AssistantChatReq) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AssistantChatReq) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *AssistantChatReq) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetLlmModel

`func (o *AssistantChatReq) GetLlmModel() string`

GetLlmModel returns the LlmModel field if non-nil, zero value otherwise.

### GetLlmModelOk

`func (o *AssistantChatReq) GetLlmModelOk() (*string, bool)`

GetLlmModelOk returns a tuple with the LlmModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmModel

`func (o *AssistantChatReq) SetLlmModel(v string)`

SetLlmModel sets LlmModel field to given value.

### HasLlmModel

`func (o *AssistantChatReq) HasLlmModel() bool`

HasLlmModel returns a boolean if a field has been set.

### GetMessage

`func (o *AssistantChatReq) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssistantChatReq) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssistantChatReq) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


