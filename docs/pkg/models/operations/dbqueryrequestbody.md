# DbQueryRequestBody


## Fields

| Field                                    | Type                                     | Required                                 | Description                              | Example                                  |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `SQL`                                    | *string*                                 | :heavy_check_mark:                       | SQL query to execute                     | SELECT * FROM users WHERE active = ?     |
| `Params`                                 | []*any*                                  | :heavy_minus_sign:                       | Query parameters for prepared statements | [<br/>true<br/>]                         |