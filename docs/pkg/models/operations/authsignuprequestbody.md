# AuthSignupRequestBody


## Fields

| Field              | Type               | Required           | Description        | Example            |
| ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| `Email`            | *string*           | :heavy_check_mark: | N/A                | user@example.com   |
| `Password`         | *string*           | :heavy_check_mark: | N/A                | SecurePass123!     |
| `Name`             | **string*          | :heavy_minus_sign: | N/A                | John Doe           |
| `Metadata`         | map[string]*any*   | :heavy_minus_sign: | N/A                |                    |