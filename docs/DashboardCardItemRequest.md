# DashboardCardItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Control** | [**DashboardControl**](DashboardControl.md) |  | 
**Visualisation** | [**DashboardVisualisationCalcRequest**](DashboardVisualisationCalcRequest.md) |  | 

## Methods

### NewDashboardCardItemRequest

`func NewDashboardCardItemRequest(control DashboardControl, visualisation DashboardVisualisationCalcRequest, ) *DashboardCardItemRequest`

NewDashboardCardItemRequest instantiates a new DashboardCardItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardItemRequestWithDefaults

`func NewDashboardCardItemRequestWithDefaults() *DashboardCardItemRequest`

NewDashboardCardItemRequestWithDefaults instantiates a new DashboardCardItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControl

`func (o *DashboardCardItemRequest) GetControl() DashboardControl`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *DashboardCardItemRequest) GetControlOk() (*DashboardControl, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *DashboardCardItemRequest) SetControl(v DashboardControl)`

SetControl sets Control field to given value.


### GetVisualisation

`func (o *DashboardCardItemRequest) GetVisualisation() DashboardVisualisationCalcRequest`

GetVisualisation returns the Visualisation field if non-nil, zero value otherwise.

### GetVisualisationOk

`func (o *DashboardCardItemRequest) GetVisualisationOk() (*DashboardVisualisationCalcRequest, bool)`

GetVisualisationOk returns a tuple with the Visualisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisation

`func (o *DashboardCardItemRequest) SetVisualisation(v DashboardVisualisationCalcRequest)`

SetVisualisation sets Visualisation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


