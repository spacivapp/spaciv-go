# ProjectPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFilter** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**KeepFilters** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**Variables**](Variables.md) |  | [optional] 

## Methods

### NewProjectPatchBody

`func NewProjectPatchBody(deleteFilter bool, filters []FilterGroup, ) *ProjectPatchBody`

NewProjectPatchBody instantiates a new ProjectPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPatchBodyWithDefaults

`func NewProjectPatchBodyWithDefaults() *ProjectPatchBody`

NewProjectPatchBodyWithDefaults instantiates a new ProjectPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFilter

`func (o *ProjectPatchBody) GetDeleteFilter() bool`

GetDeleteFilter returns the DeleteFilter field if non-nil, zero value otherwise.

### GetDeleteFilterOk

`func (o *ProjectPatchBody) GetDeleteFilterOk() (*bool, bool)`

GetDeleteFilterOk returns a tuple with the DeleteFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFilter

`func (o *ProjectPatchBody) SetDeleteFilter(v bool)`

SetDeleteFilter sets DeleteFilter field to given value.


### GetDescription

`func (o *ProjectPatchBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectPatchBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectPatchBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectPatchBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *ProjectPatchBody) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProjectPatchBody) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProjectPatchBody) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetKeepFilters

`func (o *ProjectPatchBody) GetKeepFilters() bool`

GetKeepFilters returns the KeepFilters field if non-nil, zero value otherwise.

### GetKeepFiltersOk

`func (o *ProjectPatchBody) GetKeepFiltersOk() (*bool, bool)`

GetKeepFiltersOk returns a tuple with the KeepFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFilters

`func (o *ProjectPatchBody) SetKeepFilters(v bool)`

SetKeepFilters sets KeepFilters field to given value.

### HasKeepFilters

`func (o *ProjectPatchBody) HasKeepFilters() bool`

HasKeepFilters returns a boolean if a field has been set.

### GetName

`func (o *ProjectPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariables

`func (o *ProjectPatchBody) GetVariables() Variables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProjectPatchBody) GetVariablesOk() (*Variables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProjectPatchBody) SetVariables(v Variables)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProjectPatchBody) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


