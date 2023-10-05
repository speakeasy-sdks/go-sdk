# github.com/speakeasy-sdks/go-sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create Order
* [OrderPay](docs/sdks/orders/README.md#orderpay) - Order Pay
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
