# \CacheAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheGet**](CacheAPI.md#CacheGet) | **Post** /cache/get | Get cached value
[**CacheSet**](CacheAPI.md#CacheSet) | **Post** /cache/set | Set cached value



## CacheGet

> CacheGet200Response CacheGet(ctx).CacheGetRequest(cacheGetRequest).Execute()

Get cached value

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
	cacheGetRequest := *openapiclient.NewCacheGetRequest("user:123:profile") // CacheGetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CacheAPI.CacheGet(context.Background()).CacheGetRequest(cacheGetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheAPI.CacheGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheGet`: CacheGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CacheAPI.CacheGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCacheGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cacheGetRequest** | [**CacheGetRequest**](CacheGetRequest.md) |  | 

### Return type

[**CacheGet200Response**](CacheGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheSet

> CacheSet200Response CacheSet(ctx).CacheSetRequest(cacheSetRequest).Execute()

Set cached value

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
	cacheSetRequest := *openapiclient.NewCacheSetRequest("Key_example", interface{}(123)) // CacheSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CacheAPI.CacheSet(context.Background()).CacheSetRequest(cacheSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheAPI.CacheSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheSet`: CacheSet200Response
	fmt.Fprintf(os.Stdout, "Response from `CacheAPI.CacheSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCacheSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cacheSetRequest** | [**CacheSetRequest**](CacheSetRequest.md) |  | 

### Return type

[**CacheSet200Response**](CacheSet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

