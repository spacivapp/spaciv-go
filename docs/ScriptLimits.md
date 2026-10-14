# ScriptLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentRuns** | **int32** | Runs the account may have in progress at once. | 
**MaxCreditsPerMonth** | **int64** | Credits the account&#39;s script runs may be charged in a calendar month (UTC). Once they are used, further runs are recorded as quota_exceeded and do not start. | 
**MaxEnabledScripts** | **int32** | Scripts the account may have enabled at once. | 
**MaxRunSeconds** | **int32** | Longest a single run may take; a script&#39;s own timeout is capped to it. | 
**MinCronIntervalSeconds** | **int32** | Shortest time allowed between two scheduled runs of a script. | 
**WebhookRequestsPerMinute** | **int32** | Webhook calls each script accepts per minute, and at most in a burst. | 

## Methods

### NewScriptLimits

`func NewScriptLimits(maxConcurrentRuns int32, maxCreditsPerMonth int64, maxEnabledScripts int32, maxRunSeconds int32, minCronIntervalSeconds int32, webhookRequestsPerMinute int32, ) *ScriptLimits`

NewScriptLimits instantiates a new ScriptLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptLimitsWithDefaults

`func NewScriptLimitsWithDefaults() *ScriptLimits`

NewScriptLimitsWithDefaults instantiates a new ScriptLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConcurrentRuns

`func (o *ScriptLimits) GetMaxConcurrentRuns() int32`

GetMaxConcurrentRuns returns the MaxConcurrentRuns field if non-nil, zero value otherwise.

### GetMaxConcurrentRunsOk

`func (o *ScriptLimits) GetMaxConcurrentRunsOk() (*int32, bool)`

GetMaxConcurrentRunsOk returns a tuple with the MaxConcurrentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentRuns

`func (o *ScriptLimits) SetMaxConcurrentRuns(v int32)`

SetMaxConcurrentRuns sets MaxConcurrentRuns field to given value.


### GetMaxCreditsPerMonth

`func (o *ScriptLimits) GetMaxCreditsPerMonth() int64`

GetMaxCreditsPerMonth returns the MaxCreditsPerMonth field if non-nil, zero value otherwise.

### GetMaxCreditsPerMonthOk

`func (o *ScriptLimits) GetMaxCreditsPerMonthOk() (*int64, bool)`

GetMaxCreditsPerMonthOk returns a tuple with the MaxCreditsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreditsPerMonth

`func (o *ScriptLimits) SetMaxCreditsPerMonth(v int64)`

SetMaxCreditsPerMonth sets MaxCreditsPerMonth field to given value.


### GetMaxEnabledScripts

`func (o *ScriptLimits) GetMaxEnabledScripts() int32`

GetMaxEnabledScripts returns the MaxEnabledScripts field if non-nil, zero value otherwise.

### GetMaxEnabledScriptsOk

`func (o *ScriptLimits) GetMaxEnabledScriptsOk() (*int32, bool)`

GetMaxEnabledScriptsOk returns a tuple with the MaxEnabledScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEnabledScripts

`func (o *ScriptLimits) SetMaxEnabledScripts(v int32)`

SetMaxEnabledScripts sets MaxEnabledScripts field to given value.


### GetMaxRunSeconds

`func (o *ScriptLimits) GetMaxRunSeconds() int32`

GetMaxRunSeconds returns the MaxRunSeconds field if non-nil, zero value otherwise.

### GetMaxRunSecondsOk

`func (o *ScriptLimits) GetMaxRunSecondsOk() (*int32, bool)`

GetMaxRunSecondsOk returns a tuple with the MaxRunSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunSeconds

`func (o *ScriptLimits) SetMaxRunSeconds(v int32)`

SetMaxRunSeconds sets MaxRunSeconds field to given value.


### GetMinCronIntervalSeconds

`func (o *ScriptLimits) GetMinCronIntervalSeconds() int32`

GetMinCronIntervalSeconds returns the MinCronIntervalSeconds field if non-nil, zero value otherwise.

### GetMinCronIntervalSecondsOk

`func (o *ScriptLimits) GetMinCronIntervalSecondsOk() (*int32, bool)`

GetMinCronIntervalSecondsOk returns a tuple with the MinCronIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCronIntervalSeconds

`func (o *ScriptLimits) SetMinCronIntervalSeconds(v int32)`

SetMinCronIntervalSeconds sets MinCronIntervalSeconds field to given value.


### GetWebhookRequestsPerMinute

`func (o *ScriptLimits) GetWebhookRequestsPerMinute() int32`

GetWebhookRequestsPerMinute returns the WebhookRequestsPerMinute field if non-nil, zero value otherwise.

### GetWebhookRequestsPerMinuteOk

`func (o *ScriptLimits) GetWebhookRequestsPerMinuteOk() (*int32, bool)`

GetWebhookRequestsPerMinuteOk returns a tuple with the WebhookRequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookRequestsPerMinute

`func (o *ScriptLimits) SetWebhookRequestsPerMinute(v int32)`

SetWebhookRequestsPerMinute sets WebhookRequestsPerMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


