# Orders

### Available Operations

* [CreateOrder](#createorder) - Create Order
* [OrderPay](#orderpay) - Order Pay

## CreateOrder

Use this API to create orders with Cashfree from your backend and get the payment link. To use this API S2S flag needs to be enabled from the backend. In case you want to use the cards payment option the PCI DSS flag is required, for more information email us at "care@cashfree.com".

### Example Usage

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
                CustomerBankAccountNumber: pg.String("sapiente"),
                CustomerBankCode: pg.String("quo"),
                CustomerBankIfsc: pg.String("odit"),
                CustomerEmail: pg.String("at"),
                CustomerID: "at",
                CustomerPhone: "maiores",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: pg.String("2021-07-29T00:00:00Z"),
            OrderID: pg.String("molestiae"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: pg.String("quod"),
                PaymentMethods: pg.String("quod"),
                ReturnURL: pg.String("esse"),
            },
            OrderNote: pg.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: pg.Float64(7805.29),
                    Percentage: pg.Float64(6788.8),
                    VendorID: pg.String("dicta"),
                },
                shared.VendorSplit{
                    Amount: pg.Float64(7206.33),
                    Percentage: pg.Float64(6399.21),
                    VendorID: pg.String("occaecati"),
                },
                shared.VendorSplit{
                    Amount: pg.Float64(1433.53),
                    Percentage: pg.Float64(5373.73),
                    VendorID: pg.String("hic"),
                },
            },
            OrderTags: map[string]string{
                "totam": "beatae",
                "commodi": "molestiae",
                "modi": "qui",
                "impedit": "cum",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "esse",
                TerminalPhoneNo: "ipsum",
                TerminalType: "excepturi",
            },
        },
        XAPIVersion: pg.String("aspernatur"),
        XClientID: "perferendis",
        XClientSecret: "ad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**


## OrderPay

Use this API when you have already created the orders and want Cashfree to process the payment. To use this API S2S flag needs to be enabled from the backend. In case you want to use the cards payment option the PCI DSS flag is required, for more information send an email to "care@cashfree.com".

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/go-sdk"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk/pkg/models/shared"
)

func main() {
    s := pg.New()

    ctx := context.Background()
    res, err := s.Orders.OrderPay(ctx, operations.OrderPayRequest{
        OrderPayRequest: &shared.OrderPayRequest{
            OfferID: pg.String("faa6cc05-d1e2-401c-b0cf-0c9db3ff0f0b"),
            PaymentMethod: shared.OrderPayRequestPaymentMethod{},
            PaymentSessionID: "session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn",
            SaveInstrument: pg.Bool(false),
        },
        XAPIVersion: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderPayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.OrderPayRequest](../../models/operations/orderpayrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.OrderPayResponse](../../models/operations/orderpayresponse.md), error**

