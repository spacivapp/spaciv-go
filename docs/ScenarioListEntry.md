# ScenarioListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Scenario** | [**ScenarioChangelogBody**](ScenarioChangelogBody.md) |  | 

## Methods

### NewScenarioListEntry

`func NewScenarioListEntry(id string, scenario ScenarioChangelogBody, ) *ScenarioListEntry`

NewScenarioListEntry instantiates a new ScenarioListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioListEntryWithDefaults

`func NewScenarioListEntryWithDefaults() *ScenarioListEntry`

NewScenarioListEntryWithDefaults instantiates a new ScenarioListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScenarioListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScenarioListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScenarioListEntry) SetId(v string)`

SetId sets Id field to given value.


### GetScenario

`func (o *ScenarioListEntry) GetScenario() ScenarioChangelogBody`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ScenarioListEntry) GetScenarioOk() (*ScenarioChangelogBody, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ScenarioListEntry) SetScenario(v ScenarioChangelogBody)`

SetScenario sets Scenario field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


