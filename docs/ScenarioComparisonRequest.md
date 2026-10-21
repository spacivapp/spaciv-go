# ScenarioComparisonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**From** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Function** | [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | Pointer to **[]string** |  | [optional] 
**ScenarioRefs** | Pointer to [**[]MaybeVarStr**](MaybeVarStr.md) |  | [optional] 
**To** | [**MaybeVarInt**](MaybeVarInt.md) |  | 

## Methods

### NewScenarioComparisonRequest

`func NewScenarioComparisonRequest(filters []FilterGroup, from MaybeVarInt, function MaybeVarFuncRequest, params map[string]ParameterTypeRequest, to MaybeVarInt, ) *ScenarioComparisonRequest`

NewScenarioComparisonRequest instantiates a new ScenarioComparisonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioComparisonRequestWithDefaults

`func NewScenarioComparisonRequestWithDefaults() *ScenarioComparisonRequest`

NewScenarioComparisonRequestWithDefaults instantiates a new ScenarioComparisonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ScenarioComparisonRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ScenarioComparisonRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ScenarioComparisonRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFrom

`func (o *ScenarioComparisonRequest) GetFrom() MaybeVarInt`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ScenarioComparisonRequest) GetFromOk() (*MaybeVarInt, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ScenarioComparisonRequest) SetFrom(v MaybeVarInt)`

SetFrom sets From field to given value.


### GetFunction

`func (o *ScenarioComparisonRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ScenarioComparisonRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ScenarioComparisonRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.


### GetParams

`func (o *ScenarioComparisonRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScenarioComparisonRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScenarioComparisonRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *ScenarioComparisonRequest) GetScenario() []string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ScenarioComparisonRequest) GetScenarioOk() (*[]string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ScenarioComparisonRequest) SetScenario(v []string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *ScenarioComparisonRequest) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetScenarioRefs

`func (o *ScenarioComparisonRequest) GetScenarioRefs() []MaybeVarStr`

GetScenarioRefs returns the ScenarioRefs field if non-nil, zero value otherwise.

### GetScenarioRefsOk

`func (o *ScenarioComparisonRequest) GetScenarioRefsOk() (*[]MaybeVarStr, bool)`

GetScenarioRefsOk returns a tuple with the ScenarioRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioRefs

`func (o *ScenarioComparisonRequest) SetScenarioRefs(v []MaybeVarStr)`

SetScenarioRefs sets ScenarioRefs field to given value.

### HasScenarioRefs

`func (o *ScenarioComparisonRequest) HasScenarioRefs() bool`

HasScenarioRefs returns a boolean if a field has been set.

### GetTo

`func (o *ScenarioComparisonRequest) GetTo() MaybeVarInt`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ScenarioComparisonRequest) GetToOk() (*MaybeVarInt, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ScenarioComparisonRequest) SetTo(v MaybeVarInt)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


