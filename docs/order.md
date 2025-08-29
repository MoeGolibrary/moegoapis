# Order API Documentation (`moego.business.order.v1`)

## 📌 1. Functional Overview

Order represents a financial transaction for services provided to a customer. This module provides the following
functions:

- Creating, retrieving, and updating order details
- Listing orders with support for filtering (by business ID, status, and update time)
- Managing order line items representing individual services or products
- Tracking comprehensive financial details including subtotal, tax, tips, discounts, fees, total amount, paid amount,
  and refunds
- Supporting multiple business modules and payment statuses

Applicable to scenarios such as order management, financial reporting, service delivery tracking, and customer
transaction history lookup.

---

## 🎯 2. Design Goals

- **Centralized Transaction Management**: Provides a unified interface for managing all aspects of customer orders
- **Flexible Filtering**: Supports filtering by business IDs, status, and time ranges to customize result sets
- **Comprehensive Financial Tracking**: Maintains detailed records of all financial components including taxes, tips,
  discounts, and refunds
- **Rich Itemization**: Tracks individual line items for services, products, packages, and special charges
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

---

## 🧩 3. Core Concepts

### 1. Order

Represents a financial transaction for services provided to a customer. Orders track the complete financial details of a
transaction, including services, tips, taxes, discounts, and payment status. Each order is associated with a specific
business location and customer.

#### Order.Status (enum)

| Value                | Description                                                                              |
|----------------------|------------------------------------------------------------------------------------------|
| `STATUS_UNSPECIFIED` | Unknown or invalid status; should not be used when creating new orders                   |
| `CREATED`            | Order has been created but processing hasn't started; initial state for new orders       |
| `PROCESSING`         | Order is currently being processed; services are being delivered or payments are pending |
| `COMPLETED`          | Order has been fully completed; all services delivered and payments processed            |
| `REMOVED`            | Order has been removed from active records; historical data is maintained for reporting  |

| Field Name        | Type                      | Description                                                                                            |
|-------------------|---------------------------|--------------------------------------------------------------------------------------------------------|
| `id`              | string                    | Unique identifier for the order, obfuscated ID string (format: "ord_" + random characters)             |
| `businessId`      | string                    | ID of the business location where services were provided, obfuscated ID string                         |
| `customerId`      | string                    | ID of the customer who received services, obfuscated ID string                                         |
| `status`          | Status                    | Current status of the order                                                                            |
| `title`           | string                    | Short description of the order (e.g., "Full Grooming Package - Max")                                   |
| `description`     | string                    | Detailed notes about the order                                                                         |
| `tipsAmount`      | google.type.Money         | Additional amount provided as gratuity (optional, non-negative)                                        |
| `taxAmount`       | google.type.Money         | Tax amount applied to the order (calculated based on local tax rates)                                  |
| `discountAmount`  | google.type.Money         | Total discounts applied to the order (sum of all applicable discounts)                                 |
| `extraFeeAmount`  | google.type.Money         | Additional fees applied to the order (e.g., holiday surcharge)                                         |
| `subTotalAmount`  | google.type.Money         | Subtotal before tax, tips, and adjustments (sum of all service prices)                                 |
| `tipsBasedAmount` | google.type.Money         | Amount used as basis for calculating tips (usually equals subTotalAmount)                              |
| `totalAmount`     | google.type.Money         | Total amount including all charges and adjustments (formula: subtotal + tax + tips + fees - discounts) |
| `paidAmount`      | google.type.Money         | Amount that has been paid so far (updated when payments are processed)                                 |
| `remainAmount`    | google.type.Money         | Amount still due on the order (formula: totalAmount - paidAmount)                                      |
| `refundedAmount`  | google.type.Money         | Amount that has been refunded (tracks any refunds issued to the customer)                              |
| `createdBy`       | string                    | ID of the staff member who created the order, obfuscated ID string                                     |
| `createdTime`     | google.protobuf.Timestamp | When this order was created (system-generated timestamp)                                               |
| `lastUpdatedBy`   | string                    | ID of the staff member who last modified the order, obfuscated ID string                               |
| `lastUpdatedTime` | google.protobuf.Timestamp | When this order was last modified (system-generated timestamp)                                         |
| `salesDatetime`   | google.protobuf.Timestamp | When the sale was recorded (used for financial reporting and reconciliation)                           |

