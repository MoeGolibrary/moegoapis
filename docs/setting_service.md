# 🛠️ Setting Service API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The `SettingService` provides a centralized interface for managing business configuration settings, including:

- Managing **Pet Codes** used to flag special handling requirements or medical conditions.
- Managing **Customer Tags** that categorize clients for marketing and service customization.
- Managing **Services** offered by the business, such as grooming, boarding, daycare, evaluation, and training.
- Managing **Lodging Types and Units** for standardizing boarding accommodations.
- Supporting **creation, updating, listing, and retrieving** of service definitions.
- Enabling filtering and pagination for scalable data management.

This service is essential for maintaining standardized data across all business locations and ensuring consistent
service delivery.

---

## 🎯 2. Design Goals

- **Centralized Configuration**: Provides a unified way to manage business-wide settings like services, tags, and codes.
- **Rich Data Model**: Supports complex relationships like pricing, availability, and staff assignment.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Configuring new services before launch.
- Updating service pricing or availability.
- Categorizing customers for reporting or marketing campaigns.
- Standardizing pet handling codes across multiple locations.

---

## 🧩 3. Core Concepts

### 1. Service

Represents a specific service or add-on that can be provided to pets.

| Field Name             | Type          | Description                                                  |
|------------------------|---------------|--------------------------------------------------------------|
| `id`                   | string        | Unique identifier (`srv_` prefix)                            |
| `name`                 | string        | Display name of the service                                  |
| `serviceItemType`      | ItemType      | Primary category of the service                              |
| `category`             | string        | Subcategory for further classification                       |
| `price`                | Money         | Base price for the service                                   |
| `serviceType`          | Type          | Whether it's a standalone service or an add-on               |
| `serviceTime`          | int32         | Duration in minutes                                          |
| `availableAllBusiness` | bool          | Whether available at all business locations                  |
| `availableBusinessIds` | Array(string) | List of specific business IDs where the service is available |
| `availableAllStaff`    | bool          | Whether available to all staff                               |
| `availableStaffIds`    | Array(string) | List of staff members who can perform this service           |

#### Enum: ItemType

- `SERVICE_ITEM_TYPE_UNSPECIFIED`
- `GROOMING`
- `BOARDING`
- `DAYCARE`
- `EVALUATION`
- `TRAINING`

#### Enum: Type

- `SERVICE_TYPE_UNSPECIFIED`
- `SERVICE`
- `ADDON`

---

### 2. Lodging

A lodging represents a type of accommodation and its associated units, typically used for boarding services.

| Field Name     | Type               | Description                              |
|----------------|--------------------|------------------------------------------|
| `lodgingType`  | LodgingType        | The type of lodging (e.g., room, area)   |
| `lodgingUnits` | Array(LodgingUnit) | List of individual units within the type |

---

### 3. PetCode

Represents special handling instructions or medical alerts for a pet.

| Field Name     | Type   | Description                          |
|----------------|--------|--------------------------------------|
| `id`           | string | Unique identifier                    |
| `abbreviation` | string | Short form (e.g., AG for Aggressive) |
| `description`  | string | Detailed explanation                 |
| `color`        | string | Highlight color in UI                |

---

### 4. CustomerTag

Represents a label that can be applied to customers for categorization and filtering purposes.

| Field Name        | Type      | Description                                       |
|-------------------|-----------|---------------------------------------------------|
| `id`              | string    | Unique identifier                                 |
| `name`            | string    | Display name of the tag                           |
| `lastUpdatedBy`   | string    | ID of the staff member who last modified this tag |
| `lastUpdatedTime` | Timestamp | When this tag was last modified                   |

---

### 5. ReferralSource

Represents the source or channel through which a customer was acquired.

| Field Name | Type   | Description                |
|------------|--------|----------------------------|
| `id`       | string | Unique identifier          |
| `name`     | string | Display name of the source |

---

### 6. LifeCycle

Represents a stage in the lead management process.

| Field Name | Type   | Description       |
|------------|--------|-------------------|
| `id`       | string | Unique identifier |
| `name`     | string | Display name      |

---

### 7. ActionStatus

Represents the status of an action taken on a lead.

| Field Name | Type   | Description       |
|------------|--------|-------------------|
| `id`       | string | Unique identifier |
| `name`     | string | Display name      |

---

#### Message: LodgingType

Describes the type of lodging available.

