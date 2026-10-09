# CalcErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Errors** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | Optional list of individual error details | [optional] 
**Instance** | Pointer to **string** | A URI reference that identifies the specific occurrence of the problem. | [optional] 
**InvalidModifierFormula** | Pointer to [**InvalidModifierFormulaDetail**](InvalidModifierFormulaDetail.md) |  | [optional] 
**Kind** | **string** |  | 
**MissingAttribute** | Pointer to [**MissingAttributeDetail**](MissingAttributeDetail.md) |  | [optional] 
**MissingObject** | Pointer to [**MissingObjectDetail**](MissingObjectDetail.md) |  | [optional] 
**ScenarioNotFound** | Pointer to [**ScenarioNotFoundDetail**](ScenarioNotFoundDetail.md) |  | [optional] 
**Status** | Pointer to **int64** | HTTP status code | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem type. This value should not change between occurrences of the error. | [optional] 
**Type** | Pointer to **string** | A URI reference to human-readable documentation for the error. | [optional] [default to "about:blank"]
**UnmappedParameter** | Pointer to [**UnmappedParameterDetail**](UnmappedParameterDetail.md) |  | [optional] 
**VariableNotFound** | Pointer to [**VariableNotFoundDetail**](VariableNotFoundDetail.md) |  | [optional] 
**VariableTypeMismatch** | Pointer to [**VariableTypeMismatchDetail**](VariableTypeMismatchDetail.md) |  | [optional] 

## Methods

### NewCalcErrorBody

`func NewCalcErrorBody(kind string, ) *CalcErrorBody`

NewCalcErrorBody instantiates a new CalcErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcErrorBodyWithDefaults

`func NewCalcErrorBodyWithDefaults() *CalcErrorBody`

NewCalcErrorBodyWithDefaults instantiates a new CalcErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *CalcErrorBody) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CalcErrorBody) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CalcErrorBody) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CalcErrorBody) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetErrors

`func (o *CalcErrorBody) GetErrors() []ErrorDetail`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CalcErrorBody) GetErrorsOk() (*[]ErrorDetail, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CalcErrorBody) SetErrors(v []ErrorDetail)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CalcErrorBody) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInstance

`func (o *CalcErrorBody) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CalcErrorBody) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CalcErrorBody) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CalcErrorBody) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetInvalidModifierFormula

`func (o *CalcErrorBody) GetInvalidModifierFormula() InvalidModifierFormulaDetail`

GetInvalidModifierFormula returns the InvalidModifierFormula field if non-nil, zero value otherwise.

### GetInvalidModifierFormulaOk

`func (o *CalcErrorBody) GetInvalidModifierFormulaOk() (*InvalidModifierFormulaDetail, bool)`

GetInvalidModifierFormulaOk returns a tuple with the InvalidModifierFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidModifierFormula

`func (o *CalcErrorBody) SetInvalidModifierFormula(v InvalidModifierFormulaDetail)`

SetInvalidModifierFormula sets InvalidModifierFormula field to given value.

### HasInvalidModifierFormula

`func (o *CalcErrorBody) HasInvalidModifierFormula() bool`

HasInvalidModifierFormula returns a boolean if a field has been set.

### GetKind

`func (o *CalcErrorBody) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CalcErrorBody) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CalcErrorBody) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMissingAttribute

`func (o *CalcErrorBody) GetMissingAttribute() MissingAttributeDetail`

GetMissingAttribute returns the MissingAttribute field if non-nil, zero value otherwise.

### GetMissingAttributeOk

`func (o *CalcErrorBody) GetMissingAttributeOk() (*MissingAttributeDetail, bool)`

GetMissingAttributeOk returns a tuple with the MissingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAttribute

`func (o *CalcErrorBody) SetMissingAttribute(v MissingAttributeDetail)`

SetMissingAttribute sets MissingAttribute field to given value.

### HasMissingAttribute

`func (o *CalcErrorBody) HasMissingAttribute() bool`

HasMissingAttribute returns a boolean if a field has been set.

### GetMissingObject

`func (o *CalcErrorBody) GetMissingObject() MissingObjectDetail`

