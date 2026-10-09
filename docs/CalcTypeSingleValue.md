# CalcTypeSingleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | [**MaybeVarFunc**](MaybeVarFunc.md) |  | 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 

## Methods

### NewCalcTypeSingleValue

`func NewCalcTypeSingleValue(date MaybeVarInt, filters []FilterGroup, function MaybeVarFunc, params map[string]ParameterType, scenario MaybeVarStr, ) *CalcTypeSingleValue`

NewCalcTypeSingleValue instantiates a new CalcTypeSingleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTypeSingleValueWithDefaults

`func NewCalcTypeSingleValueWithDefaults() *CalcTypeSingleValue`

NewCalcTypeSingleValueWithDefaults instantiates a new CalcTypeSingleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcTypeSingleValue) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTypeSingleValue) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTypeSingleValue) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### GetFilters

`func (o *CalcTypeSingleValue) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTypeSingleValue) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTypeSingleValue) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTypeSingleValue) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTypeSingleValue) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTypeSingleValue) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.


### GetParams

`func (o *CalcTypeSingleValue) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTypeSingleValue) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTypeSingleValue) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTypeSingleValue) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTypeSingleValue) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTypeSingleValue) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


