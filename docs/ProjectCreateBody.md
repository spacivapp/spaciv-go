# ProjectCreateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Name** | **string** |  | 
**Blank** | **map[string]interface{}** |  | 
**Duplicate** | **int32** |  | 
**Template** | [**ProjectCreateTemplate**](ProjectCreateTemplate.md) |  | 

## Methods

### NewProjectCreateBody

`func NewProjectCreateBody(description string, name string, blank map[string]interface{}, duplicate int32, template ProjectCreateTemplate, ) *ProjectCreateBody`

NewProjectCreateBody instantiates a new ProjectCreateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateBodyWithDefaults

`func NewProjectCreateBodyWithDefaults() *ProjectCreateBody`

NewProjectCreateBodyWithDefaults instantiates a new ProjectCreateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProjectCreateBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCreateBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCreateBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *ProjectCreateBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreateBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreateBody) SetName(v string)`

SetName sets Name field to given value.


### GetBlank

`func (o *ProjectCreateBody) GetBlank() map[string]interface{}`

GetBlank returns the Blank field if non-nil, zero value otherwise.

### GetBlankOk

`func (o *ProjectCreateBody) GetBlankOk() (*map[string]interface{}, bool)`

GetBlankOk returns a tuple with the Blank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlank

`func (o *ProjectCreateBody) SetBlank(v map[string]interface{})`

SetBlank sets Blank field to given value.


### GetDuplicate

`func (o *ProjectCreateBody) GetDuplicate() int32`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *ProjectCreateBody) GetDuplicateOk() (*int32, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *ProjectCreateBody) SetDuplicate(v int32)`

SetDuplicate sets Duplicate field to given value.


### GetTemplate

`func (o *ProjectCreateBody) GetTemplate() ProjectCreateTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ProjectCreateBody) GetTemplateOk() (*ProjectCreateTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ProjectCreateBody) SetTemplate(v ProjectCreateTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


