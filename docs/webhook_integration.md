# 🪝 Moego Webhook Integration Guide

This document describes how to properly integrate with Moego's webhook system. It covers:

- Required endpoint capabilities
- Request URL format
- Expected request headers and body
- Signature validation
- Best practices for secure integration

---

## ✅ 1. Endpoint Requirements

To successfully receive webhooks from Moego, your endpoint must meet the following requirements:

| Requirement                 | Description                                                              |
|-----------------------------|--------------------------------------------------------------------------|
| **HTTPS Support**           | Your endpoint must use HTTPS (HTTP will fail unless explicitly allowed). |
| **Accepts POST Requests**   | Must accept `POST` requests with JSON payload.                           |
| **Responds Within Timeout** | Must respond within 30 seconds to avoid timeout errors.                  |
| **Idempotency Support**     | Should handle duplicate deliveries using `X-Moe-Delivery-ID`.            |
| **Signature Validation**    | Optional but strongly recommended: validate HMAC signature in headers.   |

---

## 🌐 2. Webhook Request URL Format

Your webhook URL should be publicly accessible and follow this basic pattern:

```
https://your-service.com/webhook-endpoint
```

### Example:

```text
https://webhook-receiver.example.com/moego-webhook
```

> ⚠️ **Note**: Avoid using local development URLs like `localhost`. Use tools like [ngrok](https://ngrok.com)
> or [localtunnel](https://theboroer.github.io/localtunnel-www/) during testing.

---

## 📥 3. Webhook Delivery Structure

When an event occurs, Moego will send a `POST` request to your webhook URL containing:

- A structured JSON body with event details.
- Security headers for verification.
- Metadata about the delivery attempt.

---

## 🧾 4. Request Headers

| Header Name           | Sample Value                           | Description                              |
|-----------------------|----------------------------------------|------------------------------------------|
| `User-Agent`          | `Moego/Webhook-1.0`                    | Identifies Moego as the source.          |
| `X-Moe-Client-Id`     | `018e5b36-e35c-7925-a9de-321ed638b682` | Unique client identifier.                |
| `X-Moe-Event-Type`    | `HEALTH_CHECK`                         | Type of event.                           |
| `X-Moe-Delivery-ID`   | `whkdXfP`                              | Unique ID for this delivery.             |
| `X-Moe-Nonce`         | `780525611260542810`                   | Random string for replay prevention.     |
| `X-Moe-Timestamp`     | `1751284717825`                        | Unix timestamp (in milliseconds).        |
| `X-Moe-Signature`     | `base64_encoded_sha1_signature`        | SHA1-based HMAC signature.               |
| `X-Moe-Signature-256` | `base64_encoded_sha256_signature`      | SHA256-based HMAC signature (preferred). |

These headers are used for **security and idempotency**, especially when you enable signing with a `secret_token`.

---

## 📦 5. Request Body Format (JSON)

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

### Fields Explained

| Field       | Type   | Description                                               |
|-------------|--------|-----------------------------------------------------------|
| `id`        | string | Unique event ID                                           |
| `type`      | string | Event type (e.g., `HEALTH_CHECK`, `APPOINTMENT_CREATED`)  |
| `timestamp` | string | ISO 8601 formatted timestamp                              |
| `companyId` | string | Encoded company ID associated with the event              |
| `payload`   | bytes  | Varies based on event type.  The string encoded in base64 |

---

## 🔐 6. Signature Verification

If you configured a `secret_token` when creating the webhook, Moego will sign each request using HMAC.

### Signature Generation Logic

The signature is generated from the concatenation of:

```
clientID + nonce + timestamp + requestBody
```

Using the shared secret token via:

- **SHA1** → `X-Moe-Signature`
- **SHA256** → `X-Moe-Signature-256`

### Go Example: Validate Signature

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

## 🧪 7. Testing Your Webhook

You can test your webhook by triggering a health check using Moego’s API:

### Example: Trigger Test Webhook

```json
{
  "id": "whk_001",
  "eventType": "HEALTH_CHECK",
  "payload": "aGVsbG8gd29ybGQ="
}
```

You should receive a request with:

- `type`: `"HEALTH_CHECK"`
- `payload.validation`: `"hello world"`

Use this to verify your endpoint receives and processes webhooks correctly.

---

## 🛡️ 8. Security Recommendations

- ✅ Always use `HTTPS`.
- ✅ Enable `verify_ssl = true` for outbound SSL certificate validation.
- ✅ Store and validate `X-Moe-Signature` or `X-Moe-Signature-256` if using a `secret_token`.
- ✅ Log `X-Moe-Delivery-ID` and `X-Moe-Nonce` to prevent duplicate processing.
- ✅ Rate limit incoming requests to protect against abuse.

---

## 📋 9. Sample Payloads

### HEALTH_CHECK Event

```json
{
  "id": "evt_health_check",
  "type": "HEALTH_CHECK",
  "timestamp": "2025-06-30T11:58:36.740351808Z",
  "companyId": "cmp_001",
  "payload": {
    "validation": "test_validation_string"
  }
}
```

---

## 📈 10. Retry & Failure Handling

- Moego will retry failed deliveries up to 3 times.
- Failed deliveries can be manually retried via the `RedeliverWebhookDelivery` API.
- Responses outside the 2xx range are considered failures.

### Best Practices:

- Return `200 OK` immediately after receiving a webhook.
- Defer further processing asynchronously if needed.
- Monitor and log all delivery attempts for debugging and auditing.

---

## 📞 11. Delivering a Successful Response

Your service should return a successful HTTP status code:

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "status": "received"
}
```

Avoid returning error codes unless you want Moego to retry the delivery.

---

## 📌 12. Summary

To successfully integrate with Moego webhooks:

1. Set up a publicly accessible HTTPS endpoint.
2. Accept and parse `POST` requests with JSON body.
3. Validate signatures if a `secret_token` was provided.
4. Respond with a `200 OK` quickly.
5. Handle duplicates using `X-Moe-Delivery-ID` and `X-Moe-Nonce`.
6. Monitor delivery logs via `ListWebhookDeliveries`.

---

## 📎 13. Related Resources

- [Webhook API Reference](./webhook.md)
- [Event Types Reference](./event.md)
