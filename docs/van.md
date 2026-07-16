# Van API Documentation (`moego.business.van.v1`)

## 1. Functional Overview

The **Van** module provides APIs for retrieving the current mobile van configuration under a company.

This interface enables:

- Listing vans configured for one or more businesses.
- Inspecting the current staff assignments on each van.
- Mapping MoeGo internal van and staff identifiers through `rawId` fields.

This is useful for usage-based billing, operational audits, and integrations that need to understand the current van
setup.

---

## 2. Design Goals

- **Current State Snapshot**: Return the latest van and staff assignment configuration.
- **Integration Friendly**: Preserve `rawId` fields so external systems can map to product identifiers.
- **Scoped Access**: Query at company scope and optionally narrow to specific businesses.
- **Simple Consumption**: Return van and assigned staff in a single response model.

---

## 3. Core Concepts

### 1. Van

Represents a mobile service vehicle configured under a business.

| Field Name        | Type                 | Description |
|-------------------|----------------------|-------------|
| `id`              | string               | Van ID, obfuscated string |
| `rawId`           | int64                | Original numeric van ID in MoeGo product |
| `companyId`       | string               | Parent company ID |
| `businessId`      | string               | Owning business location ID |
| `displayName`     | string               | Configured van display name |
| `licensePlate`    | string               | Configured license plate |
| `assignedStaffs`  | Array(`AssignedStaff`) | Staff currently assigned to the van |

### 2. AssignedStaff

Represents a staff member currently assigned to a van.

| Field Name | Type   | Description |
|------------|--------|-------------|
| `id`       | string | Staff ID, obfuscated string |
| `rawId`    | int64  | Original numeric staff ID in MoeGo product |
| `name`     | string | Staff display name |

---

## 4. Typical Usage Flow

### Scenario: Reconcile van-based billing

1. Call `ListVans` for a company.
2. Optionally restrict the query to one or more businesses.
3. Count returned vans to support billing logic.
4. Use `assignedStaffs` to understand who is currently mapped to each van.

---

## 5. API Interface Descriptions

### 1. List Vans (`ListVans`)

- **Method**: `ListVans`
- **HTTP Method**: POST
- **Path**: `/v1/vans:list`

#### Functionality

Lists vans under a company and optionally narrows the results to specific businesses.

#### Use Cases

- Determine how many vans are currently configured.
- Audit current van-to-staff assignments.
- Export current van configuration into a third-party system.

#### Request Parameters

| Field Name    | Type                | Required | Description |
|---------------|---------------------|----------|-------------|
| `pagination`  | Pagination          | Yes      | Page size and page token |
| `companyId`   | string              | Yes      | Company ID |
| `businessIds` | Array(string)       | No       | Business IDs to narrow the result set |

#### Return Value

| Field Name      | Type            | Description |
|-----------------|-----------------|-------------|
| `nextPageToken` | string          | Token for retrieving the next page of results |
| `vans`          | Array(`Van`)    | Vans matching the request |

#### Error Codes

| Error Code          | Description |
|---------------------|-------------|
| `INVALID_ARGUMENT`  | Request pagination or IDs are invalid |
| `PERMISSION_DENIED` | Caller lacks access to the company or requested businesses |

---

## 6. Usage Example

### Example: List vans for one business

**Request**

```http
POST /v1/vans:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "companyId": "cop_001",
  "businessIds": ["biz_001"]
}
```

**Response**

```json
{
  "nextPageToken": "",
  "vans": [
    {
      "id": "van_001",
      "rawId": 101,
      "companyId": "cop_001",
      "businessId": "biz_001",
      "displayName": "Van A",
      "licensePlate": "ABC-123",
      "assignedStaffs": [
        {
          "id": "stf_001",
          "rawId": 1001,
          "name": "Jane Doe"
        }
      ]
    }
  ]
}
```
