# Online Booking API Documentation (`moego.business.online_booking.v1`)

## 📌 1. Functional Overview

The `OnlineBooking` module provides functionality for managing digital appointment requests made through the customer
portal, mobile app, or website. It supports:

- Capturing booking details including pets, services, and time preferences
- Tracking the status of online bookings (normal, waiting list, abandoned)
- Managing abandoned bookings with recovery tracking
- Listing and filtering bookings based on business needs

This API is essential for handling customer-initiated service scheduling and analyzing conversion rates from online
channels.

---

## 🎯 2. Design Goals

- **Centralized Booking Management**: Unified interface for capturing and managing all online booking data.
- **State Tracking**: Clear lifecycle management for online bookings (normal → abandoned) and recovery paths.
- **Recovery Analytics**: Tools to analyze abandonment causes and recovery strategies.
- **Integration Ready**: RESTful endpoints that support integration with marketing tools, analytics platforms, and CRM
  systems.

---

## 🧩 3. Core Concepts

### 1. OnlineBooking

Represents a service appointment request initiated through digital channels.

| Field Name          | Type                    | Description                                                       |
|---------------------|-------------------------|-------------------------------------------------------------------|
| `id`                | string                  | Unique identifier (e.g., "bkg_12345")                             |
| `businessId`        | string                  | ID of the business location where service will be performed       |
| `customerId`        | string                  | ID of the customer initiating the booking                         |
| `address`           | Address                 | Service location details (required for home service appointments) |
| `duration`          | Interval                | Requested time window for the service                             |
| `petServiceDetails` | Array(PetServiceDetail) | List of services requested for each pet                           |
| `status`            | enum(Status)            | Current state: `NORMAL`, `IN_WAIT_LIST`, `ABANDONED`              |
| `colorCode`         | string                  | UI display color in hex format                                    |
| `createdTime`       | timestamp               | When the booking was created                                      |

#### Enum Definitions

##### `OnlineBooking.Status`

- `STATUS_UNSPECIFIED`: Invalid value used only as default
- `NORMAL`: Booking is being processed normally
- `IN_WAIT_LIST`: Booking is on a waiting list due to unavailability
- `ABANDONED`: Booking was not completed by the customer

---

### 2. AbandonedBooking

Captures information about a booking that was not completed by the customer.

| Field Name        | Type                       | Description                                           |
|-------------------|----------------------------|-------------------------------------------------------|
| `id`              | string                     | Unique identifier                                     |
| `businessId`      | string                     | Business location ID                                  |
| `customer`        | AbandonedBookingCustomer   | Customer details                                      |
| `pets`            | Array(AbandonedBookingPet) | Pets and related services involved                    |
| `address`         | Address                    | Service location details                              |
| `abandonStep`     | Step                       | Stage at which the booking was abandoned              |
| `abandonTime`     | timestamp                  | Time when the booking was abandoned                   |
| `abandonStatus`   | Status                     | Current status: `ABANDONED`, `CONTACTED`, `RECOVERED` |
| `lastTextedTime`  | timestamp                  | Last SMS contact attempt regarding this booking       |
| `lastEmailedTime` | timestamp                  | Last email sent about this booking                    |
| `recoveryType`    | RecoverType                | How the booking was recovered (if applicable)         |
| `recoveryTime`    | timestamp                  | Time when the booking was recovered                   |
| `staffId`         | string                     | Staff member who handled recovery                     |
| `appointmentId`   | string                     | Appointment ID if booking was converted               |
| `appointmentTime` | timestamp                  | Scheduled appointment time if one was set             |
| `leadType`        | LeadType                   | How the customer was acquired                         |
| `createdTime`     | timestamp                  | When this record was created                          |
| `lastUpdatedTime` | timestamp                  | When this record was last updated                     |
| `additionalNote`  | string                     | Notes added by staff or system                        |
| `careType`        | ItemType                   | Type of service selected by the customer              |

#### Enum Definitions

##### `AbandonedBooking.LeadType`

- `LEAD_TYPE_UNSPECIFIED`: Default/invalid value
- `NEW_VISITOR`: First-time visitor
- `EXISTING_CLIENT`: Returning customer

##### `AbandonedBooking.RecoverType`

