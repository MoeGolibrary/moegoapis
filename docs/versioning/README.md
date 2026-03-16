# API Versioning Implementation

**Current Version: 1.1.1**

## Version Storage

Version is stored in multiple locations:

```yaml
# config/config.yaml
api:
  version: "1.1.1"
```

```go
// internal/consts/version.go (to be created)
const APIVersion = "1.1.1"
```

## Version Discovery

### Response Header

```
X-API-Version: 1.1.1
```

### Version Endpoint

```
GET /v1/version

Response:
{
  "version": "1.1.1",
  "deployed_at": "2026-02-03T10:00:00Z"
}
```

## Implementation TODO

- [ ] Add version configuration to config.yaml
- [ ] Add X-API-Version header to all responses
- [ ] Implement GET /v1/version endpoint
- [ ] Create version constant in code

---

**For changelog and version rules, see [CHANGELOG.md](../CHANGELOG.md)**
