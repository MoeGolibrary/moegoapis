# 🛠️ Setting Pet API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The pet settings provide configuration options related to pet management, including pet codes, pet sizes, pet types, and
pet coat types. These settings help standardize pet data entry and ensure consistent pet service delivery across all
business locations.

---

## 🎯 2. Design Goals

- **Centralized Configuration**: Provides a unified way to manage pet-related settings like codes and categories.
- **Rich Data Model**: Supports complex relationships like pet categorization by size, type, and coat.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Flagging special handling requirements or medical conditions for pets.
- Categorizing pets by size, type, or coat for service customization.
- Standardizing pet handling across multiple locations.

---

## 🧩 3. Core Concepts

### 1. PetCode

Represents special handling instructions or medical alerts for a pet.

| Field Name     | Type   | Description                          |
|----------------|--------|--------------------------------------|
| `id`           | string | Unique identifier                    |
| `abbreviation` | string | Short form (e.g., AG for Aggressive) |
| `description`  | string | Detailed explanation                 |
| `color`        | string | Highlight color in UI                |

### 2. PetSize

Represents a size category for pets, which can be used to customize services based on the pet's size.

| Field Name   | Type   | Description                                    |
|--------------|--------|------------------------------------------------|
| `id`         | string | Unique identifier for this pet size            |
| `name`       | string | Name of the pet size (e.g., "Small", "Medium") |
| `weightLow`  | Weight | Lower weight limit for this size category      |
| `weightHigh` | Weight | Upper weight limit for this size category      |

### 3. PetType

Represents a type or species of pet, such as Dog, Cat, Bird, etc.

| Field Name | Type   | Description          |
|------------|--------|----------------------|
| `id`       | string | Unique identifier    |
| `name`     | string | Name of the pet type |

### 4. PetCoatType

Represents a coat type or fur type for pets, such as Short, Long, Curly, etc.

| Field Name | Type   | Description               |
|------------|--------|---------------------------|
| `id`       | string | Unique identifier         |
| `name`     | string | Name of the pet coat type |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Pet Setting API

Here is a typical integration flow:

1. **List Pet Codes**
    - Retrieve all available pet codes for special handling requirements.

2. **List Pet Sizes**
    - Retrieve all available pet sizes for service customization.

3. **List Pet Types**
    - Retrieve all available pet types for categorization.

4. **List Pet Coat Types**
    - Retrieve all available pet coat types for service customization.

---

## 📦 5. API Interface Descriptions

### 1. List Pet Codes (`ListPetCodes`)

- **Method**: `ListPetCodes`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/pet/codes:list`

#### ✅ Functionality:

Lists all available pet codes for a company.

Pet codes are used to flag special handling requirements or medical conditions that staff need to be aware of when
providing services. These codes are displayed prominently in pet profiles and during service delivery.

#### 🎯 Use Cases:

- Retrieve special handling instructions
- Standardize pet care alerts across locations

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type             | Description                       |
|------------|------------------|-----------------------------------|
| `codes`    | Array(`PetCode`) | List of pet codes for the company |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 2. List Pet Sizes (`ListPetSizes`)

- **Method**: `ListPetSizes`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/pet/sizes:list`

#### ✅ Functionality:

Lists all available pet sizes for a company.

Pet sizes are used to categorize pets by general size (e.g., Small, Medium, Large) which can be used to customize
services, pricing, or handling requirements.

#### 🎯 Use Cases:

- Retrieve pet sizes for service customization
- Standardize pet categorization across locations

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type             | Description                       |
|------------|------------------|-----------------------------------|
| `petSizes` | Array(`PetSize`) | List of pet sizes for the company |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 3. List Pet Types (`ListPetTypes`)

- **Method**: `ListPetTypes`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/pet/types:list`

#### ✅ Functionality:

Lists all available pet types for a company.

Pet types are used to categorize pets by species (e.g., Dog, Cat, Bird) which can be used to customize services,
pricing, or handling requirements.

#### 🎯 Use Cases:

- Retrieve pet types for service customization
- Standardize pet categorization across locations

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type             | Description                       |
|------------|------------------|-----------------------------------|
| `petTypes` | Array(`PetType`) | List of pet types for the company |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 4. List Pet Coat Types (`ListPetCoatTypes`)

- **Method**: `ListPetCoatTypes`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/pet/coat_types:list`

#### ✅ Functionality:

Lists all available pet coat types for a company.

Pet coat types are used to categorize pets by their fur characteristics (e.g., Short, Long, Curly) which can be used to
customize services, pricing, or handling requirements.

#### 🎯 Use Cases:

- Retrieve pet coat types for service customization
- Standardize pet categorization across locations

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name     | Type                 | Description                            |
|----------------|----------------------|----------------------------------------|
| `petCoatTypes` | Array(`PetCoatType`) | List of pet coat types for the company |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

## 🧪 6. Usage Examples

### Example 1: List Pet Codes

Request Body:

```json
{
  "companyId": "cmp_001"
}
```

Response Body:

```json
{
  "codes": [
    {
      "id": "code_001",
      "abbreviation": "AG",
      "description": "Aggressive pet requiring special handling",
      "color": "#FF0000"
    },
    {
      "id": "code_002",
      "abbreviation": "MED",
      "description": "Pet with medical conditions",
      "color": "#0000FF"
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                                   | Answer                                                                                   |
|------------------------------------------------------------|------------------------------------------------------------------------------------------|
| How can I flag special handling requirements for pets?     | Use `ListPetCodes` to retrieve codes and apply them to pets for special handling alerts. |
| How can I categorize pets by size, type, or coat?          | Use `ListPetSizes`, `ListPetTypes`, and `ListPetCoatTypes` for pet categorization.       |
| What should I do if a request returns "permission denied"? | Verify that your API key has the necessary permissions to access pet settings.           |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| `PERMISSION_DENIED` | Current user has no access rights to perform the operation.                 |
| `INVALID_ARGUMENT`  | Invalid request parameters (e.g., missing required fields, invalid format). |
| `INTERNAL`          | Internal server error occurred while processing the request.                |

---

## 📎 10. Related File References

- [pet.proto](../moego/business/setting/v1/pet.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)