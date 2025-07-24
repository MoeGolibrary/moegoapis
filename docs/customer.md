# Customer API Documentation (`moego.business.customer.v1`)

## 📌 1. Functional Overview

Customer is a core entity representing clients who use your services. This interface enables:

- Creating, updating, and deleting customer profiles
- Managing customer-related information such as notes, tags, and preferences
- Listing customers with optional filtering (e.g., by last update time)
- Generating secure links for adding card-on-file information
- Retrieving detailed customer data including notes, tags, and appointment history

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all aspects of customer data
- **Rich Data Model**: Supports complex relationships like pets, addresses, notes, tags, and preferences
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

Applicable to scenarios such as customer onboarding, service history tracking, marketing campaigns, and third-party
system integration

---

## 🧩 3. Core Concepts

### 1. Customer

Represents a client who uses your services. A customer can have multiple pets, appointments, and preferences. Customers
are the core entity in the pet service business and are used throughout the system for booking appointments, managing
pets, and tracking service history.

| Field Name              | Type                                                     | Description                                                           |
|-------------------------|----------------------------------------------------------|-----------------------------------------------------------------------|
| `id`                    | string                                                   | Unique identifier                                                     |
| `first_name`            | string                                                   | Customer's first name                                                 |
| `last_name`             | string                                                   | Customer's last name                                                  |
| `avatar`                | string                                                   | URL to the customer's profile picture                                 |
| `phone`                 | string                                                   | Customer's phone number. Must be in E.164 format (e.g., +12125551234) |
| `email`                 | string                                                   | Customer's email address. Must be valid                               |
| `address`               | Array(Address)                                           | List of customer's addresses                                          |
| `status`                | Status                                                   | Current status of the customer                                        |
| `color_code`            | string                                                   | Color code for visual identification in the UI                        |
| `source`                | string                                                   | How the customer was acquired                                         |
| `last_appointment_date` | Timestamp                                                | When the customer had their last appointment                          |
| `next_appointment_date` | Timestamp                                                | When the customer's next appointment is scheduled                     |
| `created_by`            | string                                                   | ID of the staff member who created this customer                      |
| `created_time`          | Timestamp                                                | When this customer was created                                        |
| `last_updated_by`       | string                                                   | ID of the staff member who last modified this customer                |
| `last_updated_time`     | Timestamp                                                | When this customer was last modified                                  |
| `preferred_business_id` | string                                                   | ID of the customer's preferred business location                      |
| `company_id`            | string                                                   | ID of the company this customer belongs to                            |
| `notes`                 | Array(Note)                                              | List of notes about this customer                                     |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | List of tags applied to this customer                                 |
| `referral_source`       | [ReferralSource](./setting_service.md#5-referralsource)  | The source or channel through which a customer was acquired           |
| `preference`            | Preference                                               | Customer's communication and marketing preferences                    |

### 2. Note

Represents a comment or observation about a customer. Notes help track important customer information, preferences, and
history.

| Field Name          | Type      | Description                                        |
|---------------------|-----------|----------------------------------------------------|
| `id`                | string    | Unique identifier                                  |
| `note`              | string    | The content of the note                            |
| `last_updated_by`   | string    | ID of the staff member who last modified this note |
| `last_updated_time` | Timestamp | When this note was last modified                   |

### 3. Preference

Stores a customer's communication and marketing preferences. These settings determine how and when we can contact the
customer.

| Field Name                      | Type | Description                                                  |
|---------------------------------|------|--------------------------------------------------------------|
| `receive_auto_message`          | bool | Whether the customer wants to receive automated SMS messages |
| `receive_auto_email`            | bool | Whether the customer wants to receive automated emails       |
| `subscribe_to_marketing_emails` | bool | Whether the customer has opted in to marketing emails        |
| `receive_appointment_reminder`  | bool | Whether the customer wants appointment reminders             |

### 4. Address

Represents a customer's physical address.

see: [Address](common/address.md#-address-moegocommonv1address)

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Customer API

Here is a typical integration flow:

1. **Create Customer**
    - Specify required details like first name, last name, phone number, and company ID.
    - Optionally set email, address, preferences, tags, and notes.

2. **Update Customer**
    - Modify customer details like name, phone, email, or address.
    - Add or remove tags and notes.

3. **Retrieve Customer**
    - Get full details of an existing customer, including notes, tags, and preferences.

4. **List Customers**
    - View all customers matching specified criteria.
    - Filter by last update time if needed.

5. **Manage Notes & Tags**
    - Append new notes or tags to a customer.
    - Retrieve lists of notes and tags associated with a customer.

6. **Generate Card-on-File Link**
    - Generate a secure link for adding payment information.

7. **Monitoring & Maintenance**
    - Regularly retrieve customer data to monitor changes.
    - Update customer records as needed.

---

## 📦 5. API Interface Descriptions

### 1. Create Customer (`CreateCustomer`)

- **Method**: `CreateCustomer`
- **HTTP Method**: POST
- **Path**: `/v1/customers`

#### ✅ Functionality:

Registers a new customer with basic details, preferences, and optionally initial tags and notes.

#### 🎯 Use Cases:

- Users want to onboard new customers into the system.
- Third-party systems sync customer data.

#### 🔧 Request Parameters:

| Field Name              | Type                                                     | Required | Description                                           |
|-------------------------|----------------------------------------------------------|----------|-------------------------------------------------------|
| `company_id`            | string                                                   | Yes      | ID of the company creating the customer               |
| `preferred_business_id` | string                                                   | Yes      | ID of the business location preferred by the customer |
| `first_name`            | string                                                   | Yes      | Customer's first name                                 |
| `last_name`             | string                                                   | Yes      | Customer's last name                                  |
| `phone`                 | string                                                   | Yes      | Customer's phone number                               |
| `email`                 | string                                                   | No       | Customer's email address                              |
| `address`               | Address                                                  | No       | Customer's primary address                            |
| `preference`            | Preference                                               | No       | Customer's communication and marketing preferences    |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | No       | Initial tags to apply to the customer                 |
| `notes`                 | Array(Note)                                              | No       | Initial notes about the customer                      |

#### 💡 Example Request:

```json
{
  "company_id": "cmp_001",
  "preferred_business_id": "biz_001",
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
  "preference": {
    "receive_auto_message": true,
    "receive_auto_email": false,
    "subscribe_to_marketing_emails": true,
    "receive_appointment_reminder": true
  },
  "tags": [
    {
      "name": "VIP"
    }
  ],
  "notes": [
    {
      "note": "Prefers morning appointments."
    }
  ]
}
```

#### 📌 Return Value:

| Field Name              | Type                                                     | Description                                                 |
|-------------------------|----------------------------------------------------------|-------------------------------------------------------------|
| `id`                    | string                                                   | Unique identifier for the customer                          |
| `first_name`            | string                                                   | Customer's first name                                       |
| `last_name`             | string                                                   | Customer's last name                                        |
| `avatar`                | string                                                   | URL to the customer's profile picture                       |
| `phone`                 | string                                                   | Customer's phone number                                     |
| `email`                 | string                                                   | Customer's email address                                    |
| `address`               | Array(Address)                                           | List of customer's addresses                                |
| `status`                | Status                                                   | Current status of the customer                              |
| `color_code`            | string                                                   | Color code for visual identification in the UI              |
| `source`                | string                                                   | How the customer was acquired                               |
| `last_appointment_date` | Timestamp                                                | When the customer had their last appointment                |
| `next_appointment_date` | Timestamp                                                | When the customer's next appointment is scheduled           |
| `created_by`            | string                                                   | ID of the staff member who created this customer            |
| `created_time`          | Timestamp                                                | When this customer was created                              |
| `last_updated_by`       | string                                                   | ID of the staff member who last modified this customer      |
| `last_updated_time`     | Timestamp                                                | When this customer was last modified                        |
| `preferred_business_id` | string                                                   | ID of the customer's preferred business location            |
| `company_id`            | string                                                   | ID of the company this customer belongs to                  |
| `notes`                 | Array(Note)                                              | List of notes about this customer                           |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | List of tags applied to this customer                       |
| `referral_source`       | [ReferralSource](./setting_service.md#5-referralsource)  | The source or channel through which a customer was acquired |
| `preference`            | Preference                                               | Customer's communication and marketing preferences          |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields are missing or invalid.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. Get Customer (`GetCustomer`)

- **Method**: `GetCustomer`
- **HTTP Method**: GET
- **Path**: `/v1/customers/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific customer, including preferences, notes, tags, and appointment history.

#### 🎯 Use Cases:

- Check current customer data.
- Verify customer details during debugging.
- Monitor customer activity.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description             |
|------------|--------|----------|-------------------------|
| `id`       | string | Yes      | Customer ID to retrieve |

#### 📌 Return Value:

| Field Name              | Type                                                     | Description                                                 |
|-------------------------|----------------------------------------------------------|-------------------------------------------------------------|
| `id`                    | string                                                   | Unique identifier for the customer                          |
| `first_name`            | string                                                   | Customer's first name                                       |
| `last_name`             | string                                                   | Customer's last name                                        |
| `avatar`                | string                                                   | URL to the customer's profile picture                       |
| `phone`                 | string                                                   | Customer's phone number                                     |
| `email`                 | string                                                   | Customer's email address                                    |
| `address`               | Array(Address)                                           | List of customer's addresses                                |
| `status`                | Status                                                   | Current status of the customer                              |
| `color_code`            | string                                                   | Color code for visual identification in the UI              |
| `source`                | string                                                   | How the customer was acquired                               |
| `last_appointment_date` | Timestamp                                                | When the customer had their last appointment                |
| `next_appointment_date` | Timestamp                                                | When the customer's next appointment is scheduled           |
| `created_by`            | string                                                   | ID of the staff member who created this customer            |
| `created_time`          | Timestamp                                                | When this customer was created                              |
| `last_updated_by`       | string                                                   | ID of the staff member who last modified this customer      |
| `last_updated_time`     | Timestamp                                                | When this customer was last modified                        |
| `preferred_business_id` | string                                                   | ID of the customer's preferred business location            |
| `company_id`            | string                                                   | ID of the company this customer belongs to                  |
| `notes`                 | Array(Note)                                              | List of notes about this customer                           |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | List of tags applied to this customer                       |
| `referral_source`       | [ReferralSource](./setting_service.md#5-referralsource)  | The source or channel through which a customer was acquired |
| `preference`            | Preference                                               | Customer's communication and marketing preferences          |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 3. Update Customer (`UpdateCustomer`)

- **Method**: `UpdateCustomer`
- **HTTP Method**: PUT
- **Path**: `/v1/customers/{id}`

#### ✅ Functionality:

Updates an existing customer's information, including basic details, preferences, and optionally tags and notes.

#### 🎯 Use Cases:

- Change customer details like name, phone, or address.
- Add or remove tags and notes.
- Update preferences based on customer feedback.

#### 🔧 Request Parameters:

| Field Name              | Type                                                     | Required | Description                                           |
|-------------------------|----------------------------------------------------------|----------|-------------------------------------------------------|
| `id`                    | string                                                   | Yes      | Unique identifier of the customer to update           |
| `company_id`            | string                                                   | Yes      | ID of the company that owns the customer              |
| `preferred_business_id` | string                                                   | No       | ID of the business location preferred by the customer |
| `avatar_path`           | string                                                   | No       | URL to the customer's profile picture                 |
| `first_name`            | string                                                   | No       | Customer's first name                                 |
| `last_name`             | string                                                   | No       | Customer's last name                                  |
| `phone`                 | string                                                   | No       | Customer's phone number                               |
| `email`                 | string                                                   | No       | Customer's email address                              |
| `address`               | Address                                                  | No       | Customer's primary address                            |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | No       | Tags to apply to the customer                         |
| `notes`                 | Array(Note)                                              | No       | Notes about the customer                              |

#### 📌 Return Value:

| Field Name              | Type                                                     | Description                                                 |
|-------------------------|----------------------------------------------------------|-------------------------------------------------------------|
| `id`                    | string                                                   | Unique identifier for the customer                          |
| `first_name`            | string                                                   | Customer's first name                                       |
| `last_name`             | string                                                   | Customer's last name                                        |
| `avatar`                | string                                                   | URL to the customer's profile picture                       |
| `phone`                 | string                                                   | Customer's phone number                                     |
| `email`                 | string                                                   | Customer's email address                                    |
| `address`               | Array(Address)                                           | List of customer's addresses                                |
| `status`                | Status                                                   | Current status of the customer                              |
| `color_code`            | string                                                   | Color code for visual identification in the UI              |
| `source`                | string                                                   | How the customer was acquired                               |
| `last_appointment_date` | Timestamp                                                | When the customer had their last appointment                |
| `next_appointment_date` | Timestamp                                                | When the customer's next appointment is scheduled           |
| `created_by`            | string                                                   | ID of the staff member who created this customer            |
| `created_time`          | Timestamp                                                | When this customer was created                              |
| `last_updated_by`       | string                                                   | ID of the staff member who last modified this customer      |
| `last_updated_time`     | Timestamp                                                | When this customer was last modified                        |
| `preferred_business_id` | string                                                   | ID of the customer's preferred business location            |
| `company_id`            | string                                                   | ID of the company this customer belongs to                  |
| `notes`                 | Array(Note)                                              | List of notes about this customer                           |
| `tags`                  | Array([CustomerTag](./setting_service.md#4-customertag)) | List of tags applied to this customer                       |
| `referral_source`       | [ReferralSource](./setting_service.md#5-referralsource)  | The source or channel through which a customer was acquired |
| `preference`            | Preference                                               | Customer's communication and marketing preferences          |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 4. Delete Customer (**Not Implemented**)

Currently no delete operation is defined for customers. Deletion may be handled via deactivation or archival depending
on business rules.

---

### 5. List Customers (`ListCustomers`)

- **Method**: `ListCustomers`
- **HTTP Method**: POST
- **Path**: `/v1/customers:list`

#### ✅ Functionality:

Lists customers matching the specified criteria, supporting pagination and filtering by last update time.

#### 🎯 Use Cases:

- View all customers under a company.
- Audit or debug customer configurations.
- Monitor customer activity over time.

#### 🔧 Request Parameters:

| Field Name                 | Type       | Required | Description                                            |
|----------------------------|------------|----------|--------------------------------------------------------|
| `pagination`               | Pagination | Yes      | Pagination info: page_size, page_token                 |
| `company_id`               | string     | Yes      | ID of the company to list customers for                |
| `filter.last_updated_time` | Interval   | No       | Time range for filtering customers by last update time |

> **Note**:The `pagination` field is used for pagination.
> The `page_size` field specifies the number of results to return per page. Maximum value is 500.
> The `page_token` field is used to retrieve the next page of results.

#### 📌 Return Value:

| Field Name        | Type            | Description                                     |
|-------------------|-----------------|-------------------------------------------------|
| `next_page_token` | string          | Token for retrieving the next page of results   |
| `customers`       | Array(Customer) | List of customers matching the request criteria |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 6. Generate Card-on-File Link (`GenCustomerCofLink`)

- **Method**: `GenCustomerCofLink`
- **HTTP Method**: GET
- **Path**: `/v1/customers/{id}/cof/link`

#### ✅ Functionality:

Generates a secure link for adding customer card-on-file information.

#### 🎯 Use Cases:

- Allow customers to securely add payment methods.
- Integrate with external payment processors.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                      |
|------------|--------|----------|----------------------------------|
| `id`       | string | Yes      | Customer ID to generate link for |

#### 📌 Return Value:

| Field Name | Type   | Description                                    |
|------------|--------|------------------------------------------------|
| `link`     | string | Secure URL for adding card-on-file information |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 7. Append Customer Notes (`AppendCustomerNotes`)

- **Method**: `AppendCustomerNotes`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{id}/notes`

#### ✅ Functionality:

Adds new notes to a customer's profile.

#### 🎯 Use Cases:

- Record customer preferences or observations.
- Track service history or special requirements.

#### 🔧 Request Parameters:

| Field Name | Type        | Required | Description                  |
|------------|-------------|----------|------------------------------|
| `id`       | string      | Yes      | Customer ID to add notes to  |
| `notes`    | Array(Note) | Yes      | Notes to add to the customer |

#### 📌 Return Value:

| Field Name | Type        | Description                        |
|------------|-------------|------------------------------------|
| `notes`    | Array(Note) | Notes that were successfully added |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 8. List Customer Notes (`ListCustomerNotes`)

- **Method**: `ListCustomerNotes`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{id}/notes:list`

#### ✅ Functionality:

Retrieves a paginated list of notes for a specific customer.

#### 🎯 Use Cases:

- Review historical notes about a customer.
- Audit customer interactions or service history.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                            |
|--------------|------------|----------|----------------------------------------|
| `id`         | string     | Yes      | Customer ID to retrieve notes for      |
| `pagination` | Pagination | Yes      | Pagination info: page_size, page_token |

#### 📌 Return Value:

| Field Name        | Type        | Description                                   |
|-------------------|-------------|-----------------------------------------------|
| `next_page_token` | string      | Token for retrieving the next page of results |
| `notes`           | Array(Note) | List of notes for the customer                |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 9. Append Customer Tags (`AppendCustomerTags`)

- **Method**: `AppendCustomerTags`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{id}/tags`

#### ✅ Functionality:

Adds new tags to a customer's profile.

#### 🎯 Use Cases:

- Categorize customers for reporting or filtering.
- Apply labels for marketing campaigns or service tiers.

#### 🔧 Request Parameters:

| Field Name | Type                                                     | Required | Description                 |
|------------|----------------------------------------------------------|----------|-----------------------------|
| `id`       | string                                                   | Yes      | Customer ID to add tags to  |
| `tags`     | Array([CustomerTag](./setting_service.md#4-customertag)) | Yes      | Tags to add to the customer |

#### 📌 Return Value:

| Field Name | Type                                                     | Description                       |
|------------|----------------------------------------------------------|-----------------------------------|
| `tags`     | Array([CustomerTag](./setting_service.md#4-customertag)) | Tags that were successfully added |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 10. List Customer Tags (`ListCustomerTags`)

- **Method**: `ListCustomerTags`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{id}/tags:list`

#### ✅ Functionality:

Retrieves all tags associated with a specific customer.

#### 🎯 Use Cases:

- Review customer categorizations.
- Audit tag usage for reporting or analytics.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                      |
|------------|--------|----------|----------------------------------|
| `id`       | string | Yes      | Customer ID to retrieve tags for |

#### 📌 Return Value:

| Field Name | Type                                                     | Description                               |
|------------|----------------------------------------------------------|-------------------------------------------|
| `tags`     | Array([CustomerTag](./setting_service.md#4-customertag)) | List of tags associated with the customer |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: Create Customer

```json
{
  "company_id": "cmp_001",
  "preferred_business_id": "biz_001",
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
  "preference": {
    "receive_auto_message": true,
    "receive_auto_email": false,
    "subscribe_to_marketing_emails": true,
    "receive_appointment_reminder": true
  },
  "tags": [
    {
      "name": "VIP"
    }
  ],
  "notes": [
    {
      "note": "Prefers morning appointments."
    }
  ]
}
```

### Example 2: Update Customer

```json
{
  "id": "cus_001",
  "company_id": "cmp_001",
  "preferred_business_id": "biz_001",
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
  "tags": [
    {
      "name": "Frequent Visitor"
    }
  ],
  "notes": [
    {
      "note": "Prefers evening appointments now."
    }
  ]
}
```

### Example 3: List Customers

```json
{
  "company_id": "cmp_001",
  "pagination": {
    "page_size": 20
  },
  "filter": {
    "last_updated_time": {
      "start_time": "2024-08-01T00:00:00Z",
      "end_time": "2024-08-02T00:00:00Z"
    }
  }
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                                  | Answer                                                                                                                         |
|-----------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| How to verify if a customer exists?                       | Use `GetCustomer` to check if the customer ID returns a valid response                                                         |
| Can I create multiple customers at once?                  | Currently only single customer creation is supported. Use batch processing if needed                                           |
| How to filter customers by update time?                   | Use `ListCustomers` with `filter.last_updated_time`                                                                            |
| Why does creating a customer return "resource exhausted"? | The company may have reached the maximum allowed customer count. Clean up unused customers or contact admin to increase quota. |
| How to manage customer tags and notes effectively?        | Use `AppendCustomerTags` and `AppendCustomerNotes` to add new entries                                                          |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Customer ID does not exist        |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/address.md)
- [customer_service.proto](../moego/business/customer/v1/customer_service.proto)
- [customer.proto](../moego/business/customer/v1/customer.proto)