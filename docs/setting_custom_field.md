# 🛠️ Setting Custom Field API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The custom field settings provide flexible configuration options for extending customer and lead data models with business-specific fields. These settings allow businesses to capture additional information beyond standard fields, enabling customized data collection and management tailored to specific business needs.

---

## 🎯 2. Design Goals

- **Flexible Data Model**: Supports multiple field types including text, numbers, dates, selections, and relations.
- **Type-Safe Configuration**: Provides strongly-typed field definitions with validation rules.
- **Multi-Entity Support**: Can be associated with different entity types (customers, leads).
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Capturing business-specific customer information.
- Extending lead qualification data.
- Customizing data collection forms.
- Standardizing additional data fields across multiple locations.

---

## 🧩 3. Core Concepts

### 1. CustomField

A custom field represents an additional data field that can be associated with customers or leads.

| Field Name         | Type              | Description                                                            |
|--------------------|-------------------|------------------------------------------------------------------------|
| `id`               | string            | Unique identifier for the custom field                                 |
| `companyId`        | string            | ID of the company that owns this field                                 |
| `associationType`  | AssociationType   | Entity type this field is associated with (customer/lead)              |
| `code`             | string            | Unique key used when creating/updating customer/lead (e.g., field_123) |
| `label`            | string            | Display name shown to users                                            |
| `type`             | Type              | Data type of the field value                                           |
| `isRequired`       | bool              | Whether this field must be filled                                      |
| `defaultValue`     | Value             | Default value for the field                                            |
| `options`          | Array(Option)     | Available options for select/multi-select fields                       |
| `source`           | Source            | Origin of the field (custom or system)                                 |
| `displayOrder`     | int32             | Order in which the field should be displayed                           |
| `helpText`         | string            | Additional guidance text for users                                     |
| `createTime`       | Timestamp         | When this field was created                                            |
| `updateTime`       | Timestamp         | When this field was last updated                                       |

### 2. CustomField.Type

Enum representing the data type of a custom field.

| Value              | Description                                      | Value Type         |
|--------------------|--------------------------------------------------|--------------------|
| `TYPE_UNSPECIFIED` | Default value; should not be used                | -                  |
| `SHORT_TEXT`       | Short text input                                 | string             |
| `LONG_TEXT`        | Long text input (multi-line)                     | string             |
| `NUMBER`           | Numeric value                                    | int64              |
| `DATE`             | Date value                                       | Timestamp          |
| `BOOLEAN`          | True/false value                                 | bool               |
| `SELECT`           | Single selection from options                    | string             |
| `MULTI_SELECT`     | Multiple selections from options                 | StringList         |
| `RELATION`         | Reference to another entity                      | Relation           |
| `MONEY`            | Monetary value with currency                     | Money              |
| `TIME`             | Time of day                                      | TimeOfDay          |
| `DATETIME`         | Date and time value                              | Timestamp          |
| `GROUP`            | Grouping field for organization                  | -                  |

### 3. CustomField.AssociationType

Enum representing the entity type a custom field can be associated with.

| Value                          | Description                                |
|--------------------------------|--------------------------------------------|
| `ASSOCIATION_TYPE_UNSPECIFIED` | Default value; should not be used          |
| `CUSTOMER`                     | Field is associated with customer entities |
| `LEAD`                         | Field is associated with lead entities     |

### 4. CustomField.Source

Enum representing the origin of a custom field.

| Value                | Description                                    |
|----------------------|------------------------------------------------|
| `SOURCE_UNSPECIFIED` | Default value; should not be used              |
| `CUSTOM`             | Custom field created by the company            |
| `SYSTEM`             | Built-in system field                          |

### 5. CustomField.Value

Represents the value of a custom field, supporting multiple data types.

| Field Name      | Type          | Description                                |
|-----------------|---------------|--------------------------------------------|
| `string`        | string        | String value for text fields               |
| `doubleValue`   | double        | Floating-point value                       |
| `int64`         | int64         | Integer value for number fields            |
| `bool`          | bool          | Boolean value                              |
| `money`         | Money         | Monetary value with currency               |
| `timestampTime` | Timestamp     | Date or datetime value                     |
| `relation`      | Relation      | Reference to another entity                |
| `stringList`    | StringList    | List of strings for multi-select           |
| `timeOfDay`     | TimeOfDay     | Time of day value                          |

### 6. CustomField.Option

Represents an option for select or multi-select fields.

| Field Name  | Type   | Description                        |
|-------------|--------|------------------------------------|
| `value`     | Value  | The actual value of the option     |
| `label`     | string | Display text for the option        |
| `sortOrder` | int32  | Order in which option is displayed |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Custom Field API

Here is a typical integration flow:

1. **List Custom Fields**
    - Retrieve all custom fields for a specific entity type (customer or lead).
    - Filter by association type to get relevant fields.

2. **Display Custom Fields in Forms**
    - Use field metadata (label, type, options) to render appropriate input controls.
    - Apply validation rules based on `isRequired` and field type.

3. **Submit Data with Custom Fields**
    - Use the `code` field as the key when submitting custom field values.
    - Ensure values match the expected type for each field.

---

## 📦 5. API Interface Descriptions

### 1. List Custom Fields (`ListCustomFields`)

