# 🛠️ Setting Customer API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The customer settings provide configuration options related to customer management, including customer tags, referral
sources, lead lifecycle stages, and action statuses. These settings help standardize customer data entry and ensure
consistent customer service delivery across all business locations.

---

## 🎯 2. Design Goals

- **Centralized Configuration**: Provides a unified way to manage customer-related settings like tags and referral
  sources.
- **Rich Data Model**: Supports complex relationships like customer categorization and lead management.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Categorizing customers for marketing and reporting.
- Tracking lead progression through sales pipelines.
- Managing customer acquisition sources.
- Standardizing customer service across multiple locations.

---

## 🧩 3. Core Concepts

### 1. CustomerTag

Represents a label that can be applied to customers for categorization and filtering purposes.

| Field Name        | Type      | Description                                       |
|-------------------|-----------|---------------------------------------------------|
| `id`              | string    | Unique identifier                                 |
| `name`            | string    | Display name of the tag                           |
| `lastUpdatedBy`   | string    | ID of the staff member who last modified this tag |
| `lastUpdatedTime` | Timestamp | When this tag was last modified                   |

### 2. ReferralSource

Represents the source or channel through which a customer was acquired.

| Field Name | Type   | Description                |
|------------|--------|----------------------------|
| `id`       | string | Unique identifier          |
| `name`     | string | Display name of the source |

### 3. LifeCycle

Represents a stage in the lead management process.

| Field Name | Type   | Description       |
|------------|--------|-------------------|
| `id`       | string | Unique identifier |
| `name`     | string | Display name      |

### 4. ActionStatus

Represents the status of an action taken on a lead.

| Field Name | Type   | Description       |
|------------|--------|-------------------|
| `id`       | string | Unique identifier |
| `name`     | string | Display name      |
| `color`    | string | Display color     |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Customer Setting API

Here is a typical integration flow:

1. **List Customer Tags**
    - Retrieve all available customer tags for categorization.

2. **List Referral Sources**
    - Retrieve all available sources for tracking customer acquisition.

3. **List Lead Lifecycle Stages**
    - Retrieve all lifecycle stages for lead management.

4. **List Lead Action Statuses**
    - Retrieve all action statuses for lead follow-up tracking.

---

## 📦 5. API Interface Descriptions

### 1. List Customer Tags (`ListCustomerTags`)

- **Method**: `ListCustomerTags`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/customer/tags:list`

#### ✅ Functionality:

Lists all available customer tags for a company.

Customer tags help categorize clients and can be used for marketing, reporting, and service customization. Tags can
indicate preferences, loyalty status, or special handling requirements.

#### 🎯 Use Cases:

- Retrieve all active customer tags for a company
- Display customer tags in UI for categorization

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                                 |
|-------------|--------|----------|---------------------------------------------|
| `companyId` | string | Yes      | ID of the company to list customer tags for |

#### 📌 Return Value:

| Field Name | Type                 | Description                           |
|------------|----------------------|---------------------------------------|
| `tags`     | Array(`CustomerTag`) | List of customer tags for the company |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 2. List Customer Referral Sources (`ListCustomerReferralSources`)

- **Method**: `ListCustomerReferralSources`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/customer/referral_sources:list`

#### ✅ Functionality:

Lists all available customer referral sources for a company.

Referral sources help track the origin of new customers and analyze marketing effectiveness. These sources can be used
for lead generation, marketing campaigns, and customer analytics.

#### 🎯 Use Cases:

- Track the origin of new customers
- Analyze marketing effectiveness

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                           |
|-------------|--------|----------|---------------------------------------|
| `companyId` | string | Yes      | ID of the company to list sources for |

#### 📌 Return Value:

| Field Name        | Type                    | Description              |
|-------------------|-------------------------|--------------------------|
| `referralSources` | Array(`ReferralSource`) | List of referral sources |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 3. List Leads Life Cycles (`ListLeadsLifeCycles`)

- **Method**: `ListLeadsLifeCycles`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/leads/life_cycles`

#### ✅ Functionality:

Returns a list of lead life cycles.

#### 🎯 Use Cases:

- Understand the stages of a lead's journey
- Standardize lead management processes

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                          |
|-------------|--------|----------|--------------------------------------|
| `companyId` | string | Yes      | ID of the company to list cycles for |

#### 📌 Return Value:

| Field Name   | Type               | Description              |
|--------------|--------------------|--------------------------|
| `lifeCycles` | Array(`LifeCycle`) | List of lead life cycles |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |
| `INVALID_ARGUMENT`  | Malformed request |

---

### 4. List Leads Action Status (`ListLeadsActionStatus`)

- **Method**: `ListLeadsActionStatus`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/leads/action_status`

#### ✅ Functionality:

Returns a list of lead action statuses.

#### 🎯 Use Cases:

- Track the status of actions taken on leads
- Standardize lead follow-up procedures

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                            |
|-------------|--------|----------|----------------------------------------|
| `companyId` | string | Yes      | ID of the company to list statuses for |

#### 📌 Return Value:

| Field Name       | Type                  | Description                  |
|------------------|-----------------------|------------------------------|
| `actionStatuses` | Array(`ActionStatus`) | List of lead action statuses |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |
| `INVALID_ARGUMENT`  | Malformed request |

---

## 🧪 6. Usage Examples

### Example 1: List Customer Tags

Request Body:

```json
{
  "companyId": "cmp_001"
}
```

Response Body:

```json
{
  "tags": [
    {
      "id": "tag_001",
      "name": "VIP Customer",
      "lastUpdatedBy": "staff_001",
      "lastUpdatedTime": "2023-01-01T10:00:00Z"
    },
    {
      "id": "tag_002",
      "name": "First-time Customer",
      "lastUpdatedBy": "staff_002",
      "lastUpdatedTime": "2023-01-02T10:00:00Z"
    }
  ]
}
```

### Example 2: List Referral Sources

Request Body:

```json
{
  "companyId": "cmp_001"
}
```

Response Body:

```json
{
  "referralSources": [
    {
      "id": "ref_001",
      "name": "Google Search"
    },
    {
      "id": "ref_002",
      "name": "Social Media"
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                                   | Answer                                                                                     |
|------------------------------------------------------------|--------------------------------------------------------------------------------------------|
| How can I categorize customers effectively?                | Use `ListCustomerTags` to retrieve tags and apply them to customers for better management. |
| How can I track customer acquisition sources?              | Use `ListCustomerReferralSources` to get sources and associate them with new customers.    |
| How do I manage lead progression?                          | Use `ListLeadsLifeCycles` and `ListLeadsActionStatus` to track and manage leads.           |
| What should I do if a request returns "permission denied"? | Verify that your API key has the necessary permissions to access customer settings.        |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| `PERMISSION_DENIED` | Current user has no access rights to perform the operation.                 |
| `INVALID_ARGUMENT`  | Invalid request parameters (e.g., missing required fields, invalid format). |
| `INTERNAL`          | Internal server error occurred while processing the request.                |

---

## 📎 10. Related File References

- [customer.proto](../moego/business/setting/v1/customer.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)