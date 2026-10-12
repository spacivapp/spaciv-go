# ScriptKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** | Why the last automatic rotation failed or why the key needs re-authorising. | [optional] 
**OwnerUserId** | Pointer to **int64** | User whose permissions the key acts with. | [optional] 
**RotateAt** | Pointer to **time.Time** | When the key is next replaced automatically. | [optional] 
**Status** | **string** | Whether the script&#39;s Spaciv API key is usable. While \&quot;needs_reauthorisation\&quot;, runs are paused until someone with *Manage my account* re-authorises it. | 

## Methods

### NewScriptKey

`func NewScriptKey(status string, ) *ScriptKey`

NewScriptKey instantiates a new ScriptKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptKeyWithDefaults

`func NewScriptKeyWithDefaults() *ScriptKey`

NewScriptKeyWithDefaults instantiates a new ScriptKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ScriptKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ScriptKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ScriptKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ScriptKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLastError

`func (o *ScriptKey) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ScriptKey) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ScriptKey) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ScriptKey) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ScriptKey) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ScriptKey) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ScriptKey) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ScriptKey) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetRotateAt

`func (o *ScriptKey) GetRotateAt() time.Time`

GetRotateAt returns the RotateAt field if non-nil, zero value otherwise.

### GetRotateAtOk

`func (o *ScriptKey) GetRotateAtOk() (*time.Time, bool)`

GetRotateAtOk returns a tuple with the RotateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAt

`func (o *ScriptKey) SetRotateAt(v time.Time)`

SetRotateAt sets RotateAt field to given value.

### HasRotateAt

`func (o *ScriptKey) HasRotateAt() bool`

HasRotateAt returns a boolean if a field has been set.

### GetStatus

`func (o *ScriptKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScriptKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScriptKey) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


