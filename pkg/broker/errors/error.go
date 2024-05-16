package errors

import (
	"fmt"
)

// Code represents type of error code.
type Code int

// Error code enumaration.
const (
	CodeOK Code = iota
	CodeNotFound
	CodeServiceUnavailable
	CodeMalformMessage
	CodeNoHandler
	CodeNoTube
	CodeTimeout
	CodeInternal
	CodeUnknown
)

// mapCodeString ...
var mapCodeString = map[Code]string{
	CodeOK:                 "CodeOK",
	CodeNotFound:           "CodeNotFound",
	CodeServiceUnavailable: "CodeServiceUnavailable",
	CodeMalformMessage:     "CodeMalformMessage",
	CodeNoHandler:          "CodeNoHandler",
	CodeNoTube:             "CodeNoTube",
	CodeTimeout:            "CodeTimeout",
	CodeInternal:           "CodeInternal",
	CodeUnknown:            "CodeUnknown",
}

// String ...
func (c Code) String() string {
	s, ok := mapCodeString[c]
	if ok {
		return s
	}
	return fmt.Sprintf("Unknown code %d", c)
}

type handlerError struct {
	c   Code
	msg string
}

// Error ...
func (err handlerError) Error() string {
	return fmt.Sprintf("handler error: code = %d desc = %s", err.c, err.msg)
}

// Errorf returns an error containing an error code and a description; Errorf returns nil if c is CodeOK.
func Errorf(c Code, format string, a ...interface{}) error {
	return handlerError{c, fmt.Sprintf(format, a...)}
}

// ErrCode returns the error code for err if it was produced by handler system.
// Otherwise, it returns CodeUnknown.
func ErrCode(err error) Code {
	if err == nil {
		return CodeOK
	}

	if err, ok := err.(handlerError); ok {
		return err.c
	}
	return CodeUnknown
}

// ErrorDesc returns the error description of err if it was produced by the rpc system.
// Otherwise, it returns err.Error() or empty string when err is nil.
func ErrorDesc(err error) string {
	if err == nil {
		return ""
	}

	if err, ok := err.(handlerError); ok {
		return err.msg
	}
	return err.Error()
}
