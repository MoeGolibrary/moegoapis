# MoeGo APIs

This repository contains the original interface definitions of public MoeGo APIs that support REST protocols. Reading
these definitions helps you better understand the MoeGo APIs and use them more efficiently. You can also use these
definitions with open-source tools to generate client libraries, documentation, and other artifacts.

## Try it now

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 96px; height: 24px;">](https://app.getpostman.com/run-collection/30555124-2c27c225-3a6c-4234-bd4f-f818dc0a757c?source=rip_markdown&collection-url=entityId%3D30555124-2c27c225-3a6c-4234-bd4f-f818dc0a757c%26entityType%3Dcollection%26workspaceId%3Ddb8fd53e-fafa-4b5e-bf78-8dd330f46a0b)

Click the `Run in Postman` button to try out the MoeGo APIs using Postman — a popular API client that makes it easy to
explore and test APIs without writing any code.

**Note**: Replace the collection Authorization `${Base_64_API_Key}` with your own API Key encoded in Base64.

---

## 📌 Important Notes

- **Obfuscated IDs**: All our identifiers (e.g., `CompanyID`, `BusinessID`) are obfuscated for security purposes. Please
  ensure that you use the obfuscated versions of these IDs when making API requests.

- **API Limitations**: If any of the current APIs do not meet your requirements or if you need additional functionality,
  please reach out to your **Customer Success Manager** for assistance and further guidance.

---

## 🚀 Getting Started

To begin using the MoeGo APIs effectively, follow these steps:

1. **Request API Access**

    - For all API key requests, we kindly ask you to work directly with your dedicated **Customer Success Manager**.
      They
      will guide you through our API key application process.

2. **Retrieve Required Information**

    - Use the [`ListCompany`](docs/company.md) and [`ListBusiness`](docs/business.md) APIs to retrieve the necessary
      obfuscated IDs (`CompanyID`, `BusinessID`) and verify that your requests are targeting the correct entities.

3. **Validate with Postman**

    - Use the [Postman collection](#try-it-now) to test and validate API calls before integrating them into your application.
    - Replace the placeholder authorization token with your Base64-encoded API key.
    - Review example requests and responses to understand API behavior and expected data formats.

4. **Explore OpenAPI Specifications**

    - Check the [OpenAPI specifications](#-openapi-specifications) to understand detailed API contracts, including request/response schemas, error codes, and endpoint parameters.
    - Use these specifications with tools like Swagger UI or OpenAPI Generator to create client SDKs in your preferred programming language.

5. **(Optional) Use Example Scripts**

    - Use the provided example scripts like [`export_orders.py`](examples/export_orders.py) to quickly get started with common tasks such as exporting order data to CSV format.

6. **(Optional) Configure Webhooks**

    - Set up [Webhooks](#-webhook-integration) to subscribe to events you care about, such as `APPOINTMENT_CREATED` or
      `APPOINTMENT_UPDATED`. This enables real-time notifications for key business events.

---

## 🌐 Domain Structure

MoeGo APIs are organized under the following domain:

```
openapi.moego.pet
```

Each path corresponds to a specific version and resource:

- `/v1/auth`
- `/v1/appointment`
- `/v1/customer`
- `/v1/webhooks`
- `/v1/order`

Each module includes multiple APIs for handling different use cases such as authentication, data management, and event
notifications.

---

## 📦 Core Modules Overview

### 🔐 Authentication & Authorization

- [Authentication](./moego/auth/README.md): An overview of authentication methods.
- [API Key](./moego/auth/apikey/README.md): Details on how to use API keys for authentication.

### 👥 Customer & Pet Management

- [Customer](docs/customer.md): Manage customer profiles, tags, notes, and related operations.
- [Pet](docs/pet.md): Animal companion information, including coat type, behavior, and vet details.
- [Lead](docs/lead.md): Potential customers in the sales pipeline with lifecycle and action status tracking.

### 📅 Scheduling & Availability

- [Appointment](docs/appointment.md): Records of services provided to customers.
- [Block](docs/block.md): Non-bookable time slots (e.g., staff days off).
- [Online Booking](docs/online_booking.md): Enables customers to book appointments online.

### 💰 Orders & Payments

- [Order](docs/order.md): Purchase records with filtering by status, company, or business ID.
- [Payment](docs/payment.md): Transaction details between customers and businesses.

### 🛒 Retail & Inventory Management

- [Retail](docs/retail.md): Manages retail products and supplier relationships.

### ⚙️ Settings & Integrations

- [Webhook](docs/webhook.md): Real-time event notification system with HTTPS callbacks, HMAC signing, and retry
  capabilities.
- [Event](docs/event.md): List of supported events like `APPOINTMENT_CREATED`, `ORDER_UPDATED`.
- [SettingService](docs/setting_service.md): Configure business-level settings and preferences.

### 📋 Agreements & Feedback

- [Agreement](docs/agreement.md): Contract between a business and a customer.
- [Review](docs/review.md): Customer feedback on service experiences.

### 📊 Reporting & Analytics

- [Report](docs/report.md): Business insights derived from collected data.

### 🧩 Common Data Models

- [Address](docs/common/address.md): Physical location representation.
- [Pagination](docs/common/pagination.md): Mechanism for managing page parameters.
- [Weight](docs/common/weight.md): Weight unit used across data models.

---

## 🔄 Webhook Integration

Webhooks provide an efficient way to receive real-time updates when certain events occur within the MoeGo system. Key
features include:

- **Event Subscription**: Subscribe to events such as `APPOINTMENT_CREATED` and `ORDER_UPDATED`.
- **Secure Delivery**: Supports HTTPS, SSL verification, and HMAC-signed payloads.
- **Delivery Logs**: Use `ListWebhookDeliveries` to view logs and identify failed deliveries.
- **Test & Debug**: Trigger test calls using `TriggerTestWebhookDelivery`.

For more details:

- [Moego Webhook Integration Guide](docs/webhook_integration.md): A step-by-step guide on integrating with Moego's
  webhook system.
- [webhook documentation](docs/webhook.md): API reference for webhooks.
- [event documentation](docs/event.md): API reference for events.

---

## 📄 OpenAPI Specifications

To better support developers using MoeGo APIs, we have added corresponding `openapi.json` files under the `docs/openapi`
directory. These files are generated based on the Protocol Buffer definitions and provide comprehensive interface
descriptions for REST APIs, including request paths, parameters, and response formats.

With these OpenAPI specifications, developers can:

- Quickly understand how to use each API, including data structures and calling methods;
- Use open-source tools (such as Swagger UI, Postman, etc.) to generate visual documentation;
- Automatically generate client SDKs or server-side stub code;
- Use them as a reference for API testing and debugging.

It is recommended to refer to the corresponding module documentation (
e.g., [Customer](docs/customer.md), [Appointment](docs/appointment.md)) alongside the OpenAPI files for a more complete
understanding.

Each JSON file corresponds to a core module and follows the OpenAPI 3.0 specification, making it easy to integrate into
various development toolchains.