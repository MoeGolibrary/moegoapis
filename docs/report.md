# Report API Documentation (`moego.business.report.v1`)

## 📌 1. Functional Overview

Report represents a customizable business report definition. This module provides the following functions:

- Getting all available report definitions for a company
- Generating and retrieving report data based on specified conditions (supports pagination, time range filtering, field
  grouping)
- Supporting multiple data types (money, text, date, time, number, percentage, etc.)

Applicable to scenarios such as business analysis, decision-making, financial statistics, etc.

---

## 🎯 2. Design Goals

- **Flexible Configuration**: Each report can be customized with fields and formatting to adapt to different business
  needs
- **Efficient Querying**: Supports pagination, time range filtering, and field grouping aggregation
- **Unified Structure**: The returned data structure is clear and standardized, making it easy to integrate and process
- **Multiple Data Type Support**: Covers common business data types to meet diverse reporting requirements

---

## 🧩 3. Core Concepts

### 1. Report

Represents a configurable business report definition, containing field structure and description information.

| Field Name    | Type         | Description                                                           |
|---------------|--------------|-----------------------------------------------------------------------|
| `id`          | string       | Unique identifier for the report (format: "rpt_" + random characters) |
| `name`        | string       | Display name of the report, used in UI and exported documents         |
| `description` | string       | Detailed description of the report's purpose and contents             |
| `fields`      | Array(Field) | List of fields included in the report, defining the report structure  |

### 2. Field

Defines a single column of data in a report, including data type and display properties.

#### Field.Type (enum)

| Value              | Description                                                                                                                                                      |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `TYPE_UNSPECIFIED` | Default value when type is not specified; should not be used when creating new fields                                                                            |
| `MONEY`            | Represents monetary values with currency information, stored using `google.type.Money`                                                                           |
| `TEXT`             | Represents text or string values                                                                                                                                 |
| `DATE`             | Represents date-only value (yyyy-MM-dd), based on the local timezone; no conversion needed                                                                       |
| `TIME`             | Represents full datetime value (yyyy-MM-dd hh:mm) but display as time-only (hh:mm), based on UTC+0 timezone; convert according to the company's timezone setting |
| `NUMBER`           | Represents whole numbers                                                                                                                                         |
| `DECIMAL_NUMBER`   | Represents numbers with decimal points                                                                                                                           |
| `PERCENTAGE`       | Represents percentage values, stored as decimal (e.g., 0.75 for 75%)                                                                                             |
| `DURATION`         | Represents time duration in minutes                                                                                                                              |
| `DATETIME`         | Represents full datetime value (yyyy-MM-dd hh:mm), based on UTC+0 timezone; convert according to the company's timezone setting                                  |

| Field Name      | Type       | Description                                                                                          |
|-----------------|------------|------------------------------------------------------------------------------------------------------|
| `label`         | string     | Display label for the field, used in report headers and UI elements                                  |
| `type`          | Field.Type | Data type of the field, determining value formatting and validation rules                            |
| `key`           | string     | Unique identifier for the field within the report, used to reference field values in data structures |
| `groupByEnable` | bool       | Whether this field can be used as a grouping criterion, affecting report aggregation options         |

### 3. TableData

Represents a complete dataset for a report, containing both the structure (fields) and content (rows) of the report.

| Field Name | Type                | Description                                                                       |
|------------|---------------------|-----------------------------------------------------------------------------------|
| `rows`     | Array(TableRowData) | Array of data rows, each row contains values for all defined fields               |
| `fields`   | Array(Field)        | Field definitions for the table, can be from `TableMeta` or dynamically generated |

### 4. TableRowData

Represents a single row of data in a report, containing values for each field defined in the table structure.

| Field Name | Type              | Description                                                                                |
|------------|-------------------|--------------------------------------------------------------------------------------------|
| `data`     | Map(string, Data) | Map of field values indexed by field key, each entry corresponds to a column in the report |

### 5. Data

Represents a single cell value in a report table, including both the field metadata and the actual value.

| Field Name  | Type       | Description                                                                      |
|-------------|------------|----------------------------------------------------------------------------------|
| `fieldKey`  | string     | Reference to the field definition, must match a field key in the table structure |
| `fieldType` | Field.Type | Data type of the value, must match the field definition                          |
| `value`     | Value      | Actual value of the cell, type must match `fieldType`                            |

### 6. Value

Represents a typed data value used in reports, supporting multiple data types through a oneof field.

