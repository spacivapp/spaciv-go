# AccountPutOIDCClientSecretReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** | New client secret of the application registered with the identity provider. Never returned. | 

## Methods

### NewAccountPutOIDCClientSecretReqBody

`func NewAccountPutOIDCClientSecretReqBody(clientSecret string, ) *AccountPutOIDCClientSecretReqBody`

NewAccountPutOIDCClientSecretReqBody instantiates a new AccountPutOIDCClientSecretReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPutOIDCClientSecretReqBodyWithDefaults

`func NewAccountPutOIDCClientSecretReqBodyWithDefaults() *AccountPutOIDCClientSecretReqBody`

NewAccountPutOIDCClientSecretReqBodyWithDefaults instantiates a new AccountPutOIDCClientSecretReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *AccountPutOIDCClientSecretReqBody) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AccountPutOIDCClientSecretReqBody) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AccountPutOIDCClientSecretReqBody) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


