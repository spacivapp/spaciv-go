# CalcListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Function** | [**NullableMaybeVarFunc**](MaybeVarFunc.md) |  | 
**HexCode** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | [**NullableMaybeVarStr**](MaybeVarStr.md) |  | 
**Visualisation** | Pointer to [**DashboardVisualisationCalc**](DashboardVisualisationCalc.md) |  | [optional] 

## Methods

### NewCalcListEntry

`func NewCalcListEntry(function NullableMaybeVarFunc, name string, params map[string]ParameterType, scenario NullableMaybeVarStr, ) *CalcListEntry`

NewCalcListEntry instantiates a new CalcListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListEntryWithDefaults

`func NewCalcListEntryWithDefaults() *CalcListEntry`

NewCalcListEntryWithDefaults instantiates a new CalcListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcListEntry) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcListEntry) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcListEntry) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcListEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *CalcListEntry) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *CalcListEntry) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *CalcListEntry) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *CalcListEntry) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetFunction

`func (o *CalcListEntry) GetFunction() MaybeVarFunc`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcListEntry) GetFunctionOk() (*MaybeVarFunc, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcListEntry) SetFunction(v MaybeVarFunc)`

SetFunction sets Function field to given value.


### SetFunctionNil

`func (o *CalcListEntry) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *CalcListEntry) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil
### GetHexCode

`func (o *CalcListEntry) GetHexCode() string`

GetHexCode returns the HexCode field if non-nil, zero value otherwise.

### GetHexCodeOk

`func (o *CalcListEntry) GetHexCodeOk() (*string, bool)`

GetHexCodeOk returns a tuple with the HexCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexCode

`func (o *CalcListEntry) SetHexCode(v string)`

SetHexCode sets HexCode field to given value.

### HasHexCode

`func (o *CalcListEntry) HasHexCode() bool`

HasHexCode returns a boolean if a field has been set.

### GetName

`func (o *CalcListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalcListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalcListEntry) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *CalcListEntry) GetParams() map[string]ParameterType`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcListEntry) GetParamsOk() (*map[string]ParameterType, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcListEntry) SetParams(v map[string]ParameterType)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcListEntry) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcListEntry) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcListEntry) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.


### SetScenarioNil

`func (o *CalcListEntry) SetScenarioNil(b bool)`

 SetScenarioNil sets the value for Scenario to be an explicit nil

### UnsetScenario
`func (o *CalcListEntry) UnsetScenario()`

UnsetScenario ensures that no value is present for Scenario, not even an explicit nil
### GetVisualisation

`func (o *CalcListEntry) GetVisualisation() DashboardVisualisationCalc`

GetVisualisation returns the Visualisation field if non-nil, zero value otherwise.

### GetVisualisationOk

`func (o *CalcListEntry) GetVisualisationOk() (*DashboardVisualisationCalc, bool)`

GetVisualisationOk returns a tuple with the Visualisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisation

`func (o *CalcListEntry) SetVisualisation(v DashboardVisualisationCalc)`

SetVisualisation sets Visualisation field to given value.

### HasVisualisation

`func (o *CalcListEntry) HasVisualisation() bool`

HasVisualisation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


