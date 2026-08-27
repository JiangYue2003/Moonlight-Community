package knowpostlogic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/userfeed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feedcursor"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const maxPersonalFeedCandidates = 5000

const pageCacheEpochLogInterval = 10 * time.Second

var pageCacheEpochLog atomic.Int64

type feedItemCacheScope struct {
	enabled      bool
	safetyScoped bool
	safetyEpoch  uint64
}

var legacyFeedItemCacheScope = feedItemCacheScope{enabled: true}

func safetyFeedItemCacheScope(safetyEpoch uint64) feedItemCacheScope {
	return feedItemCacheScope{enabled: true, safetyScoped: true, safetyEpoch: safetyEpoch}
}

func (s feedItemCacheScope) key(postID int64) string {
	if s.safetyScoped {
		return cachekeys.PersonalFeedItemKey(postID, s.safetyEpoch)
	}
	return cachekeys.FeedItemKey(postID)
}

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFeed 获取用户的个性化Feed流（推拉混合）
func (l *GetUserFeedLogic) GetUserFeed(in *knowpost.GetUserFeedReq) (response *knowpost.FeedPage, err error) {
	observer := l.svcCtx.FeedObserver
	paginationMode := feed.PaginationModePage
	if in.GetCursor() != "" {
		paginationMode = feed.PaginationModeCursor
	}
	paginationResult := feed.PaginationResultUnknown
	paginationStarted := time.Now()
	defer func() {
		if paginationResult == feed.PaginationResultUnknown {
			paginationResult = paginationResultFromError(err)
		}
		feed.RecordPagination(observer, paginationMode, paginationResult, time.Since(paginationStarted))
	}()
	if observer != nil {
		totalStarted := time.Now()
		defer func() {
			feed.RecordStage(observer, feed.StageTotal, feed.OutcomeFromError(err), time.Since(totalStarted))
		}()
	}
	if in.GetCursor() != "" {
		if !l.svcCtx.Config.Feed.CursorPagination.Enabled {
			paginationResult = feed.PaginationResultDisabled
			return nil, status.Error(codes.InvalidArgument, "feed cursor pagination is disabled")
		}
		decodeStarted := time.Now()
		after, decodeErr := feedcursor.Decode(in.GetCursor())
		feed.RecordStage(observer, feed.StageCursorDecode, feed.OutcomeFromError(decodeErr), time.Since(decodeStarted))
		if decodeErr != nil {
			paginationResult = feed.PaginationResultInvalid
			return nil, status.Error(codes.InvalidArgument, "invalid feed cursor")
		}
		_, size := normalizePage(1, in.Size)
		feedItems := legacyFeedItemCacheScope
		if l.pageCacheConfigured() {
			feedItems = feedItemCacheScope{}
		}
		response, err := l.computeUserFeedAfter(l.ctx, in.UserId, after, size, feedItems)
		feed.RecordPageCache(l.svcCtx.FeedObserver, feed.PageCacheBypass, feed.OutcomeFromError(err))
		return response, err
	}
	page, size := normalizePage(in.Page, in.Size)
	if !l.pageCacheConfigured() {
		return l.computePageCacheBypass(in, page, size, legacyFeedItemCacheScope)
	}
	if !l.pageCacheEligible(page, size) || l.svcCtx.FeedEpochs == nil || l.svcCtx.FeedPageCache == nil {
		return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
	}
	if l.svcCtx.FeedEpochs.SafetyPending() {
		if epochErr := l.svcCtx.FeedEpochs.FlushPendingSafety(l.ctx); epochErr != nil {
			l.logPageCacheEpochBypass("safety_flush", epochErr)
			return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
		}
		if l.svcCtx.FeedEpochs.SafetyPending() {
			return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
		}
	}
	relationEpoch, epochErr := l.svcCtx.FeedEpochs.Relation(l.ctx, in.UserId)
	if epochErr != nil {
		l.logPageCacheEpochBypass("relation", epochErr)
		return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
	}
	safetyEpoch, epochErr := l.svcCtx.FeedEpochs.Safety(l.ctx)
	if epochErr != nil {
		l.logPageCacheEpochBypass("safety", epochErr)
		return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
	}
	feedItems := safetyFeedItemCacheScope(safetyEpoch)
	request := userfeed.Request{
		UserID:          in.UserId,
		StrategyVersion: userfeed.StrategyHybridV1,
		RelationEpoch:   relationEpoch,
		SafetyEpoch:     safetyEpoch,
		Page:            int32(page),
		Size:            int32(size),
	}
	response, source, err := l.svcCtx.FeedPageCache.GetOrLoad(l.ctx, request, func(loaderCtx context.Context) (*knowpost.FeedPage, error) {
		return l.computeUserFeed(loaderCtx, in, page, size, feedItems)
	})
	if err == nil {
		current, verifyErr := l.pageCacheEpochsStillCurrent(l.ctx, in.UserId, relationEpoch, safetyEpoch)
		if verifyErr != nil {
			l.logPageCacheEpochBypass("revalidate", verifyErr)
		}
		if verifyErr != nil || !current {
			return l.computePageCacheBypass(in, page, size, feedItemCacheScope{})
		}
	}
	feed.RecordPageCache(l.svcCtx.FeedObserver, pageCacheMetricSource(source), feed.OutcomeFromError(err))
	return response, err
}

