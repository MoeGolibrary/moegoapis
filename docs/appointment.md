# Appointment API Documentation (`moego.business.appointment.v1`)

## 📌 1. Functional Overview

The `Appointment` object represents a scheduled pet service booking. When an appointment is created, it will transition
through various states from unconfirmed to finished. Each appointment is associated with exactly one customer and one
business location, but can include multiple pets and services.

This API provides operations for:

- Retrieving individual appointments
- Listing appointments with filtering options (by time range, status, or last update)
- Creating new appointments with custom services and pets
- Rescheduling existing appointments
- Canceling appointments
- Managing grooming reports associated with completed services

---

## 🎯 2. Design Goals

- **Centralized Management**: Provides a unified interface for managing all aspects of pet service appointments.
- **Flexible Scheduling**: Supports complex scheduling scenarios including multi-pet bookings and staff assignments.
- **Payment Tracking**: Maintains payment state and amount tracking throughout the appointment lifecycle.
- **Service History**: Preserves detailed records of services performed, including grooming reports.
- **Audit Trail**: Tracks creation and modification times along with responsible staff members.

---

## 🧩 3. Core Concepts

### 1. Appointment

Represents a scheduled pet service booking. An appointment can include multiple pets and services, and tracks the entire
lifecycle from creation to completion.

