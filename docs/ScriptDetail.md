# ScriptDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Files** | **map[string]string** | Source files keyed by path. | 
**Id** | **string** |  | 
**Key** | [**ScriptKey**](ScriptKey.md) |  | 
**Language** | **string** |  | 
**Manifest** | Pointer to **string** | Dependency manifest: requirements.txt for Python, go.mod for Go, package.json for JavaScript. | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to **int64** | Project the script&#39;s API key is scoped to. Absent when it acts across the account. | [optional] 
**Revision** | **int64** | Pass back when updating, so a concurrent change is reported rather than overwritten. | 
**Runtime** | Pointer to [**ScriptRuntime**](ScriptRuntime.md) | Live state from the runner; absent if the runner could not be reached. | [optional] 
**Schedule** | [**ScriptSchedule**](ScriptSchedule.md) |  | 
**SecretNames** | **[]string** | Names of the secrets set on the script. Values are never returned. | 
**Type** | **string** | The handler signature the script implements. | 
**UpdatedAt** | **time.Time** |  | 
**Webhook** | **bool** |  | 
**WebhookUrl** | Pointer to **string** | Returned only when a webhook URL is first issued; it cannot be retrieved later. | [optional] 

## Methods

### NewScriptDetail

`func NewScriptDetail(createdAt time.Time, createdBy int64, enabled bool, files map[string]string, id string, key ScriptKey, language string, name string, revision int64, schedule ScriptSchedule, secretNames []string, type_ string, updatedAt time.Time, webhook bool, ) *ScriptDetail`

NewScriptDetail instantiates a new ScriptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptDetailWithDefaults

`func NewScriptDetailWithDefaults() *ScriptDetail`

NewScriptDetailWithDefaults instantiates a new ScriptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScriptDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScriptDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScriptDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ScriptDetail) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScriptDetail) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScriptDetail) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *ScriptDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptDetail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptDetail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptDetail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFiles

`func (o *ScriptDetail) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScriptDetail) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScriptDetail) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.


### GetId

`func (o *ScriptDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptDetail) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ScriptDetail) GetKey() ScriptKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScriptDetail) GetKeyOk() (*ScriptKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScriptDetail) SetKey(v ScriptKey)`

SetKey sets Key field to given value.


### GetLanguage

`func (o *ScriptDetail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptDetail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptDetail) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetManifest

`func (o *ScriptDetail) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ScriptDetail) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ScriptDetail) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ScriptDetail) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *ScriptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptDetail) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ScriptDetail) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptDetail) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptDetail) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptDetail) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptDetail) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptDetail) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetRuntime

`func (o *ScriptDetail) GetRuntime() ScriptRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ScriptDetail) GetRuntimeOk() (*ScriptRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ScriptDetail) SetRuntime(v ScriptRuntime)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ScriptDetail) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetSchedule

`func (o *ScriptDetail) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptDetail) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptDetail) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.


### GetSecretNames

`func (o *ScriptDetail) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ScriptDetail) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ScriptDetail) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetType

`func (o *ScriptDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptDetail) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ScriptDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScriptDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScriptDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWebhook

`func (o *ScriptDetail) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptDetail) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptDetail) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.


### GetWebhookUrl

`func (o *ScriptDetail) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ScriptDetail) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ScriptDetail) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ScriptDetail) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


