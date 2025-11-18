# Agreement API Documentation (`moego.business.agreement.v1`)

## 📌 1. Functional Overview

The `Agreement` object represents a legal document that requires customer consent. It is used to manage various types of
agreements such as service terms, privacy policies, and waivers. Each agreement can be configured with different signing
requirements and notification templates for customer communication.

This API provides operations for:

- Retrieving individual agreements
- Listing agreements with filtering options
- Generating secure links for customers to sign agreements
- Tracking customer interactions through `AgreementRecord`

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all types of agreements.
- **Flexible Signing Policies**: Supports different policies for when signatures are required (first interaction, each
  transaction, or optional).
- **Audit Trail**: Maintains timestamps for tracking when agreements were created, edited, or last required for signing.
- **Customer Communication**: Includes templates for SMS and email notifications related to the agreement.
- **Signing Record Tracking**: Tracks how and when customers interacted with specific agreements via `AgreementRecord`.

---

## 🧩 3. Core Concepts

### 1. Agreement

Represents a legal document requiring customer consent. A single agreement can be reused across multiple customers and
transactions.

| Field Name           | Type               | Description                                                       |
|----------------------|--------------------|-------------------------------------------------------------------|
| `id`                 | string             | Unique identifier of the agreement                                |
| `businessId`         | string             | Business location where this agreement is used                    |
| `creatorId`          | string             | Staff member who created the agreement                            |
| `status`             | enum(Status)       | Current state of the agreement: `NORMAL`, `DELETED`               |
| `signedPolicy`       | enum(SignedPolicy) | When signatures are required: `FOR_FIRST`, `FOR_EACH`, `OPTIONAL` |
| `title`              | string             | Display name of the agreement                                     |
| `content`            | string             | Full text of the agreement                                        |
| `smsTemplate`        | string             | Template for SMS notifications                                    |
| `emailTemplateTitle` | string             | Subject line for email notifications                              |
| `emailTemplateBody`  | string             | Body content for email notifications                              |
| `lastRequiredTime`   | timestamp          | Customers who signed before this time may need to re-sign         |
| `lastEditTime`       | timestamp          | When the agreement content was last modified                      |
| `createdTime`        | timestamp          | When the agreement was first created                              |
| `lastUpdatedTime`    | timestamp          | When any field was last changed                                   |

#### Enum Definitions

##### `Agreement.Status`

- `STATUS_UNSPECIFIED`
- `NORMAL`: Agreement is active and can be presented to customers
- `DELETED`: Agreement has been removed and is no longer in use

##### `Agreement.SignedPolicy`

- `SIGNED_POLICY_UNSPECIFIED`
- `FOR_FIRST`: Customer must sign only on their first interaction
- `FOR_EACH`: Customer must sign for each relevant transaction
- `OPTIONAL`: Customer may proceed without signing

---

### 2. AgreementRecord

Tracks individual instances of customer interactions with agreements. It maintains the history of when and how customers
signed agreements, including the specific version they saw and their signature method.

| Field Name     | Type               | Description                                                           |
|----------------|--------------------|-----------------------------------------------------------------------|
| `id`           | string             | Unique identifier of the record                                       |
| `uuid`         | string             | External reference UUID                                               |
| `agreementId`  | string             | Reference to the original agreement                                   |
| `businessId`   | string             | Business location where signing occurred                              |
| `companyId`    | string             | Parent company identifier                                             |
| `customerId`   | string             | Customer who signed or viewed the agreement                           |
| `targetId`     | string             | Related object ID (e.g., appointment, form)                           |
| `status`       | enum(Status)       | Current state of the record: `NORMAL`, `DELETED`                      |
| `signedStatus` | enum(SignedStatus) | Whether the agreement has been signed: `UNSIGNED`, `SIGNED`           |
| `signedType`   | enum(SignedType)   | How the agreement was signed: `CUSTOMER_SIGNED`, `BY_BUSINESS_UPLOAD` |
| `sourceType`   | enum(SourceType)   | Where the agreement was presented: `URL`, `MOBILE`, etc.              |
| `link`         | string             | URL where the agreement can be viewed                                 |
| `title`        | string             | Agreement title at time of signing                                    |
| `content`      | string             | Agreement content at time of signing                                  |
| `signature`    | string             | Customer's signature data                                             |
| `signedTime`   | timestamp          | When the agreement was signed                                         |
| `createdTime`  | timestamp          | When this record was created                                          |
| `updatedTime`  | timestamp          | When this record was last modified                                    |

#### Enum Definitions

##### `AgreementRecord.Status`

- `STATUS_UNSPECIFIED`
- `NORMAL`: Record is active and valid
- `DELETED`: Record has been removed

##### `AgreementRecord.SignedStatus`

- `SIGNED_STATUS_UNSPECIFIED`
- `UNSIGNED`: Agreement has been presented but not yet signed
- `SIGNED`: Customer has completed the signing process

##### `AgreementRecord.SignedType`

- `SIGNED_TYPE_UNSPECIFIED`
- `CUSTOMER_SIGNED`: Customer provided an electronic signature
- `BY_BUSINESS_UPLOAD`: Business uploaded a physically signed document

##### `AgreementRecord.SourceType`

- `SOURCE_TYPE_UNSPECIFIED`
- `URL`: Signed through a web browser interface
- `MOBILE`: Signed through the mobile application
- `ONLINE_BOOKING`: Signed during online booking process
- `INTAKE_FORM`: Signed as part of intake form completion

---

## 📦 4. API Interface Descriptions

### 1. Get Agreement (`GetAgreement`)

