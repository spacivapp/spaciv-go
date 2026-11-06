# ProofCompleteSignupReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptMarketing** | Pointer to **bool** |  | [optional] 
**Account** | [**SignupAccount**](SignupAccount.md) |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Proof** | **string** |  | 

## Methods

### NewProofCompleteSignupReqBody

`func NewProofCompleteSignupReqBody(account SignupAccount, firstName string, lastName string, proof string, ) *ProofCompleteSignupReqBody`

NewProofCompleteSignupReqBody instantiates a new ProofCompleteSignupReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofCompleteSignupReqBodyWithDefaults

`func NewProofCompleteSignupReqBodyWithDefaults() *ProofCompleteSignupReqBody`

NewProofCompleteSignupReqBodyWithDefaults instantiates a new ProofCompleteSignupReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptMarketing

`func (o *ProofCompleteSignupReqBody) GetAcceptMarketing() bool`

GetAcceptMarketing returns the AcceptMarketing field if non-nil, zero value otherwise.

### GetAcceptMarketingOk

`func (o *ProofCompleteSignupReqBody) GetAcceptMarketingOk() (*bool, bool)`

GetAcceptMarketingOk returns a tuple with the AcceptMarketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMarketing

`func (o *ProofCompleteSignupReqBody) SetAcceptMarketing(v bool)`

SetAcceptMarketing sets AcceptMarketing field to given value.

### HasAcceptMarketing

`func (o *ProofCompleteSignupReqBody) HasAcceptMarketing() bool`

HasAcceptMarketing returns a boolean if a field has been set.

### GetAccount

`func (o *ProofCompleteSignupReqBody) GetAccount() SignupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProofCompleteSignupReqBody) GetAccountOk() (*SignupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProofCompleteSignupReqBody) SetAccount(v SignupAccount)`

SetAccount sets Account field to given value.


### GetFirstName

`func (o *ProofCompleteSignupReqBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ProofCompleteSignupReqBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ProofCompleteSignupReqBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ProofCompleteSignupReqBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ProofCompleteSignupReqBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ProofCompleteSignupReqBody) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetProof

`func (o *ProofCompleteSignupReqBody) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ProofCompleteSignupReqBody) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ProofCompleteSignupReqBody) SetProof(v string)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


