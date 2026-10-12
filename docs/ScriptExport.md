# ScriptExport

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
**Runs** | [**[]ScriptExportRun**](ScriptExportRun.md) | Every build and run still kept, newest first. | 
**RunsError** | Pointer to **string** | Set when the run history could not be read in full, in which case runs holds only the newest part that was read. Page through the script&#39;s runs endpoint for the rest. | [optional] 
**Runtime** | Pointer to [**ScriptRuntime**](ScriptRuntime.md) | Live state from the runner; absent if the runner could not be reached. | [optional] 
**Schedule** | [**ScriptSchedule**](ScriptSchedule.md) |  | 
**SecretNames** | **[]string** | Names of the secrets set on the script. Values are never returned. | 
**Type** | **string** | The handler signature the script implements. | 
**UpdatedAt** | **time.Time** |  | 
**Webhook** | **bool** |  | 
**WebhookUrl** | Pointer to **string** | Returned only when a webhook URL is first issued; it cannot be retrieved later. | [optional] 

## Methods

### NewScriptExport

`func NewScriptExport(createdAt time.Time, createdBy int64, enabled bool, files map[string]string, id string, key ScriptKey, language string, name string, revision int64, runs []ScriptExportRun, schedule ScriptSchedule, secretNames []string, type_ string, updatedAt time.Time, webhook bool, ) *ScriptExport`

NewScriptExport instantiates a new ScriptExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExportWithDefaults

`func NewScriptExportWithDefaults() *ScriptExport`

NewScriptExportWithDefaults instantiates a new ScriptExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScriptExport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScriptExport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScriptExport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ScriptExport) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScriptExport) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScriptExport) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *ScriptExport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptExport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptExport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptExport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptExport) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptExport) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptExport) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFiles

`func (o *ScriptExport) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScriptExport) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScriptExport) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.


### GetId

`func (o *ScriptExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptExport) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ScriptExport) GetKey() ScriptKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScriptExport) GetKeyOk() (*ScriptKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScriptExport) SetKey(v ScriptKey)`

SetKey sets Key field to given value.


### GetLanguage

`func (o *ScriptExport) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptExport) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptExport) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetManifest

`func (o *ScriptExport) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ScriptExport) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ScriptExport) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ScriptExport) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *ScriptExport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptExport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptExport) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ScriptExport) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptExport) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptExport) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptExport) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptExport) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptExport) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptExport) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetRuns

`func (o *ScriptExport) GetRuns() []ScriptExportRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ScriptExport) GetRunsOk() (*[]ScriptExportRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ScriptExport) SetRuns(v []ScriptExportRun)`

SetRuns sets Runs field to given value.


### GetRunsError

`func (o *ScriptExport) GetRunsError() string`

GetRunsError returns the RunsError field if non-nil, zero value otherwise.

### GetRunsErrorOk

`func (o *ScriptExport) GetRunsErrorOk() (*string, bool)`

GetRunsErrorOk returns a tuple with the RunsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsError

`func (o *ScriptExport) SetRunsError(v string)`

SetRunsError sets RunsError field to given value.

### HasRunsError

`func (o *ScriptExport) HasRunsError() bool`

HasRunsError returns a boolean if a field has been set.

### GetRuntime

`func (o *ScriptExport) GetRuntime() ScriptRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ScriptExport) GetRuntimeOk() (*ScriptRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ScriptExport) SetRuntime(v ScriptRuntime)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ScriptExport) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetSchedule

`func (o *ScriptExport) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptExport) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptExport) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.


### GetSecretNames

`func (o *ScriptExport) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ScriptExport) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ScriptExport) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetType

`func (o *ScriptExport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptExport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptExport) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ScriptExport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScriptExport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScriptExport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWebhook

`func (o *ScriptExport) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptExport) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptExport) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.


### GetWebhookUrl

`func (o *ScriptExport) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ScriptExport) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ScriptExport) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ScriptExport) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


