# AssistantMessageSearchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | [**[]AssistantMessageSearchHit**](AssistantMessageSearchHit.md) |  | 
**NextCursor** | **string** |  | 

## Methods

### NewAssistantMessageSearchRes

`func NewAssistantMessageSearchRes(hits []AssistantMessageSearchHit, nextCursor string, ) *AssistantMessageSearchRes`

NewAssistantMessageSearchRes instantiates a new AssistantMessageSearchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantMessageSearchResWithDefaults

`func NewAssistantMessageSearchResWithDefaults() *AssistantMessageSearchRes`

NewAssistantMessageSearchResWithDefaults instantiates a new AssistantMessageSearchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *AssistantMessageSearchRes) GetHits() []AssistantMessageSearchHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *AssistantMessageSearchRes) GetHitsOk() (*[]AssistantMessageSearchHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *AssistantMessageSearchRes) SetHits(v []AssistantMessageSearchHit)`

SetHits sets Hits field to given value.


### GetNextCursor

`func (o *AssistantMessageSearchRes) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AssistantMessageSearchRes) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AssistantMessageSearchRes) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


