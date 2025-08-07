# 🏢 Company API Documentation (`moego.business.company.v1`)

## 📌 1. Functional Overview

The **Company** module provides APIs to manage business entities in the system. Each company represents a top-level
organizational unit that can contain multiple business locations, staff members, and service offerings.

This interface enables:

- Retrieving detailed information about a specific company.
- Listing companies with pagination support.
- Managing company-wide configuration settings and operational parameters.

Useful for scenarios such as company management, multi-location administration, and integration with third-party
systems.

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing company profiles.
- **Standardized Data Model**: Ensures consistent handling of company-related data across services.
- **Secure and Reliable**: Enforces access control and data integrity.
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks.

Applicable to scenarios like company onboarding, configuration management, and enterprise-level reporting.

---

## 🧩 3. Core Concepts

### 1. `Company`

Represents a business entity in the system.

| Field Name | Type       | Description                                                                  |
|------------|------------|------------------------------------------------------------------------------|
| `id`       | `string`   | Unique identifier (e.g., `"cmp_001"`).                                       |
| `name`     | `string`   | Legal name used in official communications and documents.                    |
| `country`  | `string`   | ISO 3166-1 alpha-2 country code indicating where the company is registered.  |
| `timezone` | `TimeZone` | Primary timezone for company operations (used for scheduling and reporting). |

#### Example JSON

```json
{
  "id": "cmp_001",
  "name": "Moego Pet Care Inc.",
  "country": "US",
  "timezone": {
    "id": "America/New_York"
  }
}
```

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Company API

Here is a typical integration flow:

1. **Get Company**

- Retrieve detailed information about a specific company.
- Used during system setup or debugging.

2. **List Companies**

- View all companies accessible to the current user.
- Useful for system administration and auditing.

3. **Monitoring & Maintenance**

- Regularly retrieve company data to monitor changes.
- Ensure configurations are up-to-date.

---

## 🛠️ 5. API Interface Descriptions

### 1. Get Company (`GetCompany`)

- **Method**: `GetCompany`
- **HTTP Method**: GET
- **Path**: `/v1/companies/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific company.

#### 🎯 Use Cases:

- Check current company data.
- Verify company details during debugging.

#### 🔧 Request Parameters:

| Field Name | Type     | Required | Description                      |
|------------|----------|----------|----------------------------------|
| `id`       | `string` | Yes      | Unique identifier of the company |

#### 📌 Return Value:

| Field Name | Type      | Description                  |
|------------|-----------|------------------------------|
| `company`  | `Company` | The retrieved company object |

#### ⚠️ Error Codes:

| Error Code          | Description                         |
|---------------------|-------------------------------------|
| `NOT_FOUND`         | Specified company ID does not exist |
| `PERMISSION_DENIED` | Permission denied                   |

---

### 2. List Companies (`ListCompanies`)

- **Method**: `ListCompanies`
- **HTTP Method**: POST
- **Path**: `/v1/companies:list`

#### ✅ Functionality:

Lists companies based on specified criteria.

Results are paginated and can be used for system administration and multi-company management scenarios. Companies are
returned in alphabetical order by name.

#### 🎯 Use Cases:

- View all available companies.
- Audit or debug company configurations.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                          |
|--------------|------------|----------|--------------------------------------|
| `pagination` | Pagination | Yes      | Pagination info: pageSize, pageToken |

> ⚠️ `pageToken` is optional. Leave empty for the first page. Obtain from previous response for subsequent pages.

#### 📌 Return Value:

| Field Name      | Type             | Description                         |
|-----------------|------------------|-------------------------------------|
| `nextPageToken` | `string`         | Token for retrieving the next page  |
| `companies`     | Array(`Company`) | List of companies matching criteria |

#### ⚠️ Error Codes:

| Error Code          | Description                   |
|---------------------|-------------------------------|
| `INVALID_ARGUMENT`  | Invalid pagination parameters |
| `PERMISSION_DENIED` | Permission denied             |

---

## 🧪 6. Usage Examples

### ✅ Example 1: Retrieve a Specific Company

**Request**

```http
GET /v1/companies/cmp_001
```

**Response**

```json
{
  "id": "cmp_001",
  "name": "Moego Pet Care Inc.",
  "country": "US",
  "timezone": {
    "id": "America/New_York"
  }
}
```

---

### ✅ Example 2: List Companies with Pagination

**Request**

```http
POST /v1/companies:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 20,
    "pageToken": ""
  }
}
```

**Response**

```json
{
  "nextPageToken": "CBAQAA==",
  "companies": [
    {
      "id": "cmp_001",
      "name": "Moego Pet Care Inc.",
      "country": "US",
      "timezone": {
        "id": "America/New_York"
      }
    },
    {
      "id": "cmp_002",
      "name": "PetCare Australia Pty Ltd",
      "country": "AU",
      "timezone": {
        "id": "Australia/Sydney"
      }
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                               | Answer                                                                |
|--------------------------------------------------------|-----------------------------------------------------------------------|
| How to verify if a company exists?                     | Use `GetCompany` to check if the company ID returns a valid response. |
| Can I list all companies at once?                      | No. Results are paginated. Use `pageToken` to fetch next set.         |
| Why does listing companies return "permission denied"? | Ensure you have the correct access rights for this operation.         |
| How to handle large result sets efficiently?           | Use pagination via `pageSize` and `pageToken`.                        |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                      |
|---------------------|--------------------------------------------------|
| `NOT_FOUND`         | The requested company ID does not exist.         |
| `PERMISSION_DENIED` | Caller lacks access rights to perform operation. |
| `INVALID_ARGUMENT`  | Invalid request parameters provided.             |
| `INTERNAL`          | Internal server error occurred.                  |

---

## 📎 10. Related File References

- [company.proto](../moego/business/company/v1/company.proto)
- [company_service.proto](../moego/business/company/v1/company_service.proto)