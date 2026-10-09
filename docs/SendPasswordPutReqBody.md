# SendPasswordPutReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Hostname** | Pointer to **string** | Hostname the reset was requested from. The emailed link returns the user to that account&#39;s sign-in page, and is omitted unless the user belongs to that account. | [optional] 

## Methods

### NewSendPasswordPutReqBody

`func NewSendPasswordPutReqBody(email string, ) *SendPasswordPutReqBody`

NewSendPasswordPutReqBody instantiates a new SendPasswordPutReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPasswordPutReqBodyWithDefaults

`func NewSendPasswordPutReqBodyWithDefaults() *SendPasswordPutReqBody`

NewSendPasswordPutReqBodyWithDefaults instantiates a new SendPasswordPutReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendPasswordPutReqBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendPasswordPutReqBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendPasswordPutReqBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHostname

`func (o *SendPasswordPutReqBody) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SendPasswordPutReqBody) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SendPasswordPutReqBody) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SendPasswordPutReqBody) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


