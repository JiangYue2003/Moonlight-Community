package knowpostlogic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
)

const safetyBumpLogInterval = 10 * time.Second

var safetyBumpLog atomic.Int64

// invalidateKnowPostCaches 执行 knowpost 写路径的缓存失效。
//
// 对齐策略：
// 1) 详情缓存（L1/L2）
// 2) feed:item:{postId}（L1/L2）
// 3) 公共 Feed 反向索引命中的页面（ids + hasMore + L1 整页）
// 4) 作者 mine feed（按前缀扫描并双删）
func invalidateKnowPostCaches(ctx context.Context, sc *svc.ServiceContext, postID, creatorID int64) {
	if sc == nil || postID <= 0 {
		return
	}

	if sc.DetailCache != nil {
		_ = sc.DetailCache.Invalidate(ctx, cachekeys.DetailKey(postID))
	}

	itemKey := cachekeys.FeedItemKey(postID)
	if sc.L2 != nil {
		_ = sc.L2.Del(ctx, itemKey)
	}
	if sc.L1FeedItem != nil {
		sc.L1FeedItem.Del(itemKey)
	}

	invalidatePublicFeedPagesByPost(ctx, sc, postID)
	invalidateMineFeedPagesByCreator(ctx, sc, creatorID)
}

// bumpFeedPageSafety invalidates every previously materialized personal Feed
// page after a content mutation. The database write is already committed at
// this point, so an epoch-store failure must degrade reads to cache bypass
// instead of changing the mutation's business result.
func bumpFeedPageSafety(ctx context.Context, sc *svc.ServiceContext) {
	if sc == nil || sc.FeedEpochs == nil {
		return
	}
	if _, err := sc.FeedEpochs.BumpSafety(ctx); err != nil {
		if !sc.FeedEpochs.SafetyPending() {
			sc.FeedEpochs.MarkSafetyPending()
		}
		logSafetyBumpFailure(ctx, err)
	}
}

func logSafetyBumpFailure(ctx context.Context, err error) {
	now := time.Now().UnixNano()
	for {
		last := safetyBumpLog.Load()
		if last != 0 && time.Duration(now-last) < safetyBumpLogInterval {
			return
		}
		if safetyBumpLog.CompareAndSwap(last, now) {
			logx.WithContext(ctx).Errorf("Feed PageCache safety epoch bump failed; reads will bypass cache: %v", err)
			return
		}
	}
}

func invalidatePublicFeedPagesByPost(ctx context.Context, sc *svc.ServiceContext, postID int64) {
	if sc == nil || sc.Redis == nil {
		return
	}
	curr := cachekeys.HourSlot(time.Now())
	for _, hour := range []int64{curr, curr - 1} {
		ridx := cachekeys.FeedReverseIndexKey(postID, hour)
		pageKeys, err := sc.Redis.SMembers(ctx, ridx).Result()
		if err != nil || len(pageKeys) == 0 {
			_ = sc.Redis.Del(ctx, ridx).Err()
			continue
		}

		delKeys := make([]string, 0, len(pageKeys)*3)
		for _, pageKey := range pageKeys {
			delKeys = append(delKeys, pageKey)
			_ = sc.Redis.SRem(ctx, cachekeys.FeedAllPagesKey, pageKey).Err()
			if size, page, ok := parseFeedPublicPageKey(pageKey); ok {
				if sc.L1FeedPublic != nil {
					sc.L1FeedPublic.Del(cachekeys.FeedPublicL1Key(size, page))
				}
				for _, idsKey := range findFeedPublicIDsKeys(size, page, []int64{hour, hour - 1, hour + 1}) {
					delKeys = append(delKeys, idsKey, cachekeys.FeedPublicHasMoreKey(idsKey))
				}
			}
		}
		_ = sc.Redis.Del(ctx, delKeys...).Err()
		_ = sc.Redis.Del(ctx, ridx).Err()
	}
}

func invalidateMineFeedPagesByCreator(ctx context.Context, sc *svc.ServiceContext, creatorID int64) {
	if sc == nil || sc.Redis == nil || sc.FeedMineCache == nil || creatorID <= 0 {
		return
	}
	pattern := fmt.Sprintf("feed:mine:%d:*", creatorID)
	var cursor uint64
	for {
		keys, next, err := sc.Redis.Scan(ctx, cursor, pattern, 200).Result()
		if err != nil {
			return
		}
		if len(keys) > 0 {
			_ = sc.FeedMineCache.Invalidate(ctx, keys...)
		}
		cursor = next
		if cursor == 0 {
			return
		}
	}
}

// parseFeedPublicIDsKey 解析 key: feed:public:ids:{size}:{hour}:{page}
func parseFeedPublicIDsKey(idsKey string) (size int, page int, ok bool) {
	parts := strings.Split(idsKey, ":")
	if len(parts) != 6 || parts[0] != "feed" || parts[1] != "public" || parts[2] != "ids" {
		return 0, 0, false
	}
	sz, err := strconv.Atoi(parts[3])
	if err != nil || sz <= 0 {
		return 0, 0, false
	}
	pg, err := strconv.Atoi(parts[5])
	if err != nil || pg <= 0 {
		return 0, 0, false
	}
	return sz, pg, true
}

func parseFeedPublicPageKey(pageKey string) (size int, page int, ok bool) {
	parts := strings.Split(pageKey, ":")
	if len(parts) != 5 || parts[0] != "feed" || parts[1] != "public" || parts[4] != "v1" {
		return 0, 0, false
	}
	sz, err := strconv.Atoi(parts[2])
	if err != nil || sz <= 0 {
		return 0, 0, false
	}
	pg, err := strconv.Atoi(parts[3])
	if err != nil || pg <= 0 {
		return 0, 0, false
	}
	return sz, pg, true
}

func findFeedPublicIDsKeys(size, page int, hours []int64) []string {
	keys := make([]string, 0, len(hours))
	for _, hour := range hours {
		if hour < 0 {
			continue
		}
		keys = append(keys, cachekeys.FeedPublicIdsKey(size, page, hour))
	}
	return keys
}
