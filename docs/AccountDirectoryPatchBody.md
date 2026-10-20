# AccountDirectoryPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTags** | Pointer to **[]string** | Tags to add to the directory. Each is lowercased and must then be at most 32 letters, digits and hyphens, starting with a letter or digit. A directory can have at most 20 tags. | [optional] 
**Name** | Pointer to **string** | New display name for the directory. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash, and must not be . or .. | [optional] 
**ParentId** | Pointer to **string** | Identifier of the directory to move this one into, or an empty string to move it to the root. | [optional] 
**RemoveTags** | Pointer to **[]string** | Tags to remove from the directory. Removing a tag the directory does not have changes nothing. | [optional] 

## Methods

### NewAccountDirectoryPatchBody

`func NewAccountDirectoryPatchBody() *AccountDirectoryPatchBody`

NewAccountDirectoryPatchBody instantiates a new AccountDirectoryPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDirectoryPatchBodyWithDefaults

`func NewAccountDirectoryPatchBodyWithDefaults() *AccountDirectoryPatchBody`

NewAccountDirectoryPatchBodyWithDefaults instantiates a new AccountDirectoryPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTags

`func (o *AccountDirectoryPatchBody) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *AccountDirectoryPatchBody) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *AccountDirectoryPatchBody) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *AccountDirectoryPatchBody) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetName

`func (o *AccountDirectoryPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDirectoryPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDirectoryPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountDirectoryPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *AccountDirectoryPatchBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AccountDirectoryPatchBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AccountDirectoryPatchBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AccountDirectoryPatchBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRemoveTags

`func (o *AccountDirectoryPatchBody) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *AccountDirectoryPatchBody) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *AccountDirectoryPatchBody) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *AccountDirectoryPatchBody) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


