# Message API Documentation (`moego.business.message.v1`)

## 📌 1. Functional Overview

The Message API provides customer messaging and conversation thread lifecycle management.
This interface enables:

- Sending direct messages to a customer via a specified delivery channel (SMS, email, call, or app)
- Triggering template-based automated messages for business events such as appointment reminders and payment confirmations
- Automatic conversation thread management — the system creates or reuses a thread per customer; callers do not manage threads directly

---

## 🎯 2. Design Goals

- **Channel Flexibility**: Supports multiple delivery methods (MSG, EMAIL, CALL, APP) in a unified API
- **Thread Transparency**: Thread creation and management are handled internally; callers only supply message content and delivery preferences
- **Template-Driven Automation**: Auto messages use pre-configured templates and context variables (e.g., appointment ID), reducing manual message composition
- **Multi-Channel Auto Messaging**: Auto messages can be dispatched across multiple channels simultaneously in one request

---

## 🧩 3. Core Concepts

### 1. MessageDeliveryMethod

Represents the delivery channel used for sending a message.

| Value                              | Numeric | Description                                      |
|------------------------------------|---------|--------------------------------------------------|
| `MESSAGE_DELIVERY_METHOD_UNSPECIFIED` | 0    | Unknown or unspecified delivery method           |
| `MSG`                              | 1       | In-app / internal message                        |
| `EMAIL`                            | 2       | Email                                            |
| `CALL`                             | 4       | Phone call (automated or manual)                 |
| `APP`                              | 5       | App push notification or in-app message          |

> **Note**: Value `3` is reserved and should not be used.

### 2. MessageType

Represents the content type of a message body.

| Value                    | Description                      |
|--------------------------|----------------------------------|
| `MESSAGE_TYPE_UNSPECIFIED` | Unknown or unspecified type    |
| `TEXT`                   | Plain text message               |
| `PIC`                    | Picture / image message          |

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: Staff Sends a Manual Message to a Customer

1. Identify the customer by ID
2. Choose the delivery method (e.g., `MSG` for in-app)
3. Call `SendMessageToCustomer` with the customer ID, message body, and method
4. The system finds or creates a conversation thread for the customer and delivers the message
5. Use `threadId` and `messageId` in the response for reference or tracking

### ✅ Scenario: System Sends an Appointment Reminder

1. After booking an appointment, call `SendAutoMessageToCustomer` with the customer ID, business ID, and `use_case` (e.g., `appointment_reminder`)
2. Pass the appointment ID as `contextId` for variable substitution in the template
3. Specify one or more delivery channels in `methods` (e.g., `[MSG, EMAIL]`)
4. The system resolves the template, fills in variables, and dispatches the message via all specified channels

---

## 📦 5. API Interface Descriptions

### 1. Send Message to Customer (`SendMessageToCustomer`)

- **Method**: `SendMessageToCustomer`
- **HTTP Method**: POST
- **Path**: `/v1/messages:sendToCustomer`

#### ✅ Functionality:

Creates or reuses a conversation thread for the customer, then sends the message using the specified delivery method.
Thread creation is handled internally; callers only supply message content and the delivery channel.

#### 🎯 Use Cases:

- Staff sends a manual message to a customer to confirm an appointment
- Third-party integration delivers a notification or update via the customer's preferred channel
- Send a picture or image to a customer (e.g., grooming photo)

#### 🔧 Request Parameters:

| Field Name      | Type                  | Required | Description                                                                                                       |
|-----------------|-----------------------|----------|-------------------------------------------------------------------------------------------------------------------|
| `customerId`    | string                | Yes      | Customer identifier (obfuscated) to send the message to                                                           |
| `phoneNumber`   | string                | No       | Phone number (E.164 format) when the customer has multiple phones. Required when `method = CALL`; optional otherwise |
| `messageBody`   | string                | Yes      | Message content — plain text string or image reference for `PIC` type                                            |
| `messageType`   | MessageType           | Yes      | Content type: `TEXT` or `PIC`. Defaults to `TEXT` if not set                                                     |
| `method`        | MessageDeliveryMethod | Yes      | Delivery channel: `MSG`(1), `EMAIL`(2), `CALL`(4), or `APP`(5)                                                   |

#### 💡 Example Request:

```json
{
  "customerId": "cus_abc123",
  "messageBody": "Your appointment is confirmed for tomorrow at 10 AM. See you soon!",
  "messageType": "TEXT",
  "method": "MSG"
}
```

#### 📌 Return Value:

| Field Name   | Type   | Description                                                                    |
|--------------|--------|--------------------------------------------------------------------------------|
| `threadId`   | string | Identifier of the conversation thread used or created (optional; for reference) |
| `messageId`  | string | Identifier of the sent message (optional; returned when available from backend)  |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields (`customerId`, `messageBody`, `messageType`, `method`) are missing or invalid.
- `NOT_FOUND`: The specified customer does not exist.
- `PERMISSION_DENIED`: The caller does not have access to the specified customer or business.

---

### 2. Send Auto Message to Customer (`SendAutoMessageToCustomer`)

