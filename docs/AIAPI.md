# \AIAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiChat**](AIAPI.md#AiChat) | **Post** /ai/chat | Generate AI chat completion
[**Configure**](AIAPI.md#Configure) | **Post** /ai/search/configure | Update search configuration
[**Delete**](AIAPI.md#Delete) | **Post** /ai/search/delete | Delete item by ID
[**DeleteByType**](AIAPI.md#DeleteByType) | **Post** /ai/search/deleteByType | Delete all items of a type
[**Ingest**](AIAPI.md#Ingest) | **Post** /ai/search/ingest | Ingest content into managed search index
[**ListTypes**](AIAPI.md#ListTypes) | **Get** /ai/search/listTypes | List distinct types and counts
[**Query**](AIAPI.md#Query) | **Post** /ai/search/query | Search managed index



## AiChat

> AiChat200Response AiChat(ctx).AiChatRequest(aiChatRequest).Execute()

Generate AI chat completion

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
	aiChatRequest := *openapiclient.NewAiChatRequest([]openapiclient.AiChatRequestMessagesInner{*openapiclient.NewAiChatRequestMessagesInner()}) // AiChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.AiChat(context.Background()).AiChatRequest(aiChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.AiChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiChat`: AiChat200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.AiChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiChatRequest** | [**AiChatRequest**](AiChatRequest.md) |  | 

### Return type

[**AiChat200Response**](AiChat200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Configure

> CacheSet200Response Configure(ctx).ConfigureRequest(configureRequest).Execute()

Update search configuration

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
	configureRequest := *openapiclient.NewConfigureRequest("EmbeddingModel_example") // ConfigureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.Configure(context.Background()).ConfigureRequest(configureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.Configure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Configure`: CacheSet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.Configure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureRequest** | [**ConfigureRequest**](ConfigureRequest.md) |  | 

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


## Delete

> CacheSet200Response Delete(ctx).DeleteRequest(deleteRequest).Execute()

Delete item by ID

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
	deleteRequest := *openapiclient.NewDeleteRequest("Id_example") // DeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.Delete(context.Background()).DeleteRequest(deleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: CacheSet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRequest** | [**DeleteRequest**](DeleteRequest.md) |  | 

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


## DeleteByType

> CacheSet200Response DeleteByType(ctx).DeleteByTypeRequest(deleteByTypeRequest).Execute()

Delete all items of a type

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
	deleteByTypeRequest := *openapiclient.NewDeleteByTypeRequest("Type_example") // DeleteByTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.DeleteByType(context.Background()).DeleteByTypeRequest(deleteByTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.DeleteByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteByType`: CacheSet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.DeleteByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteByTypeRequest** | [**DeleteByTypeRequest**](DeleteByTypeRequest.md) |  | 

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


## Ingest

> CacheSet200Response Ingest(ctx).IngestRequest(ingestRequest).Execute()

Ingest content into managed search index

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
	ingestRequest := *openapiclient.NewIngestRequest("Content_example", "Type_example") // IngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.Ingest(context.Background()).IngestRequest(ingestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.Ingest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ingest`: CacheSet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.Ingest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestRequest** | [**IngestRequest**](IngestRequest.md) |  | 

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


## ListTypes

> ListTypes200Response ListTypes(ctx).Execute()

List distinct types and counts

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.ListTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.ListTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTypes`: ListTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.ListTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTypesRequest struct via the builder pattern


### Return type

[**ListTypes200Response**](ListTypes200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> Query200Response Query(ctx).QueryRequest(queryRequest).Execute()

Search managed index

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
	queryRequest := *openapiclient.NewQueryRequest("Text_example") // QueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.Query(context.Background()).QueryRequest(queryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.Query``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Query`: Query200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.Query`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryRequest** | [**QueryRequest**](QueryRequest.md) |  | 

### Return type

[**Query200Response**](Query200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

