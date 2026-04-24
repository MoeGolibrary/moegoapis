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

| Field Name          | Type                    | Description                                                                             |
|---------------------|-------------------------|-----------------------------------------------------------------------------------------|
| `id`                | string                  | Unique identifier of the appointment, obfuscated ID string                              |
| `businessId`        | string                  | Business location where the service will be performed, obfuscated ID string             |
| `customerId`        | string                  | Customer who booked the appointment, obfuscated ID string                               |
| `address`           | Address                 | Service location details (required for home service appointments)                       |
| `duration`          | Interval                | Start and end time of the appointment                                                   |
| `petServiceDetails` | Array(PetServiceDetail) | List of services booked for each pet                                                    |
| `status`            | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.                             |
| `ticketComment`     | string                  | Optional notes about the appointment                                                    |
| `colorCode`         | string                  | UI display color in hex format                                                          |
| `orderId`           | string                  | Identifier of the associated payment order, obfuscated ID string                        |
| `totalAmount`       | Money                   | Total cost for all services                                                             |
| `paidAmount`        | Money                   | Amount received from customer                                                           |
| `refundAmount`      | Money                   | Amount returned to customer                                                             |
| `paymentStatus`     | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`, `PARTIAL_REFUNDED`, `FULL_REFUNDED` |
| `createdBy`         | string                  | Identifier of the appointment creator, obfuscated ID string                             |
| `createdTime`       | timestamp               | When the appointment was created                                                        |
| `lastUpdatedBy`     | string                  | Identifier of the last modifier, obfuscated ID string                                   |
| `lastUpdatedTime`   | timestamp               | When the appointment was last modified                                                  |
| `checkInTime`       | timestamp               | When the customer arrived with their pet                                                |
| `checkOutTime`      | timestamp               | When the service was completed and the pet picked up                                    |
| `bookingRequestId`  | string (optional)       | The booking request ID associated with this appointment, obfuscated ID string           |
| `rawId`             | int64 (optional)        | The raw numeric ID. Restricted and only populated for authorized users.                 |
| `confirmedTime`     | timestamp (optional)    | When the appointment was confirmed                                                      |
| `readyTime`         | timestamp (optional)    | When the appointment was ready                                                          |
| `canceledTime`      | timestamp (optional)    | When the appointment was canceled                                                       |
| `noShow`            | bool                    | bool, Indicates if the appointment was marked as no-show.                               |
| `noShowFee`         | Money                   | Amount charged for no-show                                                              |

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
- `PARTIAL_REFUNDED`: Payment partially refunded.
- `FULL_REFUNDED`: Payment fully refunded.

---

### 2. GroomingReport

Represents a grooming report generated after a pet grooming service. It contains information about the pet, report
status, sending methods and logs, and is used for customer communication and internal record keeping.

| Field Name              | Type              | Description                                                                    |
|-------------------------|-------------------|--------------------------------------------------------------------------------|
| `id`                    | string            | Unique identifier of the grooming report, obfuscated ID string                 |
| `appointmentId`         | string            | Associated appointment ID, obfuscated ID string                                |
| `businessId`            | string            | Business location ID, obfuscated ID string                                     |
| `customerId`            | string            | Owner customer ID, obfuscated ID string                                        |
| `petId`                 | string            | Pet ID, obfuscated ID string                                                   |
| `templateJson`          | string            | JSON-formatted template content                                                |
| `contentJson`           | string            | JSON-formatted actual content filled in based on the template and service data |
| `status`                | enum(Status)      | Current status: `CREATED`, `DRAFT`, `SUBMITTED`, `SENT`                        |
| `reportUrl`             | string            | URL where the report can be accessed                                           |
| `sendType`              | enum(SendType)    | How the report was sent: `MANUAL`, `AUTO`                                      |
| `templatePublishTime`   | timestamp         | Timestamp when the current template was published                              |
| `sendingMethodList`     | Array(SendMethod) | Available sending methods configured for this report                           |
| `lastSendingMethodList` | Array(SendMethod) | Sending methods selected during the last send attempt                          |
| `sendLog`               | Array(SendLog)    | Log entries for each send attempt made for this report                         |

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

- `sendMethod`: Method used for this send attempt.
- `sendStatus`: Result of this send attempt.
- `sentTime`: Timestamp when the send attempt occurred.
- `senderId`: Identifier of the staff or system that initiated the send.
- `errorCode` *(optional)*: Error code if the send failed.
- `errorMsg` *(optional)*: Error message if the send failed.

---

### 3. PetServiceDetail

Represents the services booked for a specific pet in an appointment. Each appointment can include multiple pets, and
each pet can receive multiple services.

| Field Name           | Type                   | Description                                                  |
|----------------------|------------------------|--------------------------------------------------------------|
| `pet`                | Pet                    | The pet receiving the services                               |
| `serviceDetails`     | Array(ServiceDetail)   | List of services scheduled for this pet                      |
| `evaluationDetails`  | Array(EvaluationDetail) | List of evaluation results associated with this pet (if any) |

---

### 4. EvaluationDetail

Represents an evaluation result associated with a pet in an appointment. This is typically used to indicate whether the
pet has passed a required evaluation for certain services.

| Field Name | Type                   | Description                                                                 |
|------------|------------------------|-----------------------------------------------------------------------------|
| `id`       | string                 | Unique identifier for the evaluation, obfuscated ID string                  |
| `name`     | string                 | Display name of the evaluation                                              |
| `status`   | enum(EvaluationStatus) | Evaluation status: `PASS` or `FAIL`                                         |

#### Enum Definitions

##### `EvaluationDetail.EvaluationStatus`

- `EVALUATION_STATUS_UNSPECIFIED`
- `PASS`: Pet has passed evaluation and can receive services
- `FAIL`: Pet requires modifications to receive services

---

### 5. ServiceDetail

Represents a specific service booked for a pet. It contains all the information needed to perform the service, including
timing, pricing, staff assignments, and service-specific parameters.

| Field Name           | Type          | Description                                                    |
|----------------------|---------------|----------------------------------------------------------------|
| `id`                 | string        | Unique identifier for the service                              |
| `name`               | string        | Display name of the service                                    |
| `serviceItemType`    | ItemType      | Category of service item                                       |
| `category`           | string        | Service category for grouping and filtering                    |
| `price`              | Money         | Price for this service instance                                |
| `duration`           | Interval      | Scheduled time window for the service                          |
| `staffIds`           | Array(string) | IDs of staff members assigned to this service                  |
| `serviceType`        | Type          | Type of service being provided                                 |
| `serviceTime`        | integer       | Expected duration of the service in minutes                    |
| `petServiceDetailId` | string        | Unique identifier for this specific pet's service booking      |
| `lodgingId`          | string        | Unique identifier for this specific pet's service lodging unit |
| `lodgingUnitName`    | string        | Name of the lodging unit for this pet's service                |
| `lodgingTypeName`    | string        | Type of lodging unit for this pet's service                    |

---

### 6. AppointmentNote

Represents notes associated with an appointment, used for internal communication or customer information tracking.

| Field Name        | Type       | Description                                                  |
|-------------------|------------|--------------------------------------------------------------|
| `id`              | string     | Unique identifier of the note, obfuscated ID string          |
| `businessId`      | string     | Business location ID, obfuscated ID string                   |
| `customerId`      | string     | Customer ID, obfuscated ID string                            |
| `companyId`       | string     | Company ID, obfuscated ID string                             |
| `appointmentId`   | string     | Related appointment ID, obfuscated ID string                 |
| `note`            | string     | The content of the note                                      |
| `type`            | enum(Type) | Note type: `ALERT_NOTES`, `COMMENT`, `CANCEL`, `ADDITIONAL`  |
| `createdTime`     | timestamp  | When the note was created                                    |
| `lastUpdatedTime` | timestamp  | When the note was last modified                              |
| `createdBy`       | string     | Staff member who created the note, obfuscated ID string      |
| `lastUpdatedBy`   | string     | Staff member who last updated the note, obfuscated ID string |

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

| Field Name   | Type   | Required | Description                    |
|--------------|--------|----------|--------------------------------|
| `id`         | string | Yes      | Appointment ID to retrieve     |
| `businessId` | string | Yes      | Business ID for access control |

#### 📌 Return Value:

| Field Name          | Type                    | Description                                                                   |
|---------------------|-------------------------|-------------------------------------------------------------------------------|
| `id`                | string                  | Unique identifier of the appointment, obfuscated ID string                    |
| `businessId`        | string                  | Business location where the service will be performed, obfuscated ID string   |
| `customerId`        | string                  | Customer who booked the appointment, obfuscated ID string                     |
| `address`           | Address                 | Service location details (required for home service appointments)             |
| `duration`          | Interval                | Start and end time of the appointment                                         |
| `petServiceDetails` | Array(PetServiceDetail) | List of services booked for each pet                                          |
| `status`            | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.                   |
| `ticketComment`     | string                  | Optional notes about the appointment                                          |
| `colorCode`         | string                  | UI display color in hex format                                                |
| `orderId`           | string                  | Identifier of the associated payment order, obfuscated ID string              |
| `totalAmount`       | Money                   | Total cost for all services                                                   |
| `paidAmount`        | Money                   | Amount received from customer                                                 |
| `refundAmount`      | Money                   | Amount returned to customer                                                   |
| `paymentStatus`     | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`, `PARTIAL_REFUNDED`, `FULL_REFUNDED` |
| `createdBy`         | string                  | Identifier of the appointment creator, obfuscated ID string                   |
| `createdTime`       | timestamp               | When the appointment was created                                              |
| `lastUpdatedBy`     | string                  | Identifier of the last modifier, obfuscated ID string                         |
| `lastUpdatedTime`   | timestamp               | When the appointment was last modified                                        |
| `checkInTime`       | timestamp               | When the customer arrived with their pet                                      |
| `checkOutTime`      | timestamp               | When the service was completed and the pet picked up                          |
| `bookingRequestId`  | string (optional)       | The booking request ID associated with this appointment, obfuscated ID string |

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

