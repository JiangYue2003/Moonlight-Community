package knowpostlogic

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDetail 三级缓存路径：
//  1. cachex.GetOrLoad(L1 → L2 → loader)
//  2. loader 中读 DB；不存在/已删除/私有不可见 → ErrNotFound（cachex 自动写哨兵）
//  3. 命中后 hot.Hit + Level≥LOW 时 SetWithExtension 延长 TTL
func (l *GetDetailLogic) GetDetail(in *knowpost.GetDetailReq) (*knowpost.KnowPostDetail, error) {
	if in.Id <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid id")
	}
	key := cachekeys.DetailKey(in.Id)
	sfKey := "detail:" + key

	d, err := l.svcCtx.DetailCache.GetOrLoad(l.ctx, key, sfKey, func(ctx context.Context) (*knowpost.KnowPostDetail, error) {
		row, err := l.svcCtx.KnowPostsModel.FindOne(ctx, uint64(in.Id))
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, cachex.ErrNotFound
			}
			return nil, err
		}
		if row.Status == "deleted" {
			return nil, cachex.ErrNotFound
		}
		if !(row.Status == "published" && row.Visible == "public") && int64(row.CreatorId) != in.ViewerId {
			return nil, cachex.ErrNotFound
		}
		return rowToDetail(row), nil
	})
	if err != nil {
		if errors.Is(err, cachex.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, "post not found")
		}
		return nil, err
	}

	l.svcCtx.HotDetail.Hit(key)
	if level := l.svcCtx.HotDetail.Level(key); level > 0 {
		_ = l.svcCtx.DetailCache.SetWithExtension(l.ctx, key, d, level)
		l.extendFeedItemTTL(in.Id, level)
	}
	return d, nil
}

func (l *GetDetailLogic) extendFeedItemTTL(postID int64, level hotkey.Level) {
	if l.svcCtx == nil || l.svcCtx.Redis == nil || postID <= 0 {
		return
	}
	itemKey := cachekeys.FeedItemKey(postID)
	ttl := hotkey.TTLForPublic(cachekeys.FeedItemBaseTTL, level)
	if ttl <= 0 {
		return
	}
	currentTTL, err := l.svcCtx.Redis.TTL(l.ctx, itemKey).Result()
	if err != nil {
		return
	}
	if currentTTL <= 0 || currentTTL >= ttl {
		return
	}
	_ = l.svcCtx.Redis.Expire(l.ctx, itemKey, time.Duration(ttl)).Err()
}
