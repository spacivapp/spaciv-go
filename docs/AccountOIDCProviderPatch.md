# AccountOIDCProviderPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID of the application registered with the identity provider. | [optional] 
**ProviderName** | Pointer to **string** | Identity provider product the account signs in through. | [optional] 
**Tenant** | Pointer to **string** | Organisation at the identity provider: a Microsoft Entra tenant ID or verified domain, or the PingFederate domain. | [optional] 

## Methods

### NewAccountOIDCProviderPatch

`func NewAccountOIDCProviderPatch() *AccountOIDCProviderPatch`

NewAccountOIDCProviderPatch instantiates a new AccountOIDCProviderPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOIDCProviderPatchWithDefaults

`func NewAccountOIDCProviderPatchWithDefaults() *AccountOIDCProviderPatch`

NewAccountOIDCProviderPatchWithDefaults instantiates a new AccountOIDCProviderPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountOIDCProviderPatch) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountOIDCProviderPatch) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountOIDCProviderPatch) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AccountOIDCProviderPatch) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProviderName

`func (o *AccountOIDCProviderPatch) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AccountOIDCProviderPatch) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AccountOIDCProviderPatch) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AccountOIDCProviderPatch) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetTenant

`func (o *AccountOIDCProviderPatch) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountOIDCProviderPatch) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountOIDCProviderPatch) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountOIDCProviderPatch) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


