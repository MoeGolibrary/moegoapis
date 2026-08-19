# Aggregation API Documentation (`moego.business.aggregation.v1`)

## 📌 1. Functional Overview

Aggregation provides cross-domain APIs that compose customer, pet, service, and appointment data into unified views.
This interface enables:

- Looking up a client's full profile (basic info, pets, feeding/medication schedules, vaccine records) by phone number
- Retrieving the customer's last and next appointment summaries in a single call
- Designed for receptionist and front-desk integration scenarios where a single API call replaces multiple individual requests

---

## 🎯 2. Design Goals

- **One-call Aggregation**: Eliminates the need for clients to fan out multiple API calls to customer, pet, and appointment services
- **Receptionist-Friendly**: Returns exactly the data a front-desk operator needs to greet a customer and prepare for their appointment
- **Non-blocking Care Details**: Feeding and medication schedules are fetched in parallel and omitted gracefully if unavailable, without blocking the main response
- **Business-scoped**: Appointment context (last/next) is scoped to the specific business location

---

## 🧩 3. Core Concepts

### 1. CustomerSummary

Minimal customer profile returned by aggregation APIs.

| Field Name   | Type   | Description                         |
|--------------|--------|-------------------------------------|
| `id`         | string | Unique customer identifier          |
| `firstName`  | string | Customer's first name               |
| `lastName`   | string | Customer's last name                |
| `phone`      | string | Customer's phone number             |
| `email`      | string | Customer's email address (optional) |

### 2. PetView

Pet profile enriched with care and vaccine details.

| Field Name            | Type                         | Description                                                     |
|-----------------------|------------------------------|-----------------------------------------------------------------|
| `petId`               | string                       | Unique pet identifier                                           |
| `petName`             | string                       | Pet's name                                                      |
| `petType`             | Pet.Type                     | Species/type of the pet                                         |
| `breed`               | string                       | Breed name                                                      |
| `breedMixed`          | bool                         | Whether the pet is mixed breed                                  |
| `gender`              | Pet.Gender                   | Pet's gender                                                    |
| `weight`              | string                       | Pet's weight (e.g., "10 kg")                                    |
| `coatType`            | string                       | Coat type description                                           |
| `fixed`               | string                       | Spayed/neutered status                                          |
| `birthday`            | Date                         | Pet's date of birth                                             |
| `evaluationStatus`    | Pet.EvaluationStatus         | Whether the pet has passed service evaluation                   |
| `feedingSchedules`    | Array(PetFeedingSchedule)    | List of feeding schedules; empty if none configured             |
| `medicationSchedules` | Array(PetMedicationSchedule) | List of medication schedules; empty if none configured          |
| `vaccineList`         | Array(Pet.Vaccination)       | List of vaccination records with name and expiration date       |

### 3. Pet.Type

| Value              | Description                          |
|--------------------|--------------------------------------|
| `TYPE_UNSPECIFIED` | Unknown or unspecified               |
| `OTHER`            | Pet type not in standard categories  |
| `DOG`              | Dog                                  |
| `CAT`              | Cat                                  |
| `BIRD`             | Bird                                 |

### 4. Pet.Gender

| Value                | Description                  |
|----------------------|------------------------------|
| `GENDER_UNSPECIFIED` | Unknown or unspecified       |
| `MALE`               | Male                         |
| `FEMALE`             | Female                       |
| `UNKNOWN`            | Gender not determined        |

### 5. Pet.EvaluationStatus

| Value                         | Description                                    |
|-------------------------------|------------------------------------------------|
| `EVALUATION_STATUS_UNSPECIFIED` | Not yet determined                           |
| `PASS`                        | Pet passed evaluation; eligible for services   |
| `FAIL`                        | Pet requires service modifications             |

### 6. PetFeedingSchedule

Describes a scheduled feeding routine for a pet.

