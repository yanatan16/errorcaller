// Package errorcaller provides an error type that contains its stack. This is nice for errors passed along channels instead of callstacks.
package errorcaller

import (
	"fmt"
	"runtime"
)

type ErrorCaller struct {
	msg, stack string
}

func new(msg, stack string) error {
	return &ErrorCaller{
		msg:   msg,
		stack: stack,
	}
}

func New(msg string) error {
	return new(msg, captureStack(0))
}

func Err(err error) error {
	if ecerr, ok := err.(*ErrorCaller); ok {
		// Don't modify a ErrorCaller
		return ecerr
	}
	return new(err.Error(), captureStack(0))
}

func NewDeep(msg string, n int) error {
	return new(msg, captureStack(n))
}

func ErrDeep(err error, n int) error {
	if ecerr, ok := err.(*ErrorCaller); ok {
		return ecerr
	}
	return new(err.Error(), captureStack(n))
}

func (ecerr *ErrorCaller) Error() string {
	return ecerr.msg + " " + ecerr.stack
}

func captureStack(n int) string {
	_, file, line, ok := runtime.Caller(n + 2)
	if ok {
		return fmt.Sprintf("(on %s:%d)", file, line)
	}
	return ""
}
