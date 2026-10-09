# VariablePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**VariableData**](VariableData.md) |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewVariablePatch

`func NewVariablePatch() *VariablePatch`

NewVariablePatch instantiates a new VariablePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariablePatchWithDefaults

`func NewVariablePatchWithDefaults() *VariablePatch`

NewVariablePatchWithDefaults instantiates a new VariablePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VariablePatch) GetData() VariableData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VariablePatch) GetDataOk() (*VariableData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VariablePatch) SetData(v VariableData)`

SetData sets Data field to given value.

### HasData

`func (o *VariablePatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDelete

`func (o *VariablePatch) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *VariablePatch) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *VariablePatch) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *VariablePatch) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetName

`func (o *VariablePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariablePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariablePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariablePatch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


