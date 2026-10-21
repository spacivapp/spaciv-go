# PresenceMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**Id** | **string** | Identifies one connection, so one user&#39;s tabs are separate members. | 
**LastName** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 
**User** | **int64** |  | 

## Methods

### NewPresenceMember

`func NewPresenceMember(id string, state string, user int64, ) *PresenceMember`

NewPresenceMember instantiates a new PresenceMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenceMemberWithDefaults

`func NewPresenceMemberWithDefaults() *PresenceMember`

NewPresenceMemberWithDefaults instantiates a new PresenceMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PresenceMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PresenceMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PresenceMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PresenceMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *PresenceMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresenceMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresenceMember) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *PresenceMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PresenceMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PresenceMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PresenceMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetState

`func (o *PresenceMember) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PresenceMember) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PresenceMember) SetState(v string)`

SetState sets State field to given value.


### GetTarget

`func (o *PresenceMember) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PresenceMember) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PresenceMember) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PresenceMember) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *PresenceMember) GetUser() int64`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PresenceMember) GetUserOk() (*int64, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PresenceMember) SetUser(v int64)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


