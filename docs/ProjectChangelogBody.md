# ProjectChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Name** | **string** |  | 
**Variables** | [**map[string]VariableChangelog**](VariableChangelog.md) |  | 

## Methods

### NewProjectChangelogBody

`func NewProjectChangelogBody(description string, filters []FilterGroup, name string, variables map[string]VariableChangelog, ) *ProjectChangelogBody`

NewProjectChangelogBody instantiates a new ProjectChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectChangelogBodyWithDefaults

`func NewProjectChangelogBodyWithDefaults() *ProjectChangelogBody`

NewProjectChangelogBodyWithDefaults instantiates a new ProjectChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProjectChangelogBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectChangelogBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectChangelogBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFilters

`func (o *ProjectChangelogBody) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProjectChangelogBody) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProjectChangelogBody) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetName

`func (o *ProjectChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectChangelogBody) SetName(v string)`

SetName sets Name field to given value.


### GetVariables

`func (o *ProjectChangelogBody) GetVariables() map[string]VariableChangelog`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProjectChangelogBody) GetVariablesOk() (*map[string]VariableChangelog, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProjectChangelogBody) SetVariables(v map[string]VariableChangelog)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


