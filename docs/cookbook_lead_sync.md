# 🔄 Cookbook: Sync Leads from Another Platform

This cookbook covers one end-to-end use case: **a prospect is captured somewhere that is not
MoeGo — a Meta/Facebook Lead Ad, a website form, a custom app — and you need that prospect to
become a lead in MoeGo, routed to the right location, with the conversion reported back to the
originating platform.**

Meta Lead Ads is used as the running example because it is the hardest case: Meta's native lead
form never sends the prospect to an external URL, so MoeGo's booking-page tracking cannot see it.
Everything here applies to any source that can call an HTTP endpoint.

All endpoints in this guide are **already live in production**. No configuration change, allowlist
entry, or MoeGo-side deployment is required.

---

## 📋 What You Will Build

| Step | What it does | Endpoint |
|------|--------------|----------|
| 1 | Resolve your company and its locations | `POST /v1/companies:list`, `POST /v1/businesses:list` |
| 2 | Look up referral sources (optional) | `POST /v1/setting/companies/{companyId}/customer/referral_sources:list` |
| 3 | Create the lead when the ad form is submitted | `POST /v1/leads` |
| 4 | Subscribe to appointment events for conversion reporting | `POST /v1/webhooks` |
| 5 | Resolve contact details when an appointment is booked | `GET /v1/customers/{id}` |

---

## 🔑 Prerequisites

**Base URL**

```text
https://openapi.moego.pet
```

**Authentication** — every request carries your API key in a `Basic` authorization header. There is
no token exchange or refresh step.

```http
Authorization: Basic YOUR_API_KEY
Content-Type: application/json
```

**Access scope** — your API key is bound to a set of organizations (enterprise, company, or
business). Every endpoint below is filtered by that binding, so a key scoped to one company cannot
read or write another. If a call returns `PERMISSION_DENIED` on data you expect to see, the key's
organization restrictions are the first thing to check.

> ⚠️ **Enterprise requirement for webhooks.** Appointment webhook events are only delivered for
> companies that belong to an enterprise. A standalone company that is not under an enterprise will
> create webhooks successfully but never receive appointment deliveries. Steps 1–3 (lead creation)
> have no such requirement.

---

## 1️⃣ Resolve Your Company and Locations

IDs in the MoeGo API are **obfuscated strings**, not the numeric IDs you may see elsewhere in the
product. You must fetch them rather than construct them. Do this once at plugin setup, and again
whenever a location is added — there is no need to call it on every lead.

> 🔒 **Treat every id as an opaque string.** Ids carry a short type prefix (`cop`, `biz`, `cux`,
> `ldg`, `apt`, `crs`, `whk`) followed by an encoded body, with **no separator** and no fixed
> length. Do not split on a delimiter, do not assume a length, and do not derive one id from
> another. Store them as variable-length strings and compare them whole.

### Get the company

```http
POST /v1/companies:list
```

```json
{
  "pagination": { "pageSize": 100, "pageToken": "1" }
}
```

```json
{
  "companies": [
    { "id": "copH7x2", "name": "Northside Pet Co." }
  ]
}
```

### Get the locations

```http
POST /v1/businesses:list
```

```json
{
  "companyId": "copH7x2",
  "pagination": { "pageSize": 100, "pageToken": "1" }
}
```

```json
{
  "businesses": [
    { "id": "biz9Kp4", "name": "Downtown" },
    { "id": "biz3Rt8", "name": "Westside" }
  ]
}
```

**Store the mapping from your ad-platform form to the MoeGo business ID.** For Meta Lead Ads, key it
by the Meta form ID:

| Meta form ID | MoeGo `preferredBusinessId` |
|---|---|
| `1203948570` | `biz9Kp4` |
| `1203948571` | `biz3Rt8` |

Re-running `businesses:list` at setup and on each new location is what keeps this table current
without a code change.

📎 Full reference: [business.md](./business.md), [company.md](./company.md)

---

## 2️⃣ Look Up Referral Sources (Optional)

`referralSource` records where the lead came from, so campaign attribution stays visible inside
MoeGo's own reporting. It is a **reference to a company-configured value**, not free text — you pass
an `id` that already exists.

```http
POST /v1/setting/companies/{companyId}/customer/referral_sources:list
```

```json
{ "companyId": "copH7x2" }
```

