package knowpostlogic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	"github.com/zhiguang/zhiguang-go/pkg/sfx"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

// publicFeedSF 进程内单飞，防止同 (size,page,hour) 并发回源 DB。
var publicFeedSF = sfx.New[*knowpost.FeedPage]()

type GetPublicFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublicFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicFeedLogic {
	return &GetPublicFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPublicFeed 公开 feed 流：L1 整页 + Redis 整页 + 片段辅助缓存。
//
// 阶段2 实现策略：
//  1. L1FeedPublic 命中整页 → 直接返回（最快路径）
//  2. Redis pageKey 命中整页 → 回填 L1 → 返回
//  3. L2 ids list + MGET feed:item:{id} 命中 → 重建 FeedPage，写回 Redis pageKey + L1 → 返回
//  4. 都 miss：sf 单飞 → DB 取 size+1 行 → 写 pageKey、ids list、单条 item、反向索引 → 写 L1 → 返回
func (l *GetPublicFeedLogic) GetPublicFeed(in *knowpost.GetPublicFeedReq) (*knowpost.FeedPage, error) {
	page, size := normalizePage(in.Page, in.Size)
	hourSlot := cachekeys.HourSlot(time.Now())
	pageKey := cachekeys.FeedPublicL1Key(size, page)
	redisPageKey := cachekeys.FeedPublicPageKey(size, page)
	idsKey := cachekeys.FeedPublicIdsKey(size, page, hourSlot)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	sfKey := pageKey

	// 1. L1 整页缓存
	if v, ok := l.svcCtx.L1FeedPublic.Get(pageKey); ok {
		if raw, ok := v.([]byte); ok {
			var page knowpost.FeedPage
			if err := json.Unmarshal(raw, &page); err == nil {
				l.recordHotAfterHit(redisPageKey, &page)
				return &page, nil
			}
		}
	}

	// 2. Redis 整页缓存
	if p, ok := l.tryReadFromRedisPage(redisPageKey); ok {
		l.cachePageL1(pageKey, p)
		l.recordHotAfterHit(redisPageKey, p)
		return p, nil
	}

	// 3. L2 ids + items 重建
	if p, ok := l.tryReadFromL2(idsKey, hasMoreKey, page, size); ok {
		l.writeRedisPage(redisPageKey, p)
		l.cachePageL1(pageKey, p)
		l.recordHotAfterHit(redisPageKey, p)
		return p, nil
	}

	// 4. 单飞回源
	resp, err := publicFeedSF.DoCtx(l.ctx, sfKey, func(ctx context.Context) (*knowpost.FeedPage, error) {
		if p, ok := l.tryReadFromRedisPage(redisPageKey); ok {
			return p, nil
		}
		// 双检 L2，避免在等单飞期间已经被其他 goroutine 回填
		if p, ok := l.tryReadFromL2(idsKey, hasMoreKey, page, size); ok {
			l.writeRedisPage(redisPageKey, p)
			return p, nil
		}
		rows, err := l.svcCtx.KnowPostsModel.ListPublicFeed(ctx, size+1, (page-1)*size)
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
		pageResp := &knowpost.FeedPage{
			Items: items, HasMore: hasMore,
			Size: int32(size), Page: int32(page),
		}
		// 写 L2：pageKey + ids list + 单条 item + has_more + 反向索引
		l.writeBackL2(ctx, redisPageKey, idsKey, hasMoreKey, hourSlot, items, hasMore, pageResp)
		return pageResp, nil
	})
	if err != nil {
		return nil, err
	}
	l.cachePageL1(pageKey, resp)
	l.recordHotAfterHit(redisPageKey, resp)
	return resp, nil
}

func (l *GetPublicFeedLogic) tryReadFromRedisPage(pageKey string) (*knowpost.FeedPage, bool) {
	raw, err := l.svcCtx.Redis.Get(l.ctx, pageKey).Bytes()
	if err != nil || len(raw) == 0 {
		return nil, false
	}
	var page knowpost.FeedPage
	if err := json.Unmarshal(raw, &page); err != nil {
		return nil, false
	}
	return &page, true
}

// tryReadFromL2 尝试从 ids list + item String 拼出 FeedPage；任何片段缺失即视为 miss。
func (l *GetPublicFeedLogic) tryReadFromL2(idsKey, hasMoreKey string, page, size int) (*knowpost.FeedPage, bool) {
	ctx := l.ctx
	idStrs, err := l.svcCtx.Redis.LRange(ctx, idsKey, 0, int64(size-1)).Result()
	if err != nil || len(idStrs) == 0 {
		return nil, false
	}
	hasMoreVal, err := l.svcCtx.Redis.Get(ctx, hasMoreKey).Result()
	if err != nil {
		return nil, false
	}
	itemKeys := make([]string, 0, len(idStrs))
	for _, s := range idStrs {
		uid, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, false
		}
		itemKeys = append(itemKeys, cachekeys.FeedItemKey(uid))
	}
	raws, err := l.svcCtx.Redis.MGet(ctx, itemKeys...).Result()
	if err != nil {
		return nil, false
	}
	items := make([]*knowpost.FeedItem, 0, len(raws))
	for _, r := range raws {
		s, ok := r.(string)
		if !ok || s == "" {
			return nil, false // 任意 item 缺失 → miss
		}
		var it knowpost.FeedItem
		if err := json.Unmarshal([]byte(s), &it); err != nil {
			return nil, false
		}
		items = append(items, &it)
	}
	return &knowpost.FeedPage{
		Items: items, HasMore: hasMoreVal == "1",
		Size: int32(size), Page: int32(page),
	}, true
}

