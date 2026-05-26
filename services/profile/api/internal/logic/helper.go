package logic

import (
	"regexp"
	"strings"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

var zgIdRe = regexp.MustCompile(`^[a-zA-Z0-9_]{4,32}$`)

// validatePatchReq 与原 Java ProfilePatchRequest 校验注解严格对齐。
func validatePatchReq(req *types.PatchProfileReq) error {
	hasAny := false
	if req.Nickname != nil {
		hasAny = true
		n := strings.TrimSpace(*req.Nickname)
		if len([]rune(n)) < 1 || len([]rune(n)) > 64 {
			return errorx.New(errorx.CodeBadRequest, "nickname length must be 1..64")
		}
	}
	if req.Bio != nil {
		hasAny = true
		if len([]rune(*req.Bio)) > 512 {
			return errorx.New(errorx.CodeBadRequest, "bio length must be <=512")
		}
	}
	if req.Gender != nil {
		hasAny = true
		switch strings.ToUpper(*req.Gender) {
		case "MALE", "FEMALE", "OTHER", "UNKNOWN":
		default:
			return errorx.New(errorx.CodeBadRequest, "invalid gender")
		}
	}
	if req.Birthday != nil {
		hasAny = true
		if *req.Birthday != "" {
			t, err := time.Parse("2006-01-02", *req.Birthday)
			if err != nil {
				return errorx.New(errorx.CodeBadRequest, "birthday format must be YYYY-MM-DD")
			}
			if t.After(time.Now()) {
				return errorx.New(errorx.CodeBadRequest, "birthday cannot be in the future")
			}
		}
	}
	if req.ZgId != nil {
		hasAny = true
		if *req.ZgId != "" && !zgIdRe.MatchString(*req.ZgId) {
			return errorx.New(errorx.CodeBadRequest, "zgId must match [a-zA-Z0-9_]{4,32}")
		}
	}
	if req.School != nil {
		hasAny = true
		if len([]rune(*req.School)) > 128 {
			return errorx.New(errorx.CodeBadRequest, "school length must be <=128")
		}
	}
	if req.TagsJson != nil {
		hasAny = true
	}
	if req.TagJson != nil {
		hasAny = true
	}
	if req.Phone != nil {
		hasAny = true
	}
	if req.Email != nil {
		hasAny = true
	}
	if !hasAny {
		return errorx.New(errorx.CodeBadRequest, "no field to update")
	}
	return nil
}

// toProfileResp 把 user-rpc UserInfo 翻译为 HTTP ProfileResp。
func toProfileResp(u *userpb.UserInfo) *types.ProfileResp {
	if u == nil {
		return &types.ProfileResp{}
	}
	return &types.ProfileResp{
		Id: u.Id, Nickname: u.Nickname, Avatar: u.Avatar, Bio: u.Bio,
		ZgId: u.ZgId, Gender: u.Gender, Birthday: u.Birthday, School: u.School,
		Phone: u.Phone, Email: u.Email, TagsJson: u.TagsJson, TagJson: u.TagsJson,
	}
}
