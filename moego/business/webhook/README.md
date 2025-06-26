# Webhook

## 📌 1. Functional Overview

Webhook is a lightweight event notification mechanism that actively pushes data to a specified URL when certain business
events occur. This interface enables:

- Creating, updating, and deleting webhook configurations;
- Listing all webhooks;
- Testing whether webhooks are functioning properly;
- Viewing and retrying webhook delivery logs.

---

### 🎯 2. Design Goals

- **Real-time Push**: Supports event-driven architecture for low-latency communication between systems.
- **Secure and Reliable**: Supports HTTPS, HMAC signing, and custom headers.
- **Easy Integration**: Provides RESTful interfaces compatible with mainstream development languages and frameworks.

Applicable to scenarios such as order status change notifications, user registration synchronization, exception alerts,
and third-party system integration.

---

## 🧩 3. Core Concepts

### 1. Webhook

Represents a webhook subscription configuration, including the target URL, subscribed event types, authentication
method, etc.

| Field Name      | Type                      | Description            |
|-----------------|---------------------------|------------------------|
| `id`            | string                    | Unique identifier      |
| `organizations` | Array<Organization\>      | List of organizations  |
| `event_types`   | Array<EventType\>         | Subscribed event types |
| `endpoint_url`  | string                    | URL to receive webhook |
| `secret_token`  | string                    | Optional HMAC token    |
| `content_type`  | ContentType               | Default is JSON        |
| `is_active`     | bool                      | Whether active         |
| `verify_ssl`    | bool                      | Whether to verify SSL  |
| `headers`       | map<string, HeaderValues> | Custom HTTP headers    |
| `created_time`  | Timestamp                 | Creation time          |
| `updated_time`  | Timestamp                 | Last update time       |

### 2. WebhookDelivery

Represents a specific webhook push record.

> By default, delivery logs are retained for up to **15 days**

| Field Name         | Type                      | Description                |
|--------------------|---------------------------|----------------------------|
| `id`               | string                    | Delivery log ID            |
| `webhook_id`       | string                    | Associated webhook ID      |
| `event_type`       | EventType                 | Event type                 |
| `event_id`         | string                    | Unique event ID            |
| `request_url`      | string                    | Request URL                |
| `delivered_to`     | string                    | Actual destination address |
| `request_headers`  | map<string, HeaderValues> | Request headers            |
| `request_body`     | bytes                     | Request body               |
| `response_status`  | int32                     | HTTP status code           |
| `response_headers` | map<string, HeaderValues> | Response headers           |
| `response_body`    | bytes                     | Response body              |
| `delivered_at`     | Timestamp                 | Delivery timestamp         |
| `duration_ms`      | int64                     | Duration in milliseconds   |
| `success`          | bool                      | Whether successful         |
| `error`            | string                    | Error message              |
| `retry_count`      | int32                     | Retry count                |
| `request_format`   | ContentType               | Request format             |
| `response_format`  | ContentType               | Response format            |

### 3. Event Type

Event types trigger webhook calls. Examples include:

- `HEALTH_CHECK`
- `APPOINTMENT_CREATED`
- `APPOINTMENT_FINISHED`

Refer to [event.proto](../../moego/business/event/v1/event.proto) for supported event types.

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Webhook

Here is a typical integration flow:

1. **Create Webhook**
    - Specify endpoint URL and interested event types.
    - Set HMAC token, headers, etc.

2. **Trigger Test Call**
    - Use `TriggerTestWebhookDelivery` to send a test event (e.g., ping).
    - Verify that the server receives the request.

3. **View Delivery Logs**
    - Use `ListWebhookDeliveries` to view historical push records.
    - Analyze failure reasons (e.g., network errors, permission issues).

4. **Update/Delete Webhook**
    - Modify endpoint or event_types.
    - Delete when no longer needed.

5. **Monitoring & Retrying**
    - Regularly check delivery success rate.
    - Use `RedeliverWebhookDelivery` to retry failed deliveries.

---

## 📦 5. API Interface Descriptions

### 1. Create Webhook (`CreateWebhook`)

- **Method**: `CreateWebhook`
- **HTTP Method**: POST
- **Path**: `/v1/webhooks`

#### ✅ Functionality:

Registers a new webhook, subscribes to specified event types, and configures the target URL and related parameters (
e.g., authentication token, headers).

#### 🎯 Use Cases:

- Users want to receive business events like order creation or payment success.
- Third-party developers subscribe to event streams under specific organizations.

#### 🔧 Request Parameters:

