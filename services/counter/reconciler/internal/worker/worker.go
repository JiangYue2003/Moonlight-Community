package worker

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/counter/reconciler/internal/config"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/sds"
)

// Worker 定时扫描 SDS key，与位图事实层对比，偏差超阈值时重建 SDS。
type Worker struct {
	cfg config.Config
	rdb goredis.UniversalClient
}

func New(cfg config.Config, rdb goredis.UniversalClient) *Worker {
	return &Worker{cfg: cfg, rdb: rdb}
}

func (w *Worker) Run(ctx context.Context) {
	interval := time.Duration(w.cfg.Scan.IntervalHours) * time.Hour
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := w.reconcile(ctx); err != nil {
				logx.Errorf("reconciler: reconcile error: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) reconcile(ctx context.Context) error {
	pattern := fmt.Sprintf("cnt:%s:*", schema.SchemaId)
	var cursor uint64
	batchInterval := time.Duration(w.cfg.Scan.BatchIntervalMs) * time.Millisecond

	for {
		keys, next, err := w.rdb.Scan(ctx, cursor, pattern, int64(w.cfg.Scan.BatchSize)).Result()
		if err != nil {
			return err
		}
		for _, key := range keys {
			if err := w.reconcileKey(ctx, key); err != nil {
				logx.Errorf("reconciler: key=%s err=%v", key, err)
			}
		}
		cursor = next
		if cursor == 0 {
			break
		}
		if batchInterval > 0 {
			select {
			case <-time.After(batchInterval):
			case <-ctx.Done():
				return nil
			}
		}
	}
	return nil
}

// reconcileKey 对单个 SDS key 做对账。
// key 格式：cnt:v1:{etype}:{eid}
func (w *Worker) reconcileKey(ctx context.Context, key string) error {
	parts := strings.SplitN(key, ":", 4)
	if len(parts) != 4 {
		return nil
	}
	etype, eid := parts[2], parts[3]

	raw, err := w.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return nil
		}
		return err
	}
	sdsVals := sds.Decode(raw)

	for _, metric := range schema.Supported {
		idx := schema.IdxOf(metric)
		if idx < 0 {
			continue
		}
		bitmapTotal, err := w.scanBitmapTotal(ctx, metric, etype, eid)
		if err != nil {
			return err
		}
		sdsVal := sdsVals[idx]
		diff := sdsVal - bitmapTotal
		if diff < 0 {
			diff = -diff
		}
		pct := float64(diff) / math.Max(float64(sdsVal), 1) * 100

		needRebuild := diff > w.cfg.Scan.ThresholdAbsolute || pct > w.cfg.Scan.ThresholdPercent
		logx.Infof("reconcile: etype=%s eid=%s metric=%s sds=%d bitmap=%d diff=%d pct=%.2f%% rebuild=%v",
			etype, eid, metric, sdsVal, bitmapTotal, diff, pct, needRebuild)

		if needRebuild {
			if err := w.rebuildField(ctx, key, idx, bitmapTotal); err != nil {
				logx.Errorf("reconciler: rebuild failed key=%s metric=%s: %v", key, metric, err)
			}
		}
	}
	return nil
}

// scanBitmapTotal 枚举 bm:{metric}:{etype}:{eid}:* 所有分片，pipeline BITCOUNT 求和。
func (w *Worker) scanBitmapTotal(ctx context.Context, metric, etype, eid string) (int64, error) {
	pattern := fmt.Sprintf("bm:%s:%s:%s:*", metric, etype, eid)
	var cursor uint64
	var total int64
	for {
		keys, next, err := w.rdb.Scan(ctx, cursor, pattern, 256).Result()
		if err != nil {
			return 0, err
		}
		if len(keys) > 0 {
			pipe := w.rdb.Pipeline()
			cmds := make([]*goredis.IntCmd, len(keys))
			for i, k := range keys {
				cmds[i] = pipe.BitCount(ctx, k, nil)
			}
			if _, err := pipe.Exec(ctx); err != nil {
				return 0, err
			}
			for _, c := range cmds {
				total += c.Val()
			}
		}
		cursor = next
		if cursor == 0 {
			break
		}
	}
	return total, nil
}

// rebuildField 用 bitmap 真实值覆盖 SDS 中的单个字段（绝对值替换）。
func (w *Worker) rebuildField(ctx context.Context, sdsKey string, idx int, val int64) error {
	raw, err := w.rdb.Get(ctx, sdsKey).Bytes()
	if err != nil && err != goredis.Nil {
		return err
	}
	if len(raw) < schema.SchemaLen*schema.FieldSize {
		buf := make([]byte, schema.SchemaLen*schema.FieldSize)
		copy(buf, raw)
		raw = buf
	}
	binary.BigEndian.PutUint32(raw[idx*schema.FieldSize:(idx+1)*schema.FieldSize], uint32(val))
	return w.rdb.Set(ctx, sdsKey, raw, 0).Err()
}
