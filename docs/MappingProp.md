# MappingProp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**MaybeVarProp**](MaybeVarProp.md) |  | 
**Name** | **string** |  | 
**Resolved** | **bool** |  | 

## Methods

### NewMappingProp

`func NewMappingProp(data MaybeVarProp, name string, resolved bool, ) *MappingProp`

NewMappingProp instantiates a new MappingProp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingPropWithDefaults

`func NewMappingPropWithDefaults() *MappingProp`

NewMappingPropWithDefaults instantiates a new MappingProp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MappingProp) GetData() MaybeVarProp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MappingProp) GetDataOk() (*MaybeVarProp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MappingProp) SetData(v MaybeVarProp)`

SetData sets Data field to given value.


### GetName

`func (o *MappingProp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappingProp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappingProp) SetName(v string)`

SetName sets Name field to given value.


### GetResolved

`func (o *MappingProp) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MappingProp) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MappingProp) SetResolved(v bool)`

SetResolved sets Resolved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


