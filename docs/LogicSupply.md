# LogicSupply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FilterGroup**](FilterGroup.md) |  | 
**IgnoreMissingWeights** | **bool** |  | 
**Weights** | [**map[string]LogicSupplyWeight**](LogicSupplyWeight.md) |  | 

## Methods

### NewLogicSupply

`func NewLogicSupply(filters []FilterGroup, ignoreMissingWeights bool, weights map[string]LogicSupplyWeight, ) *LogicSupply`

NewLogicSupply instantiates a new LogicSupply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicSupplyWithDefaults

`func NewLogicSupplyWithDefaults() *LogicSupply`

NewLogicSupplyWithDefaults instantiates a new LogicSupply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *LogicSupply) GetFilters() []FilterGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LogicSupply) GetFiltersOk() (*[]FilterGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LogicSupply) SetFilters(v []FilterGroup)`

SetFilters sets Filters field to given value.


### GetIgnoreMissingWeights

`func (o *LogicSupply) GetIgnoreMissingWeights() bool`

GetIgnoreMissingWeights returns the IgnoreMissingWeights field if non-nil, zero value otherwise.

### GetIgnoreMissingWeightsOk

`func (o *LogicSupply) GetIgnoreMissingWeightsOk() (*bool, bool)`

GetIgnoreMissingWeightsOk returns a tuple with the IgnoreMissingWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMissingWeights

`func (o *LogicSupply) SetIgnoreMissingWeights(v bool)`

SetIgnoreMissingWeights sets IgnoreMissingWeights field to given value.


### GetWeights

`func (o *LogicSupply) GetWeights() map[string]LogicSupplyWeight`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *LogicSupply) GetWeightsOk() (*map[string]LogicSupplyWeight, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *LogicSupply) SetWeights(v map[string]LogicSupplyWeight)`

SetWeights sets Weights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