| Field Name    | Type          | Required | Description                                    |
|---------------|---------------|----------|------------------------------------------------|
| `pagination`  | Pagination    | Yes      | Page size and token                            |
| `companyId`   | string        | Yes      | Company ID for access control                  |
| `businessIds` | Array(string) | Yes      | List of business IDs to filter appointments by |
| `filter`     | Filter        | No       | Filter options                                 |

#### Filter Options:

- `startTime`: Filter by appointment start time range.
- `endTime`: Filter by appointment end time range.
- `lastUpdatedTime`: Filter by last update time range.
- `statuses`: Filter by appointment status.
- `checkInTime`: Filter by check-in time range.
- `checkOutTime`: Filter by check-out time range.
- `serviceTypes`: Filter by service item type.
- `serviceIds`: Filter by service item ID.
- `petIds`: Filter by pet ID.
- `customerIds`: Filter by customer ID.

#### 📌 Return Value:

| Field Name      | Type               | Description                                   |
|-----------------|--------------------|-----------------------------------------------|
| `nextPageToken` | string             | Token for retrieving the next page of results |
| `appointments`  | Array(Appointment) | List of appointments matching the criteria    |

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

| Field Name       | Type              | Required | Description                                                                              |
|------------------|-------------------|----------|------------------------------------------------------------------------------------------|
| `businessId`     | string            | Yes      | Business location ID where services will be provided                                     |
| `customerId`     | string            | Yes      | Customer ID who is booking the appointment                                               |
| `petServices`    | Array(PetService) | Yes      | List of pets and their associated services                                               |
| `ignoreConflict` | bool              | No       | Whether to ignore scheduling conflicts when creating an appointment. Defaults to `true`. |

