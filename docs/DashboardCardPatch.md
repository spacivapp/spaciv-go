# DashboardCardPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveCard** | [**DashboardCardMove**](DashboardCardMove.md) |  | 
**ReorderCard** | [**DashboardCardMove**](DashboardCardMove.md) |  | 
**UpdateCard** | [**DashboardCardUpdate**](DashboardCardUpdate.md) |  | 

## Methods

### NewDashboardCardPatch

`func NewDashboardCardPatch(moveCard DashboardCardMove, reorderCard DashboardCardMove, updateCard DashboardCardUpdate, ) *DashboardCardPatch`

NewDashboardCardPatch instantiates a new DashboardCardPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardPatchWithDefaults

`func NewDashboardCardPatchWithDefaults() *DashboardCardPatch`

NewDashboardCardPatchWithDefaults instantiates a new DashboardCardPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveCard

`func (o *DashboardCardPatch) GetMoveCard() DashboardCardMove`

GetMoveCard returns the MoveCard field if non-nil, zero value otherwise.

### GetMoveCardOk

`func (o *DashboardCardPatch) GetMoveCardOk() (*DashboardCardMove, bool)`

GetMoveCardOk returns a tuple with the MoveCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCard

`func (o *DashboardCardPatch) SetMoveCard(v DashboardCardMove)`

SetMoveCard sets MoveCard field to given value.


### GetReorderCard

`func (o *DashboardCardPatch) GetReorderCard() DashboardCardMove`

GetReorderCard returns the ReorderCard field if non-nil, zero value otherwise.

### GetReorderCardOk

`func (o *DashboardCardPatch) GetReorderCardOk() (*DashboardCardMove, bool)`

GetReorderCardOk returns a tuple with the ReorderCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderCard

`func (o *DashboardCardPatch) SetReorderCard(v DashboardCardMove)`

SetReorderCard sets ReorderCard field to given value.


### GetUpdateCard

`func (o *DashboardCardPatch) GetUpdateCard() DashboardCardUpdate`

GetUpdateCard returns the UpdateCard field if non-nil, zero value otherwise.

### GetUpdateCardOk

`func (o *DashboardCardPatch) GetUpdateCardOk() (*DashboardCardUpdate, bool)`

GetUpdateCardOk returns a tuple with the UpdateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCard

`func (o *DashboardCardPatch) SetUpdateCard(v DashboardCardUpdate)`

SetUpdateCard sets UpdateCard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


