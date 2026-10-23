# EstimationRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | [**NullableAnswer**](Answer.md) |  | 
**Step** | [**NullableQuestionStep**](QuestionStep.md) |  | 

## Methods

### NewEstimationRes

`func NewEstimationRes(answer NullableAnswer, step NullableQuestionStep, ) *EstimationRes`

NewEstimationRes instantiates a new EstimationRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimationResWithDefaults

`func NewEstimationResWithDefaults() *EstimationRes`

NewEstimationResWithDefaults instantiates a new EstimationRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *EstimationRes) GetAnswer() Answer`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *EstimationRes) GetAnswerOk() (*Answer, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *EstimationRes) SetAnswer(v Answer)`

SetAnswer sets Answer field to given value.


### SetAnswerNil

`func (o *EstimationRes) SetAnswerNil(b bool)`

 SetAnswerNil sets the value for Answer to be an explicit nil

### UnsetAnswer
`func (o *EstimationRes) UnsetAnswer()`

UnsetAnswer ensures that no value is present for Answer, not even an explicit nil
### GetStep

`func (o *EstimationRes) GetStep() QuestionStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *EstimationRes) GetStepOk() (*QuestionStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *EstimationRes) SetStep(v QuestionStep)`

SetStep sets Step field to given value.


### SetStepNil

`func (o *EstimationRes) SetStepNil(b bool)`

 SetStepNil sets the value for Step to be an explicit nil

### UnsetStep
`func (o *EstimationRes) UnsetStep()`

UnsetStep ensures that no value is present for Step, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


