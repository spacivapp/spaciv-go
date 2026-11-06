# ProofCreateResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]SessionAccount**](SessionAccount.md) |  | 
**Proof** | **string** | Proves the person to the other /v1/auth/proof endpoints for 15 minutes. It is not an access token. | 

## Methods

### NewProofCreateResBody

`func NewProofCreateResBody(accounts []SessionAccount, proof string, ) *ProofCreateResBody`

NewProofCreateResBody instantiates a new ProofCreateResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofCreateResBodyWithDefaults

`func NewProofCreateResBodyWithDefaults() *ProofCreateResBody`

NewProofCreateResBodyWithDefaults instantiates a new ProofCreateResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ProofCreateResBody) GetAccounts() []SessionAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ProofCreateResBody) GetAccountsOk() (*[]SessionAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ProofCreateResBody) SetAccounts(v []SessionAccount)`

SetAccounts sets Accounts field to given value.


### GetProof

`func (o *ProofCreateResBody) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ProofCreateResBody) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ProofCreateResBody) SetProof(v string)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


