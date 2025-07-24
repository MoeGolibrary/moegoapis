# Payment API Documentation (`moego.business.payment.v1`)

## 📌 1. Functional Overview

Payment represents a financial transaction that processes funds for an order. This module provides the following
functions:

- Retrieving detailed payment information by ID
- Listing payments with support for filtering by order IDs
- Tracking payment status, processing fees, and refund amounts
- Supporting multiple payment methods and business modules
- Managing payment lifecycle from creation to completion

Applicable to scenarios such as payment tracking, financial reconciliation, refund processing, and transaction
reporting.

---

## 🎯 2. Design Goals

- **Centralized Payment Management**: Provides a unified interface for managing all aspects of payment transactions
- **Flexible Filtering**: Supports filtering by order IDs to customize result sets
- **Comprehensive Payment Tracking**: Maintains detailed records of payment status, fees, and refunds
- **Multi-Module Support**: Handles payments for different business areas (grooming, retail, membership)
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

---

## 🧩 3. Core Concepts

### 1. Payment

Represents a financial transaction that processes funds for an order. Each payment tracks the amount, processing status,
and associated fees.

#### Payment.Status (enum)

| Value                | Description                                                    |
|----------------------|----------------------------------------------------------------|
| `STATUS_UNSPECIFIED` | Unknown or invalid status; should not be used for new payments |
| `CREATED`            | Payment has been initiated but not processed                   |
| `PROCESSING`         | Payment is being processed by the payment provider             |
| `PAID`               | Funds have been successfully transferred                       |
| `COMPLETED`          | Payment has been fully reconciled                              |
| `FAILED`             | Payment processing encountered an error                        |

#### Payment.Module (enum)

| Value                | Description                                                    |
|----------------------|----------------------------------------------------------------|
| `MODULE_UNSPECIFIED` | Unknown or invalid module; should not be used for new payments |
| `GROOMING`           | Payment for pet grooming services                              |
| `RETAIL`             | Payment for retail product purchases                           |
| `MEMBERSHIP`         | Payment for membership fees or subscriptions                   |

| Field Name          | Type                      | Description                                                      |
|---------------------|---------------------------|------------------------------------------------------------------|
| `id`                | string                    | Unique identifier for the payment (format: "py_" + random chars) |
| `business_id`       | string                    | ID of the business location processing the payment               |
| `order_id`          | string                    | ID of the order being paid for                                   |
| `customer_id`       | string                    | ID of the customer making the payment                            |
| `customer_name`     | string                    | Display name of the customer                                     |
| `method`            | string                    | Payment method used (e.g., "credit_card", "cash", "check")       |
| `status`            | Status                    | Current status of the payment                                    |
| `module`            | Module                    | Business module associated with the payment                      |
| `amount`            | google.type.Money         | Amount being processed                                           |
| `processing_fee`    | google.type.Money         | Fees charged by payment processor                                |
| `refund_amount`     | google.type.Money         | Amount that has been refunded                                    |
| `created_time`      | google.protobuf.Timestamp | When this payment was created                                    |
| `last_updated_time` | google.protobuf.Timestamp | When this payment was last modified                              |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User integrates and debugs the Payment API

Here is a typical integration flow:

1. **GetPayment**
    - Retrieve detailed information about a specific payment using its unique ID.
    - Useful for verifying payment status, amounts, and associated order/customer.

2. **ListPayments**
    - List payments based on specified criteria (e.g., company ID, order IDs).
    - Use pagination to manage large datasets efficiently.
    - Filter payments by order IDs if needed.

3. **Monitoring & Reconciliation**
    - Regularly retrieve payment data to monitor transaction statuses.
    - Check payment amounts and fees for accuracy.
    - Track refunds for reconciliation purposes.

---

## 📦 5. API Interface Descriptions

### 1. GetPayment (`GetPayment`)

- **Method**: `GetPayment`
- **HTTP Method**: GET
- **Path**: `/v1/payment/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific payment by its ID.

#### 🎯 Use Cases:

- Verify payment status and details.
- Audit individual transactions.
- Track payment lifecycle.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                                  |
|------------|--------|----------|----------------------------------------------|
| `id`       | string | Yes      | Unique identifier of the payment to retrieve |

#### 📌 Return Value:

| Field Name          | Type                      | Description                                                      |
|---------------------|---------------------------|------------------------------------------------------------------|
| `id`                | string                    | Unique identifier for the payment (format: "py_" + random chars) |
| `business_id`       | string                    | ID of the business location processing the payment               |
| `order_id`          | string                    | ID of the order being paid for                                   |
| `customer_id`       | string                    | ID of the customer making the payment                            |
| `customer_name`     | string                    | Display name of the customer                                     |
| `method`            | string                    | Payment method used (e.g., "credit_card", "cash", "check")       |
| `status`            | Status                    | Current status of the payment                                    |
| `module`            | Module                    | Business module associated with the payment                      |
| `amount`            | google.type.Money         | Amount being processed                                           |
| `processing_fee`    | google.type.Money         | Fees charged by payment processor                                |
| `refund_amount`     | google.type.Money         | Amount that has been refunded                                    |
| `created_time`      | google.protobuf.Timestamp | When this payment was created                                    |
| `last_updated_time` | google.protobuf.Timestamp | When this payment was last modified                              |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified payment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. ListPayments (`ListPayments`)

- **Method**: `ListPayments`
- **HTTP Method**: POST
- **Path**: `/v1/payments:list`

#### ✅ Functionality:

Retrieves a paginated list of payments based on specified criteria. Supports filtering by order IDs.

#### 🎯 Use Cases:

- View all payments for a company.
- Filter payments by associated order IDs.
- Monitor payment activity over time.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description                                      |
|--------------------|---------------|----------|--------------------------------------------------|
| `pagination`       | Pagination    | Yes      | Pagination info: page_size, page_token           |
| `company_id`       | string        | Yes      | Company ID to scope payments                     |
| `filter.order_ids` | Array(string) | No       | Optional list of order IDs to filter payments by |

#### 📌 Return Value:

| Field Name        | Type           | Description                                                          |
|-------------------|----------------|----------------------------------------------------------------------|
| `next_page_token` | string         | Token for retrieving the next page of results (empty if none remain) |
| `payments`        | Array(Payment) | List of payments matching the request criteria                       |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Pagination parameters are invalid.
- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: GetPayment

```json
{
  "id": "py_001"
}
```

### Example 2: ListPayments

```json
{
  "pagination": {
    "page_size": 20
  },
  "company_id": "cmp_001",
  "filter": {
    "order_ids": [
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

| Question                                                 | Answer                                                                                   |
|----------------------------------------------------------|------------------------------------------------------------------------------------------|
| How to verify if a payment exists?                       | Use `GetPayment` to check if the payment ID returns a valid response                     |
| Can I list payments for multiple companies at once?      | Currently only supports listing payments for one company at a time                       |
| How to filter payments by order?                         | Use `ListPayments` with `filter.order_ids`                                               |
| Why does creating a payment return "resource exhausted"? | Not applicable — payments are typically not created via this API                         |
| How to handle failed payments?                           | Use `GetPayment` to check status and determine next steps based on business requirements |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Payment ID does not exist         |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/pagination.md)
- [payment_service.proto](../moego/business/payment/v1/payment_service.proto)
- [payment.proto](../moego/business/payment/v1/payment.proto)