- **Method**: `ListCustomFields`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/custom_fields`

#### ✅ Functionality:

Lists all available custom fields for a company, with optional filtering by association type.

Custom fields allow businesses to extend standard data models with additional fields specific to their needs.

#### 🎯 Use Cases:

- Retrieve custom fields to render in customer or lead forms
- Understand what additional data can be collected
- Build dynamic forms based on company configuration

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                                  |
|-------------|--------|----------|----------------------------------------------|
| `companyId` | string | Yes      | Company ID to retrieve custom fields for     |
| `filter`    | Filter | No       | Optional filter to narrow results            |

**Filter Parameters:**

| Field Name         | Type                      | Required | Description                                    |
|--------------------|---------------------------|----------|------------------------------------------------|
| `associationTypes` | Array(AssociationType)    | No       | Filter by entity types (customer, lead, etc.)  |

#### 📌 Return Value:

| Field Name     | Type                  | Description                      |
|----------------|-----------------------|----------------------------------|
| `customFields` | Array(`CustomField`)  | List of custom field definitions |

#### ⚠️ Error Codes:

| Error Code          | Description                |
|---------------------|----------------------------|
| `PERMISSION_DENIED` | Permission denied          |
| `INVALID_ARGUMENT`  | Malformed request          |
| `NOT_FOUND`         | The company does not exist |

---

## 🧪 6. Usage Examples

### Example 1: List All Custom Fields

Request Body:

```json
{
  "companyId": "cmp_001"
}
```

Response Body:

```json
{
  "customFields": [
    {
      "id": "cf_001",
      "companyId": "cmp_001",
      "associationType": "CUSTOMER",
      "code": "field_preferred_groomer",
      "label": "Preferred Groomer",
      "type": "SHORT_TEXT",
      "isRequired": false,
      "displayOrder": 1,
      "helpText": "Enter the name of the customer's preferred groomer",
      "source": "CUSTOM",
      "createTime": "2024-01-15T10:00:00Z",
      "updateTime": "2024-01-15T10:00:00Z"
    },
    {
      "id": "cf_002",
      "companyId": "cmp_001",
      "associationType": "CUSTOMER",
      "code": "field_membership_tier",
      "label": "Membership Tier",
      "type": "SELECT",
      "isRequired": true,
      "options": [
        {
          "value": {
            "string": "bronze"
          },
          "label": "Bronze",
          "sortOrder": 1
        },
        {
          "value": {
            "string": "silver"
          },
          "label": "Silver",
          "sortOrder": 2
        },
        {
          "value": {
            "string": "gold"
          },
          "label": "Gold",
          "sortOrder": 3
        }
      ],
      "displayOrder": 2,
      "source": "CUSTOM",
      "createTime": "2024-01-15T10:00:00Z",
      "updateTime": "2024-01-15T10:00:00Z"
    }
  ]
}
```

### Example 2: List Custom Fields for Leads Only

Request Body:

```json
{
  "companyId": "copKSMj",
  "filter": {
    "associationTypes": ["LEAD"]
  }
}
```

Response Body:

```json
{
  "customFields": [
    {
      "id": "cf_001",
      "companyId": "cmp_001",
      "associationType": "LEAD",
      "code": "field_preferred_groomer",
      "label": "Preferred Groomer",
      "type": "SHORT_TEXT",
      "isRequired": false,
      "displayOrder": 1,
      "source": "CUSTOM",
      "createTime": "2024-01-15T10:00:00Z",
      "updateTime": "2024-01-15T10:00:00Z"
    }
  ]
}
```

### Example 3: Custom Field with Multiple Types

Response showing different field types:

```json
{
  "customFields": [
    {
      "id": "cf_003",
      "code": "field_birthday",
      "label": "Pet Birthday",
      "type": "DATE",
      "associationType": "CUSTOMER"
    },
    {
      "id": "cf_004",
      "code": "field_special_needs",
      "label": "Special Needs",
      "type": "LONG_TEXT",
      "associationType": "CUSTOMER"
    },
    {
      "id": "cf_005",
      "code": "field_vip_customer",
      "label": "VIP Customer",
      "type": "BOOLEAN",
      "associationType": "CUSTOMER",
      "defaultValue": {
        "bool": false
      }
    },
    {
      "id": "cf_006",
      "code": "field_account_balance",
      "label": "Account Balance",
      "type": "MONEY",
      "associationType": "CUSTOMER",
      "defaultValue": {
        "money": {
          "currencyCode": "USD",
          "units": "0"
        }
      }
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                                      | Answer                                                                                                                |
|---------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| How can I add custom fields to customer or lead forms?        | Use `ListCustomFields` with appropriate filters to retrieve field definitions, then render them in your forms.       |
| What is the `code` field used for?                           | The `code` field is a unique key (e.g., `field_123`) used when submitting custom field values during entity creation or updates. |
| Can I filter custom fields by entity type?                   | Yes, use the `filter.associationTypes` parameter to retrieve only customer or lead fields.                           |
| What field types are supported?                               | Supports text, numbers, dates, booleans, selections, relations, money, time, and datetime fields.                    |
| What should I do if a request returns "permission denied"?    | Verify that your API key has the necessary permissions to access custom field settings.                               |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| `PERMISSION_DENIED` | Current user has no access rights to perform the operation.                 |
| `INVALID_ARGUMENT`  | Invalid request parameters (e.g., missing required fields, invalid format). |
| `NOT_FOUND`         | The requested company does not exist.                                       |
| `INTERNAL`          | Internal server error occurred while processing the request.                |

---

## 📎 10. Related File References

- [customer.proto](../moego/business/setting/v1/customer.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)
