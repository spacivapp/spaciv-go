# CalcListTwoDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axis1** | [**CalcSegmentationRequest**](CalcSegmentationRequest.md) |  | 
**Axis2** | [**CalcSegmentationRequest**](CalcSegmentationRequest.md) |  | 
**Date** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 

## Methods

### NewCalcListTwoDRequest

`func NewCalcListTwoDRequest(axis1 CalcSegmentationRequest, axis2 CalcSegmentationRequest, date MaybeVarInt, filters []FilterGroup, ) *CalcListTwoDRequest`

NewCalcListTwoDRequest instantiates a new CalcListTwoDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListTwoDRequestWithDefaults

`func NewCalcListTwoDRequestWithDefaults() *CalcListTwoDRequest`

NewCalcListTwoDRequestWithDefaults instantiates a new CalcListTwoDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis1

`func (o *CalcListTwoDRequest) GetAxis1() CalcSegmentationRequest`

GetAxis1 returns the Axis1 field if non-nil, zero value otherwise.

### GetAxis1Ok

`func (o *CalcListTwoDRequest) GetAxis1Ok() (*CalcSegmentationRequest, bool)`

GetAxis1Ok returns a tuple with the Axis1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis1

`func (o *CalcListTwoDRequest) SetAxis1(v CalcSegmentationRequest)`

SetAxis1 sets Axis1 field to given value.


### GetAxis2

`func (o *CalcListTwoDRequest) GetAxis2() CalcSegmentationRequest`

GetAxis2 returns the Axis2 field if non-nil, zero value otherwise.

### GetAxis2Ok

`func (o *CalcListTwoDRequest) GetAxis2Ok() (*CalcSegmentationRequest, bool)`

GetAxis2Ok returns a tuple with the Axis2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis2

`func (o *CalcListTwoDRequest) SetAxis2(v CalcSegmentationRequest)`

SetAxis2 sets Axis2 field to given value.


### GetDate

`func (o *CalcListTwoDRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcListTwoDRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcListTwoDRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### GetFilters

`func (o *CalcListTwoDRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcListTwoDRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcListTwoDRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


