# \GatewayAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayBillingLog**](GatewayAPI.md#GatewayBillingLog) | **Post** /gateway/billing/log | Log Gateway usage



## GatewayBillingLog

> GatewayBillingLog200Response GatewayBillingLog(ctx).GatewayBillingLogRequest(gatewayBillingLogRequest).Execute()

Log Gateway usage



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
	gatewayBillingLogRequest := *openapiclient.NewGatewayBillingLogRequest("usr_123xyz", "api_chat_bot", int32(1500)) // GatewayBillingLogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GatewayBillingLog(context.Background()).GatewayBillingLogRequest(gatewayBillingLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GatewayBillingLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayBillingLog`: GatewayBillingLog200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GatewayBillingLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayBillingLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayBillingLogRequest** | [**GatewayBillingLogRequest**](GatewayBillingLogRequest.md) |  | 

### Return type

[**GatewayBillingLog200Response**](GatewayBillingLog200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

