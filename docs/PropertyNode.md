# PropertyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]PropertyNode**](PropertyNode.md) |  | 
**HexCode** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Position** | **int32** |  | 

## Methods

### NewPropertyNode

`func NewPropertyNode(children []PropertyNode, hexCode string, id string, name string, position int32, ) *PropertyNode`

NewPropertyNode instantiates a new PropertyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyNodeWithDefaults

`func NewPropertyNodeWithDefaults() *PropertyNode`

NewPropertyNodeWithDefaults instantiates a new PropertyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *PropertyNode) GetChildren() []PropertyNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PropertyNode) GetChildrenOk() (*[]PropertyNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PropertyNode) SetChildren(v []PropertyNode)`

SetChildren sets Children field to given value.


### GetHexCode

`func (o *PropertyNode) GetHexCode() string`

GetHexCode returns the HexCode field if non-nil, zero value otherwise.

### GetHexCodeOk

`func (o *PropertyNode) GetHexCodeOk() (*string, bool)`

GetHexCodeOk returns a tuple with the HexCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexCode

`func (o *PropertyNode) SetHexCode(v string)`

SetHexCode sets HexCode field to given value.


### GetId

`func (o *PropertyNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropertyNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropertyNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PropertyNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyNode) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *PropertyNode) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PropertyNode) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PropertyNode) SetPosition(v int32)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


