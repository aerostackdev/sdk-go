# QueueEnqueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** |  |  | 
**Delay** | Pointer to **int32** | Delay in seconds before processing | [optional] 

## Methods

### NewQueueEnqueueRequest

`func NewQueueEnqueueRequest(type_ string, data map[string]interface{}, ) *QueueEnqueueRequest`

NewQueueEnqueueRequest instantiates a new QueueEnqueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueEnqueueRequestWithDefaults

`func NewQueueEnqueueRequestWithDefaults() *QueueEnqueueRequest`

NewQueueEnqueueRequestWithDefaults instantiates a new QueueEnqueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueueEnqueueRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueEnqueueRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueEnqueueRequest) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *QueueEnqueueRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueueEnqueueRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueueEnqueueRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *QueueEnqueueRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *QueueEnqueueRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDelay

`func (o *QueueEnqueueRequest) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *QueueEnqueueRequest) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *QueueEnqueueRequest) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *QueueEnqueueRequest) HasDelay() bool`

HasDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


