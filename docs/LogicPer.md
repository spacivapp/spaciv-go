# LogicPer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | **int32** |  | 
**Demand** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 
**Supply** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewLogicPer

`func NewLogicPer(basis int32, demand NullableCalcRuleRef, supply NullableCalcRuleRef, ) *LogicPer`

NewLogicPer instantiates a new LogicPer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicPerWithDefaults

`func NewLogicPerWithDefaults() *LogicPer`

NewLogicPerWithDefaults instantiates a new LogicPer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *LogicPer) GetBasis() int32`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *LogicPer) GetBasisOk() (*int32, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *LogicPer) SetBasis(v int32)`

SetBasis sets Basis field to given value.


### GetDemand

`func (o *LogicPer) GetDemand() CalcRuleRef`

GetDemand returns the Demand field if non-nil, zero value otherwise.

### GetDemandOk

`func (o *LogicPer) GetDemandOk() (*CalcRuleRef, bool)`

GetDemandOk returns a tuple with the Demand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemand

`func (o *LogicPer) SetDemand(v CalcRuleRef)`

SetDemand sets Demand field to given value.


### SetDemandNil

`func (o *LogicPer) SetDemandNil(b bool)`

 SetDemandNil sets the value for Demand to be an explicit nil

### UnsetDemand
`func (o *LogicPer) UnsetDemand()`

UnsetDemand ensures that no value is present for Demand, not even an explicit nil
### GetSupply

`func (o *LogicPer) GetSupply() CalcRuleRef`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *LogicPer) GetSupplyOk() (*CalcRuleRef, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *LogicPer) SetSupply(v CalcRuleRef)`

SetSupply sets Supply field to given value.


### SetSupplyNil

`func (o *LogicPer) SetSupplyNil(b bool)`

 SetSupplyNil sets the value for Supply to be an explicit nil

### UnsetSupply
`func (o *LogicPer) UnsetSupply()`

UnsetSupply ensures that no value is present for Supply, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


