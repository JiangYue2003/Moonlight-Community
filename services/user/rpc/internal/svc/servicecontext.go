package svc

import (
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	loginmodel "github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model_auth"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/token"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/verification"
)

type ServiceContext struct {
	Config config.Config

	UsersModel     model.UsersModel
	LoginLogsModel loginmodel.LoginLogsModel

	Redis     goredis.UniversalClient
	JwtSigner *jwtx.Signer
	Tokens    *token.Store
	Verifier  *verification.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	priv, err := jwtx.LoadPrivateKey(c.Jwt.PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	pub, err := jwtx.LoadPublicKey(c.Jwt.PublicKeyPath)
	if err != nil {
		panic(err)
	}
	signer, err := jwtx.NewSigner(jwtx.Config{
		PrivateKey: priv,
		PublicKey:  pub,
		Issuer:     c.Jwt.Issuer,
		AccessTtl:  time.Duration(c.Jwt.AccessTtlSeconds) * time.Second,
		RefreshTtl: time.Duration(c.Jwt.RefreshTtlSeconds) * time.Second,
	})
	if err != nil {
		panic(err)
	}

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})

	return &ServiceContext{
		Config:         c,
		UsersModel:     model.NewUsersModel(conn, c.CacheRedis),
		LoginLogsModel: loginmodel.NewLoginLogsModel(conn, c.CacheRedis),
		Redis:          rdb,
		JwtSigner:      signer,
		Tokens:         token.NewStore(rdb),
		Verifier: verification.New(rdb, verification.Config{
			CodeLength:     c.Verification.CodeLength,
			CodeTtl:        time.Duration(c.Verification.CodeTtlSeconds) * time.Second,
			ResendInterval: time.Duration(c.Verification.ResendIntervalSeconds) * time.Second,
			DailyLimit:     c.Verification.DailyLimit,
			MaxAttempts:    c.Verification.MaxAttempts,
		}),
	}
}
