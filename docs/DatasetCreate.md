# DatasetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DatasetChangelogData**](DatasetChangelogData.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | **int32** |  | 

## Methods

### NewDatasetCreate

`func NewDatasetCreate(data DatasetChangelogData, name string, type_ int32, ) *DatasetCreate`

NewDatasetCreate instantiates a new DatasetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetCreateWithDefaults

`func NewDatasetCreateWithDefaults() *DatasetCreate`

NewDatasetCreateWithDefaults instantiates a new DatasetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatasetCreate) GetData() DatasetChangelogData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetCreate) GetDataOk() (*DatasetChangelogData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetCreate) SetData(v DatasetChangelogData)`

SetData sets Data field to given value.


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


### GetType

`func (o *DatasetCreate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetCreate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetCreate) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


