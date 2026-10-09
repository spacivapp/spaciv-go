# ContributionChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Contribution**](Contribution.md) |  | 
**Membership** | [**ContributionMembership**](ContributionMembership.md) |  | 
**Node** | **string** |  | 
**PooledContributionId** | **string** |  | 
**State** | **int32** |  | 

## Methods

### NewContributionChangelogBody

`func NewContributionChangelogBody(data Contribution, membership ContributionMembership, node string, pooledContributionId string, state int32, ) *ContributionChangelogBody`

NewContributionChangelogBody instantiates a new ContributionChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributionChangelogBodyWithDefaults

`func NewContributionChangelogBodyWithDefaults() *ContributionChangelogBody`

NewContributionChangelogBodyWithDefaults instantiates a new ContributionChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContributionChangelogBody) GetData() Contribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContributionChangelogBody) GetDataOk() (*Contribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContributionChangelogBody) SetData(v Contribution)`

SetData sets Data field to given value.


### GetMembership

`func (o *ContributionChangelogBody) GetMembership() ContributionMembership`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *ContributionChangelogBody) GetMembershipOk() (*ContributionMembership, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *ContributionChangelogBody) SetMembership(v ContributionMembership)`

SetMembership sets Membership field to given value.


### GetNode

`func (o *ContributionChangelogBody) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ContributionChangelogBody) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ContributionChangelogBody) SetNode(v string)`

SetNode sets Node field to given value.


### GetPooledContributionId

`func (o *ContributionChangelogBody) GetPooledContributionId() string`

GetPooledContributionId returns the PooledContributionId field if non-nil, zero value otherwise.

### GetPooledContributionIdOk

`func (o *ContributionChangelogBody) GetPooledContributionIdOk() (*string, bool)`

GetPooledContributionIdOk returns a tuple with the PooledContributionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooledContributionId

`func (o *ContributionChangelogBody) SetPooledContributionId(v string)`

SetPooledContributionId sets PooledContributionId field to given value.


### GetState

`func (o *ContributionChangelogBody) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContributionChangelogBody) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContributionChangelogBody) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


