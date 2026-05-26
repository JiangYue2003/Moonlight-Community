package userlogic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *user.UpdatePasswordReq) (*user.UpdatePasswordResp, error) {
	if in.Id <= 0 || in.PasswordHash == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "id and password_hash required")
	}
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, "user not found")
		}
		return nil, err
	}
	u.PasswordHash = sql.NullString{String: in.PasswordHash, Valid: true}
	if err := l.svcCtx.UsersModel.Update(l.ctx, u); err != nil {
		return nil, err
	}
	return &user.UpdatePasswordResp{Ok: true}, nil
}
