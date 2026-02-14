# Event Documentation (`moego.business.event.v1`)

## 📌 1. Functional Overview

The `Event` service provides a centralized mechanism for handling significant system occurrences that require
notification or processing. It supports:

- System monitoring through health check events
- Appointment lifecycle tracking (creation, updates, completion, cancellation)
- Online booking notifications
- Customer data changes (creation, updates, deletion)
- Unified event metadata and payload structure

This service enables reliable event delivery and processing across system components.

---

## 🎯 2. Design Goals

- **Unified Event Handling**: Provides a standardized interface for managing different types of events
- **Rich Classification**: Supports categorized event types with clear boundaries for proper routing
- **Reliable Delivery**: Ensures proper identification and timestamping for event processing
- **Extensible Architecture**: Designed to accommodate future event types without breaking changes

Applicable to scenarios such as system monitoring, workflow automation, notification triggering, and resource
management.

---

## 🧩 3. Core Concepts

### 1. Event

Represents a significant occurrence within the system that requires notification or processing. Events can be
system-level operations, appointment status changes, or customer interactions.

| Field Name  | Type      | Description                                                               |
|-------------|-----------|---------------------------------------------------------------------------|
| `id`        | string    | Unique identifier for the event (format: "evt_" followed by random chars) |
| `type`      | Type      | Category of the event determining payload type and processing rules       |
| `timestamp` | Timestamp | When this event occurred (for ordering and processing)                    |
| `companyId` | string    | ID of the company associated with the event (for routing/access control)  |
| `payload`   | oneof     | Event-specific data based on the event type                               |

### 2. Event Types

Events are categorized into functional groups with specific ranges:

| Range   | Category              | Description                                 |
|---------|-----------------------|---------------------------------------------|
| 0-99    | System Events         | Monitoring and infrastructure events        |
| 100-199 | Appointment Events    | Lifecycle changes for appointments          |
| 200-299 | Online Booking Events | Customer self-service booking notifications |
| 300-399 | Customer Events      | Customer data changes                   |

### 3. Payload Structures

Each event type has an associated payload format:

| Event Type             | Payload Type | Description                       |
|------------------------|--------------|-----------------------------------|
| `HEALTH_CHECK`         | HealthCheck  | System monitoring verification    |
| `APPOINTMENT_*` series | Appointment  | Appointment lifecycle information |
| `CUSTOMER_*` series   | Customer     | Customer data changes        |

---

## 📦 4. API Interface Descriptions

### 1. Event Message

```protobuf
message Event {
  enum Type {
    TYPE_UNSPECIFIED = 0;

    // System Events (1-99)
    HEALTH_CHECK = 1;

    // Appointment Events (100-199)
    APPOINTMENT_CREATED = 100;
    APPOINTMENT_UPDATED = 101;
    APPOINTMENT_FINISHED = 102;
    APPOINTMENT_CANCELED = 103;
    APPOINTMENT_DELETED = 104;

    // Customer Events (300-399)
    CUSTOMER_CREATED = 300;
    CUSTOMER_UPDATED = 301;
    CUSTOMER_DELETED = 302;
  }

  string id = 1;
  Type type = 2;
  google.protobuf.Timestamp timestamp = 3;
  string company_id = 4;

  oneof payload {
    HealthCheck health_check = 5;
    moego.business.appointment.v1.Appointment appointment = 6;
    moego.business.online_booking.v1.OnlineBooking online_booking = 7;
    moego.business.customer.v1.Customer customer = 8;
  }
}
```

### 2. HealthCheck Message

```protobuf
message HealthCheck {
  string validation = 1; // Validation string to be echoed back by webhook endpoint
}
```

---

## 🧪 5. Usage Examples

### Example 1: Health Check Event

```json
{
  "id": "evt_abc123",
  "type": "HEALTH_CHECK",
  "timestamp": "2024-08-01T12:00:00Z",
  "companyId": "cmp_001",
  "healthCheck": {
    "validation": "test_validation_string"
  }
}
```

### Example 2: Appointment Created Event

```json
{
  "id": "evt_def456",
  "type": "APPOINTMENT_CREATED",
  "timestamp": "2024-08-01T12:05:00Z",
  "companyId": "cmp_001",
  "appointment": {
    // Full appointment details from moego.business.appointment.v1.Appointment
  }
}
```

### Example 3: Customer Created Event

```json
{
  "id": "evt_ghi789",
  "type": "CUSTOMER_CREATED",
  "timestamp": "2024-08-01T12:10:00Z",
  "companyId": "cmp_001",
  "customer": {
    // Full customer details from moego.business.customer.v1.Customer
  }
}
```

---

## ⚠️ 6. Usage Limitations

TODO

---

## 📎 7. FAQ

| Question                                                | Answer                                                                                                                                                          |
|---------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| How do I verify webhook endpoint availability?          | Use the `HEALTH_CHECK` event type with the provided validation string                                                                                           |
| Which event indicates a new appointment?                | The `APPOINTMENT_CREATED` event type                                                                                                                            |
| Which event indicates a customer update?                | The `CUSTOMER_UPDATED` event type triggers when customer profile is modified                                                                                         |
| How are different event types processed differently?    | Each event type has a specific payload format and processing logic - system events use HealthCheck, appointments use Appointment, customers use Customer, etc.                          |
| Why does my system need to handle multiple event types? | Different business operations require different handling procedures - appointment changes affect scheduling while customer updates trigger profile synchronization workflows |
| How to ensure proper event ordering?                    | Use the `timestamp` field for event sequencing                                                                                                                  |

---

## 📌 8. Common Error Codes

| Error Code          | Description                       |
|---------------------|-----------------------------------|
| `NOT_FOUND`         | Event ID does not exist           |
| `PERMISSION_DENIED` | Current user has no access rights |
| `INVALID_ARGUMENT`  | Invalid request parameters        |
| `INTERNAL`          | Internal server error             |

---

## 📎 9. Related File References

- [event.proto](../moego/business/event/v1/event.proto)
- [health_check.proto](../moego/business/event/v1/health_check.proto)
- [appointment.proto](../moego/business/appointment/v1/appointment.proto)
- [online_booking.proto](../moego/business/online_booking/v1/online_booking.proto)
- [customer.proto](../moego/business/customer/v1/customer.proto)
