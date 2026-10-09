# ListUserDetailsWithEmailRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptMarketing** | **bool** |  | 
**AccountId** | **int64** |  | 
**Email** | **string** |  | 
**EmployeeId** | **NullableString** |  | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**RoleName** | **string** |  | 
**Status** | **string** |  | 
**Store** | **interface{}** |  | 
**TermsConditions** | **bool** |  | 
**Title** | **NullableString** |  | 
**UserId** | **int64** |  | 

## Methods

### NewListUserDetailsWithEmailRow

`func NewListUserDetailsWithEmailRow(acceptMarketing bool, accountId int64, email string, employeeId NullableString, firstName NullableString, lastName NullableString, roleName string, status string, store interface{}, termsConditions bool, title NullableString, userId int64, ) *ListUserDetailsWithEmailRow`

NewListUserDetailsWithEmailRow instantiates a new ListUserDetailsWithEmailRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserDetailsWithEmailRowWithDefaults

`func NewListUserDetailsWithEmailRowWithDefaults() *ListUserDetailsWithEmailRow`

NewListUserDetailsWithEmailRowWithDefaults instantiates a new ListUserDetailsWithEmailRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptMarketing

`func (o *ListUserDetailsWithEmailRow) GetAcceptMarketing() bool`

GetAcceptMarketing returns the AcceptMarketing field if non-nil, zero value otherwise.

### GetAcceptMarketingOk

`func (o *ListUserDetailsWithEmailRow) GetAcceptMarketingOk() (*bool, bool)`

GetAcceptMarketingOk returns a tuple with the AcceptMarketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMarketing

`func (o *ListUserDetailsWithEmailRow) SetAcceptMarketing(v bool)`

SetAcceptMarketing sets AcceptMarketing field to given value.


### GetAccountId

`func (o *ListUserDetailsWithEmailRow) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListUserDetailsWithEmailRow) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListUserDetailsWithEmailRow) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetEmail

`func (o *ListUserDetailsWithEmailRow) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListUserDetailsWithEmailRow) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListUserDetailsWithEmailRow) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmployeeId

`func (o *ListUserDetailsWithEmailRow) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *ListUserDetailsWithEmailRow) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *ListUserDetailsWithEmailRow) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### SetEmployeeIdNil

`func (o *ListUserDetailsWithEmailRow) SetEmployeeIdNil(b bool)`

 SetEmployeeIdNil sets the value for EmployeeId to be an explicit nil

### UnsetEmployeeId
`func (o *ListUserDetailsWithEmailRow) UnsetEmployeeId()`

UnsetEmployeeId ensures that no value is present for EmployeeId, not even an explicit nil
### GetFirstName

`func (o *ListUserDetailsWithEmailRow) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ListUserDetailsWithEmailRow) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ListUserDetailsWithEmailRow) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *ListUserDetailsWithEmailRow) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ListUserDetailsWithEmailRow) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ListUserDetailsWithEmailRow) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ListUserDetailsWithEmailRow) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ListUserDetailsWithEmailRow) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *ListUserDetailsWithEmailRow) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ListUserDetailsWithEmailRow) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetRoleName

`func (o *ListUserDetailsWithEmailRow) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ListUserDetailsWithEmailRow) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ListUserDetailsWithEmailRow) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetStatus

`func (o *ListUserDetailsWithEmailRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListUserDetailsWithEmailRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListUserDetailsWithEmailRow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStore

`func (o *ListUserDetailsWithEmailRow) GetStore() interface{}`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ListUserDetailsWithEmailRow) GetStoreOk() (*interface{}, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ListUserDetailsWithEmailRow) SetStore(v interface{})`

SetStore sets Store field to given value.


### SetStoreNil

`func (o *ListUserDetailsWithEmailRow) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ListUserDetailsWithEmailRow) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetTermsConditions

`func (o *ListUserDetailsWithEmailRow) GetTermsConditions() bool`

GetTermsConditions returns the TermsConditions field if non-nil, zero value otherwise.

### GetTermsConditionsOk

`func (o *ListUserDetailsWithEmailRow) GetTermsConditionsOk() (*bool, bool)`

GetTermsConditionsOk returns a tuple with the TermsConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsConditions

`func (o *ListUserDetailsWithEmailRow) SetTermsConditions(v bool)`

SetTermsConditions sets TermsConditions field to given value.


### GetTitle

`func (o *ListUserDetailsWithEmailRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListUserDetailsWithEmailRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListUserDetailsWithEmailRow) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ListUserDetailsWithEmailRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListUserDetailsWithEmailRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *ListUserDetailsWithEmailRow) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListUserDetailsWithEmailRow) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListUserDetailsWithEmailRow) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


