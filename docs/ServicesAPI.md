# \ServicesAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesInvoke**](ServicesAPI.md#ServicesInvoke) | **Post** /services/invoke | Invoke another service



## ServicesInvoke

> ServicesInvoke200Response ServicesInvoke(ctx).ServicesInvokeRequest(servicesInvokeRequest).Execute()

Invoke another service

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
	servicesInvokeRequest := *openapiclient.NewServicesInvokeRequest("billing-webhook", map[string]interface{}(123)) // ServicesInvokeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesInvoke(context.Background()).ServicesInvokeRequest(servicesInvokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesInvoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesInvoke`: ServicesInvoke200Response
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesInvoke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesInvokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicesInvokeRequest** | [**ServicesInvokeRequest**](ServicesInvokeRequest.md) |  | 

### Return type

[**ServicesInvoke200Response**](ServicesInvoke200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

