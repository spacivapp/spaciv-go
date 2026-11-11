# GraphPreviewBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **interface{}** |  | 
**Day** | **int64** | The day to work the graph out on, as days since 1970-01-01. | 
**Scenario** | **string** | The scenario whose stack the graph&#39;s objects come from. | 

## Methods

### NewGraphPreviewBody

`func NewGraphPreviewBody(config interface{}, day int64, scenario string, ) *GraphPreviewBody`

NewGraphPreviewBody instantiates a new GraphPreviewBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewBodyWithDefaults

`func NewGraphPreviewBodyWithDefaults() *GraphPreviewBody`

NewGraphPreviewBodyWithDefaults instantiates a new GraphPreviewBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GraphPreviewBody) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GraphPreviewBody) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GraphPreviewBody) SetConfig(v interface{})`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *GraphPreviewBody) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GraphPreviewBody) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDay

`func (o *GraphPreviewBody) GetDay() int64`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *GraphPreviewBody) GetDayOk() (*int64, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *GraphPreviewBody) SetDay(v int64)`

SetDay sets Day field to given value.


### GetScenario

`func (o *GraphPreviewBody) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *GraphPreviewBody) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *GraphPreviewBody) SetScenario(v string)`

SetScenario sets Scenario field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


