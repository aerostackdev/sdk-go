# Database

## Overview

SQL database operations

### Available Operations

* [DbQuery](#dbquery) - Execute SQL query

## DbQuery

Run a SQL query against your project database

### Example Usage

<!-- UsageSnippet language="go" operationID="dbQuery" method="post" path="/db/query" -->
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

    res, err := s.Database.DbQuery(ctx, operations.DbQueryRequestBody{
        SQL: "SELECT * FROM users WHERE active = ?",
        Params: []any{
            true,
        },
    }, nil, aerostack.Pointer("0.1.0"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DbQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `requestBody`                                                                      | [operations.DbQueryRequestBody](../../pkg/models/operations/dbqueryrequestbody.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `xRequestID`                                                                       | **string*                                                                          | :heavy_minus_sign:                                                                 | Unique request tracing ID                                                          |                                                                                    |
| `xSDKVersion`                                                                      | **string*                                                                          | :heavy_minus_sign:                                                                 | SDK version string                                                                 | 0.1.0                                                                              |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DbQueryResponse](../../pkg/models/operations/dbqueryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400, 401                | application/json        |
| sdkerrors.ErrorResponse | 500                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |