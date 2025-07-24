# Block API Documentation (`moego.business.block.v1`)

## 📌 1. Functional Overview

Block represents a time period that is blocked off in a staff member's schedule. This module provides the following
functions:

- Creating and managing schedule blocks for staff members
- Representing various types of unavailable time periods (breaks, meetings, personal time off)
- Storing descriptive information about blocked time periods
- Supporting visual differentiation of blocks through color coding
- Enabling access control through business ID scoping

Applicable to scenarios such as staff scheduling, calendar management, and availability tracking.

---

## 🎯 2. Design Goals

- **Flexible Scheduling**: Provides comprehensive tools for creating and managing staff schedule blocks
- **Visual Representation**: Supports color coding for easy identification in calendar interfaces
- **Contextual Information**: Allows adding descriptions to explain the purpose of blocked time
- **Access Control**: Implements business-level scoping to ensure proper access control
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

---

## 🧩 3. Core Concepts

### 1. Block

Represents a time period that is blocked off in a staff member's schedule. This can be used for breaks, meetings, or any
other time when the staff member is not available for appointments.

| Field Name    | Type                 | Description                                                                                                                            |
|---------------|----------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| `id`          | string               | Unique identifier for the block                                                                                                        |
| `staff_id`    | string               | ID of the staff member whose schedule is being blocked                                                                                 |
| `duration`    | google.type.Interval | The time interval during which the staff member is unavailable (uses Google's standard Interval type to represent start and end times) |
| `description` | string               | Optional description of why the time is blocked off (e.g., "Lunch break", "Team meeting")                                              |
| `color_code`  | string               | Color code used to visually distinguish this block in the calendar UI (should be a valid hex color code like "#FF0000" for red)        |
| `business_id` | string               | ID of the business where this block is created (used for access control and filtering blocks by business)                              |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Block API

Here is a typical integration flow:

1. **CreateBlock**
    - Define a new block with specific duration for a staff member
    - Set an optional description explaining why the time is blocked
    - Assign a color code for visual distinction in the calendar UI

2. **GetBlock**
    - Retrieve detailed information about a specific block using its unique ID
    - Verify block configuration and check time conflicts

3. **Schedule Management & Monitoring**
    - Regularly create and update blocks to manage staff availability
    - Use blocks to prevent scheduling conflicts
    - Monitor staff schedules across multiple business locations

---

## 📦 5. API Interface Descriptions

### 1. CreateBlock (`CreateBlock`)

- **Method**: `CreateBlock`
- **HTTP Method**: POST
- **Path**: `/v1/blocks`

#### ✅ Functionality:

Creates a new block in a staff member's schedule.

#### 🎯 Use Cases:

- Scheduling staff breaks
- Blocking time for meetings or training sessions
- Managing staff unavailability periods

#### 🔧 Request Parameters:

| Field Name    | Type     | Required | Description                                                                                                                     |
|---------------|----------|----------|---------------------------------------------------------------------------------------------------------------------------------|
| `business_id` | string   | Yes      | Company identifier for scoping the block creation                                                                               |
| `staff_id`    | string   | Yes      | ID of the staff member whose schedule will be blocked                                                                           |
| `duration`    | Interval | Yes      | Time interval during which the staff member will be unavailable (must have both start and end times specified)                  |
| `description` | string   | No       | Optional description explaining why the time is being blocked off (examples: "Lunch break", "Team meeting", "Training session") |
| `color_code`  | string   | No       | Optional color code for visual distinction in the calendar UI (should be a valid hex color code like "#FF0000" for red)         |

#### 📌 Return Value:

| Field Name    | Type                 | Description                                                                                                                            |
|---------------|----------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| `id`          | string               | Unique identifier for the block                                                                                                        |
| `staff_id`    | string               | ID of the staff member whose schedule is being blocked                                                                                 |
| `duration`    | google.type.Interval | The time interval during which the staff member is unavailable (uses Google's standard Interval type to represent start and end times) |
| `description` | string               | Optional description of why the time is blocked off (e.g., "Lunch break", "Team meeting")                                              |
| `color_code`  | string               | Color code used to visually distinguish this block in the calendar UI (should be a valid hex color code like "#FF0000" for red)        |
| `business_id` | string               | ID of the business where this block is created (used for access control and filtering blocks by business)                              |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Block configuration is invalid (e.g., conflicting parameters, invalid times)
- `PERMISSION_DENIED`: Permission denied
- `NOT_FOUND`: Business ID or staff ID is invalid

---

### 2. GetBlock (`GetBlock`)

- **Method**: `GetBlock`
- **HTTP Method**: GET
- **Path**: `/v1/blocks/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific block by its ID.

#### 🎯 Use Cases:

- Verifying block configuration before scheduling appointments
- Auditing block details and staff availability
- Checking block status and timing information

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                                |
|------------|--------|----------|--------------------------------------------|
| `id`       | string | Yes      | Unique identifier of the block to retrieve |

#### 📌 Return Value:

| Field Name    | Type                 | Description                                                                                                                            |
|---------------|----------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| `id`          | string               | Unique identifier for the block                                                                                                        |
| `staff_id`    | string               | ID of the staff member whose schedule is being blocked                                                                                 |
| `duration`    | google.type.Interval | The time interval during which the staff member is unavailable (uses Google's standard Interval type to represent start and end times) |
| `description` | string               | Optional description of why the time is blocked off (e.g., "Lunch break", "Team meeting")                                              |
| `color_code`  | string               | Color code used to visually distinguish this block in the calendar UI (should be a valid hex color code like "#FF0000" for red)        |
| `business_id` | string               | ID of the business where this block is created (used for access control and filtering blocks by business)                              |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified block ID does not exist
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: CreateBlock

```json
{
  "business_id": "biz_001",
  "staff_id": "staff_001",
  "duration": {
    "start_time": "2024-09-15T12:00:00Z",
    "end_time": "2024-09-15T13:00:00Z"
  },
  "description": "Lunch break",
  "color_code": "#FFA500"
}
```

### Example 2: GetBlock

```json
{
  "id": "block_001"
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                              | Answer                                                                                                        |
|-------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| How to verify if a block exists?                      | Use `GetBlock` to check if the block ID returns a valid response                                              |
| Can I create overlapping blocks?                      | While technically possible, it's recommended to avoid creating overlapping blocks for the same staff member   |
| How to manage staff availability effectively?         | Use `CreateBlock` to define all unavailable time periods and check them before scheduling appointments        |
| Why does creating a block return "permission denied"? | You may lack the necessary permissions for the specified business ID or staff member                          |
| How to handle time zone differences?                  | All times should be specified in UTC format with proper time zone conversion handled at the application level |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                       |
|---------------------|---------------------------------------------------|
| `NOT_FOUND`         | Block ID, business ID, or staff ID does not exist |
| `PERMISSION_DENIED` | Current user has no access rights                 |
| `INVALID_ARGUMENT`  | Invalid request parameters                        |
| `INTERNAL`          | Internal server error                             |

---

## 📎 10. Related File References

- [interval.md](https://github.com/googleapis/google-types/blob/main/google/type/interval.proto)
- [block_service.proto](../moego/business/block/v1/block_service.proto)
- [block.proto](../moego/business/block/v1/block.proto)