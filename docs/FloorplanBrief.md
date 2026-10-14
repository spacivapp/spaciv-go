# FloorplanBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundFile** | Pointer to **string** | ID of the account file used as the background image, if one is set. | [optional] 
**CreatedAt** | **int64** |  | 
**ForegroundFile** | Pointer to **string** | ID of the account file used as the foreground image, if one is set. | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Property** | **string** |  | 
**PropertyNode** | **string** |  | 

## Methods

### NewFloorplanBrief

`func NewFloorplanBrief(createdAt int64, id string, name string, property string, propertyNode string, ) *FloorplanBrief`

NewFloorplanBrief instantiates a new FloorplanBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorplanBriefWithDefaults

`func NewFloorplanBriefWithDefaults() *FloorplanBrief`

NewFloorplanBriefWithDefaults instantiates a new FloorplanBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundFile

`func (o *FloorplanBrief) GetBackgroundFile() string`

GetBackgroundFile returns the BackgroundFile field if non-nil, zero value otherwise.

### GetBackgroundFileOk

`func (o *FloorplanBrief) GetBackgroundFileOk() (*string, bool)`

GetBackgroundFileOk returns a tuple with the BackgroundFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundFile

`func (o *FloorplanBrief) SetBackgroundFile(v string)`

SetBackgroundFile sets BackgroundFile field to given value.

### HasBackgroundFile

`func (o *FloorplanBrief) HasBackgroundFile() bool`

HasBackgroundFile returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FloorplanBrief) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FloorplanBrief) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FloorplanBrief) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetForegroundFile

`func (o *FloorplanBrief) GetForegroundFile() string`

GetForegroundFile returns the ForegroundFile field if non-nil, zero value otherwise.

### GetForegroundFileOk

`func (o *FloorplanBrief) GetForegroundFileOk() (*string, bool)`

GetForegroundFileOk returns a tuple with the ForegroundFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundFile

`func (o *FloorplanBrief) SetForegroundFile(v string)`

SetForegroundFile sets ForegroundFile field to given value.

### HasForegroundFile

`func (o *FloorplanBrief) HasForegroundFile() bool`

HasForegroundFile returns a boolean if a field has been set.

### GetId

`func (o *FloorplanBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloorplanBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloorplanBrief) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FloorplanBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloorplanBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloorplanBrief) SetName(v string)`

SetName sets Name field to given value.


### GetProperty

`func (o *FloorplanBrief) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *FloorplanBrief) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *FloorplanBrief) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetPropertyNode

`func (o *FloorplanBrief) GetPropertyNode() string`

GetPropertyNode returns the PropertyNode field if non-nil, zero value otherwise.

### GetPropertyNodeOk

`func (o *FloorplanBrief) GetPropertyNodeOk() (*string, bool)`

GetPropertyNodeOk returns a tuple with the PropertyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNode

`func (o *FloorplanBrief) SetPropertyNode(v string)`

SetPropertyNode sets PropertyNode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


