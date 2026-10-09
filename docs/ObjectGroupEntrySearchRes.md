# ObjectGroupEntrySearchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** |  | 
**Indexing** | **bool** |  | 
**List** | [**[]ObjectGroupEntryRef**](ObjectGroupEntryRef.md) |  | 

## Methods

### NewObjectGroupEntrySearchRes

`func NewObjectGroupEntrySearchRes(cursor string, indexing bool, list []ObjectGroupEntryRef, ) *ObjectGroupEntrySearchRes`

NewObjectGroupEntrySearchRes instantiates a new ObjectGroupEntrySearchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupEntrySearchResWithDefaults

`func NewObjectGroupEntrySearchResWithDefaults() *ObjectGroupEntrySearchRes`

NewObjectGroupEntrySearchResWithDefaults instantiates a new ObjectGroupEntrySearchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ObjectGroupEntrySearchRes) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ObjectGroupEntrySearchRes) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ObjectGroupEntrySearchRes) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetIndexing

`func (o *ObjectGroupEntrySearchRes) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *ObjectGroupEntrySearchRes) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *ObjectGroupEntrySearchRes) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.


### GetList

`func (o *ObjectGroupEntrySearchRes) GetList() []ObjectGroupEntryRef`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ObjectGroupEntrySearchRes) GetListOk() (*[]ObjectGroupEntryRef, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ObjectGroupEntrySearchRes) SetList(v []ObjectGroupEntryRef)`

SetList sets List field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


