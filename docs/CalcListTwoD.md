# CalcListTwoD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axis1** | [**CalcSegmentation**](CalcSegmentation.md) |  | 
**Axis2** | [**CalcSegmentation**](CalcSegmentation.md) |  | 
**Date** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 

## Methods

### NewCalcListTwoD

`func NewCalcListTwoD(axis1 CalcSegmentation, axis2 CalcSegmentation, date MaybeVarInt, filters []FilterGroup, ) *CalcListTwoD`

NewCalcListTwoD instantiates a new CalcListTwoD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListTwoDWithDefaults

`func NewCalcListTwoDWithDefaults() *CalcListTwoD`

NewCalcListTwoDWithDefaults instantiates a new CalcListTwoD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis1

`func (o *CalcListTwoD) GetAxis1() CalcSegmentation`

GetAxis1 returns the Axis1 field if non-nil, zero value otherwise.

### GetAxis1Ok

`func (o *CalcListTwoD) GetAxis1Ok() (*CalcSegmentation, bool)`

GetAxis1Ok returns a tuple with the Axis1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis1

`func (o *CalcListTwoD) SetAxis1(v CalcSegmentation)`

SetAxis1 sets Axis1 field to given value.


### GetAxis2

`func (o *CalcListTwoD) GetAxis2() CalcSegmentation`

GetAxis2 returns the Axis2 field if non-nil, zero value otherwise.

### GetAxis2Ok

`func (o *CalcListTwoD) GetAxis2Ok() (*CalcSegmentation, bool)`

GetAxis2Ok returns a tuple with the Axis2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis2

`func (o *CalcListTwoD) SetAxis2(v CalcSegmentation)`

SetAxis2 sets Axis2 field to given value.


### GetDate

`func (o *CalcListTwoD) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcListTwoD) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcListTwoD) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### GetFilters

`func (o *CalcListTwoD) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcListTwoD) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcListTwoD) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


