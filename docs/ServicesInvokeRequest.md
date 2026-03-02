# ServicesInvokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**Data** |  |  | 

## Methods

### NewServicesInvokeRequest

`func NewServicesInvokeRequest(serviceName string, data map[string]interface{}, ) *ServicesInvokeRequest`

NewServicesInvokeRequest instantiates a new ServicesInvokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesInvokeRequestWithDefaults

`func NewServicesInvokeRequestWithDefaults() *ServicesInvokeRequest`

NewServicesInvokeRequestWithDefaults instantiates a new ServicesInvokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ServicesInvokeRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServicesInvokeRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServicesInvokeRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetData

`func (o *ServicesInvokeRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesInvokeRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesInvokeRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *ServicesInvokeRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ServicesInvokeRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


