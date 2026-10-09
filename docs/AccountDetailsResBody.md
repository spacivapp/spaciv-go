# AccountDetailsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** |  | 
**AccountSlug** | **string** | URL-safe identifier for the account. | 
**AddressLine2** | **NullableString** |  | 
**Availability** | **string** |  | 
**City** | **NullableString** |  | 
**CompanyLegalName** | **NullableString** |  | 
**CompanySize** | **string** | Headcount band the account selected, as listed by /v1/lookups/company-sizes. | 
**Country** | **NullableString** |  | 
**Hostname** | **string** | Subdomain the account signs in on. | 
**HouseNumber** | **NullableString** |  | 
**Industry** | **string** |  | 
**Integrations** | **[]string** | Third-party integrations currently connected, for example \&quot;slack\&quot;. | 
**Name** | **string** | Display name of the account. | 
**NotificationPreferences** | **[]string** | Notification categories the account has opted into. | 
**Postcode** | **NullableString** |  | 
**Product** | **string** | Product the account is subscribed to, as listed by /v1/lookups/products. | 
**State** | **NullableString** |  | 
**Street** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**VatNumber** | **NullableString** |  | 

## Methods

### NewAccountDetailsResBody

`func NewAccountDetailsResBody(accountId int64, accountSlug string, addressLine2 NullableString, availability string, city NullableString, companyLegalName NullableString, companySize string, country NullableString, hostname string, houseNumber NullableString, industry string, integrations []string, name string, notificationPreferences []string, postcode NullableString, product string, state NullableString, street NullableString, updatedAt time.Time, vatNumber NullableString, ) *AccountDetailsResBody`

NewAccountDetailsResBody instantiates a new AccountDetailsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsResBodyWithDefaults

`func NewAccountDetailsResBodyWithDefaults() *AccountDetailsResBody`

NewAccountDetailsResBodyWithDefaults instantiates a new AccountDetailsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountDetailsResBody) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountDetailsResBody) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountDetailsResBody) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAccountSlug

`func (o *AccountDetailsResBody) GetAccountSlug() string`

GetAccountSlug returns the AccountSlug field if non-nil, zero value otherwise.

### GetAccountSlugOk

`func (o *AccountDetailsResBody) GetAccountSlugOk() (*string, bool)`

GetAccountSlugOk returns a tuple with the AccountSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlug

`func (o *AccountDetailsResBody) SetAccountSlug(v string)`

SetAccountSlug sets AccountSlug field to given value.


### GetAddressLine2

`func (o *AccountDetailsResBody) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AccountDetailsResBody) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AccountDetailsResBody) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.


### SetAddressLine2Nil

`func (o *AccountDetailsResBody) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *AccountDetailsResBody) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAvailability

`func (o *AccountDetailsResBody) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AccountDetailsResBody) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AccountDetailsResBody) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetCity

`func (o *AccountDetailsResBody) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AccountDetailsResBody) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AccountDetailsResBody) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *AccountDetailsResBody) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *AccountDetailsResBody) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCompanyLegalName

`func (o *AccountDetailsResBody) GetCompanyLegalName() string`

GetCompanyLegalName returns the CompanyLegalName field if non-nil, zero value otherwise.

### GetCompanyLegalNameOk

`func (o *AccountDetailsResBody) GetCompanyLegalNameOk() (*string, bool)`

GetCompanyLegalNameOk returns a tuple with the CompanyLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegalName

`func (o *AccountDetailsResBody) SetCompanyLegalName(v string)`

SetCompanyLegalName sets CompanyLegalName field to given value.


### SetCompanyLegalNameNil

`func (o *AccountDetailsResBody) SetCompanyLegalNameNil(b bool)`

 SetCompanyLegalNameNil sets the value for CompanyLegalName to be an explicit nil

### UnsetCompanyLegalName
`func (o *AccountDetailsResBody) UnsetCompanyLegalName()`

UnsetCompanyLegalName ensures that no value is present for CompanyLegalName, not even an explicit nil
### GetCompanySize

`func (o *AccountDetailsResBody) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountDetailsResBody) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountDetailsResBody) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.


### GetCountry

`func (o *AccountDetailsResBody) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AccountDetailsResBody) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AccountDetailsResBody) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *AccountDetailsResBody) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AccountDetailsResBody) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetHostname

`func (o *AccountDetailsResBody) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AccountDetailsResBody) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AccountDetailsResBody) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHouseNumber

`func (o *AccountDetailsResBody) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AccountDetailsResBody) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AccountDetailsResBody) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.


### SetHouseNumberNil

`func (o *AccountDetailsResBody) SetHouseNumberNil(b bool)`

 SetHouseNumberNil sets the value for HouseNumber to be an explicit nil

### UnsetHouseNumber
`func (o *AccountDetailsResBody) UnsetHouseNumber()`

UnsetHouseNumber ensures that no value is present for HouseNumber, not even an explicit nil
### GetIndustry

`func (o *AccountDetailsResBody) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *AccountDetailsResBody) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *AccountDetailsResBody) SetIndustry(v string)`

SetIndustry sets Industry field to given value.


### GetIntegrations

`func (o *AccountDetailsResBody) GetIntegrations() []string`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *AccountDetailsResBody) GetIntegrationsOk() (*[]string, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *AccountDetailsResBody) SetIntegrations(v []string)`

SetIntegrations sets Integrations field to given value.


### GetName

`func (o *AccountDetailsResBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDetailsResBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDetailsResBody) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationPreferences

`func (o *AccountDetailsResBody) GetNotificationPreferences() []string`

GetNotificationPreferences returns the NotificationPreferences field if non-nil, zero value otherwise.

### GetNotificationPreferencesOk

`func (o *AccountDetailsResBody) GetNotificationPreferencesOk() (*[]string, bool)`

GetNotificationPreferencesOk returns a tuple with the NotificationPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPreferences

`func (o *AccountDetailsResBody) SetNotificationPreferences(v []string)`

SetNotificationPreferences sets NotificationPreferences field to given value.


### GetPostcode

`func (o *AccountDetailsResBody) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AccountDetailsResBody) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AccountDetailsResBody) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### SetPostcodeNil

`func (o *AccountDetailsResBody) SetPostcodeNil(b bool)`

 SetPostcodeNil sets the value for Postcode to be an explicit nil

### UnsetPostcode
`func (o *AccountDetailsResBody) UnsetPostcode()`

UnsetPostcode ensures that no value is present for Postcode, not even an explicit nil
### GetProduct

`func (o *AccountDetailsResBody) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AccountDetailsResBody) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AccountDetailsResBody) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetState

`func (o *AccountDetailsResBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountDetailsResBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountDetailsResBody) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *AccountDetailsResBody) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AccountDetailsResBody) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *AccountDetailsResBody) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AccountDetailsResBody) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AccountDetailsResBody) SetStreet(v string)`

SetStreet sets Street field to given value.


### SetStreetNil

`func (o *AccountDetailsResBody) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *AccountDetailsResBody) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetUpdatedAt

`func (o *AccountDetailsResBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountDetailsResBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountDetailsResBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVatNumber

`func (o *AccountDetailsResBody) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *AccountDetailsResBody) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *AccountDetailsResBody) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.


### SetVatNumberNil

`func (o *AccountDetailsResBody) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *AccountDetailsResBody) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


