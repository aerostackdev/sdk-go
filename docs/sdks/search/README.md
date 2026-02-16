# Ai.Search

## Overview

### Available Operations

* [Ingest](#ingest) - Ingest content into managed search index
* [Query](#query) - Search managed index
* [Delete](#delete) - Delete item by ID
* [DeleteByType](#deletebytype) - Delete all items of a type
* [ListTypes](#listtypes) - List distinct types and counts
* [Configure](#configure) - Update search configuration

## Ingest

Ingest content into managed search index

### Example Usage

<!-- UsageSnippet language="go" operationID="ingest" method="post" path="/ai/search/ingest" -->
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

    res, err := s.Ai.Search.Ingest(ctx, operations.IngestRequestBody{
        Content: "<value>",
        Type: "<value>",
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
| `request`                                                                        | [operations.IngestRequestBody](../../pkg/models/operations/ingestrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.IngestResponse](../../pkg/models/operations/ingestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Query

Search managed index

### Example Usage

<!-- UsageSnippet language="go" operationID="query" method="post" path="/ai/search/query" -->
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

    res, err := s.Ai.Search.Query(ctx, operations.QueryRequestBody{
        Text: "<value>",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.QueryRequestBody](../../pkg/models/operations/queryrequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.QueryResponse](../../pkg/models/operations/queryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete item by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete" method="post" path="/ai/search/delete" -->
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

    res, err := s.Ai.Search.Delete(ctx, operations.DeleteRequestBody{
        ID: "<id>",
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
| `request`                                                                        | [operations.DeleteRequestBody](../../pkg/models/operations/deleterequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DeleteResponse](../../pkg/models/operations/deleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteByType

Delete all items of a type

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteByType" method="post" path="/ai/search/deleteByType" -->
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

    res, err := s.Ai.Search.DeleteByType(ctx, operations.DeleteByTypeRequestBody{
        Type: "<value>",
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
| `request`                                                                                    | [operations.DeleteByTypeRequestBody](../../pkg/models/operations/deletebytyperequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeleteByTypeResponse](../../pkg/models/operations/deletebytyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTypes

List distinct types and counts

### Example Usage

<!-- UsageSnippet language="go" operationID="listTypes" method="get" path="/ai/search/listTypes" -->
```go
package main

import(
	"context"
	"aerostack/pkg/models/shared"
	"aerostack"
	"log"
)

func main() {
    ctx := context.Background()

    s := aerostack.New(
        aerostack.WithSecurity(shared.Security{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Ai.Search.ListTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.ListTypesResponse](../../pkg/models/operations/listtypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Configure

Update search configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="configure" method="post" path="/ai/search/configure" -->
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

    res, err := s.Ai.Search.Configure(ctx, operations.ConfigureRequestBody{
        EmbeddingModel: operations.EmbeddingModelMultilingual,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ConfigureRequestBody](../../pkg/models/operations/configurerequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ConfigureResponse](../../pkg/models/operations/configureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |