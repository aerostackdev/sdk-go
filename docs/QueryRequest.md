# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**TopK** | Pointer to **int32** |  | [optional] [default to 5]
**Types** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to  |  | [optional] 

## Methods

### NewQueryRequest

`func NewQueryRequest(text string, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *QueryRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *QueryRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *QueryRequest) SetText(v string)`

SetText sets Text field to given value.


### GetTopK

`func (o *QueryRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *QueryRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *QueryRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *QueryRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetTypes

`func (o *QueryRequest) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryRequest) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryRequest) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueryRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetFilter

`func (o *QueryRequest) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryRequest) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryRequest) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueryRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *QueryRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *QueryRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


