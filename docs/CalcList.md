# CalcList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**List** | [**[]CalcListEntry**](CalcListEntry.md) |  | 
**TimeSpan** | Pointer to **int32** |  | [optional] 
**To** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 

## Methods

### NewCalcList

`func NewCalcList(filters []FilterGroup, list []CalcListEntry, ) *CalcList`

NewCalcList instantiates a new CalcList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListWithDefaults

`func NewCalcListWithDefaults() *CalcList`

NewCalcListWithDefaults instantiates a new CalcList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcList) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcList) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcList) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcList) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalcList) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcList) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcList) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetList

`func (o *CalcList) GetList() []CalcListEntry`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *CalcList) GetListOk() (*[]CalcListEntry, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *CalcList) SetList(v []CalcListEntry)`

SetList sets List field to given value.


### GetTimeSpan

`func (o *CalcList) GetTimeSpan() int32`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *CalcList) GetTimeSpanOk() (*int32, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *CalcList) SetTimeSpan(v int32)`

SetTimeSpan sets TimeSpan field to given value.

### HasTimeSpan

`func (o *CalcList) HasTimeSpan() bool`

HasTimeSpan returns a boolean if a field has been set.

### GetTo

`func (o *CalcList) GetTo() MaybeVarInt`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CalcList) GetToOk() (*MaybeVarInt, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CalcList) SetTo(v MaybeVarInt)`

SetTo sets To field to given value.

### HasTo

`func (o *CalcList) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


