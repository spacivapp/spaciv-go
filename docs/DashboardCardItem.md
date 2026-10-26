# DashboardCardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The item&#39;s stable id, unique within its card. Assigned by the server when omitted. | [optional] 
**Control** | [**DashboardControl**](DashboardControl.md) |  | 
**Visualisation** | [**DashboardVisualisationCalc**](DashboardVisualisationCalc.md) |  | 

## Methods

### NewDashboardCardItem

`func NewDashboardCardItem(control DashboardControl, visualisation DashboardVisualisationCalc, ) *DashboardCardItem`

NewDashboardCardItem instantiates a new DashboardCardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCardItemWithDefaults

`func NewDashboardCardItemWithDefaults() *DashboardCardItem`

NewDashboardCardItemWithDefaults instantiates a new DashboardCardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardCardItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardCardItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardCardItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardCardItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetControl

`func (o *DashboardCardItem) GetControl() DashboardControl`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *DashboardCardItem) GetControlOk() (*DashboardControl, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *DashboardCardItem) SetControl(v DashboardControl)`

SetControl sets Control field to given value.


### GetVisualisation

`func (o *DashboardCardItem) GetVisualisation() DashboardVisualisationCalc`

GetVisualisation returns the Visualisation field if non-nil, zero value otherwise.

### GetVisualisationOk

`func (o *DashboardCardItem) GetVisualisationOk() (*DashboardVisualisationCalc, bool)`

GetVisualisationOk returns a tuple with the Visualisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualisation

`func (o *DashboardCardItem) SetVisualisation(v DashboardVisualisationCalc)`

SetVisualisation sets Visualisation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


