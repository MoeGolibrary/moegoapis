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

| Field Name            | Type                    | Description                                                       |
|-----------------------|-------------------------|-------------------------------------------------------------------|
| `id`                  | string                  | Unique identifier (e.g., "bkg_12345")                             |
| `business_id`         | string                  | ID of the business location where service will be performed       |
| `customer_id`         | string                  | ID of the customer initiating the booking                         |
| `address`             | Address                 | Service location details (required for home service appointments) |
| `duration`            | Interval                | Requested time window for the service                             |
| `pet_service_details` | Array(PetServiceDetail) | List of services requested for each pet                           |
| `status`              | enum(Status)            | Current state: `NORMAL`, `IN_WAIT_LIST`, `ABANDONED`              |
| `color_code`          | string                  | UI display color in hex format                                    |
| `created_time`        | timestamp               | When the booking was created                                      |

#### Enum Definitions

##### `OnlineBooking.Status`

- `STATUS_UNSPECIFIED`: Invalid value used only as default
- `NORMAL`: Booking is being processed normally
- `IN_WAIT_LIST`: Booking is on a waiting list due to unavailability
- `ABANDONED`: Booking was not completed by the customer

---

### 2. AbandonedBooking

Captures information about a booking that was not completed by the customer.

| Field Name          | Type                       | Description                                           |
|---------------------|----------------------------|-------------------------------------------------------|
| `id`                | string                     | Unique identifier                                     |
| `business_id`       | string                     | Business location ID                                  |
| `customer`          | AbandonedBookingCustomer   | Customer details                                      |
| `pets`              | Array(AbandonedBookingPet) | Pets and related services involved                    |
| `address`           | Address                    | Service location details                              |
| `abandon_step`      | Step                       | Stage at which the booking was abandoned              |
| `abandon_time`      | timestamp                  | Time when the booking was abandoned                   |
| `abandon_status`    | Status                     | Current status: `ABANDONED`, `CONTACTED`, `RECOVERED` |
| `last_texted_time`  | timestamp                  | Last SMS contact attempt regarding this booking       |
| `last_emailed_time` | timestamp                  | Last email sent about this booking                    |
| `recovery_type`     | RecoverType                | How the booking was recovered (if applicable)         |
| `recovery_time`     | timestamp                  | Time when the booking was recovered                   |
| `staff_id`          | string                     | Staff member who handled recovery                     |
| `appointment_id`    | string                     | Appointment ID if booking was converted               |
| `appointment_time`  | timestamp                  | Scheduled appointment time if one was set             |
| `lead_type`         | LeadType                   | How the customer was acquired                         |
| `create_time`       | timestamp                  | When this record was created                          |
| `update_time`       | timestamp                  | When this record was last updated                     |
| `additional_note`   | string                     | Notes added by staff or system                        |
| `care_type`         | ItemType                   | Type of service selected by the customer              |

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

| Field Name                 | Type                   | Description                                                     |
|----------------------------|------------------------|-----------------------------------------------------------------|
| `customer_id`              | string                 | Customer’s unique ID                                            |
| `email`                    | string                 | Email address                                                   |
| `first_name`               | string                 | First name                                                      |
| `last_name`                | string                 | Last name                                                       |
| `phone_number`             | string                 | Phone number                                                    |
| `referer`                  | string                 | Source URL that referred the user                               |
| `referral_source`          | string                 | Referral source                                                 |
| `preferred_groomer_id`     | string                 | Preferred groomer ID                                            |
| `preferred_frequency_day`  | int32                  | Preferred frequency in days                                     |
| `preferred_frequency_type` | PreferredFrequencyType | Frequency type (daily, weekly, monthly)                         |
| `preferred_days`           | Array(int32)           | Preferred days of the week                                      |
| `preferred_time`           | Array(int32)           | Preferred times of the day                                      |
| `question_answer_list`     | Array(QuestionAnswer)  | Key-value pairs capturing user input during the booking process |

---

### 4. AbandonedBookingPet

Represents a pet involved in an abandoned booking, along with service and emergency info.

| Field Name                | Type                  | Description                        |
|---------------------------|-----------------------|------------------------------------|
| `pet_service_detail`      | PetServiceDetail      | Services booked for this pet       |
| `emergency_contact_name`  | string                | Emergency contact name             |
| `emergency_contact_phone` | string                | Emergency contact phone number     |
| `health_issues`           | string                | Known health issues or conditions  |
| `question_answer_list`    | Array(QuestionAnswer) | Pet-specific question-answer pairs |

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

Returns the complete `AbandonedBooking` object.

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

| Field Name     | Type          | Required | Description                            |
|----------------|---------------|----------|----------------------------------------|
| `pagination`   | Pagination    | Yes      | Page size and token                    |
| `company_id`   | string        | Yes      | Company ID for access control          |
| `business_ids` | Array(string) | Yes      | List of business IDs to filter results |

#### Filter Options:

- `abandon_time`: Filter by abandonment time range
- `lead_types`: Filter by acquisition source
- `steps`: Filter by the step where booking was abandoned
- `statuses`: Filter by abandonment status

#### 📌 Return Value

Returns a paginated list of `AbandonedBooking` objects.

#### ⚠️ Error Code

- `PERMISSION_DENIED`: Permission denied

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
  "business_id": "biz_001",
  "customer": {
    "customer_id": "cus_001",
    "email": "john.doe@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "+12125551234"
  },
  "pets": [
    {
      "pet_service_detail": {
        "pet": {
          "id": "pet_001",
          "name": "Buddy"
        },
        "service_details": [
          {
            "id": "svc_123",
            "name": "Premium Grooming"
          }
        ]
      },
      "emergency_contact_name": "Jane Doe",
      "emergency_contact_phone": "+12125554321",
      "health_issues": "Allergic to shampoos"
    }
  ],
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001"
  },
  "abandon_step": "SELECT_SERVICE",
  "abandon_time": "2024-08-15T10:00:00Z",
  "abandon_status": "ABANDONED",
  "lead_type": "NEW_VISITOR",
  "create_time": "2024-08-15T09:50:00Z",
  "update_time": "2024-08-15T09:55:00Z"
}
```

### Example 2: List Abandoned Bookings

```http
POST /v1/abandoned_bookings:list
```

```json
{
  "company_id": "cmp_001",
  "business_ids": [
    "biz_001",
    "biz_002"
  ],
  "pagination": {
    "page_size": 50
  },
  "filter": {
    "abandon_time": {
      "start_time": "2024-08-01T00:00:00Z",
      "end_time": "2024-08-07T23:59:59Z"
    },
    "lead_types": [
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
  "next_page_token": "",
  "bookings": [
    {
      "id": "abk_12345",
      "business_id": "biz_001",
      "customer": {
        ...
      },
      "pets": [
        ...
      ],
      "address": {
        ...
      },
      "abandon_step": "SELECT_CARE_TYPE",
      "abandon_time": "2024-08-05T14:30:00Z",
      "abandon_status": "ABANDONED",
      "lead_type": "NEW_VISITOR",
      "create_time": "2024-08-05T14:25:00Z"
    }
  ]
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

---

## 📝 Summary

This documentation covers the core types and APIs for managing online and abandoned bookings in the Moego platform. It
enables businesses to track and recover lost opportunities, understand customer behavior, and improve conversion rates
through targeted follow-ups.

If you'd like, I can also generate OpenAPI/Swagger specs or Markdown tables for proto messages to assist developers in
understanding field usage and constraints.