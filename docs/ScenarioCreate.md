# ScenarioCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Layers** | [**[]Layer**](Layer.md) |  | 
**Name** | **string** |  | 

## Methods

### NewScenarioCreate

`func NewScenarioCreate(description string, layers []Layer, name string, ) *ScenarioCreate`

NewScenarioCreate instantiates a new ScenarioCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioCreateWithDefaults

`func NewScenarioCreateWithDefaults() *ScenarioCreate`

NewScenarioCreateWithDefaults instantiates a new ScenarioCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScenarioCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScenarioCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScenarioCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLayers

`func (o *ScenarioCreate) GetLayers() []Layer`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *ScenarioCreate) GetLayersOk() (*[]Layer, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *ScenarioCreate) SetLayers(v []Layer)`

SetLayers sets Layers field to given value.


### GetName

`func (o *ScenarioCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