// writeBackL2 把回源结果落到 Redis 多 key 片段。
func (l *GetPublicFeedLogic) writeBackL2(ctx context.Context, pageKey, idsKey, hasMoreKey string,
	hourSlot int64, items []*knowpost.FeedItem, hasMore bool, page *knowpost.FeedPage) {
	if len(items) == 0 {
		return
	}
	pipe := l.svcCtx.Redis.Pipeline()
	pageRaw, _ := json.Marshal(page)
	pageTTL := cachex.Jitter(cachekeys.FeedPublicBaseTTL, cachekeys.FeedPublicJitterMax)
	pipe.Set(ctx, pageKey, pageRaw, pageTTL)

	// ids list
	idVals := make([]any, 0, len(items))
	for _, it := range items {
		idVals = append(idVals, it.Id)
	}
	pipe.Del(ctx, idsKey)
	pipe.RPush(ctx, idsKey, idVals...)
	pipe.Expire(ctx, idsKey, cachex.Jitter(cachekeys.FeedPublicBaseTTL, cachekeys.FeedPublicJitterMax))

	// has_more
	hm := "0"
	if hasMore {
		hm = "1"
	}
	pipe.Set(ctx, hasMoreKey, hm, cachex.Jitter(cachekeys.HasMoreBaseTTL, cachekeys.HasMoreJitterMax))

	// 单条 item
	for _, it := range items {
		raw, _ := json.Marshal(it)
		pipe.Set(ctx, cachekeys.FeedItemKey(parseInt64(it.Id)), raw,
			cachex.Jitter(cachekeys.FeedItemBaseTTL, cachekeys.FeedItemJitterMax))
		// 反向索引：feed:public:index:{eid}:{hour} → 包含此 item 的 pageKey
		pipe.SAdd(ctx, cachekeys.FeedReverseIndexKey(parseInt64(it.Id), hourSlot), pageKey)
		pipe.Expire(ctx, cachekeys.FeedReverseIndexKey(parseInt64(it.Id), hourSlot),
			2*cachekeys.FeedPublicBaseTTL)
	}
	pipe.SAdd(ctx, cachekeys.FeedAllPagesKey, pageKey)
	if _, err := pipe.Exec(ctx); err != nil {
		l.Logger.Errorf("public feed writeback: %v", err)
	}
}

func (l *GetPublicFeedLogic) writeRedisPage(pageKey string, p *knowpost.FeedPage) {
	raw, err := json.Marshal(p)
	if err != nil {
		return
	}
	_ = l.svcCtx.Redis.Set(l.ctx, pageKey, raw,
		cachex.Jitter(cachekeys.FeedPublicBaseTTL, cachekeys.FeedPublicJitterMax)).Err()
}

func (l *GetPublicFeedLogic) cachePageL1(pageKey string, p *knowpost.FeedPage) {
	raw, err := json.Marshal(p)
	if err != nil {
		return
	}
	l.svcCtx.L1FeedPublic.SetWithTTL(pageKey, raw, int64(len(raw)),
		cachex.Jitter(cachekeys.FeedPublicBaseTTL, cachekeys.FeedPublicJitterMax))
}

func (l *GetPublicFeedLogic) recordHotAfterHit(pageKey string, p *knowpost.FeedPage) {
	l.svcCtx.HotFeedPublic.Hit(pageKey)
	if level := l.svcCtx.HotFeedPublic.Level(pageKey); level > hotkey.LevelNone {
		ext := hotkey.TTLForPublic(cachekeys.FeedPublicBaseTTL, level)
		_ = l.svcCtx.Redis.Expire(l.ctx, pageKey, ext).Err()
	}
	for _, it := range p.Items {
		key := cachekeys.FeedItemKey(parseInt64(it.Id))
		l.svcCtx.HotFeedItem.Hit(key)
		if level := l.svcCtx.HotFeedItem.Level(key); level > hotkey.LevelNone {
			ext := hotkey.TTLForPublic(cachekeys.FeedItemBaseTTL, level)
			_ = l.svcCtx.Redis.Expire(l.ctx, key, ext).Err()
		}
	}
}

func parseInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

// 让编译器看到 goredis 类型用法
var _ = goredis.Nil
