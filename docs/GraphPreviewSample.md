# GraphPreviewSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Item** | **string** |  | 
**Missing** | **bool** |  | 
**Objects** | [**[]GraphPreviewObject**](GraphPreviewObject.md) |  | 
**Value** | **float64** |  | 

## Methods

### NewGraphPreviewSample

`func NewGraphPreviewSample(error_ string, item string, missing bool, objects []GraphPreviewObject, value float64, ) *GraphPreviewSample`

NewGraphPreviewSample instantiates a new GraphPreviewSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewSampleWithDefaults

`func NewGraphPreviewSampleWithDefaults() *GraphPreviewSample`

NewGraphPreviewSampleWithDefaults instantiates a new GraphPreviewSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GraphPreviewSample) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GraphPreviewSample) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GraphPreviewSample) SetError(v string)`

SetError sets Error field to given value.


### GetItem

`func (o *GraphPreviewSample) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GraphPreviewSample) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GraphPreviewSample) SetItem(v string)`

SetItem sets Item field to given value.


### GetMissing

`func (o *GraphPreviewSample) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *GraphPreviewSample) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *GraphPreviewSample) SetMissing(v bool)`

SetMissing sets Missing field to given value.


### GetObjects

`func (o *GraphPreviewSample) GetObjects() []GraphPreviewObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GraphPreviewSample) GetObjectsOk() (*[]GraphPreviewObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GraphPreviewSample) SetObjects(v []GraphPreviewObject)`

SetObjects sets Objects field to given value.


### GetValue

`func (o *GraphPreviewSample) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphPreviewSample) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphPreviewSample) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


