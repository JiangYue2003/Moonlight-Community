package userlogic

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create 仅持久化基础字段；幂等性由上层（auth-rpc）的 ExistsByIdentifier 保护。
func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	if in.Nickname == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "nickname required")
	}
	if in.Phone == "" && in.Email == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "phone or email required")
	}
	row := &model.Users{
		Phone:        nullable(in.Phone),
		Email:        nullable(in.Email),
		PasswordHash: nullable(in.PasswordHash),
		Nickname:     in.Nickname,
	}
	res, err := l.svcCtx.UsersModel.Insert(l.ctx, row)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &user.CreateResp{Id: id}, nil
}

func nullable(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
