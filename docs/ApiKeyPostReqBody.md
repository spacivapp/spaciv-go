# ApiKeyPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresInDays** | **int64** |  | 
**Name** | **string** |  | 
**Project** | **int64** | Restrict the key to this project ID. Must belong to the caller&#39;s account, and the key can reach no other project. | 

## Methods

### NewApiKeyPostReqBody

`func NewApiKeyPostReqBody(expiresInDays int64, name string, project int64, ) *ApiKeyPostReqBody`

NewApiKeyPostReqBody instantiates a new ApiKeyPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyPostReqBodyWithDefaults

`func NewApiKeyPostReqBodyWithDefaults() *ApiKeyPostReqBody`

NewApiKeyPostReqBodyWithDefaults instantiates a new ApiKeyPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresInDays

`func (o *ApiKeyPostReqBody) GetExpiresInDays() int64`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *ApiKeyPostReqBody) GetExpiresInDaysOk() (*int64, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *ApiKeyPostReqBody) SetExpiresInDays(v int64)`

SetExpiresInDays sets ExpiresInDays field to given value.


### GetName

`func (o *ApiKeyPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyPostReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ApiKeyPostReqBody) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ApiKeyPostReqBody) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ApiKeyPostReqBody) SetProject(v int64)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


