# AccountCalculationUsageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]AccountCalculationUsageDay**](AccountCalculationUsageDay.md) | One entry per day, oldest first, including days with no usage. | 

## Methods

### NewAccountCalculationUsageBody

`func NewAccountCalculationUsageBody(days []AccountCalculationUsageDay, ) *AccountCalculationUsageBody`

NewAccountCalculationUsageBody instantiates a new AccountCalculationUsageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCalculationUsageBodyWithDefaults

`func NewAccountCalculationUsageBodyWithDefaults() *AccountCalculationUsageBody`

NewAccountCalculationUsageBodyWithDefaults instantiates a new AccountCalculationUsageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *AccountCalculationUsageBody) GetDays() []AccountCalculationUsageDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *AccountCalculationUsageBody) GetDaysOk() (*[]AccountCalculationUsageDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *AccountCalculationUsageBody) SetDays(v []AccountCalculationUsageDay)`

SetDays sets Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


