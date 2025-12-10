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

| Field Name                   | Type                      | Description                                                                             |
|------------------------------|---------------------------|-----------------------------------------------------------------------------------------|
| `id`                         | string                    | Unique identifier for the membership plan. Format: "mem_" followed by random characters |
| `internalProductId`          | string                    | The product id in subscription service                                                  |
| `name`                       | string                    | Display name of the membership plan                                                     |
| `description`                | string                    | Detailed description of the membership plan                                             |
| `status`                     | enum Status               | Current status of the membership plan                                                   |
| `price`                      | google.type.Money         | Price of the membership plan                                                            |
| `taxId`                      | string                    | The tax id associated with the membership                                               |
| `policy`                     | string                    | Policy information for the membership                                                   |
| `companyId`                  | string                    | The company id that owns this membership                                                |
| `priceId`                    | string                    | The price id                                                                            |
| `billingCyclePeriod`         | TimePeriod                | Billing cycle period                                                                    |
| `createdTime`                | google.protobuf.Timestamp | Timestamp when the membership plan was created                                          |
| `lastUpdatedTime`            | google.protobuf.Timestamp | Timestamp when the membership plan was last updated                                     |
| `deletedTime`                | google.protobuf.Timestamp | Timestamp when the membership plan was deleted (if applicable)                          |
| `revision`                   | int32                     | Revision number for tracking updates                                                    |
| `totalPrice`                 | google.type.Money         | Total price including taxes                                                             |
| `totalTax`                   | google.type.Money         | Total tax amount                                                                        |
| `enableOnlineBooking`        | bool                      | Whether this membership can be purchased through online booking                         |
| `enableDiscountBenefits`     | bool                      | Whether discount benefits are enabled for this membership                               |
| `enableQuantityBenefits`     | bool                      | Whether quantity benefits are enabled for this membership                               |
| `billingCycleDayOfWeek`      | google.type.DayOfWeek     | Billing cycle day of week                                                               |
| `breedFilter`                | bool                      | Breed filter                                                                            |
| `customizedBreeds`           | repeated CustomizedBreed  | Customized breed                                                                        |
| `petSizeFilter`              | bool                      | Available for all pet size                                                              |
| `customizedPetSizes`         | repeated string           | Available pet size (only if is_available_for_all_pet_size is false)                     |
| `coatFilter`                 | bool                      | Available for all pet coat type                                                         |
| `customizedCoats`            | repeated string           | Available pet coat type (only if is_available_for_all_pet_coat_type is false)           |
| `source`                     | enum Source               | Source of the membership                                                                |
| `billingCycleTimeOfDay`      | google.type.TimeOfDay     | Billing cycle time of day                                                               |
| `allowBillingCycleTimeOfDay` | bool                      | Allow billing cycle time of day                                                         |

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

| Field Name          | Type                      | Description                                |
|---------------------|---------------------------|--------------------------------------------|
| `id`                | string                    | Unique identifier for the subscription     |
| `customerId`        | string                    | The customer identifier                    |
| `membershipId`      | string                    | The membership identifier                  |
| `companyId`         | string                    | The company identifier                     |
| `price`             | google.type.Money         | The price of the subscription              |
| `createdTime`       | google.protobuf.Timestamp | The create time                            |
| `lastUpdatedTime`   | google.protobuf.Timestamp | The update time                            |
| `deletedTime`       | google.protobuf.Timestamp | The delete time, non-null means is deleted |
| `validityPeriod`    | google.type.Interval      | Validity period                            |
| `nextBillingDate`   | google.protobuf.Timestamp | Next billing date                          |
| `expiredTime`       | google.protobuf.Timestamp | Expire date                                |
| `cancelAtPeriodEnd` | bool                      | Cancelled but in active status             |
| `status`            | enum Status               | Status of the subscription                 |

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

### 4. IncludeBenefitView

Represents a membership benefit view with usage details.

#### Fields

| Field Name          | Type                          | Description                              |
|---------------------|-------------------------------|------------------------------------------|
| `id`                | string                        | History ID                               |
| `membershipId`      | string                        | Membership identifier                    |
| `totalQuantity`     | int32                         | Total quantity                           |
| `remainingQuantity` | int32                         | Remaining quantity                       |
| `isLimited`         | bool                          | Whether the benefit is limited           |
| `redeemTime`        | google.protobuf.Timestamp     | Time when the benefit was redeemed       |
| `itemDetails`       | repeated RedeemItemDetailView | Details of items in the benefit          |
| `isAll`             | bool                          | Whether the benefit applies to all items |
| `discountUnit`      | enum DiscountUnit             | Unit of discount (percent or numerical)  |
| `discountValue`     | double                        | Value of the discount                    |

#### RedeemItemDetailView

| Field Name | Type              | Description                           |
|------------|-------------------|---------------------------------------|
| `itemId`   | string            | Item identifier                       |
| `itemName` | string            | Item name                             |
| `price`    | google.type.Money | Item amount                           |
| `itemType` | enum TargetType   | Type of item (service, product, etc.) |

#### TargetType Enum

