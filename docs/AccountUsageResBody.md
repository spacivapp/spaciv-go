# AccountUsageResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveObjectCount** | **int64** | Distinct objects in at least one of the account&#39;s active (non-archived) object groups. An object in several groups counts once. | 
**ObjectCount** | **int64** | Distinct objects across all of the account&#39;s object groups, archived groups included. An object in several groups counts once. | 
**ObjectCountsCalculating** | **bool** | True while the account&#39;s object count totals are first being built. | 

## Methods

### NewAccountUsageResBody

`func NewAccountUsageResBody(activeObjectCount int64, objectCount int64, objectCountsCalculating bool, ) *AccountUsageResBody`

NewAccountUsageResBody instantiates a new AccountUsageResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUsageResBodyWithDefaults

`func NewAccountUsageResBodyWithDefaults() *AccountUsageResBody`

NewAccountUsageResBodyWithDefaults instantiates a new AccountUsageResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveObjectCount

`func (o *AccountUsageResBody) GetActiveObjectCount() int64`

GetActiveObjectCount returns the ActiveObjectCount field if non-nil, zero value otherwise.

### GetActiveObjectCountOk

`func (o *AccountUsageResBody) GetActiveObjectCountOk() (*int64, bool)`

GetActiveObjectCountOk returns a tuple with the ActiveObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveObjectCount

`func (o *AccountUsageResBody) SetActiveObjectCount(v int64)`

SetActiveObjectCount sets ActiveObjectCount field to given value.


### GetObjectCount

`func (o *AccountUsageResBody) GetObjectCount() int64`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *AccountUsageResBody) GetObjectCountOk() (*int64, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *AccountUsageResBody) SetObjectCount(v int64)`

SetObjectCount sets ObjectCount field to given value.


### GetObjectCountsCalculating

`func (o *AccountUsageResBody) GetObjectCountsCalculating() bool`

GetObjectCountsCalculating returns the ObjectCountsCalculating field if non-nil, zero value otherwise.

### GetObjectCountsCalculatingOk

`func (o *AccountUsageResBody) GetObjectCountsCalculatingOk() (*bool, bool)`

GetObjectCountsCalculatingOk returns a tuple with the ObjectCountsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCountsCalculating

`func (o *AccountUsageResBody) SetObjectCountsCalculating(v bool)`

SetObjectCountsCalculating sets ObjectCountsCalculating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


