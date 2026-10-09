# ObjectGroupEntryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearProperties** | Pointer to **bool** |  | [optional] 
**Delete** | **bool** |  | 
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NewId** | Pointer to **string** |  | [optional] 
**Properties** | [**map[string]ObjectGroupEntryPropertyNodes**](ObjectGroupEntryPropertyNodes.md) |  | 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectGroupEntryUpdate

`func NewObjectGroupEntryUpdate(delete bool, id string, properties map[string]ObjectGroupEntryPropertyNodes, ) *ObjectGroupEntryUpdate`

NewObjectGroupEntryUpdate instantiates a new ObjectGroupEntryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupEntryUpdateWithDefaults

`func NewObjectGroupEntryUpdateWithDefaults() *ObjectGroupEntryUpdate`

NewObjectGroupEntryUpdateWithDefaults instantiates a new ObjectGroupEntryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearProperties

`func (o *ObjectGroupEntryUpdate) GetClearProperties() bool`

GetClearProperties returns the ClearProperties field if non-nil, zero value otherwise.

### GetClearPropertiesOk

`func (o *ObjectGroupEntryUpdate) GetClearPropertiesOk() (*bool, bool)`

GetClearPropertiesOk returns a tuple with the ClearProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearProperties

`func (o *ObjectGroupEntryUpdate) SetClearProperties(v bool)`

SetClearProperties sets ClearProperties field to given value.

### HasClearProperties

`func (o *ObjectGroupEntryUpdate) HasClearProperties() bool`

HasClearProperties returns a boolean if a field has been set.

### GetDelete

`func (o *ObjectGroupEntryUpdate) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ObjectGroupEntryUpdate) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ObjectGroupEntryUpdate) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetId

`func (o *ObjectGroupEntryUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectGroupEntryUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectGroupEntryUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ObjectGroupEntryUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectGroupEntryUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectGroupEntryUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectGroupEntryUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewId

`func (o *ObjectGroupEntryUpdate) GetNewId() string`

GetNewId returns the NewId field if non-nil, zero value otherwise.

### GetNewIdOk

`func (o *ObjectGroupEntryUpdate) GetNewIdOk() (*string, bool)`

GetNewIdOk returns a tuple with the NewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewId

`func (o *ObjectGroupEntryUpdate) SetNewId(v string)`

SetNewId sets NewId field to given value.

### HasNewId

`func (o *ObjectGroupEntryUpdate) HasNewId() bool`

HasNewId returns a boolean if a field has been set.

### GetProperties

`func (o *ObjectGroupEntryUpdate) GetProperties() map[string]ObjectGroupEntryPropertyNodes`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectGroupEntryUpdate) GetPropertiesOk() (*map[string]ObjectGroupEntryPropertyNodes, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectGroupEntryUpdate) SetProperties(v map[string]ObjectGroupEntryPropertyNodes)`

SetProperties sets Properties field to given value.


### GetUser

`func (o *ObjectGroupEntryUpdate) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectGroupEntryUpdate) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectGroupEntryUpdate) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ObjectGroupEntryUpdate) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


