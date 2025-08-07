# 👥 Staff API Documentation (`moego.business.staff.v1`)

## 📌 1. Functional Overview

The **Staff** module provides APIs to manage staff members who provide services in your business. Each staff member can
be assigned to appointments, manage customer relationships, and perform various administrative tasks based on their role
and permissions.

This interface enables:

- Retrieving detailed information about a specific staff member.
- Listing all staff members associated with a company.
- Managing roles, permissions, and work locations for each staff member.

Useful for scenarios such as staff management, scheduling, access control, and integration with third-party systems.

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all staff members.
- **Role-Based Access Control**: Supports role-based permissions and capabilities.
- **Secure and Reliable**: Ensures data integrity and access control.
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks.

Applicable to scenarios like staff assignment, permission configuration, and system administration.

---

## 🧩 3. Core Concepts

### 1. Staff

Represents an employee or contractor who provides services in your business.

| Field Name           | Type          | Description                                                 |
|----------------------|---------------|-------------------------------------------------------------|
| `id`                 | string        | Unique identifier (e.g., `"stf_001"`).                      |
| `firstName`          | string        | First name of the staff member                              |
| `lastName`           | string        | Last name of the staff member                               |
| `avatar`             | string        | Photo URL of the staff member                               |
| `phone`              | string        | Contact phone number (E.164 format)                         |
| `email`              | string        | Email address                                               |
| `role`               | Role          | Assigned role and permissions                               |
| `hiredTime`          | Timestamp     | When the staff member was hired                             |
| `workingBusinessIds` | Array(string) | List of business locations where this staff works           |
| `mobileVanId`        | string        | ID of the mobile van assigned to this staff (if applicable) |
| `companyId`          | string        | ID of the parent company                                    |

---

### 2. Role

Defines a set of permissions and capabilities assigned to staff members.

| Field Name    | Type              | Description                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | string            | Unique identifier (e.g., `"rol_001"`)      |
| `name`        | string            | Display name of the role (e.g., "Groomer") |
| `permissions` | Array(Permission) | List of permissions granted to this role   |

---

### 3. Permission

Represents a specific action that can be performed in the system.

#### Enum: Permission

- `PERMISSION_UNSPECIFIED`

> ⚠️ Note: This enum is currently empty and will be expanded with actual permissions in future versions.

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Staff API

Here is a typical integration flow:

1. **Get Staff**
    - Retrieve detailed information about a specific staff member.
    - Used during authentication, authorization, or debugging.

2. **List Staffs**
    - View all staff members associated with a company.
    - Useful for staff management and scheduling.

3. **Monitoring & Maintenance**
    - Regularly retrieve staff data to monitor changes.
    - Update staff records as needed.

---

## 🛠️ 5. API Interface Descriptions

### 1. Get Staff (`GetStaff`)

- **Method**: `GetStaff`
- **HTTP Method**: GET
- **Path**: `/v1/staffs/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific staff member.

#### 🎯 Use Cases:

- Check current staff data.
- Verify staff details during debugging.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description          |
|------------|--------|----------|----------------------|
| `id`       | string | Yes      | Staff ID to retrieve |

#### 📌 Return Value:

| Field Name | Type    | Description                |
|------------|---------|----------------------------|
| `staff`    | `Staff` | The retrieved staff object |

#### ⚠️ Error Codes:

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Specified staff ID does not exist |
| `PERMISSION_DENIED` | Permission denied                 |

---

### 2. List Staffs (`ListStaffs`)

- **Method**: `ListStaffs`
- **HTTP Method**: POST
- **Path**: `/v1/staffs:list`

#### ✅ Functionality:

Lists staff members based on specified criteria.

Results are paginated and filtered by company ID. This method is typically used for staff management, scheduling, and
reporting purposes.

#### 🎯 Use Cases:

- View all staff members for a company.
- Audit or debug staff configurations.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description               |
|--------------|------------|----------|---------------------------|
| `pagination` | Pagination | Yes      | Page size and token       |
| `companyId`  | string     | Yes      | Company ID to list staffs |

#### 📌 Return Value:

| Field Name      | Type           | Description                             |
|-----------------|----------------|-----------------------------------------|
| `nextPageToken` | string         | Token for retrieving the next page      |
| `staffs`        | Array(`Staff`) | List of staff members matching criteria |

#### ⚠️ Error Code:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

## 🧪 6. Usage Examples

### Example 1: Retrieve a Specific Staff Member

**Request**

```http
GET /v1/staffs/stf_001
```

**Response**

```json
{
  "id": "stf_001",
  "firstName": "John",
  "lastName": "Doe",
  "avatar": "https://example.com/images/john.jpg",
  "phone": "+12125551234",
  "email": "john.doe@example.com",
  "role": {
    "id": "rol_001",
    "name": "Administrator",
    "permissions": []
  },
  "hiredTime": "2020-01-15T08:00:00Z",
  "workingBusinessIds": [
    "bus_001",
    "bus_002"
  ],
  "mobileVanId": "",
  "companyId": "cmp_001"
}
```

---

### Example 2: List Staff Members with Pagination

**Request**

```http
POST /v1/staffs:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001"
}
```

**Response**

```json
{
  "nextPageToken": "CBAQAA==",
  "staffs": [
    {
      "id": "stf_001",
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "role": {
        "id": "rol_001",
        "name": "Administrator",
        "permissions": []
      },
      "hiredTime": "2020-01-15T08:00:00Z",
      "workingBusinessIds": [
        "bus_001",
        "bus_002"
      ],
      "companyId": "cmp_001"
    },
    {
      "id": "stf_002",
      "firstName": "Jane",
      "lastName": "Smith",
      "email": "jane.smith@example.com",
      "role": {
        "id": "rol_002",
        "name": "Groomer",
        "permissions": []
      },
      "hiredTime": "2021-03-10T08:00:00Z",
      "workingBusinessIds": [
        "bus_001"
      ],
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

| Question                                           | Answer                                                            |
|----------------------------------------------------|-------------------------------------------------------------------|
| How can I verify if a staff member exists?         | Use `GetStaff` to check if the staff ID returns a valid response. |
| Can I list all staff members at once?              | No. Results are paginated. Use `pageToken` to fetch next set.     |
| Why does listing staff return "permission denied"? | Ensure you have the correct access rights for this operation.     |
| How to handle large result sets efficiently?       | Use pagination via `pageSize` and `pageToken`.                    |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                      |
|---------------------|--------------------------------------------------|
| `NOT_FOUND`         | The requested staff ID does not exist.           |
| `PERMISSION_DENIED` | Caller lacks access rights to perform operation. |
| `INVALID_ARGUMENT`  | Invalid request parameters provided.             |
| `INTERNAL`          | Internal server error occurred.                  |

---

## 📎 10. Related File References

- [staff.proto](../moego/business/staff/v1/staff.proto)
- [staff_service.proto](../moego/business/staff/v1/staff_service.proto)
- [pagination.proto](../moego/common/v1/pagination.proto)