# CalcAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | **int32** | &#x60;1&#x60; is day, &#x60;2&#x60; tree node, &#x60;3&#x60; series, &#x60;4&#x60; scenario and &#x60;5&#x60; group. | 
**Members** | [**[]CalcAxisMember**](CalcAxisMember.md) |  | 

## Methods

### NewCalcAxis

`func NewCalcAxis(dimension int32, members []CalcAxisMember, ) *CalcAxis`

NewCalcAxis instantiates a new CalcAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcAxisWithDefaults

`func NewCalcAxisWithDefaults() *CalcAxis`

NewCalcAxisWithDefaults instantiates a new CalcAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *CalcAxis) GetDimension() int32`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *CalcAxis) GetDimensionOk() (*int32, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *CalcAxis) SetDimension(v int32)`

SetDimension sets Dimension field to given value.


### GetMembers

`func (o *CalcAxis) GetMembers() []CalcAxisMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CalcAxis) GetMembersOk() (*[]CalcAxisMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CalcAxis) SetMembers(v []CalcAxisMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


