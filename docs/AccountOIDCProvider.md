# AccountOIDCProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID of the application registered with the identity provider. | 
**ClientSecretSet** | **bool** | Whether a client secret is stored. The secret itself is never returned. | 
**CreatedAt** | **time.Time** |  | 
**OidcProviderId** | **int64** |  | 
**ProviderName** | **string** | Identity provider product the account signs in through. | 
**Tenant** | **string** | Organisation at the identity provider: a Microsoft Entra tenant ID or verified domain, or the PingFederate domain. | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAccountOIDCProvider

`func NewAccountOIDCProvider(clientId string, clientSecretSet bool, createdAt time.Time, oidcProviderId int64, providerName string, tenant string, updatedAt time.Time, ) *AccountOIDCProvider`

NewAccountOIDCProvider instantiates a new AccountOIDCProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOIDCProviderWithDefaults

`func NewAccountOIDCProviderWithDefaults() *AccountOIDCProvider`

NewAccountOIDCProviderWithDefaults instantiates a new AccountOIDCProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountOIDCProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountOIDCProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountOIDCProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecretSet

`func (o *AccountOIDCProvider) GetClientSecretSet() bool`

GetClientSecretSet returns the ClientSecretSet field if non-nil, zero value otherwise.

### GetClientSecretSetOk

`func (o *AccountOIDCProvider) GetClientSecretSetOk() (*bool, bool)`

GetClientSecretSetOk returns a tuple with the ClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretSet

`func (o *AccountOIDCProvider) SetClientSecretSet(v bool)`

SetClientSecretSet sets ClientSecretSet field to given value.


### GetCreatedAt

`func (o *AccountOIDCProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountOIDCProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountOIDCProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetOidcProviderId

`func (o *AccountOIDCProvider) GetOidcProviderId() int64`

GetOidcProviderId returns the OidcProviderId field if non-nil, zero value otherwise.

### GetOidcProviderIdOk

`func (o *AccountOIDCProvider) GetOidcProviderIdOk() (*int64, bool)`

GetOidcProviderIdOk returns a tuple with the OidcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProviderId

`func (o *AccountOIDCProvider) SetOidcProviderId(v int64)`

SetOidcProviderId sets OidcProviderId field to given value.


### GetProviderName

`func (o *AccountOIDCProvider) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AccountOIDCProvider) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AccountOIDCProvider) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetTenant

`func (o *AccountOIDCProvider) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountOIDCProvider) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountOIDCProvider) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdatedAt

`func (o *AccountOIDCProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountOIDCProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountOIDCProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


