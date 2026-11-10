# BuiltInSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** | Whether the account has a skill of this name, and if it does, whether its files are exactly the built-in skill&#39;s. | 

## Methods

### NewBuiltInSkill

`func NewBuiltInSkill(description string, name string, status string, ) *BuiltInSkill`

NewBuiltInSkill instantiates a new BuiltInSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltInSkillWithDefaults

`func NewBuiltInSkillWithDefaults() *BuiltInSkill`

NewBuiltInSkillWithDefaults instantiates a new BuiltInSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuiltInSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuiltInSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuiltInSkill) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *BuiltInSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuiltInSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuiltInSkill) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *BuiltInSkill) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuiltInSkill) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuiltInSkill) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


