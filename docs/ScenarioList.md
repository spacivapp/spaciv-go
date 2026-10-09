# ScenarioList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scenario** | Pointer to **[]string** |  | [optional] 
**ScenarioRefs** | Pointer to [**[]MaybeVarStr**](MaybeVarStr.md) |  | [optional] 

## Methods

### NewScenarioList

`func NewScenarioList() *ScenarioList`

NewScenarioList instantiates a new ScenarioList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioListWithDefaults

`func NewScenarioListWithDefaults() *ScenarioList`

NewScenarioListWithDefaults instantiates a new ScenarioList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenario

`func (o *ScenarioList) GetScenario() []string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ScenarioList) GetScenarioOk() (*[]string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ScenarioList) SetScenario(v []string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *ScenarioList) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetScenarioRefs

`func (o *ScenarioList) GetScenarioRefs() []MaybeVarStr`

GetScenarioRefs returns the ScenarioRefs field if non-nil, zero value otherwise.

### GetScenarioRefsOk

`func (o *ScenarioList) GetScenarioRefsOk() (*[]MaybeVarStr, bool)`

GetScenarioRefsOk returns a tuple with the ScenarioRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioRefs

`func (o *ScenarioList) SetScenarioRefs(v []MaybeVarStr)`

SetScenarioRefs sets ScenarioRefs field to given value.

### HasScenarioRefs

`func (o *ScenarioList) HasScenarioRefs() bool`

HasScenarioRefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


