# GetMappingRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**map[string]MappingProp**](MappingProp.md) |  | 
**Scenarios** | [**map[string]MappingScenario**](MappingScenario.md) |  | 

## Methods

### NewGetMappingRes

`func NewGetMappingRes(properties map[string]MappingProp, scenarios map[string]MappingScenario, ) *GetMappingRes`

NewGetMappingRes instantiates a new GetMappingRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMappingResWithDefaults

`func NewGetMappingResWithDefaults() *GetMappingRes`

NewGetMappingResWithDefaults instantiates a new GetMappingRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *GetMappingRes) GetProperties() map[string]MappingProp`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetMappingRes) GetPropertiesOk() (*map[string]MappingProp, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetMappingRes) SetProperties(v map[string]MappingProp)`

SetProperties sets Properties field to given value.


### GetScenarios

`func (o *GetMappingRes) GetScenarios() map[string]MappingScenario`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *GetMappingRes) GetScenariosOk() (*map[string]MappingScenario, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *GetMappingRes) SetScenarios(v map[string]MappingScenario)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


