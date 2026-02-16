# Ai

## Overview

AI/ML operations

### Available Operations

* [AiChat](#aichat) - Generate AI chat completion

## AiChat

Generate AI chat completion

### Example Usage

<!-- UsageSnippet language="go" operationID="aiChat" method="post" path="/ai/chat" -->
```go
package main

import(
	"context"
	"aerostack/pkg/models/shared"
	"aerostack"
	"aerostack/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := aerostack.New(
        aerostack.WithSecurity(shared.Security{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Ai.AiChat(ctx, operations.AiChatRequestBody{
        Model: aerostack.Pointer("@cf/meta/llama-3-8b-instruct"),
        Messages: []operations.Messages{
            operations.Messages{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AiChatRequestBody](../../pkg/models/operations/aichatrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AiChatResponse](../../pkg/models/operations/aichatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |