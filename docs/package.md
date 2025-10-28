# Package API Documentation (`moego.business.package.v1`)

## 📌 1. Functional Overview

The Package module manages customer packages, including package information, services included in packages, and package
usage details. It provides the following functions:

- Retrieving basic package information for customers
- Getting detailed package information including services
- Listing packages with filtering and pagination capabilities

This module is applicable to scenarios such as package management, service tracking, and customer benefits
administration.

---

## 🎯 2. Design Goals

- **Customer-Centric Package Management**: Focus on packages from the customer perspective
- **Detailed Package Information**: Provide comprehensive details about package services and usage
- **Easy Integration**: RESTful API compatible with mainstream frameworks
- **Secure Access**: Role-based access control ensures data integrity

---

## 🧩 3. Core Concepts

### 1. Package

Represents a package purchased by a customer.

#### Fields

| Field Name                 | Type                      | Description                                                                            |
|----------------------------|---------------------------|----------------------------------------------------------------------------------------|
| `id`                       | string                    | Unique identifier for the package                                                      |
| `customer_id`              | string                    | ID of the customer who purchased this package                                          |
| `business_id`              | string                    | ID of the business that owns this package                                              |
| `staff_id`                 | string                    | ID of the staff who handled this package                                               |
| `package_name`             | string                    | Name of the package                                                                    |
| `package_desc`             | string                    | Description of the package                                                             |
| `package_price`            | google.type.Money         | Price of the package                                                                   |
| `purchase_time`            | google.protobuf.Timestamp | Purchase time of the package                                                           |
| `start_time`               | google.protobuf.Timestamp | Start time of the package validity                                                     |
| `end_time`                 | google.protobuf.Timestamp | End time of the package validity                                                       |
| `create_time`              | google.protobuf.Timestamp | Creation time of the package record                                                    |
| `last_update_time`         | google.protobuf.Timestamp | Last update time of the package record                                                 |
| `expiration_date`          | google.type.Date          | Expiration date of the package in format: yyyy-MM-dd ."9999-01-01" means never expired |
| `status`                   | enum Status               | Status of the package                                                                  |
| `used`                     | bool                      | Whether the package has been used                                                      |
| `applied`                  | bool                      | Whether the package is applied                                                         |
| `total_remaining_quantity` | int32                     | Total remaining quantity of services in the package                                    |

### 2. PackageDetail

Represents detailed information about a package including its services.

#### Fields

| Field Name         | Type                    | Description                              |
|--------------------|-------------------------|------------------------------------------|
| `package_info`     | Package                 | Basic package information                |
| `package_services` | repeated PackageService | List of services included in the package |

### 3. PackageService

Represents a service included in a package.

#### Fields

| Field Name           | Type             | Description                                       |
|----------------------|------------------|---------------------------------------------------|
| `id`                 | string           | Unique identifier for the package service         |
| `package_id`         | string           | ID of the package this service belongs to         |
| `services`           | repeated Service | List of services included in this package item    |
| `total_quantity`     | int32            | Total quantity of this service in the package     |
| `remaining_quantity` | int32            | Remaining quantity of this service in the package |

### 4. Service

Represents a service included in a package.

#### Fields

| Field Name   | Type              | Description               |
|--------------|-------------------|---------------------------|
| `service_id` | string            | ID of the service         |
| `unit_price` | google.type.Money | Unit price of the service |
| `name`       | string            | Name of the service       |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Package API

Here is a typical integration flow:

1. **ListPackages**
    - Retrieve a list of packages for a specific customer.
    - Filter by business ID and customer ID.
    - Use pagination to manage large datasets.

2. **Get Package Details**
    - Retrieve detailed information about specific packages.
    - Get information about services included in packages.

---

## 📦 5. API Interface Descriptions

### 1. ListPackages (`ListPackages`)

- **Method**: `ListPackages`
- **HTTP Method**: POST
- **Path**: `/v1/packages:list`

#### ✅ Functionality:

Retrieves a list of packages for a specific customer. Supports filtering by business ID and customer ID.

#### 🎯 Use Cases:

- View all packages purchased by a customer.
- Filter packages by business location.
- Track package status and usage.

#### 🔧 Request Parameters:

| Field Name     | Type            | Required | Description                          |
|----------------|-----------------|----------|--------------------------------------|
| `pagination`   | Pagination      | Yes      | Pagination info: pageSize, pageToken |
| `company_id`   | string          | Yes      | Company ID to scope packages         |
| `customer_ids` | repeated string | Yes      | Customer IDs to filter packages      |

#### 📌 Return Value:

