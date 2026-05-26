package userlogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type GetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIdLogic) GetById(in *user.GetByIdReq) (*user.GetByIdResp, error) {
	if in.Id <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid user id")
	}
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, "user not found")
		}
		return nil, err
	}
	return &user.GetByIdResp{User: toUserInfo(u)}, nil
}
