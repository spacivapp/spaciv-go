# CalcSegmentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**CalcList**](CalcList.md) |  | 
**Property** | [**MaybeVarProp**](MaybeVarProp.md) |  | 
**Scenarios** | [**ScenarioList**](ScenarioList.md) |  | 

## Methods

### NewCalcSegmentation

`func NewCalcSegmentation(entries CalcList, property MaybeVarProp, scenarios ScenarioList, ) *CalcSegmentation`

NewCalcSegmentation instantiates a new CalcSegmentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcSegmentationWithDefaults

`func NewCalcSegmentationWithDefaults() *CalcSegmentation`

NewCalcSegmentationWithDefaults instantiates a new CalcSegmentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *CalcSegmentation) GetEntries() CalcList`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CalcSegmentation) GetEntriesOk() (*CalcList, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CalcSegmentation) SetEntries(v CalcList)`

SetEntries sets Entries field to given value.


### GetProperty

`func (o *CalcSegmentation) GetProperty() MaybeVarProp`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CalcSegmentation) GetPropertyOk() (*MaybeVarProp, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CalcSegmentation) SetProperty(v MaybeVarProp)`

SetProperty sets Property field to given value.


### GetScenarios

`func (o *CalcSegmentation) GetScenarios() ScenarioList`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *CalcSegmentation) GetScenariosOk() (*ScenarioList, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *CalcSegmentation) SetScenarios(v ScenarioList)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


