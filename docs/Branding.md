# Branding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | Empty when the hostname does not belong to an account. | 
**BackgroundVersion** | **string** | Changes whenever the background is replaced. Empty when there is no background. | 
**HasBackground** | **bool** |  | 
**HasLogo** | **bool** |  | 
**LogoVersion** | **string** | Changes whenever either logo variant is replaced. Empty when there is no logo. | 

## Methods

### NewBranding

`func NewBranding(accountName string, backgroundVersion string, hasBackground bool, hasLogo bool, logoVersion string, ) *Branding`

NewBranding instantiates a new Branding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingWithDefaults

`func NewBrandingWithDefaults() *Branding`

NewBrandingWithDefaults instantiates a new Branding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *Branding) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Branding) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Branding) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetBackgroundVersion

`func (o *Branding) GetBackgroundVersion() string`

GetBackgroundVersion returns the BackgroundVersion field if non-nil, zero value otherwise.

### GetBackgroundVersionOk

`func (o *Branding) GetBackgroundVersionOk() (*string, bool)`

GetBackgroundVersionOk returns a tuple with the BackgroundVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundVersion

`func (o *Branding) SetBackgroundVersion(v string)`

SetBackgroundVersion sets BackgroundVersion field to given value.


### GetHasBackground

`func (o *Branding) GetHasBackground() bool`

GetHasBackground returns the HasBackground field if non-nil, zero value otherwise.

### GetHasBackgroundOk

`func (o *Branding) GetHasBackgroundOk() (*bool, bool)`

GetHasBackgroundOk returns a tuple with the HasBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBackground

`func (o *Branding) SetHasBackground(v bool)`

SetHasBackground sets HasBackground field to given value.


### GetHasLogo

`func (o *Branding) GetHasLogo() bool`

GetHasLogo returns the HasLogo field if non-nil, zero value otherwise.

### GetHasLogoOk

`func (o *Branding) GetHasLogoOk() (*bool, bool)`

GetHasLogoOk returns a tuple with the HasLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLogo

`func (o *Branding) SetHasLogo(v bool)`

SetHasLogo sets HasLogo field to given value.


### GetLogoVersion

`func (o *Branding) GetLogoVersion() string`

GetLogoVersion returns the LogoVersion field if non-nil, zero value otherwise.

### GetLogoVersionOk

`func (o *Branding) GetLogoVersionOk() (*string, bool)`

GetLogoVersionOk returns a tuple with the LogoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoVersion

`func (o *Branding) SetLogoVersion(v string)`

SetLogoVersion sets LogoVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


