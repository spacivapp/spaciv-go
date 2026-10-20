# ScratchpadOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Why some of the output could not be collected, for example because it was over a limit. | [optional] 
**Files** | [**[]ScratchpadOutputFile**](ScratchpadOutputFile.md) | Files the handler wrote to its output directory. | 
**Result** | Pointer to **interface{}** |  | [optional] 
**Stderr** | **string** |  | 
**StderrTruncated** | **bool** | Set when stderr was longer than the 256 KiB kept. | 
**Stdout** | **string** |  | 
**StdoutTruncated** | **bool** | Set when stdout was longer than the 256 KiB kept. | 

## Methods

### NewScratchpadOutput

`func NewScratchpadOutput(files []ScratchpadOutputFile, stderr string, stderrTruncated bool, stdout string, stdoutTruncated bool, ) *ScratchpadOutput`

NewScratchpadOutput instantiates a new ScratchpadOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadOutputWithDefaults

`func NewScratchpadOutputWithDefaults() *ScratchpadOutput`

NewScratchpadOutputWithDefaults instantiates a new ScratchpadOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ScratchpadOutput) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScratchpadOutput) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScratchpadOutput) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScratchpadOutput) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFiles

`func (o *ScratchpadOutput) GetFiles() []ScratchpadOutputFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScratchpadOutput) GetFilesOk() (*[]ScratchpadOutputFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScratchpadOutput) SetFiles(v []ScratchpadOutputFile)`

SetFiles sets Files field to given value.


### GetResult

`func (o *ScratchpadOutput) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ScratchpadOutput) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ScratchpadOutput) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *ScratchpadOutput) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ScratchpadOutput) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ScratchpadOutput) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetStderr

`func (o *ScratchpadOutput) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *ScratchpadOutput) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *ScratchpadOutput) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStderrTruncated

`func (o *ScratchpadOutput) GetStderrTruncated() bool`

GetStderrTruncated returns the StderrTruncated field if non-nil, zero value otherwise.

### GetStderrTruncatedOk

`func (o *ScratchpadOutput) GetStderrTruncatedOk() (*bool, bool)`

GetStderrTruncatedOk returns a tuple with the StderrTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderrTruncated

`func (o *ScratchpadOutput) SetStderrTruncated(v bool)`

SetStderrTruncated sets StderrTruncated field to given value.


### GetStdout

`func (o *ScratchpadOutput) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *ScratchpadOutput) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *ScratchpadOutput) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStdoutTruncated

`func (o *ScratchpadOutput) GetStdoutTruncated() bool`

GetStdoutTruncated returns the StdoutTruncated field if non-nil, zero value otherwise.

### GetStdoutTruncatedOk

`func (o *ScratchpadOutput) GetStdoutTruncatedOk() (*bool, bool)`

GetStdoutTruncatedOk returns a tuple with the StdoutTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdoutTruncated

`func (o *ScratchpadOutput) SetStdoutTruncated(v bool)`

SetStdoutTruncated sets StdoutTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


