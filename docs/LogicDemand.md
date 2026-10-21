# LogicDemand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 

## Methods

### NewLogicDemand

`func NewLogicDemand(amount NullableCalcRuleRef, filters []FilterGroup, ) *LogicDemand`

NewLogicDemand instantiates a new LogicDemand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicDemandWithDefaults

`func NewLogicDemandWithDefaults() *LogicDemand`

NewLogicDemandWithDefaults instantiates a new LogicDemand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *LogicDemand) GetAmount() CalcRuleRef`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LogicDemand) GetAmountOk() (*CalcRuleRef, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LogicDemand) SetAmount(v CalcRuleRef)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *LogicDemand) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *LogicDemand) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetFilters

`func (o *LogicDemand) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LogicDemand) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LogicDemand) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


