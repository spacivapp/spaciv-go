# PropertyNodeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | **bool** |  | 
**HexCode** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NewId** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 

## Methods

### NewPropertyNodeUpdate

`func NewPropertyNodeUpdate(delete bool, id string, ) *PropertyNodeUpdate`

NewPropertyNodeUpdate instantiates a new PropertyNodeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyNodeUpdateWithDefaults

`func NewPropertyNodeUpdateWithDefaults() *PropertyNodeUpdate`

NewPropertyNodeUpdateWithDefaults instantiates a new PropertyNodeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *PropertyNodeUpdate) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PropertyNodeUpdate) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PropertyNodeUpdate) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetHexCode

`func (o *PropertyNodeUpdate) GetHexCode() string`

GetHexCode returns the HexCode field if non-nil, zero value otherwise.

### GetHexCodeOk

`func (o *PropertyNodeUpdate) GetHexCodeOk() (*string, bool)`

GetHexCodeOk returns a tuple with the HexCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexCode

`func (o *PropertyNodeUpdate) SetHexCode(v string)`

SetHexCode sets HexCode field to given value.

### HasHexCode

`func (o *PropertyNodeUpdate) HasHexCode() bool`

HasHexCode returns a boolean if a field has been set.

### GetId

`func (o *PropertyNodeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropertyNodeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropertyNodeUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PropertyNodeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyNodeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyNodeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyNodeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewId

`func (o *PropertyNodeUpdate) GetNewId() string`

GetNewId returns the NewId field if non-nil, zero value otherwise.

### GetNewIdOk

`func (o *PropertyNodeUpdate) GetNewIdOk() (*string, bool)`

GetNewIdOk returns a tuple with the NewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewId

`func (o *PropertyNodeUpdate) SetNewId(v string)`

SetNewId sets NewId field to given value.

### HasNewId

`func (o *PropertyNodeUpdate) HasNewId() bool`

HasNewId returns a boolean if a field has been set.

### GetParent

`func (o *PropertyNodeUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PropertyNodeUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PropertyNodeUpdate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PropertyNodeUpdate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPosition

`func (o *PropertyNodeUpdate) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PropertyNodeUpdate) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PropertyNodeUpdate) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PropertyNodeUpdate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


