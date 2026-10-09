# SingleSelectListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]ChangelogOption**](ChangelogOption.md) |  | 
**Options** | [**[]ChangelogValueOption**](ChangelogValueOption.md) |  | 

## Methods

### NewSingleSelectListData

`func NewSingleSelectListData(categories []ChangelogOption, options []ChangelogValueOption, ) *SingleSelectListData`

NewSingleSelectListData instantiates a new SingleSelectListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleSelectListDataWithDefaults

`func NewSingleSelectListDataWithDefaults() *SingleSelectListData`

NewSingleSelectListDataWithDefaults instantiates a new SingleSelectListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *SingleSelectListData) GetCategories() []ChangelogOption`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SingleSelectListData) GetCategoriesOk() (*[]ChangelogOption, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SingleSelectListData) SetCategories(v []ChangelogOption)`

SetCategories sets Categories field to given value.


### GetOptions

`func (o *SingleSelectListData) GetOptions() []ChangelogValueOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SingleSelectListData) GetOptionsOk() (*[]ChangelogValueOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SingleSelectListData) SetOptions(v []ChangelogValueOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