| Field Name            | Type                    | Description                                                       |
|-----------------------|-------------------------|-------------------------------------------------------------------|
| `id`                  | string                  | Unique identifier of the appointment                              |
| `business_id`         | string                  | Business location where the service will be performed             |
| `customer_id`         | string                  | Customer who booked the appointment                               |
| `address`             | Address                 | Service location details (required for home service appointments) |
| `duration`            | Interval                | Start and end time of the appointment                             |
| `pet_service_details` | Array(PetServiceDetail) | List of services booked for each pet                              |
| `status`              | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.       |
| `ticket_comment`      | string                  | Optional notes about the appointment                              |
| `color_code`          | string                  | UI display color in hex format                                    |
| `order_id`            | string                  | Identifier of the associated payment order                        |
| `total_amount`        | Money                   | Total cost for all services                                       |
| `paid_amount`         | Money                   | Amount received from customer                                     |
| `refund_amount`       | Money                   | Amount returned to customer                                       |
| `payment_status`      | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`              |
| `created_by`          | string                  | Identifier of the appointment creator                             |
| `created_time`        | timestamp               | When the appointment was created                                  |
| `last_updated_by`     | string                  | Identifier of the last modifier                                   |
| `last_updated_time`   | timestamp               | When the appointment was last modified                            |
| `check_in_time`       | timestamp               | When the customer arrived with their pet                          |
| `check_out_time`      | timestamp               | When the service was completed and the pet picked up              |
| `booking_request_id`  | string (optional)       | The booking request ID associated with this appointment           |

#### Enum Definitions

##### `Appointment.Status`

- `STATUS_UNSPECIFIED`
- `UNCONFIRMED`: Initial state. Appointment created but pending business confirmation.
- `CONFIRMED`: Business has accepted and allocated resources.
- `CHECKED_IN`: Customer has arrived and service is in progress.
- `READY`: Service completed, awaiting pet pickup.
- `FINISHED`: Final state. Pet picked up and service completed.
- `CANCELED`: Final state. Appointment terminated before completion.

##### `Appointment.PaymentStatus`

- `PAYMENT_STATUS_UNSPECIFIED`
- `UNPAID`: No payment received.
- `PARTIAL_PAID`: Deposit or partial payment received.
- `FULL_PAID`: Complete payment received.

---

### 2. GroomingReport

Represents a grooming report generated after a pet grooming service. It contains information about the pet, report
status, sending methods and logs, and is used for customer communication and internal record keeping.

| Field Name                 | Type              | Description                                                                    |
|----------------------------|-------------------|--------------------------------------------------------------------------------|
| `id`                       | string            | Unique identifier of the grooming report                                       |
| `appointment_id`           | string            | Associated appointment ID                                                      |
| `business_id`              | string            | Business location ID                                                           |
| `customer_id`              | string            | Owner customer ID                                                              |
| `pet_id`                   | string            | Pet ID                                                                         |
| `template_json`            | string            | JSON-formatted template content                                                |
| `content_json`             | string            | JSON-formatted actual content filled in based on the template and service data |
| `status`                   | enum(Status)      | Current status: `CREATED`, `DRAFT`, `SUBMITTED`, `SENT`                        |
| `report_url`               | string            | URL where the report can be accessed                                           |
| `send_type`                | enum(SendType)    | How the report was sent: `MANUAL`, `AUTO`                                      |
| `template_publish_time`    | timestamp         | Timestamp when the current template was published                              |
| `sending_method_list`      | Array(SendMethod) | Available sending methods configured for this report                           |
| `last_sending_method_list` | Array(SendMethod) | Sending methods selected during the last send attempt                          |
| `send_log`                 | Array(SendLog)    | Log entries for each send attempt made for this report                         |

#### Enum Definitions

##### `GroomingReport.Status`

- `STATUS_UNSPECIFIED`
- `CREATED`: Report is created but not submitted (draft).
- `DRAFT`: The report is created but not submitted.
- `SUBMITTED`: Report has been submitted and is ready to be sent.
- `SENT`: Report has been successfully sent to the customer.

##### `GroomingReport.SendType`

- `SEND_TYPE_UNSPECIFIED`
- `MANUAL`: Report was manually sent by a user.
- `AUTO`: Report was automatically sent by the system.

##### `GroomingReport.SendMethod`

- `SEND_METHOD_UNSPECIFIED`
- `SMS`: Delivery via SMS.
- `EMAIL`: Delivery via Email.

##### `GroomingReport.SendStatus`

- `SEND_STATUS_UNSPECIFIED`
- `SEND_SUCCESS`: Send attempt was successful.
- `SEND_FAIL`: Send attempt failed.

##### `GroomingReport.SendLog`

- `send_method`: Method used for this send attempt.
- `send_status`: Result of this send attempt.
- `sent_time`: Timestamp when the send attempt occurred.
- `sender_id`: Identifier of the staff or system that initiated the send.
- `error_code` *(optional)*: Error code if the send failed.
- `error_msg` *(optional)*: Error message if the send failed.

---

### 3. PetServiceDetail

Represents the services booked for a specific pet in an appointment. Each appointment can include multiple pets, and
each pet can receive multiple services.

| Field Name        | Type                 | Description                             |
|-------------------|----------------------|-----------------------------------------|
| `pet`             | Pet                  | The pet receiving the services          |
| `service_details` | Array(ServiceDetail) | List of services scheduled for this pet |

---

### 4. ServiceDetail

Represents a specific service booked for a pet. It contains all the information needed to perform the service, including
timing, pricing, staff assignments, and service-specific parameters.

| Field Name              | Type          | Description                                                    |
|-------------------------|---------------|----------------------------------------------------------------|
| `id`                    | string        | Unique identifier for the service                              |
| `name`                  | string        | Display name of the service                                    |
| `service_item_type`     | ItemType      | Category of service item                                       |
| `category`              | string        | Service category for grouping and filtering                    |
| `price`                 | Money         | Price for this service instance                                |
| `duration`              | Interval      | Scheduled time window for the service                          |
| `staff_ids`             | Array(string) | IDs of staff members assigned to this service                  |
| `service_type`          | Type          | Type of service being provided                                 |
| `service_time`          | integer       | Expected duration of the service in minutes                    |
| `pet_service_detail_id` | string        | Unique identifier for this specific pet's service booking      |
| `lodging_unit_id`       | string        | Unique identifier for this specific pet's service lodging unit |

---

### 5. AppointmentNote

Represents notes associated with an appointment, used for internal communication or customer information tracking.

| Field Name          | Type       | Description                                                 |
|---------------------|------------|-------------------------------------------------------------|
| `id`                | string     | Unique identifier of the note                               |
| `business_id`       | string     | Business location ID                                        |
| `customer_id`       | string     | Customer ID                                                 |
| `company_id`        | string     | Company ID                                                  |
| `appointment_id`    | string     | Related appointment ID                                      |
| `note`              | string     | The content of the note                                     |
| `type`              | enum(Type) | Note type: `ALERT_NOTES`, `COMMENT`, `CANCEL`, `ADDITIONAL` |
| `created_time`      | timestamp  | When the note was created                                   |
| `last_updated_time` | timestamp  | When the note was last modified                             |
| `created_by`        | string     | Staff member who created the note                           |
| `last_updated_by`   | string     | Staff member who last updated the note                      |

#### Enum Definitions

##### `AppointmentNote.Type`

- `APPOINTMENT_NOTE_TYPE_UNSPECIFIED`
- `ALERT_NOTES`: Alert or reminder note
- `COMMENT`: General comment or description
- `CANCEL`: Cancellation reason
- `ADDITIONAL`: Additional information

---

## 📦 4. API Interface Descriptions

### 1. Get Appointment (`GetAppointment`)

- **Method**: `GetAppointment`
- **HTTP Method**: GET
- **Path**: `/v1/appointments/{id}`

#### ✅ Functionality:

Retrieves a single appointment by its ID.

#### 🎯 Use Cases:

- View details of an existing appointment.
- Verify appointment configuration during debugging.

#### 🔧 Request Parameters:

| Field Name    | Type   | Required | Description                    |
|---------------|--------|----------|--------------------------------|
| `id`          | string | Yes      | Appointment ID to retrieve     |
| `business_id` | string | Yes      | Business ID for access control |

#### 📌 Return Value:

Returns the complete `Appointment` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified appointment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 2. List Appointments (`ListAppointments`)

- **Method**: `ListAppointments`
- **HTTP Method**: POST
- **Path**: `/v1/appointments:list`

#### ✅ Functionality:

Lists appointments matching specified criteria, including company ID and optional filters.

#### 🎯 Use Cases:

- View all appointments under a company.
- Audit or debug appointment configurations.

#### 🔧 Request Parameters:

| Field Name     | Type          | Required | Description                                    |
|----------------|---------------|----------|------------------------------------------------|
| `pagination`   | Pagination    | Yes      | Page size and token                            |
| `company_id`   | string        | Yes      | Company ID for access control                  |
| `business_ids` | Array(string) | Yes      | List of business IDs to filter appointments by |
| `filters`      | Filter        | No       | Filter options                                 |

#### Filter Options:

- `start_time`: Filter by appointment start time range.
- `end_time`: Filter by appointment end time range.
- `last_updated_time`: Filter by last update time range.
- `statuses`: Filter by appointment status.
- `check_in_time`: Filter by check-in time range.
- `check_out_time`: Filter by check-out time range.
- `service_types`: Filter by service item type.

#### 📌 Return Value:

Returns a list of appointments ordered by start time.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 3. Create Appointment (`CreateAppointment`)

- **Method**: `CreateAppointment`
- **HTTP Method**: POST
- **Path**: `/v1/appointments`

#### ✅ Functionality:

Creates a new appointment with services for one or more pets.

#### 🎯 Use Cases:

- Schedule a new appointment for a customer.
- Book multiple services across different pets in one appointment.

#### 🔧 Request Parameters:

| Field Name        | Type              | Required | Description                                                                              |
|-------------------|-------------------|----------|------------------------------------------------------------------------------------------|
| `business_id`     | string            | Yes      | Business location ID where services will be provided                                     |
| `customer_id`     | string            | Yes      | Customer ID who is booking the appointment                                               |
| `pet_services`    | Array(PetService) | Yes      | Services requested for each pet                                                          |
| `ignore_conflict` | bool              | No       | Whether to ignore scheduling conflicts when creating an appointment. Defaults to `true`. |

> ** Note ** : When set to 'true', the system will ignore time conflicts and directly create reservations. This field is
> mainly used to support mandatory appointments in certain special scenarios (such as manual intervention in the
> background).

#### 📌 Return Value:

Returns the newly created `Appointment` object.

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields are missing or invalid.
- `PERMISSION_DENIED`: Permission denied.

---

### 4. Reschedule Appointment (`RescheduleAppointment`)

- **Method**: `RescheduleAppointment`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/{id}:reschedule`

