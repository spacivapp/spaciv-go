# GraphFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 

## Methods

### NewGraphFilter

`func NewGraphFilter(filters []FilterGroup, ) *GraphFilter`

NewGraphFilter instantiates a new GraphFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphFilterWithDefaults

`func NewGraphFilterWithDefaults() *GraphFilter`

NewGraphFilterWithDefaults instantiates a new GraphFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GraphFilter) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GraphFilter) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GraphFilter) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