func paginationResultFromError(err error) feed.FeedPaginationResult {
	if err == nil {
		return feed.PaginationResultSuccess
	}
	if errors.Is(err, feed.ErrCursorReaderUnavailable) {
		return feed.PaginationResultReaderUnavailable
	}
	switch feed.OutcomeFromError(err) {
	case feed.OutcomeTimeout:
		return feed.PaginationResultTimeout
	case feed.OutcomeCanceled:
		return feed.PaginationResultCanceled
	default:
		return feed.PaginationResultError
	}
}

func (l *GetUserFeedLogic) pageCacheEpochsStillCurrent(
	ctx context.Context,
	userID int64,
	relationEpoch uint64,
	safetyEpoch uint64,
) (bool, error) {
	if l.svcCtx.FeedEpochs.SafetyPending() {
		return false, nil
	}
	currentRelation, err := l.svcCtx.FeedEpochs.Relation(ctx, userID)
	if err != nil {
		return false, err
	}
	if l.svcCtx.FeedEpochs.SafetyPending() {
		return false, nil
	}
	currentSafety, err := l.svcCtx.FeedEpochs.Safety(ctx)
	if err != nil {
		return false, err
	}
	if l.svcCtx.FeedEpochs.SafetyPending() {
		return false, nil
	}
	return currentRelation == relationEpoch && currentSafety == safetyEpoch, nil
}

func (l *GetUserFeedLogic) logPageCacheEpochBypass(kind string, err error) {
	now := time.Now().UnixNano()
	for {
		last := pageCacheEpochLog.Load()
		if last != 0 && time.Duration(now-last) < pageCacheEpochLogInterval {
			return
		}
		if pageCacheEpochLog.CompareAndSwap(last, now) {
			l.Logger.Errorf("Feed PageCache epoch %s failure; bypassing cache: %v", kind, err)
			return
		}
	}
}

func (l *GetUserFeedLogic) pageCacheConfigured() bool {
	mode := strings.ToLower(strings.TrimSpace(l.svcCtx.Config.Feed.PageCache.Mode))
	return mode == string(userfeed.ModeL2) || mode == string(userfeed.ModeL1L2)
}

func (l *GetUserFeedLogic) pageCacheEligible(page, size int) bool {
	strategy, err := feed.ParseStrategy(l.svcCtx.Config.Feed.Strategy)
	return err == nil && strategy == feed.StrategyHybrid && page == 1 && size == 20
}

func (l *GetUserFeedLogic) computePageCacheBypass(
	in *knowpost.GetUserFeedReq,
	page int,
	size int,
	feedItems feedItemCacheScope,
) (*knowpost.FeedPage, error) {
	response, err := l.computeUserFeed(l.ctx, in, page, size, feedItems)
	feed.RecordPageCache(l.svcCtx.FeedObserver, feed.PageCacheBypass, feed.OutcomeFromError(err))
	return response, err
}

func pageCacheMetricSource(source userfeed.Source) feed.FeedPageCacheSource {
	switch source {
	case userfeed.SourceBypass:
		return feed.PageCacheBypass
	case userfeed.SourceL1Fresh:
		return feed.PageCacheL1Fresh
	case userfeed.SourceL2Fresh:
		return feed.PageCacheL2Fresh
	case userfeed.SourceL1Stale:
		return feed.PageCacheL1Stale
	case userfeed.SourceL2Stale:
		return feed.PageCacheL2Stale
	case userfeed.SourceLoad:
		return feed.PageCacheMiss
	default:
		return feed.PageCacheUnknown
	}
}

