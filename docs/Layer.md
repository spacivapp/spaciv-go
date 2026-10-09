# Layer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDay** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**StartDay** | Pointer to **int32** |  | [optional] 
**Need** | [**LayerTypeNeed**](LayerTypeNeed.md) |  | 
**ObjectGroup** | **map[string]interface{}** |  | 
**PropertyMap** | **map[string]interface{}** |  | 
**Rule** | **map[string]interface{}** |  | 
**Stack** | **map[string]interface{}** |  | 

## Methods

### NewLayer

`func NewLayer(id string, need LayerTypeNeed, objectGroup map[string]interface{}, propertyMap map[string]interface{}, rule map[string]interface{}, stack map[string]interface{}, ) *Layer`

NewLayer instantiates a new Layer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerWithDefaults

`func NewLayerWithDefaults() *Layer`

NewLayerWithDefaults instantiates a new Layer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDay

`func (o *Layer) GetEndDay() int32`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *Layer) GetEndDayOk() (*int32, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *Layer) SetEndDay(v int32)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *Layer) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### GetId

`func (o *Layer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Layer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Layer) SetId(v string)`

SetId sets Id field to given value.


### GetStartDay

`func (o *Layer) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *Layer) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *Layer) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *Layer) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetNeed

`func (o *Layer) GetNeed() LayerTypeNeed`

GetNeed returns the Need field if non-nil, zero value otherwise.

### GetNeedOk

`func (o *Layer) GetNeedOk() (*LayerTypeNeed, bool)`

GetNeedOk returns a tuple with the Need field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeed

`func (o *Layer) SetNeed(v LayerTypeNeed)`

SetNeed sets Need field to given value.


### GetObjectGroup

`func (o *Layer) GetObjectGroup() map[string]interface{}`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *Layer) GetObjectGroupOk() (*map[string]interface{}, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *Layer) SetObjectGroup(v map[string]interface{})`

SetObjectGroup sets ObjectGroup field to given value.


### GetPropertyMap

`func (o *Layer) GetPropertyMap() map[string]interface{}`

GetPropertyMap returns the PropertyMap field if non-nil, zero value otherwise.

### GetPropertyMapOk

`func (o *Layer) GetPropertyMapOk() (*map[string]interface{}, bool)`

GetPropertyMapOk returns a tuple with the PropertyMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMap

`func (o *Layer) SetPropertyMap(v map[string]interface{})`

SetPropertyMap sets PropertyMap field to given value.


### GetRule

`func (o *Layer) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *Layer) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *Layer) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.


### GetStack

`func (o *Layer) GetStack() map[string]interface{}`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *Layer) GetStackOk() (*map[string]interface{}, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *Layer) SetStack(v map[string]interface{})`

SetStack sets Stack field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


