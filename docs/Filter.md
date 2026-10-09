# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **[]string** |  | 
**ObjectGroup** | **string** |  | 
**Property** | **string** |  | 
**PropertyNode** | **[]string** |  | 
**PropertyVar** | Pointer to **string** |  | [optional] 
**Set** | **int32** |  | 

## Methods

### NewFilter

`func NewFilter(object []string, objectGroup string, property string, propertyNode []string, set int32, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Filter) GetObject() []string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Filter) GetObjectOk() (*[]string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Filter) SetObject(v []string)`

SetObject sets Object field to given value.


### GetObjectGroup

`func (o *Filter) GetObjectGroup() string`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *Filter) GetObjectGroupOk() (*string, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *Filter) SetObjectGroup(v string)`

SetObjectGroup sets ObjectGroup field to given value.


### GetProperty

`func (o *Filter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Filter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Filter) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetPropertyNode

`func (o *Filter) GetPropertyNode() []string`

GetPropertyNode returns the PropertyNode field if non-nil, zero value otherwise.

### GetPropertyNodeOk

`func (o *Filter) GetPropertyNodeOk() (*[]string, bool)`

GetPropertyNodeOk returns a tuple with the PropertyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNode

`func (o *Filter) SetPropertyNode(v []string)`

SetPropertyNode sets PropertyNode field to given value.


### GetPropertyVar

`func (o *Filter) GetPropertyVar() string`

GetPropertyVar returns the PropertyVar field if non-nil, zero value otherwise.

### GetPropertyVarOk

`func (o *Filter) GetPropertyVarOk() (*string, bool)`

GetPropertyVarOk returns a tuple with the PropertyVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyVar

`func (o *Filter) SetPropertyVar(v string)`

SetPropertyVar sets PropertyVar field to given value.

### HasPropertyVar

`func (o *Filter) HasPropertyVar() bool`

HasPropertyVar returns a boolean if a field has been set.

### GetSet

`func (o *Filter) GetSet() int32`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *Filter) GetSetOk() (*int32, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *Filter) SetSet(v int32)`

SetSet sets Set field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


