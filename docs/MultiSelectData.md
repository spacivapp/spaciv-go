# MultiSelectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOthers** | **bool** |  | 
**MaxSelections** | **int32** |  | 
**MinSelections** | **int32** |  | 
**Options** | [**[]ChangelogOption**](ChangelogOption.md) |  | 

## Methods

### NewMultiSelectData

`func NewMultiSelectData(allowOthers bool, maxSelections int32, minSelections int32, options []ChangelogOption, ) *MultiSelectData`

NewMultiSelectData instantiates a new MultiSelectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSelectDataWithDefaults

`func NewMultiSelectDataWithDefaults() *MultiSelectData`

NewMultiSelectDataWithDefaults instantiates a new MultiSelectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOthers

`func (o *MultiSelectData) GetAllowOthers() bool`

GetAllowOthers returns the AllowOthers field if non-nil, zero value otherwise.

### GetAllowOthersOk

`func (o *MultiSelectData) GetAllowOthersOk() (*bool, bool)`

GetAllowOthersOk returns a tuple with the AllowOthers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOthers

`func (o *MultiSelectData) SetAllowOthers(v bool)`

SetAllowOthers sets AllowOthers field to given value.


### GetMaxSelections

`func (o *MultiSelectData) GetMaxSelections() int32`

GetMaxSelections returns the MaxSelections field if non-nil, zero value otherwise.

### GetMaxSelectionsOk

`func (o *MultiSelectData) GetMaxSelectionsOk() (*int32, bool)`

GetMaxSelectionsOk returns a tuple with the MaxSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSelections

`func (o *MultiSelectData) SetMaxSelections(v int32)`

SetMaxSelections sets MaxSelections field to given value.


### GetMinSelections

`func (o *MultiSelectData) GetMinSelections() int32`

GetMinSelections returns the MinSelections field if non-nil, zero value otherwise.

### GetMinSelectionsOk

`func (o *MultiSelectData) GetMinSelectionsOk() (*int32, bool)`

GetMinSelectionsOk returns a tuple with the MinSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSelections

`func (o *MultiSelectData) SetMinSelections(v int32)`

SetMinSelections sets MinSelections field to given value.


### GetOptions

`func (o *MultiSelectData) GetOptions() []ChangelogOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MultiSelectData) GetOptionsOk() (*[]ChangelogOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MultiSelectData) SetOptions(v []ChangelogOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


