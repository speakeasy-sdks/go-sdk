# App


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Channel`                                                         | *string*                                                          | :heavy_check_mark:                                                | Specify the channel through which the payment must be processed.  |
| `Phone`                                                           | *string*                                                          | :heavy_check_mark:                                                | Customer phone number associated with a wallet for payment.       |
| `Provider`                                                        | [shared.Provider](../../../pkg/models/shared/provider.md)         | :heavy_check_mark:                                                | Specify the provider through which the payment must be processed. |