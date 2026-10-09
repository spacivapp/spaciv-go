# Cookie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Expires** | **time.Time** |  | 
**HttpOnly** | **bool** |  | 
**MaxAge** | **int64** |  | 
**Name** | **string** |  | 
**Partitioned** | **bool** |  | 
**Path** | **string** |  | 
**Quoted** | **bool** |  | 
**Raw** | **string** |  | 
**RawExpires** | **string** |  | 
**SameSite** | **int64** |  | 
**Secure** | **bool** |  | 
**Unparsed** | **[]string** |  | 
**Value** | **string** |  | 

## Methods

### NewCookie

`func NewCookie(domain string, expires time.Time, httpOnly bool, maxAge int64, name string, partitioned bool, path string, quoted bool, raw string, rawExpires string, sameSite int64, secure bool, unparsed []string, value string, ) *Cookie`

NewCookie instantiates a new Cookie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieWithDefaults

`func NewCookieWithDefaults() *Cookie`

NewCookieWithDefaults instantiates a new Cookie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Cookie) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Cookie) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Cookie) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetExpires

`func (o *Cookie) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Cookie) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Cookie) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetHttpOnly

`func (o *Cookie) GetHttpOnly() bool`

GetHttpOnly returns the HttpOnly field if non-nil, zero value otherwise.

### GetHttpOnlyOk

`func (o *Cookie) GetHttpOnlyOk() (*bool, bool)`

GetHttpOnlyOk returns a tuple with the HttpOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOnly

`func (o *Cookie) SetHttpOnly(v bool)`

SetHttpOnly sets HttpOnly field to given value.


### GetMaxAge

`func (o *Cookie) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *Cookie) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *Cookie) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.


### GetName

`func (o *Cookie) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cookie) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cookie) SetName(v string)`

SetName sets Name field to given value.


### GetPartitioned

`func (o *Cookie) GetPartitioned() bool`

GetPartitioned returns the Partitioned field if non-nil, zero value otherwise.

### GetPartitionedOk

`func (o *Cookie) GetPartitionedOk() (*bool, bool)`

GetPartitionedOk returns a tuple with the Partitioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioned

`func (o *Cookie) SetPartitioned(v bool)`

SetPartitioned sets Partitioned field to given value.


### GetPath

`func (o *Cookie) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Cookie) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Cookie) SetPath(v string)`

SetPath sets Path field to given value.


### GetQuoted

`func (o *Cookie) GetQuoted() bool`

GetQuoted returns the Quoted field if non-nil, zero value otherwise.

### GetQuotedOk

`func (o *Cookie) GetQuotedOk() (*bool, bool)`

GetQuotedOk returns a tuple with the Quoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoted

`func (o *Cookie) SetQuoted(v bool)`

SetQuoted sets Quoted field to given value.


### GetRaw

`func (o *Cookie) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *Cookie) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *Cookie) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetRawExpires

`func (o *Cookie) GetRawExpires() string`

GetRawExpires returns the RawExpires field if non-nil, zero value otherwise.

### GetRawExpiresOk

`func (o *Cookie) GetRawExpiresOk() (*string, bool)`

GetRawExpiresOk returns a tuple with the RawExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawExpires

`func (o *Cookie) SetRawExpires(v string)`

SetRawExpires sets RawExpires field to given value.


### GetSameSite

`func (o *Cookie) GetSameSite() int64`

GetSameSite returns the SameSite field if non-nil, zero value otherwise.

### GetSameSiteOk

`func (o *Cookie) GetSameSiteOk() (*int64, bool)`

GetSameSiteOk returns a tuple with the SameSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameSite

`func (o *Cookie) SetSameSite(v int64)`

SetSameSite sets SameSite field to given value.


### GetSecure

`func (o *Cookie) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *Cookie) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *Cookie) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetUnparsed

`func (o *Cookie) GetUnparsed() []string`

GetUnparsed returns the Unparsed field if non-nil, zero value otherwise.

### GetUnparsedOk

`func (o *Cookie) GetUnparsedOk() (*[]string, bool)`

GetUnparsedOk returns a tuple with the Unparsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparsed

`func (o *Cookie) SetUnparsed(v []string)`

SetUnparsed sets Unparsed field to given value.


### GetValue

`func (o *Cookie) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Cookie) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Cookie) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