| Field Name           | Type                      | Description                                               |
|----------------------|---------------------------|-----------------------------------------------------------|
| `id`                 | int64                     | Feeding record identifier                                 |
| `petId`              | int64                     | Associated pet identifier                                 |
| `feedingAmount`      | string                    | Amount to feed (e.g., "1.2", "1/2")                       |
| `feedingUnit`        | string                    | Unit of measurement (e.g., cups, oz)                      |
| `feedingType`        | string                    | Type of feeding (e.g., dry, wet)                          |
| `feedingSource`      | string                    | Food source/brand                                         |
| `feedingInstruction` | string                    | Additional feeding instructions (optional)                |
| `feedingTimes`       | Array(PetScheduleTimeDef) | Times of day to feed, in minutes from midnight            |
| `feedingNote`        | string                    | Free-form note about feeding (optional)                   |

### 7. PetMedicationSchedule

Describes a scheduled medication routine for a pet.

| Field Name          | Type                      | Description                                               |
|---------------------|---------------------------|-----------------------------------------------------------|
| `id`                | int64                     | Medication record identifier                              |
| `petId`             | int64                     | Associated pet identifier                                 |
| `medicationAmount`  | string                    | Dosage amount (e.g., "1.2", "1/2")                        |
| `medicationUnit`    | string                    | Unit of measurement (e.g., ml, tablet)                    |
| `medicationName`    | string                    | Name of the medication                                    |
| `medicationNote`    | string                    | Free-form note about the medication                       |
| `medicationTimes`   | Array(PetScheduleTimeDef) | Times of day to administer, in minutes from midnight      |

### 8. Pet.Vaccination

| Field Name  | Type      | Description                                                       |
|-------------|-----------|-------------------------------------------------------------------|
| `name`      | string    | Vaccination name (e.g., "Rabies", "DHPP", "Bordetella")           |
| `expiredAt` | Timestamp | Expiration date of the vaccination; staff are notified in advance |

### 9. AppointmentView

Minimal appointment summary scoped to the queried business.

| Field Name         | Type                        | Description                                                   |
|--------------------|-----------------------------|---------------------------------------------------------------|
| `appointmentId`    | string                      | Unique appointment identifier                                 |
| `startDate`        | Date                        | Appointment start date                                        |
| `startTime`        | int32                       | Start time in minutes from midnight (e.g., 540 = 9:00 AM)    |
| `endDate`          | Date                        | Appointment end date                                          |
| `endTime`          | int32                       | End time in minutes from midnight                             |
| `mainCareType`     | Service.ItemType            | Primary service category of the appointment                   |
| `petServiceDetails`| Array(PetServiceDetailView) | Per-pet service details including staff and lodging assignment |

### 10. PetServiceDetailView

Represents a pet's selected services within an appointment.

| Field Name  | Type                      | Description                            |
|-------------|---------------------------|----------------------------------------|
| `petId`     | string                    | Pet identifier                         |
| `services`  | Array(ServiceDetailView)  | Services booked for this pet           |

### 11. ServiceDetailView

Details of a single service item within an appointment.

| Field Name        | Type             | Description                                          |
|-------------------|------------------|------------------------------------------------------|
| `serviceId`       | string           | Service identifier                                   |
| `serviceName`     | string           | Service name                                         |
| `serviceItemType` | Service.ItemType | Service category                                     |
| `serviceTime`     | int32            | Duration in minutes                                  |
| `servicePrice`    | double           | Price of the service                                 |
| `staffId`         | string           | Assigned staff member identifier (optional)          |
| `lodgingId`       | string           | Assigned lodging unit identifier (optional)          |

### 12. Service.ItemType

| Value                        | Description                       |
|------------------------------|-----------------------------------|
| `SERVICE_ITEM_TYPE_UNSPECIFIED` | Unknown or unspecified          |
| `GROOMING`                   | Pet grooming (bath, haircut, etc.) |
| `BOARDING`                   | Overnight/extended stay           |
| `DAYCARE`                    | Daytime supervised care           |
| `EVALUATION`                 | Behavior assessment               |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: Receptionist Greets a Walk-in or Caller

Here is a typical integration flow for a receptionist or front-desk system:

1. **Customer calls or walks in** — staff captures phone number
2. **Call `LookupClientPetProfile`** with the phone number and the current business ID
3. **Display customer name and contact info** from `customer` field
4. **Display pet list** from `pets` — show breed, weight, evaluation status, and any feeding/medication schedules
5. **Check upcoming appointment** from `nextAppointment` — confirm date, time, services booked, and assigned staff
6. **Review last visit** from `lastAppointment` — provide service history context for the current visit

