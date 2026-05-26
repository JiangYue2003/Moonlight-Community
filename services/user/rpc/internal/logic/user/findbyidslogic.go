package userlogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type FindByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdsLogic {
	return &FindByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindByIds 简化实现：逐 id 查询（依赖 model 缓存）。
// 后续高性能场景可替换为 IN 批量查询 + 多键 cache miss 合并。
func (l *FindByIdsLogic) FindByIds(in *user.FindByIdsReq) (*user.FindByIdsResp, error) {
	users := make([]*user.UserInfo, 0, len(in.Ids))
	for _, id := range in.Ids {
		if id <= 0 {
			continue
		}
		u, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(id))
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				continue
			}
			return nil, err
		}
		users = append(users, toUserInfo(u))
	}
	return &user.FindByIdsResp{Users: users}, nil
}
