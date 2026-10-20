# AccountDirectoryPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the directory. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash, and must not be . or .. | 
**ParentId** | Pointer to **string** | Identifier of the directory to create this one in. Leave it out to create it at the root. | [optional] 

## Methods

### NewAccountDirectoryPostBody

`func NewAccountDirectoryPostBody(name string, ) *AccountDirectoryPostBody`

NewAccountDirectoryPostBody instantiates a new AccountDirectoryPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDirectoryPostBodyWithDefaults

`func NewAccountDirectoryPostBodyWithDefaults() *AccountDirectoryPostBody`

NewAccountDirectoryPostBodyWithDefaults instantiates a new AccountDirectoryPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountDirectoryPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDirectoryPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDirectoryPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *AccountDirectoryPostBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AccountDirectoryPostBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AccountDirectoryPostBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AccountDirectoryPostBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


