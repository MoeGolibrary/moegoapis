# 📦 Address (moego.common.v1.Address)

## 📌 1. Overview

The `Address` message represents a physical location used throughout the system. It provides a standardized way to store
and validate address information for businesses, customers, and service locations. The address format follows common
postal standards and includes geocoding information for mapping.

## 🧩 2. Core Fields

| Field Name    | Type   | Description                                                      |
|---------------|--------|------------------------------------------------------------------|
| `id`          | string | Unique identifier                                                |
| `address1`    | string | Street address including house number and street name            |
| `address2`    | string | Additional address details like apartment, suite, or unit number |
| `city`        | string | City or locality name                                            |
| `state`       | string | State, province, or region                                       |
| `postal_code` | string | Postal or ZIP code                                               |
| `country`     | string | Country name or ISO country code                                 |
| `coordinate`  | LatLng | Latitude and longitude coordinate for mapping                    |
| `type`        | Type   | Represents the importance of the address (PRIMARY / SECONDARY)   |

### Enum: Type

- `TYPE_UNSPECIFIED` - Not intended for direct use. Indicates a system anomaly.
- `PRIMARY` - Primary Address, which will be used as first address choice.
- `SECONDARY` - Secondary Address, an optional address saved as back-up.

## 📦 3. Example JSON

```json
{
  "id": "addr_001",
  "address1": "123 Main Street",
  "address2": "Floor 3",
  "city": "New York",
  "state": "NY",
  "postal_code": "10001",
  "country": "US",
  "coordinate": {
    "latitude": 40.7128,
    "longitude": -74.0060
  },
  "type": "PRIMARY"
}
```

---

# 📎 Related File References

- [address.proto](../../moego/common/v1/address.proto)
