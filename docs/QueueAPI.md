# \QueueAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueueEnqueue**](QueueAPI.md#QueueEnqueue) | **Post** /queue/enqueue | Add job to queue



## QueueEnqueue

> QueueEnqueue201Response QueueEnqueue(ctx).QueueEnqueueRequest(queueEnqueueRequest).Execute()

Add job to queue

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
	queueEnqueueRequest := *openapiclient.NewQueueEnqueueRequest("send-email", map[string]interface{}(123)) // QueueEnqueueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.QueueEnqueue(context.Background()).QueueEnqueueRequest(queueEnqueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.QueueEnqueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueueEnqueue`: QueueEnqueue201Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.QueueEnqueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueueEnqueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queueEnqueueRequest** | [**QueueEnqueueRequest**](QueueEnqueueRequest.md) |  | 

### Return type

[**QueueEnqueue201Response**](QueueEnqueue201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

