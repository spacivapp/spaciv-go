# EntityChangeActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**Session** | **string** |  | 
**Type** | **string** |  | 
**User** | **int64** |  | 

## Methods

### NewEntityChangeActor

`func NewEntityChangeActor(session string, type_ string, user int64, ) *EntityChangeActor`

NewEntityChangeActor instantiates a new EntityChangeActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityChangeActorWithDefaults

`func NewEntityChangeActorWithDefaults() *EntityChangeActor`

NewEntityChangeActorWithDefaults instantiates a new EntityChangeActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *EntityChangeActor) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EntityChangeActor) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EntityChangeActor) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EntityChangeActor) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSession

`func (o *EntityChangeActor) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *EntityChangeActor) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *EntityChangeActor) SetSession(v string)`

SetSession sets Session field to given value.


### GetType

`func (o *EntityChangeActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityChangeActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityChangeActor) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *EntityChangeActor) GetUser() int64`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EntityChangeActor) GetUserOk() (*int64, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EntityChangeActor) SetUser(v int64)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


