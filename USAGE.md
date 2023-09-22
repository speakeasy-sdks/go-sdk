<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	gosdk "github.com/speakeasy-sdks/go-sdk"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/callbacks"
	"net/http"
)

func main() {
    s := gosdk.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: gosdk.String("corrupti"),
                CustomerBankCode: gosdk.String("provident"),
                CustomerBankIfsc: gosdk.String("distinctio"),
                CustomerEmail: gosdk.String("quibusdam"),
                CustomerID: "unde",
                CustomerPhone: "nulla",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: gosdk.String("2021-07-29T00:00:00Z"),
            OrderID: gosdk.String("corrupti"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: gosdk.String("illum"),
                PaymentMethods: gosdk.String("vel"),
                ReturnURL: gosdk.String("error"),
            },
            OrderNote: gosdk.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: gosdk.Float64(6458.94),
                    Percentage: gosdk.Float64(3843.82),
                    VendorID: gosdk.String("iure"),
                },
            },
            OrderTags: map[string]string{
                "magnam": "debitis",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "ipsa",
                TerminalPhoneNo: "delectus",
                TerminalType: "tempora",
            },
        },
        XAPIVersion: gosdk.String("suscipit"),
        XClientID: "molestiae",
        XClientSecret: "minus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->