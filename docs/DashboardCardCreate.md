# DashboardCardCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterCardId** | Pointer to **string** |  | [optional] 
**Card** | [**DashboardCardRequest**](DashboardCardRequest.md) |  | 
**Index** | Pointer to **int32** |  | [optional] 

## Methods

### NewDashboardCardCreate

`func NewDashboardCardCreate(card DashboardCardRequest, ) *DashboardCardCreate`

NewDashboardCardCreate instantiates a new DashboardCardCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardCreateWithDefaults

`func NewDashboardCardCreateWithDefaults() *DashboardCardCreate`

NewDashboardCardCreateWithDefaults instantiates a new DashboardCardCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterCardId

`func (o *DashboardCardCreate) GetAfterCardId() string`

GetAfterCardId returns the AfterCardId field if non-nil, zero value otherwise.

### GetAfterCardIdOk

`func (o *DashboardCardCreate) GetAfterCardIdOk() (*string, bool)`

GetAfterCardIdOk returns a tuple with the AfterCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCardId

`func (o *DashboardCardCreate) SetAfterCardId(v string)`

SetAfterCardId sets AfterCardId field to given value.

### HasAfterCardId

`func (o *DashboardCardCreate) HasAfterCardId() bool`

HasAfterCardId returns a boolean if a field has been set.

### GetCard

`func (o *DashboardCardCreate) GetCard() DashboardCardRequest`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *DashboardCardCreate) GetCardOk() (*DashboardCardRequest, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *DashboardCardCreate) SetCard(v DashboardCardRequest)`

SetCard sets Card field to given value.


### GetIndex

`func (o *DashboardCardCreate) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DashboardCardCreate) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DashboardCardCreate) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DashboardCardCreate) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


