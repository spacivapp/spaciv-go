# AccountFileByPathPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTags** | Pointer to **[]string** | Tags to add. Each is lowercased and must then be at most 32 letters, digits and hyphens, starting with a letter or digit. A file or directory can have at most 20 tags. | [optional] 
**Name** | Pointer to **string** | New name for the file or directory. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash, and must not be . or .. | [optional] 
**ParentPath** | Pointer to **string** | Path of the directory to move the file or directory into, or / to move it to the root. | [optional] 
**RemoveTags** | Pointer to **[]string** | Tags to remove. Removing a tag the item does not have changes nothing. | [optional] 

## Methods

### NewAccountFileByPathPatchBody

`func NewAccountFileByPathPatchBody() *AccountFileByPathPatchBody`

NewAccountFileByPathPatchBody instantiates a new AccountFileByPathPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFileByPathPatchBodyWithDefaults

`func NewAccountFileByPathPatchBodyWithDefaults() *AccountFileByPathPatchBody`

NewAccountFileByPathPatchBodyWithDefaults instantiates a new AccountFileByPathPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTags

`func (o *AccountFileByPathPatchBody) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *AccountFileByPathPatchBody) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *AccountFileByPathPatchBody) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *AccountFileByPathPatchBody) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetName

`func (o *AccountFileByPathPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountFileByPathPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountFileByPathPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountFileByPathPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentPath

`func (o *AccountFileByPathPatchBody) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *AccountFileByPathPatchBody) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *AccountFileByPathPatchBody) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *AccountFileByPathPatchBody) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.

### GetRemoveTags

`func (o *AccountFileByPathPatchBody) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *AccountFileByPathPatchBody) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *AccountFileByPathPatchBody) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *AccountFileByPathPatchBody) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


