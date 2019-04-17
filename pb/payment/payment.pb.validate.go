// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: payment/payment.proto

package payment

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

// Validate checks the field values on SubscribeMembershipRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubscribeMembershipRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		return SubscribeMembershipRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for GatewayId

	if utf8.RuneCountInString(m.GetMembershipId()) < 1 {
		return SubscribeMembershipRequestValidationError{
			field:  "MembershipId",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for IsRecurring

	if utf8.RuneCountInString(m.GetCardId()) < 1 {
		return SubscribeMembershipRequestValidationError{
			field:  "CardId",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// SubscribeMembershipRequestValidationError is the validation error returned
// by SubscribeMembershipRequest.Validate if the designated constraints aren't met.
type SubscribeMembershipRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeMembershipRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeMembershipRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeMembershipRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeMembershipRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeMembershipRequestValidationError) ErrorName() string {
	return "SubscribeMembershipRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeMembershipRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeMembershipRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeMembershipRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeMembershipRequestValidationError{}

// Validate checks the field values on Subscription with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Subscription) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for GatewayId

	// no validation rules for GatewayCustomerId

	// no validation rules for GatewaySubscriptionId

	// no validation rules for MembershipId

	// no validation rules for IsRecurring

	// no validation rules for CardId

	if v, ok := interface{}(m.GetValidTill()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "ValidTill",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSubscriptionStartDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "SubscriptionStartDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCurrentPeriodStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "CurrentPeriodStart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCurrentPeriodEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "CurrentPeriodEnd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SubscriptionValidationError is the validation error returned by
// Subscription.Validate if the designated constraints aren't met.
type SubscriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscriptionValidationError) ErrorName() string { return "SubscriptionValidationError" }

// Error satisfies the builtin error interface
func (e SubscriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscriptionValidationError{}

// Validate checks the field values on CancelMembershipRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CancelMembershipRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		return CancelMembershipRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetSubscriptionId()) < 1 {
		return CancelMembershipRequestValidationError{
			field:  "SubscriptionId",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// CancelMembershipRequestValidationError is the validation error returned by
// CancelMembershipRequest.Validate if the designated constraints aren't met.
type CancelMembershipRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelMembershipRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelMembershipRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelMembershipRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelMembershipRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelMembershipRequestValidationError) ErrorName() string {
	return "CancelMembershipRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CancelMembershipRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelMembershipRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelMembershipRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelMembershipRequestValidationError{}