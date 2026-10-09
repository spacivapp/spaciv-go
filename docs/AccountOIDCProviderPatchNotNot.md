# AccountOIDCProviderPatchNotNot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Identity provider product the account signs in through. | 
**ClientId** | **string** | Client ID of the application registered with the identity provider. | 
**Tenant** | **string** | Organisation at the identity provider: a Microsoft Entra tenant ID or verified domain, or the PingFederate domain. | 

## Methods

### NewAccountOIDCProviderPatchNotNot

`func NewAccountOIDCProviderPatchNotNot(providerName string, clientId string, tenant string, ) *AccountOIDCProviderPatchNotNot`

NewAccountOIDCProviderPatchNotNot instantiates a new AccountOIDCProviderPatchNotNot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOIDCProviderPatchNotNotWithDefaults

`func NewAccountOIDCProviderPatchNotNotWithDefaults() *AccountOIDCProviderPatchNotNot`

NewAccountOIDCProviderPatchNotNotWithDefaults instantiates a new AccountOIDCProviderPatchNotNot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AccountOIDCProviderPatchNotNot) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AccountOIDCProviderPatchNotNot) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AccountOIDCProviderPatchNotNot) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetClientId

`func (o *AccountOIDCProviderPatchNotNot) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountOIDCProviderPatchNotNot) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountOIDCProviderPatchNotNot) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetTenant

`func (o *AccountOIDCProviderPatchNotNot) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountOIDCProviderPatchNotNot) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountOIDCProviderPatchNotNot) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


