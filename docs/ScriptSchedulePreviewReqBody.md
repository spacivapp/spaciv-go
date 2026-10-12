# ScriptSchedulePreviewReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] [default to 5]
**Cron** | Pointer to **string** | Five-field cron expression (minute, hour, day of month, month, day of week). Empty means the script only runs from its webhook. | [optional] 
**Timezone** | Pointer to **string** | IANA time zone the cron expression is evaluated in. Defaults to UTC. | [optional] 

## Methods

### NewScriptSchedulePreviewReqBody

`func NewScriptSchedulePreviewReqBody() *ScriptSchedulePreviewReqBody`

NewScriptSchedulePreviewReqBody instantiates a new ScriptSchedulePreviewReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptSchedulePreviewReqBodyWithDefaults

`func NewScriptSchedulePreviewReqBodyWithDefaults() *ScriptSchedulePreviewReqBody`

NewScriptSchedulePreviewReqBodyWithDefaults instantiates a new ScriptSchedulePreviewReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ScriptSchedulePreviewReqBody) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScriptSchedulePreviewReqBody) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScriptSchedulePreviewReqBody) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScriptSchedulePreviewReqBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCron

`func (o *ScriptSchedulePreviewReqBody) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ScriptSchedulePreviewReqBody) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ScriptSchedulePreviewReqBody) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ScriptSchedulePreviewReqBody) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetTimezone

`func (o *ScriptSchedulePreviewReqBody) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScriptSchedulePreviewReqBody) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScriptSchedulePreviewReqBody) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ScriptSchedulePreviewReqBody) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


