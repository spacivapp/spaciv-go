# DashboardCardPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveCardAfter** | [**DashboardCardMoveAfter**](DashboardCardMoveAfter.md) |  | 
**UpdateCard** | [**DashboardCardUpdate**](DashboardCardUpdate.md) |  | 

## Methods

### NewDashboardCardPatch

`func NewDashboardCardPatch(moveCardAfter DashboardCardMoveAfter, updateCard DashboardCardUpdate, ) *DashboardCardPatch`

NewDashboardCardPatch instantiates a new DashboardCardPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardPatchWithDefaults

`func NewDashboardCardPatchWithDefaults() *DashboardCardPatch`

NewDashboardCardPatchWithDefaults instantiates a new DashboardCardPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveCardAfter

`func (o *DashboardCardPatch) GetMoveCardAfter() DashboardCardMoveAfter`

GetMoveCardAfter returns the MoveCardAfter field if non-nil, zero value otherwise.

### GetMoveCardAfterOk

`func (o *DashboardCardPatch) GetMoveCardAfterOk() (*DashboardCardMoveAfter, bool)`

GetMoveCardAfterOk returns a tuple with the MoveCardAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCardAfter

`func (o *DashboardCardPatch) SetMoveCardAfter(v DashboardCardMoveAfter)`

SetMoveCardAfter sets MoveCardAfter field to given value.


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


