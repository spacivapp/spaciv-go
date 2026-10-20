# CalcCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**CalcRuleRef**](CalcRuleRef.md) |  | 
**If** | [**CalcTest**](CalcTest.md) |  | 
**Then** | [**CalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewCalcCondition

`func NewCalcCondition(else_ CalcRuleRef, if_ CalcTest, then CalcRuleRef, ) *CalcCondition`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


