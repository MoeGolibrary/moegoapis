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

| Field Name              | Type                      | Description                                              |
|-------------------------|---------------------------|----------------------------------------------------------|
| `id`                    | string                    | Unique identifier for the product                        |
| `businessId`            | string                    | ID of the business location associated with this product |
| `name`                  | string                    | Name of the product                                      |
| `description`           | string                    | Description of the product                               |
| `sku`                   | string                    | Stock Keeping Unit (SKU)                                 |
| `barcode`               | string                    | Barcode used for scanning                                |
| `imageUrl`              | string                    | URL to the product's image                               |
| `categoryName`          | string                    | Name of the category the product belongs to              |
| `supplier`              | Supplier                  | Associated supplier information                          |
| `supplyPrice`           | google.type.Money         | Cost price from the supplier                             |
| `retailPrice`           | google.type.Money         | Selling price to customers                               |
| `specialPrice`          | google.type.Money         | Discounted price for limited time                        |
| `taxRate`               | double                    | Tax rate applied to this product                         |
| `enableStaffCommission` | bool                      | Whether staff commission is enabled                      |
| `stock`                 | int32                     | Current stock quantity                                   |
| `deleted`               | bool                      | Whether the product has been marked as deleted           |
| `createdTime`           | google.protobuf.Timestamp | When the product was created                             |
| `lastUpdatedTime`       | google.protobuf.Timestamp | When the product was last updated                        |

---

### 2. Supplier

Represents a product supplier.

#### Fields

| Field Name  | Type   | Description                        |
|-------------|--------|------------------------------------|
| `id`        | string | Unique identifier for the supplier |
| `name`      | string | Name of the supplier               |
| `firstName` | string | Contact person's first name        |
| `lastName`  | string | Contact person's last name         |
| `telephone` | string | Landline telephone number          |
| `mobile`    | string | Mobile phone number                |
| `email`     | string | Email address                      |
| `website`   | string | Website URL                        |
| `address`   | string | Physical address                   |

---

### 3. Package

Represents a package of services that can be sold together.

#### Fields

| Field Name              | Type                      | Description                                              |
|-------------------------|---------------------------|----------------------------------------------------------|
| `id`                    | string                    | Unique identifier for the package                        |
| `company_id`            | string                    | ID of the company that owns this package                 |
| `business_id`           | string                    | ID of the business location associated with this package |
| `name`                  | string                    | Name of the package                                      |
| `description`           | string                    | Description of the package                               |
| `price`                 | google.type.Money         | Price of the package                                     |
| `total_value`           | google.type.Money         | Total value of the services in the package               |
| `tax_id`                | string                    | Tax ID associated with this package                      |
| `tax_rate`              | double                    | Tax rate applied to this package                         |
| `sold_quantity`         | int32                     | Number of packages sold                                  |
| `create_time`           | google.protobuf.Timestamp | When this package was created                            |
| `update_time`           | google.protobuf.Timestamp | When this package was last updated                       |
| `is_active`             | bool                      | Whether this package is active                           |
| `enable_online_booking` | bool                      | Whether this package can be booked online                |
| `expiration_days`       | int32                     | Expiration days for the package                          |
| `source`                | int32                     | Source of the package                                    |
| `items`                 | repeated Item             | Items included in the package                            |

#### Item Object

| Field Name | Type             | Description                    |
|------------|------------------|--------------------------------|
| `quantity` | int32            | Quantity of this item          |
| `services` | repeated Service | Services included in this item |

#### Service Object

| Field Name   | Type              | Description               |
|--------------|-------------------|---------------------------|
| `service_id` | string            | ID of the service         |
| `unit_price` | google.type.Money | Unit price of the service |
| `name`       | string            | Name of the service       |

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

| Field Name    | Type          | Required | Description                                      |
|---------------|---------------|----------|--------------------------------------------------|
| `pagination`  | Pagination    | Yes      | Pagination info: pageSize, pageToken             |
| `companyId`   | string        | Yes      | Company ID to scope products                     |
| `businessIds` | Array(string) | Yes      | List of business locations to filter products by |

#### 📌 Return Value:

| Field Name      | Type           | Description                                                          |
|-----------------|----------------|----------------------------------------------------------------------|
| `nextPageToken` | string         | Token for retrieving the next page of results (empty if none remain) |
| `products`      | Array(Product) | List of products matching the request criteria                       |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid.
- `PERMISSION_DENIED`: Permission denied.

### 2. ListPackages (`ListPackages`)

- **Method**: `ListPackages`
- **HTTP Method**: POST
- **Path**: `/v1/retail/packages:list`

#### ✅ Functionality:

Retrieves a paginated list of packages based on specified criteria. Supports filtering by business locations and package
properties.

#### 🎯 Use Cases:

- View all packages in the system.
- Filter packages by business location.
- Search for specific packages by partial name match.

#### 🔧 Request Parameters:

| Field Name    | Type          | Required | Description                                      |
|---------------|---------------|----------|--------------------------------------------------|
| `pagination`  | Pagination    | Yes      | Pagination info: pageSize, pageToken             |
| `companyId`   | string        | Yes      | Company ID to scope packages                     |
| `businessIds` | Array(string) | Yes      | List of business locations to filter packages by |
| `filter`      | Filter        | No       | Optional filters to apply to the package list    |

##### Filter Object

| Field Name    | Type   | Required | Description                                                      |
|---------------|--------|----------|------------------------------------------------------------------|
| `name_like`   | string | No       | Partial name match to filter packages by name (case-insensitive) |
| `active_only` | bool   | No       | Filter to include only active packages                           |

#### 📌 Return Value:

| Field Name      | Type           | Description                                                          |
|-----------------|----------------|----------------------------------------------------------------------|
| `nextPageToken` | string         | Token for retrieving the next page of results (empty if none remain) |
| `packages`      | Array(Package) | List of packages matching the request criteria                       |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid.
- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: ListProducts

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "businessIds": [
    "biz_001",
    "biz_002"
  ]
}
```

### Example 2: ListPackages

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "businessIds": [
    "biz_001",
    "biz_002"
  ],
  "filter": {
    "name_like": "grooming",
    "active_only": true
  }
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
| How to filter products by business location?             | Use `ListProducts` with `businessIds`                              |
| Why does creating a product return "resource exhausted"? | Not applicable — products are typically not created via this API   |
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