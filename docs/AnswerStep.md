# AnswerStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependency** | [**NullableStepDependency**](StepDependency.md) |  | 
**Description** | **string** |  | 
**EstimationDefault** | [**NullableAnswer**](Answer.md) |  | 
**Heading** | **string** |  | 
**Info** | **string** |  | 
**Initial** | [**NullableAnswer**](Answer.md) |  | 
**InputLabel** | **string** |  | 
**ObjectGroup** | **string** |  | 
**Options** | [**[]AnswerOption**](AnswerOption.md) |  | 
**Property** | **string** |  | 
**SendsToCalculation** | **bool** |  | 
**SpacedHeading** | **bool** |  | 
**FreeText** | [**FreeTextType**](FreeTextType.md) |  | 
**Number** | [**NumberType**](NumberType.md) |  | 
**RankedPicks** | [**RankedPicksType**](RankedPicksType.md) |  | 
**Select** | [**SelectType**](SelectType.md) |  | 
**TimeRanges** | [**TimeRangesType**](TimeRangesType.md) |  | 
**ValuePerOption** | [**ValuePerOptionType**](ValuePerOptionType.md) |  | 

## Methods

### NewAnswerStep

`func NewAnswerStep(dependency NullableStepDependency, description string, estimationDefault NullableAnswer, heading string, info string, initial NullableAnswer, inputLabel string, objectGroup string, options []AnswerOption, property string, sendsToCalculation bool, spacedHeading bool, freeText FreeTextType, number NumberType, rankedPicks RankedPicksType, select_ SelectType, timeRanges TimeRangesType, valuePerOption ValuePerOptionType, ) *AnswerStep`

NewAnswerStep instantiates a new AnswerStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerStepWithDefaults

`func NewAnswerStepWithDefaults() *AnswerStep`

NewAnswerStepWithDefaults instantiates a new AnswerStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependency

`func (o *AnswerStep) GetDependency() StepDependency`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *AnswerStep) GetDependencyOk() (*StepDependency, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *AnswerStep) SetDependency(v StepDependency)`

SetDependency sets Dependency field to given value.


### SetDependencyNil

`func (o *AnswerStep) SetDependencyNil(b bool)`

 SetDependencyNil sets the value for Dependency to be an explicit nil

### UnsetDependency
`func (o *AnswerStep) UnsetDependency()`

UnsetDependency ensures that no value is present for Dependency, not even an explicit nil
### GetDescription

`func (o *AnswerStep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnswerStep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnswerStep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEstimationDefault

`func (o *AnswerStep) GetEstimationDefault() Answer`

GetEstimationDefault returns the EstimationDefault field if non-nil, zero value otherwise.

### GetEstimationDefaultOk

`func (o *AnswerStep) GetEstimationDefaultOk() (*Answer, bool)`

GetEstimationDefaultOk returns a tuple with the EstimationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimationDefault

`func (o *AnswerStep) SetEstimationDefault(v Answer)`

SetEstimationDefault sets EstimationDefault field to given value.


### SetEstimationDefaultNil

`func (o *AnswerStep) SetEstimationDefaultNil(b bool)`

 SetEstimationDefaultNil sets the value for EstimationDefault to be an explicit nil

### UnsetEstimationDefault
`func (o *AnswerStep) UnsetEstimationDefault()`

UnsetEstimationDefault ensures that no value is present for EstimationDefault, not even an explicit nil
### GetHeading

`func (o *AnswerStep) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *AnswerStep) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *AnswerStep) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetInfo

`func (o *AnswerStep) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AnswerStep) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AnswerStep) SetInfo(v string)`

SetInfo sets Info field to given value.


### GetInitial

`func (o *AnswerStep) GetInitial() Answer`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *AnswerStep) GetInitialOk() (*Answer, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *AnswerStep) SetInitial(v Answer)`

SetInitial sets Initial field to given value.


### SetInitialNil

`func (o *AnswerStep) SetInitialNil(b bool)`

 SetInitialNil sets the value for Initial to be an explicit nil

### UnsetInitial
`func (o *AnswerStep) UnsetInitial()`

UnsetInitial ensures that no value is present for Initial, not even an explicit nil
### GetInputLabel

`func (o *AnswerStep) GetInputLabel() string`

GetInputLabel returns the InputLabel field if non-nil, zero value otherwise.

### GetInputLabelOk

