package model

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

var reEmail = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type (
	UsersModel interface {
		usersModel
		FindOneByIdentifier(ctx context.Context, identifier string) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

// FindOneByIdentifier 按 phone 或 email 查找用户。
func (m *customUsersModel) FindOneByIdentifier(ctx context.Context, identifier string) (*Users, error) {
	if reEmail.MatchString(identifier) {
		return m.FindOneByEmail(ctx, sql.NullString{String: identifier, Valid: true})
	}
	return m.FindOneByPhone(ctx, sql.NullString{String: identifier, Valid: true})
}
