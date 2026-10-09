# ScenarioComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**From** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Function** | [**MaybeVarFunc**](MaybeVarFunc.md) |  | 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | Pointer to **[]string** |  | [optional] 
**ScenarioRefs** | Pointer to [**[]MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**To** | [**MaybeVarInt**](MaybeVarInt.md) |  | 

## Methods

### NewScenarioComparison

`func NewScenarioComparison(filters []FilterGroup, from MaybeVarInt, function MaybeVarFunc, params map[string]ParameterType, to MaybeVarInt, ) *ScenarioComparison`

NewScenarioComparison instantiates a new ScenarioComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioComparisonWithDefaults

`func NewScenarioComparisonWithDefaults() *ScenarioComparison`

NewScenarioComparisonWithDefaults instantiates a new ScenarioComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ScenarioComparison) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ScenarioComparison) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ScenarioComparison) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFrom

`func (o *ScenarioComparison) GetFrom() MaybeVarInt`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ScenarioComparison) GetFromOk() (*MaybeVarInt, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ScenarioComparison) SetFrom(v MaybeVarInt)`

SetFrom sets From field to given value.


### GetFunction

`func (o *ScenarioComparison) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ScenarioComparison) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ScenarioComparison) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.


### GetParams

`func (o *ScenarioComparison) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScenarioComparison) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScenarioComparison) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *ScenarioComparison) GetScenario() []string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ScenarioComparison) GetScenarioOk() (*[]string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ScenarioComparison) SetScenario(v []string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *ScenarioComparison) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetScenarioRefs

`func (o *ScenarioComparison) GetScenarioRefs() []MaybeVarStr`

GetScenarioRefs returns the ScenarioRefs field if non-nil, zero value otherwise.

### GetScenarioRefsOk

`func (o *ScenarioComparison) GetScenarioRefsOk() (*[]MaybeVarStr, bool)`

GetScenarioRefsOk returns a tuple with the ScenarioRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioRefs

`func (o *ScenarioComparison) SetScenarioRefs(v []MaybeVarStr)`

SetScenarioRefs sets ScenarioRefs field to given value.

### HasScenarioRefs

`func (o *ScenarioComparison) HasScenarioRefs() bool`

HasScenarioRefs returns a boolean if a field has been set.

### GetTo

`func (o *ScenarioComparison) GetTo() MaybeVarInt`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ScenarioComparison) GetToOk() (*MaybeVarInt, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ScenarioComparison) SetTo(v MaybeVarInt)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


