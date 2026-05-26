package authlogic

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type PasswordResetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordResetLogic {
	return &PasswordResetLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *PasswordResetLogic) PasswordReset(in *user.PasswordResetReq) (*user.Empty, error) {
	id := normalizeIdentifier(in.Identifier)
	if err := validateIdentifier(id); err != nil {
		return nil, err
	}
	if err := validatePassword(in.NewPassword, l.svcCtx.Config.Password.MinLength); err != nil {
		return nil, err
	}
	if err := l.svcCtx.Verifier.Verify(l.ctx, "RESET_PASSWORD", id, in.Code); err != nil {
		return nil, err
	}

	u, err := l.svcCtx.UsersModel.FindOneByIdentifier(l.ctx, id)
	if err != nil || u == nil {
		return nil, errorx.New(errorx.CodeIdentifierNotFound, "user not found")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), l.svcCtx.Config.Password.BcryptCost)
	if err != nil {
		return nil, err
	}
	u.PasswordHash = sql.NullString{String: string(hash), Valid: true}
	if err := l.svcCtx.UsersModel.Update(l.ctx, u); err != nil {
		return nil, err
	}
	_ = l.svcCtx.Tokens.RevokeAll(l.ctx, int64(u.Id))
	return &user.Empty{}, nil
}
