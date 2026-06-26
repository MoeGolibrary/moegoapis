# Changelog

All notable changes to MoeGo Open API v1 will be documented in this file.

Format: [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)

Versioning: [Semantic Versioning](https://semver.org/spec/v2.0.0.html) (`MAJOR.MINOR.PATCH`)

## How to Update

When making API changes:
1. **Classify**: MAJOR (breaking), MINOR (backward-compatible addition), or PATCH (fix)
2. **Add entry**: Update `[Unreleased]` section below
3. **Release**: Move to new version section when deploying, bump version accordingly

## [Unreleased]

### Removed
- `Aggregation:LookupClientPetProfile` API documentation and contract (revert PR #62)
- `Offering:ListServices` and `Offering:ListLodgings` API documentation and contract (revert PR #62)

### Added
- `Appointment:Create` API: added `lodgingId` field to `CreateAppointmentRequest.Service` to support lodging unit assignment for `BOARDING`/`DAYCARE` appointments (2026-03-31)
- `Appointment:Create` API: support `BOARDING` and `DAYCARE` service types; `staffIds` is no longer required for boarding/daycare (2026-03-31)
- `Appointment:Create` API: added optional `price` field to `CreateAppointmentRequest.Service`; if not set, system default price is used (2026-04-01)
- `AbandonedBooking:Get` and `AbandonedBooking:List` APIs: added optional `abandonPreference` field on `AbandonedBooking` capturing the preference snapshot submitted during the abandon flow; includes `preferredDay` (`TODAY` | `TOMORROW` | `THIS_WEEK` | `FLEXIBLE`), `preferredTime` (`MORNING` | `AFTERNOON` | `EITHER`), and `preferenceSubmittedAt` timestamp; field is omitted when the customer did not submit a preference (2026-05-12)
- `Webhook` events: added `CUSTOMER_COMPLIANCE_CHANGED` for customer communication compliance or consent changes (2026-06-18)

### Documentation
- `Appointment:Create` API: clarified per-type field requirements for `Service` object; added boarding and daycare examples (2026-03-31)
- `Appointment:Create` API: documented `price` field (2026-04-01)
- `AbandonedBooking`: documented `abandonPreference` field with nested `AbandonPreference` message, `PreferredDay` and `PreferredTime` enums (2026-05-12)
- `Webhook` and `Customer`: documented `CUSTOMER_COMPLIANCE_CHANGED` and its relationship with `complianceConfig` (2026-06-18)

---

## [1.2.0] - 2026-03-20

### Added
- `Staff:List` API now returns deleted staff with `deleted` field to identify deletion status (2026-03-16)
- `Appointment:Get` API now returns `is_deleted` field (2026-03-16)
- Webhook events for Pet: `PET_CREATED`, `PET_UPDATED`, `PET_DELETED` with Pet payload (2026-03-20)

### Changed
- `Appointment:Get` API no longer returns appointment details when appointment is deleted (consistent with `Appointment:List` behavior) (2026-03-16)

---

## [1.1.1] - 2026-02-03

Initial version with changelog tracking.
