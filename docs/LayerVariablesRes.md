# LayerVariablesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scenarios** | **map[string][]string** | For each active scenario whose layers&#39; modifier formulas use account variables, those variable ids (account:&lt;id&gt;), including the layers of stacked scenarios. | 

## Methods

### NewLayerVariablesRes

`func NewLayerVariablesRes(scenarios map[string][]string, ) *LayerVariablesRes`

NewLayerVariablesRes instantiates a new LayerVariablesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerVariablesResWithDefaults

`func NewLayerVariablesResWithDefaults() *LayerVariablesRes`

NewLayerVariablesResWithDefaults instantiates a new LayerVariablesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenarios

`func (o *LayerVariablesRes) GetScenarios() map[string][]string`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *LayerVariablesRes) GetScenariosOk() (*map[string][]string, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *LayerVariablesRes) SetScenarios(v map[string][]string)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


