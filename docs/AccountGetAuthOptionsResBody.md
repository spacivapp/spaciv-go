# AccountGetAuthOptionsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthOptions** | [**[]AccountAuthOption**](AccountAuthOption.md) | Sign-in methods configured for the account, enabled or not. | 
**SsoConfigured** | **bool** | Whether the account has exactly one OIDC provider, without which single sign-on cannot be enabled. | 

## Methods

### NewAccountGetAuthOptionsResBody

`func NewAccountGetAuthOptionsResBody(authOptions []AccountAuthOption, ssoConfigured bool, ) *AccountGetAuthOptionsResBody`

NewAccountGetAuthOptionsResBody instantiates a new AccountGetAuthOptionsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGetAuthOptionsResBodyWithDefaults

`func NewAccountGetAuthOptionsResBodyWithDefaults() *AccountGetAuthOptionsResBody`

NewAccountGetAuthOptionsResBodyWithDefaults instantiates a new AccountGetAuthOptionsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthOptions

`func (o *AccountGetAuthOptionsResBody) GetAuthOptions() []AccountAuthOption`

GetAuthOptions returns the AuthOptions field if non-nil, zero value otherwise.

### GetAuthOptionsOk

`func (o *AccountGetAuthOptionsResBody) GetAuthOptionsOk() (*[]AccountAuthOption, bool)`

GetAuthOptionsOk returns a tuple with the AuthOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOptions

`func (o *AccountGetAuthOptionsResBody) SetAuthOptions(v []AccountAuthOption)`

SetAuthOptions sets AuthOptions field to given value.


### GetSsoConfigured

`func (o *AccountGetAuthOptionsResBody) GetSsoConfigured() bool`

GetSsoConfigured returns the SsoConfigured field if non-nil, zero value otherwise.

### GetSsoConfiguredOk

`func (o *AccountGetAuthOptionsResBody) GetSsoConfiguredOk() (*bool, bool)`

GetSsoConfiguredOk returns a tuple with the SsoConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoConfigured

`func (o *AccountGetAuthOptionsResBody) SetSsoConfigured(v bool)`

SetSsoConfigured sets SsoConfigured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


