// Package zset relation-rpc 的列表 ZSet 缓存。
//
// Key 模板：
//
//	uf:flws:{userId}    score=created_at_ms  member=toUserId（关注）
//	uf:fans:{userId}    score=created_at_ms  member=fromUserId（粉丝）
//
// 阶段3 仅"读"：当 ZSet 命中时返回；缺失时 logic 层走 DB 并返回结果（回填由 syncer 在写路径完成）。
// 这样保证 rpc 写路径不直接维护 ZSet，避免双写不一致。
package zset

import (
	"context"
	"fmt"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
)

const (
	FollowingPrefix = "uf:flws:"
	FollowerPrefix  = "uf:fans:"
)

func FollowingKey(userId int64) string { return fmt.Sprintf("%s%d", FollowingPrefix, userId) }
func FollowerKey(userId int64) string  { return fmt.Sprintf("%s%d", FollowerPrefix, userId) }

// PageByOffset 用 ZREVRANGE 取倒序段。返回 (id 列表, 是否命中缓存)。
//
// 阶段3 简化语义：ZSet 中存的元素数量是该用户当前的"最近 N 条"（N 由 syncer 控制 ≤ 2000）。
// 当 offset+limit 落在 ZSet 范围内则命中；超出则视为 miss，回 DB。
func PageByOffset(ctx context.Context, rdb goredis.UniversalClient, key string, offset, limit int) ([]int64, bool, error) {
	card, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		return nil, false, err
	}
	if card == 0 {
		return nil, false, nil
	}
	if int64(offset+limit) > card {
		// 超出 ZSet 已知范围 → miss，让上层走 DB
		return nil, false, nil
	}
	members, err := rdb.ZRevRange(ctx, key, int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		return nil, false, err
	}
	return parseIds(members), true, nil
}

// PageByCursor 用 ZREVRANGEBYSCORE：score < cursor，倒序取 limit 条。
func PageByCursor(ctx context.Context, rdb goredis.UniversalClient, key string, cursor int64, limit int) ([]int64, int64, bool, error) {
	if cursor <= 0 {
		return nil, 0, false, nil
	}
	res, err := rdb.ZRevRangeByScoreWithScores(ctx, key, &goredis.ZRangeBy{
		Min:    "-inf",
		Max:    "(" + strconv.FormatInt(cursor, 10),
		Offset: 0,
		Count:  int64(limit),
	}).Result()
	if err != nil {
		return nil, 0, false, err
	}
	if len(res) == 0 {
		return nil, 0, false, nil
	}
	ids := make([]int64, 0, len(res))
	var nextCursor int64
	for _, z := range res {
		v, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		ids = append(ids, v)
		nextCursor = int64(z.Score) // 最小的 score 即下一页起点（exclusive）
	}
	return ids, nextCursor, true, nil
}

func parseIds(members []string) []int64 {
	out := make([]int64, 0, len(members))
	for _, m := range members {
		v, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			continue
		}
		out = append(out, v)
	}
	return out
}
