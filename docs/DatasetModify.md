# DatasetModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**ReplaceSteps** | **bool** |  | 
**Steps** | [**[]QuestionStepRequest**](QuestionStepRequest.md) |  | 
**VisualisationType** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasetModify

`func NewDatasetModify(replaceSteps bool, steps []QuestionStepRequest, ) *DatasetModify`

NewDatasetModify instantiates a new DatasetModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetModifyWithDefaults

`func NewDatasetModifyWithDefaults() *DatasetModify`

NewDatasetModifyWithDefaults instantiates a new DatasetModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasetModify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetModify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetModify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetModify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *DatasetModify) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *DatasetModify) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *DatasetModify) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *DatasetModify) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *DatasetModify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetModify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetModify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetModify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublished

`func (o *DatasetModify) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DatasetModify) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DatasetModify) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *DatasetModify) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReplaceSteps

`func (o *DatasetModify) GetReplaceSteps() bool`

GetReplaceSteps returns the ReplaceSteps field if non-nil, zero value otherwise.

### GetReplaceStepsOk

`func (o *DatasetModify) GetReplaceStepsOk() (*bool, bool)`

GetReplaceStepsOk returns a tuple with the ReplaceSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceSteps

`func (o *DatasetModify) SetReplaceSteps(v bool)`

SetReplaceSteps sets ReplaceSteps field to given value.


### GetSteps

`func (o *DatasetModify) GetSteps() []QuestionStepRequest`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DatasetModify) GetStepsOk() (*[]QuestionStepRequest, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DatasetModify) SetSteps(v []QuestionStepRequest)`

SetSteps sets Steps field to given value.


### GetVisualisationType

`func (o *DatasetModify) GetVisualisationType() string`

GetVisualisationType returns the VisualisationType field if non-nil, zero value otherwise.

### GetVisualisationTypeOk

`func (o *DatasetModify) GetVisualisationTypeOk() (*string, bool)`

GetVisualisationTypeOk returns a tuple with the VisualisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisationType

`func (o *DatasetModify) SetVisualisationType(v string)`

SetVisualisationType sets VisualisationType field to given value.

### HasVisualisationType

`func (o *DatasetModify) HasVisualisationType() bool`

HasVisualisationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


