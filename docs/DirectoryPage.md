# DirectoryPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | [**[]Directory**](Directory.md) |  | 
**NextCursor** | Pointer to **string** | Pass back as cursor to fetch the following page. Absent on the last page. | [optional] 

## Methods

### NewDirectoryPage

`func NewDirectoryPage(directories []Directory, ) *DirectoryPage`

NewDirectoryPage instantiates a new DirectoryPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryPageWithDefaults

`func NewDirectoryPageWithDefaults() *DirectoryPage`

NewDirectoryPageWithDefaults instantiates a new DirectoryPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *DirectoryPage) GetDirectories() []Directory`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *DirectoryPage) GetDirectoriesOk() (*[]Directory, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *DirectoryPage) SetDirectories(v []Directory)`

SetDirectories sets Directories field to given value.


### GetNextCursor

`func (o *DirectoryPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DirectoryPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DirectoryPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *DirectoryPage) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


