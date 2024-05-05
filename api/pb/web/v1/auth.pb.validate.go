// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: web/v1/auth.proto

package web

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

// Validate checks the field values on AuthLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthLoginRequestMultiError, or nil if none found.
func (m *AuthLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mobile

	// no validation rules for Password

	// no validation rules for Platform

	if len(errors) > 0 {
		return AuthLoginRequestMultiError(errors)
	}

	return nil
}

// AuthLoginRequestMultiError is an error wrapping multiple validation errors
// returned by AuthLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginRequestMultiError) AllErrors() []error { return m }

// AuthLoginRequestValidationError is the validation error returned by
// AuthLoginRequest.Validate if the designated constraints aren't met.
type AuthLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginRequestValidationError) ErrorName() string { return "AuthLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginRequestValidationError{}

// Validate checks the field values on AuthLoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthLoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthLoginResponseMultiError, or nil if none found.
func (m *AuthLoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for AccessToken

	// no validation rules for ExpiresIn

	if len(errors) > 0 {
		return AuthLoginResponseMultiError(errors)
	}

	return nil
}

// AuthLoginResponseMultiError is an error wrapping multiple validation errors
// returned by AuthLoginResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginResponseMultiError) AllErrors() []error { return m }

// AuthLoginResponseValidationError is the validation error returned by
// AuthLoginResponse.Validate if the designated constraints aren't met.
type AuthLoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginResponseValidationError) ErrorName() string {
	return "AuthLoginResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthLoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginResponseValidationError{}

// Validate checks the field values on AuthRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthRegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthRegisterRequestMultiError, or nil if none found.
func (m *AuthRegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Nickname

	// no validation rules for Mobile

	// no validation rules for Password

	// no validation rules for Platform

	// no validation rules for SmsCode

	if len(errors) > 0 {
		return AuthRegisterRequestMultiError(errors)
	}

	return nil
}

// AuthRegisterRequestMultiError is an error wrapping multiple validation
// errors returned by AuthRegisterRequest.ValidateAll() if the designated
// constraints aren't met.
type AuthRegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRegisterRequestMultiError) AllErrors() []error { return m }

// AuthRegisterRequestValidationError is the validation error returned by
// AuthRegisterRequest.Validate if the designated constraints aren't met.
type AuthRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRegisterRequestValidationError) ErrorName() string {
	return "AuthRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRegisterRequestValidationError{}

// Validate checks the field values on AuthRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthRegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthRegisterResponseMultiError, or nil if none found.
func (m *AuthRegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthRegisterResponseMultiError(errors)
	}

	return nil
}

// AuthRegisterResponseMultiError is an error wrapping multiple validation
// errors returned by AuthRegisterResponse.ValidateAll() if the designated
// constraints aren't met.
type AuthRegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRegisterResponseMultiError) AllErrors() []error { return m }

// AuthRegisterResponseValidationError is the validation error returned by
// AuthRegisterResponse.Validate if the designated constraints aren't met.
type AuthRegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRegisterResponseValidationError) ErrorName() string {
	return "AuthRegisterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthRegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRegisterResponseValidationError{}

// Validate checks the field values on AuthRefreshRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthRefreshRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRefreshRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthRefreshRequestMultiError, or nil if none found.
func (m *AuthRefreshRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRefreshRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthRefreshRequestMultiError(errors)
	}

	return nil
}

// AuthRefreshRequestMultiError is an error wrapping multiple validation errors
// returned by AuthRefreshRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthRefreshRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRefreshRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRefreshRequestMultiError) AllErrors() []error { return m }

// AuthRefreshRequestValidationError is the validation error returned by
// AuthRefreshRequest.Validate if the designated constraints aren't met.
type AuthRefreshRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRefreshRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRefreshRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRefreshRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRefreshRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRefreshRequestValidationError) ErrorName() string {
	return "AuthRefreshRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthRefreshRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRefreshRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRefreshRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRefreshRequestValidationError{}

// Validate checks the field values on AuthRefreshResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthRefreshResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRefreshResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthRefreshResponseMultiError, or nil if none found.
func (m *AuthRefreshResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRefreshResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for AccessToken

	// no validation rules for ExpiresIn

	if len(errors) > 0 {
		return AuthRefreshResponseMultiError(errors)
	}

	return nil
}

// AuthRefreshResponseMultiError is an error wrapping multiple validation
// errors returned by AuthRefreshResponse.ValidateAll() if the designated
// constraints aren't met.
type AuthRefreshResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRefreshResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRefreshResponseMultiError) AllErrors() []error { return m }

// AuthRefreshResponseValidationError is the validation error returned by
// AuthRefreshResponse.Validate if the designated constraints aren't met.
type AuthRefreshResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRefreshResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRefreshResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRefreshResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRefreshResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRefreshResponseValidationError) ErrorName() string {
	return "AuthRefreshResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthRefreshResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRefreshResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRefreshResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRefreshResponseValidationError{}

// Validate checks the field values on AuthForgetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthForgetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthForgetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthForgetRequestMultiError, or nil if none found.
func (m *AuthForgetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthForgetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mobile

	// no validation rules for Password

	// no validation rules for SmsCode

	if len(errors) > 0 {
		return AuthForgetRequestMultiError(errors)
	}

	return nil
}

// AuthForgetRequestMultiError is an error wrapping multiple validation errors
// returned by AuthForgetRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthForgetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthForgetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthForgetRequestMultiError) AllErrors() []error { return m }

// AuthForgetRequestValidationError is the validation error returned by
// AuthForgetRequest.Validate if the designated constraints aren't met.
type AuthForgetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthForgetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthForgetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthForgetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthForgetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthForgetRequestValidationError) ErrorName() string {
	return "AuthForgetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthForgetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthForgetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthForgetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthForgetRequestValidationError{}

// Validate checks the field values on AuthForgetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthForgetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthForgetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthForgetResponseMultiError, or nil if none found.
func (m *AuthForgetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthForgetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthForgetResponseMultiError(errors)
	}

	return nil
}

// AuthForgetResponseMultiError is an error wrapping multiple validation errors
// returned by AuthForgetResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthForgetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthForgetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthForgetResponseMultiError) AllErrors() []error { return m }

// AuthForgetResponseValidationError is the validation error returned by
// AuthForgetResponse.Validate if the designated constraints aren't met.
type AuthForgetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthForgetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthForgetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthForgetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthForgetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthForgetResponseValidationError) ErrorName() string {
	return "AuthForgetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthForgetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthForgetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthForgetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthForgetResponseValidationError{}