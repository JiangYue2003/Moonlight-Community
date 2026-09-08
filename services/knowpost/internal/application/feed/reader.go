package feed

import (
	"container/heap"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/feedcursor"
	"golang.org/x/sync/singleflight"
)

const (
	maxBigVPullConcurrency    = 16
	maxBigVOutboxPipelineSize = 128
	feedRouteLookupTimeout    = 2 * time.Second
)

// Post 帖子基本信息（用于排序和去重）
type Post struct {
	ID         int64
	CreatorID  int64
	CreateTime int64
}

// FeedReader 负责读取 Feed 流（推拉混合）
type FeedReader struct {
	redis                     RedisClient
	relationClient            RelationClient
	counterClient             CounterClient
	logger                    logx.Logger
	observer                  FeedObserver
	strategy                  Strategy
	tierMode                  AuthorTierMode
	tierResolver              *AuthorTierResolver
	routeCache                FeedRouteCache
	routeCacheTTL             time.Duration
	relationEpochs            RelationEpochReader
	routeSnapshot             bool
	combinedPipeline          bool
	combinedPipelineBatchSize int
	routeGroup                singleflight.Group
	routeEpochLog             atomic.Int64
}

// FeedRouteCache 保存读者关注列表对应的大V分类结果。
type FeedRouteCache interface {
	Get(key string) (any, bool)
	SetWithTTL(key string, val any, cost int64, ttl time.Duration) bool
}

type feedRouteCacheEntry struct {
	followings []int64
	bigVs      []int64
}

type redisBatchReader interface {
	ZRevRangeWithScoresBatch(ctx context.Context, requests []ZRevRangeRequest) []ZRevRangeResult
}

type redisCursorBatchReader interface {
	ZRevRangeByScoreWithScoresBatch(ctx context.Context, requests []ZRevRangeByScoreRequest) []ZRevRangeResult
}

var ErrCursorReaderUnavailable = errors.New("feed cursor reader is unavailable")

// FeedReaderOptions configures optional read-path behavior while preserving
// the existing constructors for callers that do not need instrumentation.
type FeedReaderOptions struct {
	Strategy                  Strategy
	TierMode                  AuthorTierMode
	TierResolver              *AuthorTierResolver
	RouteCache                FeedRouteCache
	RouteCacheTTL             time.Duration
	RouteSnapshotEnabled      bool
	RelationEpochs            RelationEpochReader
	CombinedPipelineEnabled   bool
	CombinedPipelineBatchSize int
	Observer                  FeedObserver
}

// FeedReadSnapshot 固化一次请求中的关注关系和大V分类，供过滤回填复用。
type FeedReadSnapshot struct {
	reader     *FeedReader
	userID     int64
	bigVs      []int64
	followings map[int64]struct{}
}

func NewFeedReader(redis RedisClient, relationClient RelationClient, counterClient CounterClient, logger logx.Logger) *FeedReader {
	return NewFeedReaderWithStrategy(redis, relationClient, counterClient, logger, StrategyHybrid)
}

func NewFeedReaderWithStrategy(redis RedisClient, relationClient RelationClient, counterClient CounterClient, logger logx.Logger, strategy Strategy) *FeedReader {
	return NewFeedReaderWithStrategyAndRouteCache(redis, relationClient, counterClient, logger, strategy, nil, 0)
}

func NewFeedReaderWithStrategyAndRouteCache(
	redis RedisClient,
	relationClient RelationClient,
	counterClient CounterClient,
	logger logx.Logger,
	strategy Strategy,
	routeCache FeedRouteCache,
	routeCacheTTL time.Duration,
) *FeedReader {
	return NewFeedReaderWithOptions(
		redis,
		relationClient,
		counterClient,
		logger,
		FeedReaderOptions{
			Strategy:      strategy,
			RouteCache:    routeCache,
			RouteCacheTTL: routeCacheTTL,
		},
	)
}

func NewFeedReaderWithOptions(
	redis RedisClient,
	relationClient RelationClient,
	counterClient CounterClient,
	logger logx.Logger,
	options FeedReaderOptions,
) *FeedReader {
	strategy := options.Strategy
	if strategy == "" {
		strategy = StrategyHybrid
	}
	tierMode := options.TierMode
	if tierMode == "" {
		tierMode = AuthorTierModeOff
	}
	combinedPipelineBatchSize := options.CombinedPipelineBatchSize
	if combinedPipelineBatchSize < 2 || combinedPipelineBatchSize > maxBigVOutboxPipelineSize {
		combinedPipelineBatchSize = maxBigVOutboxPipelineSize
	}
	return &FeedReader{
		redis:                     redis,
		relationClient:            relationClient,
		counterClient:             counterClient,
		logger:                    logger,
		observer:                  options.Observer,
		strategy:                  strategy,
		tierMode:                  tierMode,
		tierResolver:              options.TierResolver,
		routeCache:                options.RouteCache,
		routeCacheTTL:             options.RouteCacheTTL,
		relationEpochs:            options.RelationEpochs,
		routeSnapshot:             options.RouteSnapshotEnabled,
		combinedPipeline:          options.CombinedPipelineEnabled,
		combinedPipelineBatchSize: combinedPipelineBatchSize,
	}
}

