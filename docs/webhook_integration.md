# 🪝 Moego Webhook Integration Guide

Welcome to using **Moego Webhook Service**. This guide will help you step-by-step through the integration process of
Webhooks, including configuring receiving endpoints, creating Webhooks, verifying signatures, testing integration,
handling failed requests, and monitoring delivery status.

## 📋 Document Objectives

- ✅ Provide clear steps and examples for integration
- ✅ Guide developers through full configuration workflow
- ✅ Offer best practices for security validation and troubleshooting
- ✅ Help developers quickly locate API documentation and common solutions

---

## 🧭 Overview of Integration Process

| Step | Description                                    |
|------|------------------------------------------------|
| 1️⃣  | Prepare Webhook Receiving Endpoint             |
| 2️⃣  | Create Webhook (via API or Management Console) |
| 3️⃣  | Verify Request Origin Authenticity             |
| 4️⃣  | Test if Webhook Works Properly                 |
| 5️⃣  | Handle Failures and Retry Mechanisms           |
| 6️⃣  | Monitor Webhook Delivery Status                |

---

## 1️⃣ Step 1: Prepare Webhook Receiving Endpoint

To successfully receive Webhook requests from Moego, your service must meet the following basic requirements:

### ✅ Mandatory Requirements

| Requirement                       | Description                                                                      |
|-----------------------------------|----------------------------------------------------------------------------------|
| HTTPS Support                     | Must use HTTPS protocol (HTTP is not accepted)                                   |
| Accept POST Requests              | Must accept JSON-formatted POST requests                                         |
| Response Time                     | Should return a 2xx response within 30 seconds                                   |
| Idempotency Support               | Use `X-Moe-Delivery-ID` to identify unique deliveries and avoid duplicates       |
| Signature Verification (Optional) | Strongly recommend validating `X-Moe-Signature` or `X-Moe-Signature-256` headers |

### 🔐 Example Receiving URL

```text
https://your-service.com/webhook-endpoint
```

