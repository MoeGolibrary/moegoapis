# Offering API Documentation (`moego.business.offering.v1`)

## 📌 1. Functional Overview

The Offering API provides read-only queries for bookable services and lodging availability at a given business location.
This interface enables:

- Listing which services can be booked for specific pets at a business, filtered by care type or service type
- Querying add-on services that are applicable to a set of already-selected primary services
- Checking which lodging units have capacity for a given date (for boarding and daycare bookings)

---

## 🎯 2. Design Goals

- **Booking-Ready**: Returns the data needed to populate service selection UIs and AI booking flows
- **Pet-Aware**: Supports per-pet service filtering so results reflect size, breed, or type-specific availability
- **Inventory-Aware**: Lodging availability reflects real-time capacity, with reasons when a unit is unavailable
- **Flexible Filtering**: Supports filtering by care type (grooming, boarding, daycare, evaluation, training) and service type (primary vs. add-on)

---

## 🧩 3. Core Concepts

### 1. ServiceView

A flat representation of a single bookable service, suitable for list display or AI consumption.

| Field Name        | Type             | Description                                                   |
|-------------------|------------------|---------------------------------------------------------------|
| `id`              | string           | Unique service identifier (obfuscated)                        |
| `name`            | string           | Display name of the service                                   |
| `serviceItemType` | Service.ItemType | Care category: `GROOMING`, `BOARDING`, `DAYCARE`, etc.        |
| `serviceType`     | Service.Type     | Service classification: `SERVICE` (primary) or `ADDON`        |
| `price`           | Money            | Base price for the service                                    |
| `serviceTime`     | int32            | Duration of the service in minutes                            |
| `description`     | string           | Optional description or notes for the service                 |

### 2. ServicesByPet

A per-pet grouping of applicable services, returned when multiple pet IDs are requested.

| Field Name  | Type                | Description                                                         |
|-------------|---------------------|---------------------------------------------------------------------|
| `petId`     | string              | Pet identifier (obfuscated)                                         |
| `services`  | Array(ServiceView)  | Services applicable to this pet (filtered by pet type/size/breed)  |

### 3. LodgingAvailability

Represents a single lodging unit or type with its availability for the requested date.

| Field Name          | Type   | Description                                                              |
|---------------------|--------|--------------------------------------------------------------------------|
| `lodgingId`         | string | Lodging unit or type identifier (obfuscated)                             |
| `lodgingName`       | string | Display name of the lodging                                              |
| `available`         | bool   | Whether this lodging has capacity for the requested date                 |
| `unavailableReason` | string | Reason when not available (e.g., "over capacity", "date conflict"); empty when `available` is `true` |

### 4. Service.ItemType

| Value                        | Description                                                 |
|------------------------------|-------------------------------------------------------------|
| `SERVICE_ITEM_TYPE_UNSPECIFIED` | Unknown or unspecified                                   |
| `GROOMING`                   | Pet grooming services (bath, haircut, nail trim, etc.)      |
| `BOARDING`                   | Overnight or extended stay care                             |
| `DAYCARE`                    | Supervised daytime care during business hours               |
| `EVALUATION`                 | Behavior assessment and service compatibility check         |

### 5. Service.Type

| Value                    | Description                                                    |
|--------------------------|----------------------------------------------------------------|
| `SERVICE_TYPE_UNSPECIFIED` | Unknown or unspecified                                       |
| `SERVICE`                | Primary standalone service; can be booked independently        |
| `ADDON`                  | Add-on service; must be booked together with a primary service |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: Building a Service Selection UI for Booking

1. **List primary services** — call `ListServices` with `serviceType = SERVICE` and the care type (e.g., `GROOMING`), passing pet IDs to get per-pet filtered results
2. **Customer selects services** — collect the chosen `serviceId` values
3. **List add-ons** — call `ListServices` again with `serviceType = ADDON` and `selectedServiceIds` set to the previously selected primary services; only add-ons compatible with those services are returned
4. **Check lodging** (for boarding/daycare) — call `ListLodgings` with the booking date to show available units; units with `available = false` are shown as grayed-out with `unavailableReason`

---

## 📦 5. API Interface Descriptions

### 1. List Services (`ListServices`)

- **Method**: `ListServices`
- **HTTP Method**: POST
- **Path**: `/v1/offerings/services:list`

#### ✅ Functionality:

Returns a flat list of services and optionally a per-pet breakdown for the requested care type and pets. When querying
add-on services, pass `selectedServiceIds` to filter only the add-ons compatible with the chosen primary services.

