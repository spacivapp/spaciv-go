# LogicLayerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]LogicRule**](LogicRule.md) |  | 
**Type** | **string** |  | 

## Methods

### NewLogicLayerConfig

`func NewLogicLayerConfig(rules []LogicRule, type_ string, ) *LogicLayerConfig`

NewLogicLayerConfig instantiates a new LogicLayerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicLayerConfigWithDefaults

`func NewLogicLayerConfigWithDefaults() *LogicLayerConfig`

NewLogicLayerConfigWithDefaults instantiates a new LogicLayerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *LogicLayerConfig) GetRules() []LogicRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *LogicLayerConfig) GetRulesOk() (*[]LogicRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *LogicLayerConfig) SetRules(v []LogicRule)`

SetRules sets Rules field to given value.


### GetType

`func (o *LogicLayerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogicLayerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogicLayerConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


