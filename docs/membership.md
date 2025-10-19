# Membership API Documentation (`moego.business.membership.v1`)

## 📌 1. Functional Overview

The Membership module manages membership plans and customer subscriptions. It provides the following functions:

- Creating and managing membership plans
- Managing customer subscriptions to membership plans
- Listing memberships and subscriptions with filtering and pagination capabilities

This module is applicable to scenarios such as membership management, subscription tracking, and customer benefits
administration.

---

## 🎯 2. Design Goals

- **Flexible Membership Plans**: Support various membership configurations with customizable benefits
- **Subscription Management**: Track customer subscriptions with detailed status information
- **Easy Integration**: RESTful API compatible with mainstream frameworks
- **Secure Access**: Role-based access control ensures data integrity

---

## 🧩 3. Core Concepts

### 1. Membership

Represents a membership plan that customers can purchase and use for benefits.

#### Fields

| Field Name                        | Type                      | Description                                                                             |
|-----------------------------------|---------------------------|-----------------------------------------------------------------------------------------|
| `id`                              | string                    | Unique identifier for the membership plan. Format: "mem_" followed by random characters |
| `internal_product_id`             | string                    | The product id in subscription service                                                  |
| `name`                            | string                    | Display name of the membership plan                                                     |
| `description`                     | string                    | Detailed description of the membership plan                                             |
| `status`                          | enum Status               | Current status of the membership plan                                                   |
| `price`                           | google.type.Money         | Price of the membership plan                                                            |
| `tax_id`                          | string                    | The tax id associated with the membership                                               |
| `policy`                          | string                    | Policy information for the membership                                                   |
| `company_id`                      | string                    | The company id that owns this membership                                                |
| `price_id`                        | string                    | The price id                                                                            |
| `billing_cycle_period`            | TimePeriod                | Billing cycle period                                                                    |
| `created_time`                    | google.protobuf.Timestamp | Timestamp when the membership plan was created                                          |
| `last_updated_time`               | google.protobuf.Timestamp | Timestamp when the membership plan was last updated                                     |
| `deleted_time`                    | google.protobuf.Timestamp | Timestamp when the membership plan was deleted (if applicable)                          |
| `revision`                        | int32                     | Revision number for tracking updates                                                    |
| `total_price`                     | google.type.Money         | Total price including taxes                                                             |
| `total_tax`                       | google.type.Money         | Total tax amount                                                                        |
| `enable_online_booking`           | bool                      | Whether this membership can be purchased through online booking                         |
| `enable_discount_benefits`        | bool                      | Whether discount benefits are enabled for this membership                               |
| `enable_quantity_benefits`        | bool                      | Whether quantity benefits are enabled for this membership                               |
| `billing_cycle_day_of_week`       | google.type.DayOfWeek     | Billing cycle day of week                                                               |
| `breed_filter`                    | bool                      | Breed filter                                                                            |
| `customized_breeds`               | repeated CustomizedBreed  | Customized breed                                                                        |
| `pet_size_filter`                 | bool                      | Available for all pet size                                                              |
| `customized_pet_sizes`            | repeated string           | Available pet size (only if is_available_for_all_pet_size is false)                     |
| `coat_filter`                     | bool                      | Available for all pet coat type                                                         |
| `customized_coats`                | repeated string           | Available pet coat type (only if is_available_for_all_pet_coat_type is false)           |
| `source`                          | enum Source               | Source of the membership                                                                |
| `billing_cycle_time_of_day`       | google.type.TimeOfDay     | Billing cycle time of day                                                               |
| `allow_billing_cycle_time_of_day` | bool                      | Allow billing cycle time of day                                                         |

#### Status Enum

| Value                | Description                                                                                        |
|----------------------|----------------------------------------------------------------------------------------------------|
| `STATUS_UNSPECIFIED` | Indicates an unknown or invalid status                                                             |
| `ACTIVE`             | The membership plan is currently available for purchase and use                                    |
| `INACTIVE`           | The membership plan is not available for new purchases but existing memberships may still be valid |

#### Source Enum

| Value                | Description                                        |
|----------------------|----------------------------------------------------|
| `SOURCE_UNSPECIFIED` | Indicates an unknown or invalid source             |
| `MOEGO_PLATFORM`     | The membership plan is created from MoeGo platform |
| `ENTERPRISE_HUB`     | The membership plan is created from Enterprise Hub |

### 2. Subscription

Represents a customer's subscription to a membership plan.

#### Fields

| Field Name             | Type                      | Description                                |
|------------------------|---------------------------|--------------------------------------------|
| `id`                   | string                    | Unique identifier for the subscription     |
| `customer_id`          | string                    | The customer identifier                    |
| `membership_id`        | string                    | The membership identifier                  |
| `company_id`           | string                    | The company identifier                     |
| `price`                | google.type.Money         | The price of the subscription              |
| `created_time`         | google.protobuf.Timestamp | The create time                            |
| `last_updated_time`    | google.protobuf.Timestamp | The update time                            |
| `deleted_time`         | google.protobuf.Timestamp | The delete time, non-null means is deleted |
| `validity_period`      | google.type.Interval      | Validity period                            |
| `next_billing_date`    | google.protobuf.Timestamp | Next billing date                          |
| `expired_time`         | google.protobuf.Timestamp | Expire date                                |
| `cancel_at_period_end` | bool                      | Cancelled but in active status             |
| `status`               | enum Status               | Status of the subscription                 |

#### Status Enum

