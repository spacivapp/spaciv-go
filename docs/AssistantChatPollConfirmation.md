# AssistantChatPollConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** | The type&#39;s JSON body. | 
**ToolCallId** | **string** | The toolCallId of the action message that asked for it. | 
**Type** | **string** | What is to be confirmed, which says what payload holds. create_script: create a script, with the user confirming its name and description, from payload {name, description, source, project, schedule, webhook, enabled, secretNames}, where source is {accountFiles: {language, directoryPath, files, manifest}}, whose files map each of the script&#39;s file paths, and manifest the dependency manifest, to an account file whose content it takes; the settings after source are optional, project is the project the script acts in, and secretNames are secrets whose values the user enters. update_script: change a script, from payload {scriptId, name, revision, changes}, where changes holds only what changes, with the same fields and meanings as create_script&#39;s payload, a project of 0 removing the project and a schedule with an empty cron removing the schedule, and revision is the one the change was proposed against. delete_script: delete a script, from payload {scriptId, name}. Clients ignore types they do not know. | 

## Methods

### NewAssistantChatPollConfirmation

`func NewAssistantChatPollConfirmation(payload string, toolCallId string, type_ string, ) *AssistantChatPollConfirmation`

NewAssistantChatPollConfirmation instantiates a new AssistantChatPollConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantChatPollConfirmationWithDefaults

`func NewAssistantChatPollConfirmationWithDefaults() *AssistantChatPollConfirmation`

NewAssistantChatPollConfirmationWithDefaults instantiates a new AssistantChatPollConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *AssistantChatPollConfirmation) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AssistantChatPollConfirmation) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AssistantChatPollConfirmation) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetToolCallId

`func (o *AssistantChatPollConfirmation) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *AssistantChatPollConfirmation) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *AssistantChatPollConfirmation) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetType

`func (o *AssistantChatPollConfirmation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssistantChatPollConfirmation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssistantChatPollConfirmation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


