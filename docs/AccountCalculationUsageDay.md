# AccountCalculationUsageDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **int64** | Credits used aggregating: one per 10,000 objects read by a sum. | 
**Ai** | **int64** | Credits used by the AI assistant. | 
**Build** | **int64** | Credits used building layer data: one per 1,000 objects written. | 
**Date** | **string** | UTC day the usage was recorded on. | 

## Methods

### NewAccountCalculationUsageDay

`func NewAccountCalculationUsageDay(aggregation int64, ai int64, build int64, date string, ) *AccountCalculationUsageDay`

NewAccountCalculationUsageDay instantiates a new AccountCalculationUsageDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCalculationUsageDayWithDefaults

`func NewAccountCalculationUsageDayWithDefaults() *AccountCalculationUsageDay`

NewAccountCalculationUsageDayWithDefaults instantiates a new AccountCalculationUsageDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *AccountCalculationUsageDay) GetAggregation() int64`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *AccountCalculationUsageDay) GetAggregationOk() (*int64, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *AccountCalculationUsageDay) SetAggregation(v int64)`

SetAggregation sets Aggregation field to given value.


### GetAi

`func (o *AccountCalculationUsageDay) GetAi() int64`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *AccountCalculationUsageDay) GetAiOk() (*int64, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *AccountCalculationUsageDay) SetAi(v int64)`

SetAi sets Ai field to given value.


### GetBuild

`func (o *AccountCalculationUsageDay) GetBuild() int64`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *AccountCalculationUsageDay) GetBuildOk() (*int64, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *AccountCalculationUsageDay) SetBuild(v int64)`

SetBuild sets Build field to given value.


### GetDate

`func (o *AccountCalculationUsageDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountCalculationUsageDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountCalculationUsageDay) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