| Field Name        | Type            | Description                            |
|-------------------|-----------------|----------------------------------------|
| `id`              | string          | Unique identifier for the lodging type |
| `name`            | string          | Name of the lodging type               |
| `description`     | string          | Optional description                   |
| `photoList`       | Array(string)   | URLs to photos of this lodging type    |
| `maxPetNum`       | int32           | Maximum number of pets allowed         |
| `lodgingUnitType` | LodgingUnitType | Type of lodging unit                   |

---

#### Message: LodgingUnit

Describes a specific unit within a lodging type.

| Field Name | Type   | Description                |
|------------|--------|----------------------------|
| `id`       | string | Unique identifier for unit |
| `name`     | string | Display name of the unit   |

---

#### Enum: LodgingUnitType

| Value                           | Description                       |
|---------------------------------|-----------------------------------|
| `LODGING_UNIT_TYPE_UNSPECIFIED` | Default value; should not be used |
| `ROOM`                          | Room/kennel type                  |
| `AREA`                          | Open area for multiple pets       |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Setting API

Here is a typical integration flow:

1. **Create Service**
    - Define a new service with required attributes (name, type, price).
    - Optionally specify business and staff availability.

2. **Update Service**
    - Modify existing service details like price, duration, or availability.

3. **Retrieve Service**
    - Fetch full service information using its ID.

4. **List Services**
    - View all available services, optionally filtered by location or type.

5. **Manage Pet Codes & Customer Tags**
    - Retrieve active codes and tags for use in other modules.

6. Manage Lodging Types and Units
    - Retrieve lodging types and units for display or booking integration.

---

## 📦 5. API Interface Descriptions

### 1. Create Service (`CreateService`)

- **Method**: `CreateService`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/services`

#### ✅ Functionality:

Registers a new service definition with base attributes.

#### 🎯 Use Cases:

- Add a new grooming package.
- Define a new training session type.

#### 🔧 Request Parameters:

| Field Name             | Type          | Required | Description                                   |
|------------------------|---------------|----------|-----------------------------------------------|
| `companyId`            | string        | Yes      | ID of the company creating the service        |
| `name`                 | string        | Yes      | Service name                                  |
| `businessIds`          | Array(string) | No       | Business locations where service is available |
| `serviceItemType`      | ItemType      | Yes      | Service category                              |
| `price`                | Money         | Yes      | Base price                                    |
| `serviceType`          | Type          | Yes      | Whether primary or add-on                     |
| `serviceTime`          | int32         | No       | Duration in minutes                           |
| `availableAllBusiness` | bool          | No       | Available at all locations                    |
| `availableBusinessIds` | Array(string) | No       | Specific business IDs if not all              |
| `availableAllStaff`    | bool          | No       | Available to all staff                        |
| `availableStaffIds`    | Array(string) | No       | Specific staff IDs                            |

#### 📌 Return Value:

| Field Name | Type      | Description                |
|------------|-----------|----------------------------|
| `service`  | `Service` | The created service object |

#### ⚠️ Error Codes:

| Error Code          | Description               |
|---------------------|---------------------------|
| `INVALID_ARGUMENT`  | Missing or invalid fields |
| `PERMISSION_DENIED` | Permission denied         |

---

### 2. Get Service (`GetService`)

- **Method**: `GetService`
- **HTTP Method**: GET
- **Path**: `/v1/setting/companies/{company_id}/services/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific service.

#### 🎯 Use Cases:

- Verify current service details before updates.
- Audit service configurations.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company owning service |
| `id`        | string | Yes      | Service ID to retrieve |

#### 📌 Return Value:

| Field Name | Type      | Description                  |
|------------|-----------|------------------------------|
| `service`  | `Service` | The retrieved service object |

#### ⚠️ Error Codes:

| Error Code          | Description            |
|---------------------|------------------------|
| `NOT_FOUND`         | Service does not exist |
| `PERMISSION_DENIED` | Permission denied      |

---

### 3. Update Service (`UpdateService`)

- **Method**: `UpdateService`
- **HTTP Method**: PUT
- **Path**: `/v1/setting/companies/{company_id}/services/{id}`

#### ✅ Functionality:

Modifies existing service attributes.

#### 🎯 Use Cases:

- Adjust pricing due to cost changes.
- Change availability or staff assignments.

#### 🔧 Request Parameters:

