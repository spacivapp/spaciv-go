# CalcNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]CalcNode**](CalcNode.md) |  | 
**Data** | **[]float64** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewCalcNode

`func NewCalcNode(children []CalcNode, data []float64, id string, name string, ) *CalcNode`

NewCalcNode instantiates a new CalcNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcNodeWithDefaults

`func NewCalcNodeWithDefaults() *CalcNode`

NewCalcNodeWithDefaults instantiates a new CalcNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *CalcNode) GetChildren() []CalcNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CalcNode) GetChildrenOk() (*[]CalcNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CalcNode) SetChildren(v []CalcNode)`

SetChildren sets Children field to given value.


### GetData

`func (o *CalcNode) GetData() []float64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CalcNode) GetDataOk() (*[]float64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CalcNode) SetData(v []float64)`

SetData sets Data field to given value.


### GetId

`func (o *CalcNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalcNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalcNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CalcNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalcNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalcNode) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


