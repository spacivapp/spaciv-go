# Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectGroup** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** |  | [optional] 

## Methods

### NewStep

`func NewStep(id string, ) *Step`

NewStep instantiates a new Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepWithDefaults

`func NewStepWithDefaults() *Step`

NewStepWithDefaults instantiates a new Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Step) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Step) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Step) SetId(v string)`

SetId sets Id field to given value.


### GetObjectGroup

`func (o *Step) GetObjectGroup() string`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *Step) GetObjectGroupOk() (*string, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *Step) SetObjectGroup(v string)`

SetObjectGroup sets ObjectGroup field to given value.

### HasObjectGroup

`func (o *Step) HasObjectGroup() bool`

HasObjectGroup returns a boolean if a field has been set.

### GetProperty

`func (o *Step) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Step) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Step) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Step) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


