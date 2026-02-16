# Queue

## Overview

Background job queue

### Available Operations

* [QueueEnqueue](#queueenqueue) - Add job to queue

## QueueEnqueue

Add job to queue

### Example Usage

<!-- UsageSnippet language="go" operationID="queueEnqueue" method="post" path="/queue/enqueue" -->
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

    res, err := s.Queue.QueueEnqueue(ctx, operations.QueueEnqueueRequestBody{
        Type: "send-email",
        Data: map[string]any{
            "key": "<value>",
        },
        Delay: aerostack.Pointer[int64](60),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.QueueEnqueueRequestBody](../../pkg/models/operations/queueenqueuerequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.QueueEnqueueResponse](../../pkg/models/operations/queueenqueueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |