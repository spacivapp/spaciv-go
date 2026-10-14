# ScriptUsageResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditState** | **string** | near from 80% of maxCreditsPerMonth; reached once they are used, after which runs are recorded as quota_exceeded until the month ends. | 
**CreditsUsed** | **int64** | Credits the account&#39;s script runs were charged this month: the scripts figure the account Usage tab shows. | 
**EnabledScripts** | **int64** |  | 
**FailedRuns** | **int64** | Runs that ended in error, timeout or killed. | 
**Limits** | [**ScriptLimits**](ScriptLimits.md) |  | 
**Period** | **string** | The calendar month (UTC) the usage covers. | 
**PeriodEnd** | **time.Time** |  | 
**PeriodStart** | **time.Time** |  | 
**QuotaExceededRuns** | **int64** | Runs that did not start because a limit was reached. | 
**Runs** | **int64** | Runs that were not refused by a limit, whatever their outcome, including runs that failed before their sandbox started. | 
**Scripts** | [**[]ScriptUsage**](ScriptUsage.md) | Usage of each script that has run this month, deleted scripts included, sorted by scriptId. | 

## Methods

### NewScriptUsageResBody

`func NewScriptUsageResBody(creditState string, creditsUsed int64, enabledScripts int64, failedRuns int64, limits ScriptLimits, period string, periodEnd time.Time, periodStart time.Time, quotaExceededRuns int64, runs int64, scripts []ScriptUsage, ) *ScriptUsageResBody`

NewScriptUsageResBody instantiates a new ScriptUsageResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptUsageResBodyWithDefaults

`func NewScriptUsageResBodyWithDefaults() *ScriptUsageResBody`

NewScriptUsageResBodyWithDefaults instantiates a new ScriptUsageResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditState

`func (o *ScriptUsageResBody) GetCreditState() string`

GetCreditState returns the CreditState field if non-nil, zero value otherwise.

### GetCreditStateOk

`func (o *ScriptUsageResBody) GetCreditStateOk() (*string, bool)`

GetCreditStateOk returns a tuple with the CreditState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditState

`func (o *ScriptUsageResBody) SetCreditState(v string)`

SetCreditState sets CreditState field to given value.


### GetCreditsUsed

`func (o *ScriptUsageResBody) GetCreditsUsed() int64`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *ScriptUsageResBody) GetCreditsUsedOk() (*int64, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *ScriptUsageResBody) SetCreditsUsed(v int64)`

SetCreditsUsed sets CreditsUsed field to given value.


### GetEnabledScripts

`func (o *ScriptUsageResBody) GetEnabledScripts() int64`

GetEnabledScripts returns the EnabledScripts field if non-nil, zero value otherwise.

### GetEnabledScriptsOk

`func (o *ScriptUsageResBody) GetEnabledScriptsOk() (*int64, bool)`

GetEnabledScriptsOk returns a tuple with the EnabledScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledScripts

`func (o *ScriptUsageResBody) SetEnabledScripts(v int64)`

SetEnabledScripts sets EnabledScripts field to given value.


### GetFailedRuns

`func (o *ScriptUsageResBody) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *ScriptUsageResBody) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *ScriptUsageResBody) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetLimits

`func (o *ScriptUsageResBody) GetLimits() ScriptLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ScriptUsageResBody) GetLimitsOk() (*ScriptLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ScriptUsageResBody) SetLimits(v ScriptLimits)`

SetLimits sets Limits field to given value.


### GetPeriod

`func (o *ScriptUsageResBody) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ScriptUsageResBody) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ScriptUsageResBody) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetPeriodEnd

`func (o *ScriptUsageResBody) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ScriptUsageResBody) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ScriptUsageResBody) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *ScriptUsageResBody) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ScriptUsageResBody) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ScriptUsageResBody) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetQuotaExceededRuns

`func (o *ScriptUsageResBody) GetQuotaExceededRuns() int64`

GetQuotaExceededRuns returns the QuotaExceededRuns field if non-nil, zero value otherwise.

### GetQuotaExceededRunsOk

`func (o *ScriptUsageResBody) GetQuotaExceededRunsOk() (*int64, bool)`

GetQuotaExceededRunsOk returns a tuple with the QuotaExceededRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExceededRuns

`func (o *ScriptUsageResBody) SetQuotaExceededRuns(v int64)`

SetQuotaExceededRuns sets QuotaExceededRuns field to given value.


### GetRuns

`func (o *ScriptUsageResBody) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ScriptUsageResBody) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ScriptUsageResBody) SetRuns(v int64)`

SetRuns sets Runs field to given value.


### GetScripts

`func (o *ScriptUsageResBody) GetScripts() []ScriptUsage`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ScriptUsageResBody) GetScriptsOk() (*[]ScriptUsage, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ScriptUsageResBody) SetScripts(v []ScriptUsage)`

SetScripts sets Scripts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