### 2. OrderLineItem

Represents an individual item within an order. Each line item corresponds to a specific service, product, or charge that
contributes to the total order amount. Line items track individual pricing, quantities, and payment status.

#### OrderLineItem.ObjectType (enum)

| Value                     | Description                                                                     |
|---------------------------|---------------------------------------------------------------------------------|
| `OBJECT_TYPE_UNSPECIFIED` | Unknown or invalid object type; should not be used when creating new line items |
| `SERVICE`                 | Standard service offering (e.g., grooming, training session)                    |
| `PRODUCT`                 | Physical product sold (e.g., pet food, grooming supplies)                       |
| `PACKAGE`                 | Bundle of services or products (typically offered at a discount)                |
| `NO_SHOW`                 | Charge for missed appointment (based on cancellation policy)                    |
| `SERVICE_CHARGE`          | Additional service-related fee (e.g., holiday surcharge, rush fee)              |
| `EVALUATION_SERVICE`      | Initial pet evaluation service (required for new clients or services)           |

| Field Name        | Type                      | Description                                                                                   |
|-------------------|---------------------------|-----------------------------------------------------------------------------------------------|
| `id`              | string                    | Unique identifier for the line item, obfuscated ID string (format: "li_" + random characters) |
| `orderId`         | string                    | ID of the parent order, obfuscated ID string                                                  |
| `objectId`        | string                    | ID of the referenced service, product, or package, obfuscated ID string                       |
| `objectType`      | ObjectType                | Category of the line item (determines business rules and processing)                          |
| `name`            | string                    | Display name of the item (e.g., "Premium Dog Grooming")                                       |
| `description`     | string                    | Additional details about the item                                                             |
| `unitPrice`       | google.type.Money         | Price per unit of the item (before any discounts or adjustments)                              |
| `quantity`        | int32                     | Number of units purchased (must be greater than 0)                                            |
| `createdTime`     | google.protobuf.Timestamp | When this line item was created (system-generated timestamp)                                  |
| `lastUpdatedTime` | google.protobuf.Timestamp | When this line item was last modified (system-generated timestamp)                            |
| `paidAmount`      | google.type.Money         | Amount that has been paid for this item (updated when payments are processed)                 |

### 3. OrderDiscount

Represents a discount applied to an order or order line item.

#### OrderDiscount.DiscountType (enum)

| Value                       | Description                           |
|-----------------------------|---------------------------------------|
| `DISCOUNT_TYPE_UNSPECIFIED` | Unknown or invalid discount type      |
| `AMOUNT`                    | Fixed amount discount (e.g., $10 off) |
| `PERCENTAGE`                | Percentage discount (e.g., 15% off)   |
| `CREDIT`                    | Credit discount (e.g., store credit)  |

#### OrderDiscount.ApplyType (enum)

| Value                    | Description                      |
|--------------------------|----------------------------------|
| `APPLY_TYPE_UNSPECIFIED` | Unknown or invalid apply type    |
| `ALL`                    | Applied to all items             |
| `ALL_NO_TIP`             | Applied to all items except tips |
| `SERVICE`                | Applied to service items         |
| `PRODUCT`                | Applied to product items         |
| `PACKAGE`                | Applied to package items         |
| `SERVICE_CHARGE`         | Applied to service charge items  |
| `ITEM`                   | Applied to specific item         |
| `NONE`                   | No application                   |