| Field Name        | Type           | Description                                                          |
|-------------------|----------------|----------------------------------------------------------------------|
| `packages`        | Array(Package) | List of packages matching the request criteria                       |
| `next_page_token` | string         | Token for retrieving the next page of results (empty if none remain) |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid
- `PERMISSION_DENIED`: Permission denied

### 2. ListPackageDetails (`ListPackageDetails`)

- **Method**: `ListPackageDetails`
- **HTTP Method**: POST
- **Path**: `/v1/packages/details:list`

#### ✅ Functionality:

Retrieves detailed information for a list of packages including services.

#### 🎯 Use Cases:

- View detailed information about specific packages.
- Get information about services included in packages.
- Track package service usage and remaining quantities.

#### 🔧 Request Parameters:

| Field Name    | Type            | Required | Description                                 |
|---------------|-----------------|----------|---------------------------------------------|
| `pagination`  | Pagination      | Yes      | Pagination info: pageSize, pageToken        |
| `company_id`  | string          | Yes      | Company ID to scope packages                |
| `package_ids` | repeated string | Yes      | List of package IDs to retrieve details for |

#### 📌 Return Value:

| Field Name        | Type                 | Description                                                          |
|-------------------|----------------------|----------------------------------------------------------------------|
| `package_details` | Array(PackageDetail) | List of package details matching the request criteria                |
| `next_page_token` | string               | Token for retrieving the next page of results (empty if none remain) |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid or package IDs not provided
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: ListPackages

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "customerIds": [
    "cus_123"
  ]
}
```

Response:

```json
{
  "packages": [
    {
      "id": "pkg_abc123",
      "customerId": "cus_123",
      "businessId": "bus_001",
      "staffId": "staff_001",
      "packageName": "Basic Grooming Package",
      "packageDesc": "Includes 5 basic grooming sessions",
      "packagePrice": {
        "currencyCode": "USD",
        "units": 150,
        "nanos": 0
      },
      "purchaseTime": "2023-01-15T10:00:00Z",
      "startTime": "2023-01-15T10:00:00Z",
      "endTime": "2024-01-15T10:00:00Z",
      "createTime": "2023-01-15T10:00:00Z",
      "lastUpdateTime": "2023-01-15T10:00:00Z",
      "expirationDate": "2024-01-15",
      "status": "STATUS_NORMAL",
      "used": true,
      "applied": true,
      "totalRemainingQuantity": 3
    }
  ],
  "nextPageToken": ""
}
```

### Example 2: ListPackageDetails

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "packageIds": [
    "pkg_abc123"
  ]
}
```

Response:

```json
{
  "packageDetails": [
    {
      "packageInfo": {
        "id": "pkg_abc123",
        "customerId": "cus_123",
        "businessId": "bus_001",
        "staffId": "staff_001",
        "packageName": "Basic Grooming Package",
        "packageDesc": "Includes 5 basic grooming sessions",
        "packagePrice": {
          "currencyCode": "USD",
          "units": 150,
          "nanos": 0
        },
        "purchaseTime": "2023-01-15T10:00:00Z",
        "startTime": "2023-01-15T10:00:00Z",
        "endTime": "2024-01-15T10:00:00Z",
        "createTime": "2023-01-15T10:00:00Z",
        "lastUpdateTime": "2023-01-15T10:00:00Z",
        "expirationDate": "2024-01-15",
        "status": "STATUS_NORMAL",
        "used": true,
        "applied": true,
        "totalRemainingQuantity": 3
      },
      "packageServices": [
        {
          "id": "ps_001",
          "packageId": "pkg_abc123",
          "services": [
            {
              "serviceId": "srv_001",
              "unitPrice": {
                "currencyCode": "USD",
                "units": 30,
                "nanos": 0
              },
              "name": "Basic Grooming"
            }
          ],
          "totalQuantity": 5,
          "remainingQuantity": 3
        }
      ]
    }
  ],
  "nextPageToken": ""
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                            | Answer                                                              |
|-----------------------------------------------------|---------------------------------------------------------------------|
| How to verify if a package exists?                  | Use `ListPackages` with specific filters                            |
| Can I list packages for multiple customers at once? | Currently only supports listing packages for one customer at a time |
| How to get detailed information about a package?    | Use `ListPackageDetails` with specific package IDs                  |
| How to handle expired packages?                     | Check the `expirationDate` field                                    |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Package ID does not exist         |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/pagination.md)
- [package_service.proto](../moego/business/package/v1/package_service.proto)
- [package.proto](../moego/business/package/v1/package.proto)