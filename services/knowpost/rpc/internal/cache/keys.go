// Package cache 集中放置 knowpost-rpc 的 L1/L2 缓存 key 模板与 TTL 计算。
//
// 与原 Java KnowPostServiceImpl / KnowPostFeedServiceImpl 保持 key 命名一致，
// 便于阶段 N 灰度时双跑双写不冲突。
package cache

import (
	"fmt"
	"time"
)

const (
	// LayoutVer 用于在不破坏向后兼容的前提下回收所有缓存：版本号 +1 即可彻底失效旧条目。
	DetailLayoutVer = 1
	FeedLayoutVer   = 1

	// L2 基础 TTL（业务侧加抖动）。
	DetailBaseTTL     = 60 * time.Second
	FeedPublicBaseTTL = 60 * time.Second
	FeedItemBaseTTL   = 60 * time.Second
	FeedMineBaseTTL   = 30 * time.Second
	HasMoreBaseTTL    = 10 * time.Second

	// 抖动上限。
	DetailJitterMax     = 30 * time.Second
	FeedPublicJitterMax = 30 * time.Second
	FeedItemJitterMax   = 30 * time.Second
	FeedMineJitterMax   = 20 * time.Second
	HasMoreJitterMax    = 11 * time.Second

	NullTTL       = 30 * time.Second
	NullJitterMax = 30 * time.Second
)

// DetailKey 帖子详情：knowpost:detail:{id}:v{ver}
func DetailKey(id int64) string {
	return fmt.Sprintf("knowpost:detail:%d:v%d", id, DetailLayoutVer)
}

// FeedPublicL1Key 公共 Feed 的 L1（整页）key：feed:public:{size}:{page}:v{ver}
func FeedPublicL1Key(size, page int) string {
	return fmt.Sprintf("feed:public:%d:%d:v%d", size, page, FeedLayoutVer)
}

// FeedPublicPageKey 公共 Feed 的 Redis 整页缓存 key，与 Java 结构回退保持一致。
func FeedPublicPageKey(size, page int) string {
	return FeedPublicL1Key(size, page)
}

// FeedPublicIdsKey 公共 Feed 的 L2 ids 列表 key：feed:public:ids:{size}:{hourSlot}:{page}
func FeedPublicIdsKey(size, page int, hourSlot int64) string {
	return fmt.Sprintf("feed:public:ids:%d:%d:%d", size, hourSlot, page)
}

// FeedPublicHasMoreKey idsKey + ":hasMore"
func FeedPublicHasMoreKey(idsKey string) string {
	return idsKey + ":hasMore"
}

// FeedItemKey feed 流单条 item key：feed:item:{id}（跨页共享）
func FeedItemKey(id int64) string {
	return fmt.Sprintf("feed:item:%d", id)
}

// FeedReverseIndexKey 反向索引：feed:public:index:{eid}:{hourSlot} → set of pageKey
func FeedReverseIndexKey(eid int64, hourSlot int64) string {
	return fmt.Sprintf("feed:public:index:%d:%d", eid, hourSlot)
}

// FeedAllPagesKey GC 用：所有曾被填充的 pageKey 集合
const FeedAllPagesKey = "feed:public:pages"

// FeedMineKey 我的 Feed：feed:mine:{userId}:{size}:{page}
func FeedMineKey(userId int64, size, page int) string {
	return fmt.Sprintf("feed:mine:%d:%d:%d", userId, size, page)
}

// HourSlot 当前 UTC 小时槽（与 Java FeedServiceImpl 对齐）。
func HourSlot(now time.Time) int64 {
	return now.UTC().Unix() / 3600
}
