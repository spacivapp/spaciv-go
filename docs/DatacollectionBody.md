# DatacollectionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymisationProperty** | Pointer to **string** |  | [optional] 
**AuthenticationMode** | Pointer to [**DatacollectionAuthentication**](DatacollectionAuthentication.md) |  | [optional] 
**ClearSteps** | **bool** |  | 
**CompleteMessage** | Pointer to **string** |  | [optional] 
**Delete** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**InviteMessage** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResendInviteMessage** | Pointer to **string** |  | [optional] 
**ResendInvites** | **bool** |  | 
**Start** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Steps** | [**[]Step**](Step.md) |  | 
**ThankYouMessage** | Pointer to **string** |  | [optional] 
**WelcomeMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewDatacollectionBody

`func NewDatacollectionBody(clearSteps bool, delete bool, resendInvites bool, steps []Step, ) *DatacollectionBody`

NewDatacollectionBody instantiates a new DatacollectionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacollectionBodyWithDefaults

`func NewDatacollectionBodyWithDefaults() *DatacollectionBody`

NewDatacollectionBodyWithDefaults instantiates a new DatacollectionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymisationProperty

`func (o *DatacollectionBody) GetAnonymisationProperty() string`

GetAnonymisationProperty returns the AnonymisationProperty field if non-nil, zero value otherwise.

### GetAnonymisationPropertyOk

`func (o *DatacollectionBody) GetAnonymisationPropertyOk() (*string, bool)`

GetAnonymisationPropertyOk returns a tuple with the AnonymisationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymisationProperty

`func (o *DatacollectionBody) SetAnonymisationProperty(v string)`

SetAnonymisationProperty sets AnonymisationProperty field to given value.

### HasAnonymisationProperty

`func (o *DatacollectionBody) HasAnonymisationProperty() bool`

HasAnonymisationProperty returns a boolean if a field has been set.

### GetAuthenticationMode

`func (o *DatacollectionBody) GetAuthenticationMode() DatacollectionAuthentication`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *DatacollectionBody) GetAuthenticationModeOk() (*DatacollectionAuthentication, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *DatacollectionBody) SetAuthenticationMode(v DatacollectionAuthentication)`

SetAuthenticationMode sets AuthenticationMode field to given value.

### HasAuthenticationMode

`func (o *DatacollectionBody) HasAuthenticationMode() bool`

HasAuthenticationMode returns a boolean if a field has been set.

### GetClearSteps

`func (o *DatacollectionBody) GetClearSteps() bool`

GetClearSteps returns the ClearSteps field if non-nil, zero value otherwise.

### GetClearStepsOk

`func (o *DatacollectionBody) GetClearStepsOk() (*bool, bool)`

GetClearStepsOk returns a tuple with the ClearSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSteps

`func (o *DatacollectionBody) SetClearSteps(v bool)`

SetClearSteps sets ClearSteps field to given value.


### GetCompleteMessage

`func (o *DatacollectionBody) GetCompleteMessage() string`

GetCompleteMessage returns the CompleteMessage field if non-nil, zero value otherwise.

### GetCompleteMessageOk

`func (o *DatacollectionBody) GetCompleteMessageOk() (*string, bool)`

GetCompleteMessageOk returns a tuple with the CompleteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteMessage

`func (o *DatacollectionBody) SetCompleteMessage(v string)`

SetCompleteMessage sets CompleteMessage field to given value.

### HasCompleteMessage

`func (o *DatacollectionBody) HasCompleteMessage() bool`

HasCompleteMessage returns a boolean if a field has been set.

### GetDelete

`func (o *DatacollectionBody) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *DatacollectionBody) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *DatacollectionBody) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetDescription

`func (o *DatacollectionBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatacollectionBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatacollectionBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatacollectionBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *DatacollectionBody) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DatacollectionBody) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DatacollectionBody) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DatacollectionBody) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInviteMessage

`func (o *DatacollectionBody) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *DatacollectionBody) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *DatacollectionBody) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.

### HasInviteMessage

`func (o *DatacollectionBody) HasInviteMessage() bool`

HasInviteMessage returns a boolean if a field has been set.

### GetName

`func (o *DatacollectionBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacollectionBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacollectionBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatacollectionBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResendInviteMessage

`func (o *DatacollectionBody) GetResendInviteMessage() string`

GetResendInviteMessage returns the ResendInviteMessage field if non-nil, zero value otherwise.

### GetResendInviteMessageOk

`func (o *DatacollectionBody) GetResendInviteMessageOk() (*string, bool)`

GetResendInviteMessageOk returns a tuple with the ResendInviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendInviteMessage

`func (o *DatacollectionBody) SetResendInviteMessage(v string)`

SetResendInviteMessage sets ResendInviteMessage field to given value.

### HasResendInviteMessage

`func (o *DatacollectionBody) HasResendInviteMessage() bool`

HasResendInviteMessage returns a boolean if a field has been set.

### GetResendInvites

`func (o *DatacollectionBody) GetResendInvites() bool`

GetResendInvites returns the ResendInvites field if non-nil, zero value otherwise.

### GetResendInvitesOk

`func (o *DatacollectionBody) GetResendInvitesOk() (*bool, bool)`

GetResendInvitesOk returns a tuple with the ResendInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendInvites

`func (o *DatacollectionBody) SetResendInvites(v bool)`

SetResendInvites sets ResendInvites field to given value.


### GetStart

`func (o *DatacollectionBody) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DatacollectionBody) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DatacollectionBody) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *DatacollectionBody) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *DatacollectionBody) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatacollectionBody) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatacollectionBody) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatacollectionBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *DatacollectionBody) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DatacollectionBody) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DatacollectionBody) SetSteps(v []Step)`

SetSteps sets Steps field to given value.


### GetThankYouMessage

`func (o *DatacollectionBody) GetThankYouMessage() string`

GetThankYouMessage returns the ThankYouMessage field if non-nil, zero value otherwise.

### GetThankYouMessageOk

`func (o *DatacollectionBody) GetThankYouMessageOk() (*string, bool)`

GetThankYouMessageOk returns a tuple with the ThankYouMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMessage

`func (o *DatacollectionBody) SetThankYouMessage(v string)`

SetThankYouMessage sets ThankYouMessage field to given value.

### HasThankYouMessage

`func (o *DatacollectionBody) HasThankYouMessage() bool`

HasThankYouMessage returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *DatacollectionBody) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *DatacollectionBody) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *DatacollectionBody) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *DatacollectionBody) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


