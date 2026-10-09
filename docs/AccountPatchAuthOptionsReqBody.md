# AccountPatchAuthOptionsReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthOptions** | [**[]Item**](Item.md) | Sign-in methods to change. Methods left out keep their current setting. | 

## Methods

### NewAccountPatchAuthOptionsReqBody

`func NewAccountPatchAuthOptionsReqBody(authOptions []Item, ) *AccountPatchAuthOptionsReqBody`

NewAccountPatchAuthOptionsReqBody instantiates a new AccountPatchAuthOptionsReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPatchAuthOptionsReqBodyWithDefaults

`func NewAccountPatchAuthOptionsReqBodyWithDefaults() *AccountPatchAuthOptionsReqBody`

NewAccountPatchAuthOptionsReqBodyWithDefaults instantiates a new AccountPatchAuthOptionsReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthOptions

`func (o *AccountPatchAuthOptionsReqBody) GetAuthOptions() []Item`

GetAuthOptions returns the AuthOptions field if non-nil, zero value otherwise.

### GetAuthOptionsOk

`func (o *AccountPatchAuthOptionsReqBody) GetAuthOptionsOk() (*[]Item, bool)`

GetAuthOptionsOk returns a tuple with the AuthOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOptions

`func (o *AccountPatchAuthOptionsReqBody) SetAuthOptions(v []Item)`

SetAuthOptions sets AuthOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


