# ReportParamsWrapperBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboard** | **string** | ID of the dashboard whose cards are included in the report. | 
**Email** | **string** | Recipient address for the generated report document. | 

## Methods

### NewReportParamsWrapperBody

`func NewReportParamsWrapperBody(dashboard string, email string, ) *ReportParamsWrapperBody`

NewReportParamsWrapperBody instantiates a new ReportParamsWrapperBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportParamsWrapperBodyWithDefaults

`func NewReportParamsWrapperBodyWithDefaults() *ReportParamsWrapperBody`

NewReportParamsWrapperBodyWithDefaults instantiates a new ReportParamsWrapperBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboard

`func (o *ReportParamsWrapperBody) GetDashboard() string`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *ReportParamsWrapperBody) GetDashboardOk() (*string, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *ReportParamsWrapperBody) SetDashboard(v string)`

SetDashboard sets Dashboard field to given value.


### GetEmail

`func (o *ReportParamsWrapperBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReportParamsWrapperBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReportParamsWrapperBody) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