| Field Name        | Type                      | Description                                                                                                      |
|-------------------|---------------------------|------------------------------------------------------------------------------------------------------------------|
| `id`              | string                    | Unique identifier for the discount record, obfuscated ID string (format: "od_" + random characters)              |
| `businessId`      | string                    | ID of the business where the discount is applied, obfuscated ID string                                           |
| `orderId`         | string                    | ID of the order associated with this discount, obfuscated ID string                                              |
| `orderItemId`     | string                    | ID of the order line item this discount is applied to; empty string means the discount is applied at order level |
| `applyType`       | ApplyType                 | Type of discount application                                                                                     |
| `deleted`         | bool                      | Flag indicating if this discount record is deleted                                                               |
| `discountType`    | DiscountType              | Type of discount                                                                                                 |
| `discountAmount`  | google.type.Money         | Discount amount applied                                                                                          |
| `discountRate`    | double                    | Discount rate applied (e.g., 0.5 for 50%)                                                                        |
| `applyBy`         | string                    | ID of the entity that applied the discount, obfuscated ID string                                                 |
| `applySequence`   | int32                     | Sequence in which the discount is applied                                                                        |
| `createdTime`     | google.protobuf.Timestamp | When this discount record was created (system-generated timestamp)                                               |
| `lastUpdatedTime` | google.protobuf.Timestamp | When this discount record was last updated (system-generated timestamp)                                          |
| `discountCode`    | string                    | The discount code used, obfuscated ID string                                                                     |

### 4. OrderLineTax

Represents tax information for an individual line item within an order.

#### OrderLineTax.ApplyType (enum)

| Value                    | Description                      |
|--------------------------|----------------------------------|
| `APPLY_TYPE_UNSPECIFIED` | Unknown or invalid apply type    |
| `ALL`                    | Applied to all items             |
| `ALL_NO_TIP`             | Applied to all items except tips |
| `SERVICE`                | Applied to service items         |
| `PRODUCT`                | Applied to product items         |
| `PACKAGE`                | Applied to package items         |
| `SERVICE_CHARGE`         | Applied to service charge items  |
| `ITEM`                   | Applied to specific item         |
| `NONE`                   | No application                   |

| Field Name        | Type                      | Description                                                                                          |
|-------------------|---------------------------|------------------------------------------------------------------------------------------------------|
| `id`              | string                    | Unique identifier for the line tax record, obfuscated ID string (format: "olt_" + random characters) |
| `businessId`      | string                    | ID of the business where the tax is applied, obfuscated ID string                                    |
| `orderId`         | string                    | ID of the order associated with this tax, obfuscated ID string                                       |
| `orderItemId`     | string                    | ID of the order line item this tax is applied to, obfuscated ID string                               |
| `applyType`       | ApplyType                 | Type of tax application                                                                              |
| `deleted`         | bool                      | Flag indicating if this tax record is deleted                                                        |
| `taxId`           | string                    | ID of the tax configuration, obfuscated ID string                                                    |
| `taxRate`         | double                    | Tax rate (e.g., 0.08 for 8%)                                                                         |
| `taxAmount`       | google.type.Money         | Actual tax amount applied                                                                            |
| `applyBy`         | string                    | ID of the entity that applied the tax, obfuscated ID string                                          |
| `applySequence`   | int32                     | Sequence in which the tax is applied                                                                 |
| `createdTime`     | google.protobuf.Timestamp | When this tax record was created (system-generated timestamp)                                        |
| `lastUpdatedTime` | google.protobuf.Timestamp | When this tax record was last updated (system-generated timestamp)                                   |
| `taxName`         | string                    | Name of the tax                                                                                      |

### 5. OrderRedeemPackage

Represents a service within a package that has been applied to an order.

| Field Name         | Type                      | Description                                                                                                |
|--------------------|---------------------------|------------------------------------------------------------------------------------------------------------|
| `id`               | string                    | Unique identifier for the package service record, obfuscated ID string (format: "ps_" + random characters) |
| `orderId`          | string                    | ID of the order this package service is associated with, obfuscated ID string                              |
| `orderItemId`      | string                    | ID of the order line item this package service is associated with, obfuscated ID string                    |
| `packageId`        | string                    | ID of the package, obfuscated ID string                                                                    |
| `serviceId`        | string                    | ID of the service within the package, obfuscated ID string                                                 |
| `servicePrice`     | google.type.Money         | Price of the service                                                                                       |
| `packageServiceId` | string                    | ID of the package service, obfuscated ID string                                                            |
| `packageName`      | string                    | Name of the package                                                                                        |
| `serviceName`      | string                    | Name of the service                                                                                        |
| `quantity`         | int32                     | Quantity of the service                                                                                    |
| `createdTime`      | google.protobuf.Timestamp | When this package service record was created (system-generated timestamp)                                  |
| `lastUpdatedTime`  | google.protobuf.Timestamp | When this package service record was last updated (system-generated timestamp)                             |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Order API

