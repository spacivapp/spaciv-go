# CalcSelectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calc** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 
**Object** | [**MaybeVarProp**](MaybeVarProp.md) |  | 

## Methods

### NewCalcSelectRequest

`func NewCalcSelectRequest(calc CalcRuleRefRequest, object MaybeVarProp, ) *CalcSelectRequest`

NewCalcSelectRequest instantiates a new CalcSelectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcSelectRequestWithDefaults

`func NewCalcSelectRequestWithDefaults() *CalcSelectRequest`

NewCalcSelectRequestWithDefaults instantiates a new CalcSelectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalc

`func (o *CalcSelectRequest) GetCalc() CalcRuleRefRequest`

GetCalc returns the Calc field if non-nil, zero value otherwise.

### GetCalcOk

`func (o *CalcSelectRequest) GetCalcOk() (*CalcRuleRefRequest, bool)`

GetCalcOk returns a tuple with the Calc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalc

`func (o *CalcSelectRequest) SetCalc(v CalcRuleRefRequest)`

SetCalc sets Calc field to given value.


### GetObject

`func (o *CalcSelectRequest) GetObject() MaybeVarProp`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CalcSelectRequest) GetObjectOk() (*MaybeVarProp, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CalcSelectRequest) SetObject(v MaybeVarProp)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


