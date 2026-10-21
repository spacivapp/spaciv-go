# CalcTypeSingleValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 

## Methods

### NewCalcTypeSingleValueRequest

`func NewCalcTypeSingleValueRequest(date MaybeVarInt, filters []FilterGroup, function MaybeVarFuncRequest, params map[string]ParameterTypeRequest, scenario MaybeVarStr, ) *CalcTypeSingleValueRequest`

NewCalcTypeSingleValueRequest instantiates a new CalcTypeSingleValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTypeSingleValueRequestWithDefaults

`func NewCalcTypeSingleValueRequestWithDefaults() *CalcTypeSingleValueRequest`

NewCalcTypeSingleValueRequestWithDefaults instantiates a new CalcTypeSingleValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcTypeSingleValueRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTypeSingleValueRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTypeSingleValueRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### GetFilters

`func (o *CalcTypeSingleValueRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTypeSingleValueRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTypeSingleValueRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTypeSingleValueRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTypeSingleValueRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTypeSingleValueRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.


### GetParams

`func (o *CalcTypeSingleValueRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTypeSingleValueRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTypeSingleValueRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTypeSingleValueRequest) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTypeSingleValueRequest) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTypeSingleValueRequest) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


