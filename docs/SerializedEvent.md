# SerializedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **interface{}** |  | [optional] 
**Id** | **string** |  | 
**Topic** | **string** |  | 

## Methods

### NewSerializedEvent

`func NewSerializedEvent(id string, topic string, ) *SerializedEvent`

NewSerializedEvent instantiates a new SerializedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSerializedEventWithDefaults

`func NewSerializedEventWithDefaults() *SerializedEvent`

NewSerializedEventWithDefaults instantiates a new SerializedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SerializedEvent) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SerializedEvent) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SerializedEvent) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *SerializedEvent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SerializedEvent) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SerializedEvent) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetId

`func (o *SerializedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SerializedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SerializedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetTopic

`func (o *SerializedEvent) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *SerializedEvent) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *SerializedEvent) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


