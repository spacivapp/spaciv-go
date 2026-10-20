# CalcCompare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | [**CalcRuleRef**](CalcRuleRef.md) |  | 
**Op** | **int32** |  | 
**Right** | [**CalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewCalcCompare

`func NewCalcCompare(left CalcRuleRef, op int32, right CalcRuleRef, ) *CalcCompare`

NewCalcCompare instantiates a new CalcCompare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcCompareWithDefaults

`func NewCalcCompareWithDefaults() *CalcCompare`

NewCalcCompareWithDefaults instantiates a new CalcCompare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *CalcCompare) GetLeft() CalcRuleRef`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *CalcCompare) GetLeftOk() (*CalcRuleRef, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *CalcCompare) SetLeft(v CalcRuleRef)`

SetLeft sets Left field to given value.


### GetOp

`func (o *CalcCompare) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CalcCompare) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CalcCompare) SetOp(v int32)`

SetOp sets Op field to given value.


### GetRight

`func (o *CalcCompare) GetRight() CalcRuleRef`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *CalcCompare) GetRightOk() (*CalcRuleRef, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *CalcCompare) SetRight(v CalcRuleRef)`

SetRight sets Right field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


