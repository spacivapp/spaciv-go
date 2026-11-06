# GraphLookupTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **float64** |  | 
**Entries** | [**[]GraphLookupEntry**](GraphLookupEntry.md) |  | 
**KeyBy** | **int32** |  | 
**Missing** | **int32** |  | 
**RequireAllKeys** | **bool** |  | 

## Methods

### NewGraphLookupTable

`func NewGraphLookupTable(default_ float64, entries []GraphLookupEntry, keyBy int32, missing int32, requireAllKeys bool, ) *GraphLookupTable`

NewGraphLookupTable instantiates a new GraphLookupTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphLookupTableWithDefaults

`func NewGraphLookupTableWithDefaults() *GraphLookupTable`

NewGraphLookupTableWithDefaults instantiates a new GraphLookupTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *GraphLookupTable) GetDefault() float64`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GraphLookupTable) GetDefaultOk() (*float64, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GraphLookupTable) SetDefault(v float64)`

SetDefault sets Default field to given value.


### GetEntries

`func (o *GraphLookupTable) GetEntries() []GraphLookupEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GraphLookupTable) GetEntriesOk() (*[]GraphLookupEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GraphLookupTable) SetEntries(v []GraphLookupEntry)`

SetEntries sets Entries field to given value.


### GetKeyBy

`func (o *GraphLookupTable) GetKeyBy() int32`

GetKeyBy returns the KeyBy field if non-nil, zero value otherwise.

### GetKeyByOk

`func (o *GraphLookupTable) GetKeyByOk() (*int32, bool)`

GetKeyByOk returns a tuple with the KeyBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBy

`func (o *GraphLookupTable) SetKeyBy(v int32)`

SetKeyBy sets KeyBy field to given value.


### GetMissing

`func (o *GraphLookupTable) GetMissing() int32`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *GraphLookupTable) GetMissingOk() (*int32, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *GraphLookupTable) SetMissing(v int32)`

SetMissing sets Missing field to given value.


### GetRequireAllKeys

`func (o *GraphLookupTable) GetRequireAllKeys() bool`

GetRequireAllKeys returns the RequireAllKeys field if non-nil, zero value otherwise.

### GetRequireAllKeysOk

`func (o *GraphLookupTable) GetRequireAllKeysOk() (*bool, bool)`

GetRequireAllKeysOk returns a tuple with the RequireAllKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAllKeys

`func (o *GraphLookupTable) SetRequireAllKeys(v bool)`

SetRequireAllKeys sets RequireAllKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


