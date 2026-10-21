# ClientAuthMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Type** | **string** |  | 
**V** | **int32** |  | 

## Methods

### NewClientAuthMessage

`func NewClientAuthMessage(token string, type_ string, v int32, ) *ClientAuthMessage`

NewClientAuthMessage instantiates a new ClientAuthMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthMessageWithDefaults

`func NewClientAuthMessageWithDefaults() *ClientAuthMessage`

NewClientAuthMessageWithDefaults instantiates a new ClientAuthMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ClientAuthMessage) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientAuthMessage) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientAuthMessage) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *ClientAuthMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientAuthMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientAuthMessage) SetType(v string)`

SetType sets Type field to given value.


### GetV

`func (o *ClientAuthMessage) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ClientAuthMessage) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ClientAuthMessage) SetV(v int32)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


