# DatasetChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**ContributorDescription** | **string** |  | 
**ContributorInfo** | **string** |  | 
**Data** | [**DatasetChangelogData**](DatasetChangelogData.md) |  | 
**Description** | **string** |  | 
**InUse** | **bool** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Published** | **bool** |  | 
**Question** | **string** |  | 
**Type** | **int32** |  | 
**VisualisationType** | **string** |  | 

## Methods

### NewDatasetChangelogBody

`func NewDatasetChangelogBody(code string, contributorDescription string, contributorInfo string, data DatasetChangelogData, description string, inUse bool, isPublic bool, name string, published bool, question string, type_ int32, visualisationType string, ) *DatasetChangelogBody`

NewDatasetChangelogBody instantiates a new DatasetChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetChangelogBodyWithDefaults

`func NewDatasetChangelogBodyWithDefaults() *DatasetChangelogBody`

NewDatasetChangelogBodyWithDefaults instantiates a new DatasetChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DatasetChangelogBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DatasetChangelogBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DatasetChangelogBody) SetCode(v string)`

SetCode sets Code field to given value.


### GetContributorDescription

`func (o *DatasetChangelogBody) GetContributorDescription() string`

GetContributorDescription returns the ContributorDescription field if non-nil, zero value otherwise.

### GetContributorDescriptionOk

`func (o *DatasetChangelogBody) GetContributorDescriptionOk() (*string, bool)`

GetContributorDescriptionOk returns a tuple with the ContributorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorDescription

`func (o *DatasetChangelogBody) SetContributorDescription(v string)`

SetContributorDescription sets ContributorDescription field to given value.


### GetContributorInfo

`func (o *DatasetChangelogBody) GetContributorInfo() string`

GetContributorInfo returns the ContributorInfo field if non-nil, zero value otherwise.

### GetContributorInfoOk

`func (o *DatasetChangelogBody) GetContributorInfoOk() (*string, bool)`

GetContributorInfoOk returns a tuple with the ContributorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorInfo

`func (o *DatasetChangelogBody) SetContributorInfo(v string)`

SetContributorInfo sets ContributorInfo field to given value.


### GetData

`func (o *DatasetChangelogBody) GetData() DatasetChangelogData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetChangelogBody) GetDataOk() (*DatasetChangelogData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetChangelogBody) SetData(v DatasetChangelogData)`

SetData sets Data field to given value.


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


### GetInUse

`func (o *DatasetChangelogBody) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DatasetChangelogBody) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DatasetChangelogBody) SetInUse(v bool)`

SetInUse sets InUse field to given value.


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


### GetQuestion

`func (o *DatasetChangelogBody) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *DatasetChangelogBody) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *DatasetChangelogBody) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetType

`func (o *DatasetChangelogBody) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetChangelogBody) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetChangelogBody) SetType(v int32)`

SetType sets Type field to given value.


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


