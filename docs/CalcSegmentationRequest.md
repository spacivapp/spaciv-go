# CalcSegmentationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**CalcListRequest**](CalcListRequest.md) |  | 
**Property** | [**MaybeVarProp**](MaybeVarProp.md) |  | 
**Scenarios** | [**ScenarioList**](ScenarioList.md) |  | 

## Methods

### NewCalcSegmentationRequest

`func NewCalcSegmentationRequest(entries CalcListRequest, property MaybeVarProp, scenarios ScenarioList, ) *CalcSegmentationRequest`

NewCalcSegmentationRequest instantiates a new CalcSegmentationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcSegmentationRequestWithDefaults

`func NewCalcSegmentationRequestWithDefaults() *CalcSegmentationRequest`

NewCalcSegmentationRequestWithDefaults instantiates a new CalcSegmentationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *CalcSegmentationRequest) GetEntries() CalcListRequest`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CalcSegmentationRequest) GetEntriesOk() (*CalcListRequest, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CalcSegmentationRequest) SetEntries(v CalcListRequest)`

SetEntries sets Entries field to given value.


### GetProperty

`func (o *CalcSegmentationRequest) GetProperty() MaybeVarProp`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CalcSegmentationRequest) GetPropertyOk() (*MaybeVarProp, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CalcSegmentationRequest) SetProperty(v MaybeVarProp)`

SetProperty sets Property field to given value.


### GetScenarios

`func (o *CalcSegmentationRequest) GetScenarios() ScenarioList`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *CalcSegmentationRequest) GetScenariosOk() (*ScenarioList, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *CalcSegmentationRequest) SetScenarios(v ScenarioList)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