> ** Note ** : When set to 'true', the system will ignore time conflicts and directly create reservations. This field is
> mainly used to support mandatory appointments in certain special scenarios (such as manual intervention in the
> background).

##### PetService Object

| Field Name | Type           | Required | Description                                          |
|------------|----------------|----------|------------------------------------------------------|
| `petId`    | string         | Yes      | The unique identifier of the pet receiving services  |
| `services` | Array(Service) | Yes      | List of services to be provided for this pet         |

##### Service Object

| Field Name  | Type          | Required                    | Description                                                                                                    |
|-------------|---------------|-----------------------------|----------------------------------------------------------------------------------------------------------------|
| `id`        | string        | Yes                         | The unique identifier of the service. Must be an active service in the business's service list                 |
| `duration`  | Interval      | Yes                         | The scheduled time window for this service                                                                     |
| `staffIds`  | Array(string) | GROOMING only               | Staff member IDs. Required for grooming services. Leave empty for boarding/daycare                             |
| `lodgingId` | string        | BOARDING: Yes / DAYCARE: No | Lodging unit ID. Required for boarding, optional for daycare, not applicable for grooming                      |
| `price`     | Money         | No                          | Custom service price. If not set, the system default price will be used                                        |

#### 📌 Return Value:

| Field Name          | Type                    | Description                                                                   |
|---------------------|-------------------------|-------------------------------------------------------------------------------|
| `id`                | string                  | Unique identifier of the appointment, obfuscated ID string                    |
| `businessId`        | string                  | Business location where the service will be performed, obfuscated ID string   |
| `customerId`        | string                  | Customer who booked the appointment, obfuscated ID string                     |
| `address`           | Address                 | Service location details (required for home service appointments)             |
| `duration`          | Interval                | Start and end time of the appointment                                         |
| `petServiceDetails` | Array(PetServiceDetail) | List of services booked for each pet                                          |
| `status`            | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.                   |
| `ticketComment`     | string                  | Optional notes about the appointment                                          |
| `colorCode`         | string                  | UI display color in hex format                                                |
| `orderId`           | string                  | Identifier of the associated payment order, obfuscated ID string              |
| `totalAmount`       | Money                   | Total cost for all services                                                   |
| `paidAmount`        | Money                   | Amount received from customer                                                 |
| `refundAmount`      | Money                   | Amount returned to customer                                                   |
| `paymentStatus`     | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`, `PARTIAL_REFUNDED`, `FULL_REFUNDED` |
| `createdBy`         | string                  | Identifier of the appointment creator, obfuscated ID string                   |
| `createdTime`       | timestamp               | When the appointment was created                                              |
| `lastUpdatedBy`     | string                  | Identifier of the last modifier, obfuscated ID string                         |
| `lastUpdatedTime`   | timestamp               | When the appointment was last modified                                        |
| `checkInTime`       | timestamp               | When the customer arrived with their pet                                      |
| `checkOutTime`      | timestamp               | When the service was completed and the pet picked up                          |
| `bookingRequestId`  | string (optional)       | The booking request ID associated with this appointment, obfuscated ID string |

#### ⚠️ Important Notes:

> **Service type determines required fields**: The service item type (`GROOMING`, `BOARDING`, `DAYCARE`) is inferred
> from the service `id` itself.
>
> - `GROOMING`: `staffIds` is required. `lodgingId` is not applicable.
> - `BOARDING`: `staffIds` can be empty. `lodgingId` is required.
> - `DAYCARE`: `staffIds` can be empty. `lodgingId` is optional.

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields are missing or invalid (e.g., missing `staffIds` in services).
- `PERMISSION_DENIED`: Permission denied.

---

### 4. Reschedule Appointment (`RescheduleAppointment`)

- **Method**: `RescheduleAppointment`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/{id}:reschedule`

