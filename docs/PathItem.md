# PathItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to [**Directory**](Directory.md) | The directory at the path. Absent when a file is there. | [optional] 
**File** | Pointer to [**Record**](Record.md) | The file at the path. Absent when a directory is there. | [optional] 

## Methods

### NewPathItem

`func NewPathItem() *PathItem`

NewPathItem instantiates a new PathItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathItemWithDefaults

`func NewPathItemWithDefaults() *PathItem`

NewPathItemWithDefaults instantiates a new PathItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *PathItem) GetDirectory() Directory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *PathItem) GetDirectoryOk() (*Directory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *PathItem) SetDirectory(v Directory)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *PathItem) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetFile

`func (o *PathItem) GetFile() Record`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PathItem) GetFileOk() (*Record, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PathItem) SetFile(v Record)`

SetFile sets File field to given value.

### HasFile

`func (o *PathItem) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


