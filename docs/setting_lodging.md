# 🛠️ Setting Lodging API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The lodging settings provide configuration options for boarding services, including lodging types and units. These
settings help standardize boarding accommodations and ensure consistent service delivery across all business locations.

---

## 🎯 2. Design Goals

- **Centralized Configuration**: Provides a unified way to manage lodging configurations.
- **Rich Data Model**: Supports complex relationships like lodging types and units.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Configuring boarding accommodations.
- Managing different types of lodging facilities.
- Standardizing boarding services across multiple locations.

---

## 🧩 3. Core Concepts

### 1. Lodging

A lodging represents a type of accommodation and its associated units, typically used for boarding services.

| Field Name     | Type               | Description                              |
|----------------|--------------------|------------------------------------------|
| `lodgingType`  | LodgingType        | The type of lodging (e.g., room, area)   |
| `lodgingUnits` | Array(LodgingUnit) | List of individual units within the type |

### 2. LodgingType

Describes the type of lodging available.

| Field Name        | Type            | Description                            |
|-------------------|-----------------|----------------------------------------|
| `id`              | string          | Unique identifier for the lodging type |
| `name`            | string          | Name of the lodging type               |
| `description`     | string          | Optional description                   |
| `photoList`       | Array(string)   | URLs to photos of this lodging type    |
| `maxPetNum`       | int32           | Maximum number of pets allowed         |
| `lodgingUnitType` | LodgingUnitType | Type of lodging unit                   |

### 3. LodgingUnit

Describes a specific unit within a lodging type.

| Field Name | Type   | Description                |
|------------|--------|----------------------------|
| `id`       | string | Unique identifier for unit |
| `name`     | string | Display name of the unit   |

### 4. LodgingUnitType

Enum representing the type of lodging unit.

| Value                           | Description                       |
|---------------------------------|-----------------------------------|
| `LODGING_UNIT_TYPE_UNSPECIFIED` | Default value; should not be used |
| `ROOM`                          | Room/kennel type                  |
| `AREA`                          | Open area for multiple pets       |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Lodging Setting API

Here is a typical integration flow:

1. **List Lodgings**
    - Retrieve all available lodging configurations.

2. **Manage Lodging Types and Units**
    - Use lodging types and units for display or booking integration.

---

## 📦 5. API Interface Descriptions

### 1. List Lodgings (`ListLodgings`)

- **Method**: `ListLodgings`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/lodgings`

#### ✅ Functionality:

Lists all available lodgings for a company.

Lodgings represent different types of accommodation options for pet boarding services.

#### 🎯 Use Cases:

- Retrieve lodging types and units for booking or display
- Standardize lodging offerings across business locations

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type             | Description                    |
|------------|------------------|--------------------------------|
| `lodgings` | Array(`Lodging`) | List of lodging configurations |

#### ⚠️ Error Codes:

| Error Code          | Description                |
|---------------------|----------------------------|
| `PERMISSION_DENIED` | Permission denied          |
| `INVALID_ARGUMENT`  | Malformed request          |
| `NOT_FOUND`         | The company does not exist |

---

## 🧪 6. Usage Examples

### Example 1: List Lodgings

Request Body:

```json
{
  "companyId": "cmp_001"
}
```

Response Body:

```json
{
  "lodgings": [
    {
      "lodgingType": {
        "id": "lt_001",
        "name": "Deluxe Room",
        "description": "Spacious room with premium amenities.",
        "photoList": [
          "https://example.com/photo1.jpg"
        ],
        "maxPetNum": 2,
        "lodgingUnitType": "ROOM"
      },
      "lodgingUnits": [
        {
          "id": "lu_001",
          "name": "Room 101"
        },
        {
          "id": "lu_002",
          "name": "Room 102"
        }
      ]
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                                   | Answer                                                                                     |
|------------------------------------------------------------|--------------------------------------------------------------------------------------------|
| How can I manage boarding accommodations effectively?      | Use `ListLodgings` to retrieve lodging types and units for consistent boarding management. |
| What should I do if a request returns "permission denied"? | Verify that your API key has the necessary permissions to access lodging settings.         |

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

- [lodging.proto](../moego/business/setting/v1/lodging.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)