// GetFeed 读取用户的 Feed 流（推拉混合）
// 返回帖子 ID 列表，调用方负责批量查询详情
func (r *FeedReader) GetFeed(ctx context.Context, userID int64, page, size int) ([]int64, bool, error) {
	snapshot, err := r.Prepare(ctx, userID)
	if err != nil {
		return nil, false, err
	}
	return snapshot.GetFeed(ctx, page, size)
}

// Prepare 获取并分类一次关注列表；同一请求内的候选回填应复用返回的快照。
func (r *FeedReader) Prepare(ctx context.Context, userID int64) (*FeedReadSnapshot, error) {
	if r.routeSnapshot && r.relationEpochs != nil && r.routeCache != nil && r.routeCacheTTL > 0 {
		return r.prepareVersionedRoute(ctx, userID)
	}
	return r.prepareLegacyRoute(ctx, userID)
}

func (r *FeedReader) prepareLegacyRoute(ctx context.Context, userID int64) (*FeedReadSnapshot, error) {
	relationStarted := observationStarted(r.observer)
	followings, err := r.relationClient.GetFollowings(ctx, userID)
	relationOutcome := OutcomeFromError(err)
	recordStageSince(r.observer, StageRelation, relationOutcome, relationStarted)
	RecordDependency(r.observer, DependencyRelation, OperationListFollowings, relationOutcome)
	if err != nil {
		r.logger.Errorf("get followings failed: user=%d, err=%v", userID, err)
		return nil, fmt.Errorf("get followings failed: %w", err)
	}

	routeStarted := observationStarted(r.observer)
	bigVs, err := r.bigVsForUser(ctx, userID, followings)
	recordStageSince(r.observer, StageRoute, OutcomeFromError(err), routeStarted)
	if err != nil {
		return nil, err
	}
	followingSet := make(map[int64]struct{}, len(followings))
	for _, followingID := range followings {
		followingSet[followingID] = struct{}{}
	}
	return &FeedReadSnapshot{
		reader:     r,
		userID:     userID,
		bigVs:      bigVs,
		followings: followingSet,
	}, nil
}

func (r *FeedReader) bigVsForUser(ctx context.Context, userID int64, followings []int64) ([]int64, error) {
	if r.strategy != StrategyHybrid || r.routeCache == nil || r.routeCacheTTL <= 0 {
		bigVs, _, _, err := r.classifyFollowingsResult(ctx, followings)
		return bigVs, err
	}

	key := fmt.Sprintf("feed:route:%d", userID)
	if entry, ok := r.cachedRoute(key, followings); ok {
		return slices.Clone(entry.bigVs), nil
	}

	// 按“用户 + 完整关注快照”合并并发 miss。关注列表变化时不会错误共享旧分类。
	resultC := r.routeGroup.DoChan(routeFlightKey(key, followings), func() (any, error) {
		// 等待期间可能已有请求回填了缓存，进入 Counter 前再检查一次。
		if entry, ok := r.cachedRoute(key, followings); ok {
			return entry, nil
		}

		// 共享查询不能继承首个请求的取消信号，否则会污染仍有效的等待者；
		// 同时用内部超时限制独立查询的生命周期。
		lookupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), feedRouteLookupTimeout)
		defer cancel()
		bigVs, _, cacheable, err := r.classifyFollowingsResult(lookupCtx, followings)
		if err != nil {
			return nil, err
		}
		entry := feedRouteCacheEntry{
			followings: slices.Clone(followings),
			bigVs:      slices.Clone(bigVs),
		}
		// Counter 异常时沿用全普通用户降级，但不能把降级结果缓存到故障恢复之后。
		if cacheable {
			cost := int64(64 + 8*(len(entry.followings)+len(entry.bigVs)))
			r.routeCache.SetWithTTL(key, entry, cost, r.routeCacheTTL)
		}
		return entry, nil
	})
	select {
	case <-ctx.Done():
		return []int64{}, nil
	case result := <-resultC:
		if result.Err != nil {
			return nil, result.Err
		}
		entry, ok := result.Val.(feedRouteCacheEntry)
		if !ok {
			return nil, fmt.Errorf("route lookup returned %T", result.Val)
		}
		return slices.Clone(entry.bigVs), nil
	}
}

func (r *FeedReader) cachedRoute(key string, followings []int64) (feedRouteCacheEntry, bool) {
	cached, ok := r.routeCache.Get(key)
	if !ok {
		return feedRouteCacheEntry{}, false
	}
	entry, ok := cached.(feedRouteCacheEntry)
	if !ok || !slices.Equal(entry.followings, followings) {
		return feedRouteCacheEntry{}, false
	}
	return entry, true
}

func routeFlightKey(cacheKey string, followings []int64) string {
	encoded := make([]byte, len(followings)*8)
	for i, followingID := range followings {
		binary.LittleEndian.PutUint64(encoded[i*8:], uint64(followingID))
	}
	return cacheKey + ":" + string(encoded)
}

// AllowsCreator 判断候选作者是否仍属于当前用户的关注流；用户自己的内容始终允许。
func (s *FeedReadSnapshot) AllowsCreator(creatorID int64) bool {
	if creatorID == s.userID {
		return true
	}
	_, ok := s.followings[creatorID]
	return ok
}

// GetFeed 从已准备的关注快照读取候选，避免过滤回填重复调用 Relation/Counter。
func (s *FeedReadSnapshot) GetFeed(ctx context.Context, page, size int) ([]int64, bool, error) {
	posts, hasMore, err := s.GetFeedPosts(ctx, page, size)
	if err != nil {
		return nil, false, err
	}
	result := make([]int64, len(posts))
	for i, post := range posts {
		result[i] = post.ID
	}
	return result, hasMore, nil
}

// GetFeedPosts 与 GetFeed 使用相同的页码语义，并保留生成 Cursor 所需的排序键。
func (s *FeedReadSnapshot) GetFeedPosts(ctx context.Context, page, size int) ([]Post, bool, error) {
	r := s.reader

	start := (page - 1) * size
	end := start + size
	candidateLimit := end + 1
	inboxPosts, bigVPosts, err := r.readInboxAndBigVs(ctx, s.userID, s.bigVs, candidateLimit)
	if err != nil {
		r.logger.Errorf("read inbox failed: %v", err)
		// 降级：继续处理，只是收件箱为空
		inboxPosts = []Post{}
	}

	// 5. 归并、去重、排序
	mergeStarted := observationStarted(r.observer)
	allPosts := MergePosts(inboxPosts, bigVPosts)
	allPosts = Deduplicate(allPosts)
	RecordCursorWork(r.observer, CursorWorkMergeCandidates, len(allPosts))

	// 6. 平衡大V内容占比
	allPosts = BalanceFeed(allPosts, 0.5) // 大V内容最多占50%
	recordStageSince(r.observer, StageMergeDedup, OutcomeSuccess, mergeStarted)

	// 7. 分页
	hasMore := len(allPosts) > end

	if start >= len(allPosts) {
		return []Post{}, false, nil
	}

	if end > len(allPosts) {
		end = len(allPosts)
	}

	return append([]Post(nil), allPosts[start:end]...), hasMore, nil
}