- **Method**: `GetAgreement`
- **HTTP Method**: GET
- **Path**: `/v1/agreements/{id}`

#### ✅ Functionality:

Retrieves a specific agreement by its ID.

#### 🎯 Use Cases:

- View details of an existing agreement.
- Verify agreement configuration during debugging.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                   |
|-------------|--------|----------|-------------------------------|
| `id`        | string | Yes      | Agreement ID to retrieve      |
| `companyId` | string | Yes      | Company ID for access control |

#### 📌 Return Value:

| Field Name           | Type               | Description                                                       |
|----------------------|--------------------|-------------------------------------------------------------------|
| `id`                 | string             | Unique identifier of the agreement                                |
| `businessId`         | string             | Business location where this agreement is used                    |
| `creatorId`          | string             | Staff member who created the agreement                            |
| `status`             | enum(Status)       | Current state of the agreement: `NORMAL`, `DELETED`               |
| `signedPolicy`       | enum(SignedPolicy) | When signatures are required: `FOR_FIRST`, `FOR_EACH`, `OPTIONAL` |
| `title`              | string             | Display name of the agreement                                     |
| `content`            | string             | Full text of the agreement                                        |
| `smsTemplate`        | string             | Template for SMS notifications                                    |
| `emailTemplateTitle` | string             | Subject line for email notifications                              |
| `emailTemplateBody`  | string             | Body content for email notifications                              |
| `lastRequiredTime`   | timestamp          | Customers who signed before this time may need to re-sign         |
| `lastEditTime`       | timestamp          | When the agreement content was last modified                      |
| `createdTime`        | timestamp          | When the agreement was first created                              |
| `lastUpdatedTime`    | timestamp          | When any field was last changed                                   |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified agreement ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. List Agreements (`ListAgreements`)

- **Method**: `ListAgreements`
- **HTTP Method**: POST
- **Path**: `/v1/agreements:list`

#### ✅ Functionality:

Lists agreements matching specified criteria, including company ID and optional business IDs filter.

#### 🎯 Use Cases:

- View all agreements under a company.
- Audit or debug agreement configurations.

#### 🔧 Request Parameters:

| Field Name    | Type          | Required | Description                                           |
|---------------|---------------|----------|-------------------------------------------------------|
| `pagination`  | Pagination    | Yes      | Page size and token                                   |
| `companyId`   | string        | Yes      | Company ID for access control                         |
| `businessIds` | Array(string) | No       | Optional list of business IDs to filter agreements by |

> **Note**: The `pagination` field is used for pagination.
> The `pageSize` field specifies the number of results to return per page. Maximum value is 500.
> The `pageToken` field is used to retrieve the next page of results.

#### 📌 Return Value:

| Field Name      | Type             | Description                                   |
|-----------------|------------------|-----------------------------------------------|
| `nextPageToken` | string           | Token for retrieving the next page of results |
| `agreement`     | Array(Agreement) | List of agreements matching the criteria      |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 3. Generate Agreement Sign Link (`GetAgreementSignLink`)

- **Method**: `GetAgreementSignLink`
- **HTTP Method**: GET
- **Path**: `/v1/agreements/{id}/sign_link`

#### ✅ Functionality:

Generates a unique URL where a customer can view and sign the agreement.

#### 🎯 Use Cases:

- Provide a direct link for a customer to sign an agreement.
- Integrate with external systems for automated workflows.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description                          |
|--------------|--------|----------|--------------------------------------|
| `id`         | string | Yes      | Agreement ID to generate link for    |
| `customerId` | string | Yes      | Customer who will sign the agreement |
| `businessId` | string | Yes      | Business context for the agreement   |

#### 📌 Return Value:

| Field Name          | Type   | Description                                          |
|---------------------|--------|------------------------------------------------------|
| `agreementRecordId` | string | Identifier of the generated agreement record         |
| `signUrl`           | string | Unique URL where the customer can sign the agreement |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified agreement ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 5. Usage Examples

### Example 1: Get Agreement

```http
GET /v1/agreements/12345?company_id=cmp_001
```

**Response:**

```json
{
  "id": "12345",
  "businessId": "biz_001",
  "creatorId": "staff_001",
  "status": "NORMAL",
  "signedPolicy": "FOR_FIRST",
  "title": "Privacy Policy v2",
  "content": "This Privacy Policy explains how we collect...",
  "smsTemplate": "Please review and sign the latest privacy policy.",
  "emailTemplateTitle": "Action Required: New Privacy Policy",
  "emailTemplateBody": "Dear customer, please review and sign our updated privacy policy...",
  "lastRequiredTime": "2024-08-01T00:00:00Z",
  "lastEditTime": "2024-07-20T10:00:00Z",
  "createdTime": "2024-06-15T09:00:00Z",
  "lastUpdatedTime": "2024-08-01T14:30:00Z"
}
```

### Example 2: List Agreements

```json
{
  "companyId": "cmp_001",
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "businessIds": [
    "biz_001",
    "biz_002"
  ]
}
```

### Example 3: Generate Sign Link

```http
GET /v1/agreements/12345/sign_link?customer_id=cus_001&business_id=biz_001
```

**Response:**

```json
{
  "agreementRecordId": "record_001",
  "signUrl": "https://example.com/agreements/sign/abcxyz"
}
```

---

## 📌 6. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Agreement ID does not exist       |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 7. Related File References

- [agreement_service.proto](../moego/business/agreement/v1/agreement_service.proto)
- [agreement.proto](../moego/business/agreement/v1/agreement.proto)
- [agreement_record.proto](../moego/business/agreement/v1/agreement_record.proto)

---