// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rill/runtime/v1/catalog.proto

package runtimev1

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

// Validate checks the field values on Table with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Table) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Table with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TableMultiError, or nil if none found.
func (m *Table) ValidateAll() error {
	return m.validate(true)
}

func (m *Table) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetSchema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TableValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TableValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TableValidationError{
				field:  "Schema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Managed

	if len(errors) > 0 {
		return TableMultiError(errors)
	}

	return nil
}

// TableMultiError is an error wrapping multiple validation errors returned by
// Table.ValidateAll() if the designated constraints aren't met.
type TableMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TableMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TableMultiError) AllErrors() []error { return m }

// TableValidationError is the validation error returned by Table.Validate if
// the designated constraints aren't met.
type TableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TableValidationError) ErrorName() string { return "TableValidationError" }

// Error satisfies the builtin error interface
func (e TableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTable.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TableValidationError{}

// Validate checks the field values on Source with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Source) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Source with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SourceMultiError, or nil if none found.
func (m *Source) ValidateAll() error {
	return m.validate(true)
}

func (m *Source) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Connector

	if all {
		switch v := interface{}(m.GetProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SourceValidationError{
				field:  "Properties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSchema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SourceValidationError{
				field:  "Schema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolicy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SourceValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SourceValidationError{
				field:  "Policy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Timeout

	if len(errors) > 0 {
		return SourceMultiError(errors)
	}

	return nil
}

// SourceMultiError is an error wrapping multiple validation errors returned by
// Source.ValidateAll() if the designated constraints aren't met.
type SourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceMultiError) AllErrors() []error { return m }

// SourceValidationError is the validation error returned by Source.Validate if
// the designated constraints aren't met.
type SourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceValidationError) ErrorName() string { return "SourceValidationError" }

// Error satisfies the builtin error interface
func (e SourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceValidationError{}

// Validate checks the field values on Model with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Model) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Model with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ModelMultiError, or nil if none found.
func (m *Model) ValidateAll() error {
	return m.validate(true)
}

func (m *Model) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Sql

	// no validation rules for Dialect

	if all {
		switch v := interface{}(m.GetSchema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ModelValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ModelValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ModelValidationError{
				field:  "Schema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Materialize

	if len(errors) > 0 {
		return ModelMultiError(errors)
	}

	return nil
}

// ModelMultiError is an error wrapping multiple validation errors returned by
// Model.ValidateAll() if the designated constraints aren't met.
type ModelMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ModelMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ModelMultiError) AllErrors() []error { return m }

// ModelValidationError is the validation error returned by Model.Validate if
// the designated constraints aren't met.
type ModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModelValidationError) ErrorName() string { return "ModelValidationError" }

// Error satisfies the builtin error interface
func (e ModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModelValidationError{}

// Validate checks the field values on MetricsView with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetricsView) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsView with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetricsViewMultiError, or
// nil if none found.
func (m *MetricsView) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsView) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Model

	// no validation rules for TimeDimension

	for idx, item := range m.GetDimensions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetricsViewValidationError{
						field:  fmt.Sprintf("Dimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetricsViewValidationError{
						field:  fmt.Sprintf("Dimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetricsViewValidationError{
					field:  fmt.Sprintf("Dimensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetMeasures() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetricsViewValidationError{
						field:  fmt.Sprintf("Measures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetricsViewValidationError{
						field:  fmt.Sprintf("Measures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetricsViewValidationError{
					field:  fmt.Sprintf("Measures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Label

	// no validation rules for Description

	if len(errors) > 0 {
		return MetricsViewMultiError(errors)
	}

	return nil
}

// MetricsViewMultiError is an error wrapping multiple validation errors
// returned by MetricsView.ValidateAll() if the designated constraints aren't met.
type MetricsViewMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsViewMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsViewMultiError) AllErrors() []error { return m }

// MetricsViewValidationError is the validation error returned by
// MetricsView.Validate if the designated constraints aren't met.
type MetricsViewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsViewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsViewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsViewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsViewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsViewValidationError) ErrorName() string { return "MetricsViewValidationError" }

// Error satisfies the builtin error interface
func (e MetricsViewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsView.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsViewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsViewValidationError{}

// Validate checks the field values on Source_ExtractPolicy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Source_ExtractPolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Source_ExtractPolicy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Source_ExtractPolicyMultiError, or nil if none found.
func (m *Source_ExtractPolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *Source_ExtractPolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RowsStrategy

	// no validation rules for RowsLimitBytes

	// no validation rules for FilesStrategy

	// no validation rules for FilesLimit

	if len(errors) > 0 {
		return Source_ExtractPolicyMultiError(errors)
	}

	return nil
}

// Source_ExtractPolicyMultiError is an error wrapping multiple validation
// errors returned by Source_ExtractPolicy.ValidateAll() if the designated
// constraints aren't met.
type Source_ExtractPolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Source_ExtractPolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Source_ExtractPolicyMultiError) AllErrors() []error { return m }

// Source_ExtractPolicyValidationError is the validation error returned by
// Source_ExtractPolicy.Validate if the designated constraints aren't met.
type Source_ExtractPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Source_ExtractPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Source_ExtractPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Source_ExtractPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Source_ExtractPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Source_ExtractPolicyValidationError) ErrorName() string {
	return "Source_ExtractPolicyValidationError"
}

// Error satisfies the builtin error interface
func (e Source_ExtractPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSource_ExtractPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Source_ExtractPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Source_ExtractPolicyValidationError{}

// Validate checks the field values on MetricsView_Dimension with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetricsView_Dimension) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsView_Dimension with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetricsView_DimensionMultiError, or nil if none found.
func (m *MetricsView_Dimension) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsView_Dimension) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Label

	// no validation rules for Description

	if len(errors) > 0 {
		return MetricsView_DimensionMultiError(errors)
	}

	return nil
}

// MetricsView_DimensionMultiError is an error wrapping multiple validation
// errors returned by MetricsView_Dimension.ValidateAll() if the designated
// constraints aren't met.
type MetricsView_DimensionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsView_DimensionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsView_DimensionMultiError) AllErrors() []error { return m }

// MetricsView_DimensionValidationError is the validation error returned by
// MetricsView_Dimension.Validate if the designated constraints aren't met.
type MetricsView_DimensionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsView_DimensionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsView_DimensionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsView_DimensionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsView_DimensionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsView_DimensionValidationError) ErrorName() string {
	return "MetricsView_DimensionValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsView_DimensionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsView_Dimension.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsView_DimensionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsView_DimensionValidationError{}

// Validate checks the field values on MetricsView_Measure with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetricsView_Measure) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsView_Measure with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetricsView_MeasureMultiError, or nil if none found.
func (m *MetricsView_Measure) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsView_Measure) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Label

	// no validation rules for Expression

	// no validation rules for Description

	// no validation rules for Format

	if len(errors) > 0 {
		return MetricsView_MeasureMultiError(errors)
	}

	return nil
}

// MetricsView_MeasureMultiError is an error wrapping multiple validation
// errors returned by MetricsView_Measure.ValidateAll() if the designated
// constraints aren't met.
type MetricsView_MeasureMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsView_MeasureMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsView_MeasureMultiError) AllErrors() []error { return m }

// MetricsView_MeasureValidationError is the validation error returned by
// MetricsView_Measure.Validate if the designated constraints aren't met.
type MetricsView_MeasureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsView_MeasureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsView_MeasureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsView_MeasureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsView_MeasureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsView_MeasureValidationError) ErrorName() string {
	return "MetricsView_MeasureValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsView_MeasureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsView_Measure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsView_MeasureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsView_MeasureValidationError{}
