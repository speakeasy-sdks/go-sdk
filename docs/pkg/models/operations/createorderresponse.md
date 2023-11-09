# CreateOrderResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ContentType`                                                        | *string*                                                             | :heavy_check_mark:                                                   | HTTP response content type for this operation                        |
| `ErrorResponse`                                                      | [*shared.ErrorResponse](../../../pkg/models/shared/errorresponse.md) | :heavy_minus_sign:                                                   | Any bad or invalid request will lead to following error object       |
| `Headers`                                                            | map[string][]*string*                                                | :heavy_minus_sign:                                                   | N/A                                                                  |
| `OrdersEntity`                                                       | [*shared.OrdersEntity](../../../pkg/models/shared/ordersentity.md)   | :heavy_minus_sign:                                                   | OK                                                                   |
| `StatusCode`                                                         | *int*                                                                | :heavy_check_mark:                                                   | HTTP response status code for this operation                         |
| `RawResponse`                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)               | :heavy_minus_sign:                                                   | Raw HTTP response; suitable for custom response parsing              |