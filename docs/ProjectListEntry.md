# ProjectListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Filter** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Id** | **float64** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**Variables** | [**map[string]VariableChangelog**](VariableChangelog.md) |  | 

## Methods

### NewProjectListEntry

`func NewProjectListEntry(description string, filter []FilterGroup, id float64, name string, status string, variables map[string]VariableChangelog, ) *ProjectListEntry`

NewProjectListEntry instantiates a new ProjectListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListEntryWithDefaults

`func NewProjectListEntryWithDefaults() *ProjectListEntry`

NewProjectListEntryWithDefaults instantiates a new ProjectListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProjectListEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectListEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectListEntry) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFilter

`func (o *ProjectListEntry) GetFilter() []FilterGroup`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ProjectListEntry) GetFilterOk() (*[]FilterGroup, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ProjectListEntry) SetFilter(v []FilterGroup)`

SetFilter sets Filter field to given value.


### GetId

`func (o *ProjectListEntry) GetId() float64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectListEntry) GetIdOk() (*float64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectListEntry) SetId(v float64)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectListEntry) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ProjectListEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListEntry) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVariables

`func (o *ProjectListEntry) GetVariables() map[string]VariableChangelog`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProjectListEntry) GetVariablesOk() (*map[string]VariableChangelog, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProjectListEntry) SetVariables(v map[string]VariableChangelog)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


