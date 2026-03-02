# GatewayBillingLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerId** | **string** | The Consumer ID making the request | 
**ApiId** | **string** | The Developer Gateway API ID being consumed | 
**Metric** | Pointer to **string** | Optional metric name (default: &#39;units&#39;) | [optional] 
**Units** | **int32** | Amount of usage to log | 

## Methods

### NewGatewayBillingLogRequest

`func NewGatewayBillingLogRequest(consumerId string, apiId string, units int32, ) *GatewayBillingLogRequest`

NewGatewayBillingLogRequest instantiates a new GatewayBillingLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayBillingLogRequestWithDefaults

`func NewGatewayBillingLogRequestWithDefaults() *GatewayBillingLogRequest`

NewGatewayBillingLogRequestWithDefaults instantiates a new GatewayBillingLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerId

`func (o *GatewayBillingLogRequest) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *GatewayBillingLogRequest) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *GatewayBillingLogRequest) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetApiId

`func (o *GatewayBillingLogRequest) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *GatewayBillingLogRequest) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *GatewayBillingLogRequest) SetApiId(v string)`

SetApiId sets ApiId field to given value.


### GetMetric

`func (o *GatewayBillingLogRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GatewayBillingLogRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GatewayBillingLogRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GatewayBillingLogRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetUnits

`func (o *GatewayBillingLogRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GatewayBillingLogRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GatewayBillingLogRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