#### ✅ Functionality:

Updates the appointment's time slot. All services within the appointment will be moved to the new time.

#### 🎯 Use Cases:

- Move an appointment to a different time.
- Update appointment schedule due to staff availability changes.

#### 🔧 Request Parameters:

| Field Name   | Type     | Required | Description                         |
|--------------|----------|----------|-------------------------------------|
| `id`         | string   | Yes      | Appointment ID to reschedule        |
| `businessId` | string   | Yes      | Business ID for access control      |
| `duration`   | Interval | Yes      | New time window for the appointment |

#### 📌 Return Value:

| Field Name          | Type                    | Description                                                                   |
|---------------------|-------------------------|-------------------------------------------------------------------------------|
| `id`                | string                  | Unique identifier of the appointment, obfuscated ID string                    |
| `businessId`        | string                  | Business location where the service will be performed, obfuscated ID string   |
| `customerId`        | string                  | Customer who booked the appointment, obfuscated ID string                     |
| `address`           | Address                 | Service location details (required for home service appointments)             |
| `duration`          | Interval                | Start and end time of the appointment                                         |
| `petServiceDetails` | Array(PetServiceDetail) | List of services booked for each pet                                          |
| `status`            | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.                   |
| `ticketComment`     | string                  | Optional notes about the appointment                                          |
| `colorCode`         | string                  | UI display color in hex format                                                |
| `orderId`           | string                  | Identifier of the associated payment order, obfuscated ID string              |
| `totalAmount`       | Money                   | Total cost for all services                                                   |
| `paidAmount`        | Money                   | Amount received from customer                                                 |
| `refundAmount`      | Money                   | Amount returned to customer                                                   |
| `paymentStatus`     | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`, `PARTIAL_REFUNDED`, `FULL_REFUNDED` |
| `createdBy`         | string                  | Identifier of the appointment creator, obfuscated ID string                   |
| `createdTime`       | timestamp               | When the appointment was created                                              |
| `lastUpdatedBy`     | string                  | Identifier of the last modifier, obfuscated ID string                         |
| `lastUpdatedTime`   | timestamp               | When the appointment was last modified                                        |
| `checkInTime`       | timestamp               | When the customer arrived with their pet                                      |
| `checkOutTime`      | timestamp               | When the service was completed and the pet picked up                          |
| `bookingRequestId`  | string (optional)       | The booking request ID associated with this appointment, obfuscated ID string |

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

| Field Name   | Type   | Required | Description                    |
|--------------|--------|----------|--------------------------------|
| `id`         | string | Yes      | Appointment ID to cancel       |
| `businessId` | string | Yes      | Business ID for access control |

#### 📌 Return Value:

| Field Name          | Type                    | Description                                                                   |
|---------------------|-------------------------|-------------------------------------------------------------------------------|
| `id`                | string                  | Unique identifier of the appointment, obfuscated ID string                    |
| `businessId`        | string                  | Business location where the service will be performed, obfuscated ID string   |
| `customerId`        | string                  | Customer who booked the appointment, obfuscated ID string                     |
| `address`           | Address                 | Service location details (required for home service appointments)             |
| `duration`          | Interval                | Start and end time of the appointment                                         |
| `petServiceDetails` | Array(PetServiceDetail) | List of services booked for each pet                                          |
| `status`            | enum(Status)            | Current appointment state: `UNCONFIRMED`, `CONFIRMED`, etc.                   |
| `ticketComment`     | string                  | Optional notes about the appointment                                          |
| `colorCode`         | string                  | UI display color in hex format                                                |
| `orderId`           | string                  | Identifier of the associated payment order, obfuscated ID string              |
| `totalAmount`       | Money                   | Total cost for all services                                                   |
| `paidAmount`        | Money                   | Amount received from customer                                                 |
| `refundAmount`      | Money                   | Amount returned to customer                                                   |
| `paymentStatus`     | enum(PaymentStatus)     | Payment state: `UNPAID`, `PARTIAL_PAID`, `FULL_PAID`, `PARTIAL_REFUNDED`, `FULL_REFUNDED` |
| `createdBy`         | string                  | Identifier of the appointment creator, obfuscated ID string                   |
| `createdTime`       | timestamp               | When the appointment was created                                              |
| `lastUpdatedBy`     | string                  | Identifier of the last modifier, obfuscated ID string                         |
| `lastUpdatedTime`   | timestamp               | When the appointment was last modified                                        |
| `checkInTime`       | timestamp               | When the customer arrived with their pet                                      |
| `checkOutTime`      | timestamp               | When the service was completed and the pet picked up                          |
| `bookingRequestId`  | string (optional)       | The booking request ID associated with this appointment, obfuscated ID string |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified appointment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 6. Update Appointment Status (`UpdateAppointmentStatus`)

- **Method**: `UpdateAppointmentStatus`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/{id}:updateStatus`

