# GraphLayersDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | **[]string** |  | 
**Failed** | [**[]GraphLayerDeleteFailure**](GraphLayerDeleteFailure.md) |  | 

## Methods

### NewGraphLayersDeleted

`func NewGraphLayersDeleted(deleted []string, failed []GraphLayerDeleteFailure, ) *GraphLayersDeleted`

NewGraphLayersDeleted instantiates a new GraphLayersDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphLayersDeletedWithDefaults

`func NewGraphLayersDeletedWithDefaults() *GraphLayersDeleted`

NewGraphLayersDeletedWithDefaults instantiates a new GraphLayersDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *GraphLayersDeleted) GetDeleted() []string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GraphLayersDeleted) GetDeletedOk() (*[]string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GraphLayersDeleted) SetDeleted(v []string)`

SetDeleted sets Deleted field to given value.


### GetFailed

`func (o *GraphLayersDeleted) GetFailed() []GraphLayerDeleteFailure`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GraphLayersDeleted) GetFailedOk() (*[]GraphLayerDeleteFailure, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GraphLayersDeleted) SetFailed(v []GraphLayerDeleteFailure)`

SetFailed sets Failed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


