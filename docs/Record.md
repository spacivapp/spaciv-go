# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When the file was uploaded (RFC 3339). | 
**Id** | **string** | Opaque identifier of the file, used to rename or delete it. | 
**MimeType** | **string** | Media type detected from the file&#39;s contents at upload time, rather than from its name or the type the client declared. | 
**Name** | **string** | Display name of the file, taken from the uploaded file name and changeable afterwards. | 
**QuarantineSignature** | Pointer to **string** | Name of the virus signature the file matched when it was quarantined. | [optional] 
**QuarantinedAt** | Pointer to **time.Time** | When a rescan found the file matching a virus signature (RFC 3339). A quarantined file cannot be downloaded, but it can still be renamed or deleted. | [optional] 
**Size** | **int64** | Size of the stored file in bytes. | 
**UploadedBy** | **int64** | Identifier of the user who uploaded the file. | 

## Methods

### NewRecord

`func NewRecord(createdAt time.Time, id string, mimeType string, name string, size int64, uploadedBy int64, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Record) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Record) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Record) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *Record) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Record) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Record) SetId(v string)`

SetId sets Id field to given value.


### GetMimeType

`func (o *Record) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Record) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Record) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetName

`func (o *Record) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Record) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Record) SetName(v string)`

SetName sets Name field to given value.


### GetQuarantineSignature

`func (o *Record) GetQuarantineSignature() string`

GetQuarantineSignature returns the QuarantineSignature field if non-nil, zero value otherwise.

### GetQuarantineSignatureOk

`func (o *Record) GetQuarantineSignatureOk() (*string, bool)`

GetQuarantineSignatureOk returns a tuple with the QuarantineSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSignature

`func (o *Record) SetQuarantineSignature(v string)`

SetQuarantineSignature sets QuarantineSignature field to given value.

### HasQuarantineSignature

`func (o *Record) HasQuarantineSignature() bool`

HasQuarantineSignature returns a boolean if a field has been set.

### GetQuarantinedAt

`func (o *Record) GetQuarantinedAt() time.Time`

GetQuarantinedAt returns the QuarantinedAt field if non-nil, zero value otherwise.

### GetQuarantinedAtOk

`func (o *Record) GetQuarantinedAtOk() (*time.Time, bool)`

GetQuarantinedAtOk returns a tuple with the QuarantinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantinedAt

`func (o *Record) SetQuarantinedAt(v time.Time)`

SetQuarantinedAt sets QuarantinedAt field to given value.

### HasQuarantinedAt

`func (o *Record) HasQuarantinedAt() bool`

HasQuarantinedAt returns a boolean if a field has been set.

### GetSize

`func (o *Record) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Record) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Record) SetSize(v int64)`

SetSize sets Size field to given value.


### GetUploadedBy

`func (o *Record) GetUploadedBy() int64`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *Record) GetUploadedByOk() (*int64, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *Record) SetUploadedBy(v int64)`

SetUploadedBy sets UploadedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