func (l *GetUserFeedLogic) computeUserFeed(
	ctx context.Context,
	in *knowpost.GetUserFeedReq,
	page int,
	size int,
	feedItems feedItemCacheScope,
) (response *knowpost.FeedPage, err error) {
	observer := l.svcCtx.FeedObserver

	start, end, candidateLimit, ok := personalFeedWindow(page, size)
	if !ok {
		return &knowpost.FeedPage{
			Items: []*knowpost.FeedItem{},
			Page:  int32(page),
			Size:  int32(size),
		}, nil
	}
	if observer != nil {
		defer func() {
			feed.RecordColdCompute(observer, feed.OutcomeFromError(err))
		}()
	}
	loaded := make(map[int64]*knowpost.FeedItem, candidateLimit)
	seen := make(map[int64]struct{}, candidateLimit)
	var visibleItems []*knowpost.FeedItem
	var visiblePosts []feed.Post
	snapshot, err := l.svcCtx.FeedReader.Prepare(ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("FeedReader.Prepare failed: user=%d, err=%v", in.UserId, err)
		return nil, err
	}

	for {
		posts, rawHasMore, err := snapshot.GetFeedPosts(ctx, 1, candidateLimit)
		if err != nil {
			l.Logger.Errorf("FeedReader.GetFeed failed: user=%d, page=%d, size=%d, err=%v",
				in.UserId, page, size, err)
			return nil, err
		}

		newIDs := make([]int64, 0, len(posts))
		for _, post := range posts {
			if _, ok := seen[post.ID]; ok {
				continue
			}
			seen[post.ID] = struct{}{}
			newIDs = append(newIDs, post.ID)
		}
		feed.RecordCursorWork(observer, feed.CursorWorkHydrateIDs, len(newIDs))
		var hydrateStarted time.Time
		if observer != nil {
			hydrateStarted = time.Now()
		}
		batch, err := l.loadFeedItems(ctx, newIDs, feedItems)
		if observer != nil {
			feed.RecordStage(observer, feed.StageHydrate, feed.OutcomeFromError(err), time.Since(hydrateStarted))
		}
		if err != nil {
			return nil, err
		}
		for id, item := range batch {
			loaded[id] = item
		}

		visibleItems = visibleItems[:0]
		visiblePosts = visiblePosts[:0]
		for _, post := range posts {
			if item := loaded[post.ID]; feedItemAllowed(snapshot, item) {
				visibleItems = append(visibleItems, item)
				visiblePosts = append(visiblePosts, post)
			}
		}

		if len(visibleItems) > end || !rawHasMore || candidateLimit >= maxPersonalFeedCandidates {
			break
		}
		candidateLimit *= 2
		feed.RecordCursorWork(observer, feed.CursorWorkBackfillRounds, 1)
		if candidateLimit > maxPersonalFeedCandidates {
			candidateLimit = maxPersonalFeedCandidates
		}
	}

	hasMore := len(visibleItems) > end
	if start >= len(visibleItems) {
		visibleItems = []*knowpost.FeedItem{}
		visiblePosts = []feed.Post{}
	} else {
		if end > len(visibleItems) {
			end = len(visibleItems)
		}
		visibleItems = visibleItems[start:end]
		visiblePosts = visiblePosts[start:end]
	}

	nextCursor := ""
	if l.svcCtx.Config.Feed.CursorPagination.Enabled && hasMore && len(visiblePosts) > 0 {
		nextCursor, err = encodeFeedCursor(visiblePosts[len(visiblePosts)-1])
		if err != nil {
			return nil, err
		}
	}

	return &knowpost.FeedPage{
		Items:      visibleItems,
		HasMore:    hasMore,
		Page:       int32(page),
		Size:       int32(size),
		NextCursor: nextCursor,
	}, nil
}

