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
                    Amount: pg.Float64(3843.82),
                    Percentage: pg.Float64(4375.87),
                    VendorID: pg.String("magnam"),
                },
                shared.VendorSplit{
                    Amount: pg.Float64(8917.73),
                    Percentage: pg.Float64(567.13),
                    VendorID: pg.String("delectus"),
                },
                shared.VendorSplit{
                    Amount: pg.Float64(2726.56),
                    Percentage: pg.Float64(3834.41),
                    VendorID: pg.String("molestiae"),
                },
            },
            OrderTags: map[string]string{
                "placeat": "voluptatum",
                "iusto": "excepturi",
                "nisi": "recusandae",
                "temporibus": "ab",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "quis",
                TerminalPhoneNo: "veritatis",
                TerminalType: "deserunt",
            },
        },
        XAPIVersion: pg.String("perferendis"),
        XClientID: "ipsam",
        XClientSecret: "repellendus",
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