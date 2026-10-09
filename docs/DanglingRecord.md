# DanglingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **int64** |  | 
**FileId** | **string** |  | 
**Size** | **int64** |  | 

## Methods

### NewDanglingRecord

`func NewDanglingRecord(account int64, fileId string, size int64, ) *DanglingRecord`

NewDanglingRecord instantiates a new DanglingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDanglingRecordWithDefaults

`func NewDanglingRecordWithDefaults() *DanglingRecord`

NewDanglingRecordWithDefaults instantiates a new DanglingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DanglingRecord) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DanglingRecord) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DanglingRecord) SetAccount(v int64)`

SetAccount sets Account field to given value.


### GetFileId

`func (o *DanglingRecord) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DanglingRecord) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DanglingRecord) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetSize

`func (o *DanglingRecord) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DanglingRecord) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DanglingRecord) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