#### 🎯 Use Cases:

- Populate a service selection screen in a booking flow
- Let an AI assistant present bookable services for specific pets
- Filter add-ons applicable to already-selected primary services

#### 🔧 Request Parameters:

| Field Name            | Type             | Required | Description                                                                                        |
|-----------------------|------------------|----------|----------------------------------------------------------------------------------------------------|
| `companyId`           | string           | Yes      | Company identifier (obfuscated); tenant context                                                    |
| `businessId`          | string           | No       | Business location identifier (obfuscated); scopes results to this location                         |
| `petIds`              | Array(string)    | No       | Pet identifiers to filter services by (must be existing pets under the customer)                   |
| `serviceType`         | Service.Type     | No       | `SERVICE` for primary services, `ADDON` for add-ons                                               |
| `serviceItemType`     | Service.ItemType | No       | Care category filter: `GROOMING`, `BOARDING`, `DAYCARE`, `EVALUATION`                              |
| `selectedServiceIds`  | Array(string)    | No       | When `serviceType = ADDON`, the already-selected primary service IDs used to filter applicable add-ons |
| `startDate`           | Date             | No       | Start date for availability range (e.g., boarding check-in date)                                   |
| `endDate`             | Date             | No       | End date for availability range (e.g., boarding check-out date)                                    |
| `pagination`          | Pagination       | No       | Pagination parameters: `pageSize`, `pageToken`                                                     |

#### 💡 Example Request:

```json
{
  "companyId": "cmp_abc123",
  "businessId": "biz_001",
  "petIds": ["pet_001", "pet_002"],
  "serviceType": "SERVICE",
  "serviceItemType": "GROOMING"
}
```

#### 📌 Return Value:

| Field Name             | Type                | Description                                                                             |
|------------------------|---------------------|-----------------------------------------------------------------------------------------|
| `services`             | Array(ServiceView)  | All matching services as a flat list                                                    |
| `availableServiceIds`  | Array(string)       | IDs of services currently available in the requested context (for UI highlighting)      |
| `petServices`          | Array(ServicesByPet)| Per-pet list of applicable services; only present when `petIds` were provided           |
| `nextPageToken`        | string              | Token for retrieving the next page; empty if no more results                            |

#### 💡 Example Response:

```json
{
  "services": [
    {
      "id": "svc_001",
      "name": "Full Groom",
      "serviceItemType": "GROOMING",
      "serviceType": "SERVICE",
      "price": { "currencyCode": "USD", "units": 75, "nanos": 0 },
      "serviceTime": 120,
      "description": "Includes bath, haircut, and nail trim"
    },
    {
      "id": "svc_002",
      "name": "Basic Bath",
      "serviceItemType": "GROOMING",
      "serviceType": "SERVICE",
      "price": { "currencyCode": "USD", "units": 40, "nanos": 0 },
      "serviceTime": 60,
      "description": ""
    }
  ],
  "availableServiceIds": ["svc_001", "svc_002"],
  "petServices": [
    {
      "petId": "pet_001",
      "services": [
        {
          "id": "svc_001",
          "name": "Full Groom",
          "serviceItemType": "GROOMING",
          "serviceType": "SERVICE",
          "price": { "currencyCode": "USD", "units": 75, "nanos": 0 },
          "serviceTime": 120,
          "description": ""
        }
      ]
    }
  ],
  "nextPageToken": ""
}
```

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: `companyId` is missing or malformed.
- `PERMISSION_DENIED`: The caller does not have access to the requested business or company.

---

### 2. List Lodgings (`ListLodgings`)

- **Method**: `ListLodgings`
- **HTTP Method**: POST
- **Path**: `/v1/offerings/lodgings:list`

#### ✅ Functionality:

Returns lodging units with availability (inventory) for a given date. Useful when booking boarding or daycare services
to display which lodging units have capacity. Each entry includes whether the unit is available and, if not, the reason.

#### 🎯 Use Cases:

- Show available lodging units when booking a boarding appointment
- Display availability status in a lodging selection UI
- Allow AI assistants to identify suitable lodging units before confirming a booking

#### 🔧 Request Parameters:

