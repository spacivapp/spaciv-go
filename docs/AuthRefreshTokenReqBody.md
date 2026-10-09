# AuthRefreshTokenReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **int64** | Scope the new access token to this account ID. Must be an account the session is authorised for. | [optional] 
**Project** | Pointer to **int64** | Further scope the new access token to this project ID within the given account. | [optional] 

## Methods

### NewAuthRefreshTokenReqBody

`func NewAuthRefreshTokenReqBody() *AuthRefreshTokenReqBody`

NewAuthRefreshTokenReqBody instantiates a new AuthRefreshTokenReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRefreshTokenReqBodyWithDefaults

`func NewAuthRefreshTokenReqBodyWithDefaults() *AuthRefreshTokenReqBody`

NewAuthRefreshTokenReqBodyWithDefaults instantiates a new AuthRefreshTokenReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AuthRefreshTokenReqBody) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AuthRefreshTokenReqBody) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AuthRefreshTokenReqBody) SetAccount(v int64)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AuthRefreshTokenReqBody) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProject

`func (o *AuthRefreshTokenReqBody) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AuthRefreshTokenReqBody) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AuthRefreshTokenReqBody) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *AuthRefreshTokenReqBody) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


