# GraphPreviewInspectBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descending** | Pointer to **bool** | Reverses the order sort gives. | [optional] 
**Filters** | Pointer to [**[]FilterGroup**](FilterGroup.md) | Lists only the items whose objects these filter groups select, as a calculation&#39;s filters select them: any group may match, and within a group every filter must. An item of a set that isn&#39;t of objects is always listed. | [optional] 
**Limit** | **int64** | How many items to list. | 
**Node** | **string** | The node whose output to list. | 
**Offset** | Pointer to **int64** | How many matching items to skip. | [optional] 
**Search** | Pointer to **string** | Lists only the items whose name contains this, ignoring case. | [optional] 
**Socket** | **string** | The output to list. | 
**Sort** | Pointer to **string** | Orders the items by \&quot;item\&quot;, \&quot;value\&quot; or a property&#39;s id, items without one last; empty keeps them in the output&#39;s order. | [optional] 

## Methods

### NewGraphPreviewInspectBody

`func NewGraphPreviewInspectBody(limit int64, node string, socket string, ) *GraphPreviewInspectBody`

NewGraphPreviewInspectBody instantiates a new GraphPreviewInspectBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewInspectBodyWithDefaults

`func NewGraphPreviewInspectBodyWithDefaults() *GraphPreviewInspectBody`

NewGraphPreviewInspectBodyWithDefaults instantiates a new GraphPreviewInspectBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescending

`func (o *GraphPreviewInspectBody) GetDescending() bool`

GetDescending returns the Descending field if non-nil, zero value otherwise.

### GetDescendingOk

`func (o *GraphPreviewInspectBody) GetDescendingOk() (*bool, bool)`

GetDescendingOk returns a tuple with the Descending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescending

`func (o *GraphPreviewInspectBody) SetDescending(v bool)`

SetDescending sets Descending field to given value.

### HasDescending

`func (o *GraphPreviewInspectBody) HasDescending() bool`

HasDescending returns a boolean if a field has been set.

### GetFilters

`func (o *GraphPreviewInspectBody) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GraphPreviewInspectBody) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GraphPreviewInspectBody) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GraphPreviewInspectBody) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *GraphPreviewInspectBody) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GraphPreviewInspectBody) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GraphPreviewInspectBody) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetNode

`func (o *GraphPreviewInspectBody) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GraphPreviewInspectBody) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GraphPreviewInspectBody) SetNode(v string)`

SetNode sets Node field to given value.


### GetOffset

`func (o *GraphPreviewInspectBody) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GraphPreviewInspectBody) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GraphPreviewInspectBody) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GraphPreviewInspectBody) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearch

`func (o *GraphPreviewInspectBody) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *GraphPreviewInspectBody) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *GraphPreviewInspectBody) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *GraphPreviewInspectBody) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSocket

`func (o *GraphPreviewInspectBody) GetSocket() string`

GetSocket returns the Socket field if non-nil, zero value otherwise.

### GetSocketOk

`func (o *GraphPreviewInspectBody) GetSocketOk() (*string, bool)`

GetSocketOk returns a tuple with the Socket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocket

`func (o *GraphPreviewInspectBody) SetSocket(v string)`

SetSocket sets Socket field to given value.


### GetSort

`func (o *GraphPreviewInspectBody) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GraphPreviewInspectBody) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GraphPreviewInspectBody) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GraphPreviewInspectBody) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


