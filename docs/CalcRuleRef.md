# CalcRuleRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | [**MaybeVarProp**](MaybeVarProp.md) |  | 
**BinaryOperator** | [**CalcRuleFloat**](CalcRuleFloat.md) |  | 
**Num** | **float64** |  | 
**Select** | [**CalcSelect**](CalcSelect.md) |  | 
**SumOf** | [**CalcRuleRef**](CalcRuleRef.md) |  | 
**Var** | **string** |  | 

## Methods

### NewCalcRuleRef

`func NewCalcRuleRef(attr MaybeVarProp, binaryOperator CalcRuleFloat, num float64, select_ CalcSelect, sumOf CalcRuleRef, var_ string, ) *CalcRuleRef`

NewCalcRuleRef instantiates a new CalcRuleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcRuleRefWithDefaults

`func NewCalcRuleRefWithDefaults() *CalcRuleRef`

NewCalcRuleRefWithDefaults instantiates a new CalcRuleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *CalcRuleRef) GetAttr() MaybeVarProp`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *CalcRuleRef) GetAttrOk() (*MaybeVarProp, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *CalcRuleRef) SetAttr(v MaybeVarProp)`

SetAttr sets Attr field to given value.


### GetBinaryOperator

`func (o *CalcRuleRef) GetBinaryOperator() CalcRuleFloat`

GetBinaryOperator returns the BinaryOperator field if non-nil, zero value otherwise.

### GetBinaryOperatorOk

`func (o *CalcRuleRef) GetBinaryOperatorOk() (*CalcRuleFloat, bool)`

GetBinaryOperatorOk returns a tuple with the BinaryOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryOperator

`func (o *CalcRuleRef) SetBinaryOperator(v CalcRuleFloat)`

SetBinaryOperator sets BinaryOperator field to given value.


### GetNum

`func (o *CalcRuleRef) GetNum() float64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *CalcRuleRef) GetNumOk() (*float64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *CalcRuleRef) SetNum(v float64)`

SetNum sets Num field to given value.


### GetSelect

`func (o *CalcRuleRef) GetSelect() CalcSelect`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *CalcRuleRef) GetSelectOk() (*CalcSelect, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *CalcRuleRef) SetSelect(v CalcSelect)`

SetSelect sets Select field to given value.


### GetSumOf

`func (o *CalcRuleRef) GetSumOf() CalcRuleRef`

GetSumOf returns the SumOf field if non-nil, zero value otherwise.

### GetSumOfOk

`func (o *CalcRuleRef) GetSumOfOk() (*CalcRuleRef, bool)`

GetSumOfOk returns a tuple with the SumOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOf

`func (o *CalcRuleRef) SetSumOf(v CalcRuleRef)`

SetSumOf sets SumOf field to given value.


### GetVar

`func (o *CalcRuleRef) GetVar() string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *CalcRuleRef) GetVarOk() (*string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *CalcRuleRef) SetVar(v string)`

SetVar sets Var field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


