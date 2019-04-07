package errors

import (
	errors2 "errors"
	"github.com/object-it/goserv/errors"
	"testing"
)

func TestErrorHasRootCause(t *testing.T) {
	err := errors.New("caller", "error message", errors2.New("my root cause"))

	cause := errors.Cause(err)

	if cause == nil || cause.Error() != "my root cause" {
		t.Error("No Root Cause detected !!!")
	}
}
