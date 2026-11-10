# GraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Position** | [**NullableGraphPosition**](GraphPosition.md) |  | 
**Aggregate** | [**GraphAggregate**](GraphAggregate.md) |  | 
**And** | **map[string]interface{}** |  | 
**CalculationDate** | **map[string]interface{}** |  | 
**Changes** | [**GraphChanges**](GraphChanges.md) |  | 
**CreateObjects** | [**GraphCreateObjects**](GraphCreateObjects.md) |  | 
**Filter** | [**GraphFilter**](GraphFilter.md) |  | 
**Function** | [**GraphFunction**](GraphFunction.md) |  | 
**GraphInput** | [**GraphInput**](GraphInput.md) |  | 
**GraphOutput** | [**GraphOutput**](GraphOutput.md) |  | 
**InputObjects** | **map[string]interface{}** |  | 
**Join** | [**GraphJoin**](GraphJoin.md) |  | 
**LookupTable** | [**GraphLookupTable**](GraphLookupTable.md) |  | 
**MultiplyAdd** | **map[string]interface{}** |  | 
**Need** | [**GraphNeed**](GraphNeed.md) |  | 
**Not** | **map[string]interface{}** |  | 
**ObjectGroup** | [**GraphObjectGroup**](GraphObjectGroup.md) |  | 
**Or** | **map[string]interface{}** |  | 
**OverTime** | [**GraphOverTime**](GraphOverTime.md) |  | 
**Pair** | **map[string]interface{}** |  | 
**RemoveObjects** | **map[string]interface{}** |  | 
**Round** | [**GraphRound**](GraphRound.md) |  | 
**Select** | [**GraphSelect**](GraphSelect.md) |  | 
**SetProperty** | [**GraphSetProperty**](GraphSetProperty.md) |  | 
**SetVariable** | [**GraphSetVariable**](GraphSetVariable.md) |  | 
**Solve** | [**GraphSolve**](GraphSolve.md) |  | 
**SolveStep** | [**GraphSolveStep**](GraphSolveStep.md) |  | 
**SubGraph** | [**GraphSubGraph**](GraphSubGraph.md) |  | 
**Union** | [**GraphUnion**](GraphUnion.md) |  | 

## Methods

### NewGraphNode

`func NewGraphNode(id string, position NullableGraphPosition, aggregate GraphAggregate, and map[string]interface{}, calculationDate map[string]interface{}, changes GraphChanges, createObjects GraphCreateObjects, filter GraphFilter, function GraphFunction, graphInput GraphInput, graphOutput GraphOutput, inputObjects map[string]interface{}, join GraphJoin, lookupTable GraphLookupTable, multiplyAdd map[string]interface{}, need GraphNeed, not map[string]interface{}, objectGroup GraphObjectGroup, or map[string]interface{}, overTime GraphOverTime, pair map[string]interface{}, removeObjects map[string]interface{}, round GraphRound, select_ GraphSelect, setProperty GraphSetProperty, setVariable GraphSetVariable, solve GraphSolve, solveStep GraphSolveStep, subGraph GraphSubGraph, union GraphUnion, ) *GraphNode`

NewGraphNode instantiates a new GraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeWithDefaults

`func NewGraphNodeWithDefaults() *GraphNode`

NewGraphNodeWithDefaults instantiates a new GraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNode) SetId(v string)`

SetId sets Id field to given value.


### GetPosition

`func (o *GraphNode) GetPosition() GraphPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GraphNode) GetPositionOk() (*GraphPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GraphNode) SetPosition(v GraphPosition)`

SetPosition sets Position field to given value.


### SetPositionNil

`func (o *GraphNode) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *GraphNode) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAggregate

`func (o *GraphNode) GetAggregate() GraphAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *GraphNode) GetAggregateOk() (*GraphAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *GraphNode) SetAggregate(v GraphAggregate)`

SetAggregate sets Aggregate field to given value.


### GetAnd

