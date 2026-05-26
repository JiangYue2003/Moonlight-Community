package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config user-rpc 配置（合并原 auth-rpc 配置）。
type Config struct {
	zrpc.RpcServerConf

	Mysql      MysqlConf
	CacheRedis cache.CacheConf

	// 业务参数（原 auth-rpc）
	Jwt          JwtConf
	Verification VerificationConf
	Password     PasswordConf
}

type MysqlConf struct {
	DataSource string
}

type JwtConf struct {
	Issuer            string
	PrivateKeyPath    string
	PublicKeyPath     string
	AccessTtlSeconds  int
	RefreshTtlSeconds int
}

type VerificationConf struct {
	CodeLength            int
	CodeTtlSeconds        int
	ResendIntervalSeconds int
	DailyLimit            int
	MaxAttempts           int
}

type PasswordConf struct {
	BcryptCost int
	MinLength  int
}
