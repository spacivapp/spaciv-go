# TopDownModelEntryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Objects** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Properties** | Pointer to [**TopDownModelProperties**](TopDownModelProperties.md) |  | [optional] 

## Methods

### NewTopDownModelEntryUpdate

`func NewTopDownModelEntryUpdate(id string, objects []FilterGroup, ) *TopDownModelEntryUpdate`

NewTopDownModelEntryUpdate instantiates a new TopDownModelEntryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopDownModelEntryUpdateWithDefaults

`func NewTopDownModelEntryUpdateWithDefaults() *TopDownModelEntryUpdate`

NewTopDownModelEntryUpdateWithDefaults instantiates a new TopDownModelEntryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopDownModelEntryUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopDownModelEntryUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopDownModelEntryUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetObjects

`func (o *TopDownModelEntryUpdate) GetObjects() []FilterGroup`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *TopDownModelEntryUpdate) GetObjectsOk() (*[]FilterGroup, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *TopDownModelEntryUpdate) SetObjects(v []FilterGroup)`

SetObjects sets Objects field to given value.


### GetProperties

`func (o *TopDownModelEntryUpdate) GetProperties() TopDownModelProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TopDownModelEntryUpdate) GetPropertiesOk() (*TopDownModelProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TopDownModelEntryUpdate) SetProperties(v TopDownModelProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TopDownModelEntryUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


