# SpaceRulesetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Preset** | Pointer to **string** | UUID of a preset space ruleset to copy as the starting configuration. When supplied, the type, creation mode, config, and space modules are all taken from the preset. | [optional] 
**Type** | **string** | Ruleset type: \&quot;Standard\&quot;, \&quot;Amenity\&quot;, or \&quot;Advanced\&quot;. Determines the initial calculation configuration template when no preset is supplied. | 

## Methods

### NewSpaceRulesetCreate

`func NewSpaceRulesetCreate(name string, type_ string, ) *SpaceRulesetCreate`

NewSpaceRulesetCreate instantiates a new SpaceRulesetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRulesetCreateWithDefaults

`func NewSpaceRulesetCreateWithDefaults() *SpaceRulesetCreate`

NewSpaceRulesetCreateWithDefaults instantiates a new SpaceRulesetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SpaceRulesetCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpaceRulesetCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpaceRulesetCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpaceRulesetCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SpaceRulesetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceRulesetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceRulesetCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPreset

`func (o *SpaceRulesetCreate) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *SpaceRulesetCreate) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *SpaceRulesetCreate) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *SpaceRulesetCreate) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetType

`func (o *SpaceRulesetCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceRulesetCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceRulesetCreate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


