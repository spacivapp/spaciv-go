# AssistantModelListBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultModel** | **string** |  | 
**Models** | [**[]IDNameStringString**](IDNameStringString.md) |  | 

## Methods

### NewAssistantModelListBody

`func NewAssistantModelListBody(defaultModel string, models []IDNameStringString, ) *AssistantModelListBody`

NewAssistantModelListBody instantiates a new AssistantModelListBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantModelListBodyWithDefaults

`func NewAssistantModelListBodyWithDefaults() *AssistantModelListBody`

NewAssistantModelListBodyWithDefaults instantiates a new AssistantModelListBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultModel

`func (o *AssistantModelListBody) GetDefaultModel() string`

GetDefaultModel returns the DefaultModel field if non-nil, zero value otherwise.

### GetDefaultModelOk

`func (o *AssistantModelListBody) GetDefaultModelOk() (*string, bool)`

GetDefaultModelOk returns a tuple with the DefaultModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModel

`func (o *AssistantModelListBody) SetDefaultModel(v string)`

SetDefaultModel sets DefaultModel field to given value.


### GetModels

`func (o *AssistantModelListBody) GetModels() []IDNameStringString`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AssistantModelListBody) GetModelsOk() (*[]IDNameStringString, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AssistantModelListBody) SetModels(v []IDNameStringString)`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