```json
{
  "referralSources": [
    { "id": "crs2Mn", "name": "Facebook" },
    { "id": "crs7Bq", "name": "Walk-in" }
  ]
}
```

Cache the list at setup. If no source matches your platform, ask the business owner to create one in
MoeGo settings first; do not omit the field and lose attribution silently.

📎 Full reference: [setting_customer.md](./setting_customer.md#2-referralsource)

---

## 3️⃣ Create the Lead

Called once per ad-form submission. This is the only endpoint on the hot path.

```http
POST /v1/leads
```

```json
{
  "lead": {
    "companyId": "copH7x2",
    "firstName": "John",
    "lastName": "Doe",
    "phone": "+12125551234",
    "email": "john.doe@example.com",
    "preferredBusinessId": "biz9Kp4",
    "referralSource": { "id": "crs2Mn" }
  }
}
```

Returns the created `Lead`, whose `id` has the form `ldg…`. **Persist that id against your own
record of the ad submission** — you will need it for reconciliation, and it is not derivable later.

### Field notes

| Field | Notes |
|---|---|
| `companyId` | Required. From step 1. |
| `phone` | The contract specifies **E.164** (`+12125551234`). Note that this API layer does not validate the format — it passes the value straight through — so a malformed number is accepted here and may fail, or be silently stored as-is, further down. Normalize at your own boundary rather than relying on rejection. |
| `preferredBusinessId` | This is the multi-location routing key. Omitting it creates a company-level lead with no location. |
| `referralSource` | Pass `{ "id": "…" }` only. A name without a matching id is not accepted. |
| `pets`, `address`, `customFields`, `complianceConfig` | All optional. See [lead.md](./lead.md) for the full model. |

### ⚠️ Lead creation is not idempotent

There is no server-side deduplication contract. Submitting the same prospect twice can produce two
leads. **Duplicate suppression is the integration's responsibility**, and there are two layers worth
implementing:

**Layer 1 — record your own event id.** Meta delivers a `leadgen_id` with every submission. Store it
before calling MoeGo and skip any repeat. This catches the common case: the ad platform retrying
delivery to your endpoint.

**Layer 2 — check by phone before creating.** This catches the genuinely new submission from a
prospect who already exists.

> ⚠️ **`CreateLead` is not atomic when you send `pets`.** The lead record is written first, then each
> pet is created in a follow-up call. If a pet fails, the call returns an error **even though the lead
> already exists** — so a naive retry on that error produces a duplicate lead. If you send pets, treat
> a failed `CreateLead` as *state unknown* and run the phone lookup below before retrying. Sending no
> pets avoids the window entirely; you can add them later with `UpdateLead`.

```http
POST /v1/leads:list
```

```json
{
  "companyId": "copH7x2",
  "pagination": { "pageSize": 20, "pageToken": "1" },
  "filter": { "mainPhoneNumber": "+12125551234" }
}
```

An empty `leads` array means it is safe to create.

> Note that `filter.mainPhoneNumber` matches the lead's main phone number. It does not search
> secondary contacts, and it does not match on email — there is no email filter on this endpoint.

📎 Full reference: [lead.md](./lead.md)

---

## 4️⃣ Subscribe to Appointment Events

Register your endpoint once, at setup. To close the attribution loop you want `APPOINTMENT_CREATED`;
`HEALTH_CHECK` is worth subscribing to as well so you can verify the endpoint without waiting for a
real booking.

```http
POST /v1/webhooks
```

```json
{
  "endpointUrl": "https://your-service.com/moego-webhook",
  "eventTypes": ["HEALTH_CHECK", "APPOINTMENT_CREATED"],
  "secretToken": "a_long_random_string_you_generate",
  "isActive": true,
  "verifySsl": true
}
```

**Save the `secretToken` you sent.** It is what makes signature verification possible, and it is not
retrievable afterwards.

Verify the endpoint immediately with a test delivery rather than waiting for a real appointment:

```http
POST /v1/webhooks/{id}/test
```

```json
{ "eventType": "HEALTH_CHECK", "payload": "aGVsbG8gd29ybGQ=" }
```

### Endpoint requirements

| Requirement | Detail |
|---|---|
| HTTPS | Plain HTTP is rejected. |
| Method | `POST`, JSON body. |
| Response | Return `2xx` within **30 seconds**. Acknowledge first, process afterwards. |
| Retries | MoeGo retries a failed delivery up to **3 times**. Failures can also be replayed manually via `RedeliverWebhookDeliveries`. |
| Idempotency | Deduplicate on the `X-Moe-Delivery-ID` header. |
| Log retention | Delivery logs are kept **15 days**. |

### What the request actually looks like

The body is the JSON-serialized event. **The event payload is a `oneof`, which means the payload
appears under a key named after its type — there is no wrapper field called `payload`, and the
content is not base64-encoded.** For an appointment event the key is `appointment`:

```json
{
  "id": "a21eb4cd-367e-485f-932a-397a0951b709",
  "type": "APPOINTMENT_CREATED",
  "timestamp": "2026-08-21T11:58:36.740Z",
  "companyId": "copH7x2",
  "appointment": {
    "id": "apt5Wz1",
    "businessId": "biz9Kp4",
    "customerId": "cux8Qn3",
    "status": "CONFIRMED",
    "duration": { "startTime": "…", "endTime": "…" },
    "totalAmount": { "currencyCode": "USD", "units": "85" }
  }
}
```

Other event families use their own key: `customer` for `CUSTOMER_*`, `pet` for `PET_*`,
`onlineBooking` for `ONLINE_BOOKING_*`, and `healthCheck` for `HEALTH_CHECK`. Parse defensively —
read the `type` field first and select the payload key from it.

### Headers

| Header | Purpose |
|---|---|
| `X-Moe-Event-Type` | Event type, so you can route before parsing the body. |
| `X-Moe-Delivery-ID` | Unique per delivery. **Use this for idempotency.** |
| `X-Moe-Webhook-Id` | Which webhook produced this. |
| `X-Moe-Client-Id` | Your client identifier; part of the signed string. |
| `X-Moe-Nonce` | Replay protection; part of the signed string. |
| `X-Moe-Timestamp` | Unix milliseconds; part of the signed string. |
| `X-Moe-Signature-256` | HMAC-SHA256 signature. **Verify this one.** |
| `X-Moe-Signature` | HMAC-SHA1 signature. Legacy — prefer the 256 variant. |

### Verifying the signature

The signed string is the concatenation `clientId + nonce + timestamp + rawBody`, HMAC'd with your
`secretToken` and base64-encoded. Verify against the **raw request body**, before any JSON parsing —
re-serializing changes the bytes and the signature will not match.

```php
<?php
function moego_verify_signature(string $rawBody, array $headers, string $secretToken): bool {
    $raw = $headers['X-Moe-Client-Id']
         . $headers['X-Moe-Nonce']
         . $headers['X-Moe-Timestamp']
         . $rawBody;

    $expected = base64_encode(hash_hmac('sha256', $raw, $secretToken, true));

    return hash_equals($expected, $headers['X-Moe-Signature-256']);
}
```

Use a constant-time comparison (`hash_equals` above, `hmac.Equal` in Go). A plain `==` on the
signature is a timing oracle.

📎 Full reference: [webhook.md](./webhook.md), [webhook_integration.md](./webhook_integration.md),
[event.md](./event.md)

---

## 5️⃣ Resolve Contact Details for the Conversion Signal

> 🚩 **This is the step most integrations get wrong.**

Ad platforms match conversions on hashed contact details — email and phone. **The appointment
payload does not contain them.** It carries `customerId` and nothing else about the person:

```json
"appointment": { "customerId": "cux8Qn3", "businessId": "biz9Kp4", "…": "…" }
```

So the webhook handler needs a second call:

```http
GET /v1/customers/cux8Qn3
```

```json
{
  "id": "cux8Qn3",
  "firstName": "John",
  "lastName": "Doe",
  "phone": "+12125551234",
  "email": "john.doe@example.com",
  "preferredBusinessId": "biz9Kp4"
}
```

Hash `email` and `phone` per your ad platform's requirements (Meta's Conversions API expects
lowercase, trimmed, SHA-256) and send the conversion.

