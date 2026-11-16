# ImportedEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Why the entry was not saved or not confirmed. Only set when the outcome is failed or unconfirmed. | [optional] 
**Id** | **string** | The entry&#39;s ID, as sent. | 
**Outcome** | **string** | Whether the entry was created, updated, not saved (failed), or may or may not have been saved because the save ended in an unexpected error (unconfirmed). | 

## Methods

### NewImportedEntry

`func NewImportedEntry(id string, outcome string, ) *ImportedEntry`

NewImportedEntry instantiates a new ImportedEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedEntryWithDefaults

`func NewImportedEntryWithDefaults() *ImportedEntry`

NewImportedEntryWithDefaults instantiates a new ImportedEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ImportedEntry) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImportedEntry) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImportedEntry) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImportedEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *ImportedEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportedEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportedEntry) SetId(v string)`

SetId sets Id field to given value.


### GetOutcome

`func (o *ImportedEntry) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ImportedEntry) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ImportedEntry) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


