# SignupAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanySize** | Pointer to **string** | One of the values returned by /v1/lookups/company-sizes. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-3 code, as returned by /v1/lookups/countries. | [optional] 
**HouseNumber** | Pointer to **string** |  | [optional] 
**Industry** | Pointer to **string** | One of the values returned by /v1/lookups/industries. | [optional] 
**Name** | **string** | Display name of the new account. Must not already be in use. | 
**Postcode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**Template** | **string** | Starting content for the account. \&quot;blank\&quot; creates an empty account; the others copy in a prebuilt set of projects and settings. | 

## Methods

### NewSignupAccount

`func NewSignupAccount(name string, template string, ) *SignupAccount`

NewSignupAccount instantiates a new SignupAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupAccountWithDefaults

`func NewSignupAccountWithDefaults() *SignupAccount`

NewSignupAccountWithDefaults instantiates a new SignupAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine2

`func (o *SignupAccount) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *SignupAccount) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *SignupAccount) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *SignupAccount) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *SignupAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SignupAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SignupAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SignupAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanySize

`func (o *SignupAccount) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *SignupAccount) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *SignupAccount) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *SignupAccount) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetCountry

`func (o *SignupAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SignupAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SignupAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SignupAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetHouseNumber

`func (o *SignupAccount) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *SignupAccount) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *SignupAccount) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *SignupAccount) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetIndustry

`func (o *SignupAccount) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *SignupAccount) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *SignupAccount) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *SignupAccount) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetName

`func (o *SignupAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignupAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignupAccount) SetName(v string)`

SetName sets Name field to given value.


### GetPostcode

`func (o *SignupAccount) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *SignupAccount) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *SignupAccount) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *SignupAccount) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetState

`func (o *SignupAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SignupAccount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SignupAccount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SignupAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreet

`func (o *SignupAccount) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SignupAccount) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SignupAccount) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SignupAccount) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetTemplate

`func (o *SignupAccount) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SignupAccount) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SignupAccount) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


