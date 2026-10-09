# ScenarioPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Layers** | [**[]Layer**](Layer.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewScenarioPatch

`func NewScenarioPatch(layers []Layer, ) *ScenarioPatch`

NewScenarioPatch instantiates a new ScenarioPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioPatchWithDefaults

`func NewScenarioPatchWithDefaults() *ScenarioPatch`

NewScenarioPatchWithDefaults instantiates a new ScenarioPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScenarioPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScenarioPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScenarioPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScenarioPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLayers

`func (o *ScenarioPatch) GetLayers() []Layer`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *ScenarioPatch) GetLayersOk() (*[]Layer, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *ScenarioPatch) SetLayers(v []Layer)`

SetLayers sets Layers field to given value.


### GetName

`func (o *ScenarioPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScenarioPatch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


