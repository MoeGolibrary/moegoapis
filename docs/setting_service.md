# 🛠️ Setting Service API Documentation (`moego.business.setting.v1`)

## 📌 1. Functional Overview

The `SettingService` provides a centralized interface for managing business service configuration settings. This
includes defining and managing the services offered by the business, such as grooming, boarding, daycare, evaluation,
and training.

This service is essential for maintaining standardized service data across all business locations and ensuring
consistent service delivery.

---

## 🎯 2. Design Goals

- **Centralized Configuration**: Provides a unified way to manage business service definitions.
- **Rich Data Model**: Supports complex relationships like pricing, availability, and staff assignment.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful APIs compatible with mainstream development frameworks.

Applicable to scenarios such as:

- Configuring new services before launch.
- Updating service pricing or availability.
- Standardizing service offerings across multiple locations.

---

## 🧩 3. Core Concepts

### 1. Service

Represents a specific service or add-on that can be provided to pets.

| Field Name             | Type                   | Description                                                   |
|------------------------|------------------------|---------------------------------------------------------------|
| `id`                   | string                 | Unique identifier (`srv_` prefix)                             |
| `name`                 | string                 | Display name of the service                                   |
| `serviceItemType`      | ItemType               | Primary category of the service                               |
| `category`             | string                 | Subcategory for further classification                        |
| `price`                | Money                  | Base price for the service                                    |
| `serviceType`          | Type                   | Whether it's a standalone service or an add-on                |
| `serviceTime`          | int32                  | Duration in minutes                                           |
| `availableAllBusiness` | bool                   | Whether available at all business locations                   |
| `availableBusinessIds` | Array(string)          | List of specific business IDs where the service is available  |
| `availableAllStaff`    | bool                   | Whether available to all staff                                |
| `availableStaffIds`    | Array(string)          | List of staff members who can perform this service            |
| `breedFilter`          | bool                   | Whether the service is available for all pet types and breeds |
| `customizedBreeds`     | Array(CustomizedBreed) | List of pet types and their breeds that can use this service  |
| `petSizeFilter`        | bool                   | Whether the service is available for all pet sizes            |
| `customizedPetSizes`   | Array(string)          | List of pet sizes that can use this service                   |
| `coatFilter`           | bool                   | Whether the service is available for all pet coat types       |
| `customizedCoats`      | Array(string)          | List of pet coat types that can use this service              |
| `petCodeFilter`        | PetCodeFilter          | Pet code filter configuration                                 |
| `serviceFilter`        | bool                   | Whether the add-on is available for all services (only for add-ons) |
| `serviceFilterList`    | Array(ServiceFilter)   | List of service filters defining which services this add-on can be combined with (only for add-ons) |

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

#### Pet Details Configuration

Services can be configured to be available only for specific pet characteristics. This allows for fine-grained control
over which pets can use which services.

##### Type & Breed Filtering

Services can be restricted to specific pet types and breeds using:

- `breedFilter`: When set to `false`, the service is only available for specific pet types and breeds defined in
  `customizedBreeds`
- `customizedBreeds`: An array of `CustomizedBreed` objects that define which pet types and breeds can use this service

###### CustomizedBreed Object

| Field Name    | Type          | Description                                     |
|---------------|---------------|-------------------------------------------------|
| `petTypeId`   | string        | Pet type identifier                             |
| `petTypeName` | string        | Name of the pet type                            |
| `breeds`      | Array(string) | List of breed identifiers for this pet type     |
| `isAll`       | bool          | Whether all breeds of this pet type are allowed |

##### Weight/Size Filtering

Services can be configured based on pet size using:

- `petSizeFilter`: When set to `false`, the service is only available for specific pet sizes defined in
  `customizedPetSizes`
- `customizedPetSizes`: An array of size identifiers that can use this service

##### Coat Type Filtering

Services can be restricted based on pet coat types:

- `coatFilter`: When set to `false`, the service is only available for specific coat types defined in `customizedCoats`
- `customizedCoats`: An array of coat type identifiers that can use this service

##### Pet Code Filtering

Services can be configured to work only with pets that have specific pet codes:

- `petCodeFilter`: A `PetCodeFilter` object that defines which pet codes are allowed for this service

###### PetCodeFilter Object

| Field Name     | Type          | Description                                                          |
|----------------|---------------|----------------------------------------------------------------------|
| `isWhiteList`  | bool          | Whether to use whitelist (true) or blacklist (false)                 |
| `isAllPetCode` | bool          | Whether the service is available for all pet codes                   |
| `petCodeIds`   | Array(string) | List of pet code identifiers (only valid when isAllPetCode is false) |

#### Service Filter Configuration

Services can be configured to work with other services using service filters, particularly for add-ons that need to be combined with primary services:

- `serviceFilter`: When set to `true`, indicates that the service is only available for specific service types or services defined in `serviceFilterList`
- `serviceFilterList`: An array of `ServiceFilter` objects that define which service types and specific services the current service can be combined with

##### ServiceFilter Object

| Field Name               | Type                 | Description                                                 |
|--------------------------|----------------------|-------------------------------------------------------------|
| `serviceItemType`        | Service.ItemType     | Service type that this filter applies to                    |
| `availableForAllServices`| optional bool        | Whether the service is available for all services of the type|
| `availableServiceIdList` | Array(string)        | List of specific service IDs that this service can be combined with (only used when availableForAllServices is false)|

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Service API

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

| Error Code          | Description            |
|---------------------|------------------------|
| `ALREADY_EXISTS`    | Service already exists |
| `PERMISSION_DENIED` | Permission denied      |

---

### 2. Get Service (`GetService`)

- **Method**: `GetService`
- **HTTP Method**: GET
- **Path**: `/v1/setting/companies/{company_id}/services/{id}`

#### ✅ Functionality:

Retrieves detailed information about a specific service.

#### 🎯 Use Cases:

- Display full service details.
- Verify service configuration.

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
    "pageSize": 20,
    "pageToken": "1"
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
| Why does updating a service return "not found"?                      | The specified service ID does not exist. Verify the ID using `GetService` before attempting the update.             |
| How do I handle failed service operations?                           | Check the error message and logs. For rate limiting issues, implement retry logic with exponential backoff.         |

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

- [service.proto](../moego/business/setting/v1/service.proto)
- [setting_service.proto](../moego/business/setting/v1/setting_service.proto)