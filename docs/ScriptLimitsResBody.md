# ScriptLimitsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Defaults** | [**ScriptLimitValues**](ScriptLimitValues.md) | The code defaults for scripts. | 
**Effective** | [**ScriptLimitValues**](ScriptLimitValues.md) | The limits that apply to the account. | 
**Overrides** | [**ScriptLimitOverrides**](ScriptLimitOverrides.md) | This account&#39;s overrides; an absent field uses the default. | 

## Methods

### NewScriptLimitsResBody

`func NewScriptLimitsResBody(defaults ScriptLimitValues, effective ScriptLimitValues, overrides ScriptLimitOverrides, ) *ScriptLimitsResBody`

NewScriptLimitsResBody instantiates a new ScriptLimitsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptLimitsResBodyWithDefaults

`func NewScriptLimitsResBodyWithDefaults() *ScriptLimitsResBody`

NewScriptLimitsResBodyWithDefaults instantiates a new ScriptLimitsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaults

`func (o *ScriptLimitsResBody) GetDefaults() ScriptLimitValues`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ScriptLimitsResBody) GetDefaultsOk() (*ScriptLimitValues, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ScriptLimitsResBody) SetDefaults(v ScriptLimitValues)`

SetDefaults sets Defaults field to given value.


### GetEffective

`func (o *ScriptLimitsResBody) GetEffective() ScriptLimitValues`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *ScriptLimitsResBody) GetEffectiveOk() (*ScriptLimitValues, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *ScriptLimitsResBody) SetEffective(v ScriptLimitValues)`

SetEffective sets Effective field to given value.


### GetOverrides

`func (o *ScriptLimitsResBody) GetOverrides() ScriptLimitOverrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ScriptLimitsResBody) GetOverridesOk() (*ScriptLimitOverrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ScriptLimitsResBody) SetOverrides(v ScriptLimitOverrides)`

SetOverrides sets Overrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


