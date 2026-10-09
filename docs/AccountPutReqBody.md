# AccountPutReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**AccountSlug** | Pointer to **string** | URL-safe identifier for the account. | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **string** | One of the values returned by /v1/lookups/availabilities. | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanyLegalName** | Pointer to **string** | Registered company name, used on invoices. | [optional] 
**CompanySize** | Pointer to **string** | One of the values returned by /v1/lookups/company-sizes. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-3 code, as returned by /v1/lookups/countries. | [optional] 
**Hostname** | Pointer to **string** | Subdomain members sign in on. Must be unique across all accounts. | [optional] 
**HouseNumber** | Pointer to **string** |  | [optional] 
**Industry** | Pointer to **string** | One of the values returned by /v1/lookups/industries. | [optional] 
**NotificationPreferences** | Pointer to **[]string** | Notification categories to opt into. Replaces the existing set. | [optional] 
**Postcode** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** | One of the values returned by /v1/lookups/products. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountPutReqBody

`func NewAccountPutReqBody() *AccountPutReqBody`

NewAccountPutReqBody instantiates a new AccountPutReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPutReqBodyWithDefaults

`func NewAccountPutReqBodyWithDefaults() *AccountPutReqBody`

NewAccountPutReqBodyWithDefaults instantiates a new AccountPutReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *AccountPutReqBody) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountPutReqBody) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountPutReqBody) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AccountPutReqBody) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountSlug

`func (o *AccountPutReqBody) GetAccountSlug() string`

GetAccountSlug returns the AccountSlug field if non-nil, zero value otherwise.

### GetAccountSlugOk

`func (o *AccountPutReqBody) GetAccountSlugOk() (*string, bool)`

GetAccountSlugOk returns a tuple with the AccountSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlug

`func (o *AccountPutReqBody) SetAccountSlug(v string)`

SetAccountSlug sets AccountSlug field to given value.

### HasAccountSlug

`func (o *AccountPutReqBody) HasAccountSlug() bool`

HasAccountSlug returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AccountPutReqBody) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AccountPutReqBody) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AccountPutReqBody) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AccountPutReqBody) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAvailability

`func (o *AccountPutReqBody) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AccountPutReqBody) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AccountPutReqBody) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *AccountPutReqBody) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCity

`func (o *AccountPutReqBody) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AccountPutReqBody) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AccountPutReqBody) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AccountPutReqBody) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyLegalName

`func (o *AccountPutReqBody) GetCompanyLegalName() string`

GetCompanyLegalName returns the CompanyLegalName field if non-nil, zero value otherwise.

### GetCompanyLegalNameOk

`func (o *AccountPutReqBody) GetCompanyLegalNameOk() (*string, bool)`

GetCompanyLegalNameOk returns a tuple with the CompanyLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegalName

`func (o *AccountPutReqBody) SetCompanyLegalName(v string)`

SetCompanyLegalName sets CompanyLegalName field to given value.

### HasCompanyLegalName

`func (o *AccountPutReqBody) HasCompanyLegalName() bool`

HasCompanyLegalName returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountPutReqBody) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountPutReqBody) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountPutReqBody) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountPutReqBody) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetCountry

`func (o *AccountPutReqBody) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AccountPutReqBody) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AccountPutReqBody) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AccountPutReqBody) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetHostname

`func (o *AccountPutReqBody) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AccountPutReqBody) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AccountPutReqBody) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AccountPutReqBody) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AccountPutReqBody) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AccountPutReqBody) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AccountPutReqBody) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AccountPutReqBody) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetIndustry

`func (o *AccountPutReqBody) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *AccountPutReqBody) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *AccountPutReqBody) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *AccountPutReqBody) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetNotificationPreferences

`func (o *AccountPutReqBody) GetNotificationPreferences() []string`

GetNotificationPreferences returns the NotificationPreferences field if non-nil, zero value otherwise.

### GetNotificationPreferencesOk

`func (o *AccountPutReqBody) GetNotificationPreferencesOk() (*[]string, bool)`

GetNotificationPreferencesOk returns a tuple with the NotificationPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPreferences

`func (o *AccountPutReqBody) SetNotificationPreferences(v []string)`

SetNotificationPreferences sets NotificationPreferences field to given value.

### HasNotificationPreferences

`func (o *AccountPutReqBody) HasNotificationPreferences() bool`

HasNotificationPreferences returns a boolean if a field has been set.

### GetPostcode

`func (o *AccountPutReqBody) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AccountPutReqBody) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AccountPutReqBody) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *AccountPutReqBody) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetProduct

`func (o *AccountPutReqBody) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AccountPutReqBody) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AccountPutReqBody) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AccountPutReqBody) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetState

`func (o *AccountPutReqBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountPutReqBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountPutReqBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccountPutReqBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreet

`func (o *AccountPutReqBody) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AccountPutReqBody) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AccountPutReqBody) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AccountPutReqBody) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetVatNumber

`func (o *AccountPutReqBody) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *AccountPutReqBody) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *AccountPutReqBody) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *AccountPutReqBody) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


