// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: socket.proto

package socket

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

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *Request_PlayerMove:
		if v == nil {
			err := RequestValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetPlayerMove() == nil {
			err := RequestValidationError{
				field:  "PlayerMove",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPlayerMove()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestValidationError{
						field:  "PlayerMove",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestValidationError{
						field:  "PlayerMove",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlayerMove()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestValidationError{
					field:  "PlayerMove",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := RequestValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResponseMultiError, or nil
// if none found.
func (m *Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *Response_Bulk_:
		if v == nil {
			err := ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetBulk() == nil {
			err := ResponseValidationError{
				field:  "Bulk",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBulk()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Bulk",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Bulk",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBulk()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "Bulk",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Response_Error_:
		if v == nil {
			err := ResponseValidationError{
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
			err := ResponseValidationError{
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
					errors = append(errors, ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Response_Info_:
		if v == nil {
			err := ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetInfo() == nil {
			err := ResponseValidationError{
				field:  "Info",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Info",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Info",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Response_PlayerJoin_:
		if v == nil {
			err := ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetPlayerJoin() == nil {
			err := ResponseValidationError{
				field:  "PlayerJoin",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPlayerJoin()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "PlayerJoin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "PlayerJoin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlayerJoin()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "PlayerJoin",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Response_PlayerLeave:
		if v == nil {
			err := ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetPlayerLeave() <= 0 {
			err := ResponseValidationError{
				field:  "PlayerLeave",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *Response_PlayerMove_:
		if v == nil {
			err := ResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if m.GetPlayerMove() == nil {
			err := ResponseValidationError{
				field:  "PlayerMove",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPlayerMove()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "PlayerMove",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "PlayerMove",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlayerMove()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "PlayerMove",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := ResponseValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ResponseMultiError(errors)
	}

	return nil
}

// ResponseMultiError is an error wrapping multiple validation errors returned
// by Response.ValidateAll() if the designated constraints aren't met.
type ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResponseMultiError) AllErrors() []error { return m }

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on Pos with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Pos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pos with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PosMultiError, or nil if none found.
func (m *Pos) ValidateAll() error {
	return m.validate(true)
}

func (m *Pos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	if len(errors) > 0 {
		return PosMultiError(errors)
	}

	return nil
}

// PosMultiError is an error wrapping multiple validation errors returned by
// Pos.ValidateAll() if the designated constraints aren't met.
type PosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PosMultiError) AllErrors() []error { return m }

// PosValidationError is the validation error returned by Pos.Validate if the
// designated constraints aren't met.
type PosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PosValidationError) ErrorName() string { return "PosValidationError" }

// Error satisfies the builtin error interface
func (e PosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PosValidationError{}

// Validate checks the field values on Response_Bulk with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response_Bulk) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response_Bulk with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Response_BulkMultiError, or
// nil if none found.
func (m *Response_Bulk) ValidateAll() error {
	return m.validate(true)
}

func (m *Response_Bulk) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResponses()) < 1 {
		err := Response_BulkValidationError{
			field:  "Responses",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetResponses() {
		_, _ = idx, item

		if item == nil {
			err := Response_BulkValidationError{
				field:  fmt.Sprintf("Responses[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Response_BulkValidationError{
						field:  fmt.Sprintf("Responses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Response_BulkValidationError{
						field:  fmt.Sprintf("Responses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Response_BulkValidationError{
					field:  fmt.Sprintf("Responses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return Response_BulkMultiError(errors)
	}

	return nil
}

// Response_BulkMultiError is an error wrapping multiple validation errors
// returned by Response_Bulk.ValidateAll() if the designated constraints
// aren't met.
type Response_BulkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Response_BulkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Response_BulkMultiError) AllErrors() []error { return m }

// Response_BulkValidationError is the validation error returned by
// Response_Bulk.Validate if the designated constraints aren't met.
type Response_BulkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Response_BulkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Response_BulkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Response_BulkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Response_BulkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Response_BulkValidationError) ErrorName() string { return "Response_BulkValidationError" }

// Error satisfies the builtin error interface
func (e Response_BulkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse_Bulk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Response_BulkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Response_BulkValidationError{}

// Validate checks the field values on Response_Error with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response_Error) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response_Error with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Response_ErrorMultiError,
// or nil if none found.
func (m *Response_Error) ValidateAll() error {
	return m.validate(true)
}

func (m *Response_Error) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := Response_Error_Code_name[int32(m.GetCode())]; !ok {
		err := Response_ErrorValidationError{
			field:  "Code",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMessage()) < 1 {
		err := Response_ErrorValidationError{
			field:  "Message",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Response_ErrorMultiError(errors)
	}

	return nil
}

// Response_ErrorMultiError is an error wrapping multiple validation errors
// returned by Response_Error.ValidateAll() if the designated constraints
// aren't met.
type Response_ErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Response_ErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Response_ErrorMultiError) AllErrors() []error { return m }

// Response_ErrorValidationError is the validation error returned by
// Response_Error.Validate if the designated constraints aren't met.
type Response_ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Response_ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Response_ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Response_ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Response_ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Response_ErrorValidationError) ErrorName() string { return "Response_ErrorValidationError" }

// Error satisfies the builtin error interface
func (e Response_ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse_Error.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Response_ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Response_ErrorValidationError{}

// Validate checks the field values on Response_Info with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response_Info) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response_Info with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Response_InfoMultiError, or
// nil if none found.
func (m *Response_Info) ValidateAll() error {
	return m.validate(true)
}

func (m *Response_Info) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPlayerId() <= 0 {
		err := Response_InfoValidationError{
			field:  "PlayerId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPlayers() {
		_, _ = idx, item

		if item == nil {
			err := Response_InfoValidationError{
				field:  fmt.Sprintf("Players[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Response_InfoValidationError{
						field:  fmt.Sprintf("Players[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Response_InfoValidationError{
						field:  fmt.Sprintf("Players[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Response_InfoValidationError{
					field:  fmt.Sprintf("Players[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return Response_InfoMultiError(errors)
	}

	return nil
}

// Response_InfoMultiError is an error wrapping multiple validation errors
// returned by Response_Info.ValidateAll() if the designated constraints
// aren't met.
type Response_InfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Response_InfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Response_InfoMultiError) AllErrors() []error { return m }

// Response_InfoValidationError is the validation error returned by
// Response_Info.Validate if the designated constraints aren't met.
type Response_InfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Response_InfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Response_InfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Response_InfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Response_InfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Response_InfoValidationError) ErrorName() string { return "Response_InfoValidationError" }

// Error satisfies the builtin error interface
func (e Response_InfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse_Info.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Response_InfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Response_InfoValidationError{}

// Validate checks the field values on Response_PlayerJoin with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Response_PlayerJoin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response_PlayerJoin with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Response_PlayerJoinMultiError, or nil if none found.
func (m *Response_PlayerJoin) ValidateAll() error {
	return m.validate(true)
}

func (m *Response_PlayerJoin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := Response_PlayerJoinValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUsername()); l < 1 || l > 32 {
		err := Response_PlayerJoinValidationError{
			field:  "Username",
			reason: "value length must be between 1 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Response_PlayerJoinMultiError(errors)
	}

	return nil
}

// Response_PlayerJoinMultiError is an error wrapping multiple validation
// errors returned by Response_PlayerJoin.ValidateAll() if the designated
// constraints aren't met.
type Response_PlayerJoinMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Response_PlayerJoinMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Response_PlayerJoinMultiError) AllErrors() []error { return m }

// Response_PlayerJoinValidationError is the validation error returned by
// Response_PlayerJoin.Validate if the designated constraints aren't met.
type Response_PlayerJoinValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Response_PlayerJoinValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Response_PlayerJoinValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Response_PlayerJoinValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Response_PlayerJoinValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Response_PlayerJoinValidationError) ErrorName() string {
	return "Response_PlayerJoinValidationError"
}

// Error satisfies the builtin error interface
func (e Response_PlayerJoinValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse_PlayerJoin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Response_PlayerJoinValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Response_PlayerJoinValidationError{}

// Validate checks the field values on Response_PlayerMove with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Response_PlayerMove) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response_PlayerMove with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Response_PlayerMoveMultiError, or nil if none found.
func (m *Response_PlayerMove) ValidateAll() error {
	return m.validate(true)
}

func (m *Response_PlayerMove) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPlayerId() <= 0 {
		err := Response_PlayerMoveValidationError{
			field:  "PlayerId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPos() == nil {
		err := Response_PlayerMoveValidationError{
			field:  "Pos",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPos()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Response_PlayerMoveValidationError{
					field:  "Pos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Response_PlayerMoveValidationError{
					field:  "Pos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPos()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Response_PlayerMoveValidationError{
				field:  "Pos",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Response_PlayerMoveMultiError(errors)
	}

	return nil
}

// Response_PlayerMoveMultiError is an error wrapping multiple validation
// errors returned by Response_PlayerMove.ValidateAll() if the designated
// constraints aren't met.
type Response_PlayerMoveMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Response_PlayerMoveMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Response_PlayerMoveMultiError) AllErrors() []error { return m }

// Response_PlayerMoveValidationError is the validation error returned by
// Response_PlayerMove.Validate if the designated constraints aren't met.
type Response_PlayerMoveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Response_PlayerMoveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Response_PlayerMoveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Response_PlayerMoveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Response_PlayerMoveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Response_PlayerMoveValidationError) ErrorName() string {
	return "Response_PlayerMoveValidationError"
}

// Error satisfies the builtin error interface
func (e Response_PlayerMoveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse_PlayerMove.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Response_PlayerMoveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Response_PlayerMoveValidationError{}
