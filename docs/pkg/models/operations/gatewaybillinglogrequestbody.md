# GatewayBillingLogRequestBody


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `ConsumerID`                                | *string*                                    | :heavy_check_mark:                          | The Consumer ID making the request          | usr_123xyz                                  |
| `APIID`                                     | *string*                                    | :heavy_check_mark:                          | The Developer Gateway API ID being consumed | api_chat_bot                                |
| `Metric`                                    | **string*                                   | :heavy_minus_sign:                          | Optional metric name (default: 'units')     | tokens                                      |
| `Units`                                     | *int64*                                     | :heavy_check_mark:                          | Amount of usage to log                      | 1500                                        |