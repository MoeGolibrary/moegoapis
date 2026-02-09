# Pet API Documentation (`moego.business.customer.v1`)

## 📌 1. Functional Overview

Pet represents an animal companion owned by a customer, which receives services at your business. This interface
enables:

- Creating, updating, and deleting pet profiles
- Managing pet-related information such as health records, behavior notes, vaccination status, and special handling
  codes
- Listing all pets belonging to a specific customer or across the entire company
- Adding and managing special handling codes and behavioral notes for individual pets
- Retrieving detailed pet data including service history, veterinary information, and evaluation status

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all aspects of pet data
- **Rich Data Model**: Supports complex relationships like vaccinations, weight, veterinary info, and behavioral notes
- **Secure and Reliable**: Ensures data integrity and access control
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks

Applicable to scenarios such as pet grooming, medical care tracking, behavioral monitoring, and customer pet management.

---

## 🧩 3. Core Concepts

### 1. Pet

Represents a customer's pet that receives services at your business. Each pet has its own profile containing essential
information for providing appropriate care.

| Field Name         | Type                                         | Description                                        |
|--------------------|----------------------------------------------|----------------------------------------------------|
| `id`               | string                                       | Unique identifier (obfuscated string)              |
| `name`             | string                                       | Pet's given name                                   |
| `birthday`         | Date                                         | Pet's date of birth                                |
| `status`           | Status                                       | Current status (ALIVE/PASSED_AWAY)                 |
| `type`             | Type                                         | Species type (DOG/CAT/BIRD etc.)                   |
| `breed`            | string                                       | Specific breed within the pet type                 |
| `gender`           | Gender                                       | Gender of the pet                                  |
| `weight`           | Weight                                       | Pet's current weight                               |
| `fixed`            | string                                       | Spay/neuter status ("yes", "no", "unknown")        |
| `coat`             | string                                       | Description of the coat (e.g., "long double coat") |
| `behavior`         | string                                       | General temperament and behavior                   |
| `petCodes`         | Array([PetCode](./setting_pet.md#1-petcode)) | Special handling instructions or medical alerts    |
| `notes`            | Array(Note)                                  | Staff observations and special instructions        |
| `vaccinations`     | Array(Vaccination)                           | Vaccination records                                |
| `customerId`       | string                                       | ID of the pet's owner (obfuscated string)          |
| `vet`              | Vet                                          | Primary veterinary care provider                   |
| `evaluationStatus` | EvaluationStatus                             | Service eligibility status                         |
| `createdTime`      | Timestamp                                    | When the pet was created                           |
| `photo`            | string                                       | Photo URL of the pet                               |
| `deleted`          | bool                                         | Flag indicating if this pet record is deleted      |

### 2. Note

Contains staff observations about the pet's behavior, preferences, or special requirements.

| Field Name        | Type      | Description                      |
|-------------------|-----------|----------------------------------|
| `id`              | string    | Unique identifier                |
| `content`         | string    | The content of the note          |
| `lastUpdatedBy`   | string    | Staff member who last updated it |
| `lastUpdatedTime` | Timestamp | Last modification timestamp      |

### 3. Vaccination

Tracks vaccination records for compliance with service requirements.

| Field Name  | Type      | Description                     |
|-------------|-----------|---------------------------------|
| `name`      | string    | Vaccination name (e.g., Rabies) |
| `expiredAt` | Timestamp | Expiration date                 |

### 4. Vet

Represents veterinary care provider information for emergencies.

| Field Name    | Type   | Description                   |
|---------------|--------|-------------------------------|
| `name`        | string | Veterinary practice or doctor |
| `phoneNumber` | string | Contact number (E.164 format) |
| `address`     | string | Physical location             |

### 5. Weight

Represents pet weight measurements.

see: [Weight](./common/weight.md)

### 6. Enum Definitions

#### Pet.Type

- `TYPE_UNSPECIFIED`
- `OTHER`
- `DOG`
- `CAT`
- `BIRD`
- `RABBIT`
- `GUINEA_PIG`
- `HORSE`
- `HAMSTER`
- `RAT`
- `MOUSE`
- `CHINCHILLA`

#### Pet.Status

- `STATUS_UNSPECIFIED`
- `ALIVE`
- `PASSED_AWAY`

#### Pet.Gender

- `GENDER_UNSPECIFIED`
- `MALE`
- `FEMALE`
- `UNKNOWN`

#### Pet.EvaluationStatus

- `EVALUATION_STATUS_UNSPECIFIED`
- `PASS`
- `FAIL`

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Pet API

Here is a typical integration flow:

1. **Create Pet**
    - Specify required details like name, species, breed, and owner customer ID.
    - Optionally set birthday, gender, weight, coat, behavior, and initial notes/codes.

2. **Update Pet**
    - Modify pet details like name, weight, coat, behavior, or vet info.
    - Add or remove notes and codes.

3. **Retrieve Pet**
    - Get full details of an existing pet, including vaccination records, notes, and evaluation status.

4. **List Pets**
    - View all pets belonging to a specific customer.
    - Filter results if needed.

5. **Manage Notes & Codes**
    - Append new notes or codes to a pet's profile.
    - Retrieve lists of notes and codes associated with a pet.

6. **List All Pets**
    - View all pets across customers in the company (requires elevated permissions).

7. **Monitoring & Maintenance**
    - Regularly retrieve pet data to monitor changes.
    - Update pet records as needed.

---

## 📦 5. API Interface Descriptions

### 1. Create Pet (`CreatePet`)

- **Method**: `CreatePet`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{customer_id}/pets`

#### ✅ Functionality:

Registers a new pet with basic details, health records, and optionally initial notes or codes.

#### 🎯 Use Cases:

- Users want to onboard a new pet into the system.
- Third-party systems sync pet data.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description                           |
|--------------|--------|----------|---------------------------------------|
| `customerId` | string | Yes      | Owner customer ID (obfuscated string) |
| `pet`        | Pet    | Yes      | Complete pet information to create    |

#### 📌 Return Value:

| Field Name         | Type                                         | Description                                        |
|--------------------|----------------------------------------------|----------------------------------------------------|
| `id`               | string                                       | Unique identifier (obfuscated string)              |
| `name`             | string                                       | Pet's given name                                   |
| `birthday`         | Date                                         | Pet's date of birth                                |
| `status`           | Status                                       | Current status (ALIVE/PASSED_AWAY)                 |
| `type`             | Type                                         | Species type (DOG/CAT/BIRD etc.)                   |
| `breed`            | string                                       | Specific breed within the pet type                 |
| `gender`           | Gender                                       | Gender of the pet                                  |
| `weight`           | Weight                                       | Pet's current weight                               |
| `fixed`            | string                                       | Spay/neuter status ("yes", "no", "unknown")        |
| `coat`             | string                                       | Description of the coat (e.g., "long double coat") |
| `behavior`         | string                                       | General temperament and behavior                   |
| `petCodes`         | Array([PetCode](./setting_pet.md#1-petcode)) | Special handling instructions or medical alerts    |
| `notes`            | Array(Note)                                  | Staff observations and special instructions        |
| `vaccinations`     | Array(Vaccination)                           | Vaccination records                                |
| `customerId`       | string                                       | ID of the pet's owner (obfuscated string)          |
| `vet`              | Vet                                          | Primary veterinary care provider                   |
| `evaluationStatus` | EvaluationStatus                             | Service eligibility status                         |
| `createdTime`      | Timestamp                                    | When the pet was created                           |
| `photo`            | string                                       | Photo URL of the pet                               |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields are missing or invalid.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. Get Pet (`GetPet`)

- **Method**: `GetPet`
- **HTTP Method**: GET
- **Path**: `/v1/customers/{customer_id}/pets/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific pet.

#### 🎯 Use Cases:

- Check current pet data.
- Verify pet details during debugging.
- Monitor pet activity.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description                            |
|--------------|--------|----------|----------------------------------------|
| `customerId` | string | Yes      | Owner customer ID (obfuscated string)  |
| `id`         | string | Yes      | Pet ID to retrieve (obfuscated string) |

#### 📌 Return Value:

| Field Name         | Type                                         | Description                                        |
|--------------------|----------------------------------------------|----------------------------------------------------|
| `id`               | string                                       | Unique identifier (obfuscated string)              |
| `name`             | string                                       | Pet's given name                                   |
| `birthday`         | Date                                         | Pet's date of birth                                |
| `status`           | Status                                       | Current status (ALIVE/PASSED_AWAY)                 |
| `type`             | Type                                         | Species type (DOG/CAT/BIRD etc.)                   |
| `breed`            | string                                       | Specific breed within the pet type                 |
| `gender`           | Gender                                       | Gender of the pet                                  |
| `weight`           | Weight                                       | Pet's current weight                               |
| `fixed`            | string                                       | Spay/neuter status ("yes", "no", "unknown")        |
| `coat`             | string                                       | Description of the coat (e.g., "long double coat") |
| `behavior`         | string                                       | General temperament and behavior                   |
| `petCodes`         | Array([PetCode](./setting_pet.md#1-petcode)) | Special handling instructions or medical alerts    |
| `notes`            | Array(Note)                                  | Staff observations and special instructions        |
| `vaccinations`     | Array(Vaccination)                           | Vaccination records                                |
| `customerId`       | string                                       | ID of the pet's owner (obfuscated string)          |
| `vet`              | Vet                                          | Primary veterinary care provider                   |
| `evaluationStatus` | EvaluationStatus                             | Service eligibility status                         |
| `createdTime`      | Timestamp                                    | When the pet was created                           |
| `photo`            | string                                       | Photo URL of the pet                               |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified pet ID or customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 3. Update Pet (`UpdatePet`)

- **Method**: `UpdatePet`
- **HTTP Method**: PUT
- **Path**: `/v1/customers/{customer_id}/pets/{id}`

#### ✅ Functionality:

Updates an existing pet's information.

#### 🎯 Use Cases:

- Change pet details like name, weight, coat, or behavior.
- Update vaccination records or vet info.
- Modify evaluation status.

#### 🔧 Request Parameters:

| Field Name         | Type               | Required | Description                                                |
|--------------------|--------------------|----------|------------------------------------------------------------|
| `customerId`       | string             | Yes      | Owner customer ID (obfuscated string)                      |
| `id`               | string             | Yes      | Unique identifier of the pet to update (obfuscated string) |
| `name`             | string             | No       | Pet's name                                                 |
| `birthday`         | Date               | No       | Pet's date of birth                                        |
| `type`             | Pet.Type           | No       | Pet's species                                              |
| `breed`            | string             | No       | Pet's breed                                                |
| `gender`           | Pet.Gender         | No       | Pet's gender                                               |
| `weight`           | Weight             | No       | Pet's weight                                               |
| `fixed`            | string             | No       | Spay/neuter status                                         |
| `coat`             | string             | No       | Coat description                                           |
| `behavior`         | string             | No       | Behavior notes                                             |
| `vet`              | Pet.Vet            | No       | Veterinary care provider                                   |
| `evaluationStatus` | EvaluationStatus   | No       | Service eligibility status                                 |
| `photo`            | string             | No       | Pet photo URL                                              |
| `vaccinations`     | Array(Vaccination) | No       | Vaccination records                                        |

#### 📌 Return Value:

| Field Name         | Type                                         | Description                                        |
|--------------------|----------------------------------------------|----------------------------------------------------|
| `id`               | string                                       | Unique identifier (obfuscated string)              |
| `name`             | string                                       | Pet's given name                                   |
| `birthday`         | Date                                         | Pet's date of birth                                |
| `status`           | Status                                       | Current status (ALIVE/PASSED_AWAY)                 |
| `type`             | Type                                         | Species type (DOG/CAT/BIRD etc.)                   |
| `breed`            | string                                       | Specific breed within the pet type                 |
| `gender`           | Gender                                       | Gender of the pet                                  |
| `weight`           | Weight                                       | Pet's current weight                               |
| `fixed`            | string                                       | Spay/neuter status ("yes", "no", "unknown")        |
| `coat`             | string                                       | Description of the coat (e.g., "long double coat") |
| `behavior`         | string                                       | General temperament and behavior                   |
| `petCodes`         | Array([PetCode](./setting_pet.md#1-petcode)) | Special handling instructions or medical alerts    |
| `notes`            | Array(Note)                                  | Staff observations and special instructions        |
| `vaccinations`     | Array(Vaccination)                           | Vaccination records                                |
| `customerId`       | string                                       | ID of the pet's owner (obfuscated string)          |
| `vet`              | Vet                                          | Primary veterinary care provider                   |
| `evaluationStatus` | EvaluationStatus                             | Service eligibility status                         |
| `createdTime`      | Timestamp                                    | When the pet was created                           |
| `photo`            | string                                       | Photo URL of the pet                               |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified pet ID or customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 4. List Pets (`ListPets`)

- **Method**: `ListPets`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{customer_id}/pets:list`

#### ✅ Functionality:

Lists all pets belonging to a specific customer.

#### 🎯 Use Cases:

- View all pets under a customer.
- Audit or debug pet configurations.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description                                      |
|--------------|--------|----------|--------------------------------------------------|
| `customerId` | string | Yes      | Customer ID to list pets for (obfuscated string) |

#### 📌 Return Value:

| Field Name | Type       | Description                            |
|------------|------------|----------------------------------------|
| `pets`     | Array(Pet) | List of pets belonging to the customer |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 5. Append Pet Notes (`AppendPetNotes`)

- **Method**: `AppendPetNotes`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{customer_id}/pets/{id}/notes`

#### ✅ Functionality:

Adds new notes to a pet's profile.

#### 🎯 Use Cases:

- Record pet behavior or preferences.
- Track service history or special requirements.

#### 🔧 Request Parameters:

| Field Name   | Type            | Required | Description                                  |
|--------------|-----------------|----------|----------------------------------------------|
| `customerId` | string          | Yes      | Owner customer ID (obfuscated string)        |
| `id`         | string          | Yes      | Pet ID to add notes to (obfuscated string)   |
| `notes`      | Array(Pet.Note) | Yes      | Notes to add to the pet  (obfuscated string) |

#### 📌 Return Value:

| Field Name | Type            | Description                        |
|------------|-----------------|------------------------------------|
| `notes`    | Array(Pet.Note) | Notes that were successfully added |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified pet ID or customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 6. List Pet Notes (`ListPetNotes`)

- **Method**: `ListPetNotes`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{customer_id}/pets/{id}/notes:list`

#### ✅ Functionality:

Retrieves a paginated list of notes for a specific pet.

#### 🎯 Use Cases:

- Review historical notes about a pet.
- Audit pet interactions or service history.

#### 🔧 Request Parameters:

| Field Name   | Type       | Required | Description                                      |
|--------------|------------|----------|--------------------------------------------------|
| `pagination` | Pagination | Yes      | Page size and token                              |
| `customerId` | string     | Yes      | Owner customer ID (obfuscated string)            |
| `id`         | string     | Yes      | Pet ID to retrieve notes for (obfuscated string) |

#### 📌 Return Value:

| Field Name      | Type            | Description                                   |
|-----------------|-----------------|-----------------------------------------------|
| `nextPageToken` | string          | Token for retrieving the next page of results |
| `notes`         | Array(Pet.Note) | List of notes for the pet                     |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified pet ID or customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 7. Append Pet Codes (`AppendPetCodes`)

- **Method**: `AppendPetCodes`
- **HTTP Method**: POST
- **Path**: `/v1/customers/{customer_id}/pets/{id}/codes`

#### ✅ Functionality:

Adds special handling codes or medical alerts to a pet's profile.

#### 🎯 Use Cases:

- Apply tags for aggressive behavior, allergies, or special care needs.
- Ensure staff awareness during service delivery.

#### 🔧 Request Parameters:

| Field Name   | Type          | Required | Description                                      |
|--------------|---------------|----------|--------------------------------------------------|
| `customerId` | string        | Yes      | Owner customer ID (obfuscated string)            |
| `id`         | string        | Yes      | Pet ID to add codes to (obfuscated string)       |
| `petCodeIds` | Array(string) | Yes      | Code IDs to apply to the pet (obfuscated string) |

#### 📌 Return Value:

| Field Name | Type                                         | Description                        |
|------------|----------------------------------------------|------------------------------------|
| `codes`    | Array([PetCode](./setting_pet.md#1-petcode)) | Codes that were successfully added |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified pet ID or customer ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 8. List All Pets (`ListAllPets`)

- **Method**: `ListAllPets`
- **HTTP Method**: POST
- **Path**: `/v1/pets:list`

#### ✅ Functionality:

Lists all pets across all customers in the company.

#### 🎯 Use Cases:

- View all pets in the system.
- Filter by customer IDs or time intervals.

#### 🔧 Request Parameters:

| Field Name               | Type          | Required | Description                                        |
|--------------------------|---------------|----------|----------------------------------------------------|
| `pagination`             | Pagination    | Yes      | Page size and token                                |
| `companyId`              | string        | Yes      | Company ID to list pets for (obfuscated string)    |
| `filter.customerIds`     | Array(string) | No       | Optional list of customer IDs (obfuscated strings) |
| `filter.lastUpdatedTime` | Interval      | No       | Time range filter                                  |

> **Note**:The `pagination` field is used for pagination.
> The `pageSize` field specifies the number of results to return per page. Maximum value is 500.
> The `pageToken` field is used to retrieve the next page of results.

#### 📌 Return Value:

| Field Name      | Type       | Description                                   |
|-----------------|------------|-----------------------------------------------|
| `nextPageToken` | string     | Token for retrieving the next page of results |
| `pets`          | Array(Pet) | List of pets matching the request criteria    |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: Create Pet

```json
{
  "customerId": "cus_001",
  "pet": {
    "name": "Buddy",
    "birthday": {
      "year": 2020,
      "month": 5,
      "day": 12
    },
    "type": "DOG",
    "breed": "Labrador Retriever",
    "gender": "MALE",
    "weight": {
      "value": 25.5,
      "unit": "kg"
    },
    "coat": "short hair",
    "behavior": "friendly",
    "notes": [
      {
        "content": "Loves playing fetch."
      }
    ],
    "vaccinations": [
      {
        "name": "Rabies",
        "expiredAt": "2025-05-12T00:00:00Z"
      }
    ],
    "vet": {
      "name": "Dr. Smith Animal Clinic",
      "phoneNumber": "+12125551234",
      "address": "123 Vet Street, Cityville"
    },
    "evaluationStatus": "PASS",
    "photo": "https://example.com/pets/buddy.jpg"
  }
}
```

### Example 2: Update Pet

```json
{
  "customerId": "cus_001",
  "id": "pet_001",
  "name": "Buddy",
  "coat": "medium length",
  "behavior": "loves playing fetch but cautious around strangers",
  "vet": {
    "name": "New Vet Practice",
    "phoneNumber": "+12125554321",
    "address": "456 New Vet Road, Cityville"
  }
}
```

### Example 3: List Pets

```json
{
  "customerId": "cus_001"
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## 📎 8. FAQ

| Question                                             | Answer                                                                                                                |
|------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| How to verify if a pet exists?                       | Use `GetPet` to check if the pet ID returns a valid response                                                          |
| Can I create multiple pets at once?                  | Currently only single pet creation is supported. Use batch processing if needed                                       |
| How to manage pet notes effectively?                 | Use `AppendPetNotes` to add new entries                                                                               |
| Why does creating a pet return "resource exhausted"? | The customer may have reached the maximum allowed pet count. Clean up unused pets or contact admin to increase quota. |
| How to handle expired vaccinations?                  | System notifies staff before expiration; use `UpdatePet` to update vaccination records                                |

---

## 📌 9. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Pet or customer ID does not exist |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 10. Related File References

- [pagination.md](../docs/common/address.md)
- [pet_service.proto](../moego/business/customer/v1/pet_service.proto)
- [pet.proto](../moego/business/customer/v1/pet.proto)