#### ✅ Functionality:

Updates the appointment’s time slot. All services within the appointment will be moved to the new time.

#### 🎯 Use Cases:

- Move an appointment to a different time.
- Update appointment schedule due to staff availability changes.

#### 🔧 Request Parameters:

| Field Name    | Type     | Required | Description                         |
|---------------|----------|----------|-------------------------------------|
| `id`          | string   | Yes      | Appointment ID to reschedule        |
| `business_id` | string   | Yes      | Business ID for access control      |
| `duration`    | Interval | Yes      | New time window for the appointment |

#### 📌 Return Value:

Returns the updated `Appointment` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified appointment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 5. Cancel Appointment (`CancelAppointment`)

- **Method**: `CancelAppointment`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/{id}:cancel`

#### ✅ Functionality:

Sets the appointment status to `CANCELED`. This action cannot be undone.

#### 🎯 Use Cases:

- Cancel an appointment before it starts.
- Handle no-shows or customer cancellations.

#### 🔧 Request Parameters:

| Field Name    | Type   | Required | Description                    |
|---------------|--------|----------|--------------------------------|
| `id`          | string | Yes      | Appointment ID to cancel       |
| `business_id` | string | Yes      | Business ID for access control |

#### 📌 Return Value:

Returns the canceled `Appointment` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified appointment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 6. Check Appointment Availability (`CheckAppointmentAvailability`)

- **Method**: `CheckAppointmentAvailability`
- **HTTP Method**: POST
- **Path**: `/v1/appointments:check`

#### ✅ Functionality:

Performs pre-checks before creating or rescheduling an appointment, including checking for time conflicts and business
closed dates.

This method helps clients determine if it's safe to proceed with creating an appointment without causing scheduling
issues.

#### 🎯 Use Cases:

- Validate that there are no conflicting appointments.
- Ensure the selected date does not fall on a business holiday or closed day.
- Prevent duplicate or overlapping bookings.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description           |
|--------------------|---------------|----------|-----------------------|
| `business_id`      | string        | Yes      | Business location ID  |
| `date_range`       | Interval      | Yes      | Appointment time slot |
| `pet_ids`          | Array(string) | No       | List of pets ids      |
| `customer_id`      | string        | No       | Customer ID           |
| `appointment_id`   | string        | No       | Appointment ID        |
| `lodging_unit_ids` | Array(string) | No       | Lodging  Unit ID      |

#### 📌 Return Value:

Returns a `CheckAppointmentAvailabilityResponse` object containing results from:

- **Appointment Date Conflict Check**: List of conflicting appointments (if any).
- **Business Closed Date Check**: List of closed dates during the requested period.

If both checks return empty results, it means the appointment can be safely created.

##### CheckAppointmentAvailabilityResponse

Represents the result of a pre-check operation before creating or rescheduling an appointment. It includes two main
components:

###### Fields:

| Field Name                    | Type                                 | Description                                                                 |
|-------------------------------|--------------------------------------|-----------------------------------------------------------------------------|
| `appointment_conflict_check`  | `AppointmentDateConflictCheckResult` | Contains information about conflicting appointments for each pet.           |
| `business_closed_date_check`  | `BusinessClosedDateCheckResult`      | Contains information about business closed dates during the requested time. |
| `lodging_over_capacity_check` | `LodgingOverCapacityCheckResult`     | Contains information about lodging over capacity during the requested time. |

---

###### 1. `AppointmentDateConflictCheckResult`

| Field Name  | Type                             | Description                                        |
|-------------|----------------------------------|----------------------------------------------------|
| `conflicts` | Array(`PetAppointmentsOverview`) | A list of pets and their conflicting appointments. |

---

###### 2. `PetAppointmentsOverview`

| Field Name     | Type                 | Description                                      |
|----------------|----------------------|--------------------------------------------------|
| `pet`          | `Pet`                | The pet involved in the conflicting appointment. |
| `appointments` | Array(`Appointment`) | A list of conflicting appointments for this pet. |

---

###### 3. `BusinessClosedDateCheckResult`

| Field Name     | Type                      | Description                                                  |
|----------------|---------------------------|--------------------------------------------------------------|
| `closed_dates` | Array(`google.type.Date`) | A list of business closed dates during the requested period. |

---

###### 4. `LodgingOverCapacityCheckResult`

This structure is used to indicate whether lodging units are over capacity during the specified time period. When an
appointment involves pet boarding services, the system checks if the business has exceeded its lodging capacity during
the requested time.

| Field Name | Type                                       | Description                          |
|------------|--------------------------------------------|--------------------------------------|
| `lodgings` | Array(`moego.business.setting.v1.Lodging`) | List of lodging that exceed capacity |

---

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Missing required fields in request.
- `PERMISSION_DENIED`: Permission denied.

---

### 7. List Grooming Reports (`ListGroomingReports`)

- **Method**: `ListGroomingReports`
- **HTTP Method**: POST
- **Path**: `/v1/grooming_reports:list`

#### ✅ Functionality:

Lists grooming reports for specified appointments.

#### 🎯 Use Cases:

- Retrieve grooming reports for reporting or customer communication.
- Review historical grooming data for pets.

#### 🔧 Request Parameters:

| Field Name        | Type          | Required | Description                    |
|-------------------|---------------|----------|--------------------------------|
| `appointment_ids` | Array(string) | Yes      | Max 500 appointment IDs        |
| `business_id`     | string        | Yes      | Business ID for access control |

#### 📌 Return Value:

Returns a list of `GroomingReport` objects.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 8. Create Appointment Note (`CreateAppointmentNote`)

- **Method**: `CreateAppointmentNote`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/notes`

