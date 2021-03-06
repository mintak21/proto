// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sample/proto/pancake.proto

package golang

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
)

// Validate checks the field values on Pancake with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Pancake) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BakerName

	// no validation rules for Menu

	// no validation rules for TechinicalScore

	if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PancakeValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PancakeValidationError is the validation error returned by Pancake.Validate
// if the designated constraints aren't met.
type PancakeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PancakeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PancakeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PancakeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PancakeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PancakeValidationError) ErrorName() string { return "PancakeValidationError" }

// Error satisfies the builtin error interface
func (e PancakeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPancake.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PancakeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PancakeValidationError{}

// Validate checks the field values on Report with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Report) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetBakeCounts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReportValidationError{
					field:  fmt.Sprintf("BakeCounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ReportValidationError is the validation error returned by Report.Validate if
// the designated constraints aren't met.
type ReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReportValidationError) ErrorName() string { return "ReportValidationError" }

// Error satisfies the builtin error interface
func (e ReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReportValidationError{}

// Validate checks the field values on BakeRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BakeRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Menu

	return nil
}

// BakeRequestValidationError is the validation error returned by
// BakeRequest.Validate if the designated constraints aren't met.
type BakeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BakeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BakeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BakeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BakeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BakeRequestValidationError) ErrorName() string { return "BakeRequestValidationError" }

// Error satisfies the builtin error interface
func (e BakeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBakeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BakeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BakeRequestValidationError{}

// Validate checks the field values on BakeResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BakeResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPancake()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BakeResponseValidationError{
				field:  "Pancake",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BakeResponseValidationError is the validation error returned by
// BakeResponse.Validate if the designated constraints aren't met.
type BakeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BakeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BakeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BakeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BakeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BakeResponseValidationError) ErrorName() string { return "BakeResponseValidationError" }

// Error satisfies the builtin error interface
func (e BakeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBakeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BakeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BakeResponseValidationError{}

// Validate checks the field values on ReportRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReportRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ReportRequestValidationError is the validation error returned by
// ReportRequest.Validate if the designated constraints aren't met.
type ReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReportRequestValidationError) ErrorName() string { return "ReportRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReportRequestValidationError{}

// Validate checks the field values on ReportResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReportResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetReport()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReportResponseValidationError{
				field:  "Report",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReportResponseValidationError is the validation error returned by
// ReportResponse.Validate if the designated constraints aren't met.
type ReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReportResponseValidationError) ErrorName() string { return "ReportResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReportResponseValidationError{}

// Validate checks the field values on Report_BakeCount with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Report_BakeCount) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Menu

	// no validation rules for Count

	return nil
}

// Report_BakeCountValidationError is the validation error returned by
// Report_BakeCount.Validate if the designated constraints aren't met.
type Report_BakeCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Report_BakeCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Report_BakeCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Report_BakeCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Report_BakeCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Report_BakeCountValidationError) ErrorName() string { return "Report_BakeCountValidationError" }

// Error satisfies the builtin error interface
func (e Report_BakeCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReport_BakeCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Report_BakeCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Report_BakeCountValidationError{}
