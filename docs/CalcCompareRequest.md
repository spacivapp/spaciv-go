# CalcCompareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 
**Op** | **int32** |  | 
**Right** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 

## Methods

### NewCalcCompareRequest

`func NewCalcCompareRequest(left CalcRuleRefRequest, op int32, right CalcRuleRefRequest, ) *CalcCompareRequest`

NewCalcCompareRequest instantiates a new CalcCompareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcCompareRequestWithDefaults

`func NewCalcCompareRequestWithDefaults() *CalcCompareRequest`

NewCalcCompareRequestWithDefaults instantiates a new CalcCompareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *CalcCompareRequest) GetLeft() CalcRuleRefRequest`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *CalcCompareRequest) GetLeftOk() (*CalcRuleRefRequest, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *CalcCompareRequest) SetLeft(v CalcRuleRefRequest)`

SetLeft sets Left field to given value.


### GetOp

`func (o *CalcCompareRequest) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CalcCompareRequest) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CalcCompareRequest) SetOp(v int32)`

SetOp sets Op field to given value.


### GetRight

`func (o *CalcCompareRequest) GetRight() CalcRuleRefRequest`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *CalcCompareRequest) GetRightOk() (*CalcRuleRefRequest, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *CalcCompareRequest) SetRight(v CalcRuleRefRequest)`

SetRight sets Right field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


