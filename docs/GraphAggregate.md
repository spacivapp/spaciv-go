# GraphAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | **int32** |  | 
**GroupBy** | **int32** |  | 
**GroupProperty** | **string** |  | 

## Methods

### NewGraphAggregate

`func NewGraphAggregate(function int32, groupBy int32, groupProperty string, ) *GraphAggregate`

NewGraphAggregate instantiates a new GraphAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphAggregateWithDefaults

`func NewGraphAggregateWithDefaults() *GraphAggregate`

NewGraphAggregateWithDefaults instantiates a new GraphAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *GraphAggregate) GetFunction() int32`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *GraphAggregate) GetFunctionOk() (*int32, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *GraphAggregate) SetFunction(v int32)`

SetFunction sets Function field to given value.


### GetGroupBy

`func (o *GraphAggregate) GetGroupBy() int32`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphAggregate) GetGroupByOk() (*int32, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphAggregate) SetGroupBy(v int32)`

SetGroupBy sets GroupBy field to given value.


### GetGroupProperty

`func (o *GraphAggregate) GetGroupProperty() string`

GetGroupProperty returns the GroupProperty field if non-nil, zero value otherwise.

### GetGroupPropertyOk

`func (o *GraphAggregate) GetGroupPropertyOk() (*string, bool)`

GetGroupPropertyOk returns a tuple with the GroupProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProperty

`func (o *GraphAggregate) SetGroupProperty(v string)`

SetGroupProperty sets GroupProperty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


