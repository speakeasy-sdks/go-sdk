<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/speakeasy-sdks/go-sdk/v3"
	"github.com/speakeasy-sdks/go-sdk/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk/v3/pkg/models/shared"
	"log"
)

func main() {
	s := gosdk.New()

	ctx := context.Background()
	res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
		CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
			CustomerDetails: shared.CustomerDetails{
				CustomerID:    "string",
				CustomerPhone: "string",
			},
			OrderAmount:     10.15,
			OrderCurrency:   "INR",
			OrderExpiryTime: gosdk.String("2021-07-29T00:00:00Z"),
			OrderMeta:       &shared.OrderMeta{},
			OrderNote:       gosdk.String("Test order"),
			OrderSplits: []shared.VendorSplit{
				shared.VendorSplit{},
			},
			OrderTags: map[string]string{
				"key": "string",
			},
			Terminal: &shared.TerminalDetails{
				TerminalID:      "string",
				TerminalPhoneNo: "string",
				TerminalType:    "string",
			},
		},
		XClientID:     "string",
		XClientSecret: "string",
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