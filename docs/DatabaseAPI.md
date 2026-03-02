# \DatabaseAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbQuery**](DatabaseAPI.md#DbQuery) | **Post** /db/query | Execute SQL query



## DbQuery

> DbQueryResult DbQuery(ctx).DbQueryRequest(dbQueryRequest).XRequestID(xRequestID).XSDKVersion(xSDKVersion).Execute()

Execute SQL query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dbQueryRequest := *openapiclient.NewDbQueryRequest("SELECT * FROM users WHERE active = ?") // DbQueryRequest | 
	xRequestID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique request tracing ID (optional)
	xSDKVersion := "0.1.0" // string | SDK version string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseAPI.DbQuery(context.Background()).DbQueryRequest(dbQueryRequest).XRequestID(xRequestID).XSDKVersion(xSDKVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAPI.DbQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbQuery`: DbQueryResult
	fmt.Fprintf(os.Stdout, "Response from `DatabaseAPI.DbQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbQueryRequest** | [**DbQueryRequest**](DbQueryRequest.md) |  | 
 **xRequestID** | **string** | Unique request tracing ID | 
 **xSDKVersion** | **string** | SDK version string | 

### Return type

[**DbQueryResult**](DbQueryResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