#### ✅ Functionality:

Creates a new note associated with a specific appointment.

#### 🎯 Use Cases:

- Add alerts or comments related to an appointment.
- Record cancellation reasons or additional notes from customers.

#### 🔧 Request Parameters:

| Field Name       | Type                       | Required | Description                               |
|------------------|----------------------------|----------|-------------------------------------------|
| `appointment_id` | string                     | Yes      | The appointment ID to associate this note |
| `note`           | string                     | Yes      | Content of the note                       |
| `type`           | enum(AppointmentNote.Type) | Yes      | Type of note                              |

#### 📌 Return Value:

Returns the created `AppointmentNote` object.

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Missing required fields.
- `PERMISSION_DENIED`: Permission denied.

---

### 9. Update Appointment Note (`UpdateAppointmentNote`)

- **Method**: `UpdateAppointmentNote`
- **HTTP Method**: PUT
- **Path**: `/v1/appointments/notes/{id}`

#### ✅ Functionality:

Updates an existing appointment note by its ID.

#### 🎯 Use Cases:

- Modify previously added notes.
- Correct errors in alert messages or comments.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description                       |
|------------|--------|----------|-----------------------------------|
| `id`       | string | Yes      | The unique identifier of the note |
| `note`     | string | Yes      | Updated content of the note       |

