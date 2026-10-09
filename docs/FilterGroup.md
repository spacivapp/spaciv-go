# FilterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ands** | [**[]Filter**](Filter.md) |  | 

## Methods

### NewFilterGroup

`func NewFilterGroup(ands []Filter, ) *FilterGroup`

NewFilterGroup instantiates a new FilterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterGroupWithDefaults

`func NewFilterGroupWithDefaults() *FilterGroup`

NewFilterGroupWithDefaults instantiates a new FilterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnds

`func (o *FilterGroup) GetAnds() []Filter`

GetAnds returns the Ands field if non-nil, zero value otherwise.

### GetAndsOk

`func (o *FilterGroup) GetAndsOk() (*[]Filter, bool)`

GetAndsOk returns a tuple with the Ands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnds

`func (o *FilterGroup) SetAnds(v []Filter)`

SetAnds sets Ands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


