# DashboardCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Items** | Pointer to [**[]DashboardCardItemRequest**](DashboardCardItemRequest.md) | The card&#39;s visualisations and controls, in display order. Each item sets exactly one of &#x60;visualisation&#x60; or &#x60;control&#x60;. A card cannot have two controls bound to the same variable, and a row-break card has no items. | [optional] 
**Kind** | Pointer to **int32** | &#x60;0&#x60; is a card of items and &#x60;1&#x60; a row break. | [optional] 
**Name** | **string** |  | 
**Size** | **int32** |  | 

## Methods

### NewDashboardCardRequest

`func NewDashboardCardRequest(description string, name string, size int32, ) *DashboardCardRequest`

NewDashboardCardRequest instantiates a new DashboardCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardRequestWithDefaults

`func NewDashboardCardRequestWithDefaults() *DashboardCardRequest`

NewDashboardCardRequestWithDefaults instantiates a new DashboardCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DashboardCardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardCardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardCardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetItems

`func (o *DashboardCardRequest) GetItems() []DashboardCardItemRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DashboardCardRequest) GetItemsOk() (*[]DashboardCardItemRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DashboardCardRequest) SetItems(v []DashboardCardItemRequest)`

SetItems sets Items field to given value.

### HasItems

`func (o *DashboardCardRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *DashboardCardRequest) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DashboardCardRequest) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DashboardCardRequest) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DashboardCardRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DashboardCardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardCardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardCardRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *DashboardCardRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DashboardCardRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DashboardCardRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


