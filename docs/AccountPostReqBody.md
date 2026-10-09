# AccountPostReqBody

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

### NewAccountPostReqBody

`func NewAccountPostReqBody(name string, template string, ) *AccountPostReqBody`

NewAccountPostReqBody instantiates a new AccountPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPostReqBodyWithDefaults

`func NewAccountPostReqBodyWithDefaults() *AccountPostReqBody`

NewAccountPostReqBodyWithDefaults instantiates a new AccountPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine2

`func (o *AccountPostReqBody) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AccountPostReqBody) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AccountPostReqBody) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AccountPostReqBody) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *AccountPostReqBody) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AccountPostReqBody) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AccountPostReqBody) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AccountPostReqBody) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountPostReqBody) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountPostReqBody) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountPostReqBody) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountPostReqBody) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetCountry

`func (o *AccountPostReqBody) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AccountPostReqBody) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AccountPostReqBody) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AccountPostReqBody) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AccountPostReqBody) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AccountPostReqBody) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AccountPostReqBody) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AccountPostReqBody) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetIndustry

`func (o *AccountPostReqBody) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *AccountPostReqBody) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *AccountPostReqBody) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *AccountPostReqBody) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetName

`func (o *AccountPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountPostReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetPostcode

`func (o *AccountPostReqBody) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AccountPostReqBody) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AccountPostReqBody) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *AccountPostReqBody) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetState

`func (o *AccountPostReqBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountPostReqBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountPostReqBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccountPostReqBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreet

`func (o *AccountPostReqBody) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AccountPostReqBody) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AccountPostReqBody) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AccountPostReqBody) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetTemplate

`func (o *AccountPostReqBody) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AccountPostReqBody) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AccountPostReqBody) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


