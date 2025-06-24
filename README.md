# MoeGo APIs

This repository contains the original interface definitions of public
MoeGo APIs that support REST protocols. Reading the
original interface definitions can provide a better understanding of
MoeGo APIs and help you to utilize them more efficiently. You can also
use these definitions with open source tools to generate client
libraries, documentation, and other artifacts.

## Try it now

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 96px; height: 24px;">](https://app.getpostman.com/run-collection/30555124-2c27c225-3a6c-4234-bd4f-f818dc0a757c?source=rip_markdown&collection-url=entityId%3D30555124-2c27c225-3a6c-4234-bd4f-f818dc0a757c%26entityType%3Dcollection%26workspaceId%3Ddb8fd53e-fafa-4b5e-bf78-8dd330f46a0b)

When you click the `Run in Postman` button, you can try out the MoeGo APIs in Postman. Postman is a popular API client
that makes it easy to explore and test APIs. You can use the MoeGo APIs in Postman without writing any code.

**Note**: replace the collection Authorization `${Base_64_API_Key}` with your own API Key encoded with Base64.

## 🌐 Domain Structure

The MoeGo APIs are organized into one domain:

```
openapi.moego.pet
```

Each path corresponds to a specific version and resource:

- `/v1/auth`
- `/v1/appointment`
- `/v1/customer`
- `/v1/webhooks`

Each product contains one or more APIs for different use cases such as authentication, data management, and event
notifications.

---

## 🔄 Webhook Integration

Webhooks allow you to receive real-time updates when certain events occur within the MoeGo system. This provides an
efficient way to integrate external services with the platform.

### 📁 Location:

- [moego/business/webhook/README.md](moego/business/webhook/README.md)

### 📌 Features:

- Subscribe to business events like `APPOINTMENT_CREATED`, `APPOINTMENT_FINISHED`
- Configure delivery endpoints and HMAC signing
- View and retry webhook delivery logs
- Set up test calls for debugging