// GetFeedAfter 返回严格位于 Cursor 之后的候选；任一来源读取失败时整页失败。
func (s *FeedReadSnapshot) GetFeedAfter(
	ctx context.Context,
	after feedcursor.Cursor,
	limit int,
) (posts []Post, hasMore bool, err error) {
	seekStarted := observationStarted(s.reader.observer)
	defer func() {
		recordStageSince(s.reader.observer, StageCursorSeek, OutcomeFromError(err), seekStarted)
	}()
	if limit <= 0 {
		return []Post{}, false, nil
	}
	batchReader, ok := s.reader.redis.(redisCursorBatchReader)
	if !ok {
		return nil, false, ErrCursorReaderUnavailable
	}

	type source struct {
		key       string
		creatorID int64
	}
	sources := make([]source, 0, len(s.bigVs)+1)
	sources = append(sources, source{key: fmt.Sprintf(FEED_INBOX_KEY, s.userID)})
	for _, creatorID := range s.bigVs {
		sources = append(sources, source{
			key:       fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID),
			creatorID: creatorID,
		})
	}

	topLimit := limit + 1
	unique := make(map[int64]Post, topLimit*len(sources))
	maxRequests := s.reader.combinedPipelineBatchSize
	if maxRequests < 2 {
		maxRequests = maxBigVOutboxPipelineSize
	}
	sourcesPerBatch := maxRequests / 2
	if sourcesPerBatch < 1 {
		sourcesPerBatch = 1
	}
	cursorScore := strconv.FormatInt(after.SortTime, 10)

	for start := 0; start < len(sources); start += sourcesPerBatch {
		if err := ctx.Err(); err != nil {
			return nil, false, err
		}
		end := start + sourcesPerBatch
		if end > len(sources) {
			end = len(sources)
		}
		batchSources := sources[start:end]
		requests := make([]ZRevRangeByScoreRequest, 0, len(batchSources)*2)
		for _, item := range batchSources {
			requests = append(requests,
				ZRevRangeByScoreRequest{Key: item.key, Min: cursorScore, Max: cursorScore},
				ZRevRangeByScoreRequest{
					Key:   item.key,
					Min:   "-inf",
					Max:   "(" + cursorScore,
					Count: int64(topLimit),
				},
			)
		}

		results := batchReader.ZRevRangeByScoreWithScoresBatch(ctx, requests)
		RecordCursorWork(s.reader.observer, CursorWorkRedisCommands, len(requests))
		RecordCursorWork(s.reader.observer, CursorWorkRedisRoundTrips, 1)
		redisMembers := 0
		tieMembers := 0
		for i, result := range results {
			redisMembers += len(result.Scores)
			if i%2 == 0 {
				tieMembers += len(result.Scores)
			}
		}
		RecordCursorWork(s.reader.observer, CursorWorkRedisMembers, redisMembers)
		RecordCursorWork(s.reader.observer, CursorWorkTieMembers, tieMembers)
		boundaryRequests := make([]ZRevRangeByScoreRequest, 0, len(batchSources))
		boundarySources := make([]source, 0, len(batchSources))
		for i, item := range batchSources {
			tieResult := zRevRangeResultAt(results, i*2)
			if tieResult.Err != nil {
				return nil, false, fmt.Errorf("read feed cursor tie range %s: %w", item.key, tieResult.Err)
			}
			olderResult := zRevRangeResultAt(results, i*2+1)
			if olderResult.Err != nil {
				return nil, false, fmt.Errorf("read feed cursor older range %s: %w", item.key, olderResult.Err)
			}

			if err := addCursorScores(unique, tieResult.Scores, item.creatorID, after); err != nil {
				return nil, false, fmt.Errorf("read feed cursor tie range %s: %w", item.key, err)
			}
			if err := addCursorScores(unique, olderResult.Scores, item.creatorID, after); err != nil {
				return nil, false, fmt.Errorf("read feed cursor older range %s: %w", item.key, err)
			}
			if len(olderResult.Scores) == topLimit {
				boundary := olderResult.Scores[len(olderResult.Scores)-1]
				boundaryPost, err := cursorPostFromScore(boundary, item.creatorID)
				if err != nil {
					return nil, false, fmt.Errorf("read feed cursor boundary %s: %w", item.key, err)
				}
				boundaryScore := strconv.FormatInt(boundaryPost.CreateTime, 10)
				boundaryRequests = append(boundaryRequests, ZRevRangeByScoreRequest{
					Key: item.key,
					Min: boundaryScore,
					Max: boundaryScore,
				})
				boundarySources = append(boundarySources, item)
			}
		}

		if len(boundaryRequests) > 0 {
			if err := ctx.Err(); err != nil {
				return nil, false, err
			}
			boundaryResults := batchReader.ZRevRangeByScoreWithScoresBatch(ctx, boundaryRequests)
			RecordCursorWork(s.reader.observer, CursorWorkRedisCommands, len(boundaryRequests))
			RecordCursorWork(s.reader.observer, CursorWorkRedisRoundTrips, 1)
			boundaryMembers := 0
			for _, result := range boundaryResults {
				boundaryMembers += len(result.Scores)
			}
			RecordCursorWork(s.reader.observer, CursorWorkRedisMembers, boundaryMembers)
			RecordCursorWork(s.reader.observer, CursorWorkTieMembers, boundaryMembers)
			for i, item := range boundarySources {
				result := zRevRangeResultAt(boundaryResults, i)
				if result.Err != nil {
					return nil, false, fmt.Errorf("read feed cursor boundary range %s: %w", item.key, result.Err)
				}
				if err := addCursorScores(unique, result.Scores, item.creatorID, after); err != nil {
					return nil, false, fmt.Errorf("read feed cursor boundary range %s: %w", item.key, err)
				}
			}
		}
	}

	RecordCursorWork(s.reader.observer, CursorWorkMergeCandidates, len(unique))
	mergeStarted := observationStarted(s.reader.observer)
	topPosts := make(postMinHeap, 0, topLimit)
	heap.Init(&topPosts)
	for _, post := range unique {
		pushTopPost(&topPosts, post, topLimit)
	}
	sort.Slice(topPosts, func(i, j int) bool {
		return postComesBefore(topPosts[i], topPosts[j])
	})
	hasMore = len(topPosts) > limit
	if hasMore {
		topPosts = topPosts[:limit]
	}
	recordStageSince(s.reader.observer, StageMergeDedup, OutcomeSuccess, mergeStarted)
	return append([]Post(nil), topPosts...), hasMore, nil
}

func cursorPostFromScore(score ZScore, creatorID int64) (Post, error) {
	if score.Member <= 0 {
		return Post{}, fmt.Errorf("post id must be positive")
	}
	if score.Score <= 0 || math.Trunc(score.Score) != score.Score {
		return Post{}, fmt.Errorf("feed score must be a positive integer, got %v", score.Score)
	}
	return Post{
		ID:         score.Member,
		CreatorID:  creatorID,
		CreateTime: int64(score.Score),
	}, nil
}

func addCursorScores(
	unique map[int64]Post,
	scores []ZScore,
	creatorID int64,
	after feedcursor.Cursor,
) error {
	for _, score := range scores {
		post, err := cursorPostFromScore(score, creatorID)
		if err != nil {
			return err
		}
		if !postIsAfterCursor(post, after) {
			continue
		}
		if existing, exists := unique[post.ID]; !exists || postComesBefore(post, existing) {
			unique[post.ID] = post
		}
	}
	return nil
}

