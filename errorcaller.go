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
	return new(msg, captureStack())
}

func Err(err error) error {
	if ws, ok := err.(*ErrorCaller); ok {
		// Don't modify a ErrorCaller
		return ws
	}
	return new(err.Error(), captureStack())
}

func (ws *ErrorCaller) Error() string {
	return ws.msg + " " + ws.stack
}

func captureStack() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("(on %s:%d)", file, line)
	}
	return ""
}
