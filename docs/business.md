# 🏢 Business API Documentation (`moego.business.business.v1`)

## 📌 1. Functional Overview

The **Business** module provides APIs to manage business locations and their profiles. Each business location represents
a physical or virtual place where services are provided to customers, operating under a parent company.

This interface enables:

- Retrieving detailed information about a specific business location.
- Listing all business locations associated with a company.
- Managing contact details, address, and social media presence for customer engagement.

Useful for scenarios such as business management, multi-location administration, and integration with third-party
systems.

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all business locations.
- **Rich Data Model**: Supports complex relationships like contact info, address, and social media links.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks.

Applicable to scenarios like location setup, customer-facing directories, and system administration.

---

## 🧩 3. Core Concepts

### 1. Business

Represents a physical or virtual location where services are provided.

| Field Name      | Type    | Description                                                 |
|-----------------|---------|-------------------------------------------------------------|
| `id`            | string  | Unique identifier (e.g., `"bus_001"`), obfuscated ID string |
| `name`          | string  | Display name of the business location                       |
| `avatar`        | string  | Photo URL of the business                                   |
| `phone`         | string  | Primary contact phone number (E.164 format)                 |
| `businessPhone` | string  | Public business phone number                                |
| `email`         | string  | Contact email address                                       |
| `address`       | Address | Physical location details                                   |
| `website`       | string  | Official website URL                                        |
| `facebook`      | string  | Facebook page URL                                           |
| `instagram`     | string  | Instagram profile URL                                       |
| `yelp`          | string  | Yelp business page URL                                      |
| `google`        | string  | Google Business Profile URL                                 |
| `companyId`     | string  | ID of the parent company, obfuscated ID string              |

> 📝 **Note**: The `Address` type is defined in `common.v1.Address`.

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Business API

Here is a typical integration flow:

1. **Get Business**
    - Retrieve detailed information about a specific business location.
    - Used during system setup or debugging.

2. **List Businesses**
    - View all business locations associated with a company.
    - Useful for system administration and auditing.

3. **Monitoring & Maintenance**
    - Regularly retrieve business data to monitor changes.
    - Update business records as needed.

---

## 🛠️ 5. API Interface Descriptions

### 1. Get Business (`GetBusiness`)

- **Method**: `GetBusiness`
- **HTTP Method**: GET
- **Path**: `/v1/businesses/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific business location.

#### 🎯 Use Cases:

- Check current business data.
- Verify business details during debugging.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                                   |
|------------|--------|----------|-----------------------------------------------|
| `id`       | string | Yes      | Business ID to retrieve, obfuscated ID string |

#### 📌 Return Value:

| Field Name | Type       | Description                   |
|------------|------------|-------------------------------|
| `business` | `Business` | The retrieved business object |

#### ⚠️ Error Codes:

| Error Code          | Description                          |
|---------------------|--------------------------------------|
| `NOT_FOUND`         | Specified business ID does not exist |
| `PERMISSION_DENIED` | Permission denied                    |

---

### 2. List Businesses (`ListBusiness`)

- **Method**: `ListBusiness`
- **HTTP Method**: POST
- **Path**: `/v1/businesses:list`

#### ✅ Functionality:

Lists business locations based on specified criteria.

Results are paginated and can be filtered by company ID. This method is typically used for business management and
customer-facing location directories.

#### 🎯 Use Cases:

- View all business locations for a company.
- Audit or debug business configurations.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                                         |
|--------------|------------|----------|-----------------------------------------------------|
| `pagination` | Pagination | Yes      | Page size and token                                 |
| `companyId`  | string     | Yes      | Company ID to list businesses, obfuscated ID string |

#### 📌 Return Value:

| Field Name      | Type              | Description                          |
|-----------------|-------------------|--------------------------------------|
| `nextPageToken` | string            | Token for retrieving the next page   |
| `businesses`    | Array(`Business`) | List of businesses matching criteria |

#### ⚠️ Error Code:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

## 🧪 6. Usage Examples

### Example 1: Retrieve a Specific Business

**Request**

```http
GET /v1/businesses/bus_001
```

**Response**

```json
{
  "id": "bus_001",
  "name": "Moego Pet Care Downtown",
  "avatar": "https://example.com/images/downtown.jpg",
  "phone": "+12125551234",
  "businessPhone": "+12125556789",
  "email": "downtown@moegopetcare.com",
  "address": {
    "streetAddress": "123 Main Street",
    "city": "New York",
    "state": "NY",
    "postalCode": "10001",
    "country": "US"
  },
  "website": "https://moegopetcare.com/downtown",
  "facebook": "https://facebook.com/moegodowntown",
  "instagram": "https://instagram.com/moegodowntown",
  "yelp": "https://yelp.com/biz/moego-pet-care-downtown",
  "google": "https://google.com/maps/place/Moego+Pet+Care",
  "companyId": "cmp_001"
}
```

---

### Example 2: List Businesses with Pagination

**Request**

```http
POST /v1/businesses:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "companyId": "cmp_001"
}
```

**Response**

```json
{
  "nextPageToken": "CBAQAA==",
  "businesses": [
    {
      "id": "bus_001",
      "name": "Moego Pet Care Downtown",
      "address": {
        "streetAddress": "123 Main Street",
        "city": "New York",
        "state": "NY",
        "postalCode": "10001",
        "country": "US"
      },
      "companyId": "cmp_001"
    },
    {
      "id": "bus_002",
      "name": "Moego Pet Care Brooklyn",
      "address": {
        "streetAddress": "456 Atlantic Ave",
        "city": "Brooklyn",
        "state": "NY",
        "postalCode": "11201",
        "country": "US"
      },
      "companyId": "cmp_001"
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                                | Answer                                                                  |
|---------------------------------------------------------|-------------------------------------------------------------------------|
| How can I verify if a business exists?                  | Use `GetBusiness` to check if the business ID returns a valid response. |
| Can I list all businesses at once?                      | No. Results are paginated. Use `pageToken` to fetch next set.           |
| Why does listing businesses return "permission denied"? | Ensure you have the correct access rights for this operation.           |
| How to handle large result sets efficiently?            | Use pagination via `pageSize` and `pageToken`.                          |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                      |
|---------------------|--------------------------------------------------|
| `NOT_FOUND`         | The requested business ID does not exist.        |
| `PERMISSION_DENIED` | Caller lacks access rights to perform operation. |
| `INVALID_ARGUMENT`  | Invalid request parameters provided.             |
| `INTERNAL`          | Internal server error occurred.                  |

---

## 📎 10. Related File References

- [business.proto](../moego/business/business/v1/business.proto)
- [business_service.proto](../moego/business/business/v1/business_service.proto)
- [address.proto](../moego/common/v1/address.proto)
- [pagination.proto](../moego/common/v1/pagination.proto)