#### ✅ Functionality:

Updates the status of an appointment and optionally sends notifications to the customer. This method allows moving an
appointment through its lifecycle by changing its status to one of the valid states.

#### 🎯 Use Cases:

- Move an appointment from `CONFIRMED` to `CHECKED_IN` when the customer arrives.
- Update an appointment to `READY` when services are completed but pet is not yet picked up.
- Mark an appointment as `FINISHED` when the customer picks up their pet.
- Change an appointment status to `CANCELED` (alternative to Cancel Appointment).

#### 🔧 Request Parameters:

| Field Name      | Type                | Required | Description                                                |
|-----------------|---------------------|----------|------------------------------------------------------------|
| `id`            | string              | Yes      | Appointment ID to update                                   |
| `businessId`    | string              | Yes      | Business ID for access control                             |
| `status`        | enum(Status)        | Yes      | Target status for the appointment                          |
| `messageMethod` | enum(MessageMethod) | No       | Method to use for sending notifications (SMS, EMAIL, etc.) |
| `checkOut`      | object(CheckOut)    | No       | Check out details for when status is FINISHED              |

##### MessageMethod Enum:

- `MESSAGE_METHOD_UNSPECIFIED`: Do not send notifications
- `SMS`: Send notification via SMS
- `EMAIL`: Send notification via Email