| Field Name             | Type          | Required | Description                 |
|------------------------|---------------|----------|-----------------------------|
| `companyId`            | string        | Yes      | Company owning service      |
| `id`                   | string        | Yes      | Service ID to update        |
| `name`                 | string        | Yes      | Updated service name        |
| `businessIds`          | Array(string) | No       | Updated business locations  |
| `price`                | Money         | Yes      | Updated base price          |
| `serviceTime`          | int32         | No       | Updated duration            |
| `availableAllBusiness` | bool          | No       | Updated availability status |
| `availableBusinessIds` | Array(string) | No       | Updated business IDs        |
| `availableAllStaff`    | bool          | No       | Updated staff availability  |
| `availableStaffIds`    | Array(string) | No       | Updated staff member list   |
| `inactive`             | bool          | No       | Mark service as inactive    |

#### 📌 Return Value:

| Field Name | Type      | Description                |
|------------|-----------|----------------------------|
| `service`  | `Service` | The updated service object |

#### ⚠️ Error Codes:

| Error Code          | Description            |
|---------------------|------------------------|
| `NOT_FOUND`         | Service does not exist |
| `PERMISSION_DENIED` | Permission denied      |

---

### 4. List Services (`ListServices`)

- **Method**: `ListServices`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/services:list`

#### ✅ Functionality:

Lists services matching specified criteria.

#### 🎯 Use Cases:

- View all available services.
- Filter services by location or type.

#### 🔧 Request Parameters:

| Field Name         | Type            | Required | Description                  |
|--------------------|-----------------|----------|------------------------------|
| `companyId`        | string          | Yes      | Company owning services      |
| `pagination`       | Pagination      | Yes      | Page size and token          |
| `businessIds`      | Array(string)   | No       | Filter by business locations |
| `filter.itemTypes` | Array(ItemType) | No       | Filter by service types      |

#### 📌 Return Value:

| Field Name      | Type             | Description                        |
|-----------------|------------------|------------------------------------|
| `nextPageToken` | string           | Token for retrieving the next page |
| `services`      | Array(`Service`) | List of services matching criteria |

#### ⚠️ Error Code:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 5. List Pet Codes (`ListPetCodes`)

- **Method**: `ListPetCodes`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/pet/codes:list`

#### ✅ Functionality:

Retrieves all active pet codes defined for the company.

#### 🎯 Use Cases:

- Retrieve special handling instructions.
- Standardize pet care alerts across locations.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type             | Description                       |
|------------|------------------|-----------------------------------|
| `codes`    | Array(`PetCode`) | List of pet codes for the company |

#### ⚠️ Error Code:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 6. List Customer Tags (`ListCustomerTags`)

- **Method**: `ListCustomerTags`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/customer/tags:list`

#### ✅ Functionality:

Retrieves all active customer tags defined for the company.

#### 🎯 Use Cases:

- Categorize customers for marketing.
- Apply labels for loyalty programs or preferences.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description            |
|-------------|--------|----------|------------------------|
| `companyId` | string | Yes      | Company ID to retrieve |

#### 📌 Return Value:

| Field Name | Type                 | Description                           |
|------------|----------------------|---------------------------------------|
| `tags`     | Array(`CustomerTag`) | List of customer tags for the company |

#### ⚠️ Error Code:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 7. List Lodgings (`ListLodgings`)

- **Method**: `ListLodgings`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/lodgings`

#### ✅ Functionality:

Retrieves a list of all lodging configurations defined for the company.

#### 🎯 Use Cases:

- Retrieve lodging types and units for booking or display.
- Standardize lodging offerings across business locations.

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

### 8. List Customer Referral Sources (`ListCustomerReferralSources`)

- **Method**: `ListCustomerReferralSources`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/customer/referral_sources:list`

#### ✅ Functionality:

Lists all available customer referral sources for a company.

#### 🎯 Use Cases:

- Track the origin of new customers.
- Analyze marketing effectiveness.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                           |
|-------------|--------|----------|---------------------------------------|
| `companyId` | string | Yes      | ID of the company to list sources for |

#### 📌 Return Value:

| Field Name        | Type                    | Description              |
|-------------------|-------------------------|--------------------------|
| `referralSources` | Array(`ReferralSource`) | List of referral sources |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |

---

### 9. List Leads Life Cycles (`ListLeadsLifeCycles`)

- **Method**: `ListLeadsLifeCycles`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/leads/life_cycles`

