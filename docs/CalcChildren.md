# CalcChildren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**NullableMaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | [**NullableMaybeVarFunc**](MaybeVarFunc.md) |  | 
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | [**NullableMaybeVarStr**](MaybeVarStr.md) |  | 
**Segment** | Pointer to [**MaybeVarProp**](MaybeVarProp.md) |  | [optional] 

## Methods

### NewCalcChildren

`func NewCalcChildren(date NullableMaybeVarInt, filters []FilterGroup, function NullableMaybeVarFunc, params map[string]ParameterType, scenario NullableMaybeVarStr, ) *CalcChildren`

NewCalcChildren instantiates a new CalcChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcChildrenWithDefaults

`func NewCalcChildrenWithDefaults() *CalcChildren`

NewCalcChildrenWithDefaults instantiates a new CalcChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcChildren) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcChildren) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcChildren) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *CalcChildren) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *CalcChildren) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetFilters

`func (o *CalcChildren) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcChildren) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcChildren) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcChildren) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcChildren) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcChildren) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.


### SetFunctionNil

`func (o *CalcChildren) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *CalcChildren) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil
### GetObject

`func (o *CalcChildren) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CalcChildren) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CalcChildren) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *CalcChildren) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParams

`func (o *CalcChildren) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcChildren) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcChildren) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcChildren) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcChildren) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcChildren) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.


### SetScenarioNil

`func (o *CalcChildren) SetScenarioNil(b bool)`

 SetScenarioNil sets the value for Scenario to be an explicit nil

### UnsetScenario
`func (o *CalcChildren) UnsetScenario()`

UnsetScenario ensures that no value is present for Scenario, not even an explicit nil
### GetSegment

`func (o *CalcChildren) GetSegment() MaybeVarProp`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *CalcChildren) GetSegmentOk() (*MaybeVarProp, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *CalcChildren) SetSegment(v MaybeVarProp)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *CalcChildren) HasSegment() bool`

HasSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


