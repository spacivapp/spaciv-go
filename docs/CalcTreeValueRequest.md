# CalcTreeValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | Pointer to [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | [optional] 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | Pointer to [**MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**Tree** | [**PropertySource**](PropertySource.md) |  | 

## Methods

### NewCalcTreeValueRequest

`func NewCalcTreeValueRequest(filters []FilterGroup, params map[string]ParameterTypeRequest, tree PropertySource, ) *CalcTreeValueRequest`

NewCalcTreeValueRequest instantiates a new CalcTreeValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTreeValueRequestWithDefaults

`func NewCalcTreeValueRequestWithDefaults() *CalcTreeValueRequest`

NewCalcTreeValueRequestWithDefaults instantiates a new CalcTreeValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcTreeValueRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTreeValueRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTreeValueRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcTreeValueRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcTreeValueRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTreeValueRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTreeValueRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTreeValueRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTreeValueRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTreeValueRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *CalcTreeValueRequest) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetParams

`func (o *CalcTreeValueRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTreeValueRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTreeValueRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTreeValueRequest) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTreeValueRequest) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTreeValueRequest) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *CalcTreeValueRequest) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetTree

`func (o *CalcTreeValueRequest) GetTree() PropertySource`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CalcTreeValueRequest) GetTreeOk() (*PropertySource, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CalcTreeValueRequest) SetTree(v PropertySource)`

SetTree sets Tree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


