# CalcListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**List** | [**[]CalcListEntryRequest**](CalcListEntryRequest.md) |  | 
**TimeSpan** | Pointer to **int32** |  | [optional] 
**To** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 

## Methods

### NewCalcListRequest

`func NewCalcListRequest(filters []FilterGroup, list []CalcListEntryRequest, ) *CalcListRequest`

NewCalcListRequest instantiates a new CalcListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListRequestWithDefaults

`func NewCalcListRequestWithDefaults() *CalcListRequest`

NewCalcListRequestWithDefaults instantiates a new CalcListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcListRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcListRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcListRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcListRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcListRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcListRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcListRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetList

`func (o *CalcListRequest) GetList() []CalcListEntryRequest`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *CalcListRequest) GetListOk() (*[]CalcListEntryRequest, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *CalcListRequest) SetList(v []CalcListEntryRequest)`

SetList sets List field to given value.


### GetTimeSpan

`func (o *CalcListRequest) GetTimeSpan() int32`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *CalcListRequest) GetTimeSpanOk() (*int32, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *CalcListRequest) SetTimeSpan(v int32)`

SetTimeSpan sets TimeSpan field to given value.

### HasTimeSpan

`func (o *CalcListRequest) HasTimeSpan() bool`

HasTimeSpan returns a boolean if a field has been set.

### GetTo

`func (o *CalcListRequest) GetTo() MaybeVarInt`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CalcListRequest) GetToOk() (*MaybeVarInt, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CalcListRequest) SetTo(v MaybeVarInt)`

SetTo sets To field to given value.

### HasTo

`func (o *CalcListRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


