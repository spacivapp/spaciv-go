# ScriptSecretsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**Key** | [**ScriptKey**](ScriptKey.md) |  | 
**Language** | **string** |  | 
**Name** | **string** |  | 
**Project** | Pointer to **int64** | Project the script&#39;s API key is scoped to. Absent when it acts across the account. | [optional] 
**Revision** | **int64** | Pass back when updating, so a concurrent change is reported rather than overwritten. | 
**Schedule** | [**ScriptSchedule**](ScriptSchedule.md) |  | 
**SecretNames** | **[]string** |  | 
**Type** | **string** | The handler signature the script implements. | 
**UpdatedAt** | **time.Time** |  | 
**Webhook** | **bool** |  | 

## Methods

### NewScriptSecretsResBody

`func NewScriptSecretsResBody(createdAt time.Time, createdBy int64, enabled bool, id string, key ScriptKey, language string, name string, revision int64, schedule ScriptSchedule, secretNames []string, type_ string, updatedAt time.Time, webhook bool, ) *ScriptSecretsResBody`

NewScriptSecretsResBody instantiates a new ScriptSecretsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptSecretsResBodyWithDefaults

`func NewScriptSecretsResBodyWithDefaults() *ScriptSecretsResBody`

NewScriptSecretsResBodyWithDefaults instantiates a new ScriptSecretsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScriptSecretsResBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScriptSecretsResBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScriptSecretsResBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ScriptSecretsResBody) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScriptSecretsResBody) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScriptSecretsResBody) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *ScriptSecretsResBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptSecretsResBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptSecretsResBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptSecretsResBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptSecretsResBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptSecretsResBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptSecretsResBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ScriptSecretsResBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptSecretsResBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptSecretsResBody) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ScriptSecretsResBody) GetKey() ScriptKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScriptSecretsResBody) GetKeyOk() (*ScriptKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScriptSecretsResBody) SetKey(v ScriptKey)`

SetKey sets Key field to given value.


### GetLanguage

`func (o *ScriptSecretsResBody) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptSecretsResBody) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptSecretsResBody) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetName

`func (o *ScriptSecretsResBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptSecretsResBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptSecretsResBody) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ScriptSecretsResBody) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptSecretsResBody) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptSecretsResBody) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptSecretsResBody) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptSecretsResBody) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptSecretsResBody) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptSecretsResBody) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetSchedule

`func (o *ScriptSecretsResBody) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptSecretsResBody) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptSecretsResBody) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.


### GetSecretNames

`func (o *ScriptSecretsResBody) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ScriptSecretsResBody) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ScriptSecretsResBody) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetType

`func (o *ScriptSecretsResBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptSecretsResBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptSecretsResBody) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ScriptSecretsResBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScriptSecretsResBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScriptSecretsResBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWebhook

`func (o *ScriptSecretsResBody) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptSecretsResBody) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptSecretsResBody) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


