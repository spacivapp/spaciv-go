# GraphJoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | [**NullableGraphJoinKey**](GraphJoinKey.md) |  | 
**Right** | [**NullableGraphJoinKey**](GraphJoinKey.md) |  | 

## Methods

### NewGraphJoin

`func NewGraphJoin(left NullableGraphJoinKey, right NullableGraphJoinKey, ) *GraphJoin`

NewGraphJoin instantiates a new GraphJoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphJoinWithDefaults

`func NewGraphJoinWithDefaults() *GraphJoin`

NewGraphJoinWithDefaults instantiates a new GraphJoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *GraphJoin) GetLeft() GraphJoinKey`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *GraphJoin) GetLeftOk() (*GraphJoinKey, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *GraphJoin) SetLeft(v GraphJoinKey)`

SetLeft sets Left field to given value.


### SetLeftNil

`func (o *GraphJoin) SetLeftNil(b bool)`

 SetLeftNil sets the value for Left to be an explicit nil

### UnsetLeft
`func (o *GraphJoin) UnsetLeft()`

UnsetLeft ensures that no value is present for Left, not even an explicit nil
### GetRight

`func (o *GraphJoin) GetRight() GraphJoinKey`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *GraphJoin) GetRightOk() (*GraphJoinKey, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *GraphJoin) SetRight(v GraphJoinKey)`

SetRight sets Right field to given value.


### SetRightNil

`func (o *GraphJoin) SetRightNil(b bool)`

 SetRightNil sets the value for Right to be an explicit nil

### UnsetRight
`func (o *GraphJoin) UnsetRight()`

UnsetRight ensures that no value is present for Right, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


