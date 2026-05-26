// Package flusher 实现两条工作流：
//   - RunConsumer：消费 counter-events，累加到 Redis Hash 聚合桶
//   - RunFlusher：定时把聚合桶 flush 进 SDS
package flusher

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/event"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"

	kafka "github.com/segmentio/kafka-go"
)

// RunConsumer 启动 Kafka consumer，按 entityId 顺序累加到聚合桶。
func RunConsumer(ctx context.Context, sc *svc.ServiceContext) error {
	cfg := kafkax.ConsumerConfig{
		Brokers: sc.Config.Kafka.Brokers,
		Topic:   sc.Config.Kafka.Topic,
		GroupId: sc.Config.Kafka.GroupId,
	}
	return kafkax.RunConsumer(ctx, cfg, func(ctx context.Context, m kafka.Message) error {
		var ev event.CounterEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			logx.Errorf("decode event: %v body=%s", err, string(m.Value))
			return nil // 跳过坏消息
		}
		key := schema.AggKey(ev.EntityType, ev.EntityId)
		field := strconv.Itoa(ev.Idx)
		_, err := sc.Redis.HIncrBy(ctx, key, field, int64(ev.Delta)).Result()
		return err
	})
}

// RunFlusher 每隔 IntervalMs 扫描所有 agg:* key，把字段 delta 折算到 SDS 并清桶。
// 多副本环境下用 redsync 选主，只让一个 leader 真正执行 flushOnce，
// 其他副本静默放行直到下一个 tick。
func RunFlusher(ctx context.Context, sc *svc.ServiceContext) {
	interval := time.Duration(sc.Config.Flush.IntervalMs) * time.Millisecond
	lockTtl := time.Duration(sc.Config.Lock.TtlMs) * time.Millisecond
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			release, ok, err := sc.Locks.TryAcquire(ctx, sc.Config.Lock.Key, lockTtl)
			if err != nil {
				logx.Errorf("acquire flush lock: %v", err)
				continue
			}
			if !ok {
				// 不是当前 leader，跳过这次 tick
				continue
			}
			func() {
				defer func() {
					if err := release(); err != nil {
						logx.Errorf("release flush lock: %v", err)
					}
				}()
				if err := flushOnce(ctx, sc); err != nil {
					logx.Errorf("flush err: %v", err)
				}
			}()
		}
	}
}

// flushOnce 一次性扫并刷写。生产场景可加分布式锁避免多副本重复 flush。
func flushOnce(ctx context.Context, sc *svc.ServiceContext) error {
	pattern := fmt.Sprintf("agg:%s:*", schema.SchemaId)
	iter := sc.Redis.Scan(ctx, 0, pattern, int64(sc.Config.Flush.BatchSize)).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		if err := flushKey(ctx, sc, key); err != nil {
			logx.Errorf("flushKey %s err: %v", key, err)
		}
	}
	return iter.Err()
}

// flushKey 处理单个聚合 key：HGETALL → 对每字段 IncrField + DecrField。
func flushKey(ctx context.Context, sc *svc.ServiceContext, aggKey string) error {
	fields, err := sc.Redis.HGetAll(ctx, aggKey).Result()
	if err != nil || len(fields) == 0 {
		return err
	}
	etype, eid, ok := parseAggKey(aggKey)
	if !ok {
		return nil
	}
	sdsKey := schema.SdsKey(etype, eid)
	for fieldStr, deltaStr := range fields {
		idx, err := strconv.Atoi(fieldStr)
		if err != nil {
			continue
		}
		delta, err := strconv.ParseInt(deltaStr, 10, 64)
		if err != nil || delta == 0 {
			continue
		}
		// 1) 把 delta 折算到 SDS（IncrField 支持负数 delta）
		if _, err := sc.IncrFieldScript.Run(ctx, sc.Redis,
			[]string{sdsKey}, idx, delta, schema.SchemaLen, schema.FieldSize).Int64(); err != nil {
			return err
		}
		// 2) 从聚合桶扣减相同 delta（绝对值）；脚本内做 -delta
		absDelta := delta
		if absDelta < 0 {
			absDelta = -absDelta
		}
		if _, err := sc.DecrFieldScript.Run(ctx, sc.Redis,
			[]string{aggKey}, fieldStr, absDelta).Int64(); err != nil {
			return err
		}
	}
	return nil
}

// parseAggKey "agg:v1:{etype}:{eid}" → etype, eid。
func parseAggKey(s string) (etype, eid string, ok bool) {
	const prefix = "agg:" + schema.SchemaId + ":"
	if len(s) <= len(prefix) || s[:len(prefix)] != prefix {
		return "", "", false
	}
	rest := s[len(prefix):]
	for i := 0; i < len(rest); i++ {
		if rest[i] == ':' {
			return rest[:i], rest[i+1:], true
		}
	}
	return "", "", false
}
