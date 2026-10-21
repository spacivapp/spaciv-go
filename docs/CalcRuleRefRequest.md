# CalcRuleRefRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | [**MaybeVarProp**](MaybeVarProp.md) |  | 
**BinaryOperator** | [**CalcRuleFloatRequest**](CalcRuleFloatRequest.md) |  | 
**Condition** | [**CalcConditionRequest**](CalcConditionRequest.md) |  | 
**Num** | **float64** |  | 
**Select** | [**CalcSelectRequest**](CalcSelectRequest.md) |  | 
**SumOf** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 
**Var** | **string** |  | 

## Methods

### NewCalcRuleRefRequest

`func NewCalcRuleRefRequest(attr MaybeVarProp, binaryOperator CalcRuleFloatRequest, condition CalcConditionRequest, num float64, select_ CalcSelectRequest, sumOf CalcRuleRefRequest, var_ string, ) *CalcRuleRefRequest`

NewCalcRuleRefRequest instantiates a new CalcRuleRefRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcRuleRefRequestWithDefaults

`func NewCalcRuleRefRequestWithDefaults() *CalcRuleRefRequest`

NewCalcRuleRefRequestWithDefaults instantiates a new CalcRuleRefRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *CalcRuleRefRequest) GetAttr() MaybeVarProp`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *CalcRuleRefRequest) GetAttrOk() (*MaybeVarProp, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *CalcRuleRefRequest) SetAttr(v MaybeVarProp)`

SetAttr sets Attr field to given value.


### GetBinaryOperator

`func (o *CalcRuleRefRequest) GetBinaryOperator() CalcRuleFloatRequest`

GetBinaryOperator returns the BinaryOperator field if non-nil, zero value otherwise.

### GetBinaryOperatorOk

`func (o *CalcRuleRefRequest) GetBinaryOperatorOk() (*CalcRuleFloatRequest, bool)`

GetBinaryOperatorOk returns a tuple with the BinaryOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryOperator

`func (o *CalcRuleRefRequest) SetBinaryOperator(v CalcRuleFloatRequest)`

SetBinaryOperator sets BinaryOperator field to given value.


### GetCondition

`func (o *CalcRuleRefRequest) GetCondition() CalcConditionRequest`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CalcRuleRefRequest) GetConditionOk() (*CalcConditionRequest, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CalcRuleRefRequest) SetCondition(v CalcConditionRequest)`

SetCondition sets Condition field to given value.


### GetNum

`func (o *CalcRuleRefRequest) GetNum() float64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *CalcRuleRefRequest) GetNumOk() (*float64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *CalcRuleRefRequest) SetNum(v float64)`

SetNum sets Num field to given value.


### GetSelect

`func (o *CalcRuleRefRequest) GetSelect() CalcSelectRequest`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *CalcRuleRefRequest) GetSelectOk() (*CalcSelectRequest, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *CalcRuleRefRequest) SetSelect(v CalcSelectRequest)`

SetSelect sets Select field to given value.


### GetSumOf

`func (o *CalcRuleRefRequest) GetSumOf() CalcRuleRefRequest`

GetSumOf returns the SumOf field if non-nil, zero value otherwise.

### GetSumOfOk

`func (o *CalcRuleRefRequest) GetSumOfOk() (*CalcRuleRefRequest, bool)`

GetSumOfOk returns a tuple with the SumOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOf

`func (o *CalcRuleRefRequest) SetSumOf(v CalcRuleRefRequest)`

SetSumOf sets SumOf field to given value.


### GetVar

`func (o *CalcRuleRefRequest) GetVar() string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *CalcRuleRefRequest) GetVarOk() (*string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *CalcRuleRefRequest) SetVar(v string)`

SetVar sets Var field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


