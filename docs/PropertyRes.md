# PropertyRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tree** | [**NullablePropertyNode**](PropertyNode.md) |  | 
**Value** | [**NullablePropertyChangelogBody**](PropertyChangelogBody.md) |  | 

## Methods

### NewPropertyRes

`func NewPropertyRes(tree NullablePropertyNode, value NullablePropertyChangelogBody, ) *PropertyRes`

NewPropertyRes instantiates a new PropertyRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyResWithDefaults

`func NewPropertyResWithDefaults() *PropertyRes`

NewPropertyResWithDefaults instantiates a new PropertyRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTree

`func (o *PropertyRes) GetTree() PropertyNode`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *PropertyRes) GetTreeOk() (*PropertyNode, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *PropertyRes) SetTree(v PropertyNode)`

SetTree sets Tree field to given value.


### SetTreeNil

`func (o *PropertyRes) SetTreeNil(b bool)`

 SetTreeNil sets the value for Tree to be an explicit nil

### UnsetTree
`func (o *PropertyRes) UnsetTree()`

UnsetTree ensures that no value is present for Tree, not even an explicit nil
### GetValue

`func (o *PropertyRes) GetValue() PropertyChangelogBody`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyRes) GetValueOk() (*PropertyChangelogBody, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyRes) SetValue(v PropertyChangelogBody)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *PropertyRes) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PropertyRes) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


