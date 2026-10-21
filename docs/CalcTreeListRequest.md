# CalcTreeListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axis2** | [**CalcSegmentationRequest**](CalcSegmentationRequest.md) |  | 
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | Pointer to [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | [optional] 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | Pointer to [**MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**Tree** | [**PropertySource**](PropertySource.md) |  | 

## Methods

### NewCalcTreeListRequest

`func NewCalcTreeListRequest(axis2 CalcSegmentationRequest, filters []FilterGroup, params map[string]ParameterTypeRequest, tree PropertySource, ) *CalcTreeListRequest`

NewCalcTreeListRequest instantiates a new CalcTreeListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTreeListRequestWithDefaults

`func NewCalcTreeListRequestWithDefaults() *CalcTreeListRequest`

NewCalcTreeListRequestWithDefaults instantiates a new CalcTreeListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis2

`func (o *CalcTreeListRequest) GetAxis2() CalcSegmentationRequest`

GetAxis2 returns the Axis2 field if non-nil, zero value otherwise.

### GetAxis2Ok

`func (o *CalcTreeListRequest) GetAxis2Ok() (*CalcSegmentationRequest, bool)`

GetAxis2Ok returns a tuple with the Axis2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis2

`func (o *CalcTreeListRequest) SetAxis2(v CalcSegmentationRequest)`

SetAxis2 sets Axis2 field to given value.


### GetDate

`func (o *CalcTreeListRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTreeListRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTreeListRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcTreeListRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcTreeListRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTreeListRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTreeListRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTreeListRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTreeListRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTreeListRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *CalcTreeListRequest) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetParams

`func (o *CalcTreeListRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTreeListRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTreeListRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTreeListRequest) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTreeListRequest) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTreeListRequest) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *CalcTreeListRequest) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetTree

`func (o *CalcTreeListRequest) GetTree() PropertySource`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CalcTreeListRequest) GetTreeOk() (*PropertySource, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CalcTreeListRequest) SetTree(v PropertySource)`

SetTree sets Tree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


