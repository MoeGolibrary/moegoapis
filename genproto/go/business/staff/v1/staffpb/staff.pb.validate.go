// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: moego/business/staff/v1/staff.proto

package staffpb

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
)

// Validate checks the field values on Staff with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Staff) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Staff with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StaffMultiError, or nil if none found.
func (m *Staff) ValidateAll() error {
	return m.validate(true)
}

func (m *Staff) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Avatar

	// no validation rules for Phone

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetRole()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StaffValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StaffValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StaffValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHiredTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StaffValidationError{
					field:  "HiredTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StaffValidationError{
					field:  "HiredTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHiredTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StaffValidationError{
				field:  "HiredTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MobileVanIds

	if len(errors) > 0 {
		return StaffMultiError(errors)
	}

	return nil
}

// StaffMultiError is an error wrapping multiple validation errors returned by
// Staff.ValidateAll() if the designated constraints aren't met.
type StaffMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StaffMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StaffMultiError) AllErrors() []error { return m }

// StaffValidationError is the validation error returned by Staff.Validate if
// the designated constraints aren't met.
type StaffValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StaffValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StaffValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StaffValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StaffValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StaffValidationError) ErrorName() string { return "StaffValidationError" }

// Error satisfies the builtin error interface
func (e StaffValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStaff.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StaffValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StaffValidationError{}

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Role) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Role with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RoleMultiError, or nil if none found.
func (m *Role) ValidateAll() error {
	return m.validate(true)
}

func (m *Role) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return RoleMultiError(errors)
	}

	return nil
}

// RoleMultiError is an error wrapping multiple validation errors returned by
// Role.ValidateAll() if the designated constraints aren't met.
type RoleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleMultiError) AllErrors() []error { return m }

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}
