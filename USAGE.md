<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/go-sdk"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/callbacks"
	"net/http"
)

func main() {
    s := pg.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: pg.String("corrupti"),
                CustomerBankCode: pg.String("provident"),
                CustomerBankIfsc: pg.String("distinctio"),
                CustomerEmail: pg.String("quibusdam"),
                CustomerID: "unde",
                CustomerPhone: "nulla",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: pg.String("2021-07-29T00:00:00Z"),
            OrderID: pg.String("corrupti"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: pg.String("illum"),
                PaymentMethods: pg.String("vel"),
                ReturnURL: pg.String("error"),
            },
            OrderNote: pg.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: pg.Float64(6458.94),
                    Percentage: pg.Float64(3843.82),
                    VendorID: pg.String("iure"),
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
        XAPIVersion: pg.String("suscipit"),
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