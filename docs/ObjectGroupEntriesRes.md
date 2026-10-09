# ObjectGroupEntriesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** |  | 
**List** | [**[]ObjectGroupEntry**](ObjectGroupEntry.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewObjectGroupEntriesRes

`func NewObjectGroupEntriesRes(cursor string, list []ObjectGroupEntry, total int32, ) *ObjectGroupEntriesRes`

NewObjectGroupEntriesRes instantiates a new ObjectGroupEntriesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupEntriesResWithDefaults

`func NewObjectGroupEntriesResWithDefaults() *ObjectGroupEntriesRes`

NewObjectGroupEntriesResWithDefaults instantiates a new ObjectGroupEntriesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ObjectGroupEntriesRes) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ObjectGroupEntriesRes) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ObjectGroupEntriesRes) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetList

`func (o *ObjectGroupEntriesRes) GetList() []ObjectGroupEntry`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ObjectGroupEntriesRes) GetListOk() (*[]ObjectGroupEntry, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ObjectGroupEntriesRes) SetList(v []ObjectGroupEntry)`

SetList sets List field to given value.


### GetTotal

`func (o *ObjectGroupEntriesRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ObjectGroupEntriesRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ObjectGroupEntriesRes) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


