// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: moego/business/business/v1/business.proto

package businesspb

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

// Validate checks the field values on Business with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Business) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Business with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BusinessMultiError, or nil
// if none found.
func (m *Business) ValidateAll() error {
	return m.validate(true)
}

func (m *Business) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Avatar

	// no validation rules for Phone

	// no validation rules for BusinessPhone

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Website

	// no validation rules for Facebook

	// no validation rules for Instagram

	// no validation rules for Yelp

	// no validation rules for Google

	if len(errors) > 0 {
		return BusinessMultiError(errors)
	}

	return nil
}

// BusinessMultiError is an error wrapping multiple validation errors returned
// by Business.ValidateAll() if the designated constraints aren't met.
type BusinessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BusinessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BusinessMultiError) AllErrors() []error { return m }

// BusinessValidationError is the validation error returned by
// Business.Validate if the designated constraints aren't met.
type BusinessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BusinessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BusinessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BusinessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BusinessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BusinessValidationError) ErrorName() string { return "BusinessValidationError" }

// Error satisfies the builtin error interface
func (e BusinessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBusiness.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BusinessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BusinessValidationError{}