| Value                | Description                                                       |
|----------------------|-------------------------------------------------------------------|
| `STATUS_UNSPECIFIED` | Indicates an unknown or invalid status                            |
| `TRIAL`              | The subscription is in trial period                               |
| `PENDING`            | The subscription is waiting for the first charge result           |
| `ACTIVE`             | The subscription is charged and in validity period                |
| `GRACE`              | The subscription is in grace period                               |
| `EXPIRED`            | The subscription is expired and out of validity period            |
| `CANCELLED`          | The subscription is manually cancelled and out of validity period |
| `PAUSED`             | The subscription is currently paused                              |
| `INCOMPLETE`         | The subscription is not completed                                 |

### 3. TimePeriod

Represents a period of time with a unit and value.

#### Fields

| Field Name | Type                       | Description                                   |
|------------|----------------------------|-----------------------------------------------|
| `unit`     | google.type.CalendarPeriod | The unit of time for this period              |
| `value`    | int32                      | The value of the period in the specified unit |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Membership API

Here is a typical integration flow:

1. **ListMemberships**
    - Retrieve a list of all membership plans.
    - Filter by name, status, or IDs if needed.
    - Use pagination to manage large datasets.

2. **Manage Subscriptions**
    - List existing subscriptions for customers.
    - Filter subscriptions by customer, membership, or status.
    - Track subscription lifecycle through different statuses.

---

## 📦 5. API Interface Descriptions

### 1. ListMemberships (`ListMemberships`)

- **Method**: `ListMemberships`
- **HTTP Method**: POST
- **Path**: `/v1/memberships:list`

#### ✅ Functionality:

Retrieves a paginated list of memberships based on specified criteria. Supports filtering by name, status, and IDs to
facilitate targeted queries.

#### 🎯 Use Cases:

- View all membership plans in the system.
- Filter membership plans by status (active/inactive).
- Search for specific membership plans by partial name match.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                                      |
|--------------|------------|----------|--------------------------------------------------|
| `pagination` | Pagination | Yes      | Pagination info: pageSize, pageToken             |
| `companyId`  | string     | Yes      | Company ID to scope memberships                  |
| `filter`     | Filter     | No       | Optional filters to apply to the membership list |

##### Filter Object

| Field Name  | Type               | Required | Description                                                         |
|-------------|--------------------|----------|---------------------------------------------------------------------|
| `name_like` | string             | No       | Partial name match to filter memberships by name (case-insensitive) |
| `statuses`  | Array(enum Status) | No       | List of membership statuses to include in results                   |

#### 📌 Return Value:

| Field Name        | Type              | Description                                                          |
|-------------------|-------------------|----------------------------------------------------------------------|
| `memberships`     | Array(Membership) | List of memberships matching the request criteria                    |
| `next_page_token` | string            | Token for retrieving the next page of results (empty if none remain) |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid
- `PERMISSION_DENIED`: Permission denied

### 2. ListSubscriptions (`ListSubscriptions`)

- **Method**: `ListSubscriptions`
- **HTTP Method**: POST
- **Path**: `/v1/memberships/subscriptions:list`

#### ✅ Functionality:

Retrieves a paginated list of subscriptions based on specified criteria. Supports filtering by customer, membership, and
status to facilitate targeted queries.

#### 🎯 Use Cases:

- View all subscriptions in the system.
- Filter subscriptions by customer or membership.
- Track subscription status across the platform.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                                        |
|--------------|------------|----------|----------------------------------------------------|
| `pagination` | Pagination | Yes      | Pagination info: pageSize, pageToken               |
| `companyId`  | string     | Yes      | Company ID to scope subscriptions                  |
| `filter`     | Filter     | No       | Optional filters to apply to the subscription list |

##### Filter Object

| Field Name       | Type               | Required | Description                                         |
|------------------|--------------------|----------|-----------------------------------------------------|
| `customer_ids`   | Array(string)      | No       | List of customer IDs to include in results          |
| `membership_ids` | Array(string)      | No       | List of membership IDs to include in results        |
| `statuses`       | Array(enum Status) | No       | List of subscription statuses to include in results |

#### 📌 Return Value:

| Field Name        | Type                | Description                                                          |
|-------------------|---------------------|----------------------------------------------------------------------|
| `subscriptions`   | Array(Subscription) | List of subscriptions matching the request criteria                  |
| `next_page_token` | string              | Token for retrieving the next page of results (empty if none remain) |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: ListMemberships

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "filter": {
    "name_like": "gold",
    "statuses": [
      "ACTIVE"
    ]
  }
}
```

### Example 2: ListSubscriptions

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "filter": {
    "customer_ids": [
      "cus_123",
      "cus_456"
    ],
    "statuses": [
      "ACTIVE",
      "TRIAL"
    ]
  }
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                               | Answer                                                                |
|--------------------------------------------------------|-----------------------------------------------------------------------|
| How to verify if a membership exists?                  | Use `ListMemberships` with specific filters                           |
| Can I list memberships for multiple companies at once? | Currently only supports listing memberships for one company at a time |
| How to filter memberships by status?                   | Use `ListMemberships` with `statuses` filter                          |
| How to handle expired subscriptions?                   | Use `ListSubscriptions` with `statuses` filter including `EXPIRED`    |

---

## 📌 9. Common Error Codes

| Error Code          | Description                               |
|---------------------|-------------------------------------------|
| `NOT_FOUND`         | Membership/Subscription ID does not exist |
| `PERMISSION_DENIED` | Current user has no access rights         |
| `INVALID_ARGUMENT`  | Invalid request parameters                |
| `INTERNAL`          | Internal server error                     |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/pagination.md)
- [membership_service.proto](../moego/business/membership/v1/membership_service.proto)
- [membership.proto](../moego/business/membership/v1/membership.proto)