- `ABANDON_RECOVER_TYPE_UNSPECIFIED`: Default/invalid value
- `RECOVERED_BY_SCHEDULE_APPOINTMENT`: Converted into an appointment
- `RECOVERED_BY_EMAIL`: Recovered via email communication
- `RECOVERED_BY_MESSAGE`: Recovered via messaging (e.g., SMS)

##### `AbandonedBooking.Step`

Indicates the stage in the booking flow where the customer abandoned the process.

| Value                 | Description                       |
|-----------------------|-----------------------------------|
| `WELCOME_PAGE`        | Customer visited the welcome page |
| `BASIC_INFO`          | Customer filled out basic info    |
| `SELECT_CARE_TYPE`    | Selected care type                |
| `SELECT_ADDRESS`      | Selected address                  |
| `SELECT_PET`          | Selected pet                      |
| `SELECT_DATE`         | Selected date                     |
| `SELECT_SERVICE`      | Selected service                  |
| `SELECT_GROOMER`      | Selected groomer                  |
| `SELECT_TIME`         | Selected time                     |
| `ADDITIONAL_PET_INFO` | Provided additional pet info      |
| `PERSONAL_INFO`       | Entered personal info             |
| `CARD_ON_FILE`        | Entered card-on-file info         |
| `PREPAY`              | Initiated prepay                  |
| `PRE_AUTH`            | Initiated pre-auth                |
| `SUBMIT_APPT`         | Submitted the appointment         |

##### `AbandonedBooking.Status`

- `ABANDON_STATUS_UNSPECIFIED`: Default/invalid
- `ABANDONED`: The booking was abandoned
- `CONTACTED`: The customer has been contacted
- `RECOVERED`: The booking was successfully recovered

##### `AbandonedBooking.PreferredFrequencyType`

- `PREFERRED_FREQUENCY_TYPE_UNSPECIFIED`
- `PREFERRED_FREQUENCY_TYPE_DAY`
- `PREFERRED_FREQUENCY_TYPE_WEEK`
- `PREFERRED_FREQUENCY_TYPE_MONTH`

---

### 3. AbandonedBookingCustomer

Contains customer-specific data collected during the booking flow.

| Field Name               | Type                   | Description                                                     |
|--------------------------|------------------------|-----------------------------------------------------------------|
| `customerId`             | string                 | Customer's unique ID                                            |
| `email`                  | string                 | Email address                                                   |
| `firstName`              | string                 | First name                                                      |
| `lastName`               | string                 | Last name                                                       |
| `phoneNumber`            | string                 | Phone number                                                    |
| `referer`                | string                 | Source URL that referred the user                               |
| `referralSource`         | string                 | Referral source                                                 |
| `preferredGroomerId`     | string                 | Preferred groomer ID                                            |
| `preferredFrequencyDay`  | int32                  | Preferred frequency in days                                     |
| `preferredFrequencyType` | PreferredFrequencyType | Frequency type (daily, weekly, monthly)                         |
| `preferredDays`          | Array(int32)           | Preferred days of the week                                      |
| `preferredTime`          | Array(int32)           | Preferred times of the day                                      |
| `questionAnswerList`     | Array(QuestionAnswer)  | Key-value pairs capturing user input during the booking process |

---

### 4. AbandonedBookingPet

Represents a pet involved in an abandoned booking, along with service and emergency info.

| Field Name              | Type                  | Description                        |
|-------------------------|-----------------------|------------------------------------|
| `petServiceDetail`      | PetServiceDetail      | Services booked for this pet       |
| `emergencyContactName`  | string                | Emergency contact name             |
| `emergencyContactPhone` | string                | Emergency contact phone number     |
| `healthIssues`          | string                | Known health issues or conditions  |
| `questionAnswerList`    | Array(QuestionAnswer) | Pet-specific question-answer pairs |

---

### 5. QuestionAnswer

A key-value pair capturing user input during the booking process.

| Field Name | Type   | Description                        |
|------------|--------|------------------------------------|
| `key`      | string | Unique identifier for the question |
| `question` | string | The question shown to the user     |
| `answer`   | string | The answer provided by the user    |

---

## 📦 4. API Interface Descriptions

