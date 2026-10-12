# ScriptPatchReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Files** | Pointer to **map[string]string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Manifest** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **int64** | 0 removes the project. | [optional] 
**Revision** | **int64** | The revision last read; a different current revision is rejected with 409. | 
**Schedule** | Pointer to [**ScriptSchedule**](ScriptSchedule.md) |  | [optional] 
**Webhook** | Pointer to **bool** |  | [optional] 

## Methods

### NewScriptPatchReqBody

`func NewScriptPatchReqBody(revision int64, ) *ScriptPatchReqBody`

NewScriptPatchReqBody instantiates a new ScriptPatchReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptPatchReqBodyWithDefaults

`func NewScriptPatchReqBodyWithDefaults() *ScriptPatchReqBody`

NewScriptPatchReqBodyWithDefaults instantiates a new ScriptPatchReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScriptPatchReqBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptPatchReqBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptPatchReqBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptPatchReqBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptPatchReqBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptPatchReqBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptPatchReqBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScriptPatchReqBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFiles

`func (o *ScriptPatchReqBody) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScriptPatchReqBody) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScriptPatchReqBody) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ScriptPatchReqBody) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLanguage

`func (o *ScriptPatchReqBody) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptPatchReqBody) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptPatchReqBody) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ScriptPatchReqBody) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetManifest

`func (o *ScriptPatchReqBody) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ScriptPatchReqBody) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ScriptPatchReqBody) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ScriptPatchReqBody) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *ScriptPatchReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptPatchReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptPatchReqBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptPatchReqBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *ScriptPatchReqBody) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptPatchReqBody) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptPatchReqBody) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptPatchReqBody) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptPatchReqBody) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptPatchReqBody) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptPatchReqBody) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetSchedule

`func (o *ScriptPatchReqBody) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptPatchReqBody) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptPatchReqBody) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ScriptPatchReqBody) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetWebhook

`func (o *ScriptPatchReqBody) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptPatchReqBody) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptPatchReqBody) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *ScriptPatchReqBody) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


