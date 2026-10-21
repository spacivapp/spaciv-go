# CalcListEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | [**MaybeVarFuncRequest**](MaybeVarFuncRequest.md) |  | 
**HexCode** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Params** | [**map[string]ParameterTypeRequest**](ParameterTypeRequest.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


