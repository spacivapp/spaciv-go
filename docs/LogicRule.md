# LogicRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Demand** | [**NullableLogicDemand**](LogicDemand.md) |  | 
**ExactNodes** | **bool** |  | 
**Name** | **string** |  | 
**Output** | [**[]LogicAssignment**](LogicAssignment.md) |  | 
**Per** | [**NullableLogicPer**](LogicPer.md) |  | 
**Supply** | [**NullableLogicSupply**](LogicSupply.md) |  | 

## Methods

### NewLogicRule

`func NewLogicRule(demand NullableLogicDemand, exactNodes bool, name string, output []LogicAssignment, per NullableLogicPer, supply NullableLogicSupply, ) *LogicRule`

NewLogicRule instantiates a new LogicRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicRuleWithDefaults

`func NewLogicRuleWithDefaults() *LogicRule`

NewLogicRuleWithDefaults instantiates a new LogicRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemand

`func (o *LogicRule) GetDemand() LogicDemand`

GetDemand returns the Demand field if non-nil, zero value otherwise.

### GetDemandOk

`func (o *LogicRule) GetDemandOk() (*LogicDemand, bool)`

GetDemandOk returns a tuple with the Demand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemand

`func (o *LogicRule) SetDemand(v LogicDemand)`

SetDemand sets Demand field to given value.


### SetDemandNil

`func (o *LogicRule) SetDemandNil(b bool)`

 SetDemandNil sets the value for Demand to be an explicit nil

### UnsetDemand
`func (o *LogicRule) UnsetDemand()`

UnsetDemand ensures that no value is present for Demand, not even an explicit nil
### GetExactNodes

`func (o *LogicRule) GetExactNodes() bool`

GetExactNodes returns the ExactNodes field if non-nil, zero value otherwise.

### GetExactNodesOk

`func (o *LogicRule) GetExactNodesOk() (*bool, bool)`

GetExactNodesOk returns a tuple with the ExactNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNodes

`func (o *LogicRule) SetExactNodes(v bool)`

SetExactNodes sets ExactNodes field to given value.


### GetName

`func (o *LogicRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicRule) SetName(v string)`

SetName sets Name field to given value.


### GetOutput

`func (o *LogicRule) GetOutput() []LogicAssignment`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *LogicRule) GetOutputOk() (*[]LogicAssignment, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *LogicRule) SetOutput(v []LogicAssignment)`

SetOutput sets Output field to given value.


### GetPer

`func (o *LogicRule) GetPer() LogicPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *LogicRule) GetPerOk() (*LogicPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *LogicRule) SetPer(v LogicPer)`

SetPer sets Per field to given value.


### SetPerNil

`func (o *LogicRule) SetPerNil(b bool)`

 SetPerNil sets the value for Per to be an explicit nil

### UnsetPer
`func (o *LogicRule) UnsetPer()`

UnsetPer ensures that no value is present for Per, not even an explicit nil
### GetSupply

`func (o *LogicRule) GetSupply() LogicSupply`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *LogicRule) GetSupplyOk() (*LogicSupply, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *LogicRule) SetSupply(v LogicSupply)`

SetSupply sets Supply field to given value.


### SetSupplyNil

`func (o *LogicRule) SetSupplyNil(b bool)`

 SetSupplyNil sets the value for Supply to be an explicit nil

### UnsetSupply
`func (o *LogicRule) UnsetSupply()`

UnsetSupply ensures that no value is present for Supply, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


