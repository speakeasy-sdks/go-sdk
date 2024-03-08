<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/speakeasy-sdks/go-sdk/v3"
	"github.com/speakeasy-sdks/go-sdk/v3/pkg/models/operations"
	"log"
)

func main() {
	s := gosdk.New()

	ctx := context.Background()
	res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
		XClientID:     "<value>",
		XClientSecret: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OrdersEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->