# MetaVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumFrontendVersion** | **int64** | Frontend clients whose compatibility version is below this must reload before continuing | 
**Version** | **string** | Release version of the Spaciv platform | 

## Methods

### NewMetaVersionResponse

`func NewMetaVersionResponse(minimumFrontendVersion int64, version string, ) *MetaVersionResponse`

NewMetaVersionResponse instantiates a new MetaVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaVersionResponseWithDefaults

`func NewMetaVersionResponseWithDefaults() *MetaVersionResponse`

NewMetaVersionResponseWithDefaults instantiates a new MetaVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumFrontendVersion

`func (o *MetaVersionResponse) GetMinimumFrontendVersion() int64`

GetMinimumFrontendVersion returns the MinimumFrontendVersion field if non-nil, zero value otherwise.

### GetMinimumFrontendVersionOk

`func (o *MetaVersionResponse) GetMinimumFrontendVersionOk() (*int64, bool)`

GetMinimumFrontendVersionOk returns a tuple with the MinimumFrontendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFrontendVersion

`func (o *MetaVersionResponse) SetMinimumFrontendVersion(v int64)`

SetMinimumFrontendVersion sets MinimumFrontendVersion field to given value.


### GetVersion

`func (o *MetaVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetaVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetaVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


