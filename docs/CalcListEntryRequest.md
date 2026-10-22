# CalcListEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**MaybeVarInt**](MaybeVarInt.md) |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Function** | [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | 
**HexCode** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 
**Visualisation** | Pointer to [**DashboardVisualisationCalcRequest**](DashboardVisualisationCalcRequest.md) |  | [optional] 

## Methods

### NewCalcListEntryRequest

`func NewCalcListEntryRequest(function MaybeVarFuncRequest, name string, params map[string]ParameterTypeRequest, scenario MaybeVarStr, ) *CalcListEntryRequest`

NewCalcListEntryRequest instantiates a new CalcListEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListEntryRequestWithDefaults

`func NewCalcListEntryRequestWithDefaults() *CalcListEntryRequest`

NewCalcListEntryRequestWithDefaults instantiates a new CalcListEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalcListEntryRequest) GetDate() MaybeVarInt`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalcListEntryRequest) GetDateOk() (*MaybeVarInt, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalcListEntryRequest) SetDate(v MaybeVarInt)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalcListEntryRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *CalcListEntryRequest) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *CalcListEntryRequest) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *CalcListEntryRequest) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *CalcListEntryRequest) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetFunction

`func (o *CalcListEntryRequest) GetFunction() MaybeVarFuncRequest`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CalcListEntryRequest) GetFunctionOk() (*MaybeVarFuncRequest, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CalcListEntryRequest) SetFunction(v MaybeVarFuncRequest)`

SetFunction sets Function field to given value.


### GetHexCode

`func (o *CalcListEntryRequest) GetHexCode() string`

GetHexCode returns the HexCode field if non-nil, zero value otherwise.

### GetHexCodeOk

`func (o *CalcListEntryRequest) GetHexCodeOk() (*string, bool)`

GetHexCodeOk returns a tuple with the HexCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexCode

`func (o *CalcListEntryRequest) SetHexCode(v string)`

SetHexCode sets HexCode field to given value.

### HasHexCode

`func (o *CalcListEntryRequest) HasHexCode() bool`

HasHexCode returns a boolean if a field has been set.

### GetName

`func (o *CalcListEntryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalcListEntryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalcListEntryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *CalcListEntryRequest) GetParams() map[string]ParameterTypeRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CalcListEntryRequest) GetParamsOk() (*map[string]ParameterTypeRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CalcListEntryRequest) SetParams(v map[string]ParameterTypeRequest)`

SetParams sets Params field to given value.


### GetScenario

`func (o *CalcListEntryRequest) GetScenario() MaybeVarStr`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CalcListEntryRequest) GetScenarioOk() (*MaybeVarStr, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CalcListEntryRequest) SetScenario(v MaybeVarStr)`

SetScenario sets Scenario field to given value.


### GetVisualisation

`func (o *CalcListEntryRequest) GetVisualisation() DashboardVisualisationCalcRequest`

GetVisualisation returns the Visualisation field if non-nil, zero value otherwise.

### GetVisualisationOk

`func (o *CalcListEntryRequest) GetVisualisationOk() (*DashboardVisualisationCalcRequest, bool)`

GetVisualisationOk returns a tuple with the Visualisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisation

`func (o *CalcListEntryRequest) SetVisualisation(v DashboardVisualisationCalcRequest)`

SetVisualisation sets Visualisation field to given value.

### HasVisualisation

`func (o *CalcListEntryRequest) HasVisualisation() bool`

HasVisualisation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


