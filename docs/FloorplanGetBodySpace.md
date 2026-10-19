# FloorplanGetBodySpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndHeight** | Pointer to **float64** | Height the space ends at in 3D, in the spaces&#39; path units, as written by editors that predate metres. Absent once converted to metres. | [optional] 
**EndHeightMetres** | Pointer to **float64** | Height the space ends at in 3D, in metres. Absent when the space uses the default heights. | [optional] 
**Location** | **string** |  | 
**Paths** | [**[][]FloorplanGetBodySpacePoint**]([]FloorplanGetBodySpacePoint.md) |  | 
**StartHeight** | Pointer to **float64** | Height the space starts at in 3D, in the spaces&#39; path units, as written by editors that predate metres. Absent once converted to metres. | [optional] 
**StartHeightMetres** | Pointer to **float64** | Height the space starts at in 3D, in metres. Absent when the space uses the default heights. | [optional] 

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

### GetEndHeight

`func (o *FloorplanGetBodySpace) GetEndHeight() float64`

GetEndHeight returns the EndHeight field if non-nil, zero value otherwise.

### GetEndHeightOk

`func (o *FloorplanGetBodySpace) GetEndHeightOk() (*float64, bool)`

GetEndHeightOk returns a tuple with the EndHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHeight

`func (o *FloorplanGetBodySpace) SetEndHeight(v float64)`

SetEndHeight sets EndHeight field to given value.

### HasEndHeight

`func (o *FloorplanGetBodySpace) HasEndHeight() bool`

HasEndHeight returns a boolean if a field has been set.

### GetEndHeightMetres

`func (o *FloorplanGetBodySpace) GetEndHeightMetres() float64`

GetEndHeightMetres returns the EndHeightMetres field if non-nil, zero value otherwise.

### GetEndHeightMetresOk

`func (o *FloorplanGetBodySpace) GetEndHeightMetresOk() (*float64, bool)`

GetEndHeightMetresOk returns a tuple with the EndHeightMetres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHeightMetres

`func (o *FloorplanGetBodySpace) SetEndHeightMetres(v float64)`

SetEndHeightMetres sets EndHeightMetres field to given value.

### HasEndHeightMetres

`func (o *FloorplanGetBodySpace) HasEndHeightMetres() bool`

HasEndHeightMetres returns a boolean if a field has been set.

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


### GetStartHeight

`func (o *FloorplanGetBodySpace) GetStartHeight() float64`

GetStartHeight returns the StartHeight field if non-nil, zero value otherwise.

### GetStartHeightOk

`func (o *FloorplanGetBodySpace) GetStartHeightOk() (*float64, bool)`

GetStartHeightOk returns a tuple with the StartHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHeight

`func (o *FloorplanGetBodySpace) SetStartHeight(v float64)`

SetStartHeight sets StartHeight field to given value.

### HasStartHeight

`func (o *FloorplanGetBodySpace) HasStartHeight() bool`

HasStartHeight returns a boolean if a field has been set.

### GetStartHeightMetres

`func (o *FloorplanGetBodySpace) GetStartHeightMetres() float64`

GetStartHeightMetres returns the StartHeightMetres field if non-nil, zero value otherwise.

### GetStartHeightMetresOk

`func (o *FloorplanGetBodySpace) GetStartHeightMetresOk() (*float64, bool)`

GetStartHeightMetresOk returns a tuple with the StartHeightMetres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHeightMetres

`func (o *FloorplanGetBodySpace) SetStartHeightMetres(v float64)`

SetStartHeightMetres sets StartHeightMetres field to given value.

### HasStartHeightMetres

`func (o *FloorplanGetBodySpace) HasStartHeightMetres() bool`

HasStartHeightMetres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


