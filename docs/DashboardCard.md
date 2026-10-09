# DashboardCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Items** | Pointer to [**[]DashboardCardItem**](DashboardCardItem.md) | The card&#39;s visualisations and controls, in display order. Each item sets exactly one of &#x60;visualisation&#x60; or &#x60;control&#x60;. A card cannot have two controls bound to the same variable, and a row-break card has no items. | [optional] 
**Kind** | Pointer to **int32** | &#x60;0&#x60; is a card of items and &#x60;1&#x60; a row break. | [optional] 
**Name** | **string** |  | 
**Size** | **int32** |  | 

## Methods

### NewDashboardCard

`func NewDashboardCard(description string, name string, size int32, ) *DashboardCard`

NewDashboardCard instantiates a new DashboardCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardWithDefaults

`func NewDashboardCardWithDefaults() *DashboardCard`

NewDashboardCardWithDefaults instantiates a new DashboardCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DashboardCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardCard) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetItems

`func (o *DashboardCard) GetItems() []DashboardCardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DashboardCard) GetItemsOk() (*[]DashboardCardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DashboardCard) SetItems(v []DashboardCardItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DashboardCard) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *DashboardCard) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DashboardCard) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DashboardCard) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DashboardCard) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DashboardCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardCard) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *DashboardCard) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DashboardCard) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DashboardCard) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


