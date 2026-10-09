# CalcBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DataShapes**](DataShapes.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**IncludeObjects** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**VariableOverrides** | Pointer to [**map[string]VariableData**](VariableData.md) |  | [optional] 

## Methods

### NewCalcBody

`func NewCalcBody(data DataShapes, filters []FilterGroup, ) *CalcBody`

NewCalcBody instantiates a new CalcBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcBodyWithDefaults

`func NewCalcBodyWithDefaults() *CalcBody`

NewCalcBodyWithDefaults instantiates a new CalcBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CalcBody) GetData() DataShapes`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CalcBody) GetDataOk() (*DataShapes, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CalcBody) SetData(v DataShapes)`

SetData sets Data field to given value.


### GetFilters

`func (o *CalcBody) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcBody) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcBody) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetIncludeObjects

`func (o *CalcBody) GetIncludeObjects() bool`

GetIncludeObjects returns the IncludeObjects field if non-nil, zero value otherwise.

### GetIncludeObjectsOk

`func (o *CalcBody) GetIncludeObjectsOk() (*bool, bool)`

GetIncludeObjectsOk returns a tuple with the IncludeObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjects

`func (o *CalcBody) SetIncludeObjects(v bool)`

SetIncludeObjects sets IncludeObjects field to given value.

### HasIncludeObjects

`func (o *CalcBody) HasIncludeObjects() bool`

HasIncludeObjects returns a boolean if a field has been set.

### GetLimit

`func (o *CalcBody) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CalcBody) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CalcBody) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CalcBody) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetVariableOverrides

`func (o *CalcBody) GetVariableOverrides() map[string]VariableData`

GetVariableOverrides returns the VariableOverrides field if non-nil, zero value otherwise.

### GetVariableOverridesOk

`func (o *CalcBody) GetVariableOverridesOk() (*map[string]VariableData, bool)`

GetVariableOverridesOk returns a tuple with the VariableOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableOverrides

`func (o *CalcBody) SetVariableOverrides(v map[string]VariableData)`

SetVariableOverrides sets VariableOverrides field to given value.

### HasVariableOverrides

`func (o *CalcBody) HasVariableOverrides() bool`

HasVariableOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


