# DashboardChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | [**[]DashboardCard**](DashboardCard.md) |  | 
**Description** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewDashboardChangelogBody

`func NewDashboardChangelogBody(cards []DashboardCard, description string, name string, ) *DashboardChangelogBody`

NewDashboardChangelogBody instantiates a new DashboardChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardChangelogBodyWithDefaults

`func NewDashboardChangelogBodyWithDefaults() *DashboardChangelogBody`

NewDashboardChangelogBodyWithDefaults instantiates a new DashboardChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *DashboardChangelogBody) GetCards() []DashboardCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *DashboardChangelogBody) GetCardsOk() (*[]DashboardCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *DashboardChangelogBody) SetCards(v []DashboardCard)`

SetCards sets Cards field to given value.


### GetDescription

`func (o *DashboardChangelogBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardChangelogBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardChangelogBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *DashboardChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardChangelogBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


