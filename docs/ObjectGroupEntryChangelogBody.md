# ObjectGroupEntryChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Properties** | [**map[string]ObjectGroupEntryPropertyNodes**](ObjectGroupEntryPropertyNodes.md) |  | 
**User** | **string** |  | 

## Methods

### NewObjectGroupEntryChangelogBody

`func NewObjectGroupEntryChangelogBody(name string, properties map[string]ObjectGroupEntryPropertyNodes, user string, ) *ObjectGroupEntryChangelogBody`

NewObjectGroupEntryChangelogBody instantiates a new ObjectGroupEntryChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupEntryChangelogBodyWithDefaults

`func NewObjectGroupEntryChangelogBodyWithDefaults() *ObjectGroupEntryChangelogBody`

NewObjectGroupEntryChangelogBodyWithDefaults instantiates a new ObjectGroupEntryChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectGroupEntryChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectGroupEntryChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectGroupEntryChangelogBody) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *ObjectGroupEntryChangelogBody) GetProperties() map[string]ObjectGroupEntryPropertyNodes`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectGroupEntryChangelogBody) GetPropertiesOk() (*map[string]ObjectGroupEntryPropertyNodes, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectGroupEntryChangelogBody) SetProperties(v map[string]ObjectGroupEntryPropertyNodes)`

SetProperties sets Properties field to given value.


### GetUser

`func (o *ObjectGroupEntryChangelogBody) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectGroupEntryChangelogBody) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectGroupEntryChangelogBody) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


