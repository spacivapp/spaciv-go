# FloorplanGetBodySpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** |  | 
**Paths** | [**[][]FloorplanGetBodySpacePoint**]([]FloorplanGetBodySpacePoint.md) |  | 

## Methods

### NewFloorplanGetBodySpace

`func NewFloorplanGetBodySpace(location string, paths [][]FloorplanGetBodySpacePoint, ) *FloorplanGetBodySpace`

NewFloorplanGetBodySpace instantiates a new FloorplanGetBodySpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorplanGetBodySpaceWithDefaults

`func NewFloorplanGetBodySpaceWithDefaults() *FloorplanGetBodySpace`

NewFloorplanGetBodySpaceWithDefaults instantiates a new FloorplanGetBodySpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FloorplanGetBodySpace) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FloorplanGetBodySpace) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FloorplanGetBodySpace) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetPaths

`func (o *FloorplanGetBodySpace) GetPaths() [][]FloorplanGetBodySpacePoint`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *FloorplanGetBodySpace) GetPathsOk() (*[][]FloorplanGetBodySpacePoint, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *FloorplanGetBodySpace) SetPaths(v [][]FloorplanGetBodySpacePoint)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


