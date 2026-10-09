# ImportBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**EmployeeID** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Role** | **string** | Role name to assign to this user within the account (e.g. \&quot;Account-Owner\&quot;, \&quot;Contributor\&quot;). Must be a valid role returned by GET /v1/user-roles. | 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewImportBody

`func NewImportBody(email string, firstName string, lastName string, role string, ) *ImportBody`

NewImportBody instantiates a new ImportBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBodyWithDefaults

`func NewImportBodyWithDefaults() *ImportBody`

NewImportBodyWithDefaults instantiates a new ImportBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ImportBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ImportBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ImportBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmployeeID

`func (o *ImportBody) GetEmployeeID() string`

GetEmployeeID returns the EmployeeID field if non-nil, zero value otherwise.

### GetEmployeeIDOk

`func (o *ImportBody) GetEmployeeIDOk() (*string, bool)`

GetEmployeeIDOk returns a tuple with the EmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeID

`func (o *ImportBody) SetEmployeeID(v string)`

SetEmployeeID sets EmployeeID field to given value.

### HasEmployeeID

`func (o *ImportBody) HasEmployeeID() bool`

HasEmployeeID returns a boolean if a field has been set.

### GetFirstName

`func (o *ImportBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ImportBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ImportBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ImportBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ImportBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ImportBody) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetRole

`func (o *ImportBody) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ImportBody) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ImportBody) SetRole(v string)`

SetRole sets Role field to given value.


### GetTitle

`func (o *ImportBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ImportBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


