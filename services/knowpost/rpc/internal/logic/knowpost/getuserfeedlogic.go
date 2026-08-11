package knowpostlogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const maxPersonalFeedCandidates = 5000

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFeed 获取用户的个性化Feed流（推拉混合）
func (l *GetUserFeedLogic) GetUserFeed(in *knowpost.GetUserFeedReq) (*knowpost.FeedPage, error) {
	page, size := normalizePage(in.Page, in.Size)

	start, end, candidateLimit, ok := personalFeedWindow(page, size)
	if !ok {
		return &knowpost.FeedPage{
			Items: []*knowpost.FeedItem{},
			Page:  int32(page),
			Size:  int32(size),
		}, nil
	}
	loaded := make(map[int64]*knowpost.FeedItem, candidateLimit)
	seen := make(map[int64]struct{}, candidateLimit)
	var visibleItems []*knowpost.FeedItem
	snapshot, err := l.svcCtx.FeedReader.Prepare(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("FeedReader.Prepare failed: user=%d, err=%v", in.UserId, err)
		return nil, err
	}

	for {
		postIDs, rawHasMore, err := snapshot.GetFeed(l.ctx, 1, candidateLimit)
		if err != nil {
			l.Logger.Errorf("FeedReader.GetFeed failed: user=%d, page=%d, size=%d, err=%v",
				in.UserId, page, size, err)
			return nil, err
		}

		newIDs := make([]int64, 0, len(postIDs))
		for _, postID := range postIDs {
			if _, ok := seen[postID]; ok {
				continue
			}
			seen[postID] = struct{}{}
			newIDs = append(newIDs, postID)
		}
		batch, err := l.loadFeedItems(newIDs)
		if err != nil {
			return nil, err
		}
		for id, item := range batch {
			loaded[id] = item
		}

		visibleItems = visibleItems[:0]
		for _, postID := range postIDs {
			if item := loaded[postID]; feedItemAllowed(snapshot, item) {
				visibleItems = append(visibleItems, item)
			}
		}

		if len(visibleItems) > end || !rawHasMore || candidateLimit >= maxPersonalFeedCandidates {
			break
		}
		candidateLimit *= 2
		if candidateLimit > maxPersonalFeedCandidates {
			candidateLimit = maxPersonalFeedCandidates
		}
	}

	hasMore := len(visibleItems) > end
	if start >= len(visibleItems) {
		visibleItems = []*knowpost.FeedItem{}
	} else {
		if end > len(visibleItems) {
			end = len(visibleItems)
		}
		visibleItems = visibleItems[start:end]
	}

	return &knowpost.FeedPage{
		Items:   visibleItems,
		HasMore: hasMore,
		Page:    int32(page),
		Size:    int32(size),
	}, nil
}

func personalFeedWindow(page, size int) (start, end, candidateLimit int, ok bool) {
	if page <= 0 || size <= 0 {
		return 0, 0, 0, false
	}
	start64 := int64(page-1) * int64(size)
	if start64 >= maxPersonalFeedCandidates {
		return 0, 0, 0, false
	}
	end64 := start64 + int64(size)
	if end64 > maxPersonalFeedCandidates {
		end64 = maxPersonalFeedCandidates
	}
	candidate64 := end64 + int64(size)
	if candidate64 > maxPersonalFeedCandidates {
		candidate64 = maxPersonalFeedCandidates
	}
	return int(start64), int(end64), int(candidate64), true
}

func (l *GetUserFeedLogic) loadFeedItems(postIDs []int64) (map[int64]*knowpost.FeedItem, error) {
	items := make(map[int64]*knowpost.FeedItem, len(postIDs))
	if len(postIDs) == 0 {
		return items, nil
	}

	misses := append([]int64(nil), postIDs...)
	if l.svcCtx.Redis != nil {
		keys := make([]string, len(postIDs))
		for i, postID := range postIDs {
			keys[i] = cachekeys.FeedItemKey(postID)
		}
		raws, err := l.svcCtx.Redis.MGet(l.ctx, keys...).Result()
		if err == nil {
			misses = misses[:0]
			for i, raw := range raws {
				encoded, ok := raw.(string)
				if !ok || encoded == "" {
					misses = append(misses, postIDs[i])
					continue
				}
				var item knowpost.FeedItem
				if err := json.Unmarshal([]byte(encoded), &item); err != nil || !feedItemVisibilityAllowed(item.Visible) {
					misses = append(misses, postIDs[i])
					continue
				}
				items[postIDs[i]] = &item
			}
		}
	}

	if len(misses) == 0 {
		return items, nil
	}
	loader := l.svcCtx.FeedPostLoader
	if loader == nil {
		loader = l.svcCtx.KnowPostsModel
	}
	if loader == nil {
		return nil, fmt.Errorf("feed post loader is not configured")
	}
	ids := make([]uint64, len(misses))
	for i, id := range misses {
		ids[i] = uint64(id)
	}
	rows, err := loader.FindPublishedFeedByIDs(l.ctx, ids)
	if err != nil {
		return nil, err
	}
	dbItems := make([]*knowpost.FeedItem, 0, len(rows))
	for _, row := range rows {
		if row == nil {
			continue
		}
		item := rowToFeedItem(row)
		items[int64(row.Id)] = item
		dbItems = append(dbItems, item)
	}
	l.writeFeedItemCache(dbItems)
	return items, nil
}

func (l *GetUserFeedLogic) writeFeedItemCache(items []*knowpost.FeedItem) {
	if l.svcCtx.Redis == nil || len(items) == 0 {
		return
	}
	pipe := l.svcCtx.Redis.Pipeline()
	for _, item := range items {
		raw, err := json.Marshal(item)
		if err != nil {
			continue
		}
		pipe.Set(l.ctx, cachekeys.FeedItemKey(parseInt64(item.Id)), raw,
			cachex.Jitter(cachekeys.FeedItemBaseTTL, cachekeys.FeedItemJitterMax))
	}
	if _, err := pipe.Exec(l.ctx); err != nil {
		l.Logger.Errorf("personal feed item cache writeback failed: %v", err)
	}
}

func feedItemVisibilityAllowed(visible string) bool {
	return visible == "public" || visible == "followers"
}

func feedItemAllowed(snapshot *feed.FeedReadSnapshot, item *knowpost.FeedItem) bool {
	return item != nil && snapshot.AllowsCreator(item.CreatorId) && feedItemVisibilityAllowed(item.Visible)
}
