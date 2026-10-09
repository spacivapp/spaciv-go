# VariableCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**VariableData**](VariableData.md) |  | 
**Name** | **string** |  | 

## Methods

### NewVariableCreate

`func NewVariableCreate(data VariableData, name string, ) *VariableCreate`

NewVariableCreate instantiates a new VariableCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCreateWithDefaults

`func NewVariableCreateWithDefaults() *VariableCreate`

NewVariableCreateWithDefaults instantiates a new VariableCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VariableCreate) GetData() VariableData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VariableCreate) GetDataOk() (*VariableData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VariableCreate) SetData(v VariableData)`

SetData sets Data field to given value.


### GetName

`func (o *VariableCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


