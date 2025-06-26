# ⚖️ Weight (moego.common.v1.Weight)

## 📌 1. Overview

The `Weight` message provides a standardized way to represent and convert weight measurements throughout the system. It
supports multiple units and ensures consistent weight tracking for pets, products, and other resources that require
weight measurements.

## 🧩 2. Core Fields

| Field Name | Type   | Description                     |
|------------|--------|---------------------------------|
| `value`    | uint32 | The numeric value of the weight |
| `unit`     | Unit   | The unit of measurement         |

### Enum: Unit

- `UNIT_UNSPECIFIED` - Not intended for direct use.
- `KILOGRAM` - Base unit for weight (1 kg = 2.20462 lbs)
- `POUND` - Common in US and some other regions (1 lb = 0.453592 kg)

## 📦 3. Example JSON

```json
{
  "value": 10,
  "unit": "KILOGRAM"
}
```

---

# 📎 Related File References

- [weight.proto](../../moego/common/v1/weight.proto)
