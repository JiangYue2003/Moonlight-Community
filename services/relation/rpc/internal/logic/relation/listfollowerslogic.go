package relationlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/zset"
)

type ListFollowersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowersLogic {
	return &ListFollowersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListFollowers 粉丝列表：先查 ZSet（uf:fans:{userId}）；miss 走 follower 反查表 DB。
func (l *ListFollowersLogic) ListFollowers(in *relation.ListReq) (*relation.ListResp, error) {
	if in.UserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "user id required")
	}
	limit := clampLimit(in.Limit)
	key := zset.FollowerKey(in.UserId)

	var ids []int64
	var nextCursor int64
	if in.Cursor == 0 {
		if top := l.svcCtx.FollowerTopCache[in.UserId]; len(top) > 0 {
			start := int(in.Offset)
			if start < 0 {
				start = 0
			}
			if start < len(top) {
				end := start + limit
				if end > len(top) {
					end = len(top)
				}
				ids = append(ids, top[start:end]...)
			}
		}
	}

	if in.Cursor > 0 {
		hit, cursor, ok, err := zset.PageByCursor(l.ctx, l.svcCtx.Redis, key, in.Cursor, limit)
		if err == nil && ok {
			ids, nextCursor = hit, cursor
		}
	} else if ids == nil && in.Offset >= 0 {
		hit, ok, err := zset.PageByOffset(l.ctx, l.svcCtx.Redis, key, int(in.Offset), limit)
		if err == nil && ok {
			ids = hit
		}
	}

	if ids == nil {
		rows, err := l.svcCtx.FollowerModel.PageActive(l.ctx, in.UserId, limit, int(in.Offset), in.Cursor)
		if err != nil {
			return nil, err
		}
		for _, r := range rows {
			ids = append(ids, int64(r.FromUserId))
			nextCursor = r.CreatedAt.UnixMilli()
		}
	}

	users, err := hydrateUsers(l.ctx, l.svcCtx, ids)
	if err != nil {
		return nil, err
	}
	return &relation.ListResp{
		Items:      users,
		NextCursor: nextCursor,
		HasMore:    len(ids) == limit,
	}, nil
}
