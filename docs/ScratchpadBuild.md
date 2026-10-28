# ScratchpadBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The build&#39;s id, to give as buildId when running it. | 
**Message** | Pointer to **string** | Why the build failed. | [optional] 
**Output** | Pointer to **string** | What the build printed, kept up to 256 KiB from the end. Set when it failed. | [optional] 
**State** | **string** | pending while building, then ready or failed. | 

## Methods

### NewScratchpadBuild

`func NewScratchpadBuild(id string, state string, ) *ScratchpadBuild`

NewScratchpadBuild instantiates a new ScratchpadBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScratchpadBuildWithDefaults

`func NewScratchpadBuildWithDefaults() *ScratchpadBuild`

NewScratchpadBuildWithDefaults instantiates a new ScratchpadBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScratchpadBuild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScratchpadBuild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScratchpadBuild) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *ScratchpadBuild) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScratchpadBuild) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScratchpadBuild) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScratchpadBuild) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *ScratchpadBuild) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ScratchpadBuild) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ScratchpadBuild) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ScratchpadBuild) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetState

`func (o *ScratchpadBuild) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScratchpadBuild) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScratchpadBuild) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


