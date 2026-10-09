# DatasetModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisuaisationType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContributorDescription** | Pointer to **string** |  | [optional] 
**ContributorInfo** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**DatasetData**](DatasetData.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InUse** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewDatasetModify

`func NewDatasetModify() *DatasetModify`

NewDatasetModify instantiates a new DatasetModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetModifyWithDefaults

`func NewDatasetModifyWithDefaults() *DatasetModify`

NewDatasetModifyWithDefaults instantiates a new DatasetModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisuaisationType

`func (o *DatasetModify) GetVisuaisationType() string`

GetVisuaisationType returns the VisuaisationType field if non-nil, zero value otherwise.

### GetVisuaisationTypeOk

`func (o *DatasetModify) GetVisuaisationTypeOk() (*string, bool)`

GetVisuaisationTypeOk returns a tuple with the VisuaisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisuaisationType

`func (o *DatasetModify) SetVisuaisationType(v string)`

SetVisuaisationType sets VisuaisationType field to given value.

### HasVisuaisationType

`func (o *DatasetModify) HasVisuaisationType() bool`

HasVisuaisationType returns a boolean if a field has been set.

### GetCode

`func (o *DatasetModify) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DatasetModify) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DatasetModify) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DatasetModify) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContributorDescription

`func (o *DatasetModify) GetContributorDescription() string`

GetContributorDescription returns the ContributorDescription field if non-nil, zero value otherwise.

### GetContributorDescriptionOk

`func (o *DatasetModify) GetContributorDescriptionOk() (*string, bool)`

GetContributorDescriptionOk returns a tuple with the ContributorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorDescription

`func (o *DatasetModify) SetContributorDescription(v string)`

SetContributorDescription sets ContributorDescription field to given value.

### HasContributorDescription

`func (o *DatasetModify) HasContributorDescription() bool`

HasContributorDescription returns a boolean if a field has been set.

### GetContributorInfo

`func (o *DatasetModify) GetContributorInfo() string`

GetContributorInfo returns the ContributorInfo field if non-nil, zero value otherwise.

### GetContributorInfoOk

`func (o *DatasetModify) GetContributorInfoOk() (*string, bool)`

GetContributorInfoOk returns a tuple with the ContributorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorInfo

`func (o *DatasetModify) SetContributorInfo(v string)`

SetContributorInfo sets ContributorInfo field to given value.

### HasContributorInfo

`func (o *DatasetModify) HasContributorInfo() bool`

HasContributorInfo returns a boolean if a field has been set.

### GetData

`func (o *DatasetModify) GetData() DatasetData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetModify) GetDataOk() (*DatasetData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetModify) SetData(v DatasetData)`

SetData sets Data field to given value.

### HasData

`func (o *DatasetModify) HasData() bool`

HasData returns a boolean if a field has been set.

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

### GetInUse

`func (o *DatasetModify) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DatasetModify) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DatasetModify) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *DatasetModify) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

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

### GetQuestion

`func (o *DatasetModify) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *DatasetModify) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *DatasetModify) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *DatasetModify) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetType

`func (o *DatasetModify) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetModify) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetModify) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DatasetModify) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


