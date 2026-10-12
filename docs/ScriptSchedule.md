# ScriptSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to **string** | Five-field cron expression (minute, hour, day of month, month, day of week). Empty means the script only runs from its webhook. | [optional] 
**Timezone** | Pointer to **string** | IANA time zone the cron expression is evaluated in. Defaults to UTC. | [optional] 

## Methods

### NewScriptSchedule

`func NewScriptSchedule() *ScriptSchedule`

NewScriptSchedule instantiates a new ScriptSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptScheduleWithDefaults

`func NewScriptScheduleWithDefaults() *ScriptSchedule`

NewScriptScheduleWithDefaults instantiates a new ScriptSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *ScriptSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ScriptSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ScriptSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ScriptSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetTimezone

`func (o *ScriptSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScriptSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScriptSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ScriptSchedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


