# PropertyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | **bool** |  | 
**DeleteAllNodes** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Node** | [**[]PropertyNodeUpdate**](PropertyNodeUpdate.md) |  | 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewPropertyBody

`func NewPropertyBody(delete bool, deleteAllNodes bool, node []PropertyNodeUpdate, ) *PropertyBody`

NewPropertyBody instantiates a new PropertyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyBodyWithDefaults

`func NewPropertyBodyWithDefaults() *PropertyBody`

NewPropertyBodyWithDefaults instantiates a new PropertyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *PropertyBody) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PropertyBody) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PropertyBody) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetDeleteAllNodes

`func (o *PropertyBody) GetDeleteAllNodes() bool`

GetDeleteAllNodes returns the DeleteAllNodes field if non-nil, zero value otherwise.

### GetDeleteAllNodesOk

`func (o *PropertyBody) GetDeleteAllNodesOk() (*bool, bool)`

GetDeleteAllNodesOk returns a tuple with the DeleteAllNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllNodes

`func (o *PropertyBody) SetDeleteAllNodes(v bool)`

SetDeleteAllNodes sets DeleteAllNodes field to given value.


### GetDescription

`func (o *PropertyBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *PropertyBody) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PropertyBody) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PropertyBody) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *PropertyBody) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *PropertyBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNode

`func (o *PropertyBody) GetNode() []PropertyNodeUpdate`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PropertyBody) GetNodeOk() (*[]PropertyNodeUpdate, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PropertyBody) SetNode(v []PropertyNodeUpdate)`

SetNode sets Node field to given value.


### GetType

`func (o *PropertyBody) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyBody) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyBody) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyBody) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