---

## 📦 5. API Interface Descriptions

### 1. Lookup Client and Pet Profile (`LookupClientPetProfile`)

- **Method**: `LookupClientPetProfile`
- **HTTP Method**: POST
- **Path**: `/v1/aggregation/clients:lookupClientPetProfile`

#### ✅ Functionality:

Resolves a customer by phone number within the given business context, then returns:
- Minimal customer profile
- All pets belonging to the customer, enriched with feeding schedules, medication schedules, and vaccination records
- The customer's last and next appointment at the specified business location

Feeding and medication schedule fetches run in parallel and are omitted gracefully on failure without blocking the response.

#### 🎯 Use Cases:

- Receptionist system identifies a caller by phone number before they arrive
- Front-desk kiosk displays pet care notes and upcoming appointments for a check-in
- Integration system pre-populates service forms based on pet profile

#### 🔧 Request Parameters:

| Field Name    | Type   | Required | Description                                                            |
|---------------|--------|----------|------------------------------------------------------------------------|
| `phoneNumber` | string | Yes      | Customer's phone number. Normalized by the caller. Length: 7–30 chars  |
| `businessId`  | string | Yes      | Obfuscated business ID to scope the appointment context                |

#### 💡 Example Request:

```json
{
  "phoneNumber": "+12125551234",
  "businessId": "biz_abc123"
}
```

#### 📌 Return Value:

| Field Name        | Type                  | Description                                                            |
|-------------------|-----------------------|------------------------------------------------------------------------|
| `customer`        | CustomerSummary       | Matched customer's basic profile (optional; absent if not found)       |
| `pets`            | Array(PetView)        | All pets belonging to the customer, with care and vaccine details      |
| `lastAppointment` | AppointmentView       | Most recent completed appointment at the business (optional)           |
| `nextAppointment` | AppointmentView       | Next upcoming appointment at the business (optional)                   |

#### 💡 Example Response:

```json
{
  "customer": {
    "id": "cus_abc123",
    "firstName": "Jane",
    "lastName": "Smith",
    "phone": "+12125551234",
    "email": "jane.smith@example.com"
  },
  "pets": [
    {
      "petId": "pet_xyz789",
      "petName": "Buddy",
      "petType": "DOG",
      "breed": "Golden Retriever",
      "breedMixed": false,
      "gender": "MALE",
      "weight": "30 kg",
      "coatType": "Long",
      "fixed": "Yes",
      "birthday": { "year": 2019, "month": 4, "day": 15 },
      "evaluationStatus": "PASS",
      "feedingSchedules": [
        {
          "id": 1001,
          "petId": 5001,
          "feedingAmount": "1",
          "feedingUnit": "cup",
          "feedingType": "Dry",
          "feedingSource": "Royal Canin",
          "feedingInstruction": "Mix with warm water",
          "feedingTimes": [{ "timeMinutes": 480 }, { "timeMinutes": 1080 }],
          "feedingNote": "Do not overfeed"
        }
      ],
      "medicationSchedules": [
        {
          "id": 2001,
          "petId": 5001,
          "medicationAmount": "1",
          "medicationUnit": "tablet",
          "medicationName": "Heartgard",
          "medicationNote": "Give with food",
          "medicationTimes": [{ "timeMinutes": 480 }]
        }
      ],
      "vaccineList": [
        { "name": "Rabies", "expiredAt": "2026-09-01T00:00:00Z" },
        { "name": "DHPP", "expiredAt": "2025-12-01T00:00:00Z" }
      ]
    }
  ],
  "lastAppointment": {
    "appointmentId": "appt_111",
    "startDate": { "year": 2025, "month": 11, "day": 5 },
    "startTime": 540,
    "endDate": { "year": 2025, "month": 11, "day": 5 },
    "endTime": 660,
    "mainCareType": "GROOMING",
    "petServiceDetails": [
      {
        "petId": "pet_xyz789",
        "services": [
          {
            "serviceId": "svc_001",
            "serviceName": "Full Groom",
            "serviceItemType": "GROOMING",
            "serviceTime": 120,
            "servicePrice": 75.0,
            "staffId": "stf_001"
          }
        ]
      }
    ]
  },
  "nextAppointment": {
    "appointmentId": "appt_222",
    "startDate": { "year": 2026, "month": 4, "day": 10 },
    "startTime": 600,
    "endDate": { "year": 2026, "month": 4, "day": 10 },
    "endTime": 720,
    "mainCareType": "GROOMING",
    "petServiceDetails": [
      {
        "petId": "pet_xyz789",
        "services": [
          {
            "serviceId": "svc_001",
            "serviceName": "Full Groom",
            "serviceItemType": "GROOMING",
            "serviceTime": 120,
            "servicePrice": 80.0,
            "staffId": "stf_002"
          }
        ]
      }
    ]
  }
}
```

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: `phone_number` or `business_id` is missing or malformed.
- `NOT_FOUND`: No customer found with the given phone number, or the business ID does not exist.
- `PERMISSION_DENIED`: The caller does not have access to the requested business or company.