#### 📌 Return Value:

Returns the updated `AppointmentNote` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified note ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 10. List Appointment Notes (`ListAppointmentNotes`)

- **Method**: `ListAppointmentNotes`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/notes:list`

#### ✅ Functionality:

Lists all notes matching specified criteria, including company, customer, and optional filters like note type or
appointment ID.

#### 🎯 Use Cases:

- Retrieve notes for reporting or audit purposes.
- View historical notes related to specific appointments.

#### 🔧 Request Parameters:

| Field Name       | Type                       | Required | Description                       |
|------------------|----------------------------|----------|-----------------------------------|
| `appointment_id` | string                     | Yes      | Appointment ID for access control |
| `type`           | enum(AppointmentNote.Type) | Yes      | Filter by note type               |

#### 📌 Return Value:

Returns a list of `AppointmentNote` objects.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 5. Usage Examples

### Example 1: Get Appointment

```http
GET /v1/appointments/12345?business_id=biz_001
```

**Response:**

```json
{
  "id": "apt_12345",
  "business_id": "biz_001",
  "customer_id": "cus_001",
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001"
  },
  "duration": {
    "start_time": "2024-08-15T10:00:00Z",
    "end_time": "2024-08-15T12:00:00Z"
  },
  "pet_service_details": [
    {
      "pet": {
        "id": "pet_001",
        "name": "Buddy"
      },
      "service_details": [
        {
          "id": "svc_123",
          "name": "Premium Grooming",
          "price": {
            "currency_code": "USD",
            "units": 50,
            "nanos": 0
          },
          "duration": {
            "start_time": "2024-08-15T10:00:00Z",
            "end_time": "2024-08-15T11:00:00Z"
          }
        }
      ]
    }
  ],
  "status": "CONFIRMED",
  "payment_status": "FULL_PAID",
  "total_amount": {
    "currency_code": "USD",
    "units": 75,
    "nanos": 0
  },
  "paid_amount": {
    "currency_code": "USD",
    "units": 75,
    "nanos": 0
  },
  "created_time": "2024-08-10T09:00:00Z",
  "last_updated_time": "2024-08-12T14:30:00Z"
}
```

### Example 2: Create Appointment

```json
{
  "business_id": "biz_001",
  "customer_id": "cus_001",
  "pet_services": [
    {
      "pet_id": "pet_001",
      "services": [
        {
          "id": "svc_grooming",
          "duration": {
            "start_time": "2024-08-15T10:00:00Z",
            "end_time": "2024-08-15T11:00:00Z"
          },
          "staff_ids": [
            "staff_001"
          ]
        }
      ]
    }
  ]
}
```

### Example 3: Reschedule Appointment

```http
POST /v1/appointments/apt_12345:reschedule
```

```json
{
  "business_id": "biz_001",
  "duration": {
    "start_time": "2024-08-16T11:00:00Z",
    "end_time": "2024-08-16T13:00:00Z"
  }
}
```

### Example 4: Check Create Appointment

```http
POST /v1/appointments:check
```

```json
{
  "business_id": "biz_001",
  "customer_id": "cus_001",
  "pet_services": [
    {
      "pet_id": "pet_001",
      "services": [
        {
          "id": "svc_grooming",
          "duration": {
            "start_time": "2024-08-15T10:00:00Z",
            "end_time": "2024-08-15T11:00:00Z"
          },
          "staff_ids": [
            "staff_001"
          ]
        }
      ]
    }
  ]
}
```

**Response:**

```json
{
  "appointment_conflict_check": {
    "conflicts": [
      {
        "pet": {
          "id": "pet_001",
          "name": "Buddy"
        },
        "appointments": [
          {
            "id": "apt_98765",
            "business_id": "biz_001",
            "customer_id": "cus_002",
            "duration": {
              "start_time": "2024-08-15T10:30:00Z",
              "end_time": "2024-08-15T11:30:00Z"
            },
            "status": "CONFIRMED"
          }
        ]
      }
    ]
  },
  "business_closed_date_check": {
    "closed_dates": []
  }
}

```

> The above response indicates that there is currently a confirmed appointment that overlaps with the new appointment
> time, but there is no conflict with the business closure date.


---

## 📌 6. Common Error Codes

| Error Code          | Description                                                               |
|---------------------|---------------------------------------------------------------------------|
| `NOT_FOUND`         | Appointment or customer ID does not exist                                 |
| `PERMISSION_DENIED` | Current user has no access rights                                         |
| `INVALID_ARGUMENT`  | Invalid request parameters                                                |
| `INTERNAL`          | Internal server error                                                     |
| `CONFLICT`          | Detected scheduling conflict (returned by `CheckAppointmentAvailability`) |

---

## 📎 7. Related File References

- [appointment_service.proto](../moego/business/appointment/v1/appointment_service.proto)
- [appointment.proto](../moego/business/appointment/v1/appointment.proto)
- [grooming_report.proto](../moego/business/appointment/v1/grooming_report.proto)
- [pet_service_detail.proto](../moego/business/appointment/v1/pet_service_detail.proto)


