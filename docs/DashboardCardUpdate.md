# DashboardCardUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**DashboardCardRequest**](DashboardCardRequest.md) |  | 
**CardId** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 

## Methods

### NewDashboardCardUpdate

`func NewDashboardCardUpdate(card DashboardCardRequest, ) *DashboardCardUpdate`

NewDashboardCardUpdate instantiates a new DashboardCardUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardUpdateWithDefaults

`func NewDashboardCardUpdateWithDefaults() *DashboardCardUpdate`

NewDashboardCardUpdateWithDefaults instantiates a new DashboardCardUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *DashboardCardUpdate) GetCard() DashboardCardRequest`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *DashboardCardUpdate) GetCardOk() (*DashboardCardRequest, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *DashboardCardUpdate) SetCard(v DashboardCardRequest)`

SetCard sets Card field to given value.


### GetCardId

`func (o *DashboardCardUpdate) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *DashboardCardUpdate) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *DashboardCardUpdate) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *DashboardCardUpdate) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetIndex

`func (o *DashboardCardUpdate) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DashboardCardUpdate) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DashboardCardUpdate) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DashboardCardUpdate) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