---

## 🧪 6. Usage Examples

### Example 1: Look Up Profile by Phone

```json
POST /v1/aggregation/clients:lookupClientPetProfile
{
  "phoneNumber": "+12125551234",
  "businessId": "biz_abc123"
}
```

Response includes customer name, all pets with care schedules and vaccines, and last/next appointments.

### Example 2: Handle "No Upcoming Appointment"

If `nextAppointment` is absent in the response, the customer has no upcoming appointment at this business. You can
prompt staff to schedule one via the Appointment API.

### Example 3: Handle "No Last Appointment"

If `lastAppointment` is absent, this is a new customer at this business with no prior visit history.

---

## ⚠️ 7. Usage Limitations

- The phone number lookup matches on the customer's **main phone number** only. Secondary/alternate numbers are not searched.
- A maximum of **100 pets** are returned per customer.
- Feeding and medication schedules are fetched on a **best-effort basis**: if those services are temporarily unavailable, the pet list is still returned but `feedingSchedules` and `medicationSchedules` will be empty.
- Appointment context (`lastAppointment` / `nextAppointment`) is scoped to the specified `businessId`. Appointments at other business locations under the same company are not included.

---

## 📎 8. FAQ

| Question                                                               | Answer                                                                                                                                   |
|------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| What if multiple customers share the same phone number?                | The API returns the first matched customer. Ensure phone numbers are unique when managing customer data.                                  |
| Why are `feedingSchedules` empty even though I set them?               | Schedule data is fetched in parallel and omitted gracefully on service errors. Retry the request or check the pet care configuration.     |
| What format is `startTime` / `endTime`?                                | Minutes from midnight. For example, 540 = 9:00 AM, 570 = 9:30 AM, 1080 = 6:00 PM.                                                       |
| Can I filter which pets are returned?                                  | No. All pets belonging to the customer are returned. Filter client-side if needed.                                                        |
| Does `lastAppointment` include cancelled appointments?                 | No. Only completed appointments are considered when determining the last visit.                                                           |
| What does it mean if `evaluationStatus` is `FAIL`?                     | The pet did not pass the behavior evaluation and may require service modifications. Confirm with the business before booking.             |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                      |
|---------------------|------------------------------------------------------------------|
| `NOT_FOUND`         | Customer not found by phone number, or business ID is invalid    |
| `INVALID_ARGUMENT`  | Missing or malformed `phone_number` or `business_id`             |
| `PERMISSION_DENIED` | Caller does not have access to the requested business or company |
| `INTERNAL`          | Internal server error                                            |

---

## 📎 10. Related File References

- [customer.md](./customer.md)
- [pet.md](./pet.md)
- [appointment.md](./appointment.md)
- [aggregation_service.proto](../moego/business/aggregation/v1/aggregation_service.proto)
- [customer.proto](../moego/business/customer/v1/customer.proto)
- [pet.proto](../moego/business/customer/v1/pet.proto)
- [service.proto](../moego/business/setting/v1/service.proto)
