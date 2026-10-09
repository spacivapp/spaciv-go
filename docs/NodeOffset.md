# NodeOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **string** |  | 
**Root** | [**CalcNode**](CalcNode.md) |  | 

## Methods

### NewNodeOffset

`func NewNodeOffset(day string, root CalcNode, ) *NodeOffset`

NewNodeOffset instantiates a new NodeOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeOffsetWithDefaults

`func NewNodeOffsetWithDefaults() *NodeOffset`

NewNodeOffsetWithDefaults instantiates a new NodeOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *NodeOffset) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *NodeOffset) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *NodeOffset) SetDay(v string)`

SetDay sets Day field to given value.


### GetRoot

`func (o *NodeOffset) GetRoot() CalcNode`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *NodeOffset) GetRootOk() (*CalcNode, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *NodeOffset) SetRoot(v CalcNode)`

SetRoot sets Root field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


