# DatasetRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DatasetChangelogBody**](DatasetChangelogBody.md) |  | 
**Id** | **string** |  | 

## Methods

### NewDatasetRes

`func NewDatasetRes(data DatasetChangelogBody, id string, ) *DatasetRes`

NewDatasetRes instantiates a new DatasetRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetResWithDefaults

`func NewDatasetResWithDefaults() *DatasetRes`

NewDatasetResWithDefaults instantiates a new DatasetRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatasetRes) GetData() DatasetChangelogBody`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetRes) GetDataOk() (*DatasetChangelogBody, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetRes) SetData(v DatasetChangelogBody)`

SetData sets Data field to given value.


### GetId

`func (o *DatasetRes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetRes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetRes) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


