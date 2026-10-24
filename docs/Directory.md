# Directory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When the directory was created (RFC 3339). | 
**CreatedBy** | **int64** | Identifier of the user who created the directory. | 
**Id** | **string** | Opaque identifier of the directory. | 
**Name** | **string** | Display name of the directory. Nothing else in the same parent has the same name, compared exactly and after normalising to Unicode NFC. | 
**ParentId** | Pointer to **string** | Identifier of the directory this one is in. Absent for a directory at the root. | [optional] 
**Path** | **string** | Where the directory is: the names of the directories it is in, from the root, then its own name, separated by /, such as Leases/2026. It changes when the directory or one it is in is renamed or moved; the identifier does not. | 
**TagProblems** | Pointer to **map[string][]string** | Problems found with the directory by its reserved tags&#39; checks, by tag. Absent when there are none. | [optional] 
**Tags** | Pointer to **[]string** | The directory&#39;s tags, in alphabetical order. Absent when the directory has none. | [optional] 

## Methods

### NewDirectory

`func NewDirectory(createdAt time.Time, createdBy int64, id string, name string, path string, ) *Directory`

NewDirectory instantiates a new Directory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryWithDefaults

`func NewDirectoryWithDefaults() *Directory`

NewDirectoryWithDefaults instantiates a new Directory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Directory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Directory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Directory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Directory) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Directory) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Directory) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *Directory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Directory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Directory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Directory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Directory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Directory) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Directory) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Directory) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Directory) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Directory) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPath

`func (o *Directory) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Directory) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Directory) SetPath(v string)`

SetPath sets Path field to given value.


### GetTagProblems

`func (o *Directory) GetTagProblems() map[string][]string`

GetTagProblems returns the TagProblems field if non-nil, zero value otherwise.

### GetTagProblemsOk

`func (o *Directory) GetTagProblemsOk() (*map[string][]string, bool)`

GetTagProblemsOk returns a tuple with the TagProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagProblems

`func (o *Directory) SetTagProblems(v map[string][]string)`

SetTagProblems sets TagProblems field to given value.

### HasTagProblems

`func (o *Directory) HasTagProblems() bool`

HasTagProblems returns a boolean if a field has been set.

### GetTags

`func (o *Directory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Directory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Directory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Directory) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


