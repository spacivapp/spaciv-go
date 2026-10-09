# FloorplanGetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**Name** | **string** |  | 
**OverlayOffset** | [**FloorplanGetBodyOverlayOffset**](FloorplanGetBodyOverlayOffset.md) |  | 
**Property** | **string** |  | 
**PropertyNode** | **string** |  | 
**Spaces** | [**[]FloorplanGetBodySpace**](FloorplanGetBodySpace.md) |  | 

## Methods

### NewFloorplanGetBody

`func NewFloorplanGetBody(createdAt int64, name string, overlayOffset FloorplanGetBodyOverlayOffset, property string, propertyNode string, spaces []FloorplanGetBodySpace, ) *FloorplanGetBody`

NewFloorplanGetBody instantiates a new FloorplanGetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorplanGetBodyWithDefaults

`func NewFloorplanGetBodyWithDefaults() *FloorplanGetBody`

NewFloorplanGetBodyWithDefaults instantiates a new FloorplanGetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FloorplanGetBody) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FloorplanGetBody) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FloorplanGetBody) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *FloorplanGetBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloorplanGetBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloorplanGetBody) SetName(v string)`

SetName sets Name field to given value.


### GetOverlayOffset

`func (o *FloorplanGetBody) GetOverlayOffset() FloorplanGetBodyOverlayOffset`

GetOverlayOffset returns the OverlayOffset field if non-nil, zero value otherwise.

### GetOverlayOffsetOk

`func (o *FloorplanGetBody) GetOverlayOffsetOk() (*FloorplanGetBodyOverlayOffset, bool)`

GetOverlayOffsetOk returns a tuple with the OverlayOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayOffset

`func (o *FloorplanGetBody) SetOverlayOffset(v FloorplanGetBodyOverlayOffset)`

SetOverlayOffset sets OverlayOffset field to given value.


### GetProperty

`func (o *FloorplanGetBody) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *FloorplanGetBody) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *FloorplanGetBody) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetPropertyNode

`func (o *FloorplanGetBody) GetPropertyNode() string`

GetPropertyNode returns the PropertyNode field if non-nil, zero value otherwise.

### GetPropertyNodeOk

`func (o *FloorplanGetBody) GetPropertyNodeOk() (*string, bool)`

GetPropertyNodeOk returns a tuple with the PropertyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNode

`func (o *FloorplanGetBody) SetPropertyNode(v string)`

SetPropertyNode sets PropertyNode field to given value.


### GetSpaces

`func (o *FloorplanGetBody) GetSpaces() []FloorplanGetBodySpace`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *FloorplanGetBody) GetSpacesOk() (*[]FloorplanGetBodySpace, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *FloorplanGetBody) SetSpaces(v []FloorplanGetBodySpace)`

SetSpaces sets Spaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


