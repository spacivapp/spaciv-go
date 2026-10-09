# SysAdminUpdateFeatureGroupsParamsWrapperBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to **string** | Availability tier to toggle this feature flag for. | [optional] 
**Product** | Pointer to **string** | Product tier to toggle this feature flag for. | [optional] 
**Role** | Pointer to **string** | Role name to toggle this feature flag for. Exactly one of role, availability, or product must be supplied. | [optional] 
**SortOrder** | **int32** | Feature flag index identifying the flag to update. | 

## Methods

### NewSysAdminUpdateFeatureGroupsParamsWrapperBody

`func NewSysAdminUpdateFeatureGroupsParamsWrapperBody(sortOrder int32, ) *SysAdminUpdateFeatureGroupsParamsWrapperBody`

NewSysAdminUpdateFeatureGroupsParamsWrapperBody instantiates a new SysAdminUpdateFeatureGroupsParamsWrapperBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysAdminUpdateFeatureGroupsParamsWrapperBodyWithDefaults

`func NewSysAdminUpdateFeatureGroupsParamsWrapperBodyWithDefaults() *SysAdminUpdateFeatureGroupsParamsWrapperBody`

NewSysAdminUpdateFeatureGroupsParamsWrapperBodyWithDefaults instantiates a new SysAdminUpdateFeatureGroupsParamsWrapperBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetProduct

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRole

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSortOrder

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SysAdminUpdateFeatureGroupsParamsWrapperBody) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


