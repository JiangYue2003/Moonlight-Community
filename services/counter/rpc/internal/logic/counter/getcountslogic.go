package counterlogic

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/sds"
)

// 阶段4 重建覆盖范围：使用 Redis SCAN 全量枚举 chunk key，覆盖任意 entityId。
// SCAN 单批 ScanCount 个，pipeline BITCOUNT 累加。
const (
	RebuildLockTtl   = 5 * time.Second
	RebuildWaitMs    = 50
	RebuildSdsTtl    = 0 // 持久（与位图同生命周期）
	RebuildScanCount = 256
)

func (l *GetCountsLogic) rebuildLockTTL() time.Duration {
	ttlMs := l.svcCtx.Config.Rebuild.Lock.TtlMs
	if ttlMs <= 0 {
		ttlMs = int(RebuildLockTtl / time.Millisecond)
	}
	return time.Duration(ttlMs) * time.Millisecond
}

type GetCountsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountsLogic {
	return &GetCountsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCounts 读 SDS；缺失时走 lockx + BITCOUNT 自愈重建。
//
// 重建策略（与原 Java 一致）：
//  1. SDS hit → 解码返回
//  2. SDS miss + 抢到锁 → 双检 → BITCOUNT 重建并写回
//  3. SDS miss + 抢锁失败 → 短等 50ms 后再 GET：拿到则返回；仍无则返回零值（不阻塞）
func (l *GetCountsLogic) GetCounts(in *counter.GetCountsReq) (*counter.GetCountsResp, error) {
	if in.EntityType == "" || in.EntityId == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "missing entity")
	}
	sdsKey := schema.SdsKey(in.EntityType, in.EntityId)

	raw, err := l.svcCtx.Redis.Get(l.ctx, sdsKey).Bytes()
	if err == nil && len(raw) > 0 {
		return decodeToResp(raw), nil
	}
	if err != nil && !errors.Is(err, goredis.Nil) {
		return nil, err
	}

	// SDS 缺失，进入重建路径
	return l.rebuildAndReturn(sdsKey, in.EntityType, in.EntityId)
}

func (l *GetCountsLogic) rebuildAndReturn(sdsKey, etype, eid string) (*counter.GetCountsResp, error) {
	lockKey := fmt.Sprintf("lock:rebuild:%s:%s", etype, eid)
	release, ok, err := l.svcCtx.Locks.TryAcquire(l.ctx, lockKey, l.rebuildLockTTL())
	if err != nil {
		l.Logger.Errorf("acquire rebuild lock: %v", err)
	}
	if !ok {
		// 抢锁失败 → 短等 → 再 GET，仍 miss 就返回零值
		time.Sleep(RebuildWaitMs * time.Millisecond)
		if raw, err := l.svcCtx.Redis.Get(l.ctx, sdsKey).Bytes(); err == nil && len(raw) > 0 {
			return decodeToResp(raw), nil
		}
		return zeroResp(), nil
	}
	defer func() {
		if err := release(); err != nil {
			l.Logger.Errorf("release rebuild lock: %v", err)
		}
	}()

	// 双检：可能其他副本刚刚写好
	if raw, err := l.svcCtx.Redis.Get(l.ctx, sdsKey).Bytes(); err == nil && len(raw) > 0 {
		return decodeToResp(raw), nil
	}

	// 真重建：用 SCAN 全量枚举 chunk key，pipeline BITCOUNT。
	rebuilt := make([]int64, schema.SchemaLen)
	for _, metric := range schema.Supported {
		idx := schema.IdxOf(metric)
		if idx < 0 {
			continue
		}
		total, err := l.scanBitmapTotal(metric, etype, eid)
		if err != nil {
			return nil, err
		}
		rebuilt[idx] = total
	}

	// 写回 SDS（直接 SET 完整 20 字节，不走 IncrField Lua —— 重建是绝对值替换）
	buf := make([]byte, schema.SchemaLen*schema.FieldSize)
	for i, v := range rebuilt {
		binary.BigEndian.PutUint32(buf[i*schema.FieldSize:(i+1)*schema.FieldSize], uint32(v))
	}
	if err := l.svcCtx.Redis.Set(l.ctx, sdsKey, buf, RebuildSdsTtl).Err(); err != nil {
		l.Logger.Errorf("rebuild SDS write: %v", err)
		// 写失败不阻塞读路径，照常返回
	}

	out := map[string]int64{}
	for _, m := range schema.Supported {
		idx := schema.IdxOf(m)
		if idx >= 0 {
			out[m] = rebuilt[idx]
		}
	}
	return &counter.GetCountsResp{Counts: out}, nil
}

func decodeToResp(raw []byte) *counter.GetCountsResp {
	fields := sds.Decode(raw)
	out := map[string]int64{}
	for _, m := range schema.Supported {
		idx := schema.IdxOf(m)
		if idx >= 0 {
			out[m] = fields[idx]
		}
	}
	return &counter.GetCountsResp{Counts: out}
}

func zeroResp() *counter.GetCountsResp {
	out := map[string]int64{}
	for _, m := range schema.Supported {
		out[m] = 0
	}
	return &counter.GetCountsResp{Counts: out}
}

// scanBitmapTotal 用 Redis SCAN 全量枚举该 (metric, etype, eid) 下的所有 chunk key，pipeline BITCOUNT 求和。
//
// 与 Java 端 BitmapShard 命名一致：bm:{metric}:{etype}:{eid}:{chunk}
func (l *GetCountsLogic) scanBitmapTotal(metric, etype, eid string) (int64, error) {
	pattern := fmt.Sprintf("bm:%s:%s:%s:*", metric, etype, eid)
	var cursor uint64
	var total int64
	for {
		keys, next, err := l.svcCtx.Redis.Scan(l.ctx, cursor, pattern, RebuildScanCount).Result()
		if err != nil {
			return 0, err
		}
		if len(keys) > 0 {
			pipe := l.svcCtx.Redis.Pipeline()
			cmds := make([]*goredis.IntCmd, len(keys))
			for i, k := range keys {
				cmds[i] = pipe.BitCount(l.ctx, k, nil)
			}
			if _, err := pipe.Exec(l.ctx); err != nil {
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
