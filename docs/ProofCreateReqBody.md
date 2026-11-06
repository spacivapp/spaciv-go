# ProofCreateReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A magic link token for the person, from a signup, invitation or sign-in email. | 

## Methods

### NewProofCreateReqBody

`func NewProofCreateReqBody(token string, ) *ProofCreateReqBody`

NewProofCreateReqBody instantiates a new ProofCreateReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofCreateReqBodyWithDefaults

`func NewProofCreateReqBodyWithDefaults() *ProofCreateReqBody`

NewProofCreateReqBodyWithDefaults instantiates a new ProofCreateReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ProofCreateReqBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProofCreateReqBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProofCreateReqBody) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


