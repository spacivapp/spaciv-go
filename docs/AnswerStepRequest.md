# AnswerStepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependency** | [**StepDependency**](StepDependency.md) |  | 
**Description** | **string** |  | 
**EstimationDefault** | [**Answer**](Answer.md) |  | 
**Heading** | **string** |  | 
**Info** | **string** |  | 
**Initial** | [**Answer**](Answer.md) |  | 
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

### NewAnswerStepRequest

`func NewAnswerStepRequest(dependency StepDependency, description string, estimationDefault Answer, heading string, info string, initial Answer, inputLabel string, objectGroup string, options []AnswerOption, property string, sendsToCalculation bool, spacedHeading bool, freeText FreeTextType, number NumberType, rankedPicks RankedPicksType, select_ SelectType, timeRanges TimeRangesType, valuePerOption ValuePerOptionType, ) *AnswerStepRequest`

NewAnswerStepRequest instantiates a new AnswerStepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerStepRequestWithDefaults

`func NewAnswerStepRequestWithDefaults() *AnswerStepRequest`

NewAnswerStepRequestWithDefaults instantiates a new AnswerStepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependency

`func (o *AnswerStepRequest) GetDependency() StepDependency`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *AnswerStepRequest) GetDependencyOk() (*StepDependency, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *AnswerStepRequest) SetDependency(v StepDependency)`

SetDependency sets Dependency field to given value.


### GetDescription

`func (o *AnswerStepRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnswerStepRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnswerStepRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEstimationDefault

`func (o *AnswerStepRequest) GetEstimationDefault() Answer`

GetEstimationDefault returns the EstimationDefault field if non-nil, zero value otherwise.

### GetEstimationDefaultOk

`func (o *AnswerStepRequest) GetEstimationDefaultOk() (*Answer, bool)`

GetEstimationDefaultOk returns a tuple with the EstimationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimationDefault

`func (o *AnswerStepRequest) SetEstimationDefault(v Answer)`

SetEstimationDefault sets EstimationDefault field to given value.


### GetHeading

`func (o *AnswerStepRequest) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *AnswerStepRequest) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *AnswerStepRequest) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetInfo

`func (o *AnswerStepRequest) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AnswerStepRequest) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AnswerStepRequest) SetInfo(v string)`

SetInfo sets Info field to given value.


### GetInitial

`func (o *AnswerStepRequest) GetInitial() Answer`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *AnswerStepRequest) GetInitialOk() (*Answer, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *AnswerStepRequest) SetInitial(v Answer)`

SetInitial sets Initial field to given value.


### GetInputLabel

`func (o *AnswerStepRequest) GetInputLabel() string`

GetInputLabel returns the InputLabel field if non-nil, zero value otherwise.

### GetInputLabelOk

`func (o *AnswerStepRequest) GetInputLabelOk() (*string, bool)`

GetInputLabelOk returns a tuple with the InputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabel

`func (o *AnswerStepRequest) SetInputLabel(v string)`

SetInputLabel sets InputLabel field to given value.


### GetObjectGroup

`func (o *AnswerStepRequest) GetObjectGroup() string`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *AnswerStepRequest) GetObjectGroupOk() (*string, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *AnswerStepRequest) SetObjectGroup(v string)`

SetObjectGroup sets ObjectGroup field to given value.


### GetOptions

`func (o *AnswerStepRequest) GetOptions() []AnswerOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AnswerStepRequest) GetOptionsOk() (*[]AnswerOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AnswerStepRequest) SetOptions(v []AnswerOption)`

SetOptions sets Options field to given value.


### GetProperty

`func (o *AnswerStepRequest) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AnswerStepRequest) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AnswerStepRequest) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetSendsToCalculation

`func (o *AnswerStepRequest) GetSendsToCalculation() bool`

GetSendsToCalculation returns the SendsToCalculation field if non-nil, zero value otherwise.

### GetSendsToCalculationOk

`func (o *AnswerStepRequest) GetSendsToCalculationOk() (*bool, bool)`

GetSendsToCalculationOk returns a tuple with the SendsToCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendsToCalculation

`func (o *AnswerStepRequest) SetSendsToCalculation(v bool)`

SetSendsToCalculation sets SendsToCalculation field to given value.


### GetSpacedHeading

`func (o *AnswerStepRequest) GetSpacedHeading() bool`

GetSpacedHeading returns the SpacedHeading field if non-nil, zero value otherwise.

### GetSpacedHeadingOk

`func (o *AnswerStepRequest) GetSpacedHeadingOk() (*bool, bool)`

GetSpacedHeadingOk returns a tuple with the SpacedHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacedHeading

`func (o *AnswerStepRequest) SetSpacedHeading(v bool)`

SetSpacedHeading sets SpacedHeading field to given value.


### GetFreeText

`func (o *AnswerStepRequest) GetFreeText() FreeTextType`

GetFreeText returns the FreeText field if non-nil, zero value otherwise.

### GetFreeTextOk

`func (o *AnswerStepRequest) GetFreeTextOk() (*FreeTextType, bool)`

GetFreeTextOk returns a tuple with the FreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeText

`func (o *AnswerStepRequest) SetFreeText(v FreeTextType)`

SetFreeText sets FreeText field to given value.


### GetNumber

`func (o *AnswerStepRequest) GetNumber() NumberType`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AnswerStepRequest) GetNumberOk() (*NumberType, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AnswerStepRequest) SetNumber(v NumberType)`

SetNumber sets Number field to given value.


### GetRankedPicks

`func (o *AnswerStepRequest) GetRankedPicks() RankedPicksType`

GetRankedPicks returns the RankedPicks field if non-nil, zero value otherwise.

### GetRankedPicksOk

`func (o *AnswerStepRequest) GetRankedPicksOk() (*RankedPicksType, bool)`

GetRankedPicksOk returns a tuple with the RankedPicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedPicks

`func (o *AnswerStepRequest) SetRankedPicks(v RankedPicksType)`

SetRankedPicks sets RankedPicks field to given value.


### GetSelect

`func (o *AnswerStepRequest) GetSelect() SelectType`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *AnswerStepRequest) GetSelectOk() (*SelectType, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *AnswerStepRequest) SetSelect(v SelectType)`

SetSelect sets Select field to given value.


### GetTimeRanges

`func (o *AnswerStepRequest) GetTimeRanges() TimeRangesType`

GetTimeRanges returns the TimeRanges field if non-nil, zero value otherwise.

### GetTimeRangesOk

`func (o *AnswerStepRequest) GetTimeRangesOk() (*TimeRangesType, bool)`

GetTimeRangesOk returns a tuple with the TimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRanges

`func (o *AnswerStepRequest) SetTimeRanges(v TimeRangesType)`

SetTimeRanges sets TimeRanges field to given value.


### GetValuePerOption

`func (o *AnswerStepRequest) GetValuePerOption() ValuePerOptionType`

GetValuePerOption returns the ValuePerOption field if non-nil, zero value otherwise.

### GetValuePerOptionOk

`func (o *AnswerStepRequest) GetValuePerOptionOk() (*ValuePerOptionType, bool)`

GetValuePerOptionOk returns a tuple with the ValuePerOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePerOption

`func (o *AnswerStepRequest) SetValuePerOption(v ValuePerOptionType)`

SetValuePerOption sets ValuePerOption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


