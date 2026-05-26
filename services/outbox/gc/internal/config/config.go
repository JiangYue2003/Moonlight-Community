package config

import "github.com/zeromicro/go-zero/core/stores/sqlx"

type Config struct {
	Name           string
	DataSource     string
	IntervalHours  int `json:",default=1"`
	RetainDays     int `json:",default=7"`
	BatchSize      int `json:",default=500"`
	BatchIntervalMs int `json:",default=50"`
}

// SqlConn 返回数据库连接（go-zero sqlx）。
func (c Config) SqlConn() sqlx.SqlConn {
	return sqlx.NewMysql(c.DataSource)
}
