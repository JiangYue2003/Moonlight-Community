package errorx

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestNew_FormatsCodeAndMessage(t *testing.T) {
	e := New(CodeBadRequest, "missing field")
	if e.Code != CodeBadRequest {
		t.Fatalf("code mismatch: %s", e.Code)
	}
	if e.Message != "missing field" {
		t.Fatalf("message mismatch: %s", e.Message)
	}
	if e.Cause != nil {
		t.Fatalf("expected nil cause")
	}
	got := e.Error()
	if !strings.Contains(got, CodeBadRequest) || !strings.Contains(got, "missing field") {
		t.Fatalf("Error() should contain code and message: %s", got)
	}
}

func TestWrap_PreservesCauseAndAppendsToString(t *testing.T) {
	cause := errors.New("io: closed")
	e := Wrap(CodeInternalError, "load failed", cause)
	if e.Cause != cause {
		t.Fatal("cause not preserved")
	}
	if !strings.Contains(e.Error(), "io: closed") {
		t.Fatalf("Error() should embed cause: %s", e.Error())
	}
}

func TestUnwrap_ReturnsCause(t *testing.T) {
	cause := errors.New("root cause")
	e := Wrap(CodeNotFound, "wrapper", cause)
	if errors.Unwrap(e) != cause {
		t.Fatal("Unwrap should return the original cause")
	}
}

func TestErrorsIs_FollowsCauseChain(t *testing.T) {
	sentinel := errors.New("sentinel")
	wrapped := Wrap(CodeInternalError, "outer", fmt.Errorf("middle: %w", sentinel))
	if !errors.Is(wrapped, sentinel) {
		t.Fatal("errors.Is should walk through BizError -> middle -> sentinel")
	}
}

func TestAs_ExtractsBizError(t *testing.T) {
	original := New(CodeUnauthorized, "no token")
	wrapped := fmt.Errorf("handler: %w", original)

	be, ok := As(wrapped)
	if !ok {
		t.Fatal("As should report success for wrapped BizError")
	}
	if be.Code != CodeUnauthorized {
		t.Fatalf("As returned wrong code: %s", be.Code)
	}
}

func TestAs_RejectsPlainError(t *testing.T) {
	be, ok := As(errors.New("plain"))
	if ok {
		t.Fatalf("As should report failure for non-BizError; got %v", be)
	}
}

func TestAs_HandlesNilGracefully(t *testing.T) {
	if be, ok := As(nil); ok || be != nil {
		t.Fatal("As(nil) should return (nil, false)")
	}
}
