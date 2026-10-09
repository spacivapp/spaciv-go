# ProjectCreateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PropertyMap** | [**map[string]MaybeVarProp**](MaybeVarProp.md) |  | 
**ScenarioMap** | [**map[string]MaybeVarStr**](MaybeVarStr.md) |  | 
**Variables** | [**[]VariableRes**](VariableRes.md) |  | 

## Methods

### NewProjectCreateTemplate

`func NewProjectCreateTemplate(id int32, propertyMap map[string]MaybeVarProp, scenarioMap map[string]MaybeVarStr, variables []VariableRes, ) *ProjectCreateTemplate`

NewProjectCreateTemplate instantiates a new ProjectCreateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateTemplateWithDefaults

`func NewProjectCreateTemplateWithDefaults() *ProjectCreateTemplate`

NewProjectCreateTemplateWithDefaults instantiates a new ProjectCreateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCreateTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCreateTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCreateTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetPropertyMap

`func (o *ProjectCreateTemplate) GetPropertyMap() map[string]MaybeVarProp`

GetPropertyMap returns the PropertyMap field if non-nil, zero value otherwise.

### GetPropertyMapOk

`func (o *ProjectCreateTemplate) GetPropertyMapOk() (*map[string]MaybeVarProp, bool)`

GetPropertyMapOk returns a tuple with the PropertyMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMap

`func (o *ProjectCreateTemplate) SetPropertyMap(v map[string]MaybeVarProp)`

SetPropertyMap sets PropertyMap field to given value.


### GetScenarioMap

`func (o *ProjectCreateTemplate) GetScenarioMap() map[string]MaybeVarStr`

GetScenarioMap returns the ScenarioMap field if non-nil, zero value otherwise.

### GetScenarioMapOk

`func (o *ProjectCreateTemplate) GetScenarioMapOk() (*map[string]MaybeVarStr, bool)`

GetScenarioMapOk returns a tuple with the ScenarioMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioMap

`func (o *ProjectCreateTemplate) SetScenarioMap(v map[string]MaybeVarStr)`

SetScenarioMap sets ScenarioMap field to given value.


### GetVariables

`func (o *ProjectCreateTemplate) GetVariables() []VariableRes`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProjectCreateTemplate) GetVariablesOk() (*[]VariableRes, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProjectCreateTemplate) SetVariables(v []VariableRes)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


