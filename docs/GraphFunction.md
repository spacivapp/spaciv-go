# GraphFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExactNodes** | **bool** |  | 
**Function** | [**NullableCalcRuleRef**](CalcRuleRef.md) |  | 

## Methods

### NewGraphFunction

`func NewGraphFunction(exactNodes bool, function NullableCalcRuleRef, ) *GraphFunction`

NewGraphFunction instantiates a new GraphFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphFunctionWithDefaults

`func NewGraphFunctionWithDefaults() *GraphFunction`

NewGraphFunctionWithDefaults instantiates a new GraphFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExactNodes

`func (o *GraphFunction) GetExactNodes() bool`

GetExactNodes returns the ExactNodes field if non-nil, zero value otherwise.

### GetExactNodesOk

`func (o *GraphFunction) GetExactNodesOk() (*bool, bool)`

GetExactNodesOk returns a tuple with the ExactNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNodes

`func (o *GraphFunction) SetExactNodes(v bool)`

SetExactNodes sets ExactNodes field to given value.


### GetFunction

`func (o *GraphFunction) GetFunction() CalcRuleRef`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *GraphFunction) GetFunctionOk() (*CalcRuleRef, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *GraphFunction) SetFunction(v CalcRuleRef)`

SetFunction sets Function field to given value.


### SetFunctionNil

`func (o *GraphFunction) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *GraphFunction) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


