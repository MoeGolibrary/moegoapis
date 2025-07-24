# Discount API Documentation (`moego.business.discount.v1`)

## 📌 1. Functional Overview

Discount represents a promotional offer that reduces the price of services. This module provides the following
functions:

- Creating, retrieving, and managing discount configurations
- Applying discounts either automatically or manually to appointments
- Defining various discount types (fixed amount or percentage-based)
- Setting validity periods and usage limitations
- Controlling discount eligibility through business locations and customer segments
- Supporting both online booking integration and staff application scenarios

Applicable to scenarios such as seasonal promotions, loyalty rewards, targeted marketing campaigns, and special offers.

---

## 🎯 2. Design Goals

- **Flexible Promotion Management**: Provides comprehensive tools for creating and managing various types of discounts
- **Targeted Application**: Supports precise targeting of discounts through customer segmentation and business location
  restrictions
- **Usage Control**: Implements robust limitations to prevent abuse and manage discount budgets
- **Time-based Management**: Enables scheduling of discounts for specific time periods
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

---

## 🧩 3. Core Concepts

### 1. Discount

Represents a promotional offer that reduces the price of services. Discounts can be applied automatically or manually to
appointments and can be limited by various criteria such as usage count, validity period, and eligible customers or
business locations.

#### Discount.Value (oneof)

| Field Type   | Description                                                                                 |
|--------------|---------------------------------------------------------------------------------------------|
| `amount`     | Fixed amount to deduct from the service price (must be in the same currency as the service) |
| `percentage` | Percentage to deduct from the service price (range: 1-100)                                  |

| Field Name     | Type                      | Description                                                                                                |
|----------------|---------------------------|------------------------------------------------------------------------------------------------------------|
| `code`         | string                    | Unique identifier for the discount (alphanumeric code, case-sensitive)                                     |
| `description`  | string                    | Human-readable explanation of the discount used in customer communications and UI                          |
| `valid_period` | google.type.Interval      | Time period during which the discount is valid for controlling seasonal or promotional offers              |
| `limitation`   | DiscountLimitation        | Usage restrictions and eligibility criteria that control who can use the discount and how often            |
| `settings`     | DiscountSettings          | Configuration options for discount application that control how the discount behaves in different contexts |
| `expiry_time`  | google.protobuf.Timestamp | Optional field indicating when this discount becomes invalid (if not set, discount never expires)          |

### 2. DiscountLimitation

Defines usage restrictions for a discount. These limitations help control discount usage and target specific customer
segments or business locations.

| Field Name                 | Type          | Description                                                                                 |
|----------------------------|---------------|---------------------------------------------------------------------------------------------|
| `max_redeem_times`         | uint32        | Maximum number of times this discount can be used (set to 0 for unlimited usage)            |
| `business_ids`             | Array(string) | Business locations where this discount is valid (empty list means valid at all locations)   |
| `redeem_once_per_customer` | bool          | Whether each customer can use this discount only once (helps prevent discount abuse)        |
| `customer_ids`             | Array(string) | Specific customers eligible for this discount (empty list means all customers are eligible) |

### 3. DiscountSettings

Configures how the discount is applied and presented. These settings control the discount's behavior in different
booking channels and scenarios.

| Field Name                           | Type | Description                                                                                                              |
|--------------------------------------|------|--------------------------------------------------------------------------------------------------------------------------|
| `auto_apply_on_eligible_appointment` | bool | Whether to apply the discount automatically when conditions are met (used for promotional campaigns and loyalty rewards) |
| `allow_for_online_booking`           | bool | Whether this discount can be used in online bookings (controls discount visibility in customer portal)                   |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Discount API

Here is a typical integration flow:

1. **CreateDiscount**
    - Define a new discount with specific configuration including code, value type (amount or percentage), validity
      period, and settings.
    - Set usage limitations like maximum redemption times and eligible business locations or customers.

2. **GetDiscount**
    - Retrieve detailed information about a specific discount using its unique code.
    - Verify discount configuration and check current usage statistics.

3. **ListDiscounts**
    - View all available discounts for a company.
    - Monitor active promotions or manage existing discounts.
    - Filter results if needed based on usage or validity criteria.

4. **Promotion Management & Monitoring**
    - Regularly review discount usage to ensure compliance with business goals.
    - Track redemption rates and adjust limitations as needed.
    - Analyze effectiveness of different discount strategies.

---

## 📦 5. API Interface Descriptions

### 1. CreateDiscount (`CreateDiscount`)

- **Method**: `CreateDiscount`
- **HTTP Method**: POST
- **Path**: `/v1/discounts`

#### ✅ Functionality:

Creates a new discount with the specified configuration.

#### 🎯 Use Cases:

- Launching seasonal promotions
- Setting up loyalty reward discounts
- Creating targeted marketing offers

#### 🔧 Request Parameters:

| Field Name   | Type     | Required | Description                                                                                |
|--------------|----------|----------|--------------------------------------------------------------------------------------------|
| `company_id` | string   | Yes      | Company identifier for scoping the discount creation                                       |
| `discount`   | Discount | Yes      | Complete discount configuration to create including all required fields and valid settings |

#### 📌 Return Value:

