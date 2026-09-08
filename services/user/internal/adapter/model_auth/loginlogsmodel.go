package model_auth

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LoginLogsModel = (*customLoginLogsModel)(nil)

type (
	// LoginLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginLogsModel.
	LoginLogsModel interface {
		loginLogsModel
	}

	customLoginLogsModel struct {
		*defaultLoginLogsModel
	}
)

// NewLoginLogsModel returns a model for the database table.
func NewLoginLogsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LoginLogsModel {
	return &customLoginLogsModel{
		defaultLoginLogsModel: newLoginLogsModel(conn, c, opts...),
	}
}
