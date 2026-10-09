# AccountFilePatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | New display name for the file. Must not contain a forward slash or a backslash. | 

## Methods

### NewAccountFilePatchBody

`func NewAccountFilePatchBody(name string, ) *AccountFilePatchBody`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


