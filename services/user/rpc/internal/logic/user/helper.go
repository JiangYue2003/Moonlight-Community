package userlogic

import (
	"database/sql"
	"regexp"
	"strings"

	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

var (
	phoneRe = regexp.MustCompile(`^1\d{10}$`)
	emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)
)

// IdentifierType 与 Java 端 IdentifierType 枚举对齐。
type IdentifierType int

const (
	IdentifierUnknown IdentifierType = iota
	IdentifierPhone
	IdentifierEmail
)

// detectIdentifier 推断标识符类型；非法输入返回 Unknown。
func detectIdentifier(s string) (string, IdentifierType) {
	s = strings.TrimSpace(s)
	switch {
	case phoneRe.MatchString(s):
		return s, IdentifierPhone
	case emailRe.MatchString(strings.ToLower(s)):
		return strings.ToLower(s), IdentifierEmail
	default:
		return s, IdentifierUnknown
	}
}

// nullStr 把可空 sql.NullString 转换为 string，nil/无效→""。
func nullStr(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}

// toUserInfo 将 model.Users 转为 pb.UserInfo。
func toUserInfo(u *model.Users) *user.UserInfo {
	if u == nil {
		return nil
	}
	birthday := ""
	if u.Birthday.Valid {
		birthday = u.Birthday.Time.Format("2006-01-02")
	}
	return &user.UserInfo{
		Id:           int64(u.Id),
		Phone:        nullStr(u.Phone),
		Email:        nullStr(u.Email),
		Nickname:     u.Nickname,
		Avatar:       nullStr(u.Avatar),
		ZgId:         nullStr(u.ZgId),
		Gender:       nullStr(u.Gender),
		Birthday:     birthday,
		School:       nullStr(u.School),
		Bio:          nullStr(u.Bio),
		TagsJson:     nullStr(u.TagsJson),
		PasswordHash: nullStr(u.PasswordHash),
		CreatedAt:    u.CreatedAt.UnixMilli(),
		UpdatedAt:    u.UpdatedAt.UnixMilli(),
	}
}
