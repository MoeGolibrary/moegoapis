// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: moego/business/appointment/v1/appointment.proto

package appointmentpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	paymentpb "github.com/MoeGolibrary/moegoapis/genproto/go/business/payment/v1/paymentpb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = paymentpb.Payment_Status(0)
)

// Validate checks the field values on Appointment with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Appointment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Appointment with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppointmentMultiError, or
// nil if none found.
func (m *Appointment) ValidateAll() error {
	return m.validate(true)
}

func (m *Appointment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for BusinessId

	// no validation rules for CustomerId

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPetServiceDetails() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AppointmentValidationError{
						field:  fmt.Sprintf("PetServiceDetails[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AppointmentValidationError{
						field:  fmt.Sprintf("PetServiceDetails[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AppointmentValidationError{
					field:  fmt.Sprintf("PetServiceDetails[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Status

	// no validation rules for TicketComment

	// no validation rules for ColorCode

	// no validation rules for InvoiceId

	if all {
		switch v := interface{}(m.GetTotalAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "TotalAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "TotalAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTotalAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "TotalAmount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPaidAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "PaidAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "PaidAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaidAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "PaidAmount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRefundAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "RefundAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "RefundAmount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefundAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "RefundAmount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PaymentStatus

	// no validation rules for CreatedBy

	if all {
		switch v := interface{}(m.GetCreatedTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "CreatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "CreatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "CreatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LastUpdatedBy

	if all {
		switch v := interface{}(m.GetLastUpdatedTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "LastUpdatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppointmentValidationError{
					field:  "LastUpdatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastUpdatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppointmentValidationError{
				field:  "LastUpdatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AppointmentMultiError(errors)
	}

	return nil
}

// AppointmentMultiError is an error wrapping multiple validation errors
// returned by Appointment.ValidateAll() if the designated constraints aren't met.
type AppointmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppointmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppointmentMultiError) AllErrors() []error { return m }

// AppointmentValidationError is the validation error returned by
// Appointment.Validate if the designated constraints aren't met.
type AppointmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppointmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppointmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppointmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppointmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppointmentValidationError) ErrorName() string { return "AppointmentValidationError" }

// Error satisfies the builtin error interface
func (e AppointmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppointment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppointmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppointmentValidationError{}
