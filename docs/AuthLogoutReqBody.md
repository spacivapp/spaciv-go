# AuthLogoutReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **int64** | Sign out of this account only, keeping the session&#39;s logins to other accounts. Omit to sign out of every account. | [optional] 

## Methods

### NewAuthLogoutReqBody

`func NewAuthLogoutReqBody() *AuthLogoutReqBody`

NewAuthLogoutReqBody instantiates a new AuthLogoutReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLogoutReqBodyWithDefaults

`func NewAuthLogoutReqBodyWithDefaults() *AuthLogoutReqBody`

NewAuthLogoutReqBodyWithDefaults instantiates a new AuthLogoutReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AuthLogoutReqBody) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AuthLogoutReqBody) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AuthLogoutReqBody) SetAccount(v int64)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AuthLogoutReqBody) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


