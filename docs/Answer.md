# Answer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **float64** |  | 
**Other** | **string** |  | 
**Ranges** | [**map[string]TimeRange**](TimeRange.md) |  | 
**Selected** | **[]string** |  | 
**SelectedByRow** | **map[string]string** |  | 
**Text** | **string** |  | 
**Values** | **map[string]float64** |  | 

## Methods

### NewAnswer

`func NewAnswer(number float64, other string, ranges map[string]TimeRange, selected []string, selectedByRow map[string]string, text string, values map[string]float64, ) *Answer`

NewAnswer instantiates a new Answer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWithDefaults

`func NewAnswerWithDefaults() *Answer`

NewAnswerWithDefaults instantiates a new Answer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Answer) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Answer) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Answer) SetNumber(v float64)`

SetNumber sets Number field to given value.


### GetOther

`func (o *Answer) GetOther() string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Answer) GetOtherOk() (*string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Answer) SetOther(v string)`

SetOther sets Other field to given value.


### GetRanges

`func (o *Answer) GetRanges() map[string]TimeRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *Answer) GetRangesOk() (*map[string]TimeRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *Answer) SetRanges(v map[string]TimeRange)`

SetRanges sets Ranges field to given value.


### GetSelected

`func (o *Answer) GetSelected() []string`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *Answer) GetSelectedOk() (*[]string, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *Answer) SetSelected(v []string)`

SetSelected sets Selected field to given value.


### GetSelectedByRow

`func (o *Answer) GetSelectedByRow() map[string]string`

GetSelectedByRow returns the SelectedByRow field if non-nil, zero value otherwise.

### GetSelectedByRowOk

`func (o *Answer) GetSelectedByRowOk() (*map[string]string, bool)`

GetSelectedByRowOk returns a tuple with the SelectedByRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByRow

`func (o *Answer) SetSelectedByRow(v map[string]string)`

SetSelectedByRow sets SelectedByRow field to given value.


### GetText

`func (o *Answer) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Answer) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Answer) SetText(v string)`

SetText sets Text field to given value.


### GetValues

`func (o *Answer) GetValues() map[string]float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Answer) GetValuesOk() (*map[string]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Answer) SetValues(v map[string]float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


