# ScriptSummary

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
**Type** | **string** | The handler signature the script implements. | 
**UpdatedAt** | **time.Time** |  | 
**Webhook** | **bool** |  | 

## Methods

### NewScriptSummary

`func NewScriptSummary(createdAt time.Time, createdBy int64, enabled bool, id string, key ScriptKey, language string, name string, revision int64, schedule ScriptSchedule, type_ string, updatedAt time.Time, webhook bool, ) *ScriptSummary`

NewScriptSummary instantiates a new ScriptSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptSummaryWithDefaults

`func NewScriptSummaryWithDefaults() *ScriptSummary`

NewScriptSummaryWithDefaults instantiates a new ScriptSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScriptSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScriptSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScriptSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ScriptSummary) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScriptSummary) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScriptSummary) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *ScriptSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptSummary) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptSummary) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptSummary) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ScriptSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptSummary) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ScriptSummary) GetKey() ScriptKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScriptSummary) GetKeyOk() (*ScriptKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScriptSummary) SetKey(v ScriptKey)`

SetKey sets Key field to given value.


### GetLanguage

`func (o *ScriptSummary) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScriptSummary) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScriptSummary) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetName

`func (o *ScriptSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptSummary) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *ScriptSummary) GetProject() int64`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScriptSummary) GetProjectOk() (*int64, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScriptSummary) SetProject(v int64)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScriptSummary) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ScriptSummary) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScriptSummary) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScriptSummary) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetSchedule

`func (o *ScriptSummary) GetSchedule() ScriptSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScriptSummary) GetScheduleOk() (*ScriptSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScriptSummary) SetSchedule(v ScriptSchedule)`

SetSchedule sets Schedule field to given value.


### GetType

`func (o *ScriptSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptSummary) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ScriptSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScriptSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScriptSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWebhook

`func (o *ScriptSummary) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ScriptSummary) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ScriptSummary) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