| Field Name  | Type                      | Description                              |
|-------------|---------------------------|------------------------------------------|
| `string`    | string                    | Text value for string fields             |
| `double`    | double                    | Decimal value for numerical fields       |
| `int64`     | int64                     | Integer value for whole number fields    |
| `bool`      | bool                      | Boolean value                            |
| `money`     | google.type.Money         | Monetary value with currency information |
| `timestamp` | google.protobuf.Timestamp | Date/time value for temporal fields      |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User calls report interface for data analysis

The typical usage flow is as follows:

1. **ListReports**
    - Get all report definitions available to the current company
    - Can be used to select which report to view or analyze

2. **FetchReportData**
    - Generate specific data based on the selected report ID and time range
    - Can specify grouping fields for data aggregation
    - Returns structured data that can be used for front-end display or export

---

## 📦 5. API Interface Descriptions

### 1. ListReports (`ListReports`)

- **Method**: `ListReports`
- **HTTP Method**: POST
- **Path**: `/v1/reports:list`

#### ✅ Functionality:

Gets the list of all report definitions available to the current company.

#### 🎯 Use Cases:

- View all reports available to the current company
- Used for users to select which report to analyze

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                             |
|-------------|--------|----------|-----------------------------------------|
| `companyId` | string | Yes      | Company ID, used for permission control |

#### 📌 Return Value:

| Field Name | Type          | Description                          |
|------------|---------------|--------------------------------------|
| `reports`  | Array(Report) | List of available report definitions |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: companyId is invalid
- `PERMISSION_DENIED`: Permission denied

---

### 2. FetchReportData (`FetchReportData`)

- **Method**: `FetchReportData`
- **HTTP Method**: POST
- **Path**: `/v1/reports/data:fetch`

#### ✅ Functionality:

Generates and returns report data based on the specified report ID and parameters.

#### 🎯 Use Cases:

- Generate specific data for a given report
- Supports pagination, time range filtering, and custom field grouping

#### 🔧 Request Parameters:

| Field Name                   | Type       | Required | Description                                                                |
|------------------------------|------------|----------|----------------------------------------------------------------------------|
| `pagination`                 | Pagination | Yes      | Pagination parameters (pageSize, pageToken)                                |
| `companyId`                  | string     | Yes      | Company ID, used for permission control                                    |
| `businessIds`                | string[]   | No       | List of business unit IDs to include in the report (if null, includes all) |
| `condition.id`               | string     | Yes      | Report definition ID                                                       |
| `condition.queryPeriod`      | Interval   | Yes      | Time period for data collection                                            |
| `condition.groupByFieldKeys` | string[]   | No       | Optional list of field keys to group data by                               |

#### 📌 Return Value:

| Field Name      | Type      | Description                                      |
|-----------------|-----------|--------------------------------------------------|
| `nextPageToken` | string    | Token for retrieving the next page of results    |
| `tableData`     | TableData | Structured report data including fields and rows |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Request parameters are invalid
- `PERMISSION_DENIED`: Permission denied
- `NOT_FOUND`: Report definition does not exist

---

## 🧪 6. Usage Examples

### Example 1: ListReports

```json
{
  "companyId": "cmp_001"
}
```

### Example 2: FetchReportData

```json
{
  "pagination": {
    "pageSize": 20
  },
  "companyId": "cmp_001",
  "condition": {
    "id": "rpt_sales_summary",
    "queryPeriod": {
      "startTime": "2024-09-01T00:00:00Z",
      "endTime": "2024-09-30T23:59:59Z"
    },
    "groupByFieldKeys": [
      "product",
      "region"
    ]
  }
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                                 | Answer                                                                                           |
|----------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| How to determine if a report exists?                     | Use the `ListReports` interface to check if the report with the corresponding ID exists          |
| Is it possible to get data for multiple reports at once? | Currently only supports querying a single report, need to call `FetchReportData` multiple times  |
| How to group report data?                                | Pass the `groupByFieldKeys` parameter in the `FetchReportData` request                           |
| What data types are supported for report fields?         | Includes money, text, date, time, integer, decimal, percentage, duration, and complete date-time |
| How to implement paginated queries?                      | Use `pagination.pageSize` and `pagination.pageToken` to implement pagination                     |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                    |
|---------------------|------------------------------------------------|
| `NOT_FOUND`         | Report definition or company ID does not exist |
| `PERMISSION_DENIED` | Current user has no access rights              |
| `INVALID_ARGUMENT`  | Request parameters are invalid                 |
| `INTERNAL`          | Internal server error                          |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/address.md)
- [report_service.proto](../moego/business/report/v1/report_service.proto)
- [report.proto](../moego/business/report/v1/report.proto)