> ⚠️ **Note**: Avoid using local addresses like `localhost`. It's recommended to use tools
> like [ngrok](https://ngrok.com) or [localtunnel](https://theboroer.github.io/localtunnel-www/) to expose local
> services
> as public URLs for testing.

### 🧾 Request Headers

| Header Name           | Sample Value                           | Description                              |
|-----------------------|----------------------------------------|------------------------------------------|
| `User-Agent`          | `Moego/Webhook-1.0`                    | Identifies Moego as the source.          |
| `X-Moe-Client-Id`     | `018e5b36-e35c-7925-a9de-321ed638b682` | Unique client identifier.                |
| `X-Moe-Event-Type`    | `HEALTH_CHECK`                         | Type of event.                           |
| `X-Moe-Delivery-ID`   | `whkdXfP`                              | Unique ID for this delivery.             |
| `X-Moe-Webhook-ID`    | `whk01`                                | Unique ID for this Webhook.              |
| `X-Moe-Nonce`         | `780525611260542810`                   | Random string for replay prevention.     |
| `X-Moe-Timestamp`     | `1751284717825`                        | Unix timestamp (in milliseconds).        |
| `X-Moe-Signature`     | `base64_encoded_sha1_signature`        | SHA1-based HMAC signature.               |
| `X-Moe-Signature-256` | `base64_encoded_sha256_signature`      | SHA256-based HMAC signature (preferred). |

These headers are used for **security and idempotency**, especially when you enable signing with a `secret_token`.

### 📦 Request Body Format (JSON)

```json
{
  "id": "a21eb4cd-367e-485f-932a-397a0951b709",
  "type": "HEALTH_CHECK",
  "timestamp": "2025-06-30T11:58:36.740351808Z",
  "companyId": "encoded_company_id",
  "payload": {
    "validation": "hello world"
  }
}
```

#### Fields Explained

| Field       | Type   | Description                                               |
|-------------|--------|-----------------------------------------------------------|
| `id`        | string | Unique event ID                                           |
| `type`      | string | Event type (e.g., `HEALTH_CHECK`, `APPOINTMENT_CREATED`, `CUSTOMER_CREATED`)  |
| `timestamp` | string | ISO 8601 formatted timestamp                              |
| `companyId` | string | Encoded company ID associated with the event              |
| `payload`   | bytes  | Varies based on event type.  The string encoded in base64 |

---


---

## 2️⃣ Step 2: Create Webhook via API

You can register a new Webhook using Moego's RESTful API.

### 📡 CreateWebhook API

- **Method**: `POST`
- **Path**: `/v1/webhooks`
- **See details**: [CreateWebhook - webhook.md](./webhook.md#1-create-webhook-createwebhook)

#### ✅ Required Fields

| Field         | Type   | Description                      |
|---------------|--------|----------------------------------|
| `endpointUrl` | string | Webhook receiving endpoint URL   |
| `eventTypes`  | array  | List of event types to subscribe |

#### 🎨 Example Request

```json
{
  "endpointUrl": "https://your-service.com/webhook-endpoint",
  "eventTypes": [
    "HEALTH_CHECK",
    "APPOINTMENT_CREATED",
    "CUSTOMER_CREATED",
    "CUSTOMER_UPDATED",
    "CUSTOMER_DELETED"
  ],
  "secretToken": "your_secure_token_here",
  "isActive": true,
  "verifySsl": true,
  "headers": {
    "Authorization": {
      "values": [
        "Bearer your_api_key"
      ]
    }
  }
}
```

#### 📦 Example Response

```json
{
  "id": "whk_001",
  "endpointUrl": "https://your-service.com/webhook-endpoint",
  "eventTypes": [
    "HEALTH_CHECK",
    "APPOINTMENT_CREATED"
  ],
  "isActive": true,
  "createdTime": "2025-06-30T11:58:36Z"
}
```

> 💡 **Tip**: Save the returned `id` field, which can be used later for triggering tests or querying logs.
> ⚠️ **Note**: When creating a Webhook, make sure to save the `secret_token`, which is used to verify request
> signatures. You can also query existing Webhooks using the `ListWebhooks`
> API. [ListWebhooks - webhook.md](./webhook.md#5-list-webhooks-listwebhooks)

---

## 3️⃣ Step 3: Verify Webhook Request Source (Optional)

To prevent forged requests, Moego includes signature fields in outgoing Webhook requests.

### 🔐 Signature Header Descriptions

| Header Name           | Description                                |
|-----------------------|--------------------------------------------|
| `X-Moe-Signature`     | SHA1-generated signature (not recommended) |
| `X-Moe-Signature-256` | SHA256-generated signature (recommended)   |

### 🔁 Signature Generation Logic

The signature is generated by concatenating the following fields and performing HMAC encryption:

```
clientID + nonce + timestamp + body
```

### 🧪 Go Example Code

```go
func isValidWebhookSignature(receivedSig string, secretToken string,
clientID string, nonce string, timestamp string, body []byte) bool {
raw := clientID + nonce + timestamp + string(body)
h := hmac.New(sha256.New, []byte(secretToken))
h.Write([]byte(raw))
expectedSig := base64.StdEncoding.EncodeToString(h.Sum(nil))

return hmac.Equal([]byte(receivedSig), []byte(expectedSig))
}
```

---

## 4️⃣ Step 4: Test Webhook Functionality

### 🧪 TriggerTestWebhookDelivery API

- **Method**: `POST`
- **Path**: `/v1/webhooks/{id}/test`
- **See details
  **: [TriggerTestWebhookDelivery - webhook.md](./webhook.md#6-trigger-test-webhook-triggertestwebhookdelivery)

#### 📥 Sample Request

```json
{
  "eventType": "HEALTH_CHECK",
  "payload": "aGVsbG8gd29ybGQ="
}
```

#### 📤 Expected Request Body

```json
{
  "id": "evt_health_check",
  "type": "HEALTH_CHECK",
  "timestamp": "2025-06-30T11:58:36.740351808Z",
  "companyId": "cmp_001",
  "payload": {
    "validation": "hello world"
  }
}
```

---

## 5️⃣ Step 5: Handle Failures and Retries

### 🔁 Automatic Retry Policy

- Moego automatically retries failed requests up to **3 times**
- You can manually redeliver failed deliveries using `RedeliverWebhookDelivery`.
  See [RedeliverWebhookDelivery - webhook.md](./webhook.md#9-redeliver-webhook-delivery-redeliverwebhookdeliveries)

### 📈 View Delivery Logs

Use `ListWebhookDeliveries` to retrieve historical records:

```json
{
  "webhookId": "whk_001",
  "filter": {
    "success": false
  }
}
```

- **See details
  **: [ListWebhookDeliveries - webhook.md](./webhook.md#8-list-all-webhook-deliveries-listwebhookdeliveries)

### 📞 Recommended Responses

- Return HTTP status code `200 OK` immediately upon successful request receipt
- Avoid asynchronous processing before returning a 2xx response, as it may cause timeouts

---

## 6️⃣ Step 6: Monitor Webhook Operational Status

### 📊 Recommended Practices

- Periodically call `ListWebhookDeliveries` to check delivery success rate
- Record `X-Moe-Delivery-ID` and `X-Moe-Nonce` to prevent duplicate processing
- Set up alert rules to monitor failure rates or delays

---

## 🛡️ Security Best Practices

- ✅ Enable `HTTPS` and `verify_ssl` for encrypted communication
- ✅ Use `secret_token` for request signature verification
- ✅ Set appropriate custom request headers (e.g., Bearer Token)
- ✅ Control Webhook access permissions
- ✅ Regularly clean up invalid or expired Webhooks
- ✅ Implement replay attack protection on the receiver side

---

## 📌 Summary

Key steps to successfully integrate with Moego Webhook:

1. Prepare an HTTPS endpoint that supports POST requests
2. Create a Webhook via API and subscribe to required event types
3. Implement signature verification to ensure requests originate from Moego
4. Use test APIs to verify if the Webhook receives data correctly
5. Handle failed requests and enable automatic/manual retry mechanisms
6. Monitor Webhook delivery status and performance metrics

---

## 📌 Error Codes

| Error Code           | Description                                                               |
|----------------------|---------------------------------------------------------------------------|
| `ALREADY_EXISTS`     | A Webhook with the same `endpointUrl` already exists                      |
| `NOT_FOUND`          | Requested Webhook ID or Delivery ID does not exist                        |
| `PERMISSION_DENIED`  | Current user lacks permission to perform this action                      |
| `INVALID_ARGUMENT`   | Invalid request parameters (e.g., malformed URL or missing fields)        |
| `INTERNAL`           | Internal system error; try again later                                    |
| `RESOURCE_EXHAUSTED` | System limit exceeded (e.g., max Webhook count or push frequency reached) |

> For more error codes, see [Common Error Codes - webhook.md](./webhook.md#-10-common-error-codes)

---

## ❓ Frequently Asked Questions (FAQ)

### Q: How do I know if my Webhook is working?

A: Use the `TriggerTestWebhookDelivery` API to send a test event and check if your endpoint receives the request.

### Q: How can I prevent duplicate Webhook deliveries?

A: Use the `X-Moe-Delivery-ID` header to uniquely identify each delivery and avoid reprocessing.

### Q: What does RESOURCE_EXHAUSTED mean when creating a Webhook?

A: This means your account has reached the maximum number of allowed Webhooks. Contact your admin to adjust quotas or
delete unused Webhooks.

### Q: What content formats are supported by Webhooks?

A: JSON (`CONTENT_TYPE_JSON`) is currently supported. Additional formats may be added in the future.

### Q: How long are Webhook logs retained?

A: Logs are retained for 15 days by default. They are automatically deleted after this period, so we recommend regular
export for auditing purposes.

### Q: Are failed Webhook requests retried?

A: Yes, Moego automatically retries failed requests up to 3 times. You can also manually redeliver them using
`RedeliverWebhookDelivery`.

---

## 📎 Related Resources

- [Webhook API Reference (Full Interface Definition)](./webhook.md)
- [Event Types Reference (Supported Events)](./event.md)
- [Protocol Buffer Definition File](../moego/business/webhook/v1/webhook.proto)

---