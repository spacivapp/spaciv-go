# Selection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]Group**](Group.md) | The object groups holding a selected object, by name, each with its selected objects by name. | 
**Truncated** | **bool** | More objects matched than were returned; narrow the filters to see the rest. | 

## Methods

### NewSelection

`func NewSelection(groups []Group, truncated bool, ) *Selection`

NewSelection instantiates a new Selection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectionWithDefaults

`func NewSelectionWithDefaults() *Selection`

NewSelectionWithDefaults instantiates a new Selection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *Selection) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Selection) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Selection) SetGroups(v []Group)`

SetGroups sets Groups field to given value.


### GetTruncated

`func (o *Selection) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *Selection) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *Selection) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


