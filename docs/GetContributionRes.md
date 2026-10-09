# GetContributionRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContributionId** | **string** |  | 
**Data** | [**Contribution**](Contribution.md) |  | 
**IsAnonymous** | **bool** |  | 
**Node** | **string** |  | 
**State** | **int32** |  | 

## Methods

### NewGetContributionRes

`func NewGetContributionRes(contributionId string, data Contribution, isAnonymous bool, node string, state int32, ) *GetContributionRes`

NewGetContributionRes instantiates a new GetContributionRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContributionResWithDefaults

`func NewGetContributionResWithDefaults() *GetContributionRes`

NewGetContributionResWithDefaults instantiates a new GetContributionRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributionId

`func (o *GetContributionRes) GetContributionId() string`

GetContributionId returns the ContributionId field if non-nil, zero value otherwise.

### GetContributionIdOk

`func (o *GetContributionRes) GetContributionIdOk() (*string, bool)`

GetContributionIdOk returns a tuple with the ContributionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionId

`func (o *GetContributionRes) SetContributionId(v string)`

SetContributionId sets ContributionId field to given value.


### GetData

`func (o *GetContributionRes) GetData() Contribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContributionRes) GetDataOk() (*Contribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContributionRes) SetData(v Contribution)`

SetData sets Data field to given value.


### GetIsAnonymous

`func (o *GetContributionRes) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *GetContributionRes) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *GetContributionRes) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.


### GetNode

`func (o *GetContributionRes) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GetContributionRes) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GetContributionRes) SetNode(v string)`

SetNode sets Node field to given value.


### GetState

`func (o *GetContributionRes) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetContributionRes) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetContributionRes) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


