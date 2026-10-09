# EntryPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of an entry that already exists in the object group. | 
**Name** | Pointer to **string** | Renames the entry. | [optional] 
**Properties** | Pointer to [**map[string]EntryPatchPropertiesValue**](EntryPatchPropertiesValue.md) | The properties to merge into the entry, keyed by property id. Properties the request does not name are left as they are. | [optional] 
**User** | Pointer to **string** | Replaces the entry&#39;s user. | [optional] 

## Methods

### NewEntryPatch

`func NewEntryPatch(id string, ) *EntryPatch`

NewEntryPatch instantiates a new EntryPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryPatchWithDefaults

`func NewEntryPatchWithDefaults() *EntryPatch`

NewEntryPatchWithDefaults instantiates a new EntryPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryPatch) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntryPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *EntryPatch) GetProperties() map[string]EntryPatchPropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EntryPatch) GetPropertiesOk() (*map[string]EntryPatchPropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EntryPatch) SetProperties(v map[string]EntryPatchPropertiesValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EntryPatch) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUser

`func (o *EntryPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EntryPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EntryPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *EntryPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


