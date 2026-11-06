# GraphConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**NullableGraphSocketRef**](GraphSocketRef.md) |  | 
**To** | [**NullableGraphSocketRef**](GraphSocketRef.md) |  | 

## Methods

### NewGraphConnection

`func NewGraphConnection(from NullableGraphSocketRef, to NullableGraphSocketRef, ) *GraphConnection`

NewGraphConnection instantiates a new GraphConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphConnectionWithDefaults

`func NewGraphConnectionWithDefaults() *GraphConnection`

NewGraphConnectionWithDefaults instantiates a new GraphConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *GraphConnection) GetFrom() GraphSocketRef`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphConnection) GetFromOk() (*GraphSocketRef, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphConnection) SetFrom(v GraphSocketRef)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *GraphConnection) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *GraphConnection) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetTo

`func (o *GraphConnection) GetTo() GraphSocketRef`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphConnection) GetToOk() (*GraphSocketRef, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphConnection) SetTo(v GraphSocketRef)`

SetTo sets To field to given value.


### SetToNil

`func (o *GraphConnection) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GraphConnection) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


