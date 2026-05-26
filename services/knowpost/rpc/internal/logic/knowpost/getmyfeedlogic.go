package knowpostlogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetMyFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFeedLogic {
	return &GetMyFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMyFeed 整页缓存：feed:mine:{uid}:{size}:{page}
func (l *GetMyFeedLogic) GetMyFeed(in *knowpost.GetMyFeedReq) (*knowpost.FeedPage, error) {
	page, size := normalizePage(in.Page, in.Size)
	key := cachekeys.FeedMineKey(in.CreatorId, size, page)

	resp, err := l.svcCtx.FeedMineCache.GetOrLoad(l.ctx, key, key, func(ctx context.Context) (*knowpost.FeedPage, error) {
		rows, err := l.svcCtx.KnowPostsModel.ListMyFeed(ctx, uint64(in.CreatorId), size+1, (page-1)*size)
		if err != nil {
			return nil, err
		}
		hasMore := len(rows) > size
		if hasMore {
			rows = rows[:size]
		}
		items := make([]*knowpost.FeedItem, 0, len(rows))
		for _, r := range rows {
			items = append(items, rowToFeedItem(r))
		}
		return &knowpost.FeedPage{
			Items:   items,
			HasMore: hasMore,
			Size:    int32(size),
			Page:    int32(page),
		}, nil
	})
	if err != nil && !errors.Is(err, cachex.ErrNotFound) {
		return nil, err
	}
	if resp == nil {
		return &knowpost.FeedPage{Size: int32(size), Page: int32(page)}, nil
	}
	l.svcCtx.HotFeedMine.Hit(key)
	if level := l.svcCtx.HotFeedMine.Level(key); level > 0 {
		_ = l.svcCtx.FeedMineCache.SetWithExtension(l.ctx, key, resp, level)
	}
	return resp, nil
}

func normalizePage(page, size int32) (int, int) {
	p, s := int(page), int(size)
	if p <= 0 {
		p = 1
	}
	if s <= 0 {
		s = 20
	}
	if s > 50 {
		s = 50
	}
	return p, s
}
