# UserV2GetResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**EmployeeId** | **NullableString** |  | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**LocaleCode** | **string** | BCP 47 locale code selected by the user, or an empty string if none has been set. | 
**Store** | **map[string]string** | Arbitrary key-value pairs persisted by the client for UI state (e.g. dismissed banners, feature tour progress). | 
**TermsConditions** | **bool** |  | 
**Title** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 

## Methods

### NewUserV2GetResBody

`func NewUserV2GetResBody(email string, employeeId NullableString, firstName NullableString, lastName NullableString, localeCode string, store map[string]string, termsConditions bool, title NullableString, updatedAt time.Time, userId int64, ) *UserV2GetResBody`

NewUserV2GetResBody instantiates a new UserV2GetResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserV2GetResBodyWithDefaults

`func NewUserV2GetResBodyWithDefaults() *UserV2GetResBody`

NewUserV2GetResBodyWithDefaults instantiates a new UserV2GetResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserV2GetResBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserV2GetResBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserV2GetResBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmployeeId

`func (o *UserV2GetResBody) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *UserV2GetResBody) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *UserV2GetResBody) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### SetEmployeeIdNil

`func (o *UserV2GetResBody) SetEmployeeIdNil(b bool)`

 SetEmployeeIdNil sets the value for EmployeeId to be an explicit nil

### UnsetEmployeeId
`func (o *UserV2GetResBody) UnsetEmployeeId()`

UnsetEmployeeId ensures that no value is present for EmployeeId, not even an explicit nil
### GetFirstName

`func (o *UserV2GetResBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserV2GetResBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserV2GetResBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *UserV2GetResBody) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserV2GetResBody) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserV2GetResBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserV2GetResBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserV2GetResBody) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *UserV2GetResBody) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserV2GetResBody) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLocaleCode

`func (o *UserV2GetResBody) GetLocaleCode() string`

GetLocaleCode returns the LocaleCode field if non-nil, zero value otherwise.

### GetLocaleCodeOk

`func (o *UserV2GetResBody) GetLocaleCodeOk() (*string, bool)`

GetLocaleCodeOk returns a tuple with the LocaleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaleCode

`func (o *UserV2GetResBody) SetLocaleCode(v string)`

SetLocaleCode sets LocaleCode field to given value.


### GetStore

`func (o *UserV2GetResBody) GetStore() map[string]string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *UserV2GetResBody) GetStoreOk() (*map[string]string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *UserV2GetResBody) SetStore(v map[string]string)`

SetStore sets Store field to given value.


### GetTermsConditions

`func (o *UserV2GetResBody) GetTermsConditions() bool`

GetTermsConditions returns the TermsConditions field if non-nil, zero value otherwise.

### GetTermsConditionsOk

`func (o *UserV2GetResBody) GetTermsConditionsOk() (*bool, bool)`

GetTermsConditionsOk returns a tuple with the TermsConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsConditions

`func (o *UserV2GetResBody) SetTermsConditions(v bool)`

SetTermsConditions sets TermsConditions field to given value.


### GetTitle

`func (o *UserV2GetResBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserV2GetResBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserV2GetResBody) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *UserV2GetResBody) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserV2GetResBody) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdatedAt

`func (o *UserV2GetResBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserV2GetResBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserV2GetResBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *UserV2GetResBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserV2GetResBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserV2GetResBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


