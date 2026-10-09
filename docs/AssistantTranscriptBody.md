# AssistantTranscriptBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]AssistantTranscriptMessage**](AssistantTranscriptMessage.md) |  | 
**Summary** | **string** | What the assistant is told of the messages it no longer replays, once the conversation has been compacted. Empty until then. | 
**SummaryUpTo** | **int64** | How many of the leading messages the summary stands in for. | 

## Methods

### NewAssistantTranscriptBody

`func NewAssistantTranscriptBody(messages []AssistantTranscriptMessage, summary string, summaryUpTo int64, ) *AssistantTranscriptBody`

NewAssistantTranscriptBody instantiates a new AssistantTranscriptBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantTranscriptBodyWithDefaults

`func NewAssistantTranscriptBodyWithDefaults() *AssistantTranscriptBody`

NewAssistantTranscriptBodyWithDefaults instantiates a new AssistantTranscriptBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


