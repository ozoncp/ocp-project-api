// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-project-api/ocp-project-api.proto

package ocp_project_api

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

// Validate checks the field values on ListProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProjectsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListProjectsRequestValidationError is the validation error returned by
// ListProjectsRequest.Validate if the designated constraints aren't met.
type ListProjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsRequestValidationError) ErrorName() string {
	return "ListProjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsRequestValidationError{}

// Validate checks the field values on ListProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProjectsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProjectsResponseValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListProjectsResponseValidationError is the validation error returned by
// ListProjectsResponse.Validate if the designated constraints aren't met.
type ListProjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsResponseValidationError) ErrorName() string {
	return "ListProjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsResponseValidationError{}

// Validate checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCourseId() <= 0 {
		return CreateProjectRequestValidationError{
			field:  "CourseId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Name

	return nil
}

// CreateProjectRequestValidationError is the validation error returned by
// CreateProjectRequest.Validate if the designated constraints aren't met.
type CreateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectRequestValidationError) ErrorName() string {
	return "CreateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectRequestValidationError{}

// Validate checks the field values on CreateProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProjectId

	return nil
}

// CreateProjectResponseValidationError is the validation error returned by
// CreateProjectResponse.Validate if the designated constraints aren't met.
type CreateProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectResponseValidationError) ErrorName() string {
	return "CreateProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectResponseValidationError{}

// Validate checks the field values on MultiCreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateProjectRequestValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateProjectRequestValidationError is the validation error returned by
// MultiCreateProjectRequest.Validate if the designated constraints aren't met.
type MultiCreateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateProjectRequestValidationError) ErrorName() string {
	return "MultiCreateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateProjectRequestValidationError{}

// Validate checks the field values on MultiCreateProjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CountOfCreated

	return nil
}

// MultiCreateProjectResponseValidationError is the validation error returned
// by MultiCreateProjectResponse.Validate if the designated constraints aren't met.
type MultiCreateProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateProjectResponseValidationError) ErrorName() string {
	return "MultiCreateProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateProjectResponseValidationError{}

// Validate checks the field values on RemoveProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProjectId() <= 0 {
		return RemoveProjectRequestValidationError{
			field:  "ProjectId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveProjectRequestValidationError is the validation error returned by
// RemoveProjectRequest.Validate if the designated constraints aren't met.
type RemoveProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveProjectRequestValidationError) ErrorName() string {
	return "RemoveProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveProjectRequestValidationError{}

// Validate checks the field values on RemoveProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveProjectResponseValidationError is the validation error returned by
// RemoveProjectResponse.Validate if the designated constraints aren't met.
type RemoveProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveProjectResponseValidationError) ErrorName() string {
	return "RemoveProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveProjectResponseValidationError{}

// Validate checks the field values on DescribeProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProjectId() <= 0 {
		return DescribeProjectRequestValidationError{
			field:  "ProjectId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeProjectRequestValidationError is the validation error returned by
// DescribeProjectRequest.Validate if the designated constraints aren't met.
type DescribeProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeProjectRequestValidationError) ErrorName() string {
	return "DescribeProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeProjectRequestValidationError{}

// Validate checks the field values on DescribeProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeProjectResponseValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeProjectResponseValidationError is the validation error returned by
// DescribeProjectResponse.Validate if the designated constraints aren't met.
type DescribeProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeProjectResponseValidationError) ErrorName() string {
	return "DescribeProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeProjectResponseValidationError{}

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Project) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for CourseId

	// no validation rules for Name

	return nil
}

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on NewProject with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NewProject) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CourseId

	// no validation rules for Name

	return nil
}

// NewProjectValidationError is the validation error returned by
// NewProject.Validate if the designated constraints aren't met.
type NewProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewProjectValidationError) ErrorName() string { return "NewProjectValidationError" }

// Error satisfies the builtin error interface
func (e NewProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewProjectValidationError{}
