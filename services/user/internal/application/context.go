package application

import (
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	"github.com/zhiguang/zhiguang-go/services/user/internal/adapter/model"
	loginmodel "github.com/zhiguang/zhiguang-go/services/user/internal/adapter/model_auth"
	"github.com/zhiguang/zhiguang-go/services/user/internal/adapter/token"
	"github.com/zhiguang/zhiguang-go/services/user/internal/adapter/verification"
)

type PasswordPolicy struct {
	BcryptCost int
	MinLength  int
}

type ServiceContext struct {
	UsersModel     model.UsersModel
	LoginLogsModel loginmodel.LoginLogsModel

	Redis     goredis.UniversalClient
	JwtSigner *jwtx.Signer
	Tokens    *token.Store
	Verifier  *verification.Service
	Password  PasswordPolicy
}
