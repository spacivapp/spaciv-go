# DatasetChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Published** | **bool** |  | 
**Steps** | [**[]QuestionStep**](QuestionStep.md) |  | 
**VisualisationType** | **string** |  | 

## Methods

### NewDatasetChangelogBody

`func NewDatasetChangelogBody(description string, isPublic bool, name string, published bool, steps []QuestionStep, visualisationType string, ) *DatasetChangelogBody`

NewDatasetChangelogBody instantiates a new DatasetChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetChangelogBodyWithDefaults

`func NewDatasetChangelogBodyWithDefaults() *DatasetChangelogBody`

NewDatasetChangelogBodyWithDefaults instantiates a new DatasetChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasetChangelogBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetChangelogBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetChangelogBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsPublic

`func (o *DatasetChangelogBody) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *DatasetChangelogBody) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *DatasetChangelogBody) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *DatasetChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetChangelogBody) SetName(v string)`

SetName sets Name field to given value.


### GetPublished

`func (o *DatasetChangelogBody) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DatasetChangelogBody) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DatasetChangelogBody) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetSteps

`func (o *DatasetChangelogBody) GetSteps() []QuestionStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DatasetChangelogBody) GetStepsOk() (*[]QuestionStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DatasetChangelogBody) SetSteps(v []QuestionStep)`

SetSteps sets Steps field to given value.


### GetVisualisationType

`func (o *DatasetChangelogBody) GetVisualisationType() string`

GetVisualisationType returns the VisualisationType field if non-nil, zero value otherwise.

### GetVisualisationTypeOk

`func (o *DatasetChangelogBody) GetVisualisationTypeOk() (*string, bool)`

GetVisualisationTypeOk returns a tuple with the VisualisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisationType

`func (o *DatasetChangelogBody) SetVisualisationType(v string)`

SetVisualisationType sets VisualisationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


