# Moego APIs Documentation

Welcome to the official documentation for Moego APIs. This documentation covers all the gRPC services and protocols used in the Moego platform.

## 📚 API Documentation

### Core Business Modules

| Module | Description | Status |
|--------|-------------|--------|
| [Appointment](./appointment.md) | Manage service appointments and scheduling | ✅ Active |
| [Customer](./customer.md) | Customer management and profiles | ✅ Active |
| [Pet](./pet.md) | Pet information and care records | ✅ Active |
| [Business](./business.md) | Business location and settings | ✅ Active |
| [Staff](./staff.md) | Staff management and scheduling | ✅ Active |
| [Order](./order.md) | Order processing and management | ✅ Active |
| [Payment](./payment.md) | Payment processing and transactions | ✅ Active |

### Specialized Services

| Module | Description | Status |
|--------|-------------|--------|
| [Online Booking](./online_booking.md) | Digital appointment booking system | ✅ Active |
| [Membership](./membership.md) | Membership programs and benefits | ✅ Active |
| [Package](./package.md) | Service packages and bundles | ✅ Active |
| [Discount](./discount.md) | Discount codes and promotions | ✅ Active |
| [Review](./review.md) | Customer reviews and ratings | ✅ Active |
| [Report](./report.md) | Business reporting and analytics | ✅ Active |

### Administrative Modules

| Module | Description | Status |
|--------|-------------|--------|
| [Company](./company.md) | Company-level administration | ✅ Active |
| [Setting](./setting_service.md) | System configuration and settings | ✅ Active |
| [Webhook](./webhook.md) | Webhook integrations and notifications | ✅ Active |
| [Event](./event.md) | System events and health checks | ✅ Active |

## 🛠️ Developer Resources

### Protocol Buffers
- [Common Types](./common/)
- [Proto Definitions](../moego/)

### Integration Guides
- [Webhook Integration](./webhook_integration.md)
- [API Authentication](../moego/auth/)
- [Example Implementations](../examples/)

## 📖 Getting Started

1. **Authentication**: Learn about [API Key authentication](../moego/auth/apikey/v1/apikey.proto)
2. **Core Concepts**: Review the [common protobuf types](./common/)
3. **Quick Start**: Check out the [example scripts](../examples/)

## 🤝 Contributing

Documentation improvements are welcome! Please submit pull requests with:
- Clear explanations of changes
- Updated examples where applicable
- Proper formatting and structure

## 📞 Support

For questions or issues with the APIs, please:
- Check the relevant module documentation first
- Review the [examples](../examples/) for implementation guidance
- Contact the development team for specific integration questions

---
*Last updated: February 2026*