##### CheckOut Object:

| Field Name | Type | Required | Description                                                                                      |
|------------|------|----------|--------------------------------------------------------------------------------------------------|
| `endDate`  | Date | No       | Actual check out date. Updates boarding appointment check out date and corresponding order price |

#### 📌 Return Value:

Returns the updated `Appointment` object with the new status and all associated details.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified appointment ID does not exist.
- `PERMISSION_DENIED`: Permission denied.
- `INVALID_ARGUMENT`: Invalid status transition or other validation errors.

---

### 7. Check Appointment Availability (`CheckAppointmentAvailability`)

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

| Field Name       | Type          | Required | Description           |
|------------------|---------------|----------|-----------------------|
| `businessId`     | string        | Yes      | Business location ID  |
| `dateRange`      | Interval      | Yes      | Appointment time slot |
| `petIds`         | Array(string) | No       | List of pets ids      |
| `customerId`     | string        | No       | Customer ID           |
| `appointmentId`  | string        | No       | Appointment ID        |
| `lodgingUnitIds` | Array(string) | No       | Lodging  Unit ID      |

#### 📌 Return Value:

| Field Name                 | Type                                 | Description                                                                 |
|----------------------------|--------------------------------------|-----------------------------------------------------------------------------|
| `appointmentConflictCheck` | `AppointmentDateConflictCheckResult` | Contains information about conflicting appointments for each pet.           |
| `businessClosedDateCheck`  | `BusinessClosedDateCheckResult`      | Contains information about business closed dates during the requested time. |
| `lodgingOverCapacityCheck` | `LodgingOverCapacityCheckResult`     | Contains information about lodging over capacity during the requested time. |

##### CheckAppointmentAvailabilityResponse

Represents the result of a pre-check operation before creating or rescheduling an appointment. It includes two main
components:

###### 1. `AppointmentDateConflictCheckResult`

| Field Name  | Type                             | Description                                        |
|-------------|----------------------------------|----------------------------------------------------|
| `conflicts` | Array(`PetAppointmentsOverview`) | A list of pets and their conflicting appointments. |

###### 2. `PetAppointmentsOverview`

| Field Name     | Type                 | Description                                      |
|----------------|----------------------|--------------------------------------------------|
| `pet`          | `Pet`                | The pet involved in the conflicting appointment. |
| `appointments` | Array(`Appointment`) | A list of conflicting appointments for this pet. |

###### 3. `BusinessClosedDateCheckResult`

| Field Name    | Type                      | Description                                                  |
|---------------|---------------------------|--------------------------------------------------------------|
| `closedDates` | Array(`google.type.Date`) | A list of business closed dates during the requested period. |

###### 4. `LodgingOverCapacityCheckResult`

This structure is used to indicate whether lodging units are over capacity during the specified time period. When an
appointment involves pet boarding services, the system checks if the business has exceeded its lodging capacity during
the requested time.

| Field Name | Type                                       | Description                          |
|------------|--------------------------------------------|--------------------------------------|
| `lodgings` | Array(`moego.business.setting.v1.Lodging`) | List of lodging that exceed capacity |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Missing required fields in request.
- `PERMISSION_DENIED`: Permission denied.

---

### 8. List Grooming Reports (`ListGroomingReports`)

- **Method**: `ListGroomingReports`
- **HTTP Method**: POST
- **Path**: `/v1/grooming_reports:list`

#### ✅ Functionality:

Lists grooming reports for specified appointments.

#### 🎯 Use Cases:

