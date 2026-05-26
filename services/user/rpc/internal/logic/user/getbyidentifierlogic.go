package userlogic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type GetByIdentifierLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIdentifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdentifierLogic {
	return &GetByIdentifierLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIdentifierLogic) GetByIdentifier(in *user.GetByIdentifierReq) (*user.GetByIdentifierResp, error) {
	id, t := detectIdentifier(in.Identifier)
	var (
		u   *model.Users
		err error
	)
	switch t {
	case IdentifierPhone:
		u, err = l.svcCtx.UsersModel.FindOneByPhone(l.ctx, sql.NullString{String: id, Valid: true})
	case IdentifierEmail:
		u, err = l.svcCtx.UsersModel.FindOneByEmail(l.ctx, sql.NullString{String: id, Valid: true})
	default:
		return nil, errorx.New(errorx.CodeBadRequest, "invalid identifier")
	}
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.New(errorx.CodeIdentifierNotFound, "user not found")
		}
		return nil, err
	}
	return &user.GetByIdentifierResp{User: toUserInfo(u)}, nil
}
