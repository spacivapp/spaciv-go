# ContributionDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | [**ActivityData**](ActivityData.md) |  | 
**ActivityLocation** | [**ActivityLocationData**](ActivityLocationData.md) |  | 
**ActivitySpaceModule** | [**ActivitySpaceModuleData**](ActivitySpaceModuleData.md) |  | 
**AdditionalInfo** | **string** |  | 
**Confidentiality** | [**ConfidentialityData**](ConfidentialityData.md) |  | 
**DaysOff** | [**DaysOffData**](DaysOffData.md) |  | 
**FreeText** | **string** |  | 
**InOfficeDays** | [**InOfficeDaysData**](InOfficeDaysData.md) |  | 
**InOfficeHours** | [**InOfficeHoursDataRequest**](InOfficeHoursDataRequest.md) |  | 
**Location** | [**LocationData**](LocationData.md) |  | 
**MeetingDetails** | [**MeetingDetailsData**](MeetingDetailsData.md) |  | 
**MeetingSize** | [**MeetingSizeData**](MeetingSizeData.md) |  | 
**MultiSelect** | [**MultiSelectContributionData**](MultiSelectContributionData.md) |  | 
**OutOfOffice** | [**OutOfOfficeData**](OutOfOfficeData.md) |  | 
**Rating** | **float64** |  | 
**Relationship** | [**RelationshipData**](RelationshipData.md) |  | 
**RoleTags** | [**RoleTagsData**](RoleTagsData.md) |  | 
**SingleSelect** | [**SingleSelectContributionData**](SingleSelectContributionData.md) |  | 
**SingleSelectList** | [**SingleSelectListContributionData**](SingleSelectListContributionData.md) |  | 
**Workingtime** | **float64** |  | 

## Methods

### NewContributionDataRequest

`func NewContributionDataRequest(activity ActivityData, activityLocation ActivityLocationData, activitySpaceModule ActivitySpaceModuleData, additionalInfo string, confidentiality ConfidentialityData, daysOff DaysOffData, freeText string, inOfficeDays InOfficeDaysData, inOfficeHours InOfficeHoursDataRequest, location LocationData, meetingDetails MeetingDetailsData, meetingSize MeetingSizeData, multiSelect MultiSelectContributionData, outOfOffice OutOfOfficeData, rating float64, relationship RelationshipData, roleTags RoleTagsData, singleSelect SingleSelectContributionData, singleSelectList SingleSelectListContributionData, workingtime float64, ) *ContributionDataRequest`

NewContributionDataRequest instantiates a new ContributionDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributionDataRequestWithDefaults

`func NewContributionDataRequestWithDefaults() *ContributionDataRequest`

NewContributionDataRequestWithDefaults instantiates a new ContributionDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ContributionDataRequest) GetActivity() ActivityData`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ContributionDataRequest) GetActivityOk() (*ActivityData, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ContributionDataRequest) SetActivity(v ActivityData)`

SetActivity sets Activity field to given value.


### GetActivityLocation

`func (o *ContributionDataRequest) GetActivityLocation() ActivityLocationData`

GetActivityLocation returns the ActivityLocation field if non-nil, zero value otherwise.

### GetActivityLocationOk

`func (o *ContributionDataRequest) GetActivityLocationOk() (*ActivityLocationData, bool)`

GetActivityLocationOk returns a tuple with the ActivityLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLocation

`func (o *ContributionDataRequest) SetActivityLocation(v ActivityLocationData)`

SetActivityLocation sets ActivityLocation field to given value.


### GetActivitySpaceModule

`func (o *ContributionDataRequest) GetActivitySpaceModule() ActivitySpaceModuleData`

GetActivitySpaceModule returns the ActivitySpaceModule field if non-nil, zero value otherwise.

### GetActivitySpaceModuleOk

`func (o *ContributionDataRequest) GetActivitySpaceModuleOk() (*ActivitySpaceModuleData, bool)`

GetActivitySpaceModuleOk returns a tuple with the ActivitySpaceModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySpaceModule

`func (o *ContributionDataRequest) SetActivitySpaceModule(v ActivitySpaceModuleData)`

SetActivitySpaceModule sets ActivitySpaceModule field to given value.


### GetAdditionalInfo

`func (o *ContributionDataRequest) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ContributionDataRequest) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ContributionDataRequest) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.


### GetConfidentiality

`func (o *ContributionDataRequest) GetConfidentiality() ConfidentialityData`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *ContributionDataRequest) GetConfidentialityOk() (*ConfidentialityData, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *ContributionDataRequest) SetConfidentiality(v ConfidentialityData)`

SetConfidentiality sets Confidentiality field to given value.


### GetDaysOff

`func (o *ContributionDataRequest) GetDaysOff() DaysOffData`

GetDaysOff returns the DaysOff field if non-nil, zero value otherwise.

### GetDaysOffOk

`func (o *ContributionDataRequest) GetDaysOffOk() (*DaysOffData, bool)`

GetDaysOffOk returns a tuple with the DaysOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOff

`func (o *ContributionDataRequest) SetDaysOff(v DaysOffData)`

SetDaysOff sets DaysOff field to given value.


### GetFreeText

`func (o *ContributionDataRequest) GetFreeText() string`

GetFreeText returns the FreeText field if non-nil, zero value otherwise.

### GetFreeTextOk

`func (o *ContributionDataRequest) GetFreeTextOk() (*string, bool)`

GetFreeTextOk returns a tuple with the FreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeText

`func (o *ContributionDataRequest) SetFreeText(v string)`

SetFreeText sets FreeText field to given value.


### GetInOfficeDays

`func (o *ContributionDataRequest) GetInOfficeDays() InOfficeDaysData`

GetInOfficeDays returns the InOfficeDays field if non-nil, zero value otherwise.

### GetInOfficeDaysOk

`func (o *ContributionDataRequest) GetInOfficeDaysOk() (*InOfficeDaysData, bool)`

GetInOfficeDaysOk returns a tuple with the InOfficeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOfficeDays

`func (o *ContributionDataRequest) SetInOfficeDays(v InOfficeDaysData)`

SetInOfficeDays sets InOfficeDays field to given value.


### GetInOfficeHours

`func (o *ContributionDataRequest) GetInOfficeHours() InOfficeHoursDataRequest`

GetInOfficeHours returns the InOfficeHours field if non-nil, zero value otherwise.

### GetInOfficeHoursOk

`func (o *ContributionDataRequest) GetInOfficeHoursOk() (*InOfficeHoursDataRequest, bool)`

GetInOfficeHoursOk returns a tuple with the InOfficeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOfficeHours

`func (o *ContributionDataRequest) SetInOfficeHours(v InOfficeHoursDataRequest)`

SetInOfficeHours sets InOfficeHours field to given value.


### GetLocation