### 1. Get Abandoned Booking (`GetAbandonedBooking`)

- **Method**: `GetAbandonedBooking`
- **HTTP Method**: GET
- **Path**: `/v1/abandoned_bookings/{id}`

#### ✅ Functionality

Retrieves detailed information about a specific abandoned booking.

#### 🔧 Request Parameters

| Field Name | Type   | Required | Description                      |
|------------|--------|----------|----------------------------------|
| `id`       | string | Yes      | Abandoned booking ID to retrieve |

#### 📌 Return Value

| Field Name        | Type                       | Description                                           |
|-------------------|----------------------------|-------------------------------------------------------|
| `id`              | string                     | Unique identifier                                     |
| `businessId`      | string                     | Business location ID                                  |
| `customer`        | AbandonedBookingCustomer   | Customer details                                      |
| `pets`            | Array(AbandonedBookingPet) | Pets and related services involved                    |
| `address`         | Address                    | Service location details                              |
| `abandonStep`     | Step                       | Stage at which the booking was abandoned              |
| `abandonTime`     | timestamp                  | Time when the booking was abandoned                   |
| `abandonStatus`   | Status                     | Current status: `ABANDONED`, `CONTACTED`, `RECOVERED` |
| `lastTextedTime`  | timestamp                  | Last SMS contact attempt regarding this booking       |
| `lastEmailedTime` | timestamp                  | Last email sent about this booking                    |
| `recoveryType`    | RecoverType                | How the booking was recovered (if applicable)         |
| `recoveryTime`    | timestamp                  | Time when the booking was recovered                   |
| `staffId`         | string                     | Staff member who handled recovery                     |
| `appointmentId`   | string                     | Appointment ID if booking was converted               |
| `appointmentTime` | timestamp                  | Scheduled appointment time if one was set             |
| `leadType`        | LeadType                   | How the customer was acquired                         |
| `createdTime`     | timestamp                  | When this record was created                          |
| `lastUpdatedTime` | timestamp                  | When this record was last updated                     |
| `additionalNote`  | string                     | Notes added by staff or system                        |
| `careType`        | ItemType                   | Type of service selected by the customer              |

#### ⚠️ Error Codes

- `NOT_FOUND`: Specified booking ID does not exist
- `PERMISSION_DENIED`: Permission denied

---

### 2. List Abandoned Bookings (`ListAbandonedBookings`)

- **Method**: `ListAbandonedBookings`
- **HTTP Method**: POST
- **Path**: `/v1/abandoned_bookings:list`

#### ✅ Functionality

Lists abandoned bookings matching specified criteria including abandon time range, lead type, step, and status.

#### 🔧 Request Parameters

| Field Name    | Type          | Required | Description                            |
|---------------|---------------|----------|----------------------------------------|
| `pagination`  | Pagination    | Yes      | Page size and token                    |
| `companyId`   | string        | Yes      | Company ID for access control          |
| `businessIds` | Array(string) | Yes      | List of business IDs to filter results |
| `filters`     | Filter        | No       | Filter options                         |

#### Filter Options:

- `abandonTime`: Filter by abandonment time range
- `leadTypes`: Filter by acquisition source
- `steps`: Filter by the step where booking was abandoned
- `statuses`: Filter by abandonment status

#### 📌 Return Value

| Field Name      | Type                    | Description                                   |
|-----------------|-------------------------|-----------------------------------------------|
| `nextPageToken` | string                  | Token for retrieving the next page of results |
| `bookings`      | Array(AbandonedBooking) | List of abandoned bookings matching criteria  |

#### ⚠️ Error Code

- `PERMISSION_DENIED`: Permission denied

---

### 3. Get Booking Availability (`GetBookingAvailability`)

- **Method**: `GetBookingAvailability`
- **HTTP Method**: POST
- **Path**: `/v1/online_booking/availability`

#### ✅ Functionality

Gets available dates and times for online booking based on business hours, staff availability, and other scheduling
constraints. This endpoint helps customers find suitable time slots when booking services online.

#### 🔧 Request Parameters

