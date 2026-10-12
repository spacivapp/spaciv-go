# ScriptRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildRunId** | Pointer to **string** |  | [optional] 
**BuildState** | Pointer to **string** | State of the latest dependency build. | [optional] 
**Cpu** | Pointer to **int32** | vCPU units per run (1024 &#x3D; 1 vCPU). | [optional] 
**MemoryMb** | Pointer to **int32** |  | [optional] 
**NextRunAt** | Pointer to **time.Time** |  | [optional] 
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewScriptRuntime

`func NewScriptRuntime() *ScriptRuntime`

NewScriptRuntime instantiates a new ScriptRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptRuntimeWithDefaults

`func NewScriptRuntimeWithDefaults() *ScriptRuntime`

NewScriptRuntimeWithDefaults instantiates a new ScriptRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildRunId

`func (o *ScriptRuntime) GetBuildRunId() string`

GetBuildRunId returns the BuildRunId field if non-nil, zero value otherwise.

### GetBuildRunIdOk

`func (o *ScriptRuntime) GetBuildRunIdOk() (*string, bool)`

GetBuildRunIdOk returns a tuple with the BuildRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRunId

`func (o *ScriptRuntime) SetBuildRunId(v string)`

SetBuildRunId sets BuildRunId field to given value.

### HasBuildRunId

`func (o *ScriptRuntime) HasBuildRunId() bool`

HasBuildRunId returns a boolean if a field has been set.

### GetBuildState

`func (o *ScriptRuntime) GetBuildState() string`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ScriptRuntime) GetBuildStateOk() (*string, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ScriptRuntime) SetBuildState(v string)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *ScriptRuntime) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetCpu

`func (o *ScriptRuntime) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ScriptRuntime) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ScriptRuntime) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ScriptRuntime) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemoryMb

`func (o *ScriptRuntime) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *ScriptRuntime) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *ScriptRuntime) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *ScriptRuntime) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetNextRunAt

`func (o *ScriptRuntime) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *ScriptRuntime) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *ScriptRuntime) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *ScriptRuntime) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ScriptRuntime) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ScriptRuntime) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ScriptRuntime) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ScriptRuntime) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


