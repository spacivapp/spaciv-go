# AccountPostOIDCProviderReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID of the application registered with the identity provider. | 
**ClientSecret** | **string** | Client secret of the application registered with the identity provider. Never returned. | 
**ProviderName** | **string** | Identity provider product the account signs in through. | 
**Tenant** | **string** | Organisation at the identity provider: a Microsoft Entra tenant ID or verified domain, or the PingFederate domain. | 

## Methods

### NewAccountPostOIDCProviderReqBody

`func NewAccountPostOIDCProviderReqBody(clientId string, clientSecret string, providerName string, tenant string, ) *AccountPostOIDCProviderReqBody`

NewAccountPostOIDCProviderReqBody instantiates a new AccountPostOIDCProviderReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPostOIDCProviderReqBodyWithDefaults

`func NewAccountPostOIDCProviderReqBodyWithDefaults() *AccountPostOIDCProviderReqBody`

NewAccountPostOIDCProviderReqBodyWithDefaults instantiates a new AccountPostOIDCProviderReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountPostOIDCProviderReqBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountPostOIDCProviderReqBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountPostOIDCProviderReqBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AccountPostOIDCProviderReqBody) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AccountPostOIDCProviderReqBody) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AccountPostOIDCProviderReqBody) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetProviderName

`func (o *AccountPostOIDCProviderReqBody) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AccountPostOIDCProviderReqBody) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AccountPostOIDCProviderReqBody) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetTenant

`func (o *AccountPostOIDCProviderReqBody) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountPostOIDCProviderReqBody) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountPostOIDCProviderReqBody) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


