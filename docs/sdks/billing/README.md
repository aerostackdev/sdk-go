# Gateway.Billing

## Overview

### Available Operations

* [GatewayBillingLog](#gatewaybillinglog) - Log Gateway usage

## GatewayBillingLog

Manually log tokens or custom metric usage for a Gateway API

### Example Usage

<!-- UsageSnippet language="go" operationID="gatewayBillingLog" method="post" path="/gateway/billing/log" -->
```go
package main

import(
	"context"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := aerostack.New(
        aerostack.WithSecurity(shared.Security{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Gateway.Billing.GatewayBillingLog(ctx, operations.GatewayBillingLogRequestBody{
        ConsumerID: "usr_123xyz",
        APIID: "api_chat_bot",
        Metric: aerostack.Pointer("tokens"),
        Units: 1500,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GatewayBillingLogRequestBody](../../pkg/models/operations/gatewaybillinglogrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GatewayBillingLogResponse](../../pkg/models/operations/gatewaybillinglogresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400, 401                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |