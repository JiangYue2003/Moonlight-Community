package userlogic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ExistsByIdentifierLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistsByIdentifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistsByIdentifierLogic {
	return &ExistsByIdentifierLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExistsByIdentifierLogic) ExistsByIdentifier(in *user.ExistsByIdentifierReq) (*user.ExistsByIdentifierResp, error) {
	id, t := detectIdentifier(in.Identifier)
	if t == IdentifierUnknown {
		return &user.ExistsByIdentifierResp{Exists: false}, nil
	}
	switch t {
	case IdentifierPhone:
		u, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, sql.NullString{String: id, Valid: true})
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return &user.ExistsByIdentifierResp{Exists: false}, nil
			}
			return nil, err
		}
		return &user.ExistsByIdentifierResp{Exists: true, UserId: int64(u.Id)}, nil
	case IdentifierEmail:
		u, err := l.svcCtx.UsersModel.FindOneByEmail(l.ctx, sql.NullString{String: id, Valid: true})
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return &user.ExistsByIdentifierResp{Exists: false}, nil
			}
			return nil, err
		}
		return &user.ExistsByIdentifierResp{Exists: true, UserId: int64(u.Id)}, nil
	}
	return &user.ExistsByIdentifierResp{Exists: false}, nil
}
