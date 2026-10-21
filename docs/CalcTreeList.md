# CalcTreeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axis2** | [**NullableCalcSegmentation**](CalcSegmentation.md) |  | 
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | Pointer to [**MaybeVarFunc**](MaybeVarFunc.md) |  | [optional] 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | Pointer to [**MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**Tree** | [**NullablePropertySource**](PropertySource.md) |  | 

## Methods

### NewCalcTreeList

`func NewCalcTreeList(axis2 NullableCalcSegmentation, filters []FilterGroup, params map[string]ParameterType, tree NullablePropertySource, ) *CalcTreeList`

NewCalcTreeList instantiates a new CalcTreeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTreeListWithDefaults

`func NewCalcTreeListWithDefaults() *CalcTreeList`

NewCalcTreeListWithDefaults instantiates a new CalcTreeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis2

`func (o *CalcTreeList) GetAxis2() CalcSegmentation`

GetAxis2 returns the Axis2 field if non-nil, zero value otherwise.

### GetAxis2Ok

`func (o *CalcTreeList) GetAxis2Ok() (*CalcSegmentation, bool)`

GetAxis2Ok returns a tuple with the Axis2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis2

`func (o *CalcTreeList) SetAxis2(v CalcSegmentation)`

SetAxis2 sets Axis2 field to given value.


### SetAxis2Nil

`func (o *CalcTreeList) SetAxis2Nil(b bool)`

 SetAxis2Nil sets the value for Axis2 to be an explicit nil

### UnsetAxis2
`func (o *CalcTreeList) UnsetAxis2()`

UnsetAxis2 ensures that no value is present for Axis2, not even an explicit nil
### GetDate

`func (o *CalcTreeList) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcTreeList) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcTreeList) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcTreeList) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcTreeList) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcTreeList) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcTreeList) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcTreeList) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcTreeList) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcTreeList) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *CalcTreeList) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetParams

`func (o *CalcTreeList) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcTreeList) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcTreeList) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcTreeList) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcTreeList) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcTreeList) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *CalcTreeList) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetTree

`func (o *CalcTreeList) GetTree() PropertySource`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CalcTreeList) GetTreeOk() (*PropertySource, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CalcTreeList) SetTree(v PropertySource)`

SetTree sets Tree field to given value.


### SetTreeNil

`func (o *CalcTreeList) SetTreeNil(b bool)`

 SetTreeNil sets the value for Tree to be an explicit nil

### UnsetTree
`func (o *CalcTreeList) UnsetTree()`

UnsetTree ensures that no value is present for Tree, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


