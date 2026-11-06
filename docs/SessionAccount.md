# SessionAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** |  | 
**AccountName** | **string** |  | 
**Enterable** | **bool** | A token for this account can be refreshed now, from an existing login or by carrying one of the session&#39;s sign-ins over. When false, the person signs in with one of signInMethods on the account&#39;s sign-in page. | 
**Hostname** | **string** | The account&#39;s sign-in hostname; empty when the account has none. | 
**SignInMethods** | **[]string** | The sign-in methods the account accepts. | 
**SignedIn** | **bool** | The session holds a login to this account. | 

## Methods

### NewSessionAccount

`func NewSessionAccount(accountId int64, accountName string, enterable bool, hostname string, signInMethods []string, signedIn bool, ) *SessionAccount`

NewSessionAccount instantiates a new SessionAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAccountWithDefaults

`func NewSessionAccountWithDefaults() *SessionAccount`

NewSessionAccountWithDefaults instantiates a new SessionAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SessionAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SessionAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SessionAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *SessionAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SessionAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SessionAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetEnterable

`func (o *SessionAccount) GetEnterable() bool`

GetEnterable returns the Enterable field if non-nil, zero value otherwise.

### GetEnterableOk

`func (o *SessionAccount) GetEnterableOk() (*bool, bool)`

GetEnterableOk returns a tuple with the Enterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterable

`func (o *SessionAccount) SetEnterable(v bool)`

SetEnterable sets Enterable field to given value.


### GetHostname

`func (o *SessionAccount) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SessionAccount) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SessionAccount) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSignInMethods

`func (o *SessionAccount) GetSignInMethods() []string`

GetSignInMethods returns the SignInMethods field if non-nil, zero value otherwise.

### GetSignInMethodsOk

`func (o *SessionAccount) GetSignInMethodsOk() (*[]string, bool)`

GetSignInMethodsOk returns a tuple with the SignInMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMethods

`func (o *SessionAccount) SetSignInMethods(v []string)`

SetSignInMethods sets SignInMethods field to given value.


### GetSignedIn

`func (o *SessionAccount) GetSignedIn() bool`

GetSignedIn returns the SignedIn field if non-nil, zero value otherwise.

### GetSignedInOk

`func (o *SessionAccount) GetSignedInOk() (*bool, bool)`

GetSignedInOk returns a tuple with the SignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedIn

`func (o *SessionAccount) SetSignedIn(v bool)`

SetSignedIn sets SignedIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


