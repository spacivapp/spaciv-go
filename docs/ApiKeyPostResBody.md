# ApiKeyPostResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | [**APIKeyDetails**](APIKeyDetails.md) |  | 
**Key** | **string** | The secret, returned only once | 

## Methods

### NewApiKeyPostResBody

`func NewApiKeyPostResBody(apiKey APIKeyDetails, key string, ) *ApiKeyPostResBody`

NewApiKeyPostResBody instantiates a new ApiKeyPostResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyPostResBodyWithDefaults

`func NewApiKeyPostResBodyWithDefaults() *ApiKeyPostResBody`

NewApiKeyPostResBodyWithDefaults instantiates a new ApiKeyPostResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiKeyPostResBody) GetApiKey() APIKeyDetails`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiKeyPostResBody) GetApiKeyOk() (*APIKeyDetails, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiKeyPostResBody) SetApiKey(v APIKeyDetails)`

SetApiKey sets ApiKey field to given value.


### GetKey

`func (o *ApiKeyPostResBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKeyPostResBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKeyPostResBody) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


