package feed

import (
	"container/heap"
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxBigVPullConcurrency = 16

// Post 帖子基本信息（用于排序和去重）
type Post struct {
	ID         int64
	CreatorID  int64
	CreateTime int64
}

// FeedReader 负责读取 Feed 流（推拉混合）
type FeedReader struct {
	redis          RedisClient
	relationClient RelationClient
	counterClient  CounterClient
	logger         logx.Logger
}

// FeedReadSnapshot 固化一次请求中的关注关系和大V分类，供过滤回填复用。
type FeedReadSnapshot struct {
	reader     *FeedReader
	userID     int64
	bigVs      []int64
	followings map[int64]struct{}
}

func NewFeedReader(redis RedisClient, relationClient RelationClient, counterClient CounterClient, logger logx.Logger) *FeedReader {
	return &FeedReader{
		redis:          redis,
		relationClient: relationClient,
		counterClient:  counterClient,
		logger:         logger,
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
	followings, err := r.relationClient.GetFollowings(ctx, userID)
	if err != nil {
		r.logger.Errorf("get followings failed: user=%d, err=%v", userID, err)
		return nil, fmt.Errorf("get followings failed: %w", err)
	}

	bigVs, normalUsers := r.classifyFollowings(ctx, followings)
	r.logger.Infof("user %d followings: %d bigVs, %d normal users",
		userID, len(bigVs), len(normalUsers))
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
	r := s.reader

	start := (page - 1) * size
	end := start + size
	candidateLimit := end + 1
	inboxPosts, err := r.readInbox(ctx, s.userID, candidateLimit)
	if err != nil {
		r.logger.Errorf("read inbox failed: %v", err)
		// 降级：继续处理，只是收件箱为空
		inboxPosts = []Post{}
	}

	bigVPosts := r.pullFromBigVs(ctx, s.bigVs, candidateLimit)

	// 5. 归并、去重、排序
	allPosts := MergePosts(inboxPosts, bigVPosts)
	allPosts = Deduplicate(allPosts)

	// 6. 平衡大V内容占比
	allPosts = BalanceFeed(allPosts, 0.5) // 大V内容最多占50%

	// 7. 分页
	hasMore := len(allPosts) > end

	if start >= len(allPosts) {
		return []int64{}, false, nil
	}

	if end > len(allPosts) {
		end = len(allPosts)
	}

	result := make([]int64, 0, end-start)
	for i := start; i < end; i++ {
		result = append(result, allPosts[i].ID)
	}

	return result, hasMore, nil
}

// classifyFollowings 区分大V和普通用户
func (r *FeedReader) classifyFollowings(ctx context.Context, followings []int64) (bigVs, normalUsers []int64) {
	bigVs = []int64{}
	normalUsers = []int64{}
	if len(followings) == 0 {
		return bigVs, normalUsers
	}

	followerCounts, err := r.counterClient.BatchGetFollowerCounts(ctx, followings)
	if err != nil {
		r.logger.Errorf("batch get follower counts failed: users=%d, err=%v", len(followings), err)
		return bigVs, append(normalUsers, followings...)
	}

	for _, uid := range followings {
		if followerCounts[uid] > BIGV_THRESHOLD {
			bigVs = append(bigVs, uid)
		} else {
			normalUsers = append(normalUsers, uid)
		}
	}

	return
}

// readInbox 读取用户的收件箱
func (r *FeedReader) readInbox(ctx context.Context, userID int64, limit int) ([]Post, error) {
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)

	// 从 ZSet 读取（按时间倒序）
	zs, err := r.redis.ZRevRangeWithScores(ctx, key, 0, int64(limit-1))
	if err != nil {
		return nil, fmt.Errorf("zrevrange failed: %w", err)
	}

	posts := make([]Post, 0, len(zs))
	for _, z := range zs {
		posts = append(posts, Post{
			ID:         z.Member,
			CreateTime: int64(z.Score),
		})
	}

	return posts, nil
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

// pullFromBigV 从单个大V的发件箱拉取
func (r *FeedReader) pullFromBigV(ctx context.Context, creatorID int64, limit int) ([]Post, error) {
	key := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID)

	// 从 ZSet 读取（按时间倒序）
	zs, err := r.redis.ZRevRangeWithScores(ctx, key, 0, int64(limit-1))
	if err != nil {
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
