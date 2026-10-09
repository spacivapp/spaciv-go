# PropertyListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | **bool** |  | 
**Id** | **string** |  | 
**Value** | [**PropertyChangelogBody**](PropertyChangelogBody.md) |  | 

## Methods

### NewPropertyListEntry

`func NewPropertyListEntry(builtIn bool, id string, value PropertyChangelogBody, ) *PropertyListEntry`

NewPropertyListEntry instantiates a new PropertyListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyListEntryWithDefaults

`func NewPropertyListEntryWithDefaults() *PropertyListEntry`

NewPropertyListEntryWithDefaults instantiates a new PropertyListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *PropertyListEntry) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *PropertyListEntry) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *PropertyListEntry) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.


### GetId

`func (o *PropertyListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropertyListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropertyListEntry) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *PropertyListEntry) GetValue() PropertyChangelogBody`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyListEntry) GetValueOk() (*PropertyChangelogBody, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyListEntry) SetValue(v PropertyChangelogBody)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


