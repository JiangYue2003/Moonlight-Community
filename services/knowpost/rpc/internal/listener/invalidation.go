// Package listener 在 knowpost-rpc 进程内启动 Kafka consumer，
// 监听 counter-events 并失效相关缓存（详情、feed:item、反向索引列出的 feed 页）。
package listener

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/event"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

// Run 启动消费 goroutine，直到 ctx 被取消。
func Run(ctx context.Context, sc *svc.ServiceContext) error {
	cfg := kafkax.ConsumerConfig{
		Brokers: sc.Config.Kafka.Brokers,
		Topic:   sc.Config.Kafka.CounterEventsTopic,
		GroupId: sc.Config.Kafka.GroupId,
	}
	return kafkax.RunConsumer(ctx, cfg, func(ctx context.Context, m kafka.Message) error {
		var ev event.CounterEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			logx.Errorf("decode counter-event: %v body=%s", err, string(m.Value))
			return nil // skip 坏消息
		}
		handleCounterEvent(ctx, sc, ev)
		return nil
	})
}

func handleCounterEvent(ctx context.Context, sc *svc.ServiceContext, ev event.CounterEvent) {
	if ev.EntityType != "knowpost" || !isTargetCounterMetric(ev.Metric) {
		return
	}
	eid, err := strconv.ParseInt(ev.EntityId, 10, 64)
	if err != nil || eid <= 0 {
		return
	}

	// 详情缓存在 API 层拼装计数前可能被复用，沿用失效策略确保一致性。
	if sc.DetailCache != nil {
		_ = sc.DetailCache.Invalidate(ctx, cachekeys.DetailKey(eid))
	}

	// 计数事件优先做“页缓存修正”：成功则保留 Redis/L1 页缓存；失败才回退删除。
	patchPublicFeedPagesForCounterEvent(ctx, sc, eid)

	// usercounter 同步：作者获赞数累加（fav 字段暂未接入 usercounter schema）。
	if ev.Delta != 0 && ev.Metric == "like" && sc.KnowPostsModel != nil && sc.UserCounterRpc != nil {
		row, err := sc.KnowPostsModel.FindOne(ctx, uint64(eid))
		if err == nil && row != nil {
			if _, err := sc.UserCounterRpc.UserIncrement(ctx, &counterpb.UserIncrementReq{
				UserId: int64(row.CreatorId),
				Field:  "likes_received",
				Delta:  ev.Delta,
			}); err != nil {
				logx.Errorf("usercounter increment likes_received: %v", err)
			}
		}
	}
}

func isTargetCounterMetric(metric string) bool {
	return metric == "like" || metric == "fav"
}

func patchPublicFeedPagesForCounterEvent(ctx context.Context, sc *svc.ServiceContext, eid int64) {
	if sc == nil || sc.Redis == nil || sc.L1FeedPublic == nil {
		return
	}

	curr := cachekeys.HourSlot(time.Now())
	for _, hour := range []int64{curr, curr - 1} {
		ridx := cachekeys.FeedReverseIndexKey(eid, hour)
		pageKeys, err := sc.Redis.SMembers(ctx, ridx).Result()
		if err != nil || len(pageKeys) == 0 {
			continue
		}
		for _, pageKey := range pageKeys {
			if refreshPublicFeedPageCaches(ctx, sc, pageKey) {
				continue
			}
			invalidatePublicFeedPageByKey(ctx, sc, ridx, pageKey)
		}
	}
}

func refreshPublicFeedPageCaches(ctx context.Context, sc *svc.ServiceContext, pageKey string) bool {
	size, page, ok := parseFeedPublicPageKey(pageKey)
	if !ok {
		return false
	}
	idsKey, hasMoreKey, ok := loadFeedPublicFragmentsByPage(ctx, sc, size, page)
	if !ok {
		return false
	}

	idStrs, err := sc.Redis.LRange(ctx, idsKey, 0, int64(size-1)).Result()
	if err != nil || len(idStrs) == 0 {
		return false
	}
	hasMoreVal, err := sc.Redis.Get(ctx, hasMoreKey).Result()
	if err != nil {
		return false
	}

	itemKeys := make([]string, 0, len(idStrs))
	for _, idStr := range idStrs {
		pid, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || pid <= 0 {
			return false
		}
		itemKeys = append(itemKeys, cachekeys.FeedItemKey(pid))
	}
	rawItems, err := sc.Redis.MGet(ctx, itemKeys...).Result()
	if err != nil {
		return false
	}

	items := make([]*pb.FeedItem, 0, len(rawItems))
	for _, raw := range rawItems {
		s, ok := raw.(string)
		if !ok || s == "" {
			return false
		}
		var it pb.FeedItem
		if err := json.Unmarshal([]byte(s), &it); err != nil {
			return false
		}
		items = append(items, &it)
	}

	pagePayload := pb.FeedPage{
		Items:   items,
		HasMore: hasMoreVal == "1",
		Size:    int32(size),
		Page:    int32(page),
	}
	rawPage, err := json.Marshal(&pagePayload)
	if err != nil {
		return false
	}

	ttl := cachex.Jitter(cachekeys.FeedPublicBaseTTL, cachekeys.FeedPublicJitterMax)
	if redisTTL, err := sc.Redis.TTL(ctx, pageKey).Result(); err == nil && redisTTL > 0 {
		ttl = redisTTL
	}
	_ = sc.Redis.Set(ctx, pageKey, rawPage, ttl).Err()
	sc.L1FeedPublic.SetWithTTL(cachekeys.FeedPublicL1Key(size, page), rawPage, int64(len(rawPage)), ttl)
	return true
}

func invalidatePublicFeedPageByKey(ctx context.Context, sc *svc.ServiceContext, ridx, pageKey string) {
	if sc == nil || sc.Redis == nil {
		return
	}
	_ = sc.Redis.Del(ctx, pageKey).Err()
	_ = sc.Redis.SRem(ctx, cachekeys.FeedAllPagesKey, pageKey).Err()
	_ = sc.Redis.SRem(ctx, ridx, pageKey).Err()

	if sc.L1FeedPublic != nil {
		if size, page, ok := parseFeedPublicPageKey(pageKey); ok {
			sc.L1FeedPublic.Del(cachekeys.FeedPublicL1Key(size, page))
			for _, hour := range []int64{cachekeys.HourSlot(time.Now()), cachekeys.HourSlot(time.Now()) - 1} {
				idsKey := cachekeys.FeedPublicIdsKey(size, page, hour)
				_ = sc.Redis.Del(ctx, idsKey, cachekeys.FeedPublicHasMoreKey(idsKey)).Err()
			}
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

func loadFeedPublicFragmentsByPage(ctx context.Context, sc *svc.ServiceContext, size, page int) (idsKey, hasMoreKey string, ok bool) {
	curr := cachekeys.HourSlot(time.Now())
	for _, hour := range []int64{curr, curr - 1} {
		idsKey = cachekeys.FeedPublicIdsKey(size, page, hour)
		hasMoreKey = cachekeys.FeedPublicHasMoreKey(idsKey)
		if sc.Redis.Exists(ctx, idsKey).Val() > 0 {
			return idsKey, hasMoreKey, true
		}
	}
	return "", "", false
}
