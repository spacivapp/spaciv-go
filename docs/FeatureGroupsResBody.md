# FeatureGroupsResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureToDetails** | **map[string][]string** |  | 
**Groups** | [**[]FeatureGroup**](FeatureGroup.md) |  | 

## Methods

### NewFeatureGroupsResBody

`func NewFeatureGroupsResBody(featureToDetails map[string][]string, groups []FeatureGroup, ) *FeatureGroupsResBody`

NewFeatureGroupsResBody instantiates a new FeatureGroupsResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureGroupsResBodyWithDefaults

`func NewFeatureGroupsResBodyWithDefaults() *FeatureGroupsResBody`

NewFeatureGroupsResBodyWithDefaults instantiates a new FeatureGroupsResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureToDetails

`func (o *FeatureGroupsResBody) GetFeatureToDetails() map[string][]string`

GetFeatureToDetails returns the FeatureToDetails field if non-nil, zero value otherwise.

### GetFeatureToDetailsOk

`func (o *FeatureGroupsResBody) GetFeatureToDetailsOk() (*map[string][]string, bool)`

GetFeatureToDetailsOk returns a tuple with the FeatureToDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureToDetails

`func (o *FeatureGroupsResBody) SetFeatureToDetails(v map[string][]string)`

SetFeatureToDetails sets FeatureToDetails field to given value.


### GetGroups

`func (o *FeatureGroupsResBody) GetGroups() []FeatureGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FeatureGroupsResBody) GetGroupsOk() (*[]FeatureGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FeatureGroupsResBody) SetGroups(v []FeatureGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


