# UserV2PatchReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptMarketing** | Pointer to **bool** |  | [optional] 
**EmployeeID** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LocaleCode** | Pointer to **string** | BCP 47 locale code to set for the authenticated user. Pass an empty string to clear the preference. | [optional] 
**Role** | Pointer to **string** | Role name to assign to the target user. Omit (with roleUser or roleUserEmail present) to revoke the user&#39;s current role. | [optional] 
**RoleUser** | Pointer to **int64** | ID of the user whose role should be updated. Requires the *Manage users* permission. | [optional] 
**RoleUserEmail** | Pointer to **string** | Email of the user whose role should be updated. If no account exists for this address, an invitation is sent. Requires the *Manage users* permission. | [optional] 
**Store** | Pointer to **map[string]string** | Arbitrary key-value pairs persisted for client UI state. Merged with any existing entries. | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewUserV2PatchReqBody

`func NewUserV2PatchReqBody() *UserV2PatchReqBody`

NewUserV2PatchReqBody instantiates a new UserV2PatchReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserV2PatchReqBodyWithDefaults

`func NewUserV2PatchReqBodyWithDefaults() *UserV2PatchReqBody`

NewUserV2PatchReqBodyWithDefaults instantiates a new UserV2PatchReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptMarketing

`func (o *UserV2PatchReqBody) GetAcceptMarketing() bool`

GetAcceptMarketing returns the AcceptMarketing field if non-nil, zero value otherwise.

### GetAcceptMarketingOk

`func (o *UserV2PatchReqBody) GetAcceptMarketingOk() (*bool, bool)`

GetAcceptMarketingOk returns a tuple with the AcceptMarketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMarketing

`func (o *UserV2PatchReqBody) SetAcceptMarketing(v bool)`

SetAcceptMarketing sets AcceptMarketing field to given value.

### HasAcceptMarketing

`func (o *UserV2PatchReqBody) HasAcceptMarketing() bool`

HasAcceptMarketing returns a boolean if a field has been set.

### GetEmployeeID

`func (o *UserV2PatchReqBody) GetEmployeeID() string`

GetEmployeeID returns the EmployeeID field if non-nil, zero value otherwise.

### GetEmployeeIDOk

`func (o *UserV2PatchReqBody) GetEmployeeIDOk() (*string, bool)`

GetEmployeeIDOk returns a tuple with the EmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeID

`func (o *UserV2PatchReqBody) SetEmployeeID(v string)`

SetEmployeeID sets EmployeeID field to given value.

### HasEmployeeID

`func (o *UserV2PatchReqBody) HasEmployeeID() bool`

HasEmployeeID returns a boolean if a field has been set.

### GetFirstName

`func (o *UserV2PatchReqBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserV2PatchReqBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserV2PatchReqBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserV2PatchReqBody) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserV2PatchReqBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserV2PatchReqBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserV2PatchReqBody) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserV2PatchReqBody) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLocaleCode

`func (o *UserV2PatchReqBody) GetLocaleCode() string`

GetLocaleCode returns the LocaleCode field if non-nil, zero value otherwise.

### GetLocaleCodeOk

`func (o *UserV2PatchReqBody) GetLocaleCodeOk() (*string, bool)`

GetLocaleCodeOk returns a tuple with the LocaleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaleCode

`func (o *UserV2PatchReqBody) SetLocaleCode(v string)`

SetLocaleCode sets LocaleCode field to given value.

### HasLocaleCode

`func (o *UserV2PatchReqBody) HasLocaleCode() bool`

HasLocaleCode returns a boolean if a field has been set.

### GetRole

`func (o *UserV2PatchReqBody) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserV2PatchReqBody) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserV2PatchReqBody) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserV2PatchReqBody) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleUser

`func (o *UserV2PatchReqBody) GetRoleUser() int64`

GetRoleUser returns the RoleUser field if non-nil, zero value otherwise.

### GetRoleUserOk

`func (o *UserV2PatchReqBody) GetRoleUserOk() (*int64, bool)`

GetRoleUserOk returns a tuple with the RoleUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUser

`func (o *UserV2PatchReqBody) SetRoleUser(v int64)`

SetRoleUser sets RoleUser field to given value.

### HasRoleUser

`func (o *UserV2PatchReqBody) HasRoleUser() bool`

HasRoleUser returns a boolean if a field has been set.

### GetRoleUserEmail

`func (o *UserV2PatchReqBody) GetRoleUserEmail() string`

GetRoleUserEmail returns the RoleUserEmail field if non-nil, zero value otherwise.

### GetRoleUserEmailOk

`func (o *UserV2PatchReqBody) GetRoleUserEmailOk() (*string, bool)`

GetRoleUserEmailOk returns a tuple with the RoleUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUserEmail

`func (o *UserV2PatchReqBody) SetRoleUserEmail(v string)`

SetRoleUserEmail sets RoleUserEmail field to given value.

### HasRoleUserEmail

`func (o *UserV2PatchReqBody) HasRoleUserEmail() bool`

HasRoleUserEmail returns a boolean if a field has been set.

### GetStore

`func (o *UserV2PatchReqBody) GetStore() map[string]string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *UserV2PatchReqBody) GetStoreOk() (*map[string]string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *UserV2PatchReqBody) SetStore(v map[string]string)`

SetStore sets Store field to given value.

### HasStore

`func (o *UserV2PatchReqBody) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTitle

`func (o *UserV2PatchReqBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserV2PatchReqBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserV2PatchReqBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserV2PatchReqBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


