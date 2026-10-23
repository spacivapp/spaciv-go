# QuestionStepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Answer** | [**AnswerStepRequest**](AnswerStepRequest.md) |  | 
**Content** | [**ContentStep**](ContentStep.md) |  | 

## Methods

### NewQuestionStepRequest

`func NewQuestionStepRequest(id string, answer AnswerStepRequest, content ContentStep, ) *QuestionStepRequest`

NewQuestionStepRequest instantiates a new QuestionStepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionStepRequestWithDefaults

`func NewQuestionStepRequestWithDefaults() *QuestionStepRequest`

NewQuestionStepRequestWithDefaults instantiates a new QuestionStepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionStepRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionStepRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionStepRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAnswer

`func (o *QuestionStepRequest) GetAnswer() AnswerStepRequest`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *QuestionStepRequest) GetAnswerOk() (*AnswerStepRequest, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *QuestionStepRequest) SetAnswer(v AnswerStepRequest)`

SetAnswer sets Answer field to given value.


### GetContent

`func (o *QuestionStepRequest) GetContent() ContentStep`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *QuestionStepRequest) GetContentOk() (*ContentStep, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *QuestionStepRequest) SetContent(v ContentStep)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