| Field Name      | Type                      | Required | Description                         |
|-----------------|---------------------------|----------|-------------------------------------|
| `endpoint_url`  | string                    | Yes      | URL to receive webhook              |
| `organizations` | Array<Organization\>      | No       | Organization list (empty means all) |
| `event_types`   | Array<EventType\>         | Yes      | List of event types (empty = none)  |
| `secret_token`  | string                    | No       | Optional HMAC token                 |
| `content_type`  | ContentType               | No       | Content format, default JSON        |
| `is_active`     | bool                      | No       | Whether active, default true        |
| `verify_ssl`    | bool                      | No       | Whether to verify SSL (recommended) |
| `headers`       | map<string, HeaderValues> | No       | Custom HTTP headers                 |

#### 💡 Example Request:

```json
{
  "endpoint_url": "https://your-service.com/webhook",
  "organizations": [
    {
      "id": "org_001"
    }
  ],
  "event_types": [
    "ORDER_CREATED",
    "PAYMENT_SUCCESS"
  ],
  "secret_token": "my-secret-token",
  "is_active": true,
  "verify_ssl": true,
  "headers": {
    "Authorization": {
      "values": [
        "Bearer your_token"
      ]
    }
  }
}
```

#### 📌 Return Value:

Returns the created `Webhook` object.

- Returns `RESOURCE_EXHAUSTED` if the client has reached the maximum number of webhooks allowed.

#### ⚠️ Error Codes:

- `ALREADY_EXISTS`: A webhook with the same endpoint URL already exists.
- `PERMISSION_DENIED`: Permission denied.
- `INVALID_ARGUMENT`: Invalid request parameters.
- `NOT_FOUND`: Specified organization does not exist.
- `RESOURCE_EXHAUSTED`: Maximum webhook count reached.

---

### 2. Get Webhook (`GetWebhook`)

- **Method**: `GetWebhook`
- **HTTP Method**: GET
- **Path**: `/v1/webhooks/{id}`

#### ✅ Functionality:

Retrieves a registered webhook's full configuration, including subscribed event types, target address, and status.

#### 🎯 Use Cases:

- Check current webhook configuration.
- Verify subscribed event types or organizations.
- Confirm webhook during debugging.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description            |
|------------|--------|----------|------------------------|
| `id`       | string | Yes      | Webhook ID to retrieve |

#### 📌 Return Value:

Returns the complete `Webhook` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified webhook ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 3. Update Webhook (`UpdateWebhook`)

- **Method**: `UpdateWebhook`
- **HTTP Method**: PUT
- **Path**: `/v1/webhooks/{id}`

#### ✅ Functionality:

Updates an existing webhook configuration, e.g., modify endpoint URL, adjust subscribed event types, update headers or
secret token.

#### 🎯 Use Cases:

- Change callback address (e.g., deploy new service).
- Add new event subscriptions.
- Update signature token or headers due to security changes.

#### 🔧 Request Parameters:

| Field Name      | Type                      | Required | Description               |
|-----------------|---------------------------|----------|---------------------------|
| `id`            | string                    | Yes      | Webhook ID to update      |
| `endpoint_url`  | string                    | Yes      | New endpoint URL          |
| `organizations` | Array<Organization\>      | No       | Updated organization list |
| `event_types`   | Array<EventType\>         | Yes      | Updated event types       |
| `secret_token`  | string                    | No       | Updated secret token      |
| `content_type`  | ContentType               | No       | Updated content type      |
| `is_active`     | bool                      | No       | Whether active            |
| `verify_ssl`    | bool                      | No       | Whether to verify SSL     |
| `headers`       | map<string, HeaderValues> | No       | Custom HTTP headers       |

#### 📌 Return Value:

Returns the updated `Webhook` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified webhook ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 4. Delete Webhook (`DeleteWebhook`)

- **Method**: `DeleteWebhook`
- **HTTP Method**: DELETE
- **Path**: `/v1/webhooks/{id}`

#### ✅ Functionality:

Deletes a registered webhook and stops all its event pushes.

#### 🎯 Use Cases:

- Stop unnecessary notifications.
- Clean up invalid webhooks in testing environments.

#### 🔧 Request Parameters:

| Field Name | Type   | Required | Description          |
|------------|--------|----------|----------------------|
| `id`       | string | Yes      | Webhook ID to delete |

#### 📌 Return Value:

Returns an empty `DeleteWebhookResponse`.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified webhook ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 5. List Webhooks (`ListWebhooks`)

- **Method**: `ListWebhooks`
- **HTTP Method**: POST
- **Path**: `/v1/webhooks:list`

#### ✅ Functionality:

Lists all webhooks under the current account, supporting filtering by status, event types, and time range.

#### 🎯 Use Cases:

- View all webhooks' status and subscriptions.
- Audit or debug webhook configurations.
- Monitor active/inactive webhook counts.

