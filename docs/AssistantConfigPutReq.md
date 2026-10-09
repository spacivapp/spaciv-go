# AssistantConfigPutReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetLimit** | Pointer to [**AssistantBudgetLimit**](AssistantBudgetLimit.md) |  | [optional] 
**LlmModel** | Pointer to **string** |  | [optional] 
**PermittedModels** | **[]string** |  | 
**ToolsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAssistantConfigPutReq

`func NewAssistantConfigPutReq(permittedModels []string, ) *AssistantConfigPutReq`

NewAssistantConfigPutReq instantiates a new AssistantConfigPutReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantConfigPutReqWithDefaults

`func NewAssistantConfigPutReqWithDefaults() *AssistantConfigPutReq`

NewAssistantConfigPutReqWithDefaults instantiates a new AssistantConfigPutReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetLimit

`func (o *AssistantConfigPutReq) GetBudgetLimit() AssistantBudgetLimit`

GetBudgetLimit returns the BudgetLimit field if non-nil, zero value otherwise.

### GetBudgetLimitOk

`func (o *AssistantConfigPutReq) GetBudgetLimitOk() (*AssistantBudgetLimit, bool)`

GetBudgetLimitOk returns a tuple with the BudgetLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLimit

`func (o *AssistantConfigPutReq) SetBudgetLimit(v AssistantBudgetLimit)`

SetBudgetLimit sets BudgetLimit field to given value.

### HasBudgetLimit

`func (o *AssistantConfigPutReq) HasBudgetLimit() bool`

HasBudgetLimit returns a boolean if a field has been set.

### GetLlmModel

`func (o *AssistantConfigPutReq) GetLlmModel() string`

GetLlmModel returns the LlmModel field if non-nil, zero value otherwise.

### GetLlmModelOk

`func (o *AssistantConfigPutReq) GetLlmModelOk() (*string, bool)`

GetLlmModelOk returns a tuple with the LlmModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmModel

`func (o *AssistantConfigPutReq) SetLlmModel(v string)`

SetLlmModel sets LlmModel field to given value.

### HasLlmModel

`func (o *AssistantConfigPutReq) HasLlmModel() bool`

HasLlmModel returns a boolean if a field has been set.

### GetPermittedModels

`func (o *AssistantConfigPutReq) GetPermittedModels() []string`

GetPermittedModels returns the PermittedModels field if non-nil, zero value otherwise.

### GetPermittedModelsOk

`func (o *AssistantConfigPutReq) GetPermittedModelsOk() (*[]string, bool)`

GetPermittedModelsOk returns a tuple with the PermittedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedModels

`func (o *AssistantConfigPutReq) SetPermittedModels(v []string)`

SetPermittedModels sets PermittedModels field to given value.


### GetToolsEnabled

`func (o *AssistantConfigPutReq) GetToolsEnabled() bool`

GetToolsEnabled returns the ToolsEnabled field if non-nil, zero value otherwise.

### GetToolsEnabledOk

`func (o *AssistantConfigPutReq) GetToolsEnabledOk() (*bool, bool)`

GetToolsEnabledOk returns a tuple with the ToolsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsEnabled

`func (o *AssistantConfigPutReq) SetToolsEnabled(v bool)`

SetToolsEnabled sets ToolsEnabled field to given value.

### HasToolsEnabled

`func (o *AssistantConfigPutReq) HasToolsEnabled() bool`

HasToolsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


