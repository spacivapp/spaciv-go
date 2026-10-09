# MappingScenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**MaybeVarStr**](MaybeVarStr.md) |  | 
**Name** | **string** |  | 
**Resolved** | **bool** |  | 

## Methods

### NewMappingScenario

`func NewMappingScenario(data MaybeVarStr, name string, resolved bool, ) *MappingScenario`

NewMappingScenario instantiates a new MappingScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingScenarioWithDefaults

`func NewMappingScenarioWithDefaults() *MappingScenario`

NewMappingScenarioWithDefaults instantiates a new MappingScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MappingScenario) GetData() MaybeVarStr`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MappingScenario) GetDataOk() (*MaybeVarStr, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MappingScenario) SetData(v MaybeVarStr)`

SetData sets Data field to given value.


### GetName

`func (o *MappingScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappingScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappingScenario) SetName(v string)`

SetName sets Name field to given value.


### GetResolved

`func (o *MappingScenario) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MappingScenario) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MappingScenario) SetResolved(v bool)`

SetResolved sets Resolved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


