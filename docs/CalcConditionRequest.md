# CalcConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 
**If** | [**CalcTestRequest**](CalcTestRequest.md) |  | 
**Then** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 

## Methods

### NewCalcConditionRequest

`func NewCalcConditionRequest(else_ CalcRuleRefRequest, if_ CalcTestRequest, then CalcRuleRefRequest, ) *CalcConditionRequest`

NewCalcConditionRequest instantiates a new CalcConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcConditionRequestWithDefaults

`func NewCalcConditionRequestWithDefaults() *CalcConditionRequest`

NewCalcConditionRequestWithDefaults instantiates a new CalcConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElse

`func (o *CalcConditionRequest) GetElse() CalcRuleRefRequest`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *CalcConditionRequest) GetElseOk() (*CalcRuleRefRequest, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *CalcConditionRequest) SetElse(v CalcRuleRefRequest)`

SetElse sets Else field to given value.


### GetIf

`func (o *CalcConditionRequest) GetIf() CalcTestRequest`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CalcConditionRequest) GetIfOk() (*CalcTestRequest, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CalcConditionRequest) SetIf(v CalcTestRequest)`

SetIf sets If field to given value.


### GetThen

`func (o *CalcConditionRequest) GetThen() CalcRuleRefRequest`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *CalcConditionRequest) GetThenOk() (*CalcRuleRefRequest, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *CalcConditionRequest) SetThen(v CalcRuleRefRequest)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


