# AccountDirectoryPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New display name for the directory. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash, and must not be . or .. | [optional] 
**ParentId** | Pointer to **string** | Identifier of the directory to move this one into, or an empty string to move it to the root. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


