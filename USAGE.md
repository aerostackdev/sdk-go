<!-- Start SDK Example Usage [usage] -->
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