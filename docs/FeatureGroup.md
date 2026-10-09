# FeatureGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryLu** | **string** |  | 
**CategoryName** | **string** |  | 
**RoleCategoryOptions** | [**[]Options**](Options.md) |  | 

## Methods

### NewFeatureGroup

`func NewFeatureGroup(categoryLu string, categoryName string, roleCategoryOptions []Options, ) *FeatureGroup`

NewFeatureGroup instantiates a new FeatureGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureGroupWithDefaults

`func NewFeatureGroupWithDefaults() *FeatureGroup`

NewFeatureGroupWithDefaults instantiates a new FeatureGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryLu

`func (o *FeatureGroup) GetCategoryLu() string`

GetCategoryLu returns the CategoryLu field if non-nil, zero value otherwise.

### GetCategoryLuOk

`func (o *FeatureGroup) GetCategoryLuOk() (*string, bool)`

GetCategoryLuOk returns a tuple with the CategoryLu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryLu

`func (o *FeatureGroup) SetCategoryLu(v string)`

SetCategoryLu sets CategoryLu field to given value.


### GetCategoryName

`func (o *FeatureGroup) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *FeatureGroup) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *FeatureGroup) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetRoleCategoryOptions

`func (o *FeatureGroup) GetRoleCategoryOptions() []Options`

GetRoleCategoryOptions returns the RoleCategoryOptions field if non-nil, zero value otherwise.

### GetRoleCategoryOptionsOk

`func (o *FeatureGroup) GetRoleCategoryOptionsOk() (*[]Options, bool)`

GetRoleCategoryOptionsOk returns a tuple with the RoleCategoryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCategoryOptions

`func (o *FeatureGroup) SetRoleCategoryOptions(v []Options)`

SetRoleCategoryOptions sets RoleCategoryOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


