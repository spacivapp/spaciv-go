# Orphan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **int64** |  | 
**Action** | **string** |  | 
**FileId** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Size** | **int64** |  | 

## Methods

### NewOrphan

`func NewOrphan(account int64, action string, fileId string, lastModified time.Time, size int64, ) *Orphan`

NewOrphan instantiates a new Orphan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrphanWithDefaults

`func NewOrphanWithDefaults() *Orphan`

NewOrphanWithDefaults instantiates a new Orphan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Orphan) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Orphan) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Orphan) SetAccount(v int64)`

SetAccount sets Account field to given value.


### GetAction

`func (o *Orphan) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Orphan) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Orphan) SetAction(v string)`

SetAction sets Action field to given value.


### GetFileId

`func (o *Orphan) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Orphan) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Orphan) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetLastModified

`func (o *Orphan) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Orphan) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Orphan) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetSize

`func (o *Orphan) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Orphan) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Orphan) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


