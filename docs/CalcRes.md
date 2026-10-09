# CalcRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Legend** | [**Legend**](Legend.md) |  | 
**Timeline** | [**[]NodeOffset**](NodeOffset.md) |  | 

## Methods

### NewCalcRes

`func NewCalcRes(legend Legend, timeline []NodeOffset, ) *CalcRes`

NewCalcRes instantiates a new CalcRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcResWithDefaults

`func NewCalcResWithDefaults() *CalcRes`

NewCalcResWithDefaults instantiates a new CalcRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegend

`func (o *CalcRes) GetLegend() Legend`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *CalcRes) GetLegendOk() (*Legend, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *CalcRes) SetLegend(v Legend)`

SetLegend sets Legend field to given value.


### GetTimeline

`func (o *CalcRes) GetTimeline() []NodeOffset`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *CalcRes) GetTimelineOk() (*[]NodeOffset, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *CalcRes) SetTimeline(v []NodeOffset)`

SetTimeline sets Timeline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


