# VariableData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalcRule** | **string** |  | 
**Date** | **int64** |  | 
**Floorplan** | **string** |  | 
**Node** | [**VariableProperty**](VariableProperty.md) |  | 
**Number** | [**VariableNumber**](VariableNumber.md) |  | 
**Object** | [**VariableObject**](VariableObject.md) |  | 
**Property** | **string** |  | 
**Scenario** | **string** |  | 

## Methods

### NewVariableData

`func NewVariableData(calcRule string, date int64, floorplan string, node VariableProperty, number VariableNumber, object VariableObject, property string, scenario string, ) *VariableData`

NewVariableData instantiates a new VariableData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableDataWithDefaults

`func NewVariableDataWithDefaults() *VariableData`

NewVariableDataWithDefaults instantiates a new VariableData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalcRule

`func (o *VariableData) GetCalcRule() string`

GetCalcRule returns the CalcRule field if non-nil, zero value otherwise.

### GetCalcRuleOk

`func (o *VariableData) GetCalcRuleOk() (*string, bool)`

GetCalcRuleOk returns a tuple with the CalcRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalcRule

`func (o *VariableData) SetCalcRule(v string)`

SetCalcRule sets CalcRule field to given value.


### GetDate

`func (o *VariableData) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *VariableData) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *VariableData) SetDate(v int64)`

SetDate sets Date field to given value.


### GetFloorplan

`func (o *VariableData) GetFloorplan() string`

GetFloorplan returns the Floorplan field if non-nil, zero value otherwise.

### GetFloorplanOk

`func (o *VariableData) GetFloorplanOk() (*string, bool)`

GetFloorplanOk returns a tuple with the Floorplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorplan

`func (o *VariableData) SetFloorplan(v string)`

SetFloorplan sets Floorplan field to given value.


### GetNode

`func (o *VariableData) GetNode() VariableProperty`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *VariableData) GetNodeOk() (*VariableProperty, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *VariableData) SetNode(v VariableProperty)`

SetNode sets Node field to given value.


### GetNumber

`func (o *VariableData) GetNumber() VariableNumber`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VariableData) GetNumberOk() (*VariableNumber, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VariableData) SetNumber(v VariableNumber)`

SetNumber sets Number field to given value.


### GetObject

`func (o *VariableData) GetObject() VariableObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VariableData) GetObjectOk() (*VariableObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VariableData) SetObject(v VariableObject)`

SetObject sets Object field to given value.


### GetProperty

`func (o *VariableData) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *VariableData) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *VariableData) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetScenario

`func (o *VariableData) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *VariableData) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *VariableData) SetScenario(v string)`

SetScenario sets Scenario field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


