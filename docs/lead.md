# Lead API Documentation (`moego.business.customer.v1`)

## 📌 1. Functional Overview

`Lead` represents a potential customer in the sales pipeline. This interface enables:

- Creating, updating, and retrieving lead information
- Managing lead contact details, pets, and acquisition source
- Listing leads with optional filtering
- Promoting leads to customers

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

| Field Name            | Type                                                     | Description                                               |
|-----------------------|----------------------------------------------------------|-----------------------------------------------------------|
| `id`                  | string                                                   | Unique identifier                                         |
| `companyId`           | string                                                   | ID of the company the lead belongs to                     |
| `firstName`           | string                                                   | Lead's first name                                         |
| `lastName`            | string                                                   | Lead's last name                                          |
| `avatar`              | string                                                   | URL to the lead's profile picture                         |
| `phone`               | string                                                   | Lead's phone number. Must be in E.164 format              |
| `email`               | string                                                   | Lead's email address. Must be valid                       |
| `address`             | Address                                                  | Lead's address                                            |
| `pets`                | Array([Pet](./pet.md#1-pet))                             | List of pets belonging to the lead                        |
| `preferredBusinessId` | string                                                   | ID of the lead's preferred business location              |
| `allocateStaffId`     | string                                                   | ID of the staff member assigned to the lead               |
| `lifeCycle`           | [LifeCycle](./setting_customer.md#3-lifecycle)           | Lead's current lifecycle stage                            |
| `actionStatus`        | [ActionStatus](./setting_customer.md#4-actionstatus)     | Lead's action status                                      |
| `referralSource`      | [ReferralSource](./setting_customer.md#2-referralsource) | The source or channel through which the lead was acquired |
| `createdTime`         | Timestamp                                                | Creation timestamp                                        |
| `lastUpdatedTime`     | Timestamp                                                | Last modification timestamp                               |
| `complianceConfig`    | [CustomerComplianceConfig](./customer.md#5-customercomplianceconfig) | Lead's compliance configuration for communication channels. See [Customer ComplianceConfig](./customer.md#5-customercomplianceconfig) for details |
| `customFields`        | Map<string, [CustomField.Value](./setting_custom_field.md#5-customfieldvalue)> | Custom field values. Key is the custom field code (e.g., `field_123`), value is the field value based on the field type. See [Custom Field Documentation](./setting_custom_field.md) for details |

> **Note**: Lead uses the same compliance configuration structure as Customer. For detailed information about compliance channels and configuration options, please refer to the [Customer API Documentation](./customer.md#5-customercomplianceconfig).

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

5. **Promote Lead**
    - Promote a lead to a customer

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

| Field Name         | Type                                  | Required | Description                                               |
|--------------------|---------------------------------------|----------|-----------------------------------------------------------|
| `lead`             | Lead                                  | Yes      | Lead information to create                                |
| `complianceConfig` | CustomerComplianceConfigUpdateDef     | No       | Lead's compliance configuration for communication channels. See [Customer ComplianceConfig](./customer.md#5-customercomplianceconfig) |

**Lead Object Fields:**

In addition to standard fields (firstName, lastName, phone, email, etc.), the `lead` object supports:

| Field Name     | Type                                  | Required | Description                                               |
|----------------|---------------------------------------|----------|-----------------------------------------------------------|
| `customFields` | Map<string, CustomField.Value>        | No       | Custom field values. Key is the custom field code (from [ListCustomFields](./setting_custom_field.md)), value matches the field type. See [Custom Field Documentation](./setting_custom_field.md) for details |

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

| Field Name         | Type                                  | Required | Description                                              |
|--------------------|---------------------------------------|----------|----------------------------------------------------------|
| `id`               | string                                | Yes      | Lead ID to update                                        |
| `lead`             | Lead                                  | Yes      | Updated lead information                                 |
| `complianceConfig` | CustomerComplianceConfigUpdateDef     | No       | Lead's compliance configuration updates. See [Customer ComplianceConfig](./customer.md#5-customercomplianceconfig) |

#### 📌 Return Value:

Returns the updated `Lead` object

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified lead ID does not exist
- `INVALID_ARGUMENT`: Invalid request parameters
- `PERMISSION_DENIED`: Permission denied

---

### 4. List Leads (`ListLeads`)

- **Method**: `ListLeads`
- **HTTP Method**: POST
- **Path**: `/v1/leads:list`

#### ✅ Functionality:

Lists leads with pagination and optional filters

#### 🎯 Use Cases:

- View all leads in the company
- Filter leads by criteria

#### 🔧 Request Parameters:

| Field Name               | Type       | Required | Description                          |
|--------------------------|------------|----------|--------------------------------------|
| `pagination`             | Pagination | Yes      | Pagination info: pageSize, pageToken |
| `companyId`              | string     | Yes      | ID of the company to list leads for  |
| `filter.lifeCycleId`     | string     | No       | ID of the lead's life cycle stage    |
| `filter.actionStatusId`  | string     | No       | ID of the lead's action status       |
| `filter.mainPhoneNumber` | string     | No       | Main phone number of the lead        |

> **Note**: The `pagination` field is used for pagination.
> The `pageSize` field specifies the number of results to return per page. Maximum value is 500.
> The `pageToken` field is used to retrieve the next page of results.

#### 📌 Return Value:

Returns paginated results and lead list. Each lead in the response includes all standard fields plus any custom field values that have been set.

| Field Name       | Type          | Description                                    |
|------------------|---------------|------------------------------------------------|
| `nextPageToken`  | string        | Token for retrieving the next page of results  |
| `leads`          | Array(Lead)   | List of leads with all fields including custom fields |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied

---

### 5. Promote Lead (`PromoteLead`)

- **Method**: `PromoteLead`
- **HTTP Method**: POST
- **Path**: `/v1/leads:promote`

#### ✅ Functionality:

Promotes a lead to a customer

#### 🎯 Use Cases:

- Move a lead from the sales pipeline to become a customer

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description        |
|------------|--------|----------|--------------------|
| `id`       | string | Yes      | Lead ID to promote |

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
    "firstName": "John",
    "lastName": "Doe",
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
    ],
    "customFields": {
      "field_lead_source_detail": {
        "string": "Facebook Ad Campaign"
      },
      "field_estimated_budget": {
        "money": {
          "currencyCode": "USD",
          "units": "500"
        }
      },
      "field_interested_services": {
        "stringList": {
          "values": ["Grooming", "Boarding", "Training"]
        }
      },
      "field_follow_up_date": {
        "timestampTime": "2024-02-15T10:00:00Z"
      }
    }
  },
  "complianceConfig": {
    "serviceRelatedChannels": {
      "channels": ["COMPLIANCE_CHANNEL_SMS", "COMPLIANCE_CHANNEL_EMAIL"]
    },
    "marketingCampaignsChannels": {
      "channels": ["COMPLIANCE_CHANNEL_EMAIL"]
    },
    "brandedAppEnabled": true,
    "isAgreedMarketingPolicy": true,
    "isConsented": true
  }
}
```

### Example 2: Update Lead

```json
{
  "id": "lcus_001",
  "lead": {
    "firstName": "John",
    "lastName": "Doe",
    "phone": "+12125551234",
    "email": "john.doe@example.com",
    "lifeCycle": {
      "id": "lc_001",
      "name": "Qualified"
    },
    "actionStatus": {
      "id": "as_001",
      "name": "Contacted",
      "color": "#00FF00"
    }
  },
  "complianceConfig": {
    "serviceRelatedChannels": {
      "channels": ["COMPLIANCE_CHANNEL_SMS"]
    },
    "marketingCampaignsChannels": {
      "channels": []
    },
    "isAgreedMarketingPolicy": false,
    "isConsented": false
  }
}
```

### Example 3: List Leads

Request:

```json
{
  "companyId": "cmp_001",
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "filter": {
    "lifeCycleId": "lc_001"
  }
}
```

Response:

```json
{
  "nextPageToken": "2",
  "leads": [
    {
      "id": "lcus_001",
      "firstName": "John",
      "lastName": "Doe",
      "phone": "+12125551234",
      "email": "john.doe@example.com",
      "lifeCycle": {
        "id": "lc_001",
        "name": "Qualified"
      },
      "customFields": {
        "field_lead_source_detail": {
          "string": "Facebook Ad Campaign"
        },
        "field_estimated_budget": {
          "money": {
            "currencyCode": "USD",
            "units": "500"
          }
        },
        "field_interested_services": {
          "stringList": {
            "values": ["Grooming", "Boarding"]
          }
        }
      },
      "createdTime": "2024-01-15T10:00:00Z",
      "lastUpdatedTime": "2024-01-20T14:30:00Z"
    }
  ]
}
```

### Example 4: Promote Lead

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

| Question                             | Answer                                                              |
|--------------------------------------|---------------------------------------------------------------------|
| How to verify if a lead exists?      | Use `GetLead` to check if the lead ID returns a valid response      |
| Can I create multiple leads at once? | Currently only single lead creation is supported                    |
| How to filter leads effectively?     | Use `ListLeads` with appropriate filter parameters                  |
| What happens when promoting a lead?  | The lead is promoted to a customer and removed from the lead system |
| How to control which communication channels can be used to contact a lead? | Use the `complianceConfig` field to specify allowed channels for service-related and marketing communications. Pass an empty array to clear a channel configuration. |
| How do I use custom fields with leads? | First use [ListCustomFields](./setting_custom_field.md) to get available fields and their codes, then include them in the `customFields` map when creating or updating leads. The key is the field code (e.g., `field_123`) and the value type must match the field's defined type. |
| What custom field types are supported? | Supports text, numbers, dates, booleans, selections, relations, money, time, and datetime. See [Custom Field Documentation](./setting_custom_field.md) for details. |

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