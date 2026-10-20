# SelectObjectsParamsWrapperBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FilterGroup**](FilterGroup.md) | Filter groups, of which any may match; within a group every filter must. None selects every object. | 

## Methods

### NewSelectObjectsParamsWrapperBody

`func NewSelectObjectsParamsWrapperBody(filters []FilterGroup, ) *SelectObjectsParamsWrapperBody`

NewSelectObjectsParamsWrapperBody instantiates a new SelectObjectsParamsWrapperBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectObjectsParamsWrapperBodyWithDefaults

`func NewSelectObjectsParamsWrapperBodyWithDefaults() *SelectObjectsParamsWrapperBody`

NewSelectObjectsParamsWrapperBodyWithDefaults instantiates a new SelectObjectsParamsWrapperBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SelectObjectsParamsWrapperBody) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SelectObjectsParamsWrapperBody) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SelectObjectsParamsWrapperBody) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