| Value                     | Description             |
|---------------------------|-------------------------|
| `TARGET_TYPE_UNSPECIFIED` | Unspecified target type |
| `SERVICE`                 | Service                 |
| `ADDON`                   | Add-on                  |
| `PRODUCT`                 | Product                 |
| `PACKAGE`                 | Package                 |

#### DiscountUnit Enum

| Value              | Description           |
|--------------------|-----------------------|
| `UNIT_UNSPECIFIED` | Unspecified unit      |
| `PERCENT`          | Percentage discount   |
| `NUMERICAL`        | Fixed amount discount |

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

| Field Name | Type               | Required | Description                                                         |
|------------|--------------------|----------|---------------------------------------------------------------------|
| `nameLike` | string             | No       | Partial name match to filter memberships by name (case-insensitive) |
| `statuses` | Array(enum Status) | No       | List of membership statuses to include in results                   |

#### 📌 Return Value:

| Field Name      | Type              | Description                                                          |
|-----------------|-------------------|----------------------------------------------------------------------|
| `memberships`   | Array(Membership) | List of memberships matching the request criteria                    |
| `nextPageToken` | string            | Token for retrieving the next page of results (empty if none remain) |

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

| Field Name      | Type               | Required | Description                                         |
|-----------------|--------------------|----------|-----------------------------------------------------|
| `customerIds`   | Array(string)      | No       | List of customer IDs to include in results          |
| `membershipIds` | Array(string)      | No       | List of membership IDs to include in results        |
| `statuses`      | Array(enum Status) | No       | List of subscription statuses to include in results |

#### 📌 Return Value:

| Field Name      | Type                | Description                                                          |
|-----------------|---------------------|----------------------------------------------------------------------|
| `subscriptions` | Array(Subscription) | List of subscriptions matching the request criteria                  |
| `nextPageToken` | string              | Token for retrieving the next page of results (empty if none remain) |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid
- `PERMISSION_DENIED`: Permission denied

### 3. GetPerkUsageDetail (`GetPerkUsageDetail`)

- **Method**: `GetPerkUsageDetail`
- **HTTP Method**: POST
- **Path**: `/v1/memberships/perkUsageDetail`

#### ✅ Functionality:

Retrieves detailed information about perk usage for a specific membership and customer. This includes both included
benefits and discount benefits that are associated with the membership.

#### 🎯 Use Cases:

- View detailed perk usage information for a customer's membership
- Check remaining quantities of service benefits
- Review discount benefits available to a customer

#### 🔧 Request Parameters:

| Field Name     | Type   | Required | Description                      |
|----------------|--------|----------|----------------------------------|
| `filter`       | Filter | No       | Optional filters for the request |
| `companyId`    | string | Yes      | Company ID                       |
| `membershipId` | string | Yes      | Membership ID                    |
| `customerId`   | string | Yes      | Customer ID                      |

##### Filter Object

| Field Name          | Type                      | Required | Description         |
|---------------------|---------------------------|----------|---------------------|
| `validityStartTime` | google.protobuf.Timestamp | No       | Validity start time |

#### 📌 Return Value:

| Field Name         | Type                      | Description           |
|--------------------|---------------------------|-----------------------|
| `includedBenefits` | Array(IncludeBenefitView) | The included benefits |
| `discountBenefits` | Array(IncludeBenefitView) | Discount benefits     |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Request parameters are invalid
- `NOT_FOUND`: Membership or customer not found
- `PERMISSION_DENIED`: Permission denied

---

## 🧪 6. Usage Examples

### Example 1: ListMemberships

```json
{
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "companyId": "cmp_001",
  "filter": {
    "nameLike": "gold",
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
    "pageSize": 20,
    "pageToken": "1"
  },
  "companyId": "cmp_001",
  "filter": {
    "customerIds": [
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

### Example 3: GetPerkUsageDetail

```json
{
  "companyId": "cmp_001",
  "membershipId": "mem_abc123",
  "customerId": "cus_xyz789"
}

```

Response:

```json
{
  "includedBenefits": [
    {
      "id": "inc_001",
      "membershipId": "mem_abc123",
      "totalQuantity": 10,
      "remainingQuantity": 5,
      "isLimited": true,
      "redeemTime": "2023-01-15T10:00:00Z",
      "itemDetails": [
        {
          "itemId": "srv_001",
          "itemName": "Basic Grooming",
          "price": {
            "currencyCode": "USD",
            "units": 50,
            "nanos": 0
          },
          "itemType": "SERVICE"
        }
      ],
      "isAll": false,
      "discountUnit": "UNIT_UNSPECIFIED",
      "discountValue": 0
    }
  ],
  "discountBenefits": [
    {
      "id": "disc_001",
      "membershipId": "mem_abc123",
      "totalQuantity": 0,
      "remainingQuantity": 0,
      "isLimited": false,
      "redeemTime": "2023-01-15T10:00:00Z",
      "itemDetails": [],
      "isAll": true,
      "discountUnit": "PERCENT",
      "discountValue": 15.0
    }
  ]
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