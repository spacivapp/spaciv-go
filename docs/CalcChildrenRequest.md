# CalcChildrenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**MaybeVarInt**](MaybeVarInt.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**Function** | [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | 
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 
**Segment** | Pointer to [**MaybeVarProp**](MaybeVarProp.md) |  | [optional] 

## Methods

### NewCalcChildrenRequest

`func NewCalcChildrenRequest(date MaybeVarInt, filters []FilterGroup, function MaybeVarFuncRequest, params map[string]ParameterTypeRequest, scenario MaybeVarStr, ) *CalcChildrenRequest`

NewCalcChildrenRequest instantiates a new CalcChildrenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcChildrenRequestWithDefaults

`func NewCalcChildrenRequestWithDefaults() *CalcChildrenRequest`

NewCalcChildrenRequestWithDefaults instantiates a new CalcChildrenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcChildrenRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcChildrenRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcChildrenRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.


### GetFilters

`func (o *CalcChildrenRequest) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalcChildrenRequest) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalcChildrenRequest) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetFunction

`func (o *CalcChildrenRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcChildrenRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcChildrenRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.


### GetObject

`func (o *CalcChildrenRequest) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CalcChildrenRequest) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CalcChildrenRequest) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *CalcChildrenRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParams

`func (o *CalcChildrenRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcChildrenRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcChildrenRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcChildrenRequest) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcChildrenRequest) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcChildrenRequest) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.


### GetSegment

`func (o *CalcChildrenRequest) GetSegment() MaybeVarProp`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *CalcChildrenRequest) GetSegmentOk() (*MaybeVarProp, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *CalcChildrenRequest) SetSegment(v MaybeVarProp)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *CalcChildrenRequest) HasSegment() bool`

HasSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


