# Lead API Documentation (`moego.business.customer.v1`)

## 📌 1. Functional Overview

`Lead` represents a potential customer in the sales pipeline. This interface enables:

- Creating, updating, and retrieving lead information
- Managing lead contact details, pets, and acquisition source
- Listing leads with optional filtering
- Converting leads to customers

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all aspects of lead data
- **Rich Data Model**: Supports complex lead information including pets, lifecycle stages, and action status
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

Applicable to scenarios such as lead management, sales pipeline tracking, and marketing campaigns.

---

## 🧩 3. Core Concepts

### 1. Lead

Represents a potential customer in the sales pipeline

| Field Name              | Type                                                    | Description                                               |
|-------------------------|---------------------------------------------------------|-----------------------------------------------------------|
| `id`                    | string                                                  | Unique identifier                                         |
| `first_name`            | string                                                  | Lead's first name                                         |
| `last_name`             | string                                                  | Lead's last name                                          |
| `avatar`                | string                                                  | URL to the lead's profile picture                         |
| `phone`                 | string                                                  | Lead's phone number. Must be in E.164 format              |
| `email`                 | string                                                  | Lead's email address. Must be valid                       |
| `address`               | Address                                                 | Lead's address                                            |
| `pets`                  | Array(Pet)                                              | List of pets belonging to the lead                        |
| `preferred_business_id` | string                                                  | ID of the lead's preferred business location              |
| `allocate_staff_id`     | string                                                  | ID of the staff member assigned to the lead               |
| `life_cycle`            | [LifeCycle](./setting_service.md#6-lifecycle)           | Lead's current lifecycle stage                            |
| `action_status`         | [ActionStatus](./setting_service.md#7-actionstatus)     | Lead's action status                                      |
| `referral_source`       | [ReferralSource](./setting_service.md#5-referralsource) | The source or channel through which the lead was acquired |
| `created_time`          | Timestamp                                               | Creation timestamp                                        |
| `last_updated_time`     | Timestamp                                               | Last modification timestamp                               |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Lead API

Here is a typical integration flow:

1. **Create Lead**
    - Specify required details like name, phone number, etc.
    - Optionally add pet information, address, etc.

2. **Update Lead**
    - Modify lead details
    - Update lifecycle stage or action status

3. **Retrieve Lead**
    - Get full details of an existing lead

4. **List Leads**
    - View all leads in the company
    - Filter by criteria if needed

5. **Convert Lead**
    - Convert a lead to a customer

6. **Monitoring & Maintenance**
    - Regularly retrieve lead data to monitor changes
    - Update lead records as needed

---

## 📦 5. API Interface Descriptions

### 1. Create Lead (`CreateLead`)

- **Method**: `CreateLead`
- **HTTP Method**: POST
- **Path**: `/v1/leads`

#### ✅ Functionality:

Creates a new lead

#### 🎯 Use Cases:

- Users want to add new leads to the system
- Third-party systems sync lead data

#### 🔧 Request Parameters:

| Field Name | Type | Required | Description                |
|------------|------|----------|----------------------------|
| `lead`     | Lead | Yes      | Lead information to create |

#### 📌 Return Value:

Returns the created `Lead` object

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields are missing or invalid
- `PERMISSION_DENIED`: Permission denied

---

### 2. Get Lead (`GetLead`)

- **Method**: `GetLead`
- **HTTP Method**: GET
- **Path**: `/v1/leads/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific lead

#### 🎯 Use Cases:

- Check current lead data
- Verify lead details during debugging

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description         |
|------------|--------|----------|---------------------|
| `id`       | string | Yes      | Lead ID to retrieve |

#### 📌 Return Value:

Returns the complete `Lead` object

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified lead ID does not exist
- `PERMISSION_DENIED`: Permission denied

---

### 3. Update Lead (`UpdateLead`)

- **Method**: `UpdateLead`
- **HTTP Method**: PUT
- **Path**: `/v1/leads/{id}`

#### ✅ Functionality:

Updates an existing lead's information

#### 🎯 Use Cases:

- Change lead details
- Update lifecycle stage or action status

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description              |
|------------|--------|----------|--------------------------|
| `id`       | string | Yes      | Lead ID to update        |
| `lead`     | Lead   | Yes      | Updated lead information |

#### 📌 Return Value:

Returns the updated `Lead` object

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified lead ID does not exist
- `INVALID_ARGUMENT`: Invalid request parameters
- `PERMISSION_DENIED`: Permission denied

---

### 4. List Leads (`ListLeads`)

- **Method**: `ListLeads`
- **HTTP Method**: GET
- **Path**: `/v1/leads:list`

#### ✅ Functionality:

Lists leads with pagination and optional filters

#### 🎯 Use Cases:

- View all leads in the company
- Filter leads by criteria

#### 🔧 Request Parameters:

| Field Name                 | Type       | Required | Description                            |
|----------------------------|------------|----------|----------------------------------------|
| `pagination`               | Pagination | Yes      | Pagination info: page_size, page_token |
| `company_id`               | string     | Yes      | ID of the company to list leads for    |
| `filter.life_cycle_id`     | string     | No       | ID of the lead's life cycle stage      |
| `filter.action_status_id`  | string     | No       | ID of the lead's action status         |
| `filter.main_phone_number` | string     | No       | Main phone number of the lead          |

> **Note**: The `pagination` field is used for pagination.
> The `page_size` field specifies the number of results to return per page. Maximum value is 500.
> The `page_token` field is used to retrieve the next page of results.

#### 📌 Return Value:

Returns paginated results and lead list

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied

---

### 5. Convert Lead (`ConvertLead`)

- **Method**: `ConvertLead`
- **HTTP Method**: POST
- **Path**: `/v1/leads:convert`

#### ✅ Functionality:

Converts a lead to a customer

#### 🎯 Use Cases:

- Move a lead from the sales pipeline to become a customer

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description        |
|------------|--------|----------|--------------------|
| `id`       | string | Yes      | Lead ID to convert |

#### 📌 Return Value:

Returns the newly created `Customer` object

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified lead ID does not exist
- `INVALID_ARGUMENT`: Invalid request parameters
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: Create Lead

```json
{
  "lead": {
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+12125551234",
    "email": "john.doe@example.com",
    "address": {
      "street": "123 Main St",
      "city": "New York",
      "state": "NY",
      "zip": "10001"
    },
    "pets": [
      {
        "name": "Buddy",
        "type": "DOG",
        "breed": "Labrador Retriever"
      }
    ]
  }
}
```

### Example 2: Update Lead

```json
{
  "id": "lcus_001",
  "lead": {
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+12125551234",
    "email": "john.doe@example.com",
    "life_cycle": {
      "id": "lc_001",
      "name": "Qualified"
    },
    "action_status": {
      "id": "as_001",
      "name": "Contacted",
      "color": "#00FF00"
    }
  }
}
```

### Example 3: List Leads

```json
{
  "company_id": "cmp_001",
  "pagination": {
    "page_size": 20
  },
  "filter": {
    "life_cycle_id": "lc_001"
  }
}
```

### Example 4: Convert Lead

```json
{
  "id": "lcus_001"
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                             | Answer                                                               |
|--------------------------------------|----------------------------------------------------------------------|
| How to verify if a lead exists?      | Use `GetLead` to check if the lead ID returns a valid response       |
| Can I create multiple leads at once? | Currently only single lead creation is supported                     |
| How to filter leads effectively?     | Use `ListLeads` with appropriate filter parameters                   |
| What happens when converting a lead? | The lead is converted to a customer and removed from the lead system |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Lead ID does not exist            |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [lead_service.proto](../moego/business/customer/v1/lead_service.proto)
- [lead.proto](../moego/business/customer/v1/lead.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)