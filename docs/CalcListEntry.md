# CalcListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | [**MaybeVarFunc**](MaybeVarFunc.md) |  | 
**Name** | **string** |  | 
**Params** | [**map[string]ParameterType**](ParameterType.md) |  | 
**Scenario** | [**MaybeVarStr**](MaybeVarStr.md) |  | 

## Methods

### NewCalcListEntry

`func NewCalcListEntry(function MaybeVarFunc, name string, params map[string]ParameterType, scenario MaybeVarStr, ) *CalcListEntry`

NewCalcListEntry instantiates a new CalcListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcListEntryWithDefaults

`func NewCalcListEntryWithDefaults() *CalcListEntry`

NewCalcListEntryWithDefaults instantiates a new CalcListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


