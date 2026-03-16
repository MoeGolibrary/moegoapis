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

### Added
- `Staff:List` API now returns deleted staff with `deleted` field to identify deletion status (2026-03-16)
- `Appointment:Get` API now returns `is_deleted` field (2026-03-16)

### Changed
- `Appointment:Get` API no longer returns appointment details when appointment is deleted (consistent with `Appointment:List` behavior) (2026-03-16)

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

---

## [1.1.1] - 2026-02-03

Initial version with changelog tracking.