func postIsAfterCursor(post Post, after feedcursor.Cursor) bool {
	return post.CreateTime < after.SortTime ||
		(post.CreateTime == after.SortTime && post.ID < after.PostID)
}

// classifyFollowings 区分大V和普通用户
func (r *FeedReader) classifyFollowings(ctx context.Context, followings []int64) (bigVs, normalUsers []int64) {
	bigVs, normalUsers, _, _ = r.classifyFollowingsResult(ctx, followings)
	return bigVs, normalUsers
}

// classifyFollowingsResult 的 cacheable 表示分类结果来自有效路由数据，而非故障降级。
func (r *FeedReader) classifyFollowingsResult(ctx context.Context, followings []int64) (bigVs, normalUsers []int64, cacheable bool, resolveErr error) {
	bigVs = []int64{}
	normalUsers = []int64{}
	if len(followings) == 0 {
		return bigVs, normalUsers, true, nil
	}
	switch r.strategy {
	case StrategyPush:
		return bigVs, append(normalUsers, followings...), true, nil
	case StrategyPull:
		return append(bigVs, followings...), normalUsers, true, nil
	}
	if r.tierMode == AuthorTierModeEnforce {
		return r.classifyFollowingsFromTier(ctx, followings)
	}
	counterStarted := observationStarted(r.observer)
	followerCounts, err := r.counterClient.BatchGetFollowerCounts(ctx, followings)
	counterOutcome := OutcomeFromError(err)
	recordStageSince(r.observer, StageCounter, counterOutcome, counterStarted)
	RecordDependency(r.observer, DependencyCounter, OperationBatchFollowerCounts, counterOutcome)
	if err != nil {
		r.logger.Errorf("batch get follower counts failed: users=%d, err=%v", len(followings), err)
		normalUsers = append(normalUsers, followings...)
		if r.tierMode == AuthorTierModeShadow {
			r.observeShadowTier(ctx, followings, bigVs)
		}
		return bigVs, normalUsers, false, nil
	}

	for _, uid := range followings {
		if followerCounts[uid] > BIGV_THRESHOLD {
			bigVs = append(bigVs, uid)
		} else {
			normalUsers = append(normalUsers, uid)
		}
	}
	if r.tierMode == AuthorTierModeShadow {
		r.observeShadowTier(ctx, followings, bigVs)
	}

	return bigVs, normalUsers, true, nil
}

func (r *FeedReader) classifyFollowingsFromTier(
	ctx context.Context,
	followings []int64,
) (bigVs, normalUsers []int64, cacheable bool, resolveErr error) {
	bigVs = []int64{}
	normalUsers = []int64{}
	if r.tierResolver == nil {
		return bigVs, normalUsers, false, newTierResolutionError(
			"configuration",
			fmt.Errorf("author tier resolver is not configured"),
		)
	}

	resolutions, err := r.tierResolver.ResolveBatch(ctx, followings)
	if err != nil {
		RecordTierResolution(r.observer, TierEvidenceUnknown, OutcomeFromError(err), true, false)
		return bigVs, normalUsers, false, err
	}
	for _, authorID := range followings {
		resolution, ok := resolutions[authorID]
		if !ok {
			err := newTierResolutionError(
				"batch-result",
				fmt.Errorf("tier result missing for author %d", authorID),
			)
			RecordTierResolution(r.observer, TierEvidenceUnknown, OutcomeFromError(err), true, false)
			return []int64{}, []int64{}, false, err
		}
		RecordTierResolution(
			r.observer,
			resolution.Evidence,
			OutcomeSuccess,
			resolution.Fallback,
			false,
		)
		if resolution.BigV {
			bigVs = append(bigVs, authorID)
		} else {
			normalUsers = append(normalUsers, authorID)
		}
	}
	return bigVs, normalUsers, true, nil
}

func (r *FeedReader) observeShadowTier(ctx context.Context, followings, legacyBigVs []int64) {
	if r.tierResolver == nil {
		return
	}
	resolutions, err := r.tierResolver.ResolveBatch(ctx, followings)
	if err != nil {
		RecordTierResolution(r.observer, TierEvidenceUnknown, OutcomeFromError(err), true, false)
		r.logger.Errorf("author tier shadow batch resolution failed: %v", err)
		return
	}

	legacyBigVSet := make(map[int64]struct{}, len(legacyBigVs))
	for _, authorID := range legacyBigVs {
		legacyBigVSet[authorID] = struct{}{}
	}
	for _, authorID := range followings {
		resolution, ok := resolutions[authorID]
		if !ok {
			RecordTierResolution(r.observer, TierEvidenceUnknown, OutcomeError, true, false)
			r.logger.Errorf("author tier shadow result missing for author %d", authorID)
			continue
		}
		_, legacyBigV := legacyBigVSet[authorID]
		mismatch := resolution.BigV != legacyBigV
		RecordTierResolution(
			r.observer,
			resolution.Evidence,
			OutcomeSuccess,
			resolution.Fallback,
			mismatch,
		)
	}
}