| Field Name   | Type   | Required | Description                                  |
|--------------|--------|----------|----------------------------------------------|
| `companyId`  | string | Yes      | Company identifier for multi-tenancy support |
| `businessId` | string | Yes      | Business location where services provided    |
| `filter`     | object | No       | Filter parameters for the availability check |

##### Filter Options:

- `startDate`: Start date for availability check (defaults to today). Maximum range between startDate and endDate is 3 months.
- `endDate`: End date for availability check (defaults to startDate). Maximum range between startDate and endDate is 3 months.
- `serviceIds`: Filter by specific service IDs
- `staffIds`: Filter by specific staff IDs
- `customerId`: Filter by specific customer ID
- `coordinate`: Filter by location coordinates (latitude and longitude)
- `zipcode`: Filter by postal code
- `pets`: Array of pet parameters including:
    - `id`: Pet ID (for existing pets)
    - `name`: Pet name
    - `type`: Pet type/species (see Pet Types table below)
    - `breed`: Pet breed
    - `birthday`: Pet birthday
    - `weight`: Pet weight
    - `staffId`: Preferred staff ID for this pet
    - `serviceIds`: Service IDs for this pet

###### Pet Types

When specifying the pet type in the `type` field, use one of the following values:

| Pet Type           | Description                                        |
|--------------------|----------------------------------------------------|
| `TYPE_UNSPECIFIED` | Unknown or unspecified pet type (default value)    |
| `OTHER`            | Pet type not listed in standard categories         |
| `DOG`              | Canine companion (e.g., Labrador, German Shepherd) |
| `CAT`              | Feline companion (domestic cat breeds)             |
| `BIRD`             | Avian pet (parrots, canaries, finches, etc.)       |
| `RABBIT`           | Domestic rabbit                                    |
| `GUINEA_PIG`       | Guinea pig                                         |
| `HORSE`            | Equine                                             |
| `HAMSTER`          | Hamster                                            |
| `RAT`              | Rat                                                |
| `MOUSE`            | Mouse                                              |
| `CHINCHILLA`       | Chinchilla                                         |

#### 📌 Return Value

| Field Name       | Type                      | Description                                    |
|------------------|---------------------------|------------------------------------------------|
| `availableDates` | Array(Date)               | List of dates with at least one available slot |
| `availability`   | Array(AvailabilityByDate) | Detailed availability information by date      |

##### AvailabilityByDate Object:

| Field Name | Type                     | Description                             |
|------------|--------------------------|-----------------------------------------|
| `date`     | Date                     | Date for which availability is provided |
| `staff`    | Array(StaffAvailability) | Staff and their available time slots    |

##### StaffAvailability Object:

| Field Name       | Type            | Description                                 |
|------------------|-----------------|---------------------------------------------|
| `staffId`        | string          | Unique identifier of the staff member       |
| `first_name`     | string          | First name of the staff member              |
| `last_name`      | string          | Last name of the staff member               |
| `availableSlots` | Array(Interval) | Available time slots (google.type.Interval) |

#### ⚠️ Error Codes

- `PERMISSION_DENIED`: Permission denied
- `INVALID_ARGUMENT`: Invalid request parameters

---

### 4. Lookup Franchise By Zipcode (`LookupFranchiseByZipcode`)

- **Method**: `LookupFranchiseByZipcode`
- **HTTP Method**: POST
- **Path**: `/v1/online_booking/franchise_lookup`

#### ✅ Functionality

Looks up franchise branch based on zipcode for territory routing. This endpoint allows routing customers to the correct
franchisee based on their service address zipcode against territory mappings. It is specifically designed for the Aussie
Pet Mobile franchise model where each franchisee can only serve users within their assigned territory.

#### 🔧 Request Parameters

| Field Name | Type   | Required | Description                                                  |
|------------|--------|----------|--------------------------------------------------------------|
| `zipcode`  | string | Yes      | The zipcode to look up for territory-based franchise routing |

#### 📌 Return Value

| Field Name       | Type    | Description                                                |
|------------------|---------|------------------------------------------------------------|
| `businessId`     | string  | The business identifier of the franchise branch            |
| `companyId`      | string  | The company identifier of the franchise branch             |
| `businessName`   | string  | The name of the franchise branch                           |
| `bookOnlineName` | string  | The book online name for the franchise branch              |
| `isEnable`       | boolean | Whether online booking is enabled for the franchise branch |

