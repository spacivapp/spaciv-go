# CalcRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axes** | [**[]CalcAxis**](CalcAxis.md) |  | 
**ColumnDecimalPlaces** | **map[string]int32** |  | 
**ColumnHexCodes** | **[]string** |  | 
**Truncated** | **bool** |  | 
**Truncations** | **[]int32** |  | 
**Values** | **[]float64** |  | 

## Methods

### NewCalcRes

`func NewCalcRes(axes []CalcAxis, columnDecimalPlaces map[string]int32, columnHexCodes []string, truncated bool, truncations []int32, values []float64, ) *CalcRes`

NewCalcRes instantiates a new CalcRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcResWithDefaults

`func NewCalcResWithDefaults() *CalcRes`

NewCalcResWithDefaults instantiates a new CalcRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxes

`func (o *CalcRes) GetAxes() []CalcAxis`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *CalcRes) GetAxesOk() (*[]CalcAxis, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxes

`func (o *CalcRes) SetAxes(v []CalcAxis)`

SetAxes sets Axes field to given value.


### GetColumnDecimalPlaces

`func (o *CalcRes) GetColumnDecimalPlaces() map[string]int32`

GetColumnDecimalPlaces returns the ColumnDecimalPlaces field if non-nil, zero value otherwise.

### GetColumnDecimalPlacesOk

`func (o *CalcRes) GetColumnDecimalPlacesOk() (*map[string]int32, bool)`

GetColumnDecimalPlacesOk returns a tuple with the ColumnDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDecimalPlaces

`func (o *CalcRes) SetColumnDecimalPlaces(v map[string]int32)`

SetColumnDecimalPlaces sets ColumnDecimalPlaces field to given value.


### GetColumnHexCodes

`func (o *CalcRes) GetColumnHexCodes() []string`

GetColumnHexCodes returns the ColumnHexCodes field if non-nil, zero value otherwise.

### GetColumnHexCodesOk

`func (o *CalcRes) GetColumnHexCodesOk() (*[]string, bool)`

GetColumnHexCodesOk returns a tuple with the ColumnHexCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnHexCodes

`func (o *CalcRes) SetColumnHexCodes(v []string)`

SetColumnHexCodes sets ColumnHexCodes field to given value.


### GetTruncated

`func (o *CalcRes) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *CalcRes) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *CalcRes) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetTruncations

`func (o *CalcRes) GetTruncations() []int32`

GetTruncations returns the Truncations field if non-nil, zero value otherwise.

### GetTruncationsOk

`func (o *CalcRes) GetTruncationsOk() (*[]int32, bool)`

GetTruncationsOk returns a tuple with the Truncations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncations

`func (o *CalcRes) SetTruncations(v []int32)`

SetTruncations sets Truncations field to given value.


### GetValues

`func (o *CalcRes) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CalcRes) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CalcRes) SetValues(v []float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


