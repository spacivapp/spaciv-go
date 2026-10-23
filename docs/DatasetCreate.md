# DatasetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Steps** | [**[]QuestionStepRequest**](QuestionStepRequest.md) |  | 

## Methods

### NewDatasetCreate

`func NewDatasetCreate(isPublic bool, name string, steps []QuestionStepRequest, ) *DatasetCreate`

NewDatasetCreate instantiates a new DatasetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetCreateWithDefaults

`func NewDatasetCreateWithDefaults() *DatasetCreate`

NewDatasetCreateWithDefaults instantiates a new DatasetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasetCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *DatasetCreate) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *DatasetCreate) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *DatasetCreate) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *DatasetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *DatasetCreate) GetSteps() []QuestionStepRequest`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DatasetCreate) GetStepsOk() (*[]QuestionStepRequest, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DatasetCreate) SetSteps(v []QuestionStepRequest)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


