# @aerostack/sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *Aerostack* API.

[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

<br /><br />
> [!IMPORTANT]

<!-- Start Summary [summary] -->
## Summary

Aerostack API: Aerostack Platform API - Unified access to database, authentication, 
caching, queues, storage, and AI services.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [@aerostack/sdk-go](#aerostacksdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Backend Service Integration](#backend-service-integration)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/aerostackdev/sdks/packages/go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type   | Scheme  |
| ------------ | ------ | ------- |
| `APIKeyAuth` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Ai](docs/sdks/ai/README.md)

* [AiChat](docs/sdks/ai/README.md#aichat) - Generate AI chat completion

### [Ai.Search](docs/sdks/search/README.md)

* [Ingest](docs/sdks/search/README.md#ingest) - Ingest content into managed search index
* [Query](docs/sdks/search/README.md#query) - Search managed index
* [Delete](docs/sdks/search/README.md#delete) - Delete item by ID
* [DeleteByType](docs/sdks/search/README.md#deletebytype) - Delete all items of a type
* [ListTypes](docs/sdks/search/README.md#listtypes) - List distinct types and counts
* [Configure](docs/sdks/search/README.md#configure) - Update search configuration

### [Authentication](docs/sdks/authentication/README.md)

* [AuthSignup](docs/sdks/authentication/README.md#authsignup) - Sign up new user
* [AuthSignin](docs/sdks/authentication/README.md#authsignin) - Sign in user

### [Cache](docs/sdks/cache/README.md)

* [CacheGet](docs/sdks/cache/README.md#cacheget) - Get cached value
* [CacheSet](docs/sdks/cache/README.md#cacheset) - Set cached value

### [Database](docs/sdks/database/README.md)

* [DbQuery](docs/sdks/database/README.md#dbquery) - Execute SQL query

### [Gateway.Billing](docs/sdks/billing/README.md)

* [GatewayBillingLog](docs/sdks/billing/README.md#gatewaybillinglog) - Log Gateway usage

### [Queue](docs/sdks/queue/README.md)

* [QueueEnqueue](docs/sdks/queue/README.md#queueenqueue) - Add job to queue

### [Services](docs/sdks/services/README.md)

* [ServicesInvoke](docs/sdks/services/README.md#servicesinvoke) - Invoke another service

### [Storage](docs/sdks/storage/README.md)

* [StorageUpload](docs/sdks/storage/README.md#storageupload) - Upload file to storage

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	"github.com/aerostackdev/sdks/packages/go/pkg/retry"
	"log"
	"pkg/models/operations"
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
	}, nil, aerostack.Pointer("0.1.0"), operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.DbQueryResult != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	"github.com/aerostackdev/sdks/packages/go/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := aerostack.New(
		aerostack.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `DbQuery` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| sdkerrors.ErrorResponse | 400, 401    | application/json |
| sdkerrors.ErrorResponse | 500         | application/json |
| sdkerrors.SDKError      | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/sdkerrors"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
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

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                        | Description       |
| --- | ----------------------------- | ----------------- |
| 0   | `https://api.aerocall.ai/v1` | Production        |
| 1   | `http://localhost:8787/v1`    | Local Development |

#### Example

```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := aerostack.New(
		aerostack.WithServerIndex(0),
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := aerostack.New(
		aerostack.WithServerURL("http://localhost:8787/v1"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/aerostackdev/sdks/packages/go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = aerostack.New(aerostack.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->


## Backend Service Integration

> **SDK Type**: HTTP Client (API Calls)

This Go SDK is an **HTTP client** for calling Aerostack APIs from Go backends:

✅ **Use cases**:
- Go HTTP servers (net/http, Gin, Echo, Fiber) calling Aerostack APIs
- Go microservices integrating with Aerostack Auth/E-commerce
- CLI tools and background workers

❌ **Not for**:
- Direct Cloudflare binding access (use TypeScript `@aerostack/sdk` Server SDK in Workers)

For Go backends, this SDK is all you need. For Worker bindings, use TypeScript Workers.

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 


<!-- Placeholder for Future Speakeasy SDK Sections -->
