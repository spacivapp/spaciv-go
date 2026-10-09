# FloorplanCreateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** |  | 
**Paths** | [**[]FloorplanSpaceChangelogPath**](FloorplanSpaceChangelogPath.md) |  | 

## Methods

### NewFloorplanCreateSpace

`func NewFloorplanCreateSpace(location string, paths []FloorplanSpaceChangelogPath, ) *FloorplanCreateSpace`

NewFloorplanCreateSpace instantiates a new FloorplanCreateSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorplanCreateSpaceWithDefaults

`func NewFloorplanCreateSpaceWithDefaults() *FloorplanCreateSpace`

NewFloorplanCreateSpaceWithDefaults instantiates a new FloorplanCreateSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FloorplanCreateSpace) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FloorplanCreateSpace) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FloorplanCreateSpace) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetPaths

`func (o *FloorplanCreateSpace) GetPaths() []FloorplanSpaceChangelogPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *FloorplanCreateSpace) GetPathsOk() (*[]FloorplanSpaceChangelogPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *FloorplanCreateSpace) SetPaths(v []FloorplanSpaceChangelogPath)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