`func (o *GraphNode) GetAnd() map[string]interface{}`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *GraphNode) GetAndOk() (*map[string]interface{}, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *GraphNode) SetAnd(v map[string]interface{})`

SetAnd sets And field to given value.


### GetCalculationDate

`func (o *GraphNode) GetCalculationDate() map[string]interface{}`

GetCalculationDate returns the CalculationDate field if non-nil, zero value otherwise.

### GetCalculationDateOk

`func (o *GraphNode) GetCalculationDateOk() (*map[string]interface{}, bool)`

GetCalculationDateOk returns a tuple with the CalculationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationDate

`func (o *GraphNode) SetCalculationDate(v map[string]interface{})`

SetCalculationDate sets CalculationDate field to given value.


### GetChanges

`func (o *GraphNode) GetChanges() GraphChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *GraphNode) GetChangesOk() (*GraphChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *GraphNode) SetChanges(v GraphChanges)`

SetChanges sets Changes field to given value.


### GetCreateObjects

`func (o *GraphNode) GetCreateObjects() GraphCreateObjects`

GetCreateObjects returns the CreateObjects field if non-nil, zero value otherwise.

### GetCreateObjectsOk

`func (o *GraphNode) GetCreateObjectsOk() (*GraphCreateObjects, bool)`

GetCreateObjectsOk returns a tuple with the CreateObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObjects

`func (o *GraphNode) SetCreateObjects(v GraphCreateObjects)`

SetCreateObjects sets CreateObjects field to given value.


### GetFilter

`func (o *GraphNode) GetFilter() GraphFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GraphNode) GetFilterOk() (*GraphFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GraphNode) SetFilter(v GraphFilter)`

SetFilter sets Filter field to given value.


### GetFunction

`func (o *GraphNode) GetFunction() GraphFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *GraphNode) GetFunctionOk() (*GraphFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *GraphNode) SetFunction(v GraphFunction)`

SetFunction sets Function field to given value.


### GetGraphInput

`func (o *GraphNode) GetGraphInput() GraphInput`

GetGraphInput returns the GraphInput field if non-nil, zero value otherwise.

### GetGraphInputOk

`func (o *GraphNode) GetGraphInputOk() (*GraphInput, bool)`

GetGraphInputOk returns a tuple with the GraphInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphInput

`func (o *GraphNode) SetGraphInput(v GraphInput)`

SetGraphInput sets GraphInput field to given value.


### GetGraphOutput

`func (o *GraphNode) GetGraphOutput() GraphOutput`

GetGraphOutput returns the GraphOutput field if non-nil, zero value otherwise.

### GetGraphOutputOk

`func (o *GraphNode) GetGraphOutputOk() (*GraphOutput, bool)`

GetGraphOutputOk returns a tuple with the GraphOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphOutput

`func (o *GraphNode) SetGraphOutput(v GraphOutput)`

SetGraphOutput sets GraphOutput field to given value.


### GetInputObjects

`func (o *GraphNode) GetInputObjects() map[string]interface{}`

GetInputObjects returns the InputObjects field if non-nil, zero value otherwise.

### GetInputObjectsOk

`func (o *GraphNode) GetInputObjectsOk() (*map[string]interface{}, bool)`

GetInputObjectsOk returns a tuple with the InputObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputObjects

`func (o *GraphNode) SetInputObjects(v map[string]interface{})`

SetInputObjects sets InputObjects field to given value.


### GetJoin

`func (o *GraphNode) GetJoin() GraphJoin`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *GraphNode) GetJoinOk() (*GraphJoin, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *GraphNode) SetJoin(v GraphJoin)`

SetJoin sets Join field to given value.


### GetLookupTable

`func (o *GraphNode) GetLookupTable() GraphLookupTable`

GetLookupTable returns the LookupTable field if non-nil, zero value otherwise.

### GetLookupTableOk

`func (o *GraphNode) GetLookupTableOk() (*GraphLookupTable, bool)`

GetLookupTableOk returns a tuple with the LookupTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupTable

`func (o *GraphNode) SetLookupTable(v GraphLookupTable)`

SetLookupTable sets LookupTable field to given value.


### GetMultiplyAdd

`func (o *GraphNode) GetMultiplyAdd() map[string]interface{}`

GetMultiplyAdd returns the MultiplyAdd field if non-nil, zero value otherwise.

### GetMultiplyAddOk

`func (o *GraphNode) GetMultiplyAddOk() (*map[string]interface{}, bool)`

GetMultiplyAddOk returns a tuple with the MultiplyAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplyAdd

`func (o *GraphNode) SetMultiplyAdd(v map[string]interface{})`

SetMultiplyAdd sets MultiplyAdd field to given value.


### GetNeed

`func (o *GraphNode) GetNeed() GraphNeed`

GetNeed returns the Need field if non-nil, zero value otherwise.

### GetNeedOk

`func (o *GraphNode) GetNeedOk() (*GraphNeed, bool)`

GetNeedOk returns a tuple with the Need field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeed

`func (o *GraphNode) SetNeed(v GraphNeed)`

SetNeed sets Need field to given value.


### GetNot

`func (o *GraphNode) GetNot() map[string]interface{}`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *GraphNode) GetNotOk() (*map[string]interface{}, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *GraphNode) SetNot(v map[string]interface{})`

SetNot sets Not field to given value.


### GetObjectGroup

`func (o *GraphNode) GetObjectGroup() GraphObjectGroup`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *GraphNode) GetObjectGroupOk() (*GraphObjectGroup, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *GraphNode) SetObjectGroup(v GraphObjectGroup)`

