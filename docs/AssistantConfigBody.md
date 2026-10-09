# AssistantConfigBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetLimitCredits** | **NullableInt64** |  | 
**BudgetLimitMaxCredits** | **NullableInt64** |  | 
**BudgetLimitMinCredits** | **NullableInt64** |  | 
**Models** | [**[]IDNameStringString**](IDNameStringString.md) |  | 
**PermittedModels** | **[]string** |  | 
**PricePerCreditEur** | **float64** |  | 
**SpentCredits** | **NullableInt64** |  | 
**ToolsEnabled** | **bool** |  | 

## Methods

### NewAssistantConfigBody

`func NewAssistantConfigBody(budgetLimitCredits NullableInt64, budgetLimitMaxCredits NullableInt64, budgetLimitMinCredits NullableInt64, models []IDNameStringString, permittedModels []string, pricePerCreditEur float64, spentCredits NullableInt64, toolsEnabled bool, ) *AssistantConfigBody`

NewAssistantConfigBody instantiates a new AssistantConfigBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantConfigBodyWithDefaults

`func NewAssistantConfigBodyWithDefaults() *AssistantConfigBody`

NewAssistantConfigBodyWithDefaults instantiates a new AssistantConfigBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetLimitCredits

`func (o *AssistantConfigBody) GetBudgetLimitCredits() int64`

GetBudgetLimitCredits returns the BudgetLimitCredits field if non-nil, zero value otherwise.

### GetBudgetLimitCreditsOk

`func (o *AssistantConfigBody) GetBudgetLimitCreditsOk() (*int64, bool)`

GetBudgetLimitCreditsOk returns a tuple with the BudgetLimitCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLimitCredits

`func (o *AssistantConfigBody) SetBudgetLimitCredits(v int64)`

SetBudgetLimitCredits sets BudgetLimitCredits field to given value.


### SetBudgetLimitCreditsNil

`func (o *AssistantConfigBody) SetBudgetLimitCreditsNil(b bool)`

 SetBudgetLimitCreditsNil sets the value for BudgetLimitCredits to be an explicit nil

### UnsetBudgetLimitCredits
`func (o *AssistantConfigBody) UnsetBudgetLimitCredits()`

UnsetBudgetLimitCredits ensures that no value is present for BudgetLimitCredits, not even an explicit nil
### GetBudgetLimitMaxCredits

`func (o *AssistantConfigBody) GetBudgetLimitMaxCredits() int64`

GetBudgetLimitMaxCredits returns the BudgetLimitMaxCredits field if non-nil, zero value otherwise.

### GetBudgetLimitMaxCreditsOk

`func (o *AssistantConfigBody) GetBudgetLimitMaxCreditsOk() (*int64, bool)`

GetBudgetLimitMaxCreditsOk returns a tuple with the BudgetLimitMaxCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLimitMaxCredits

`func (o *AssistantConfigBody) SetBudgetLimitMaxCredits(v int64)`

SetBudgetLimitMaxCredits sets BudgetLimitMaxCredits field to given value.


### SetBudgetLimitMaxCreditsNil

`func (o *AssistantConfigBody) SetBudgetLimitMaxCreditsNil(b bool)`

 SetBudgetLimitMaxCreditsNil sets the value for BudgetLimitMaxCredits to be an explicit nil

### UnsetBudgetLimitMaxCredits
`func (o *AssistantConfigBody) UnsetBudgetLimitMaxCredits()`

UnsetBudgetLimitMaxCredits ensures that no value is present for BudgetLimitMaxCredits, not even an explicit nil
### GetBudgetLimitMinCredits

`func (o *AssistantConfigBody) GetBudgetLimitMinCredits() int64`

GetBudgetLimitMinCredits returns the BudgetLimitMinCredits field if non-nil, zero value otherwise.

### GetBudgetLimitMinCreditsOk

`func (o *AssistantConfigBody) GetBudgetLimitMinCreditsOk() (*int64, bool)`

GetBudgetLimitMinCreditsOk returns a tuple with the BudgetLimitMinCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLimitMinCredits

`func (o *AssistantConfigBody) SetBudgetLimitMinCredits(v int64)`

SetBudgetLimitMinCredits sets BudgetLimitMinCredits field to given value.


### SetBudgetLimitMinCreditsNil

`func (o *AssistantConfigBody) SetBudgetLimitMinCreditsNil(b bool)`

 SetBudgetLimitMinCreditsNil sets the value for BudgetLimitMinCredits to be an explicit nil

### UnsetBudgetLimitMinCredits
`func (o *AssistantConfigBody) UnsetBudgetLimitMinCredits()`

UnsetBudgetLimitMinCredits ensures that no value is present for BudgetLimitMinCredits, not even an explicit nil
### GetModels

`func (o *AssistantConfigBody) GetModels() []IDNameStringString`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AssistantConfigBody) GetModelsOk() (*[]IDNameStringString, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AssistantConfigBody) SetModels(v []IDNameStringString)`

SetModels sets Models field to given value.


### GetPermittedModels

`func (o *AssistantConfigBody) GetPermittedModels() []string`

GetPermittedModels returns the PermittedModels field if non-nil, zero value otherwise.

### GetPermittedModelsOk

`func (o *AssistantConfigBody) GetPermittedModelsOk() (*[]string, bool)`

GetPermittedModelsOk returns a tuple with the PermittedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedModels

`func (o *AssistantConfigBody) SetPermittedModels(v []string)`

SetPermittedModels sets PermittedModels field to given value.


### GetPricePerCreditEur

`func (o *AssistantConfigBody) GetPricePerCreditEur() float64`

GetPricePerCreditEur returns the PricePerCreditEur field if non-nil, zero value otherwise.

### GetPricePerCreditEurOk

`func (o *AssistantConfigBody) GetPricePerCreditEurOk() (*float64, bool)`

GetPricePerCreditEurOk returns a tuple with the PricePerCreditEur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCreditEur

`func (o *AssistantConfigBody) SetPricePerCreditEur(v float64)`

SetPricePerCreditEur sets PricePerCreditEur field to given value.


### GetSpentCredits

`func (o *AssistantConfigBody) GetSpentCredits() int64`

GetSpentCredits returns the SpentCredits field if non-nil, zero value otherwise.

### GetSpentCreditsOk

`func (o *AssistantConfigBody) GetSpentCreditsOk() (*int64, bool)`

GetSpentCreditsOk returns a tuple with the SpentCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentCredits

`func (o *AssistantConfigBody) SetSpentCredits(v int64)`

SetSpentCredits sets SpentCredits field to given value.


### SetSpentCreditsNil

`func (o *AssistantConfigBody) SetSpentCreditsNil(b bool)`

 SetSpentCreditsNil sets the value for SpentCredits to be an explicit nil

### UnsetSpentCredits
`func (o *AssistantConfigBody) UnsetSpentCredits()`

UnsetSpentCredits ensures that no value is present for SpentCredits, not even an explicit nil
### GetToolsEnabled

`func (o *AssistantConfigBody) GetToolsEnabled() bool`

GetToolsEnabled returns the ToolsEnabled field if non-nil, zero value otherwise.

### GetToolsEnabledOk

`func (o *AssistantConfigBody) GetToolsEnabledOk() (*bool, bool)`

GetToolsEnabledOk returns a tuple with the ToolsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsEnabled

`func (o *AssistantConfigBody) SetToolsEnabled(v bool)`

SetToolsEnabled sets ToolsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