#### 📋 Business Rules

- If the zipcode belongs to an assigned territory or gray area, return information for the corresponding franchise
  branch
- If the zipcode belongs to an unassigned territory or doesn't belong to any territory, return empty results

#### ⚠️ Error Codes

- `PERMISSION_DENIED`: Permission denied
- `INVALID_ARGUMENT`: Invalid request parameters (e.g., missing or malformed zipcode)

---

## 🧪 5. Usage Examples

### Example 1: Get Abandoned Booking

```http
GET /v1/abandoned_bookings/abk_12345
```

**Response:**

```json
{
  "id": "abk_12345",
  "businessId": "biz_001",
  "customer": {
    "customerId": "cus_001",
    "email": "john.doe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "phoneNumber": "+12125551234"
  },
  "pets": [
    {
      "petServiceDetail": {
        "pet": {
          "id": "pet_001",
          "name": "Buddy"
        },
        "serviceDetails": [
          {
            "id": "svc_123",
            "name": "Premium Grooming"
          }
        ]
      },
      "emergencyContactName": "Jane Doe",
      "emergencyContactPhone": "+12125554321",
      "healthIssues": "Allergic to shampoos"
    }
  ],
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001"
  },
  "abandonStep": "SELECT_SERVICE",
  "abandonTime": "2024-08-15T10:00:00Z",
  "abandonStatus": "ABANDONED",
  "leadType": "NEW_VISITOR",
  "createdTime": "2024-08-15T09:50:00Z",
  "lastUpdatedTime": "2024-08-15T09:55:00Z"
}
```

### Example 2: List Abandoned Bookings

```http
POST /v1/abandoned_bookings:list
```

```json
{
  "companyId": "cmp_001",
  "businessIds": [
    "biz_001",
    "biz_002"
  ],
  "pagination": {
    "pageSize": 20,
    "pageToken": "1"
  },
  "filter": {
    "abandonTime": {
      "startTime": "2024-08-01T00:00:00Z",
      "endTime": "2024-08-07T23:59:59Z"
    },
    "leadTypes": [
      "NEW_VISITOR"
    ],
    "steps": [
      "SELECT_CARE_TYPE",
      "SELECT_PET"
    ],
    "statuses": [
      "ABANDONED"
    ]
  }
}
```

**Response:**

```json
{
  "nextPageToken": "",
  "bookings": [
    {
      "id": "abk_12345",
      "businessId": "biz_001",
      "customer": {
        ...
      },
      "pets": [
        ...
      ],
      "address": {
        ...
      },
      "abandonStep": "SELECT_CARE_TYPE",
      "abandonTime": "2024-08-05T14:30:00Z",
      "abandonStatus": "ABANDONED",
      "leadType": "NEW_VISITOR",
      "createdTime": "2024-08-05T14:25:00Z"
    }
  ]
}
```

### Example 3: Get Booking Availability

```http
POST /v1/online_booking/availability

```

**Request Body:**

```json
{
  "companyId": "cmp_001",
  "businessId": "biz_001"
}
```

**Response:**

```json
{
  "availableDates": [
    "2024-08-20",
    "2024-08-21",
    "2024-08-22"
  ],
  "availability": [
    {
      "date": "2024-08-20",
      "staff": [
        {
          "staffId": "stf_123",
          "firstName": "John",
          "lastName": "Doe",
          "availableSlots": [
            {
              "startTime": "2024-08-20T09:00:00Z",
              "endTime": "2024-08-20T09:30:00Z"
            },
            {
              "startTime": "2024-08-20T09:30:00Z",
              "endTime": "2024-08-20T10:00:00Z"
            },
            {
              "startTime": "2024-08-20T10:00:00Z",
              "endTime": "2024-08-20T10:30:00Z"
            },
            {
              "startTime": "2024-08-20T14:00:00Z",
              "endTime": "2024-08-20T14:30:00Z"
            },
            {
              "startTime": "2024-08-20T14:30:00Z",
              "endTime": "2024-08-20T15:00:00Z"
            }
          ]
        },
        {
          "staffId": "stf_456",
          "firstName": "John",
          "lastName": "Staff",
          "availableSlots": [
            {
              "startTime": "2024-08-20T10:00:00Z",
              "endTime": "2024-08-20T10:30:00Z"
            },
            {
              "startTime": "2024-08-20T10:30:00Z",
              "endTime": "2024-08-20T11:00:00Z"
            },
            {
              "startTime": "2024-08-20T11:00:00Z",
              "endTime": "2024-08-20T11:30:00Z"
            },
            {
              "startTime": "2024-08-20T15:00:00Z",
              "endTime": "2024-08-20T15:30:00Z"
            },
            {
              "startTime": "2024-08-20T15:30:00Z",
              "endTime": "2024-08-20T16:00:00Z"
            },
            {
              "startTime": "2024-08-20T16:00:00Z",
              "endTime": "2024-08-20T16:30:00Z"
            }
          ]
        }
      ]
    }
  ]
}
```