`func (o *AnswerStep) GetInputLabelOk() (*string, bool)`

GetInputLabelOk returns a tuple with the InputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabel

`func (o *AnswerStep) SetInputLabel(v string)`

SetInputLabel sets InputLabel field to given value.


### GetObjectGroup

`func (o *AnswerStep) GetObjectGroup() string`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *AnswerStep) GetObjectGroupOk() (*string, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *AnswerStep) SetObjectGroup(v string)`

SetObjectGroup sets ObjectGroup field to given value.


### GetOptions

`func (o *AnswerStep) GetOptions() []AnswerOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AnswerStep) GetOptionsOk() (*[]AnswerOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AnswerStep) SetOptions(v []AnswerOption)`

SetOptions sets Options field to given value.


### GetProperty

`func (o *AnswerStep) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AnswerStep) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AnswerStep) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetSendsToCalculation

`func (o *AnswerStep) GetSendsToCalculation() bool`

GetSendsToCalculation returns the SendsToCalculation field if non-nil, zero value otherwise.

### GetSendsToCalculationOk

`func (o *AnswerStep) GetSendsToCalculationOk() (*bool, bool)`

GetSendsToCalculationOk returns a tuple with the SendsToCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendsToCalculation

`func (o *AnswerStep) SetSendsToCalculation(v bool)`

SetSendsToCalculation sets SendsToCalculation field to given value.


### GetSpacedHeading

`func (o *AnswerStep) GetSpacedHeading() bool`

GetSpacedHeading returns the SpacedHeading field if non-nil, zero value otherwise.

### GetSpacedHeadingOk

`func (o *AnswerStep) GetSpacedHeadingOk() (*bool, bool)`

GetSpacedHeadingOk returns a tuple with the SpacedHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacedHeading

`func (o *AnswerStep) SetSpacedHeading(v bool)`

SetSpacedHeading sets SpacedHeading field to given value.


### GetFreeText

`func (o *AnswerStep) GetFreeText() FreeTextType`

GetFreeText returns the FreeText field if non-nil, zero value otherwise.

### GetFreeTextOk

`func (o *AnswerStep) GetFreeTextOk() (*FreeTextType, bool)`

GetFreeTextOk returns a tuple with the FreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeText

`func (o *AnswerStep) SetFreeText(v FreeTextType)`

SetFreeText sets FreeText field to given value.


### GetNumber

`func (o *AnswerStep) GetNumber() NumberType`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AnswerStep) GetNumberOk() (*NumberType, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AnswerStep) SetNumber(v NumberType)`

SetNumber sets Number field to given value.


### GetRankedPicks

`func (o *AnswerStep) GetRankedPicks() RankedPicksType`

GetRankedPicks returns the RankedPicks field if non-nil, zero value otherwise.

### GetRankedPicksOk

`func (o *AnswerStep) GetRankedPicksOk() (*RankedPicksType, bool)`

GetRankedPicksOk returns a tuple with the RankedPicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedPicks

`func (o *AnswerStep) SetRankedPicks(v RankedPicksType)`

SetRankedPicks sets RankedPicks field to given value.


### GetSelect

`func (o *AnswerStep) GetSelect() SelectType`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *AnswerStep) GetSelectOk() (*SelectType, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *AnswerStep) SetSelect(v SelectType)`

SetSelect sets Select field to given value.


### GetTimeRanges

`func (o *AnswerStep) GetTimeRanges() TimeRangesType`

GetTimeRanges returns the TimeRanges field if non-nil, zero value otherwise.

### GetTimeRangesOk

`func (o *AnswerStep) GetTimeRangesOk() (*TimeRangesType, bool)`

GetTimeRangesOk returns a tuple with the TimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRanges

`func (o *AnswerStep) SetTimeRanges(v TimeRangesType)`

SetTimeRanges sets TimeRanges field to given value.


### GetValuePerOption

`func (o *AnswerStep) GetValuePerOption() ValuePerOptionType`

GetValuePerOption returns the ValuePerOption field if non-nil, zero value otherwise.

### GetValuePerOptionOk

`func (o *AnswerStep) GetValuePerOptionOk() (*ValuePerOptionType, bool)`

GetValuePerOptionOk returns a tuple with the ValuePerOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePerOption

`func (o *AnswerStep) SetValuePerOption(v ValuePerOptionType)`

SetValuePerOption sets ValuePerOption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


