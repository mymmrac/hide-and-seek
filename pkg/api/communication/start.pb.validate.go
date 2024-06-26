// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: start.proto

package communication

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

// Validate checks the field values on Start with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Start) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Start with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StartMultiError, or nil if none found.
func (m *Start) ValidateAll() error {
	return m.validate(true)
}

func (m *Start) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StartMultiError(errors)
	}

	return nil
}

// StartMultiError is an error wrapping multiple validation errors returned by
// Start.ValidateAll() if the designated constraints aren't met.
type StartMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartMultiError) AllErrors() []error { return m }

// StartValidationError is the validation error returned by Start.Validate if
// the designated constraints aren't met.
type StartValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartValidationError) ErrorName() string { return "StartValidationError" }

// Error satisfies the builtin error interface
func (e StartValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStart.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartValidationError{}

// Validate checks the field values on Start_Request with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Start_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Start_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Start_RequestMultiError, or
// nil if none found.
func (m *Start_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Start_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 1 || l > 32 {
		err := Start_RequestValidationError{
			field:  "Username",
			reason: "value length must be between 1 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Start_Request_Username_Pattern.MatchString(m.GetUsername()) {
		err := Start_RequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"^[a-zA-Z][a-zA-Z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Start_RequestMultiError(errors)
	}

	return nil
}

// Start_RequestMultiError is an error wrapping multiple validation errors
// returned by Start_Request.ValidateAll() if the designated constraints
// aren't met.
type Start_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Start_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Start_RequestMultiError) AllErrors() []error { return m }

// Start_RequestValidationError is the validation error returned by
// Start_Request.Validate if the designated constraints aren't met.
type Start_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Start_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Start_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Start_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Start_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Start_RequestValidationError) ErrorName() string { return "Start_RequestValidationError" }

// Error satisfies the builtin error interface
func (e Start_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStart_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Start_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Start_RequestValidationError{}

var _Start_Request_Username_Pattern = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$")

// Validate checks the field values on Start_Response with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Start_Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Start_Response with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Start_ResponseMultiError,
// or nil if none found.
func (m *Start_Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Start_Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *Start_Response_Error:
		if v == nil {
			err := Start_ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetError() == nil {
			err := Start_ResponseValidationError{
				field:  "Error",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetError()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Start_ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Start_ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Start_ResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Start_Response_Result_:
		if v == nil {
			err := Start_ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetResult() == nil {
			err := Start_ResponseValidationError{
				field:  "Result",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetResult()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Start_ResponseValidationError{
						field:  "Result",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Start_ResponseValidationError{
						field:  "Result",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Start_ResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := Start_ResponseValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Start_ResponseMultiError(errors)
	}

	return nil
}

// Start_ResponseMultiError is an error wrapping multiple validation errors
// returned by Start_Response.ValidateAll() if the designated constraints
// aren't met.
type Start_ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Start_ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Start_ResponseMultiError) AllErrors() []error { return m }

// Start_ResponseValidationError is the validation error returned by
// Start_Response.Validate if the designated constraints aren't met.
type Start_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Start_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Start_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Start_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Start_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Start_ResponseValidationError) ErrorName() string { return "Start_ResponseValidationError" }

// Error satisfies the builtin error interface
func (e Start_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStart_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Start_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Start_ResponseValidationError{}

// Validate checks the field values on Start_Response_Result with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Start_Response_Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Start_Response_Result with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Start_Response_ResultMultiError, or nil if none found.
func (m *Start_Response_Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Start_Response_Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetToken()) < 1 {
		err := Start_Response_ResultValidationError{
			field:  "Token",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Start_Response_ResultMultiError(errors)
	}

	return nil
}

// Start_Response_ResultMultiError is an error wrapping multiple validation
// errors returned by Start_Response_Result.ValidateAll() if the designated
// constraints aren't met.
type Start_Response_ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Start_Response_ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Start_Response_ResultMultiError) AllErrors() []error { return m }

// Start_Response_ResultValidationError is the validation error returned by
// Start_Response_Result.Validate if the designated constraints aren't met.
type Start_Response_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Start_Response_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Start_Response_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Start_Response_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Start_Response_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Start_Response_ResultValidationError) ErrorName() string {
	return "Start_Response_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e Start_Response_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStart_Response_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Start_Response_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Start_Response_ResultValidationError{}
