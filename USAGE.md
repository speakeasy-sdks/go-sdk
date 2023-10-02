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
                CustomerBankAccountNumber: gosdk.String("North double"),
                CustomerBankCode: gosdk.String("spherical woman burdensome"),
                CustomerBankIfsc: gosdk.String("interfaces Smart"),
                CustomerEmail: gosdk.String("Doyle brown toast"),
                CustomerID: "Bedfordshire",
                CustomerPhone: "Mohr North",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: gosdk.String("2021-07-29T00:00:00Z"),
            OrderID: gosdk.String("deploy South"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: gosdk.String("Road male Berkshire"),
                PaymentMethods: gosdk.String("parsing female middleware"),
                ReturnURL: gosdk.String("Bedfordshire navigating"),
            },
            OrderNote: gosdk.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: gosdk.Float64(5942.72),
                    Percentage: gosdk.Float64(3302.96),
                    VendorID: gosdk.String("dearly remount"),
                },
            },
            OrderTags: map[string]string{
                "expedita": "South",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "Southwest",
                TerminalPhoneNo: "violet Chips Porsche",
                TerminalType: "mobile",
            },
        },
        XAPIVersion: gosdk.String("ROI bypassing vero"),
        XClientID: "Solutions Ferrari Accountability",
        XClientSecret: "Folk ampere",
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