# VariableRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**VariableChangelog**](VariableChangelog.md) |  | 
**Id** | **string** |  | 

## Methods

### NewVariableRes

`func NewVariableRes(data VariableChangelog, id string, ) *VariableRes`

NewVariableRes instantiates a new VariableRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableResWithDefaults

`func NewVariableResWithDefaults() *VariableRes`

NewVariableResWithDefaults instantiates a new VariableRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VariableRes) GetData() VariableChangelog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VariableRes) GetDataOk() (*VariableChangelog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VariableRes) SetData(v VariableChangelog)`

SetData sets Data field to given value.


### GetId

`func (o *VariableRes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableRes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableRes) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


