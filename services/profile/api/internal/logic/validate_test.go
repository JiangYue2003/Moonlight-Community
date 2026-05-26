package logic

import (
	"strings"
	"testing"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/types"
)

func sp(s string) *string { return &s }

func TestValidatePatchReq_RequiresAtLeastOneField(t *testing.T) {
	if err := validatePatchReq(&types.PatchProfileReq{}); err == nil {
		t.Fatal("empty patch should fail")
	}
}

func TestValidatePatchReq_NicknameLength(t *testing.T) {
	if err := validatePatchReq(&types.PatchProfileReq{Nickname: sp("")}); err == nil {
		t.Fatal("empty nickname must fail")
	}
	long := strings.Repeat("a", 65)
	if err := validatePatchReq(&types.PatchProfileReq{Nickname: sp(long)}); err == nil {
		t.Fatal(">64 chars must fail")
	}
	if err := validatePatchReq(&types.PatchProfileReq{Nickname: sp("ok")}); err != nil {
		t.Fatalf("valid nickname should pass: %v", err)
	}
}

func TestValidatePatchReq_BioMaxLen(t *testing.T) {
	long := strings.Repeat("中", 513)
	err := validatePatchReq(&types.PatchProfileReq{Bio: sp(long)})
	if err == nil {
		t.Fatal(">512 runes must fail")
	}
}

func TestValidatePatchReq_GenderEnum(t *testing.T) {
	for _, g := range []string{"MALE", "FEMALE", "Other", "unknown"} {
		if err := validatePatchReq(&types.PatchProfileReq{Gender: sp(g)}); err != nil {
			t.Errorf("gender %q should pass: %v", g, err)
		}
	}
	if err := validatePatchReq(&types.PatchProfileReq{Gender: sp("???")}); err == nil {
		t.Fatal("invalid gender must fail")
	}
}

func TestValidatePatchReq_BirthdayFormatAndPast(t *testing.T) {
	if err := validatePatchReq(&types.PatchProfileReq{Birthday: sp("2099-01-01")}); err == nil {
		t.Fatal("future birthday must fail")
	}
	if err := validatePatchReq(&types.PatchProfileReq{Birthday: sp("2000-01-02")}); err != nil {
		t.Fatalf("past birthday should pass: %v", err)
	}
	if err := validatePatchReq(&types.PatchProfileReq{Birthday: sp("01/02/2000")}); err == nil {
		t.Fatal("non-ISO format must fail")
	}
}

func TestValidatePatchReq_ZgIdRegex(t *testing.T) {
	bads := []string{"a", "abc", "with space", "long-with-dash", strings.Repeat("a", 33), "包含中文"}
	for _, z := range bads {
		if err := validatePatchReq(&types.PatchProfileReq{ZgId: sp(z)}); err == nil {
			t.Errorf("zgId %q should fail validation", z)
		}
	}
	for _, z := range []string{"abcd", "user_42", "ABC123_xyz"} {
		if err := validatePatchReq(&types.PatchProfileReq{ZgId: sp(z)}); err != nil {
			t.Errorf("zgId %q should pass: %v", z, err)
		}
	}
	// 显式空 zgId（清空）应该通过 regex 检查（regex 仅在非空时校验）
	if err := validatePatchReq(&types.PatchProfileReq{ZgId: sp("")}); err != nil {
		t.Fatalf("empty zgId (clear) should pass: %v", err)
	}
}

func TestValidatePatchReq_SchoolMax(t *testing.T) {
	long := strings.Repeat("a", 129)
	if err := validatePatchReq(&types.PatchProfileReq{School: sp(long)}); err == nil {
		t.Fatal(">128 must fail")
	}
}

func TestValidatePatchReq_BadRequestCode(t *testing.T) {
	err := validatePatchReq(&types.PatchProfileReq{Gender: sp("xxx")})
	be, ok := errorx.As(err)
	if !ok || be.Code != errorx.CodeBadRequest {
		t.Fatalf("validation error must be CodeBadRequest, got %v", err)
	}
}
