// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: moego/business/block/v1/block_service.proto

package blockpb

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

// Validate checks the field values on CreateBlockRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateBlockRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBlockRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBlockRequestMultiError, or nil if none found.
func (m *CreateBlockRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBlockRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BusinessId

	// no validation rules for StaffId

	if all {
		switch v := interface{}(m.GetDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateBlockRequestValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateBlockRequestValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateBlockRequestValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	// no validation rules for ColorCode

	if len(errors) > 0 {
		return CreateBlockRequestMultiError(errors)
	}

	return nil
}

// CreateBlockRequestMultiError is an error wrapping multiple validation errors
// returned by CreateBlockRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateBlockRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBlockRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBlockRequestMultiError) AllErrors() []error { return m }

// CreateBlockRequestValidationError is the validation error returned by
// CreateBlockRequest.Validate if the designated constraints aren't met.
type CreateBlockRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBlockRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBlockRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBlockRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBlockRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBlockRequestValidationError) ErrorName() string {
	return "CreateBlockRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBlockRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBlockRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBlockRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBlockRequestValidationError{}

// Validate checks the field values on GetBlockRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetBlockRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBlockRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetBlockRequestMultiError, or nil if none found.
func (m *GetBlockRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBlockRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetBlockRequestMultiError(errors)
	}

	return nil
}

// GetBlockRequestMultiError is an error wrapping multiple validation errors
// returned by GetBlockRequest.ValidateAll() if the designated constraints
// aren't met.
type GetBlockRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBlockRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBlockRequestMultiError) AllErrors() []error { return m }

// GetBlockRequestValidationError is the validation error returned by
// GetBlockRequest.Validate if the designated constraints aren't met.
type GetBlockRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBlockRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBlockRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBlockRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBlockRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBlockRequestValidationError) ErrorName() string { return "GetBlockRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetBlockRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBlockRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBlockRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBlockRequestValidationError{}
