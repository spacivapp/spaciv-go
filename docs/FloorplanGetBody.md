# FloorplanGetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundFile** | Pointer to **string** | ID of the account file used as the background image, if one is set. | [optional] 
**CreatedAt** | **int64** |  | 
**ForegroundFile** | Pointer to **string** | ID of the account file used as the foreground image, if one is set. | [optional] 
**Name** | **string** |  | 
**OverlayOffset** | [**FloorplanGetBodyOverlayOffset**](FloorplanGetBodyOverlayOffset.md) |  | 
**Property** | **string** |  | 
**PropertyNode** | **string** |  | 
**Spaces** | [**[]FloorplanGetBodySpace**](FloorplanGetBodySpace.md) |  | 
**UnitsPerMetre** | Pointer to **float64** | How many of the spaces&#39; path units make one metre, which scales space heights in 3D. Absent until the floor plan&#39;s scale has been set. | [optional] 
**ViewMode** | **string** | Whether the floor plan can be shown in 2D, in 3D, or both. Defaults to 2d-and-3d. | 

## Methods

### NewFloorplanGetBody

`func NewFloorplanGetBody(createdAt int64, name string, overlayOffset FloorplanGetBodyOverlayOffset, property string, propertyNode string, spaces []FloorplanGetBodySpace, viewMode string, ) *FloorplanGetBody`

NewFloorplanGetBody instantiates a new FloorplanGetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorplanGetBodyWithDefaults

`func NewFloorplanGetBodyWithDefaults() *FloorplanGetBody`

NewFloorplanGetBodyWithDefaults instantiates a new FloorplanGetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundFile

`func (o *FloorplanGetBody) GetBackgroundFile() string`

GetBackgroundFile returns the BackgroundFile field if non-nil, zero value otherwise.

### GetBackgroundFileOk

`func (o *FloorplanGetBody) GetBackgroundFileOk() (*string, bool)`

GetBackgroundFileOk returns a tuple with the BackgroundFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundFile

`func (o *FloorplanGetBody) SetBackgroundFile(v string)`

SetBackgroundFile sets BackgroundFile field to given value.

### HasBackgroundFile

`func (o *FloorplanGetBody) HasBackgroundFile() bool`

HasBackgroundFile returns a boolean if a field has been set.

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


### GetForegroundFile

`func (o *FloorplanGetBody) GetForegroundFile() string`

GetForegroundFile returns the ForegroundFile field if non-nil, zero value otherwise.

### GetForegroundFileOk

`func (o *FloorplanGetBody) GetForegroundFileOk() (*string, bool)`

GetForegroundFileOk returns a tuple with the ForegroundFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundFile

`func (o *FloorplanGetBody) SetForegroundFile(v string)`

SetForegroundFile sets ForegroundFile field to given value.

### HasForegroundFile

`func (o *FloorplanGetBody) HasForegroundFile() bool`

HasForegroundFile returns a boolean if a field has been set.

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


### GetUnitsPerMetre

`func (o *FloorplanGetBody) GetUnitsPerMetre() float64`

GetUnitsPerMetre returns the UnitsPerMetre field if non-nil, zero value otherwise.

### GetUnitsPerMetreOk

`func (o *FloorplanGetBody) GetUnitsPerMetreOk() (*float64, bool)`

GetUnitsPerMetreOk returns a tuple with the UnitsPerMetre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsPerMetre

`func (o *FloorplanGetBody) SetUnitsPerMetre(v float64)`

SetUnitsPerMetre sets UnitsPerMetre field to given value.

### HasUnitsPerMetre

`func (o *FloorplanGetBody) HasUnitsPerMetre() bool`

HasUnitsPerMetre returns a boolean if a field has been set.

### GetViewMode

`func (o *FloorplanGetBody) GetViewMode() string`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *FloorplanGetBody) GetViewModeOk() (*string, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *FloorplanGetBody) SetViewMode(v string)`

SetViewMode sets ViewMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


