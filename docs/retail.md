# Retail API Documentation (`moego.business.retail.v1`)

## 📌 1. Functional Overview

The Retail module manages retail products, including product information and supplier details. It provides the following
functions:

- Listing retail products with pagination
- Filtering by business locations and company ID
- Retrieving detailed product and supplier information

This module is applicable to scenarios such as inventory management, product catalog maintenance, and sales tracking.

---

## 🎯 2. Design Goals

- **Centralized Product Management**: Unified interface for retrieving product data
- **Flexible Filtering**: Supports filtering by business IDs
- **Easy Integration**: RESTful API compatible with mainstream frameworks
- **Secure Access**: Role-based access control ensures data integrity

---

## 🧩 3. Core Concepts

### 1. Product

Represents a retail product available for sale.

#### Fields

| Field Name                | Type                      | Description                                              |
|---------------------------|---------------------------|----------------------------------------------------------|
| `id`                      | string                    | Unique identifier for the product                        |
| `business_id`             | string                    | ID of the business location associated with this product |
| `name`                    | string                    | Name of the product                                      |
| `description`             | string                    | Description of the product                               |
| `sku`                     | string                    | Stock Keeping Unit (SKU)                                 |
| `barcode`                 | string                    | Barcode used for scanning                                |
| `image_url`               | string                    | URL to the product's image                               |
| `category_name`           | string                    | Name of the category the product belongs to              |
| `supplier`                | Supplier                  | Associated supplier information                          |
| `supply_price`            | google.type.Money         | Cost price from the supplier                             |
| `retail_price`            | google.type.Money         | Selling price to customers                               |
| `special_price`           | google.type.Money         | Discounted price for limited time                        |
| `tax_rate`                | double                    | Tax rate applied to this product                         |
| `enable_staff_commission` | bool                      | Whether staff commission is enabled                      |
| `stock`                   | int32                     | Current stock quantity                                   |
| `deleted`                 | bool                      | Whether the product has been marked as deleted           |
| `create_time`             | google.protobuf.Timestamp | When the product was created                             |
| `update_time`             | google.protobuf.Timestamp | When the product was last updated                        |

---

### 2. Supplier

Represents a product supplier.

#### Fields

| Field Name   | Type   | Description                        |
|--------------|--------|------------------------------------|
| `id`         | string | Unique identifier for the supplier |
| `name`       | string | Name of the supplier               |
| `first_name` | string | Contact person’s first name        |
| `last_name`  | string | Contact person’s last name         |
| `telephone`  | string | Landline telephone number          |
| `mobile`     | string | Mobile phone number                |
| `email`      | string | Email address                      |
| `website`    | string | Website URL                        |
| `address`    | string | Physical address                   |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Retail API

Here is a typical integration flow:

1. **ListProducts**
    - Retrieve a list of all products.
    - Filter by business locations or company ID if needed.
    - Use pagination to manage large datasets.

2. **Monitoring Inventory**
    - Regularly check stock levels.
    - Track product status (active/deleted).

3. **Supplier Management**
    - View associated supplier information.
    - Update supplier contact details.

---

## 📦 5. API Interface Descriptions

### 1. ListProducts (`ListProducts`)

- **Method**: `ListProducts`
- **HTTP Method**: POST
- **Path**: `/v1/retail/products:list`

#### ✅ Functionality:

Retrieves a paginated list of products based on specified criteria. Supports filtering by business locations.

#### 🎯 Use Cases:

- View all products in the system.
- Filter products by business location.
- Monitor product inventory.

#### 🔧 Request Parameters:

| Field Name     | Type          | Required | Description                                      |
|----------------|---------------|----------|--------------------------------------------------|
| `pagination`   | Pagination    | Yes      | Pagination info: page_size, page_token           |
| `company_id`   | string        | Yes      | Company ID to scope products                     |
| `business_ids` | Array(string) | Yes      | List of business locations to filter products by |

#### 📌 Return Value:

Returns a paginated list of products and a token for retrieving the next page.

| Field Name        | Type           | Description                                                          |
|-------------------|----------------|----------------------------------------------------------------------|
| `next_page_token` | string         | Token for retrieving the next page of results (empty if none remain) |
| `products`        | Array(Product) | List of products matching the request criteria                       |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid.
- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: ListProducts

```json
{
  "pagination": {
    "page_size": 20
  },
  "company_id": "cmp_001",
  "business_ids": [
    "biz_001",
    "biz_002"
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                                 | Answer                                                             |
|----------------------------------------------------------|--------------------------------------------------------------------|
| How to verify if a product exists?                       | Use `ListProducts` with specific filters                           |
| Can I list products for multiple companies at once?      | Currently only supports listing products for one company at a time |
| How to filter products by business location?             | Use `ListProducts` with `business_ids`                             |
| Why does creating a product return “resource exhausted”? | Not applicable — products are typically not created via this API   |
| How to handle out-of-stock products?                     | Use `ListProducts` with `stock = 0` filter                         |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Product ID does not exist         |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/pagination.md)
- [retail_service.proto](../moego/business/retail/v1/retail_service.proto)
- [product.proto](../moego/business/retail/v1/product.proto)
