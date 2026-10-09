# I18nTranslationsImportParamsWrapperBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | **string** |  | 
**Translations** | **[][]string** | Rows of key/value pairs. The first row must be a header ([\&quot;key\&quot;,\&quot;value\&quot;]) and is ignored; each subsequent row must contain exactly two strings. | 

## Methods

### NewI18nTranslationsImportParamsWrapperBody

`func NewI18nTranslationsImportParamsWrapperBody(locale string, translations [][]string, ) *I18nTranslationsImportParamsWrapperBody`

NewI18nTranslationsImportParamsWrapperBody instantiates a new I18nTranslationsImportParamsWrapperBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewI18nTranslationsImportParamsWrapperBodyWithDefaults

`func NewI18nTranslationsImportParamsWrapperBodyWithDefaults() *I18nTranslationsImportParamsWrapperBody`

NewI18nTranslationsImportParamsWrapperBodyWithDefaults instantiates a new I18nTranslationsImportParamsWrapperBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *I18nTranslationsImportParamsWrapperBody) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *I18nTranslationsImportParamsWrapperBody) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *I18nTranslationsImportParamsWrapperBody) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetTranslations

`func (o *I18nTranslationsImportParamsWrapperBody) GetTranslations() [][]string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *I18nTranslationsImportParamsWrapperBody) GetTranslationsOk() (*[][]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *I18nTranslationsImportParamsWrapperBody) SetTranslations(v [][]string)`

SetTranslations sets Translations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