// readInbox 读取用户的收件箱
func (r *FeedReader) readInbox(ctx context.Context, userID int64, limit int) ([]Post, error) {
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)

	// 从 ZSet 读取（按时间倒序）
	started := observationStarted(r.observer)
	zs, err := r.redis.ZRevRangeWithScores(ctx, key, 0, int64(limit-1))
	if err != nil {
		outcome := OutcomeFromError(err)
		recordStageSince(r.observer, StageInbox, outcome, started)
		RecordDependency(r.observer, DependencyRedis, OperationInboxRead, outcome)
		return nil, fmt.Errorf("zrevrange failed: %w", err)
	}

	posts := make([]Post, 0, len(zs))
	for _, z := range zs {
		posts = append(posts, Post{
			ID:         z.Member,
			CreateTime: int64(z.Score),
		})
	}
	recordStageSince(r.observer, StageInbox, OutcomeSuccess, started)
	RecordDependency(r.observer, DependencyRedis, OperationInboxRead, OutcomeSuccess)
	RecordCursorWork(r.observer, CursorWorkRedisCommands, 1)
	RecordCursorWork(r.observer, CursorWorkRedisRoundTrips, 1)
	RecordCursorWork(r.observer, CursorWorkRedisMembers, len(zs))

	return posts, nil
}

// readInboxAndBigVs combines the inbox and the first bounded BigV batch into
// one Redis pipeline when the adapter supports it. The fallback preserves the
// legacy behavior for mocks and alternative Redis implementations.
func (r *FeedReader) readInboxAndBigVs(
	ctx context.Context,
	userID int64,
	bigVs []int64,
	limit int,
) (inboxPosts, bigVPosts []Post, inboxErr error) {
	batchReader, ok := r.redis.(redisBatchReader)
	if !r.combinedPipeline || !ok || len(bigVs) == 0 || limit <= 0 {
		inboxPosts, inboxErr = r.readInbox(ctx, userID, limit)
		if ctx.Err() != nil {
			return inboxPosts, []Post{}, inboxErr
		}
		return inboxPosts, r.pullFromBigVs(ctx, bigVs, limit), inboxErr
	}

	return r.readInboxAndBigVsBatch(ctx, batchReader, userID, bigVs, limit)
}

func (r *FeedReader) readInboxAndBigVsBatch(
	ctx context.Context,
	batchReader redisBatchReader,
	userID int64,
	bigVs []int64,
	limit int,
) ([]Post, []Post, error) {
	perBigV := limit
	if perBigV > BIGV_OUTBOX_MAX_SIZE {
		perBigV = BIGV_OUTBOX_MAX_SIZE
	}

	inboxPosts := []Post{}
	var inboxErr error
	topPosts := make(postMinHeap, 0, limit)
	heap.Init(&topPosts)
	nextCreator := 0
	firstBatch := true

	for firstBatch || nextCreator < len(bigVs) {
		if !firstBatch && ctx.Err() != nil {
			break
		}

		creatorCapacity := r.combinedPipelineBatchSize
		requests := make([]ZRevRangeRequest, 0, r.combinedPipelineBatchSize)
		if firstBatch {
			requests = append(requests, ZRevRangeRequest{
				Key:   fmt.Sprintf(FEED_INBOX_KEY, userID),
				Start: 0,
				Stop:  int64(limit - 1),
			})
			creatorCapacity--
		}

		end := nextCreator + creatorCapacity
		if end > len(bigVs) {
			end = len(bigVs)
		}
		creators := bigVs[nextCreator:end]
		for _, creatorID := range creators {
			requests = append(requests, ZRevRangeRequest{
				Key:   fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID),
				Start: 0,
				Stop:  int64(perBigV - 1),
			})
		}

		started := observationStarted(r.observer)
		results := batchReader.ZRevRangeWithScoresBatch(ctx, requests)
		RecordCursorWork(r.observer, CursorWorkRedisCommands, len(requests))
		RecordCursorWork(r.observer, CursorWorkRedisRoundTrips, 1)
		redisMembers := 0
		for _, result := range results {
			redisMembers += len(result.Scores)
		}
		RecordCursorWork(r.observer, CursorWorkRedisMembers, redisMembers)
		dependencyErr := firstZRevRangeError(results, len(requests))
		operation := OperationBigVPipeline
		if firstBatch {
			operation = OperationInboxBigVPipeline
		}
		RecordDependency(r.observer, DependencyRedis, operation, OutcomeFromError(dependencyErr))

		resultOffset := 0
		if firstBatch {
			inboxResult := zRevRangeResultAt(results, 0)
			inboxOutcome := OutcomeFromError(inboxResult.Err)
			recordStageSince(r.observer, StageInbox, inboxOutcome, started)
			if inboxResult.Err != nil {
				inboxErr = fmt.Errorf("zrevrange failed: %w", inboxResult.Err)
			} else {
				inboxPosts = postsFromScores(inboxResult.Scores, 0)
			}
			resultOffset = 1
		}

		bigVErr := firstZRevRangeErrorFrom(results, resultOffset, len(creators))
		recordStageSince(r.observer, StageBigVPipeline, OutcomeFromError(bigVErr), started)
		for i, creatorID := range creators {
			result := zRevRangeResultAt(results, resultOffset+i)
			if result.Err != nil {
				if ctx.Err() == nil {
					r.logger.Errorf("pull from bigv pipeline failed: err=%v", result.Err)
				}
				continue
			}
			for _, score := range result.Scores {
				pushTopPost(&topPosts, Post{
					ID:         score.Member,
					CreatorID:  creatorID,
					CreateTime: int64(score.Score),
				}, limit)
			}
		}

		nextCreator = end
		firstBatch = false
	}

	sort.Slice(topPosts, func(i, j int) bool {
		return postComesBefore(topPosts[i], topPosts[j])
	})
	return inboxPosts, topPosts, inboxErr
}

