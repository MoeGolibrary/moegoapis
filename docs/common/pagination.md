# 📐 Pagination (moego.common.v1.Pagination)

## 📌 1. Overview

The `Pagination` message provides a standardized way to handle paginated results across all API endpoints. It supports
both cursor-based and offset-based pagination, with configurable page sizes and sorting options. This ensures consistent
and efficient data retrieval for large result sets.

## 🧩 2. Core Fields

| Field Name  | Type   | Required | Description                                                      |
|-------------|--------|----------|------------------------------------------------------------------|
| `pageSize`  | int32  | Yes      | Number of items to return per page (1~500)                       |
| `pageToken` | string | No       | Token for retrieving the next page (min length 1, max length 64) |

> ⚠️ `pageToken` is optional. Leave empty for the first page. Obtain from previous response for subsequent pages.

## 📦 3. Example JSON

```json
{
  "pageSize": 20,
  "pageToken": "next_page_token_here"
}
```

---

# 📎 Related File References

- [pagination.proto](../../moego/common/v1/pagination.proto)
