# InstalledSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | [**[]Directory**](Directory.md) | The directories inside the skill&#39;s directory that the load needed. | 
**Directory** | [**Directory**](Directory.md) |  | 
**Files** | [**[]Record**](Record.md) | Every file the load wrote into the skill&#39;s directory, at any depth. | 
**RemovedFileIds** | **[]string** | The files the load deleted: those it replaced, and those the built-in skill doesn&#39;t have. | 

## Methods

### NewInstalledSkill

`func NewInstalledSkill(directories []Directory, directory Directory, files []Record, removedFileIds []string, ) *InstalledSkill`

NewInstalledSkill instantiates a new InstalledSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstalledSkillWithDefaults

`func NewInstalledSkillWithDefaults() *InstalledSkill`

NewInstalledSkillWithDefaults instantiates a new InstalledSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *InstalledSkill) GetDirectories() []Directory`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *InstalledSkill) GetDirectoriesOk() (*[]Directory, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *InstalledSkill) SetDirectories(v []Directory)`

SetDirectories sets Directories field to given value.


### GetDirectory

`func (o *InstalledSkill) GetDirectory() Directory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *InstalledSkill) GetDirectoryOk() (*Directory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *InstalledSkill) SetDirectory(v Directory)`

SetDirectory sets Directory field to given value.


### GetFiles

`func (o *InstalledSkill) GetFiles() []Record`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *InstalledSkill) GetFilesOk() (*[]Record, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *InstalledSkill) SetFiles(v []Record)`

SetFiles sets Files field to given value.


### GetRemovedFileIds

`func (o *InstalledSkill) GetRemovedFileIds() []string`

GetRemovedFileIds returns the RemovedFileIds field if non-nil, zero value otherwise.

### GetRemovedFileIdsOk

`func (o *InstalledSkill) GetRemovedFileIdsOk() (*[]string, bool)`

GetRemovedFileIdsOk returns a tuple with the RemovedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFileIds

`func (o *InstalledSkill) SetRemovedFileIds(v []string)`

SetRemovedFileIds sets RemovedFileIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


