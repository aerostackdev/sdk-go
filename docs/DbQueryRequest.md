# DbQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | **string** | SQL query to execute | 
**Params** | Pointer to **[]interface{}** | Query parameters for prepared statements | [optional] 

## Methods

### NewDbQueryRequest

`func NewDbQueryRequest(sql string, ) *DbQueryRequest`

NewDbQueryRequest instantiates a new DbQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbQueryRequestWithDefaults

`func NewDbQueryRequestWithDefaults() *DbQueryRequest`

NewDbQueryRequestWithDefaults instantiates a new DbQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *DbQueryRequest) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *DbQueryRequest) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *DbQueryRequest) SetSql(v string)`

SetSql sets Sql field to given value.


### GetParams

`func (o *DbQueryRequest) GetParams() []interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DbQueryRequest) GetParamsOk() (*[]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DbQueryRequest) SetParams(v []interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DbQueryRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


