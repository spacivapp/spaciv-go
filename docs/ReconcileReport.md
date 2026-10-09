# ReconcileReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **int64** |  | 
**DanglingRecords** | [**[]DanglingRecord**](DanglingRecord.md) |  | 
**Orphans** | [**[]Orphan**](Orphan.md) |  | 
**UnexpectedKeys** | **[]string** |  | 

## Methods

### NewReconcileReport

`func NewReconcileReport(accounts int64, danglingRecords []DanglingRecord, orphans []Orphan, unexpectedKeys []string, ) *ReconcileReport`

NewReconcileReport instantiates a new ReconcileReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconcileReportWithDefaults

`func NewReconcileReportWithDefaults() *ReconcileReport`

NewReconcileReportWithDefaults instantiates a new ReconcileReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ReconcileReport) GetAccounts() int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ReconcileReport) GetAccountsOk() (*int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ReconcileReport) SetAccounts(v int64)`

SetAccounts sets Accounts field to given value.


### GetDanglingRecords

`func (o *ReconcileReport) GetDanglingRecords() []DanglingRecord`

GetDanglingRecords returns the DanglingRecords field if non-nil, zero value otherwise.

### GetDanglingRecordsOk

`func (o *ReconcileReport) GetDanglingRecordsOk() (*[]DanglingRecord, bool)`

GetDanglingRecordsOk returns a tuple with the DanglingRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDanglingRecords

`func (o *ReconcileReport) SetDanglingRecords(v []DanglingRecord)`

SetDanglingRecords sets DanglingRecords field to given value.


### GetOrphans

`func (o *ReconcileReport) GetOrphans() []Orphan`

GetOrphans returns the Orphans field if non-nil, zero value otherwise.

### GetOrphansOk

`func (o *ReconcileReport) GetOrphansOk() (*[]Orphan, bool)`

GetOrphansOk returns a tuple with the Orphans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphans

`func (o *ReconcileReport) SetOrphans(v []Orphan)`

SetOrphans sets Orphans field to given value.


### GetUnexpectedKeys

`func (o *ReconcileReport) GetUnexpectedKeys() []string`

GetUnexpectedKeys returns the UnexpectedKeys field if non-nil, zero value otherwise.

### GetUnexpectedKeysOk

`func (o *ReconcileReport) GetUnexpectedKeysOk() (*[]string, bool)`

GetUnexpectedKeysOk returns a tuple with the UnexpectedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedKeys

`func (o *ReconcileReport) SetUnexpectedKeys(v []string)`

SetUnexpectedKeys sets UnexpectedKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


