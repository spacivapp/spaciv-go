# CalcBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DataShapesRequest**](DataShapesRequest.md) |  | 
**Dimensions** | Pointer to **[]int32** | The order to lay the result&#39;s values out in, outermost first. &#x60;1&#x60; is day, &#x60;2&#x60; tree node, &#x60;3&#x60; series, &#x60;4&#x60; scenario and &#x60;5&#x60; group. Every result has day and tree node; a card adds series (single values, categorisation mixes, lists and tree lists), scenario (scenario comparisons and tree lists over scenarios) or, for a 2D list, group for its rows and series or scenario for its columns. List each of the card&#39;s dimensions exactly once, in any order; leave this out for day, tree node, then the card&#39;s own dimensions. Any other list is rejected with &#x60;400&#x60; and &#x60;kind&#x60; &#x60;invalid_dimensions&#x60;. | [optional] 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**IncludeObjects** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**VariableOverrides** | Pointer to [**map[string]VariableData**](VariableData.md) |  | [optional] 

## Methods

### NewCalcBody

`func NewCalcBody(data DataShapesRequest, filters []FilterGroup, ) *CalcBody`

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

`func (o *CalcBody) GetData() DataShapesRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CalcBody) GetDataOk() (*DataShapesRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CalcBody) SetData(v DataShapesRequest)`

SetData sets Data field to given value.


### GetDimensions

`func (o *CalcBody) GetDimensions() []int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CalcBody) GetDimensionsOk() (*[]int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CalcBody) SetDimensions(v []int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CalcBody) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

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


