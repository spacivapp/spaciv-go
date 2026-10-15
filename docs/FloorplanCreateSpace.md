# FloorplanCreateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndHeight** | Pointer to **float64** |  | [optional] 
**Location** | **string** |  | 
**Paths** | [**[]FloorplanSpaceChangelogPath**](FloorplanSpaceChangelogPath.md) |  | 
**StartHeight** | Pointer to **float64** |  | [optional] 

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

### GetEndHeight

`func (o *FloorplanCreateSpace) GetEndHeight() float64`

GetEndHeight returns the EndHeight field if non-nil, zero value otherwise.

### GetEndHeightOk

`func (o *FloorplanCreateSpace) GetEndHeightOk() (*float64, bool)`

GetEndHeightOk returns a tuple with the EndHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHeight

`func (o *FloorplanCreateSpace) SetEndHeight(v float64)`

SetEndHeight sets EndHeight field to given value.

### HasEndHeight

`func (o *FloorplanCreateSpace) HasEndHeight() bool`

HasEndHeight returns a boolean if a field has been set.

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


### GetStartHeight

`func (o *FloorplanCreateSpace) GetStartHeight() float64`

GetStartHeight returns the StartHeight field if non-nil, zero value otherwise.

### GetStartHeightOk

`func (o *FloorplanCreateSpace) GetStartHeightOk() (*float64, bool)`

GetStartHeightOk returns a tuple with the StartHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHeight

`func (o *FloorplanCreateSpace) SetStartHeight(v float64)`

SetStartHeight sets StartHeight field to given value.

### HasStartHeight

`func (o *FloorplanCreateSpace) HasStartHeight() bool`

HasStartHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