func (l *GetUserFeedLogic) computeUserFeedAfter(
	ctx context.Context,
	userID int64,
	after feedcursor.Cursor,
	size int,
	feedItems feedItemCacheScope,
) (response *knowpost.FeedPage, err error) {
	observer := l.svcCtx.FeedObserver
	if observer != nil {
		defer func() {
			feed.RecordColdCompute(observer, feed.OutcomeFromError(err))
		}()
	}

	candidateLimit := size + 1
	loaded := make(map[int64]*knowpost.FeedItem, candidateLimit)
	seen := make(map[int64]struct{}, candidateLimit)
	var visibleItems []*knowpost.FeedItem
	var visiblePosts []feed.Post
	snapshot, err := l.svcCtx.FeedReader.Prepare(ctx, userID)
	if err != nil {
		l.Logger.Errorf("FeedReader.Prepare failed: user=%d, err=%v", userID, err)
		return nil, err
	}

	for {
		posts, rawHasMore, err := snapshot.GetFeedAfter(ctx, after, candidateLimit)
		if err != nil {
			l.Logger.Errorf("FeedReader.GetFeedAfter failed: user=%d, size=%d, err=%v", userID, size, err)
			return nil, err
		}

		newIDs := make([]int64, 0, len(posts))
		for _, post := range posts {
			if _, ok := seen[post.ID]; ok {
				continue
			}
			seen[post.ID] = struct{}{}
			newIDs = append(newIDs, post.ID)
		}
		feed.RecordCursorWork(observer, feed.CursorWorkHydrateIDs, len(newIDs))
		var hydrateStarted time.Time
		if observer != nil {
			hydrateStarted = time.Now()
		}
		batch, err := l.loadFeedItems(ctx, newIDs, feedItems)
		if observer != nil {
			feed.RecordStage(observer, feed.StageHydrate, feed.OutcomeFromError(err), time.Since(hydrateStarted))
		}
		if err != nil {
			return nil, err
		}
		for id, item := range batch {
			loaded[id] = item
		}

		visibleItems = visibleItems[:0]
		visiblePosts = visiblePosts[:0]
		for _, post := range posts {
			if item := loaded[post.ID]; feedItemAllowed(snapshot, item) {
				visibleItems = append(visibleItems, item)
				visiblePosts = append(visiblePosts, post)
			}
		}

		if len(visibleItems) > size || !rawHasMore || candidateLimit >= maxPersonalFeedCandidates {
			break
		}
		candidateLimit *= 2
		feed.RecordCursorWork(observer, feed.CursorWorkBackfillRounds, 1)
		if candidateLimit > maxPersonalFeedCandidates {
			candidateLimit = maxPersonalFeedCandidates
		}
	}

	hasMore := len(visibleItems) > size
	if len(visibleItems) > size {
		visibleItems = visibleItems[:size]
		visiblePosts = visiblePosts[:size]
	}
	nextCursor := ""
	if hasMore && len(visiblePosts) > 0 {
		nextCursor, err = encodeFeedCursor(visiblePosts[len(visiblePosts)-1])
		if err != nil {
			return nil, err
		}
	}

	return &knowpost.FeedPage{
		Items:      visibleItems,
		HasMore:    hasMore,
		Size:       int32(size),
		Page:       0,
		NextCursor: nextCursor,
	}, nil
}

func encodeFeedCursor(post feed.Post) (string, error) {
	return feedcursor.Encode(feedcursor.Cursor{
		SortTime: post.CreateTime,
		PostID:   post.ID,
	})
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

func (l *GetUserFeedLogic) loadFeedItems(
	ctx context.Context,
	postIDs []int64,
	feedItems feedItemCacheScope,
) (map[int64]*knowpost.FeedItem, error) {
	items := make(map[int64]*knowpost.FeedItem, len(postIDs))
	if len(postIDs) == 0 {
		return items, nil
	}

	misses := append([]int64(nil), postIDs...)
	if feedItems.enabled && l.svcCtx.Redis != nil {
		keys := make([]string, len(postIDs))
		for i, postID := range postIDs {
			keys[i] = feedItems.key(postID)
		}
		raws, err := l.svcCtx.Redis.MGet(ctx, keys...).Result()
		feed.RecordDependency(
			l.svcCtx.FeedObserver,
			feed.DependencyRedis,
			feed.OperationFeedItemMGet,
			feed.OutcomeFromError(err),
		)
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
	rows, err := loader.FindPublishedFeedByIDs(ctx, ids)
	feed.RecordDependency(
		l.svcCtx.FeedObserver,
		feed.DependencyMySQL,
		feed.OperationFeedItemDB,
		feed.OutcomeFromError(err),
	)
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
	l.writeFeedItemCache(ctx, dbItems, feedItems)
	return items, nil
}

func (l *GetUserFeedLogic) writeFeedItemCache(
	ctx context.Context,
	items []*knowpost.FeedItem,
	feedItems feedItemCacheScope,
) {
	if !feedItems.enabled || l.svcCtx.Redis == nil || len(items) == 0 {
		return
	}
	pipe := l.svcCtx.Redis.Pipeline()
	for _, item := range items {
		raw, err := json.Marshal(item)
		if err != nil {
			continue
		}
		pipe.Set(ctx, feedItems.key(parseInt64(item.Id)), raw,
			cachex.Jitter(cachekeys.FeedItemBaseTTL, cachekeys.FeedItemJitterMax))
	}
	_, err := pipe.Exec(ctx)
	feed.RecordDependency(
		l.svcCtx.FeedObserver,
		feed.DependencyRedis,
		feed.OperationFeedItemCacheWrite,
		feed.OutcomeFromError(err),
	)
	if err != nil {
		l.Logger.Errorf("personal feed item cache writeback failed: %v", err)
	}
}

func feedItemVisibilityAllowed(visible string) bool {
	return visible == "public" || visible == "followers"
}

func feedItemAllowed(snapshot *feed.FeedReadSnapshot, item *knowpost.FeedItem) bool {
	return item != nil && snapshot.AllowsCreator(item.CreatorId) && feedItemVisibilityAllowed(item.Visible)
}
