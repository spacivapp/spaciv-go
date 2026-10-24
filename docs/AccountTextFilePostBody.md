# AccountTextFilePostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new file. Its extension chooses its type. At most 255 bytes after normalising to Unicode NFC. Must not contain a forward slash or a backslash, and must not be . or .. | 
**ParentId** | Pointer to **string** | Identifier of the directory to create the file in. Leave it and parentPath out to create it at the root. | [optional] 
**ParentPath** | Pointer to **string** | Path of the directory to create the file in, such as Leases/2026. Cannot be given with parentId. | [optional] 

## Methods

### NewAccountTextFilePostBody

`func NewAccountTextFilePostBody(name string, ) *AccountTextFilePostBody`

NewAccountTextFilePostBody instantiates a new AccountTextFilePostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTextFilePostBodyWithDefaults

`func NewAccountTextFilePostBodyWithDefaults() *AccountTextFilePostBody`

NewAccountTextFilePostBodyWithDefaults instantiates a new AccountTextFilePostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountTextFilePostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTextFilePostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTextFilePostBody) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *AccountTextFilePostBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AccountTextFilePostBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AccountTextFilePostBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AccountTextFilePostBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentPath

`func (o *AccountTextFilePostBody) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *AccountTextFilePostBody) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *AccountTextFilePostBody) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *AccountTextFilePostBody) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


