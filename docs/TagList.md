# TagList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reserved** | [**[]ReservedTagInfo**](ReservedTagInfo.md) | The tags that give an item a meaning Spaciv acts on. Adding or removing one, or changing anything in an item that has one, needs a particular permission. | 
**Tags** | **[]string** | Every tag on the account&#39;s files and directories, in alphabetical order. | 

## Methods

### NewTagList

`func NewTagList(reserved []ReservedTagInfo, tags []string, ) *TagList`

NewTagList instantiates a new TagList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagListWithDefaults

`func NewTagListWithDefaults() *TagList`

NewTagListWithDefaults instantiates a new TagList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReserved

`func (o *TagList) GetReserved() []ReservedTagInfo`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *TagList) GetReservedOk() (*[]ReservedTagInfo, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *TagList) SetReserved(v []ReservedTagInfo)`

SetReserved sets Reserved field to given value.


### GetTags

`func (o *TagList) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagList) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagList) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