Here is a typical integration flow:

1. **GetOrder**
    - Retrieve detailed information about a specific order using its unique ID.
    - Useful for verifying order status, financial details, and associated line items.

2. **ListOrders**
    - List orders based on specified criteria (e.g., company ID, business locations, status).
    - Use pagination to manage large datasets efficiently.
    - Filter orders by status or last updated time if needed.

3. **ListOrderLineItems**
    - Retrieve detailed information about individual items within orders.
    - Useful for itemized reporting and revenue analysis.
    - Filter line items by order IDs if needed.
    - Optionally include tax information for each line item.

4. **ListOrderDiscounts**
    - Retrieve detailed information about discounts applied to orders or order line items.
    - Useful for discount reporting and analysis.
    - Filter discounts by order IDs if needed.

5. **ListOrderRedeemPackages**
    - Retrieve detailed information about packages that have been redeemed and applied to specific orders.
    - Useful for tracking package usage and reporting.
    - Filter redeemed packages by order IDs if needed.

6. **Monitoring & Reconciliation**
    - Regularly retrieve order data to monitor transaction statuses.
    - Check financial amounts and ensure accuracy.
    - Track payments and refunds for reconciliation purposes.

---

## 📦 5. API Interface Descriptions

### 1. GetOrder (`GetOrder`)

