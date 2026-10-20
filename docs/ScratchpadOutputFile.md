# ScratchpadOutputFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The file&#39;s content, base64-encoded. | 
**ContentType** | Pointer to **string** |  | [optional] 
**Path** | **string** | Path relative to the output directory. | 

## Methods

### NewScratchpadOutputFile

`func NewScratchpadOutputFile(content string, path string, ) *ScratchpadOutputFile`

NewScratchpadOutputFile instantiates a new ScratchpadOutputFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadOutputFileWithDefaults

`func NewScratchpadOutputFileWithDefaults() *ScratchpadOutputFile`

NewScratchpadOutputFileWithDefaults instantiates a new ScratchpadOutputFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ScratchpadOutputFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ScratchpadOutputFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ScratchpadOutputFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *ScratchpadOutputFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ScratchpadOutputFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ScratchpadOutputFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ScratchpadOutputFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetPath

`func (o *ScratchpadOutputFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ScratchpadOutputFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ScratchpadOutputFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


