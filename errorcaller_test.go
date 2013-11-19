package errorcaller

import (
	"strings"
	"testing"
)

func TestErrorCaller(t *testing.T) {
	err := New("hello")
	if !strings.Contains(err.Error(), "hello") {
		t.Error("doesn't contain message")
	}
	if !strings.Contains(err.Error(), "errorcaller_test.go:9") {
		t.Error("doesn't contain caller")
	}
	if strings.Contains(err.Error(), "errorcaller.go") {
		t.Error("contains errocaller.go instead of real caller")
	}
}