`func (o *ContributionDataRequest) GetLocation() LocationData`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContributionDataRequest) GetLocationOk() (*LocationData, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContributionDataRequest) SetLocation(v LocationData)`

SetLocation sets Location field to given value.


### GetMeetingDetails

`func (o *ContributionDataRequest) GetMeetingDetails() MeetingDetailsData`

GetMeetingDetails returns the MeetingDetails field if non-nil, zero value otherwise.

### GetMeetingDetailsOk

`func (o *ContributionDataRequest) GetMeetingDetailsOk() (*MeetingDetailsData, bool)`

GetMeetingDetailsOk returns a tuple with the MeetingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingDetails

`func (o *ContributionDataRequest) SetMeetingDetails(v MeetingDetailsData)`

SetMeetingDetails sets MeetingDetails field to given value.


### GetMeetingSize

`func (o *ContributionDataRequest) GetMeetingSize() MeetingSizeData`

GetMeetingSize returns the MeetingSize field if non-nil, zero value otherwise.

### GetMeetingSizeOk

`func (o *ContributionDataRequest) GetMeetingSizeOk() (*MeetingSizeData, bool)`

GetMeetingSizeOk returns a tuple with the MeetingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingSize

`func (o *ContributionDataRequest) SetMeetingSize(v MeetingSizeData)`

SetMeetingSize sets MeetingSize field to given value.


### GetMultiSelect

`func (o *ContributionDataRequest) GetMultiSelect() MultiSelectContributionData`

GetMultiSelect returns the MultiSelect field if non-nil, zero value otherwise.

### GetMultiSelectOk

`func (o *ContributionDataRequest) GetMultiSelectOk() (*MultiSelectContributionData, bool)`

GetMultiSelectOk returns a tuple with the MultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSelect

`func (o *ContributionDataRequest) SetMultiSelect(v MultiSelectContributionData)`

SetMultiSelect sets MultiSelect field to given value.


### GetOutOfOffice

`func (o *ContributionDataRequest) GetOutOfOffice() OutOfOfficeData`

GetOutOfOffice returns the OutOfOffice field if non-nil, zero value otherwise.

### GetOutOfOfficeOk

`func (o *ContributionDataRequest) GetOutOfOfficeOk() (*OutOfOfficeData, bool)`

GetOutOfOfficeOk returns a tuple with the OutOfOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfOffice

`func (o *ContributionDataRequest) SetOutOfOffice(v OutOfOfficeData)`

SetOutOfOffice sets OutOfOffice field to given value.


### GetRating

`func (o *ContributionDataRequest) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ContributionDataRequest) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ContributionDataRequest) SetRating(v float64)`

SetRating sets Rating field to given value.


### GetRelationship

`func (o *ContributionDataRequest) GetRelationship() RelationshipData`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *ContributionDataRequest) GetRelationshipOk() (*RelationshipData, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *ContributionDataRequest) SetRelationship(v RelationshipData)`

SetRelationship sets Relationship field to given value.


### GetRoleTags

`func (o *ContributionDataRequest) GetRoleTags() RoleTagsData`

GetRoleTags returns the RoleTags field if non-nil, zero value otherwise.

### GetRoleTagsOk

`func (o *ContributionDataRequest) GetRoleTagsOk() (*RoleTagsData, bool)`

GetRoleTagsOk returns a tuple with the RoleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTags

`func (o *ContributionDataRequest) SetRoleTags(v RoleTagsData)`

SetRoleTags sets RoleTags field to given value.


### GetSingleSelect

`func (o *ContributionDataRequest) GetSingleSelect() SingleSelectContributionData`

GetSingleSelect returns the SingleSelect field if non-nil, zero value otherwise.

### GetSingleSelectOk

`func (o *ContributionDataRequest) GetSingleSelectOk() (*SingleSelectContributionData, bool)`

GetSingleSelectOk returns a tuple with the SingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelect

`func (o *ContributionDataRequest) SetSingleSelect(v SingleSelectContributionData)`

SetSingleSelect sets SingleSelect field to given value.


### GetSingleSelectList

`func (o *ContributionDataRequest) GetSingleSelectList() SingleSelectListContributionData`

GetSingleSelectList returns the SingleSelectList field if non-nil, zero value otherwise.

### GetSingleSelectListOk

`func (o *ContributionDataRequest) GetSingleSelectListOk() (*SingleSelectListContributionData, bool)`

GetSingleSelectListOk returns a tuple with the SingleSelectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelectList

`func (o *ContributionDataRequest) SetSingleSelectList(v SingleSelectListContributionData)`

SetSingleSelectList sets SingleSelectList field to given value.


### GetWorkingtime

`func (o *ContributionDataRequest) GetWorkingtime() float64`

GetWorkingtime returns the Workingtime field if non-nil, zero value otherwise.

### GetWorkingtimeOk

`func (o *ContributionDataRequest) GetWorkingtimeOk() (*float64, bool)`

GetWorkingtimeOk returns a tuple with the Workingtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingtime

`func (o *ContributionDataRequest) SetWorkingtime(v float64)`

SetWorkingtime sets Workingtime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


