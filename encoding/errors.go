package encoding

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type errType string

const (
	errTypeEncoding = errType("encode")
	errTypeDecoding = errType("decode")
)

// Error is the kind of error to return when there is an encoding/decoding issue.
type Error struct {
	key   errType
	Field string
	Err   error
}

// NewEncodingError creates a new error for an encoding failure.
func NewEncodingError(field string, err error) Error {
	return Error{
		key:   errTypeEncoding,
		Field: field,
		Err:   err,
	}
}

// NewDecodingError creates a new error for a decoding failure.
func NewDecodingError(field string, err error) Error {
	return Error{
		key:   errTypeDecoding,
		Field: field,
		Err:   err,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("couldn't %s %s: %v", e.key, e.Field, e.Err)
}

// Is returns true when both errors are similar.
func (e Error) Is(err error) bool {
	errenc, ok := err.(Error)

	return ok && e.key == errenc.key && errenc.Field == e.Field
}

// Unwrap returns the error wrapped.
func (e Error) Unwrap() error {
	return e.Err
}

// AnyError is the kind of error to return when there is an encoding/decoding
// failure when marshaling/unmarshaling to any.
type AnyError struct {
	key     errType
	Message proto.Message
	Err     error
}

// NewAnyEncodingError creates an error for a marshaling failure.
func NewAnyEncodingError(msg proto.Message, err error) AnyError {
	return AnyError{
		key:     errTypeEncoding,
		Message: msg,
		Err:     err,
	}
}

// NewAnyDecodingError creates an error for a unmarshaling failure.
func NewAnyDecodingError(msg proto.Message, err error) AnyError {
	return AnyError{
		key:     errTypeDecoding,
		Message: msg,
		Err:     err,
	}
}

func (e AnyError) Error() string {
	return fmt.Sprintf("couldn't %s any %T: %v", e.key, e.Message, e.Err)
}

// Is returns true when both errors are similar.
func (e AnyError) Is(err error) bool {
	errenc, ok := err.(AnyError)

	return ok && e.key == errenc.key &&
		reflect.TypeOf(errenc.Message) == reflect.TypeOf(e.Message)
}

// Unwrap returns the error wrapped.
func (e AnyError) Unwrap() error {
	return e.Err
}
