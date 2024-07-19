# API Key

The API key is a unique identifier that authenticates requests associated with your project for usage and billing
purposes. You can view and manage your API Key in the MoeGo Open Platform (coming soon).

When you use an API key to authenticate to an API, the API key does not identify a principal, nor does it provide any
authorization information. Therefore, the request does not use Identity and Access Management (IAM) to check whether the
caller has permission to perform the requested operation.

Moreover, the API key carries many permissions, so **be sure to keep them secure!** Do not share your secret API keys in
publicly accessible areas such as GitHub, client-side code, and so forth. It is recommended to use a server-side proxy
to handle API key requests. Alternatively, you can use restricted API keys for granular permissions.

## How to get an API Key

You can only reveal an API key **once** when you create it. If you lose it, you can’t retrieve it from the Dashboard. In
that case, rotate it or delete it and create a new one.

An API key has the following components, which you use to manage and use the key:

| Component   | Description                                                                                                                                                                |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Secret      | The API key Secret is an encrypted string, for example, `AIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGewQe`. When you use an API key to authenticate, you always use this string.    |
| ID          | The API key ID is used to uniquely identify the key. The key ID cannot be used to authenticate.                                                                            |
| Name        | The API key Name is an optional, descriptive name for the key, which you can set when you create or update the key.                                                        |
| Restriction | The API key Restriction is an optional, but recommended, setting that you can use to control the requests that can be made with the key. See [Restrictions](#restrictions) |

> **Note**: The MoeGo Open Platform is not available yet. This section will be completed when the platform is available.

## How to use an API Key

All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will
also fail.

To use an API key, you need to include it in the `Authorization` header when you make a request to the API (Refer to
the [Bacis HTTP Authentication](https://datatracker.ietf.org/doc/html/rfc7617)).

For example:

```https
GET /v1/endpoint HTTP/1.1
Host: openapi.moego.pet
Authorization: Basic <API_KEY_BASE64>
```

`<API_KEY_BASE64>` is the base64 encoded string of the API key Secret.

## Restrictions

You can restrict the use of an API key to specific websites, IP addresses, or mobile apps. This can help prevent
unauthorized use of your key. For example, you can restrict a key so that it can only be used from a specific IP address
or range of IP addresses, or only from a specific website.

The following restrictions are available:

| Restriction  | Description                                                               |
|--------------|---------------------------------------------------------------------------|
| None         | No restrictions. The API key can be used from any server.                 |
| Scopes       | The API key can only access the APIs that require the scopes.             |
| IP addresses | The API key can only be used from an IP address or range of IP addresses. |