func postsFromScores(scores []ZScore, creatorID int64) []Post {
	posts := make([]Post, 0, len(scores))
	for _, score := range scores {
		posts = append(posts, Post{
			ID:         score.Member,
			CreatorID:  creatorID,
			CreateTime: int64(score.Score),
		})
	}
	return posts
}

func zRevRangeResultAt(results []ZRevRangeResult, index int) ZRevRangeResult {
	if index < len(results) {
		return results[index]
	}
	return ZRevRangeResult{Err: fmt.Errorf("redis pipeline returned %d results, missing index %d", len(results), index)}
}

func firstZRevRangeError(results []ZRevRangeResult, expected int) error {
	return firstZRevRangeErrorFrom(results, 0, expected)
}

func firstZRevRangeErrorFrom(results []ZRevRangeResult, offset, expected int) error {
	for i := 0; i < expected; i++ {
		if result := zRevRangeResultAt(results, offset+i); result.Err != nil {
			return result.Err
		}
	}
	return nil
}

func pushTopPost(topPosts *postMinHeap, post Post, limit int) {
	if topPosts.Len() < limit {
		heap.Push(topPosts, post)
		return
	}
	if postComesBefore(post, (*topPosts)[0]) {
		heap.Pop(topPosts)
		heap.Push(topPosts, post)
	}
}

