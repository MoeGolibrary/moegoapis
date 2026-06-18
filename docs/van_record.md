# Van Staff Record API Documentation (`moego.business.van_record.v1`)

## 1. Functional Overview

The **Van Staff Record** module provides APIs for retrieving historical day-level snapshots of van-to-staff assignments.

This interface enables:

- Looking up which staff were assigned to each van on a specific date.
- Preserving raw product identifiers for both vans and staff.
- Building drill-down workflows from van to staff to appointments and services.

This is useful for revenue attribution, operational investigations, and historical reporting.

---

## 2. Design Goals

- **Historical Snapshot**: Return the van-to-staff state for a specific day.
- **Business Scoped**: Query records within a single business.
- **Readable Date Model**: Use `yyyy-MM-dd` instead of exposing internal timestamp storage details.
- **Bridge Friendly**: Include `rawId` values to help external systems map downstream data.

---

## 3. Core Concepts

### 1. VanStaffRecord

Represents the assignment snapshot for one van on one date.

| Field Name        | Type                   | Description |
|-------------------|------------------------|-------------|
| `vanId`           | string                 | Van ID, obfuscated string |
| `vanRawId`        | int64                  | Original numeric van ID in MoeGo product |
| `businessId`      | string                 | Owning business location ID |
| `date`            | string                 | Snapshot date in `yyyy-MM-dd` format |
| `assignedStaffs`  | Array(`AssignedStaff`) | Staff assigned to the van on that date |

### 2. AssignedStaff

Reuses the `AssignedStaff` model from `moego.business.van.v1`.

| Field Name | Type   | Description |
|------------|--------|-------------|
| `id`       | string | Staff ID, obfuscated string |
| `rawId`    | int64  | Original numeric staff ID in MoeGo product |
| `name`     | string | Staff display name |

---

## 4. Typical Usage Flow

### Scenario: Attribute revenue by van

1. Call `ListVanStaffRecords` for a business and a target date.
2. For each returned record, identify the staff assigned to the van that day.
3. Use those staff IDs to query appointments and downstream service revenue.
4. Aggregate results by `vanId` or `vanRawId`.

---

## 5. API Interface Descriptions

### 1. List Van Staff Records (`ListVanStaffRecords`)

- **Method**: `ListVanStaffRecords`
- **HTTP Method**: POST
- **Path**: `/v1/van_staff_records:list`

#### Functionality

Lists the historical van-to-staff assignment snapshot for a specific business on a specific date.

#### Use Cases

- Analyze which staff contributed to a van on a given day.
- Build van-level historical revenue drill-down.
- Debug assignment changes over time.

#### Request Parameters

| Field Name   | Type         | Required | Description |
|--------------|--------------|----------|-------------|
| `pagination` | Pagination   | Yes      | Page size and page token |
| `businessId` | string       | Yes      | Business ID |
| `date`       | string       | Yes      | Snapshot date in `yyyy-MM-dd` format |

#### Return Value

| Field Name      | Type                   | Description |
|-----------------|------------------------|-------------|
| `nextPageToken` | string                 | Token for retrieving the next page of results |
| `records`       | Array(`VanStaffRecord`) | Historical records matching the request |

#### Error Codes

| Error Code          | Description |
|---------------------|-------------|
| `INVALID_ARGUMENT`  | Request date, pagination, or IDs are invalid |
| `PERMISSION_DENIED` | Caller lacks access to the business |

---

## 6. Usage Example

### Example: List van staff records for one day

**Request**

```http
POST /v1/van_staff_records:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 50,
    "pageToken": "1"
  },
  "businessId": "biz_001",
  "date": "2026-06-17"
}
```

**Response**

```json
{
  "nextPageToken": "",
  "records": [
    {
      "vanId": "van_001",
      "vanRawId": 101,
      "businessId": "biz_001",
      "date": "2026-06-17",
      "assignedStaffs": [
        {
          "id": "stf_001",
          "rawId": 1001,
          "name": "Jane Doe"
        },
        {
          "id": "stf_002",
          "rawId": 1002,
          "name": "John Smith"
        }
      ]
    }
  ]
}
```