#### 🔧 Request Parameters:

| Field Name            | Type              | Required | Description                            |
|-----------------------|-------------------|----------|----------------------------------------|
| `pagination`          | Pagination        | Yes      | Pagination info: page_size, page_token |
| `filter.is_active`    | bool              | No       | Filter by active status                |
| `filter.event_types`  | Array<EventType\> | No       | Filter by event types                  |
| `filter.created_time` | Interval          | No       | Filter by creation time                |
| `filter.updated_time` | Interval          | No       | Filter by update time                  |

#### 📌 Return Value:

Returns paginated results and webhook list.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 6. Trigger Test Webhook (`TriggerTestWebhookDelivery`)

- **Method**: `TriggerTestWebhookDelivery`
- **HTTP Method**: POST
- **Path**: `/v1/webhooks/{id}/test`

#### ✅ Functionality:

Manually triggers a test event delivery to verify whether the webhook is working correctly. Sends a "ping" event by
default.

#### 🎯 Use Cases:

- Test if the webhook endpoint can receive and process requests.
- Verify that the signature token and headers are configured correctly.

#### 🔧 Request Parameters:

| Field Name   | Type      | Required | Description                       |
|--------------|-----------|----------|-----------------------------------|
| `id`         | string    | Yes      | Webhook ID to test                |
| `event_type` | EventType | No       | Custom event type (default: PING) |
| `payload`    | bytes     | No       | Custom payload content            |

#### 📌 Return Value:

Returns the `WebhookDelivery` log object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified webhook ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 7. Get Webhook Delivery Log (`GetWebhookDelivery`)

- **Method**: `GetWebhookDelivery`
- **HTTP Method**: GET
- **Path**: `/v1/webhook/deliveries/{id}`

#### ✅ Functionality:

Retrieves a specific webhook push record, including detailed request/response information.

#### 🎯 Use Cases:

- Check whether a delivery was successful.
- Troubleshoot failures (e.g., network errors, permissions).
- Audit delivery history.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description           |
|--------------|--------|----------|-----------------------|
| `id`         | string | Yes      | Delivery log ID       |
| `webhook_id` | string | Yes      | Associated webhook ID |

#### 📌 Return Value:

Returns the `WebhookDelivery` object.

#### ⚠️ Error Codes:

- `NOT_FOUND`: Specified log ID does not exist.
- `PERMISSION_DENIED`: Permission denied.

---

### 8. List All Webhook Deliveries (`ListWebhookDeliveries`)

- **Method**: `ListWebhookDeliveries`
- **HTTP Method**: POST
- **Path**: `/v1/webhook/deliveries:list`

#### ✅ Functionality:

Lists all push records for a given webhook, supports filtering by event type, success status, and time range.

#### 🎯 Use Cases:

- Monitor webhook delivery success rate.
- Analyze failures (e.g., timeouts, status codes).
- Audit push history.

#### 🔧 Request Parameters:

| Field Name             | Type              | Required | Description                   |
|------------------------|-------------------|----------|-------------------------------|
| `pagination`           | Pagination        | Yes      | Pagination info               |
| `webhook_id`           | string            | Yes      | Associated webhook ID         |
| `filter.event_types`   | Array<EventType\> | No       | Filter by event types         |
| `filter.success`       | bool              | No       | Filter by success status      |
| `filter.delivery_time` | Interval          | No       | Filter by delivery time range |

#### 📌 Return Value:

Returns pagination result and `WebhookDelivery` list.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

### 9. Redeliver Webhook Delivery (`RedeliverWebhookDeliveries`)

- **Method**: `RedeliverWebhookDeliveries`
- **HTTP Method**: POST
- **Path**: `/v1/webhooks/deliveries/{id}/redeliver`

#### ✅ Functionality:

Retries a failed webhook delivery, useful after fixing endpoint issues.

#### 🎯 Use Cases:

- Resend failed messages after endpoint recovery.
- Debug webhook delivery logic.

#### 🔧 Request Parameters:

| Field Name   | Type   | Required | Description              |
|--------------|--------|----------|--------------------------|
| `id`         | string | Yes      | Delivery log ID to retry |
| `webhook_id` | string | Yes      | Associated webhook ID    |

#### 📌 Return Value:

Returns the redelivered `WebhookDelivery` object.

#### ⚠️ Error Code:

- `PERMISSION_DENIED`: Permission denied.

---

## 🧪 6. Usage Examples

### Example 1: Create Webhook