| Field Name     | Type          | Required | Description                                                               |
|----------------|---------------|----------|---------------------------------------------------------------------------|
| `companyId`    | string        | Yes      | Company identifier (obfuscated); tenant context                           |
| `businessId`   | string        | No       | Business location identifier (obfuscated); scopes results to this location |
| `date`         | Date          | No       | Date to check lodging availability (inventory)                            |
| `petId`        | string        | No       | Pet identifier for size-based or type-based lodging filtering             |
| `serviceIds`   | Array(string) | No       | Selected service IDs to check availability for a specific booking context |
| `pagination`   | Pagination    | No       | Pagination parameters: `pageSize`, `pageToken`                            |

#### 💡 Example Request:

```json
{
  "companyId": "cmp_abc123",
  "businessId": "biz_001",
  "date": { "year": 2026, "month": 5, "day": 10 },
  "petId": "pet_001",
  "serviceIds": ["svc_boarding_001"]
}
```

#### 📌 Return Value:

| Field Name      | Type                        | Description                                           |
|-----------------|-----------------------------|-------------------------------------------------------|
| `lodgings`      | Array(LodgingAvailability)  | List of lodging units with availability for the date  |
| `nextPageToken` | string                      | Token for retrieving the next page; empty if no more  |

#### 💡 Example Response:

```json
{
  "lodgings": [
    {
      "lodgingId": "lodgu_001",
      "lodgingName": "Suite A",
      "available": true,
      "unavailableReason": ""
    },
    {
      "lodgingId": "lodgu_002",
      "lodgingName": "Suite B",
      "available": false,
      "unavailableReason": "Over capacity"
    }
  ],
  "nextPageToken": ""
}
```

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: `companyId` is missing or malformed.
- `PERMISSION_DENIED`: The caller does not have access to the requested business or company.

---

## 🧪 6. Usage Examples

### Example 1: List Grooming Services for Two Pets

```json
POST /v1/offerings/services:list
{
  "companyId": "cmp_abc123",
  "businessId": "biz_001",
  "petIds": ["pet_001", "pet_002"],
  "serviceType": "SERVICE",
  "serviceItemType": "GROOMING"
}
```

### Example 2: List Add-ons for Selected Primary Services

```json
POST /v1/offerings/services:list
{
  "companyId": "cmp_abc123",
  "businessId": "biz_001",
  "petIds": ["pet_001"],
  "serviceType": "ADDON",
  "selectedServiceIds": ["svc_001"]
}
```

### Example 3: Check Lodging Availability for a Boarding Date

```json
POST /v1/offerings/lodgings:list
{
  "companyId": "cmp_abc123",
  "businessId": "biz_001",
  "date": { "year": 2026, "month": 5, "day": 10 },
  "petId": "pet_001",
  "serviceIds": ["svc_boarding_001"]
}
```

---

## ⚠️ 7. Usage Limitations

- `ListServices` returns a flat list with no category hierarchy; nesting must be handled client-side if needed.
- When `serviceType = ADDON`, `selectedServiceIds` should be provided to filter add-ons correctly. Without it, all add-ons may be returned regardless of primary service compatibility.
- `ListLodgings` checks availability for a single date only; multi-day boarding range checks require calling the Appointment availability API (`CheckAppointmentAvailability`).

---

## 📎 8. FAQ

| Question                                                                  | Answer                                                                                                                         |
|---------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| How do I get add-on services for a specific pet?                          | Pass `petIds` and `serviceType = ADDON` with `selectedServiceIds` set to the chosen primary services.                         |
| What does `availableServiceIds` mean vs the full `services` list?         | `services` is the full catalog; `availableServiceIds` highlights which of those are currently bookable in the given context.   |
| Can I list all services without filtering by pet?                         | Yes. Omit `petIds`. The response will include `services` only, not `petServices`.                                              |
| Why does a lodging show `available = false`?                              | Check `unavailableReason` for details such as "over capacity" or "date conflict".                                              |
| Does `ListLodgings` support multi-day availability checks?                | No. Use `CheckAppointmentAvailability` from the Appointment API for multi-day boarding date ranges.                            |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                      |
|---------------------|------------------------------------------------------------------|
| `INVALID_ARGUMENT`  | Missing or malformed `companyId`                                 |
| `PERMISSION_DENIED` | Caller does not have access to the requested business or company |
| `INTERNAL`          | Internal server error                                            |

---

## 📎 10. Related File References

- [appointment.md](./appointment.md)
- [setting_service.md](./setting_service.md)
- [setting_lodging.md](./setting_lodging.md)
- [offering_service.proto](../moego/business/offering/v1/offering_service.proto)
- [service.proto](../moego/business/setting/v1/service.proto)
