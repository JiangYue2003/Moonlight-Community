package feed

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisAdapter 实现 RedisClient 接口，适配 go-redis
type RedisAdapter struct {
	client redis.UniversalClient
}

func NewRedisAdapter(client redis.UniversalClient) *RedisAdapter {
	return &RedisAdapter{client: client}
}

func (r *RedisAdapter) ZAdd(ctx context.Context, key string, members ...interface{}) error {
	// members 格式：score1, member1, score2, member2...
	// 转换为 redis.Z 结构
	zs := make([]redis.Z, 0, len(members)/2)
	for i := 0; i < len(members); i += 2 {
		score, ok1 := members[i].(float64)
		member, ok2 := members[i+1].(int64)
		if !ok1 || !ok2 {
			continue
		}
		zs = append(zs, redis.Z{
			Score:  score,
			Member: member,
		})
	}
	return r.client.ZAdd(ctx, key, zs...).Err()
}

func (r *RedisAdapter) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) error {
	return r.client.ZRemRangeByRank(ctx, key, start, stop).Err()
}

func (r *RedisAdapter) Expire(ctx context.Context, key string, seconds int) error {
	return r.client.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()
}

func (r *RedisAdapter) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	return n > 0, err
}

func (r *RedisAdapter) SetEx(ctx context.Context, key string, value interface{}, seconds int) error {
	return r.client.SetEx(ctx, key, value, time.Duration(seconds)*time.Second).Err()
}

func (r *RedisAdapter) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZScore, error) {
	zs, err := r.client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	result := make([]ZScore, len(zs))
	for i, z := range zs {
		memberID, ok := z.Member.(string)
		if !ok {
			continue
		}
		id, err := strconv.ParseInt(memberID, 10, 64)
		if err != nil {
			continue
		}
		result[i] = ZScore{
			Member: id,
			Score:  z.Score,
		}
	}
	return result, nil
}

func (r *RedisAdapter) SCard(ctx context.Context, key string) (int64, error) {
	return r.client.SCard(ctx, key).Result()
}
