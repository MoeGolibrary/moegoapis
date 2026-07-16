# Van Staff Record API Documentation (`moego.business.van_record.v1`)

## 1. Functional Overview

The **Van Staff Record** module provides APIs for retrieving historical day-level snapshots of van-to-staff assignments.

This interface is used to:

- Look up which staff were assigned to specific vans on a specific date.
- Preserve raw product identifiers for downstream mapping.
- Support historical attribution, investigation, and reporting flows at the van level.

---

## 2. Design Goals

- **Historical Snapshot**: Return the van-to-staff state for a specific day.
- **Van-Scoped Query**: Scope requests by `van_ids` instead of querying an entire business blindly.
- **List Contract Consistency**: Keep the public API aligned with standard list-style pagination semantics.
- **Readable Date Model**: Use `yyyy-MM-dd` instead of exposing internal timestamp storage details.
- **Bridge Friendly**: Include `rawId` values so downstream systems can map to internal entities.

---

## 3. Core Models

### 3.1 VanStaffRecord

Represents the assignment snapshot for one van on one date.

| Field Name | Type | Description |
|------------|------|-------------|
| `vanId` | string | Van ID, obfuscated string |
| `vanRawId` | int64 | Original numeric van ID in MoeGo product |
| `businessId` | string | Owning business location ID |
| `date` | string | Snapshot date in `yyyy-MM-dd` format |
| `assignedStaffs` | Array(`AssignedStaff`) | Staff assigned to the van on that date |

### 3.2 AssignedStaff

Reuses the `AssignedStaff` model from `moego.business.van.v1`.

| Field Name | Type | Description |
|------------|------|-------------|
| `id` | string | Staff ID, obfuscated string |
| `rawId` | int64 | Original numeric staff ID in MoeGo product |
| `name` | string | Staff display name |

---

## 4. Typical Usage Flow

### Scenario: Attribute historical data by van

1. Call `ListVanStaffRecords` with the target `van_ids` and date.
2. For each returned record, identify the staff assigned to that van on that day.
3. Use those staff IDs to query downstream appointments, services, or revenue data.
4. Aggregate results by `vanId` or `vanRawId`.

---

## 5. API Interface Description

### 5.1 List Van Staff Records (`ListVanStaffRecords`)

- **Method**: `ListVanStaffRecords`
- **HTTP Method**: POST
- **Path**: `/v1/van_staff_records:list`

#### Functionality

Lists historical van-to-staff assignment snapshots for the requested vans on a specific date.

#### Use Cases

- Analyze which staff contributed to specific vans on a given day.
- Build van-level historical drill-down flows.
- Debug historical assignment changes.

#### Request Parameters

| Field Name | Type | Required | Description |
|------------|------|----------|-------------|
| `pagination` | Pagination | Yes | Page size and page token |
| `vanIds` | Array(string) | Yes | Target van IDs, obfuscated strings |
| `date` | string | Yes | Snapshot date in `yyyy-MM-dd` format |

#### Return Value

| Field Name | Type | Description |
|------------|------|-------------|
| `nextPageToken` | string | Token for retrieving the next page of results |
| `records` | Array(`VanStaffRecord`) | Historical records matching the request |

#### Response Behavior

- Results are ordered by `van_id` ascending after pagination is applied.

#### Error Codes

| Error Code | Description |
|------------|-------------|
| `INVALID_ARGUMENT` | `van_ids`, `date`, or `pagination` are invalid |
| `PERMISSION_DENIED` | Caller lacks access to one or more requested vans |

---

## 6. Usage Example

### Example: List historical records for multiple vans on one day

**Request**

```http
POST /v1/van_staff_records:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 50,
    "pageToken": "1"
  },
  "vanIds": ["van_001", "van_002"],
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
    },
    {
      "vanId": "van_002",
      "vanRawId": 102,
      "businessId": "biz_001",
      "date": "2026-06-17",
      "assignedStaffs": [
        {
          "id": "stf_003",
          "rawId": 1003,
          "name": "Alice Lee"
        }
      ]
    }
  ]
}
```
