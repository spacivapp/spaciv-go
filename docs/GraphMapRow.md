# GraphMapRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | **[]string** |  | 
**Node** | **string** |  | 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewGraphMapRow

`func NewGraphMapRow(codes []string, node string, ) *GraphMapRow`

NewGraphMapRow instantiates a new GraphMapRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphMapRowWithDefaults

`func NewGraphMapRowWithDefaults() *GraphMapRow`

NewGraphMapRowWithDefaults instantiates a new GraphMapRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *GraphMapRow) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *GraphMapRow) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *GraphMapRow) SetCodes(v []string)`

SetCodes sets Codes field to given value.


### GetNode

`func (o *GraphMapRow) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GraphMapRow) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GraphMapRow) SetNode(v string)`

SetNode sets Node field to given value.


### GetValue

`func (o *GraphMapRow) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphMapRow) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphMapRow) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *GraphMapRow) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


