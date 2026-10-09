# GetDatacollectionRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **int32** |  | 
**Datacollection** | [**DatacollectionChangelogBody**](DatacollectionChangelogBody.md) |  | 
**Id** | **string** |  | 
**Started** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewGetDatacollectionRes

`func NewGetDatacollectionRes(completed int32, datacollection DatacollectionChangelogBody, id string, started int32, total int32, ) *GetDatacollectionRes`

NewGetDatacollectionRes instantiates a new GetDatacollectionRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatacollectionResWithDefaults

`func NewGetDatacollectionResWithDefaults() *GetDatacollectionRes`

NewGetDatacollectionResWithDefaults instantiates a new GetDatacollectionRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *GetDatacollectionRes) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *GetDatacollectionRes) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *GetDatacollectionRes) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.


### GetDatacollection

`func (o *GetDatacollectionRes) GetDatacollection() DatacollectionChangelogBody`

GetDatacollection returns the Datacollection field if non-nil, zero value otherwise.

### GetDatacollectionOk

`func (o *GetDatacollectionRes) GetDatacollectionOk() (*DatacollectionChangelogBody, bool)`

GetDatacollectionOk returns a tuple with the Datacollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacollection

`func (o *GetDatacollectionRes) SetDatacollection(v DatacollectionChangelogBody)`

SetDatacollection sets Datacollection field to given value.


### GetId

`func (o *GetDatacollectionRes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatacollectionRes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatacollectionRes) SetId(v string)`

SetId sets Id field to given value.


### GetStarted

`func (o *GetDatacollectionRes) GetStarted() int32`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetDatacollectionRes) GetStartedOk() (*int32, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetDatacollectionRes) SetStarted(v int32)`

SetStarted sets Started field to given value.


### GetTotal

`func (o *GetDatacollectionRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDatacollectionRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDatacollectionRes) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


