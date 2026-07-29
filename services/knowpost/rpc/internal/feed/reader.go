package feed

import (
	"context"
	"fmt"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	// 1. 获取关注列表
	followings, err := r.relationClient.GetFollowings(ctx, userID)
	if err != nil {
		r.logger.Errorf("get followings failed: user=%d, err=%v", userID, err)
		return nil, false, fmt.Errorf("get followings failed: %w", err)
	}

	if len(followings) == 0 {
		// 没有关注任何人，返回空
		r.logger.Infof("user %d has no followings", userID)
		return []int64{}, false, nil
	}

	// 2. 区分大V和普通用户
	bigVs, normalUsers := r.classifyFollowings(ctx, followings)
	r.logger.Infof("user %d followings: %d bigVs, %d normal users",
		userID, len(bigVs), len(normalUsers))

	// 3. 读收件箱（推模式的内容）
	inboxPosts, err := r.readInbox(ctx, userID, size*2) // 多读一些，用于去重后分页
	if err != nil {
		r.logger.Errorf("read inbox failed: %v", err)
		// 降级：继续处理，只是收件箱为空
		inboxPosts = []Post{}
	}

	// 4. 实时拉取大V内容
	bigVPosts := r.pullFromBigVs(ctx, bigVs, size)

	// 5. 归并、去重、排序
	allPosts := MergePosts(inboxPosts, bigVPosts)
	allPosts = Deduplicate(allPosts)

	// 6. 平衡大V内容占比
	allPosts = BalanceFeed(allPosts, 0.5) // 大V内容最多占50%

	// 7. 分页
	start := (page - 1) * size
	end := start + size
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

	for _, uid := range followings {
		// 从 Counter 获取粉丝数
		followerCount, err := r.counterClient.GetFollowerCount(ctx, uid)
		if err != nil {
			r.logger.Errorf("get follower count failed: user=%d, err=%v", uid, err)
			// 降级：当作普通用户处理
			normalUsers = append(normalUsers, uid)
			continue
		}

		if followerCount > BIGV_THRESHOLD {
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

// pullFromBigVs 实时拉取大V的最新内容（并行）
func (r *FeedReader) pullFromBigVs(ctx context.Context, bigVs []int64, limit int) []Post {
	if len(bigVs) == 0 {
		return []Post{}
	}

	// 计算每个大V拉取多少条
	perBigV := limit / len(bigVs)
	if perBigV == 0 {
		perBigV = 1
	}

	// 并行拉取
	type result struct {
		posts []Post
		err   error
	}

	ch := make(chan result, len(bigVs))

	for _, bigV := range bigVs {
		go func(creatorID int64) {
			posts, err := r.pullFromBigV(ctx, creatorID, perBigV)
			ch <- result{posts: posts, err: err}
		}(bigV)
	}

	// 收集结果
	allPosts := []Post{}
	for i := 0; i < len(bigVs); i++ {
		res := <-ch
		if res.err != nil {
			r.logger.Errorf("pull from bigv failed: %v", res.err)
			continue
		}
		allPosts = append(allPosts, res.posts...)
	}

	return allPosts
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

// MergePosts 归并两个帖子列表（按时间倒序）
func MergePosts(list1, list2 []Post) []Post {
	result := make([]Post, 0, len(list1)+len(list2))
	result = append(result, list1...)
	result = append(result, list2...)

	// 按时间戳降序排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreateTime > result[j].CreateTime
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
