# ScriptExportResBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportedAt** | **time.Time** |  | 
**Scripts** | [**[]ScriptExport**](ScriptExport.md) |  | 

## Methods

### NewScriptExportResBody

`func NewScriptExportResBody(exportedAt time.Time, scripts []ScriptExport, ) *ScriptExportResBody`

NewScriptExportResBody instantiates a new ScriptExportResBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExportResBodyWithDefaults

`func NewScriptExportResBodyWithDefaults() *ScriptExportResBody`

NewScriptExportResBodyWithDefaults instantiates a new ScriptExportResBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportedAt

`func (o *ScriptExportResBody) GetExportedAt() time.Time`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *ScriptExportResBody) GetExportedAtOk() (*time.Time, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *ScriptExportResBody) SetExportedAt(v time.Time)`

SetExportedAt sets ExportedAt field to given value.


### GetScripts

`func (o *ScriptExportResBody) GetScripts() []ScriptExport`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ScriptExportResBody) GetScriptsOk() (*[]ScriptExport, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ScriptExportResBody) SetScripts(v []ScriptExport)`

SetScripts sets Scripts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