#### ✅ Functionality:

Returns a list of lead life cycles.

#### 🎯 Use Cases:

- Understand the stages of a lead's journey.
- Standardize lead management processes.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                          |
|-------------|--------|----------|--------------------------------------|
| `companyId` | string | Yes      | ID of the company to list cycles for |

#### 📌 Return Value:

| Field Name   | Type               | Description              |
|--------------|--------------------|--------------------------|
| `lifeCycles` | Array(`LifeCycle`) | List of lead life cycles |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |
| `INVALID_ARGUMENT`  | Malformed request |

---

### 10. List Leads Action Status (`ListLeadsActionStatus`)

- **Method**: `ListLeadsActionStatus`
- **HTTP Method**: POST
- **Path**: `/v1/setting/companies/{company_id}/leads/action_status`

#### ✅ Functionality:

Returns a list of lead action statuses.

#### 🎯 Use Cases:

- Track the status of actions taken on leads.
- Standardize lead follow-up procedures.

#### 🔧 Request Parameters:

| Field Name  | Type   | Required | Description                            |
|-------------|--------|----------|----------------------------------------|
| `companyId` | string | Yes      | ID of the company to list statuses for |

#### 📌 Return Value:

| Field Name       | Type                  | Description                  |
|------------------|-----------------------|------------------------------|
| `actionStatuses` | Array(`ActionStatus`) | List of lead action statuses |

#### ⚠️ Error Codes:

| Error Code          | Description       |
|---------------------|-------------------|
| `PERMISSION_DENIED` | Permission denied |
| `INVALID_ARGUMENT`  | Malformed request |

---

## 🧪 6. Usage Examples

### Example 1: Create Service

Request Body:

```json
{
  "companyId": "cmp_001",
  "name": "Premium Grooming",
  "serviceItemType": "GROOMING",
  "price": {
    "currencyCode": "USD",
    "units": 75,
    "nanos": 0
  },
  "serviceType": "SERVICE",
  "serviceTime": 90,
  "availableAllBusiness": true
}
```

### Example 2: List Services

Request Body:

```json
{
  "companyId": "cmp_001",
  "pagination": {
    "pageSize": 20
  },
  "businessIds": [
    "biz_001",
    "biz_002"
  ],
  "filter": {
    "itemTypes": [
      "GROOMING",
      "DAYCARE"
    ]
  }
}
```

### Example 3: List Lodgings

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

| Question                                                             | Answer                                                                                                              |
|----------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|
| How can I verify if a service configuration is effective?            | Use `ListServices` to check if the service appears in the response with expected attributes.                        |
| How can I prevent duplicate services from being created?             | Ensure that your system checks for existing services before creating new ones.                                      |
| What should I do if a service creation returns "resource exhausted"? | Clean up unused services or contact support to request a quota increase.                                            |
| Can I restrict services to specific business locations?              | Yes, via the `availableBusinessIds` field. Set `availableAllBusiness = false` and specify the allowed business IDs. |
| How can I manage service tags effectively?                           | Use `ListCustomerTags` to retrieve available tags and ensure consistency across services.                           |
| Why does updating a service return "not found"?                      | The specified service ID does not exist. Verify the ID using `GetService` before attempting the update.             |
| How do I handle failed service operations?                           | Check the error message and logs. For rate limiting issues, implement retry logic with exponential backoff.         |
| How can I manage boarding accommodations effectively?                | Use `ListLodgings` to retrieve lodging types and units for consistent boarding management.                          |

---

## 📌 9. Common Error Codes

| Error Code           | Description                                                                 |
|----------------------|-----------------------------------------------------------------------------|
| `ALREADY_EXISTS`     | A setting (e.g., service, tag) with the same name or ID already exists.     |
| `NOT_FOUND`          | The requested setting (e.g., service, pet code) does not exist.             |
| `PERMISSION_DENIED`  | Current user has no access rights to perform the operation.                 |
| `INVALID_ARGUMENT`   | Invalid request parameters (e.g., missing required fields, invalid format). |
| `INTERNAL`           | Internal server error occurred while processing the request.                |
| `RESOURCE_EXHAUSTED` | Request exceeds system-defined quotas or limits.                            |

---

## 📎 10. Related File References

- [lodging.proto](../moego/business/setting/v1/lodging.proto)
- [service.proto](../moego/business/setting/v1/service.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)