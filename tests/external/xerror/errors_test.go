package xerror

import (
	"errors"
	"github.com/object-it/goserv/xerror"
	"testing"
)

func TestErrorHasRootCause(t *testing.T) {
	err := xerror.New("caller", "error message", errors.New("my root cause"))

	cause := xerror.Cause(err)

	if cause == nil || cause.Error() != "my root cause" {
		t.Error("No Root Cause detected !!!")
	}
}

func TestErrorStringOutput(t *testing.T) {
	err := xerror.New("caller", "error message", errors.New("my root cause"))

	if err.Error() != "caller : error message - Caused by : my root cause" {
		t.Error("Bad error output !!!")
	}
}
