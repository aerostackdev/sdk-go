# DbQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewDbQueryResult

`func NewDbQueryResult() *DbQueryResult`

NewDbQueryResult instantiates a new DbQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbQueryResultWithDefaults

`func NewDbQueryResultWithDefaults() *DbQueryResult`

NewDbQueryResultWithDefaults instantiates a new DbQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *DbQueryResult) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DbQueryResult) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DbQueryResult) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *DbQueryResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *DbQueryResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DbQueryResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DbQueryResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DbQueryResult) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


