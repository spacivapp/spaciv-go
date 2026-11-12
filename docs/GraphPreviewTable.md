# GraphPreviewTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | [**[]GraphPreviewSample**](GraphPreviewSample.md) |  | 
**Total** | **string** |  | 

## Methods

### NewGraphPreviewTable

`func NewGraphPreviewTable(rows []GraphPreviewSample, total string, ) *GraphPreviewTable`

NewGraphPreviewTable instantiates a new GraphPreviewTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphPreviewTableWithDefaults

`func NewGraphPreviewTableWithDefaults() *GraphPreviewTable`

NewGraphPreviewTableWithDefaults instantiates a new GraphPreviewTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GraphPreviewTable) GetRows() []GraphPreviewSample`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GraphPreviewTable) GetRowsOk() (*[]GraphPreviewSample, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GraphPreviewTable) SetRows(v []GraphPreviewSample)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *GraphPreviewTable) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GraphPreviewTable) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GraphPreviewTable) SetTotal(v string)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


