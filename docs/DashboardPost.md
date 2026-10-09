# DashboardPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | [**[]DashboardCard**](DashboardCard.md) |  | 
**Description** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewDashboardPost

`func NewDashboardPost(cards []DashboardCard, description string, name string, ) *DashboardPost`

NewDashboardPost instantiates a new DashboardPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardPostWithDefaults

`func NewDashboardPostWithDefaults() *DashboardPost`

NewDashboardPostWithDefaults instantiates a new DashboardPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *DashboardPost) GetCards() []DashboardCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *DashboardPost) GetCardsOk() (*[]DashboardCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *DashboardPost) SetCards(v []DashboardCard)`

SetCards sets Cards field to given value.


### GetDescription

`func (o *DashboardPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardPost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DashboardPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DashboardPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardPost) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