- **Method**: `GetOrder`
- **HTTP Method**: GET
- **Path**: `/v1/orders/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific order by its ID.

#### 🎯 Use Cases:

- Verify order status and details.
- Audit individual transactions.
- Track order lifecycle.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                                                      |
|------------|--------|----------|------------------------------------------------------------------|
| `id`       | string | Yes      | Unique identifier of the order to retrieve, obfuscated ID string |

#### 📌 Return Value:

| Field Name        | Type                      | Description                                                                                            |
|-------------------|---------------------------|--------------------------------------------------------------------------------------------------------|
| `id`              | string                    | Unique identifier for the order, obfuscated ID string (format: "ord_" + random characters)             |
| `businessId`      | string                    | ID of the business location where services were provided, obfuscated ID string                         |
| `customerId`      | string                    | ID of the customer who received services, obfuscated ID string                                         |
| `status`          | Status                    | Current status of the order                                                                            |
| `title`           | string                    | Short description of the order (e.g., "Full Grooming Package - Max")                                   |
| `description`     | string                    | Detailed notes about the order                                                                         |
| `tipsAmount`      | google.type.Money         | Additional amount provided as gratuity (optional, non-negative)                                        |
| `taxAmount`       | google.type.Money         | Tax amount applied to the order (calculated based on local tax rates)                                  |
| `discountAmount`  | google.type.Money         | Total discounts applied to the order (sum of all applicable discounts)                                 |
| `extraFeeAmount`  | google.type.Money         | Additional fees applied to the order (e.g., holiday surcharge)                                         |
| `subTotalAmount`  | google.type.Money         | Subtotal before tax, tips, and adjustments (sum of all service prices)                                 |
| `tipsBasedAmount` | google.type.Money         | Amount used as basis for calculating tips (usually equals subTotalAmount)                              |
| `totalAmount`     | google.type.Money         | Total amount including all charges and adjustments (formula: subtotal + tax + tips + fees - discounts) |
| `paidAmount`      | google.type.Money         | Amount that has been paid so far (updated when payments are processed)                                 |
| `remainAmount`    | google.type.Money         | Amount still due on the order (formula: totalAmount - paidAmount)                                      |
| `refundedAmount`  | google.type.Money         | Amount that has been refunded (tracks any refunds issued to the customer)                              |
| `createdBy`       | string                    | ID of the staff member who created the order, obfuscated ID string                                     |
| `createdTime`     | google.protobuf.Timestamp | When this order was created (system-generated timestamp)                                               |
| `lastUpdatedBy`   | string                    | ID of the staff member who last modified the order, obfuscated ID string                               |
| `lastUpdatedTime` | google.protobuf.Timestamp | When this order was last modified (system-generated timestamp)                                         |
| `salesDatetime`   | google.protobuf.Timestamp | When the sale was recorded (used for financial reporting and reconciliation)                           |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified order ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. ListOrders (`ListOrders`)

- **Method**: `ListOrders`
- **HTTP Method**: POST
- **Path**: `/v1/orders:list`

#### ✅ Functionality:

Retrieves a paginated list of orders based on specified criteria. Supports filtering by business locations, status, and
update time.

#### 🎯 Use Cases:

- View all orders for a company.
- Filter orders by business location or status.
- Monitor order activity over time.

#### 🔧 Request Parameters:

| Field Name                 | Type                | Required | Description                                                                                    |
|----------------------------|---------------------|----------|------------------------------------------------------------------------------------------------|
| `pagination`               | Pagination          | Yes      | Pagination info: pageSize, pageToken                                                           |
| `companyId`                | string              | Yes      | Company ID to scope orders, obfuscated ID string                                               |
| `businessIds`              | Array(string)       | Yes      | List of business locations to filter orders by, obfuscated ID strings                          |
| `filter.ids`               | Array(string)       | No       | Specific order IDs to retrieve (if provided, other filters are ignored), obfuscated ID strings |
| `filter.statuses`          | Array(Order.Status) | No       | Order statuses to include in results                                                           |
| `filter.last_updated_time` | Interval            | No       | Time range for filtering orders by last update time                                            |

#### 📌 Return Value:

| Field Name      | Type         | Description                                                          |
|-----------------|--------------|----------------------------------------------------------------------|
| `nextPageToken` | string       | Token for retrieving the next page of results (empty if none remain) |
| `orders`        | Array(Order) | List of orders matching the request criteria                         |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid.
- `PERMISSION_DENIED`: Permission denied.

---

### 3. ListOrderLineItems (`ListOrderLineItems`)

- **Method**: `ListOrderLineItems`
- **HTTP Method**: POST
- **Path**: `/v1/orders/items:list`

#### ✅ Functionality:

Retrieves detailed information about individual items within orders. This method is useful for itemized reporting and
revenue analysis. When the `include_tax` parameter is set to true, the response will also include tax details for each
line item.

#### 🎯 Use Cases:

- Analyze revenue by service type or product.
- Track usage of specific services or products.
- Generate detailed financial reports.
- Retrieve tax information for line items when needed.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description                                                               |
|--------------------|---------------|----------|---------------------------------------------------------------------------|
| `companyId`        | string        | Yes      | Company ID to scope line items                                            |
| `includeTax`       | bool          | No       | Include tax information in the response (default: false)                  |
| `filter.order_ids` | Array(string) | No       | Optional list of order IDs to filter line items by, obfuscated ID strings |

#### 📌 Return Value:

| Field Name           | Type                 | Description                                                             |
|----------------------|----------------------|-------------------------------------------------------------------------|
| `orderLineItems`     | Array(OrderLineItem) | List of line items matching the criteria                                |
| `orderLineItemTaxes` | Array(OrderLineTax)  | List of line item taxes matching the criteria (when includeTax is true) |

#### ⚠️ Error Codes:

- `PERMISSION_DENIED`: Permission denied.

---

### 4. ListOrderDiscounts (`ListOrderDiscounts`)

- **Method**: `ListOrderDiscounts`
- **HTTP Method**: POST
- **Path**: `/v1/orders/discounts:list`

#### ✅ Functionality:

Retrieves detailed information about discounts applied to orders or order line items. This method is useful for discount
reporting and analysis.

#### 🎯 Use Cases:

- Analyze discount usage and effectiveness.
- Track discount application across orders.
- Generate discount-related financial reports.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description                                                              |
|--------------------|---------------|----------|--------------------------------------------------------------------------|
| `companyId`        | string        | Yes      | Company ID to scope discounts                                            |
| `filter.order_ids` | Array(string) | No       | Optional list of order IDs to filter discounts by, obfuscated ID strings |

#### 📌 Return Value:

| Field Name       | Type                 | Description                             |
|------------------|----------------------|-----------------------------------------|
| `orderDiscounts` | Array(OrderDiscount) | List of discounts matching the criteria |

#### ⚠️ Error Codes:

- `PERMISSION_DENIED`: Permission denied.

---

### 5. ListOrderRedeemPackages (`ListOrderRedeemPackages`)

- **Method**: `ListOrderRedeemPackages`
- **HTTP Method**: POST
- **Path**: `/v1/orders/redeemPackages:list`

#### ✅ Functionality:

Retrieves detailed information about packages that have been redeemed and applied to specific orders. This method is
useful for tracking package usage and reporting.

#### 🎯 Use Cases:

- Track package redemption and usage.
- Analyze package popularity and effectiveness.
- Generate package-related reports.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description                                                                      |
|--------------------|---------------|----------|----------------------------------------------------------------------------------|
| `companyId`        | string        | Yes      | Company ID to scope redeemed packages                                            |
| `filter.order_ids` | Array(string) | No       | Optional list of order IDs to filter redeemed packages by, obfuscated ID strings |

#### 📌 Return Value:

| Field Name            | Type                      | Description                                     |
|-----------------------|---------------------------|-------------------------------------------------|
| `orderRedeemPackages` | Array(OrderRedeemPackage) | List of redeemed packages matching the criteria |

#### ⚠️ Error Codes:

- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: GetOrder

```json
{
  "id": "ord_001"
}
```

### Example 2: ListOrders

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "businessIds": [
    "biz_001",
    "biz_002"
  ],
  "filter": {
    "statuses": [
      "CREATED",
      "PROCESSING"
    ],
    "lastUpdatedTime": {
      "startTime": "2024-08-01T00:00:00Z",
      "endTime": "2024-08-02T00:00:00Z"
    }
  }
}

```

