# GetEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorType** | **string** |  | 
**Entity** | **string** |  | 
**EntityId** | **string** |  | 
**Kind** | **string** |  | 
**ModifiedAt** | **time.Time** |  | 
**ModifiedBy** | **int64** |  | 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Session** | Pointer to **string** |  | [optional] 
**To** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGetEntry

`func NewGetEntry(actorType string, entity string, entityId string, kind string, modifiedAt time.Time, modifiedBy int64, ) *GetEntry`

NewGetEntry instantiates a new GetEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntryWithDefaults

`func NewGetEntryWithDefaults() *GetEntry`

NewGetEntryWithDefaults instantiates a new GetEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorType

`func (o *GetEntry) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *GetEntry) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *GetEntry) SetActorType(v string)`

SetActorType sets ActorType field to given value.


### GetEntity

`func (o *GetEntry) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetEntry) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetEntry) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEntityId

`func (o *GetEntry) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GetEntry) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GetEntry) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetKind

`func (o *GetEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetEntry) SetKind(v string)`

SetKind sets Kind field to given value.


### GetModifiedAt

`func (o *GetEntry) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GetEntry) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GetEntry) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *GetEntry) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GetEntry) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GetEntry) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.


### GetParams

`func (o *GetEntry) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetEntry) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetEntry) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *GetEntry) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSession

`func (o *GetEntry) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetEntry) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetEntry) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *GetEntry) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTo

`func (o *GetEntry) GetTo() map[string]string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetEntry) GetToOk() (*map[string]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetEntry) SetTo(v map[string]string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetEntry) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


