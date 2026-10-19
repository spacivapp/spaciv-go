# EntityChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **int64** |  | 
**Actor** | [**EntityChangeActor**](EntityChangeActor.md) |  | 
**At** | **string** |  | 
**Id** | **string** |  | 
**Op** | **string** |  | 
**Project** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityChangedEvent

`func NewEntityChangedEvent(account int64, actor EntityChangeActor, at string, id string, op string, type_ string, ) *EntityChangedEvent`

NewEntityChangedEvent instantiates a new EntityChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityChangedEventWithDefaults

`func NewEntityChangedEventWithDefaults() *EntityChangedEvent`

NewEntityChangedEventWithDefaults instantiates a new EntityChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *EntityChangedEvent) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *EntityChangedEvent) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *EntityChangedEvent) SetAccount(v int64)`

SetAccount sets Account field to given value.


### GetActor

`func (o *EntityChangedEvent) GetActor() EntityChangeActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EntityChangedEvent) GetActorOk() (*EntityChangeActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EntityChangedEvent) SetActor(v EntityChangeActor)`

SetActor sets Actor field to given value.


### GetAt

`func (o *EntityChangedEvent) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *EntityChangedEvent) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *EntityChangedEvent) SetAt(v string)`

SetAt sets At field to given value.


### GetId

`func (o *EntityChangedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityChangedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityChangedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *EntityChangedEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EntityChangedEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EntityChangedEvent) SetOp(v string)`

SetOp sets Op field to given value.


### GetProject

`func (o *EntityChangedEvent) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EntityChangedEvent) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EntityChangedEvent) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *EntityChangedEvent) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetType

`func (o *EntityChangedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityChangedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityChangedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *EntityChangedEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EntityChangedEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EntityChangedEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EntityChangedEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


