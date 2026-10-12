# ScriptPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether schedules and webhooks start runs. | [optional] 
**Files** | **map[string]string** | Source files keyed by path. Python needs script.py defining handle(event, ctx), JavaScript needs script.js exporting handle(event, ctx), and Go needs script.go in a non-main package exporting Handle(ctx context.Context, event handler.Event, hc handler.Context) (any, error) from spaciv/handler. | 
**Language** | **string** |  | 
**Manifest** | Pointer to **string** | Dependency manifest: requirements.txt for Python, go.mod for Go, package.json for JavaScript. | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to **int64** | Optional project to scope the script&#39;s API key to, for example to update its dashboards. Omit for account-wide access. | [optional] 
**Schedule** | Pointer to [**ScriptSchedule**](ScriptSchedule.md) |  | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets exposed to the script as environment variables. Names must be upper case and must not start with SPACIV_. | [optional] 
**Type** | Pointer to **string** | The handler signature the script implements; defaults to integration and cannot be changed. An integration handles each run as handle(event, ctx), where event has the trigger and, for webhook runs, the request, and ctx has the account, project, API key and secrets. | [optional] 
**Webhook** | Pointer to **bool** | Whether the script can be triggered by an HTTP POST to its webhook URL. | [optional] 

## Methods

### NewScriptPostReqBody

`func NewScriptPostReqBody(files map[string]string, language string, name string, ) *ScriptPostReqBody`

NewScriptPostReqBody instantiates a new ScriptPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptPostReqBodyWithDefaults

`func NewScriptPostReqBodyWithDefaults() *ScriptPostReqBody`

NewScriptPostReqBodyWithDefaults instantiates a new ScriptPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScriptPostReqBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptPostReqBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptPostReqBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptPostReqBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptPostReqBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptPostReqBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptPostReqBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScriptPostReqBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFiles

`func (o *ScriptPostReqBody) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScriptPostReqBody) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScriptPostReqBody) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.


### GetLanguage

`func (o *ScriptPostReqBody) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptPostReqBody) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptPostReqBody) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetManifest

`func (o *ScriptPostReqBody) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ScriptPostReqBody) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ScriptPostReqBody) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ScriptPostReqBody) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *ScriptPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptPostReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ScriptPostReqBody) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptPostReqBody) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptPostReqBody) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptPostReqBody) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSchedule

`func (o *ScriptPostReqBody) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptPostReqBody) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptPostReqBody) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ScriptPostReqBody) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSecrets

`func (o *ScriptPostReqBody) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ScriptPostReqBody) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ScriptPostReqBody) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ScriptPostReqBody) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetType

`func (o *ScriptPostReqBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptPostReqBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptPostReqBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScriptPostReqBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebhook

`func (o *ScriptPostReqBody) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptPostReqBody) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptPostReqBody) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *ScriptPostReqBody) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


