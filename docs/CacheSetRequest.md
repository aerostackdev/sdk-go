# CacheSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | **interface{}** |  | 
**Ttl** | Pointer to **int32** | Time to live in seconds | [optional] 

## Methods

### NewCacheSetRequest

`func NewCacheSetRequest(key string, value interface{}, ) *CacheSetRequest`

NewCacheSetRequest instantiates a new CacheSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSetRequestWithDefaults

`func NewCacheSetRequestWithDefaults() *CacheSetRequest`

NewCacheSetRequestWithDefaults instantiates a new CacheSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CacheSetRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CacheSetRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CacheSetRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CacheSetRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CacheSetRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CacheSetRequest) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *CacheSetRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CacheSetRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTtl

`func (o *CacheSetRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CacheSetRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CacheSetRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CacheSetRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