- Retrieve grooming reports for reporting or customer communication.
- Review historical grooming data for pets.

#### 🔧 Request Parameters:

| Field Name       | Type          | Required | Description                    |
|------------------|---------------|----------|--------------------------------|
| `appointmentIds` | Array(string) | Yes      | Max 500 appointment IDs        |
| `businessId`     | string        | Yes      | Business ID for access control |

#### 📌 Return Value:

| Field Name        | Type                  | Description              |
|-------------------|-----------------------|--------------------------|
| `groomingReports` | Array(GroomingReport) | List of grooming reports |

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 9. Create Appointment Note (`CreateAppointmentNote`)

- **Method**: `CreateAppointmentNote`
- **HTTP Method**: POST
- **Path**: `/v1/appointments/notes`

#### ✅ Functionality:

Creates a new note associated with a specific appointment.

#### 🎯 Use Cases:

- Add alerts or comments related to an appointment.
- Record cancellation reasons or additional notes from customers.

#### 🔧 Request Parameters:

| Field Name      | Type                       | Required | Description                               |
|-----------------|----------------------------|----------|-------------------------------------------|
| `appointmentId` | string                     | Yes      | The appointment ID to associate this note |
| `note`          | string                     | Yes      | Content of the note                       |
| `type`          | enum(AppointmentNote.Type) | Yes      | Type of note                              |

#### 📌 Return Value:

| Field Name        | Type       | Description                                                  |
|-------------------|------------|--------------------------------------------------------------|
| `id`              | string     | Unique identifier of the note, obfuscated ID string          |
| `businessId`      | string     | Business location ID, obfuscated ID string                   |
| `customerId`      | string     | Customer ID, obfuscated ID string                            |
| `companyId`       | string     | Company ID, obfuscated ID string                             |
| `appointmentId`   | string     | Related appointment ID, obfuscated ID string                 |
| `note`            | string     | The content of the note                                      |
| `type`            | enum(Type) | Note type: `ALERT_NOTES`, `COMMENT`, `CANCEL`, `ADDITIONAL`  |
| `createdTime`     | timestamp  | When the note was created                                    |
| `lastUpdatedTime` | timestamp  | When the note was last modified                              |
| `createdBy`       | string     | Staff member who created the note, obfuscated ID string      |
| `lastUpdatedBy`   | string     | Staff member who last updated the note, obfuscated ID string |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Missing required fields.
- `PERMISSION_DENIED`: Permission denied.

---

### 10. Update Appointment Note (`UpdateAppointmentNote`)

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

| Field Name        | Type       | Description                                                  |
|-------------------|------------|--------------------------------------------------------------|
| `id`              | string     | Unique identifier of the note, obfuscated ID string          |
| `businessId`      | string     | Business location ID, obfuscated ID string                   |
| `customerId`      | string     | Customer ID, obfuscated ID string                            |
| `companyId`       | string     | Company ID, obfuscated ID string                             |
| `appointmentId`   | string     | Related appointment ID, obfuscated ID string                 |
| `note`            | string     | The content of the note                                      |
| `type`            | enum(Type) | Note type: `ALERT_NOTES`, `COMMENT`, `CANCEL`, `ADDITIONAL`  |
| `createdTime`     | timestamp  | When the note was created                                    |
| `lastUpdatedTime` | timestamp  | When the note was last modified                              |
| `createdBy`       | string     | Staff member who created the note, obfuscated ID string      |
| `lastUpdatedBy`   | string     | Staff member who last updated the note, obfuscated ID string |

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified note ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 11. List Appointment Notes (`ListAppointmentNotes`)

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

| Field Name      | Type                       | Required | Description                       |
|-----------------|----------------------------|----------|-----------------------------------|
| `appointmentId` | string                     | Yes      | Appointment ID for access control |
| `type`          | enum(AppointmentNote.Type) | No       | Filter by note type               |

#### 📌 Return Value:

