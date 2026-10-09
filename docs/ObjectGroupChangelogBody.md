# ObjectGroupChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Properties** | [**map[string]ObjectGroupEntryPropertyNodes**](ObjectGroupEntryPropertyNodes.md) |  | 

## Methods

### NewObjectGroupChangelogBody

`func NewObjectGroupChangelogBody(description string, isPublic bool, name string, properties map[string]ObjectGroupEntryPropertyNodes, ) *ObjectGroupChangelogBody`

NewObjectGroupChangelogBody instantiates a new ObjectGroupChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupChangelogBodyWithDefaults

`func NewObjectGroupChangelogBodyWithDefaults() *ObjectGroupChangelogBody`

NewObjectGroupChangelogBodyWithDefaults instantiates a new ObjectGroupChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ObjectGroupChangelogBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectGroupChangelogBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectGroupChangelogBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsPublic

`func (o *ObjectGroupChangelogBody) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ObjectGroupChangelogBody) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ObjectGroupChangelogBody) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *ObjectGroupChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectGroupChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectGroupChangelogBody) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *ObjectGroupChangelogBody) GetProperties() map[string]ObjectGroupEntryPropertyNodes`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectGroupChangelogBody) GetPropertiesOk() (*map[string]ObjectGroupEntryPropertyNodes, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectGroupChangelogBody) SetProperties(v map[string]ObjectGroupEntryPropertyNodes)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


