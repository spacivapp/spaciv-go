# VariableChangelogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**VariableData**](VariableData.md) |  | 
**Name** | **string** |  | 

## Methods

### NewVariableChangelogRequest

`func NewVariableChangelogRequest(data VariableData, name string, ) *VariableChangelogRequest`

NewVariableChangelogRequest instantiates a new VariableChangelogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableChangelogRequestWithDefaults

`func NewVariableChangelogRequestWithDefaults() *VariableChangelogRequest`

NewVariableChangelogRequestWithDefaults instantiates a new VariableChangelogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VariableChangelogRequest) GetData() VariableData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VariableChangelogRequest) GetDataOk() (*VariableData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VariableChangelogRequest) SetData(v VariableData)`

SetData sets Data field to given value.


### GetName

`func (o *VariableChangelogRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableChangelogRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableChangelogRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


