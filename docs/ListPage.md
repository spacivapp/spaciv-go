# ListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | Pointer to [**[]Directory**](Directory.md) | Directories directly inside the listed directory. Returned only when a parent is given, and only on the pages that hold some. | [optional] 
**Files** | [**[]Record**](Record.md) |  | 
**NextCursor** | Pointer to **string** | Pass back as cursor to fetch the following page. Absent on the last page. | [optional] 

## Methods

### NewListPage

`func NewListPage(files []Record, ) *ListPage`

NewListPage instantiates a new ListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPageWithDefaults

`func NewListPageWithDefaults() *ListPage`

NewListPageWithDefaults instantiates a new ListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *ListPage) GetDirectories() []Directory`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *ListPage) GetDirectoriesOk() (*[]Directory, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *ListPage) SetDirectories(v []Directory)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *ListPage) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### GetFiles

`func (o *ListPage) GetFiles() []Record`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ListPage) GetFilesOk() (*[]Record, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ListPage) SetFiles(v []Record)`

SetFiles sets Files field to given value.


### GetNextCursor

`func (o *ListPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListPage) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


