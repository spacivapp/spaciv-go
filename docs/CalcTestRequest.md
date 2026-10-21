# CalcTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compare** | [**CalcCompareRequest**](CalcCompareRequest.md) |  | 
**Has** | [**CalcHasNode**](CalcHasNode.md) |  | 
**IsObject** | [**CalcIsObject**](CalcIsObject.md) |  | 

## Methods

### NewCalcTestRequest

`func NewCalcTestRequest(compare CalcCompareRequest, has CalcHasNode, isObject CalcIsObject, ) *CalcTestRequest`

NewCalcTestRequest instantiates a new CalcTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcTestRequestWithDefaults

`func NewCalcTestRequestWithDefaults() *CalcTestRequest`

NewCalcTestRequestWithDefaults instantiates a new CalcTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompare

`func (o *CalcTestRequest) GetCompare() CalcCompareRequest`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *CalcTestRequest) GetCompareOk() (*CalcCompareRequest, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *CalcTestRequest) SetCompare(v CalcCompareRequest)`

SetCompare sets Compare field to given value.


### GetHas

`func (o *CalcTestRequest) GetHas() CalcHasNode`

GetHas returns the Has field if non-nil, zero value otherwise.

### GetHasOk

`func (o *CalcTestRequest) GetHasOk() (*CalcHasNode, bool)`

GetHasOk returns a tuple with the Has field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas

`func (o *CalcTestRequest) SetHas(v CalcHasNode)`

SetHas sets Has field to given value.


### GetIsObject

`func (o *CalcTestRequest) GetIsObject() CalcIsObject`

GetIsObject returns the IsObject field if non-nil, zero value otherwise.

### GetIsObjectOk

`func (o *CalcTestRequest) GetIsObjectOk() (*CalcIsObject, bool)`

GetIsObjectOk returns a tuple with the IsObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObject

`func (o *CalcTestRequest) SetIsObject(v CalcIsObject)`

SetIsObject sets IsObject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