// pullFromBigVs 实时拉取大V的最新内容（有界并行）
func (r *FeedReader) pullFromBigVs(ctx context.Context, bigVs []int64, limit int) []Post {
	if len(bigVs) == 0 || limit <= 0 {
		return []Post{}
	}

	// 每个大V都必须提供足够深度，才能形成正确的全局 Top-N。
	// 单个 outbox 最多保留 BIGV_OUTBOX_MAX_SIZE 条，因此读取量有明确上限。
	perBigV := limit
	if perBigV > BIGV_OUTBOX_MAX_SIZE {
		perBigV = BIGV_OUTBOX_MAX_SIZE
	}
	if batchReader, ok := r.redis.(redisBatchReader); ok {
		return r.pullFromBigVsBatch(ctx, batchReader, bigVs, perBigV, limit)
	}

	type result struct {
		posts []Post
		err   error
	}

	workerCount := len(bigVs)
	if workerCount > maxBigVPullConcurrency {
		workerCount = maxBigVPullConcurrency
	}
	jobs := make(chan int64)
	results := make(chan result, workerCount)
	var workers sync.WaitGroup
	workers.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer workers.Done()
			for creatorID := range jobs {
				posts, err := r.pullFromBigV(ctx, creatorID, perBigV)
				results <- result{posts: posts, err: err}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, creatorID := range bigVs {
			select {
			case jobs <- creatorID:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		workers.Wait()
		close(results)
	}()

	// 每个 worker 最多暂存一个 outbox；聚合侧只维护全局 Top-N。
	topPosts := make(postMinHeap, 0, limit)
	heap.Init(&topPosts)
	for res := range results {
		if res.err != nil {
			r.logger.Errorf("pull from bigv failed: %v", res.err)
			continue
		}
		for _, post := range res.posts {
			if topPosts.Len() < limit {
				heap.Push(&topPosts, post)
				continue
			}
			if postComesBefore(post, topPosts[0]) {
				heap.Pop(&topPosts)
				heap.Push(&topPosts, post)
			}
		}
	}

	sort.Slice(topPosts, func(i, j int) bool {
		return postComesBefore(topPosts[i], topPosts[j])
	})
	return topPosts
}

func (r *FeedReader) pullFromBigVsBatch(
	ctx context.Context,
	batchReader redisBatchReader,
	bigVs []int64,
	perBigV int,
	limit int,
) []Post {
	topPosts := make(postMinHeap, 0, limit)
	heap.Init(&topPosts)
	for start := 0; start < len(bigVs); start += maxBigVOutboxPipelineSize {
		if ctx.Err() != nil {
			break
		}
		end := start + maxBigVOutboxPipelineSize
		if end > len(bigVs) {
			end = len(bigVs)
		}
		creators := bigVs[start:end]
		requests := make([]ZRevRangeRequest, len(creators))
		for i, creatorID := range creators {
			requests[i] = ZRevRangeRequest{
				Key:   fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID),
				Start: 0,
				Stop:  int64(perBigV - 1),
			}
		}

		started := observationStarted(r.observer)
		results := batchReader.ZRevRangeWithScoresBatch(ctx, requests)
		RecordCursorWork(r.observer, CursorWorkRedisCommands, len(requests))
		RecordCursorWork(r.observer, CursorWorkRedisRoundTrips, 1)
		redisMembers := 0
		for _, result := range results {
			redisMembers += len(result.Scores)
		}
		RecordCursorWork(r.observer, CursorWorkRedisMembers, redisMembers)
		var batchErr error
		for _, result := range results {
			if result.Err != nil {
				batchErr = result.Err
				break
			}
		}
		outcome := OutcomeFromError(batchErr)
		recordStageSince(r.observer, StageBigVPipeline, outcome, started)
		RecordDependency(r.observer, DependencyRedis, OperationBigVPipeline, outcome)
		for i, result := range results {
			if i >= len(creators) {
				break
			}
			if result.Err != nil {
				r.logger.Errorf("pull from bigv failed: creator=%d, err=%v", creators[i], result.Err)
				continue
			}
			for _, score := range result.Scores {
				post := Post{
					ID:         score.Member,
					CreatorID:  creators[i],
					CreateTime: int64(score.Score),
				}
				if topPosts.Len() < limit {
					heap.Push(&topPosts, post)
					continue
				}
				if postComesBefore(post, topPosts[0]) {
					heap.Pop(&topPosts)
					heap.Push(&topPosts, post)
				}
			}
		}
	}

	sort.Slice(topPosts, func(i, j int) bool {
		return postComesBefore(topPosts[i], topPosts[j])
	})
	return topPosts
}

// pullFromBigV 从单个大V的发件箱拉取
func (r *FeedReader) pullFromBigV(ctx context.Context, creatorID int64, limit int) ([]Post, error) {
	key := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID)

	// 从 ZSet 读取（按时间倒序）
	started := observationStarted(r.observer)
	zs, err := r.redis.ZRevRangeWithScores(ctx, key, 0, int64(limit-1))
	if err != nil {
		outcome := OutcomeFromError(err)
		recordStageSince(r.observer, StageBigVPipeline, outcome, started)
		RecordDependency(r.observer, DependencyRedis, OperationBigVPipeline, outcome)
		return nil, fmt.Errorf("zrevrange bigv outbox failed: %w", err)
	}

	posts := make([]Post, 0, len(zs))
	for _, z := range zs {
		posts = append(posts, Post{
			ID:         z.Member,
			CreatorID:  creatorID,
			CreateTime: int64(z.Score),
		})
	}
	recordStageSince(r.observer, StageBigVPipeline, OutcomeSuccess, started)
	RecordDependency(r.observer, DependencyRedis, OperationBigVPipeline, OutcomeSuccess)
	RecordCursorWork(r.observer, CursorWorkRedisCommands, 1)
	RecordCursorWork(r.observer, CursorWorkRedisRoundTrips, 1)
	RecordCursorWork(r.observer, CursorWorkRedisMembers, len(zs))

	return posts, nil
}

// ===== 辅助函数（纯函数，容易测试）=====

// postMinHeap 把最旧的候选放在堆顶，便于在固定空间内维护全局 Top-N。
type postMinHeap []Post

func (h postMinHeap) Len() int { return len(h) }
func (h postMinHeap) Less(i, j int) bool {
	return postComesBefore(h[j], h[i])
}
func (h postMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *postMinHeap) Push(value interface{}) {
	*h = append(*h, value.(Post))
}
func (h *postMinHeap) Pop() interface{} {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

func postComesBefore(left, right Post) bool {
	if left.CreateTime == right.CreateTime {
		return left.ID > right.ID
	}
	return left.CreateTime > right.CreateTime
}

// MergePosts 归并两个帖子列表（按时间倒序）
func MergePosts(list1, list2 []Post) []Post {
	result := make([]Post, 0, len(list1)+len(list2))
	result = append(result, list1...)
	result = append(result, list2...)

	// 按时间戳降序排序
	sort.Slice(result, func(i, j int) bool {
		return postComesBefore(result[i], result[j])
	})

	return result
}

// Deduplicate 去重（保留第一次出现的）
func Deduplicate(posts []Post) []Post {
	seen := make(map[int64]bool)
	result := []Post{}

	for _, post := range posts {
		if !seen[post.ID] {
			seen[post.ID] = true
			result = append(result, post)
		}
	}

	return result
}

// BalanceFeed 平衡大V内容占比
// maxBigVRatio: 大V内容最多占比（0-1之间）
func BalanceFeed(posts []Post, maxBigVRatio float64) []Post {
	// 简化版本：直接返回，不做平衡
	// TODO: 需要知道哪些是大V的帖子才能平衡
	// 这需要传入大V列表或者在Post结构中标记
	return posts
}