With filter parameters:

```http
POST /v1/online_booking/availability
```

**Request Body:**

```json
{
  "companyId": "cmp_001",
  "businessId": "biz_001",
  "filter": {
    "startDate": "2024-08-20",
    "endDate": "2024-08-22",
    "serviceIds": [
      "svc_123"
    ],
    "staffIds": [
      "stf_123"
    ],
    "customerId": "cus_001",
    "coordinate": {
      "latitude": 40.7128,
      "longitude": -74.0060
    },
    "zipcode": "10001"
  }
}
```

**Response:**

```json
{
  "availableDates": [
    "2024-08-20",
    "2024-08-21"
  ],
  "availability": [
    {
      "date": "2024-08-20",
      "staff": [
        {
          "staffId": "stf_123",
          "firstName": "John",
          "lastName": "Staff",
          "availableSlots": [
            {
              "startTime": "2024-08-20T09:00:00Z",
              "endTime": "2024-08-20T09:30:00Z"
            },
            {
              "startTime": "2024-08-20T09:30:00Z",
              "endTime": "2024-08-20T10:00:00Z"
            },
            {
              "startTime": "2024-08-20T10:00:00Z",
              "endTime": "2024-08-20T10:30:00Z"
            }
          ]
        }
      ]
    }
  ]
}
```

### Example 4: Lookup Franchise By Zipcode

```http
POST /v1/online_booking/business/info
```

**Request Body:**

```json
{
  "zipcode": "90210"
}
```

**Response (for a valid zipcode in assigned territory):**

```json
{
  "businessId": "biz_001",
  "companyId": "cmp_001",
  "businessName": "Beverly Hills Pet Grooming",
  "bookOnlineName": "beverly-hills",
  "isEnable": true
}
```

**Response (for an invalid or unassigned zipcode):**

```json
{
  "businessId": "",
  "companyId": "",
  "businessName": "",
  "bookOnlineName": "",
  "isEnable": false
}
```

---

## 📌 6. Common Error Codes

| Error Code          | Description                                     |
|---------------------|-------------------------------------------------|
| `NOT_FOUND`         | Specified booking or customer ID does not exist |
| `PERMISSION_DENIED` | Permission denied                               |
| `INVALID_ARGUMENT`  | Invalid request parameters                      |
| `INTERNAL`          | Internal server error                           |

---

## 📎 7. Related File References

- [appointment.proto](../moego/business/appointment/v1/appointment.proto)
- [pet_service_detail.proto](../moego/business/appointment/v1/pet_service_detail.proto)
- [address.proto](../../common/v1/address.proto)

---

## 🧪 8. Typical Use Cases

### ✅ Scenario: Analyze Abandoned Bookings

Use `ListAbandonedBookings` with filters to identify patterns in where customers drop off in the booking flow. This can
help optimize the user experience and improve conversion rates.

### ✅ Scenario: Follow-up on Abandoned Bookings

Use `GetAbandonedBooking` to review details of an abandoned booking and determine whether it can be recovered via
outreach or re-engagement campaigns.

### ✅ Scenario: Check Booking Availability

Use `GetBookingAvailability` to retrieve available dates and time slots for online booking. This allows customers to
find suitable appointment times based on business hours, staff availability, and service requirements. The API supports
filtering by specific services, staff members, or pets to provide more targeted availability information.

---