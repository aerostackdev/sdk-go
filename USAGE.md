<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/pkg/models/operations"
	"undefined/pkg/models/shared"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(shared.Security{
			APIKeyAuth: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.Database.DbQuery(ctx, operations.DbQueryRequestBody{
		SQL: "SELECT * FROM users WHERE active = ?",
		Params: []any{
			true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.DbQueryResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->