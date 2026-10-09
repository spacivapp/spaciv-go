# EstimationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Contribution**](Contribution.md) |  | [optional] 
**Delete** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SetDefault** | **bool** |  | 

## Methods

### NewEstimationPatch

`func NewEstimationPatch(delete bool, setDefault bool, ) *EstimationPatch`

NewEstimationPatch instantiates a new EstimationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimationPatchWithDefaults

`func NewEstimationPatchWithDefaults() *EstimationPatch`

NewEstimationPatchWithDefaults instantiates a new EstimationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EstimationPatch) GetData() Contribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EstimationPatch) GetDataOk() (*Contribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EstimationPatch) SetData(v Contribution)`

SetData sets Data field to given value.

### HasData

`func (o *EstimationPatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDelete

`func (o *EstimationPatch) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *EstimationPatch) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *EstimationPatch) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetDescription

`func (o *EstimationPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EstimationPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EstimationPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EstimationPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *EstimationPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EstimationPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EstimationPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EstimationPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSetDefault

`func (o *EstimationPatch) GetSetDefault() bool`

GetSetDefault returns the SetDefault field if non-nil, zero value otherwise.

### GetSetDefaultOk

`func (o *EstimationPatch) GetSetDefaultOk() (*bool, bool)`

GetSetDefaultOk returns a tuple with the SetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDefault

`func (o *EstimationPatch) SetSetDefault(v bool)`

SetSetDefault sets SetDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


