# AccountAuthOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** |  | 
**AuthMethod** | **string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewAccountAuthOption

`func NewAccountAuthOption(accountId int64, authMethod string, enabled bool, ) *AccountAuthOption`

NewAccountAuthOption instantiates a new AccountAuthOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAuthOptionWithDefaults

`func NewAccountAuthOptionWithDefaults() *AccountAuthOption`

NewAccountAuthOptionWithDefaults instantiates a new AccountAuthOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountAuthOption) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountAuthOption) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountAuthOption) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAuthMethod

`func (o *AccountAuthOption) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AccountAuthOption) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AccountAuthOption) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetEnabled

`func (o *AccountAuthOption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountAuthOption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountAuthOption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


