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

type ExistsByZgIdExceptIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistsByZgIdExceptIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistsByZgIdExceptIdLogic {
	return &ExistsByZgIdExceptIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ExistsByZgIdExceptId 用唯一索引 zg_id 直查；命中且不是 except_id 自己 → 冲突。
// 命中自己（resoufce update no-op）→ 不冲突，返回 false。
func (l *ExistsByZgIdExceptIdLogic) ExistsByZgIdExceptId(in *user.ExistsByZgIdExceptIdReq) (*user.ExistsByZgIdExceptIdResp, error) {
	if in.ZgId == "" {
		return &user.ExistsByZgIdExceptIdResp{Exists: false}, nil
	}
	u, err := l.svcCtx.UsersModel.FindOneByZgId(l.ctx, sql.NullString{String: in.ZgId, Valid: true})
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &user.ExistsByZgIdExceptIdResp{Exists: false}, nil
		}
		return nil, err
	}
	if int64(u.Id) == in.ExceptId {
		return &user.ExistsByZgIdExceptIdResp{Exists: false}, nil
	}
	return &user.ExistsByZgIdExceptIdResp{Exists: true}, nil
}