### Example 3: ListOrderLineItems

```json
{
  "companyId": "cmp_001",
  "includeTax": true,
  "filter": {
    "orderIds": [
      "ord_001",
      "ord_002"
    ]
  }
}
```

### Example 4: ListOrderDiscounts

```json
{
  "companyId": "cmp_001",
  "filter": {
    "orderIds": [
      "ord_001",
      "ord_002"
    ]
  }
}
```

### Example 5: ListOrderRedeemPackages

```json
{
  "companyId": "cmp_001",
  "filter": {
    "orderIds": [
      "ord_001",
      "ord_002"
    ]
  }
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                                | Answer                                                                                   |
|---------------------------------------------------------|------------------------------------------------------------------------------------------|
| How to verify if an order exists?                       | Use `GetOrder` to check if the order ID returns a valid response                         |
| Can I list orders for multiple companies at once?       | Currently only supports listing orders for one company at a time                         |
| How to filter orders by business location?              | Use `ListOrders` with `businessIds`                                                      |
| Why does creating an order return "resource exhausted"? | Not applicable — orders are typically not created via this API                           |
| How to handle incomplete orders?                        | Use `ListOrders` with `STATUS_PROCESSING` to identify and follow up on incomplete orders |
| How to get tax information for line items?              | Use `ListOrderLineItems` with `includeTax` set to true                                   |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Order ID does not exist           |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/pagination.md)
- [order_service.proto](../moego/business/order/v1/order_service.proto)
- [order.proto](../moego/business/order/v1/order.proto)
- [order_line_item.proto](../moego/business/order/v1/order_line_item.proto)
- [order_discount.proto](../moego/business/order/v1/order_discount.proto)
- [order_line_tax.proto](../moego/business/order/v1/order_line_tax.proto)