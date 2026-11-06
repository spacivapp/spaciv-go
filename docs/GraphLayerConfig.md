# GraphLayerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]GraphConnection**](GraphConnection.md) |  | 
**Nodes** | [**[]GraphNode**](GraphNode.md) |  | 
**Outputs** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewGraphLayerConfig

`func NewGraphLayerConfig(connections []GraphConnection, nodes []GraphNode, outputs []string, type_ string, ) *GraphLayerConfig`

NewGraphLayerConfig instantiates a new GraphLayerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphLayerConfigWithDefaults

`func NewGraphLayerConfigWithDefaults() *GraphLayerConfig`

NewGraphLayerConfigWithDefaults instantiates a new GraphLayerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *GraphLayerConfig) GetConnections() []GraphConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *GraphLayerConfig) GetConnectionsOk() (*[]GraphConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *GraphLayerConfig) SetConnections(v []GraphConnection)`

SetConnections sets Connections field to given value.


### GetNodes

`func (o *GraphLayerConfig) GetNodes() []GraphNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GraphLayerConfig) GetNodesOk() (*[]GraphNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GraphLayerConfig) SetNodes(v []GraphNode)`

SetNodes sets Nodes field to given value.


### GetOutputs

`func (o *GraphLayerConfig) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *GraphLayerConfig) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *GraphLayerConfig) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.


### GetType

`func (o *GraphLayerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphLayerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphLayerConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


