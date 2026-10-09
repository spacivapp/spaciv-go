# Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureToggles** | **map[string]bool** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewOptions

`func NewOptions(featureToggles map[string]bool, id string, name string, ) *Options`

NewOptions instantiates a new Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsWithDefaults

`func NewOptionsWithDefaults() *Options`

NewOptionsWithDefaults instantiates a new Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureToggles

`func (o *Options) GetFeatureToggles() map[string]bool`

GetFeatureToggles returns the FeatureToggles field if non-nil, zero value otherwise.

### GetFeatureTogglesOk

`func (o *Options) GetFeatureTogglesOk() (*map[string]bool, bool)`

GetFeatureTogglesOk returns a tuple with the FeatureToggles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureToggles

`func (o *Options) SetFeatureToggles(v map[string]bool)`

SetFeatureToggles sets FeatureToggles field to given value.


### GetId

`func (o *Options) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Options) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Options) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Options) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Options) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Options) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


