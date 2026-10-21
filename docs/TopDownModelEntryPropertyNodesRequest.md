# TopDownModelEntryPropertyNodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Formula** | [**CalcRuleRefRequest**](CalcRuleRefRequest.md) |  | 
**NodeId** | **string** |  | 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewTopDownModelEntryPropertyNodesRequest

`func NewTopDownModelEntryPropertyNodesRequest(formula CalcRuleRefRequest, nodeId string, ) *TopDownModelEntryPropertyNodesRequest`

NewTopDownModelEntryPropertyNodesRequest instantiates a new TopDownModelEntryPropertyNodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopDownModelEntryPropertyNodesRequestWithDefaults

`func NewTopDownModelEntryPropertyNodesRequestWithDefaults() *TopDownModelEntryPropertyNodesRequest`

NewTopDownModelEntryPropertyNodesRequestWithDefaults instantiates a new TopDownModelEntryPropertyNodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *TopDownModelEntryPropertyNodesRequest) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TopDownModelEntryPropertyNodesRequest) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TopDownModelEntryPropertyNodesRequest) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TopDownModelEntryPropertyNodesRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFormula

`func (o *TopDownModelEntryPropertyNodesRequest) GetFormula() CalcRuleRefRequest`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *TopDownModelEntryPropertyNodesRequest) GetFormulaOk() (*CalcRuleRefRequest, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *TopDownModelEntryPropertyNodesRequest) SetFormula(v CalcRuleRefRequest)`

SetFormula sets Formula field to given value.


### GetNodeId

`func (o *TopDownModelEntryPropertyNodesRequest) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TopDownModelEntryPropertyNodesRequest) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TopDownModelEntryPropertyNodesRequest) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetStart

`func (o *TopDownModelEntryPropertyNodesRequest) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TopDownModelEntryPropertyNodesRequest) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TopDownModelEntryPropertyNodesRequest) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *TopDownModelEntryPropertyNodesRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


