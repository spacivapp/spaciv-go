# CalcCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 
**If** | [**NullableCalcTest**](CalcTest.md) |  | 
**Then** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewCalcCondition

`func NewCalcCondition(else_ NullableCalcRuleRef, if_ NullableCalcTest, then NullableCalcRuleRef, ) *CalcCondition`

NewCalcCondition instantiates a new CalcCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcConditionWithDefaults

`func NewCalcConditionWithDefaults() *CalcCondition`

NewCalcConditionWithDefaults instantiates a new CalcCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElse

`func (o *CalcCondition) GetElse() CalcRuleRef`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *CalcCondition) GetElseOk() (*CalcRuleRef, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *CalcCondition) SetElse(v CalcRuleRef)`

SetElse sets Else field to given value.


### SetElseNil

`func (o *CalcCondition) SetElseNil(b bool)`

 SetElseNil sets the value for Else to be an explicit nil

### UnsetElse
`func (o *CalcCondition) UnsetElse()`

UnsetElse ensures that no value is present for Else, not even an explicit nil
### GetIf

`func (o *CalcCondition) GetIf() CalcTest`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CalcCondition) GetIfOk() (*CalcTest, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CalcCondition) SetIf(v CalcTest)`

SetIf sets If field to given value.


### SetIfNil

`func (o *CalcCondition) SetIfNil(b bool)`

 SetIfNil sets the value for If to be an explicit nil

### UnsetIf
`func (o *CalcCondition) UnsetIf()`

UnsetIf ensures that no value is present for If, not even an explicit nil
### GetThen

`func (o *CalcCondition) GetThen() CalcRuleRef`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *CalcCondition) GetThenOk() (*CalcRuleRef, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *CalcCondition) SetThen(v CalcRuleRef)`

SetThen sets Then field to given value.


### SetThenNil

`func (o *CalcCondition) SetThenNil(b bool)`

 SetThenNil sets the value for Then to be an explicit nil

### UnsetThen
`func (o *CalcCondition) UnsetThen()`

UnsetThen ensures that no value is present for Then, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