| Field Name     | Type                      | Description                                                                                                |
|----------------|---------------------------|------------------------------------------------------------------------------------------------------------|
| `code`         | string                    | Unique identifier for the discount (alphanumeric code, case-sensitive)                                     |
| `description`  | string                    | Human-readable explanation of the discount used in customer communications and UI                          |
| `value`        | oneof(amount, percentage) | Discount value, either a fixed amount or percentage                                                        |
| `valid_period` | google.type.Interval      | Time period during which the discount is valid for controlling seasonal or promotional offers              |
| `limitation`   | DiscountLimitation        | Usage restrictions and eligibility criteria that control who can use the discount and how often            |
| `settings`     | DiscountSettings          | Configuration options for discount application that control how the discount behaves in different contexts |
| `expiry_time`  | google.protobuf.Timestamp | Optional field indicating when this discount becomes invalid (if not set, discount never expires)          |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Discount configuration is invalid (e.g., conflicting parameters, invalid values)
- `ALREADY_EXISTS`: Discount code is already in use
- `PERMISSION_DENIED`: Permission denied

---

### 2. GetDiscount (`GetDiscount`)

- **Method**: `GetDiscount`
- **HTTP Method**: GET
- **Path**: `/v1/discounts/{code}`

#### ✅ Functionality:

Retrieves detailed information about a specific discount by its code.

#### 🎯 Use Cases:

- Verifying discount configuration before applying it to appointments
- Auditing discount details and usage rules
- Checking discount status and validity

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description                                   |
|--------------|--------|----------|-----------------------------------------------|
| `company_id` | string | Yes      | Company identifier for scoping the request    |
| `code`       | string | Yes      | Unique identifier of the discount to retrieve |

#### 📌 Return Value:

| Field Name     | Type                      | Description                                                                                                |
|----------------|---------------------------|------------------------------------------------------------------------------------------------------------|
| `code`         | string                    | Unique identifier for the discount (alphanumeric code, case-sensitive)                                     |
| `description`  | string                    | Human-readable explanation of the discount used in customer communications and UI                          |
| `value`        | oneof(amount, percentage) | Discount value, either a fixed amount or percentage                                                        |
| `valid_period` | google.type.Interval      | Time period during which the discount is valid for controlling seasonal or promotional offers              |
| `limitation`   | DiscountLimitation        | Usage restrictions and eligibility criteria that control who can use the discount and how often            |
| `settings`     | DiscountSettings          | Configuration options for discount application that control how the discount behaves in different contexts |
| `expiry_time`  | google.protobuf.Timestamp | Optional field indicating when this discount becomes invalid (if not set, discount never expires)          |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified discount code does not exist
- `PERMISSION_DENIED`: Permission denied

---

### 3. ListDiscounts (`ListDiscounts`)

- **Method**: `ListDiscounts`
- **HTTP Method**: POST
- **Path**: `/v1/discounts:list`

#### ✅ Functionality:

Retrieves a paginated list of discounts based on specified criteria. Results include both active and expired discounts
for historical reference.

#### 🎯 Use Cases:

- Viewing all available discounts for a company
- Managing existing discount inventory
- Analyzing discount performance across different time periods

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                                         |
|--------------|------------|----------|-----------------------------------------------------|
| `pagination` | Pagination | Yes      | Pagination info: page_size, page_token              |
| `company_id` | string     | Yes      | Company identifier for scoping the discount listing |

#### 📌 Return Value:

| Field Name        | Type            | Description                                                          |
|-------------------|-----------------|----------------------------------------------------------------------|
| `next_page_token` | string          | Token for retrieving the next page of results (empty if none remain) |
| `discounts`       | Array(Discount) | List of discounts matching the request criteria                      |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: CreateDiscount

```json
{
  "company_id": "cmp_001",
  "discount": {
    "code": "SUMMER25",
    "description": "Summer Special - 25% off grooming services",
    "value": {
      "percentage": 25
    },
    "valid_period": {
      "start_time": "2024-06-01T00:00:00Z",
      "end_time": "2024-08-31T23:59:59Z"
    },
    "limitation": {
      "max_redeem_times": 500,
      "business_ids": [
        "biz_001",
        "biz_002"
      ],
      "redeem_once_per_customer": true
    },
    "settings": {
      "auto_apply_on_eligible_appointment": true,
      "allow_for_online_booking": true
    },
    "expiry_time": "2024-09-15T00:00:00Z"
  }
}
```

### Example 2: GetDiscount

```json
{
  "company_id": "cmp_001",
  "code": "SUMMER25"
}
```

### Example 3: ListDiscounts

```json
{
  "pagination": {
    "page_size": 20
  },
  "company_id": "cmp_001"
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                                  | Answer                                                                                                             |
|-----------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| How to verify if a discount exists?                       | Use `GetDiscount` to check if the discount code returns a valid response                                           |
| Can I create multiple discounts with the same code?       | No, each discount code must be unique within a company                                                             |
| How to limit discount usage to specific customers?        | Use `limitation.customer_ids` to specify eligible customers                                                        |
| Why does creating a discount return "resource exhausted"? | Not applicable — discounts typically don't have hard limits unless configured                                      |
| How to handle expired discounts?                          | Use `ListDiscounts` to view expired discounts for historical reference; they cannot be applied to new appointments |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Discount code does not exist      |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `ALREADY_EXISTS`    | Discount code is already in use   |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/address.md)
- [discount_service.proto](../moego/business/discount/v1/discount_service.proto)
- [discount.proto](../moego/business/discount/v1/discount.proto)