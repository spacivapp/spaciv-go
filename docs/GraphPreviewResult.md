# GraphPreviewResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Node** | **string** |  | 
**Sockets** | [**[]GraphPreviewSocket**](GraphPreviewSocket.md) |  | 
**Table** | Pointer to [**GraphPreviewTable**](GraphPreviewTable.md) |  | [optional] 

## Methods

### NewGraphPreviewResult

`func NewGraphPreviewResult(error_ string, node string, sockets []GraphPreviewSocket, ) *GraphPreviewResult`

NewGraphPreviewResult instantiates a new GraphPreviewResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewResultWithDefaults

`func NewGraphPreviewResultWithDefaults() *GraphPreviewResult`

NewGraphPreviewResultWithDefaults instantiates a new GraphPreviewResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GraphPreviewResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GraphPreviewResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GraphPreviewResult) SetError(v string)`

SetError sets Error field to given value.


### GetNode

`func (o *GraphPreviewResult) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GraphPreviewResult) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GraphPreviewResult) SetNode(v string)`

SetNode sets Node field to given value.


### GetSockets

`func (o *GraphPreviewResult) GetSockets() []GraphPreviewSocket`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *GraphPreviewResult) GetSocketsOk() (*[]GraphPreviewSocket, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *GraphPreviewResult) SetSockets(v []GraphPreviewSocket)`

SetSockets sets Sockets field to given value.


### GetTable

`func (o *GraphPreviewResult) GetTable() GraphPreviewTable`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *GraphPreviewResult) GetTableOk() (*GraphPreviewTable, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *GraphPreviewResult) SetTable(v GraphPreviewTable)`

SetTable sets Table field to given value.

### HasTable

`func (o *GraphPreviewResult) HasTable() bool`

HasTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