Do this **after** returning `2xx` to MoeGo. Two outbound calls — `GetCustomer` plus the ad
platform — will not reliably finish inside the 30-second delivery window, and a timeout costs you a
retry you did not need.

### 🔗 Linking the appointment back to the original lead

**Conversion happens in place.** A lead and the customer it becomes are the *same underlying record* —
converting flips its type from lead to customer, and the record's internal id does not change. This is
why a converted lead stops appearing in `leads:list`: that endpoint filters on the lead type, not on
the record's existence.

What this means in practice is narrower than it sounds. The two ids are **not interchangeable as
strings**: a lead id carries the `ldg` prefix and a customer id carries `cux`, each encoded
separately, so passing a lead id to `GET /v1/customers/{id}` fails to decode and returns
`INVALID_ARGUMENT`. There is also no endpoint that translates one into the other, and the public
event vocabulary contains no lead or conversion event at all — nothing notifies you that a conversion
happened.

So reconcile on contact details — match the customer's `email` or `phone` against the lead you stored
in step 3:

```text
ad submission → lead (ldg…)      ─┐
                                  ├─ matched on email / phone
appointment   → customer (cux…)  ─┘
```

This matches more reliably than a contact-detail join usually would, precisely because of the
in-place conversion: the customer carries forward the very `email` and `phone` that were written when
the lead was created, so barring a later edit the values are identical rather than merely similar.
And the conversion signal you send is *already* keyed on hashed email and phone, so the same details
that identify the customer are the ones the ad platform matches on — you never need the internal id.

