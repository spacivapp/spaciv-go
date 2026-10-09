# DatacollectionChangelogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymisationProperty** | **string** |  | 
**AuthenticationMode** | [**DatacollectionAuthentication**](DatacollectionAuthentication.md) |  | 
**CompleteMessage** | **string** |  | 
**Description** | **string** |  | 
**End** | **string** |  | 
**InviteMessage** | **string** |  | 
**Name** | **string** |  | 
**ResendInviteMessage** | **string** |  | 
**Start** | **string** |  | 
**Status** | **int32** |  | 
**Steps** | [**[]Step**](Step.md) |  | 
**ThankYouMessage** | **string** |  | 
**WelcomeMessage** | **string** |  | 

## Methods

### NewDatacollectionChangelogBody

`func NewDatacollectionChangelogBody(anonymisationProperty string, authenticationMode DatacollectionAuthentication, completeMessage string, description string, end string, inviteMessage string, name string, resendInviteMessage string, start string, status int32, steps []Step, thankYouMessage string, welcomeMessage string, ) *DatacollectionChangelogBody`

NewDatacollectionChangelogBody instantiates a new DatacollectionChangelogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacollectionChangelogBodyWithDefaults

`func NewDatacollectionChangelogBodyWithDefaults() *DatacollectionChangelogBody`

NewDatacollectionChangelogBodyWithDefaults instantiates a new DatacollectionChangelogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymisationProperty

`func (o *DatacollectionChangelogBody) GetAnonymisationProperty() string`

GetAnonymisationProperty returns the AnonymisationProperty field if non-nil, zero value otherwise.

### GetAnonymisationPropertyOk

`func (o *DatacollectionChangelogBody) GetAnonymisationPropertyOk() (*string, bool)`

GetAnonymisationPropertyOk returns a tuple with the AnonymisationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymisationProperty

`func (o *DatacollectionChangelogBody) SetAnonymisationProperty(v string)`

SetAnonymisationProperty sets AnonymisationProperty field to given value.


### GetAuthenticationMode

`func (o *DatacollectionChangelogBody) GetAuthenticationMode() DatacollectionAuthentication`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *DatacollectionChangelogBody) GetAuthenticationModeOk() (*DatacollectionAuthentication, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *DatacollectionChangelogBody) SetAuthenticationMode(v DatacollectionAuthentication)`

SetAuthenticationMode sets AuthenticationMode field to given value.


### GetCompleteMessage

`func (o *DatacollectionChangelogBody) GetCompleteMessage() string`

GetCompleteMessage returns the CompleteMessage field if non-nil, zero value otherwise.

### GetCompleteMessageOk

`func (o *DatacollectionChangelogBody) GetCompleteMessageOk() (*string, bool)`

GetCompleteMessageOk returns a tuple with the CompleteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteMessage

`func (o *DatacollectionChangelogBody) SetCompleteMessage(v string)`

SetCompleteMessage sets CompleteMessage field to given value.


### GetDescription

`func (o *DatacollectionChangelogBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatacollectionChangelogBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatacollectionChangelogBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnd

`func (o *DatacollectionChangelogBody) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DatacollectionChangelogBody) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DatacollectionChangelogBody) SetEnd(v string)`

SetEnd sets End field to given value.


### GetInviteMessage

`func (o *DatacollectionChangelogBody) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *DatacollectionChangelogBody) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *DatacollectionChangelogBody) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.


### GetName

`func (o *DatacollectionChangelogBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacollectionChangelogBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacollectionChangelogBody) SetName(v string)`

SetName sets Name field to given value.


### GetResendInviteMessage

`func (o *DatacollectionChangelogBody) GetResendInviteMessage() string`

GetResendInviteMessage returns the ResendInviteMessage field if non-nil, zero value otherwise.

### GetResendInviteMessageOk

`func (o *DatacollectionChangelogBody) GetResendInviteMessageOk() (*string, bool)`

GetResendInviteMessageOk returns a tuple with the ResendInviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendInviteMessage

`func (o *DatacollectionChangelogBody) SetResendInviteMessage(v string)`

SetResendInviteMessage sets ResendInviteMessage field to given value.


### GetStart

`func (o *DatacollectionChangelogBody) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DatacollectionChangelogBody) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DatacollectionChangelogBody) SetStart(v string)`

SetStart sets Start field to given value.


### GetStatus

`func (o *DatacollectionChangelogBody) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatacollectionChangelogBody) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatacollectionChangelogBody) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetSteps

`func (o *DatacollectionChangelogBody) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DatacollectionChangelogBody) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DatacollectionChangelogBody) SetSteps(v []Step)`

SetSteps sets Steps field to given value.


### GetThankYouMessage

`func (o *DatacollectionChangelogBody) GetThankYouMessage() string`

GetThankYouMessage returns the ThankYouMessage field if non-nil, zero value otherwise.

### GetThankYouMessageOk

`func (o *DatacollectionChangelogBody) GetThankYouMessageOk() (*string, bool)`

GetThankYouMessageOk returns a tuple with the ThankYouMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMessage

`func (o *DatacollectionChangelogBody) SetThankYouMessage(v string)`

SetThankYouMessage sets ThankYouMessage field to given value.


### GetWelcomeMessage

`func (o *DatacollectionChangelogBody) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *DatacollectionChangelogBody) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *DatacollectionChangelogBody) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


