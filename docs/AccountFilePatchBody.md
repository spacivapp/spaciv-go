# AccountFilePatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New display name for the file. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash. | [optional] 
**ParentId** | Pointer to **string** | Identifier of the directory to move the file into, or an empty string to move it to the root. | [optional] 

## Methods

### NewAccountFilePatchBody

`func NewAccountFilePatchBody() *AccountFilePatchBody`

NewAccountFilePatchBody instantiates a new AccountFilePatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFilePatchBodyWithDefaults

`func NewAccountFilePatchBodyWithDefaults() *AccountFilePatchBody`

NewAccountFilePatchBodyWithDefaults instantiates a new AccountFilePatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountFilePatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountFilePatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountFilePatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountFilePatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *AccountFilePatchBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AccountFilePatchBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AccountFilePatchBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AccountFilePatchBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


