# QuestionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Answer** | [**AnswerStep**](AnswerStep.md) |  | 
**Content** | [**ContentStep**](ContentStep.md) |  | 

## Methods

### NewQuestionStep

`func NewQuestionStep(id string, answer AnswerStep, content ContentStep, ) *QuestionStep`

NewQuestionStep instantiates a new QuestionStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionStepWithDefaults

`func NewQuestionStepWithDefaults() *QuestionStep`

NewQuestionStepWithDefaults instantiates a new QuestionStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionStep) SetId(v string)`

SetId sets Id field to given value.


### GetAnswer

`func (o *QuestionStep) GetAnswer() AnswerStep`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *QuestionStep) GetAnswerOk() (*AnswerStep, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *QuestionStep) SetAnswer(v AnswerStep)`

SetAnswer sets Answer field to given value.


### GetContent

`func (o *QuestionStep) GetContent() ContentStep`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *QuestionStep) GetContentOk() (*ContentStep, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *QuestionStep) SetContent(v ContentStep)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