| Field Name | Type                   | Description               |
|------------|------------------------|---------------------------|
| `notes`    | Array(AppointmentNote) | List of appointment notes |

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
  "businessId": "biz_001",
  "customerId": "cus_001",
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001"
  },
  "duration": {
    "startTime": "2024-08-15T10:00:00Z",
    "endTime": "2024-08-15T12:00:00Z"
  },
  "petServiceDetails": [
    {
      "pet": {
        "id": "pet_001",
        "name": "Buddy"
      },
      "serviceDetails": [
        {
          "id": "svc_123",
          "name": "Premium Grooming",
          "price": {
            "currencyCode": "USD",
            "units": 50,
            "nanos": 0
          },
          "duration": {
            "startTime": "2024-08-15T10:00:00Z",
            "endTime": "2024-08-15T11:00:00Z"
          }
        }
      ]
    }
  ],
  "status": "CONFIRMED",
  "paymentStatus": "FULL_PAID",
  "totalAmount": {
    "currencyCode": "USD",
    "units": 75,
    "nanos": 0
  },
  "paidAmount": {
    "currencyCode": "USD",
    "units": 75,
    "nanos": 0
  },
  "createdTime": "2024-08-10T09:00:00Z",
  "lastUpdatedTime": "2024-08-12T14:30:00Z"
}
```

### Example 2: Create Grooming Appointment

```
{
  "businessId": "biz_001",
  "customerId": "cus_001",
  "petServices": [
    {
      "petId": "pet_001",
      "services": [
        {
          "id": "svc_grooming",
          "duration": {
            "startTime": "2024-08-15T10:00:00Z",
            "endTime": "2024-08-15T11:00:00Z"
          },
          "staffIds": [
            "staff_001"
          ]
        }
      ]
    }
  ]
}
```

### Example 2.1: Create Boarding Appointment

> Note: `staffIds` is optional for boarding. `lodgingId` specifies the lodging unit.

```json
{
  "businessId": "bizIT9f",
  "customerId": "cuxVOTi9O",
  "petServices": [
    {
      "petId": "petOvAv9M",
      "services": [
        {
          "id": "svcPOX8H",
          "duration": {
            "startTime": "2026-03-30T20:00:00Z",
            "endTime": "2026-03-31T06:00:00Z"
          },
          "price": {
            "currencyCode": "USD",
            "units": "46",
            "nanos": 500000000
          },
          "staffIds": ["stfUNfY"],
          "lodgingId": "lodgu3PAP"
        }
      ]
    }
  ],
  "ignoreConflict": true
}
```

### Example 2.2: Create Daycare Appointment

> Note: `staffIds` is optional for daycare. `lodgingId` is optional.

```json
{
  "businessId": "bizIT9f",
  "customerId": "cuxVOTi9O",
  "petServices": [
    {
      "petId": "petOvAv9M",
      "services": [
        {
          "id": "svcSO0bK",
          "duration": {
            "startTime": "2026-03-29T16:00:00Z",
            "endTime": "2026-03-30T02:00:00Z"
          },
          "price": {
            "currencyCode": "USD",
            "units": "63",
            "nanos": 500000000
          },
          "staffIds": ["stfUNfY"],
          "lodgingId": "lodgu3PAP"
        }
      ]
    }
  ],
  "ignoreConflict": true
}
```

### Example 3: Reschedule Appointment

```
POST /v1/appointments/apt_12345:reschedule
```

```
{
  "businessId": "biz_001",
  "duration": {
    "startTime": "2024-08-16T11:00:00Z",
    "endTime": "2024-08-16T13:00:00Z"
  }
}
```

### Example 4: Check Create Appointment

```
POST /v1/appointments:check
```

```
{
  "businessId": "biz_001",
  "customerId": "cus_001",
  "dateRange": {
    "startTime": "2024-08-15T10:00:00Z",
    "endTime": "2024-08-15T11:00:00Z"
  },
  "petIds": [
    "pet_001"
  ]
}
```

**Response:**

```
{
  "appointmentConflictCheck": {
    "conflicts": [
      {
        "pet": {
          "id": "pet_001",
          "name": "Buddy"
        },
        "appointments": [
          {
            "id": "apt_98765",
            "businessId": "biz_001",
            "customerId": "cus_002",
            "duration": {
              "startTime": "2024-08-15T10:30:00Z",
              "endTime": "2024-08-15T11:30:00Z"
            },
            "status": "CONFIRMED"
          }
        ]
      }
    ]
  },
  "businessClosedDateCheck": {
    "closedDates": []
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