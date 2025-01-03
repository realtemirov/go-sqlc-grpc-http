// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: general/request.proto

package general

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

// Validate checks the field values on GetAllRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetAllRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAllRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetAllRequestMultiError, or
// nil if none found.
func (m *GetAllRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAllRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPageSize() <= 0 {
		err := GetAllRequestValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() <= 0 {
		err := GetAllRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Search

	// no validation rules for UserId

	// no validation rules for Limit

	// no validation rules for Offset

	// no validation rules for Lang

	if m.GetCountryId() <= 0 {
		err := GetAllRequestValidationError{
			field:  "CountryId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetAllRequestMultiError(errors)
	}

	return nil
}

// GetAllRequestMultiError is an error wrapping multiple validation errors
// returned by GetAllRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAllRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAllRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAllRequestMultiError) AllErrors() []error { return m }

// GetAllRequestValidationError is the validation error returned by
// GetAllRequest.Validate if the designated constraints aren't met.
type GetAllRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllRequestValidationError) ErrorName() string { return "GetAllRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetAllRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllRequestValidationError{}
