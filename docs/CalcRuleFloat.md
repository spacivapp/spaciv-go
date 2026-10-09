# CalcRuleFloat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | [**CalcRuleRef**](CalcRuleRef.md) |  | 
**Op** | **int32** |  | 
**Right** | [**CalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewCalcRuleFloat

`func NewCalcRuleFloat(left CalcRuleRef, op int32, right CalcRuleRef, ) *CalcRuleFloat`

NewCalcRuleFloat instantiates a new CalcRuleFloat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcRuleFloatWithDefaults

`func NewCalcRuleFloatWithDefaults() *CalcRuleFloat`

NewCalcRuleFloatWithDefaults instantiates a new CalcRuleFloat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *CalcRuleFloat) GetLeft() CalcRuleRef`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *CalcRuleFloat) GetLeftOk() (*CalcRuleRef, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *CalcRuleFloat) SetLeft(v CalcRuleRef)`

SetLeft sets Left field to given value.


### GetOp

`func (o *CalcRuleFloat) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CalcRuleFloat) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CalcRuleFloat) SetOp(v int32)`

SetOp sets Op field to given value.


### GetRight

`func (o *CalcRuleFloat) GetRight() CalcRuleRef`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *CalcRuleFloat) GetRightOk() (*CalcRuleRef, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *CalcRuleFloat) SetRight(v CalcRuleRef)`

SetRight sets Right field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


