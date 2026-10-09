# CalcTreeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | Pointer to [**MaybeVarFunc**](MaybeVarFunc.md) |  | [optional] 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | Pointer to [**MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**Tree** | [**PropertySource**](PropertySource.md) |  | 

## Methods

### NewCalcTreeValue

`func NewCalcTreeValue(filters []FilterGroup, params map[string]ParameterType, tree PropertySource, ) *CalcTreeValue`

NewCalcTreeValue instantiates a new CalcTreeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTreeValueWithDefaults

`func NewCalcTreeValueWithDefaults() *CalcTreeValue`

NewCalcTreeValueWithDefaults instantiates a new CalcTreeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcTreeValue) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTreeValue) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTreeValue) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcTreeValue) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcTreeValue) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTreeValue) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTreeValue) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTreeValue) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTreeValue) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTreeValue) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *CalcTreeValue) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetParams

`func (o *CalcTreeValue) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTreeValue) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTreeValue) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTreeValue) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTreeValue) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTreeValue) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *CalcTreeValue) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetTree

`func (o *CalcTreeValue) GetTree() PropertySource`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CalcTreeValue) GetTreeOk() (*PropertySource, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CalcTreeValue) SetTree(v PropertySource)`

SetTree sets Tree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


