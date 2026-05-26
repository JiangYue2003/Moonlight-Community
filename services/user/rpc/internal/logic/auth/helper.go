package authlogic

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	loginmodel "github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model_auth"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

var (
	phoneRe = regexp.MustCompile(`^1\d{10}$`)
	emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)
)

func normalizeIdentifier(s string) string {
	s = strings.TrimSpace(s)
	if emailRe.MatchString(strings.ToLower(s)) {
		return strings.ToLower(s)
	}
	return s
}

func validateIdentifier(s string) error {
	if phoneRe.MatchString(s) || emailRe.MatchString(strings.ToLower(s)) {
		return nil
	}
	return errorx.New(errorx.CodeBadRequest, "invalid identifier")
}

func validatePassword(p string, minLen int) error {
	if len(p) < minLen {
		return errorx.New(errorx.CodePasswordPolicyViolation, "password too short")
	}
	hasLetter, hasDigit := false, false
	for _, r := range p {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
			hasLetter = true
		case r >= '0' && r <= '9':
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return errorx.New(errorx.CodePasswordPolicyViolation, "password must contain letters and digits")
	}
	return nil
}

// modelToAuthUser 将 model.Users 转换为 user.AuthUser（合并后直接用 model，不再走 RPC）。
func modelToAuthUser(u *model.Users) *user.AuthUser {
	if u == nil {
		return nil
	}
	birthday := ""
	if u.Birthday.Valid {
		birthday = u.Birthday.Time.Format("2006-01-02")
	}
	return &user.AuthUser{
		Id:       int64(u.Id),
		Nickname: u.Nickname,
		Avatar:   u.Avatar.String,
		Phone:    u.Phone.String,
		ZgId:     u.ZgId.String,
		Birthday: birthday,
		School:   u.School.String,
		Bio:      u.Bio.String,
		Gender:   u.Gender.String,
		TagsJson: u.TagsJson.String,
	}
}

func issueAndPersist(ctx context.Context, sc *svc.ServiceContext, uid int64, nickname string) (*user.TokenPair, error) {
	pair, err := sc.JwtSigner.IssuePair(uid, nickname)
	if err != nil {
		return nil, err
	}
	if err := sc.Tokens.Save(ctx, uid, pair.RefreshTokenId, time.Until(pair.RefreshExpiresAt)); err != nil {
		return nil, err
	}
	return &user.TokenPair{
		AccessToken:      pair.AccessToken,
		AccessExpiresAt:  pair.AccessExpiresAt.UnixMilli(),
		RefreshToken:     pair.RefreshToken,
		RefreshExpiresAt: pair.RefreshExpiresAt.UnixMilli(),
		RefreshTokenId:   pair.RefreshTokenId,
	}, nil
}

func recordLoginLog(ctx context.Context, sc *svc.ServiceContext, userId int64, identifier, channel, ip, ua, status string) {
	row := &loginmodel.LoginLogs{
		UserId:     sql.NullInt64{Int64: userId, Valid: userId > 0},
		Identifier: identifier,
		Channel:    channel,
		Ip:         sql.NullString{String: ip, Valid: ip != ""},
		UserAgent:  sql.NullString{String: ua, Valid: ua != ""},
		Status:     status,
		CreatedAt:  time.Now(),
	}
	if _, err := sc.LoginLogsModel.Insert(ctx, row); err != nil {
		_ = err
	}
}