- **Method**: `SendAutoMessageToCustomer`
- **HTTP Method**: POST
- **Path**: `/v1/messages:sendAutoMessageToCustomer`

#### ✅ Functionality:

Triggers a template-based automated message to a customer via one or more configured delivery channels. Message content
is derived from the configured template and optional context (e.g., `appointmentId`). Useful for appointment reminders,
payment confirmations, and other event-driven notifications.

#### 🎯 Use Cases:

- Send an appointment confirmation or reminder after booking
- Notify a customer of a payment received
- Trigger a post-service follow-up or review request

#### 🔧 Request Parameters:

| Field Name   | Type                          | Required | Description                                                                                              |
|--------------|-------------------------------|----------|----------------------------------------------------------------------------------------------------------|
| `customerId` | string                        | Yes      | Customer identifier (obfuscated) to send the auto message to                                             |
| `businessId` | string                        | Yes      | Business identifier (obfuscated) for context and template lookup                                         |
| `useCase`    | string                        | Yes      | Use case or template type (e.g., `appointment_reminder`, `payment_confirmation`)                         |
| `methods`    | Array(MessageDeliveryMethod)  | Yes      | One or more delivery channels: `MSG`(1), `EMAIL`(2), `CALL`(4), `APP`(5). At least one required         |
| `contextId`  | string                        | No       | Optional context for template variable substitution (e.g., appointment ID for appointment reminders)    |

#### 💡 Example Request:

```json
{
  "customerId": "cus_abc123",
  "businessId": "biz_001",
  "useCase": "appointment_reminder",
  "methods": ["MSG", "EMAIL"],
  "contextId": "apt_xyz789"
}
```

#### 📌 Return Value:

This method returns an empty response on success.

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Required fields (`customerId`, `businessId`, `useCase`, `methods`) are missing or invalid, or `methods` is empty.
- `NOT_FOUND`: The specified customer or business does not exist.
- `PERMISSION_DENIED`: The caller does not have access to the specified business.

---

## 🧪 6. Usage Examples

### Example 1: Send a Text Message via In-App Channel

```json
POST /v1/messages:sendToCustomer
{
  "customerId": "cus_abc123",
  "messageBody": "Hi Jane, your grooming appointment is confirmed for April 10 at 10 AM.",
  "messageType": "TEXT",
  "method": "MSG"
}
```

Response:
```json
{
  "threadId": "thread_001",
  "messageId": "msg_111"
}
```

### Example 2: Send an Auto Reminder via SMS and Email

```json
POST /v1/messages:sendAutoMessageToCustomer
{
  "customerId": "cus_abc123",
  "businessId": "biz_001",
  "useCase": "appointment_reminder",
  "methods": ["MSG", "EMAIL"],
  "contextId": "apt_xyz789"
}
```

### Example 3: Send a Call Notification with Specific Phone Number

```json
POST /v1/messages:sendToCustomer
{
  "customerId": "cus_abc123",
  "phoneNumber": "+12125551234",
  "messageBody": "Your pet is ready for pickup.",
  "messageType": "TEXT",
  "method": "CALL"
}
```

---

## ⚠️ 7. Usage Limitations

- `SendMessageToCustomer` delivers to a single channel per call; to send via multiple channels, call the API multiple times or use `SendAutoMessageToCustomer` instead.
- The `CALL` delivery method requires `phoneNumber` to be set when the customer has multiple registered phones.
- `SendAutoMessageToCustomer` requires the `useCase` to correspond to a template configured for the business; using an unknown `useCase` value may result in no message being sent.
- `contextId` substitution only applies when the template is configured to use that variable; passing an unrecognized `contextId` is silently ignored.

---

## 📎 8. FAQ

| Question                                                                         | Answer                                                                                                                                      |
|----------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Do I need to create a conversation thread before sending a message?              | No. The system automatically creates or reuses a thread per customer. Just provide the customer ID and message content.                     |
| Can I send a message to multiple customers at once?                              | No. Each call targets a single customer. Call the API once per customer for batch sending.                                                  |
| What delivery method should I use for appointment reminders?                     | Use `SendAutoMessageToCustomer` with the relevant `useCase`. Specify `methods` to control channels (e.g., `[MSG, EMAIL]`).                  |
| What is the difference between `MSG` and `APP`?                                  | `MSG` is an in-app / internal channel (numeric value 1); `APP` targets push notifications or in-app alerts (numeric value 5).               |
| What happens if the template for `useCase` is not configured for the business?   | No message will be sent. Ensure the use case template is configured in the business settings before calling this API.                       |

---

## 📌 9. Common Error Codes

| Error Code          | Description                                                          |
|---------------------|----------------------------------------------------------------------|
| `INVALID_ARGUMENT`  | Missing or invalid required fields, or `methods` array is empty      |
| `NOT_FOUND`         | Customer or business ID does not exist                               |
| `PERMISSION_DENIED` | Caller does not have access to the specified customer or business    |
| `INTERNAL`          | Internal server error                                                |

---

## 📎 10. Related File References

- [appointment.md](./appointment.md)
- [message_service.proto](../moego/business/message/v1/message_service.proto)
- [message.proto](../moego/business/message/v1/message.proto)