SetObjectGroup sets ObjectGroup field to given value.


### GetOr

`func (o *GraphNode) GetOr() map[string]interface{}`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *GraphNode) GetOrOk() (*map[string]interface{}, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *GraphNode) SetOr(v map[string]interface{})`

SetOr sets Or field to given value.


### GetOverTime

`func (o *GraphNode) GetOverTime() GraphOverTime`

GetOverTime returns the OverTime field if non-nil, zero value otherwise.

### GetOverTimeOk

`func (o *GraphNode) GetOverTimeOk() (*GraphOverTime, bool)`

GetOverTimeOk returns a tuple with the OverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTime

`func (o *GraphNode) SetOverTime(v GraphOverTime)`

SetOverTime sets OverTime field to given value.


### GetPair

`func (o *GraphNode) GetPair() map[string]interface{}`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *GraphNode) GetPairOk() (*map[string]interface{}, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *GraphNode) SetPair(v map[string]interface{})`

SetPair sets Pair field to given value.


### GetRemoveObjects

`func (o *GraphNode) GetRemoveObjects() map[string]interface{}`

GetRemoveObjects returns the RemoveObjects field if non-nil, zero value otherwise.

### GetRemoveObjectsOk

`func (o *GraphNode) GetRemoveObjectsOk() (*map[string]interface{}, bool)`

GetRemoveObjectsOk returns a tuple with the RemoveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveObjects

`func (o *GraphNode) SetRemoveObjects(v map[string]interface{})`

SetRemoveObjects sets RemoveObjects field to given value.


### GetRound

`func (o *GraphNode) GetRound() GraphRound`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *GraphNode) GetRoundOk() (*GraphRound, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *GraphNode) SetRound(v GraphRound)`

SetRound sets Round field to given value.


### GetSelect

`func (o *GraphNode) GetSelect() GraphSelect`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *GraphNode) GetSelectOk() (*GraphSelect, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *GraphNode) SetSelect(v GraphSelect)`

SetSelect sets Select field to given value.


### GetSetProperty

`func (o *GraphNode) GetSetProperty() GraphSetProperty`

GetSetProperty returns the SetProperty field if non-nil, zero value otherwise.

### GetSetPropertyOk

`func (o *GraphNode) GetSetPropertyOk() (*GraphSetProperty, bool)`

GetSetPropertyOk returns a tuple with the SetProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetProperty

`func (o *GraphNode) SetSetProperty(v GraphSetProperty)`

SetSetProperty sets SetProperty field to given value.


### GetSetVariable

`func (o *GraphNode) GetSetVariable() GraphSetVariable`

GetSetVariable returns the SetVariable field if non-nil, zero value otherwise.

### GetSetVariableOk

`func (o *GraphNode) GetSetVariableOk() (*GraphSetVariable, bool)`

GetSetVariableOk returns a tuple with the SetVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetVariable

`func (o *GraphNode) SetSetVariable(v GraphSetVariable)`

SetSetVariable sets SetVariable field to given value.


### GetSolve

`func (o *GraphNode) GetSolve() GraphSolve`

GetSolve returns the Solve field if non-nil, zero value otherwise.

### GetSolveOk

`func (o *GraphNode) GetSolveOk() (*GraphSolve, bool)`

GetSolveOk returns a tuple with the Solve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolve

`func (o *GraphNode) SetSolve(v GraphSolve)`

SetSolve sets Solve field to given value.


### GetSolveStep

`func (o *GraphNode) GetSolveStep() GraphSolveStep`

GetSolveStep returns the SolveStep field if non-nil, zero value otherwise.

### GetSolveStepOk

`func (o *GraphNode) GetSolveStepOk() (*GraphSolveStep, bool)`

GetSolveStepOk returns a tuple with the SolveStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolveStep

`func (o *GraphNode) SetSolveStep(v GraphSolveStep)`

SetSolveStep sets SolveStep field to given value.


### GetSubGraph

`func (o *GraphNode) GetSubGraph() GraphSubGraph`

GetSubGraph returns the SubGraph field if non-nil, zero value otherwise.

### GetSubGraphOk

`func (o *GraphNode) GetSubGraphOk() (*GraphSubGraph, bool)`

GetSubGraphOk returns a tuple with the SubGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGraph

`func (o *GraphNode) SetSubGraph(v GraphSubGraph)`

SetSubGraph sets SubGraph field to given value.


### GetUnion

`func (o *GraphNode) GetUnion() GraphUnion`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *GraphNode) GetUnionOk() (*GraphUnion, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *GraphNode) SetUnion(v GraphUnion)`

SetUnion sets Union field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