GetMissingObject returns the MissingObject field if non-nil, zero value otherwise.

### GetMissingObjectOk

`func (o *CalcErrorBody) GetMissingObjectOk() (*MissingObjectDetail, bool)`

GetMissingObjectOk returns a tuple with the MissingObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingObject

`func (o *CalcErrorBody) SetMissingObject(v MissingObjectDetail)`

SetMissingObject sets MissingObject field to given value.

### HasMissingObject

`func (o *CalcErrorBody) HasMissingObject() bool`

HasMissingObject returns a boolean if a field has been set.

### GetScenarioNotFound

`func (o *CalcErrorBody) GetScenarioNotFound() ScenarioNotFoundDetail`

GetScenarioNotFound returns the ScenarioNotFound field if non-nil, zero value otherwise.

### GetScenarioNotFoundOk

`func (o *CalcErrorBody) GetScenarioNotFoundOk() (*ScenarioNotFoundDetail, bool)`

GetScenarioNotFoundOk returns a tuple with the ScenarioNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioNotFound

`func (o *CalcErrorBody) SetScenarioNotFound(v ScenarioNotFoundDetail)`

SetScenarioNotFound sets ScenarioNotFound field to given value.

### HasScenarioNotFound

`func (o *CalcErrorBody) HasScenarioNotFound() bool`

HasScenarioNotFound returns a boolean if a field has been set.

### GetStatus

`func (o *CalcErrorBody) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CalcErrorBody) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CalcErrorBody) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CalcErrorBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CalcErrorBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CalcErrorBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CalcErrorBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CalcErrorBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CalcErrorBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CalcErrorBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CalcErrorBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CalcErrorBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnmappedParameter

`func (o *CalcErrorBody) GetUnmappedParameter() UnmappedParameterDetail`

GetUnmappedParameter returns the UnmappedParameter field if non-nil, zero value otherwise.

### GetUnmappedParameterOk

`func (o *CalcErrorBody) GetUnmappedParameterOk() (*UnmappedParameterDetail, bool)`

GetUnmappedParameterOk returns a tuple with the UnmappedParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmappedParameter

`func (o *CalcErrorBody) SetUnmappedParameter(v UnmappedParameterDetail)`

SetUnmappedParameter sets UnmappedParameter field to given value.

### HasUnmappedParameter

`func (o *CalcErrorBody) HasUnmappedParameter() bool`

HasUnmappedParameter returns a boolean if a field has been set.

### GetVariableNotFound

`func (o *CalcErrorBody) GetVariableNotFound() VariableNotFoundDetail`

GetVariableNotFound returns the VariableNotFound field if non-nil, zero value otherwise.

### GetVariableNotFoundOk

`func (o *CalcErrorBody) GetVariableNotFoundOk() (*VariableNotFoundDetail, bool)`

GetVariableNotFoundOk returns a tuple with the VariableNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNotFound

`func (o *CalcErrorBody) SetVariableNotFound(v VariableNotFoundDetail)`

SetVariableNotFound sets VariableNotFound field to given value.

### HasVariableNotFound

`func (o *CalcErrorBody) HasVariableNotFound() bool`

HasVariableNotFound returns a boolean if a field has been set.

### GetVariableTypeMismatch

`func (o *CalcErrorBody) GetVariableTypeMismatch() VariableTypeMismatchDetail`

GetVariableTypeMismatch returns the VariableTypeMismatch field if non-nil, zero value otherwise.

### GetVariableTypeMismatchOk

`func (o *CalcErrorBody) GetVariableTypeMismatchOk() (*VariableTypeMismatchDetail, bool)`

GetVariableTypeMismatchOk returns a tuple with the VariableTypeMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableTypeMismatch

`func (o *CalcErrorBody) SetVariableTypeMismatch(v VariableTypeMismatchDetail)`

SetVariableTypeMismatch sets VariableTypeMismatch field to given value.

### HasVariableTypeMismatch

`func (o *CalcErrorBody) HasVariableTypeMismatch() bool`

HasVariableTypeMismatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


