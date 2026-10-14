# CalcRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axes** | [**[]CalcAxis**](CalcAxis.md) |  | 
**Values** | **[]float64** |  | 

## Methods

### NewCalcRes

`func NewCalcRes(axes []CalcAxis, values []float64, ) *CalcRes`

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