```json
{
  "endpoint_url": "https://your-service.com/webhook",
  "organizations": [
    {
      "id": "org_001"
    },
    {
      "id": "org_002"
    }
  ],
  "event_types": [
    "ORDER_CREATED",
    "PAYMENT_SUCCESS"
  ],
  "secret_token": "my-secret-token",
  "is_active": true,
  "verify_ssl": true,
  "headers": {
    "Authorization": {
      "values": [
        "Bearer your_token"
      ]
    }
  }
}
```

### Example 2: Trigger Test Call

```json
{
  "id": "whk_001",
  "event_type": "PING",
  "payload": "{ \"test\": \"hello world\" }"
}
```

### Example 3: Query Delivery Logs

```json
{
  "webhook_id": "whk_001",
  "pagination": {
    "page_size": 20
  },
  "filter": {
    "success": false,
    "delivery_time": {
      "start_time": "2024-08-01T00:00:00Z",
      "end_time": "2024-08-02T00:00:00Z"
    }
  }
}
```

---

## ⚠️ 7. Usage Limitations

To ensure system stability and fair resource usage, the following default limits are applied to webhook usage:

### 1. Maximum Number of Webhook Configurations per Client

- **Limit**: Up to 10
- **Description**: Prevents abuse or malicious registration of large numbers of webhooks. Contact admin or adjust quota
  if higher limits are required.

### 2. Maximum Pushes per Minute per Webhook

- **Limit**: Up to 100 times/minute
- **Description**: Avoids downstream service overload due to high-frequency pushes. Evaluate service capacity before
  requesting higher thresholds.

### 3. Daily Total Push Count per Webhook

- **Limit**: Up to 10,000 times/day
- **Description**: Controls long-term resource consumption to ensure overall system stability and fairness.

### 4. Monthly Total Push Count per Webhook

- **Limit**: Up to 300,000 times/month
- **Description**: Long-term resource allocation basis, suitable for multi-tenant SaaS resource management strategies.

### 5. Delivery Log Retention Period

- **Limit**: Delivery logs (`WebhookDelivery`) are retained for a maximum of **15 days** by default.
- **Description**: Logs older than 15 days are automatically deleted to manage storage and performance. For long-term
  auditing, export logs periodically.

> ⚠️ Note: If limits are exceeded, the API may return `RESOURCE_EXHAUSTED`. It is recommended that clients implement
> reasonable rate-limiting fallback logic, such as queue caching and delayed retries.

---

## 📎 8. FAQ

| Question                                                        | Answer                                                                                                                                           |
|-----------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|
| How to determine if a webhook is effective?                     | Use `TriggerTestWebhookDelivery` to send a test and observe delivery logs                                                                        |
| How to prevent duplicate webhooks?                              | System automatically detects identical `endpoint_url`, returns `ALREADY_EXISTS`                                                                  |
| What content formats does webhook support?                      | Currently supports JSON, default `CONTENT_TYPE_JSON`                                                                                             |
| How to view failed records?                                     | Use `ListWebhookDeliveries` with `filter.success = false`                                                                                        |
| Can webhook limit events from specific companies or businesses? | Yes, via the `organizations` field                                                                                                               |
| Why does creating a webhook return “resource exhausted”?        | The client may have reached the maximum allowed webhook count. Clean up unused webhooks or contact admin to increase quota.                      |
| Why does webhook delivery fail with status RESOURCE_EXHAUSTED?  | Indicates push frequency or total volume has exceeded platform limits. Try again later or optimize push logic to avoid triggering rate limiting. |

---

## 🛡️ 9. Security Recommendations

- ✅ Enable `verify_ssl` to prevent man-in-the-middle attacks.
- ✅ Use `secret_token` to sign payloads (e.g., HMAC).
- ✅ Set reasonable HTTP headers, such as Bearer Token.
- ✅ Control access to ensure only authorized users can operate webhooks.
- ✅ Avoid exposing sensitive information in `headers` or `payload`.
- ✅ Recommend cleaning up invalid webhooks regularly.
- ✅ Recommend anti-replay attack handling on the receiving side.

---

## 📌 10. Common Error Codes

| Error Code           | Description                                |
|----------------------|--------------------------------------------|
| `ALREADY_EXISTS`     | A webhook with the same URL already exists |
| `NOT_FOUND`          | Webhook or delivery ID does not exist      |
| `PERMISSION_DENIED`  | Current user has no access rights          |
| `INVALID_ARGUMENT`   | Invalid request parameters                 |
| `INTERNAL`           | Internal server error                      |
| `RESOURCE_EXHAUSTED` | Request exceeds system quota limits        |

---

## 📎 11. Related File References

- [webhook_service.proto](../../../moego/business/webhook/v1/webhook_service.proto)
- [webhook.proto](../../../moego/business/webhook/v1/webhook.proto)
- [event.proto](../../../moego/business/event/v1/event.proto)
