# ScriptAuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorType** | **string** |  | 
**Kind** | **string** | What changed, for example created, updated, api-key-rotated or needs-reauthorisation. | 
**ModifiedAt** | **time.Time** |  | 
**ModifiedBy** | Pointer to **int64** | User who made the change, when actorType is user. | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScriptAuditEntry

`func NewScriptAuditEntry(actorType string, kind string, modifiedAt time.Time, ) *ScriptAuditEntry`

NewScriptAuditEntry instantiates a new ScriptAuditEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptAuditEntryWithDefaults

`func NewScriptAuditEntryWithDefaults() *ScriptAuditEntry`

NewScriptAuditEntryWithDefaults instantiates a new ScriptAuditEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorType

`func (o *ScriptAuditEntry) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *ScriptAuditEntry) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *ScriptAuditEntry) SetActorType(v string)`

SetActorType sets ActorType field to given value.


### GetKind

`func (o *ScriptAuditEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScriptAuditEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScriptAuditEntry) SetKind(v string)`

SetKind sets Kind field to given value.


### GetModifiedAt

`func (o *ScriptAuditEntry) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ScriptAuditEntry) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ScriptAuditEntry) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ScriptAuditEntry) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ScriptAuditEntry) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ScriptAuditEntry) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ScriptAuditEntry) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetParams

`func (o *ScriptAuditEntry) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScriptAuditEntry) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScriptAuditEntry) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ScriptAuditEntry) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


