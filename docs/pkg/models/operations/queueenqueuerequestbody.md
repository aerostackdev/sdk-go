# QueueEnqueueRequestBody


## Fields

| Field                              | Type                               | Required                           | Description                        | Example                            |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `Type`                             | *string*                           | :heavy_check_mark:                 | N/A                                | send-email                         |
| `Data`                             | map[string]*any*                   | :heavy_check_mark:                 | N/A                                |                                    |
| `Delay`                            | **int64*                           | :heavy_minus_sign:                 | Delay in seconds before processing | 60                                 |