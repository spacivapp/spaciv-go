# SelectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOther** | **bool** |  | 
**Display** | **int32** |  | 
**MaxSelections** | **int32** |  | 
**MinSelections** | **int32** |  | 
**Multiple** | **bool** |  | 
**OtherLabel** | **string** |  | 
**OtherOption** | **string** |  | 
**Rows** | [**[]AnswerOption**](AnswerOption.md) |  | 
**Wrap** | **bool** |  | 

## Methods

### NewSelectType

`func NewSelectType(allowOther bool, display int32, maxSelections int32, minSelections int32, multiple bool, otherLabel string, otherOption string, rows []AnswerOption, wrap bool, ) *SelectType`

NewSelectType instantiates a new SelectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectTypeWithDefaults

`func NewSelectTypeWithDefaults() *SelectType`

NewSelectTypeWithDefaults instantiates a new SelectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOther

`func (o *SelectType) GetAllowOther() bool`

GetAllowOther returns the AllowOther field if non-nil, zero value otherwise.

### GetAllowOtherOk

`func (o *SelectType) GetAllowOtherOk() (*bool, bool)`

GetAllowOtherOk returns a tuple with the AllowOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOther

`func (o *SelectType) SetAllowOther(v bool)`

SetAllowOther sets AllowOther field to given value.


### GetDisplay

`func (o *SelectType) GetDisplay() int32`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SelectType) GetDisplayOk() (*int32, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SelectType) SetDisplay(v int32)`

SetDisplay sets Display field to given value.


### GetMaxSelections

`func (o *SelectType) GetMaxSelections() int32`

GetMaxSelections returns the MaxSelections field if non-nil, zero value otherwise.

### GetMaxSelectionsOk

`func (o *SelectType) GetMaxSelectionsOk() (*int32, bool)`

GetMaxSelectionsOk returns a tuple with the MaxSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSelections

`func (o *SelectType) SetMaxSelections(v int32)`

SetMaxSelections sets MaxSelections field to given value.


### GetMinSelections

`func (o *SelectType) GetMinSelections() int32`

GetMinSelections returns the MinSelections field if non-nil, zero value otherwise.

### GetMinSelectionsOk

`func (o *SelectType) GetMinSelectionsOk() (*int32, bool)`

GetMinSelectionsOk returns a tuple with the MinSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSelections

`func (o *SelectType) SetMinSelections(v int32)`

SetMinSelections sets MinSelections field to given value.


### GetMultiple

`func (o *SelectType) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *SelectType) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *SelectType) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.


### GetOtherLabel

`func (o *SelectType) GetOtherLabel() string`

GetOtherLabel returns the OtherLabel field if non-nil, zero value otherwise.

### GetOtherLabelOk

`func (o *SelectType) GetOtherLabelOk() (*string, bool)`

GetOtherLabelOk returns a tuple with the OtherLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLabel

`func (o *SelectType) SetOtherLabel(v string)`

SetOtherLabel sets OtherLabel field to given value.


### GetOtherOption

`func (o *SelectType) GetOtherOption() string`

GetOtherOption returns the OtherOption field if non-nil, zero value otherwise.

### GetOtherOptionOk

`func (o *SelectType) GetOtherOptionOk() (*string, bool)`

GetOtherOptionOk returns a tuple with the OtherOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOption

`func (o *SelectType) SetOtherOption(v string)`

SetOtherOption sets OtherOption field to given value.


### GetRows

`func (o *SelectType) GetRows() []AnswerOption`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SelectType) GetRowsOk() (*[]AnswerOption, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SelectType) SetRows(v []AnswerOption)`

SetRows sets Rows field to given value.


### GetWrap

`func (o *SelectType) GetWrap() bool`

GetWrap returns the Wrap field if non-nil, zero value otherwise.

### GetWrapOk

`func (o *SelectType) GetWrapOk() (*bool, bool)`

GetWrapOk returns a tuple with the Wrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrap

`func (o *SelectType) SetWrap(v bool)`

SetWrap sets Wrap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


