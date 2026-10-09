# ContributionPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Contribution**](Contribution.md) |  | 
**Node** | **string** |  | 
**State** | **int32** |  | 

## Methods

### NewContributionPost

`func NewContributionPost(data Contribution, node string, state int32, ) *ContributionPost`

NewContributionPost instantiates a new ContributionPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributionPostWithDefaults

`func NewContributionPostWithDefaults() *ContributionPost`

NewContributionPostWithDefaults instantiates a new ContributionPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContributionPost) GetData() Contribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContributionPost) GetDataOk() (*Contribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContributionPost) SetData(v Contribution)`

SetData sets Data field to given value.


### GetNode

`func (o *ContributionPost) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ContributionPost) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ContributionPost) SetNode(v string)`

SetNode sets Node field to given value.


### GetState

`func (o *ContributionPost) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContributionPost) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContributionPost) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


