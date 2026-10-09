# PropertySource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Floorplan** | **string** |  | 
**Property** | [**MaybeVarProp**](MaybeVarProp.md) |  | 

## Methods

### NewPropertySource

`func NewPropertySource(floorplan string, property MaybeVarProp, ) *PropertySource`

NewPropertySource instantiates a new PropertySource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertySourceWithDefaults

`func NewPropertySourceWithDefaults() *PropertySource`

NewPropertySourceWithDefaults instantiates a new PropertySource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloorplan

`func (o *PropertySource) GetFloorplan() string`

GetFloorplan returns the Floorplan field if non-nil, zero value otherwise.

### GetFloorplanOk

`func (o *PropertySource) GetFloorplanOk() (*string, bool)`

GetFloorplanOk returns a tuple with the Floorplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorplan

`func (o *PropertySource) SetFloorplan(v string)`

SetFloorplan sets Floorplan field to given value.


### GetProperty

`func (o *PropertySource) GetProperty() MaybeVarProp`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PropertySource) GetPropertyOk() (*MaybeVarProp, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PropertySource) SetProperty(v MaybeVarProp)`

SetProperty sets Property field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