Two cases still miss. Someone edits the contact details in MoeGo between capture and booking, which
breaks the join; or the prospect was never converted at all and was entered as a fresh customer
instead, which produces a genuinely unrelated record. Treat unmatched conversions as expected
background noise rather than a bug, and log them for periodic review.

---

## ⚠️ Error Handling

| Status | Meaning | Recommended action |
|---|---|---|
| `INVALID_ARGUMENT` | An id that could not be decoded — a `companyId`, `preferredBusinessId`, `referralSource.id`, or `lead.id` that is malformed or carries the wrong type prefix. Also raised when a `customFields` key does not match a configured field, or its value type does not match the field's type. | **Do not retry.** Log the full response body; this is a data or mapping bug. |
| `PERMISSION_DENIED` | The API key's organization restrictions do not cover this company or business. | **Do not retry.** Configuration issue; alert an admin. |
| `NOT_FOUND` | The id does not exist, or is outside your key's scope. | Do not retry. |
| `ALREADY_EXISTS` | For webhooks: an active webhook with the same `endpointUrl` already exists. | Reuse the existing webhook rather than creating another. |
| `RESOURCE_EXHAUSTED` | Webhook count or delivery rate limit reached. | Back off. Check for orphaned webhooks from earlier setup attempts. |
| `INTERNAL` / `5xx` | Transient server-side failure. | **Retry with exponential backoff** — the only class worth retrying. One exception: on `CreateLead` *with pets*, the lead may already exist despite the error, so run the phone lookup before retrying. |

**On the inbound side**, acknowledge the ad platform immediately — return `200` before you attempt
the MoeGo call — so the platform does not retry on your processing latency. Then create the lead
asynchronously with your own retry schedule. If the retries are exhausted, notify an admin: an ad
lead that never reached the CRM is invisible otherwise.

---

## ✅ Setup Checklist

- [ ] API key issued, and its organization restrictions cover every location you will write to
- [ ] `companies:list` and `businesses:list` called; form-to-business mapping stored
- [ ] Referral source id resolved, or deliberately skipped
- [ ] Ad-platform event id persisted before each `CreateLead`, for duplicate suppression
- [ ] Phone numbers normalized to E.164 at the boundary
- [ ] Webhook registered; `secretToken` stored securely
- [ ] Endpoint verified with a `HEALTH_CHECK` test delivery
- [ ] Signature verification implemented against the **raw** body, with constant-time comparison
- [ ] `X-Moe-Delivery-ID` deduplication in place
- [ ] Webhook handler returns `2xx` **before** calling `GetCustomer` and the ad platform
- [ ] Unmatched-conversion logging in place

---

## 📎 Related Documentation

- [lead.md](./lead.md) — full Lead model and endpoints
- [customer.md](./customer.md) — Customer model
- [business.md](./business.md) / [company.md](./company.md) — organization resolution
- [setting_customer.md](./setting_customer.md) — referral sources, life cycle, action status
- [webhook.md](./webhook.md) — webhook API reference
- [webhook_integration.md](./webhook_integration.md) — general webhook integration guide
- [event.md](./event.md) — event types and payload shapes
