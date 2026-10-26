# ScriptSecretsReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remove** | Pointer to **[]string** | Names of secrets to delete. | [optional] 
**Revision** | Pointer to **int64** | The revision last read. When given, a different current revision is rejected with 409 and nothing is set, and a change to the script&#39;s code or settings while the secrets are being set is reported with 409 after they are set; either way, read the script again before changing anything else. Omit it to set the secrets whatever the revision. | [optional] 
**Set** | Pointer to **map[string]string** | Secrets to add or replace, keyed by environment variable name. | [optional] 

## Methods

### NewScriptSecretsReqBody

`func NewScriptSecretsReqBody() *ScriptSecretsReqBody`

NewScriptSecretsReqBody instantiates a new ScriptSecretsReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptSecretsReqBodyWithDefaults

`func NewScriptSecretsReqBodyWithDefaults() *ScriptSecretsReqBody`

NewScriptSecretsReqBodyWithDefaults instantiates a new ScriptSecretsReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemove

`func (o *ScriptSecretsReqBody) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *ScriptSecretsReqBody) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *ScriptSecretsReqBody) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *ScriptSecretsReqBody) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptSecretsReqBody) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptSecretsReqBody) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptSecretsReqBody) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ScriptSecretsReqBody) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSet

`func (o *ScriptSecretsReqBody) GetSet() map[string]string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ScriptSecretsReqBody) GetSetOk() (*map[string]string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ScriptSecretsReqBody) SetSet(v map[string]string)`

SetSet sets Set field to given value.

### HasSet

`func (o *ScriptSecretsReqBody) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


