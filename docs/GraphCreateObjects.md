# GraphCreateObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignments** | [**[]GraphCreatedLink**](GraphCreatedLink.md) |  | 
**Base** | **int32** |  | 
**IdSuffix** | **string** |  | 
**Overlay** | **int32** |  | 
**Values** | [**[]GraphCreatedValue**](GraphCreatedValue.md) |  | 

## Methods

### NewGraphCreateObjects

`func NewGraphCreateObjects(assignments []GraphCreatedLink, base int32, idSuffix string, overlay int32, values []GraphCreatedValue, ) *GraphCreateObjects`

NewGraphCreateObjects instantiates a new GraphCreateObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphCreateObjectsWithDefaults

`func NewGraphCreateObjectsWithDefaults() *GraphCreateObjects`

NewGraphCreateObjectsWithDefaults instantiates a new GraphCreateObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignments

`func (o *GraphCreateObjects) GetAssignments() []GraphCreatedLink`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *GraphCreateObjects) GetAssignmentsOk() (*[]GraphCreatedLink, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *GraphCreateObjects) SetAssignments(v []GraphCreatedLink)`

SetAssignments sets Assignments field to given value.


### GetBase

`func (o *GraphCreateObjects) GetBase() int32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *GraphCreateObjects) GetBaseOk() (*int32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *GraphCreateObjects) SetBase(v int32)`

SetBase sets Base field to given value.


### GetIdSuffix

`func (o *GraphCreateObjects) GetIdSuffix() string`

GetIdSuffix returns the IdSuffix field if non-nil, zero value otherwise.

### GetIdSuffixOk

`func (o *GraphCreateObjects) GetIdSuffixOk() (*string, bool)`

GetIdSuffixOk returns a tuple with the IdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSuffix

`func (o *GraphCreateObjects) SetIdSuffix(v string)`

SetIdSuffix sets IdSuffix field to given value.


### GetOverlay

`func (o *GraphCreateObjects) GetOverlay() int32`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *GraphCreateObjects) GetOverlayOk() (*int32, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *GraphCreateObjects) SetOverlay(v int32)`

SetOverlay sets Overlay field to given value.


### GetValues

`func (o *GraphCreateObjects) GetValues() []GraphCreatedValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GraphCreateObjects) GetValuesOk() (*[]GraphCreatedValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GraphCreateObjects) SetValues(v []GraphCreatedValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


