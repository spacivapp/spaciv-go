# GraphPreviewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Properties** | [**[]GraphPreviewProperty**](GraphPreviewProperty.md) |  | 

## Methods

### NewGraphPreviewObject

`func NewGraphPreviewObject(id string, name string, properties []GraphPreviewProperty, ) *GraphPreviewObject`

NewGraphPreviewObject instantiates a new GraphPreviewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewObjectWithDefaults

`func NewGraphPreviewObjectWithDefaults() *GraphPreviewObject`

NewGraphPreviewObjectWithDefaults instantiates a new GraphPreviewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphPreviewObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphPreviewObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphPreviewObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GraphPreviewObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphPreviewObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphPreviewObject) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *GraphPreviewObject) GetProperties() []GraphPreviewProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GraphPreviewObject) GetPropertiesOk() (*[]GraphPreviewProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GraphPreviewObject) SetProperties(v []GraphPreviewProperty)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


