// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: credit/credit.proto

package credit

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Credit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Credit) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Credits

	if v, ok := interface{}(m.GetExpireOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreditValidationError{
				field:  "ExpireOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreditValidationError is the validation error returned by Credit.Validate if
// the designated constraints aren't met.
type CreditValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreditValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreditValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreditValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreditValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreditValidationError) ErrorName() string { return "CreditValidationError" }

// Error satisfies the builtin error interface
func (e CreditValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCredit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreditValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreditValidationError{}

// Validate checks the field values on ChargeCreditsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ChargeCreditsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		return ChargeCreditsRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Credits

	return nil
}

// ChargeCreditsRequestValidationError is the validation error returned by
// ChargeCreditsRequest.Validate if the designated constraints aren't met.
type ChargeCreditsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChargeCreditsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChargeCreditsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChargeCreditsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChargeCreditsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChargeCreditsRequestValidationError) ErrorName() string {
	return "ChargeCreditsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChargeCreditsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChargeCreditsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChargeCreditsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChargeCreditsRequestValidationError{}
