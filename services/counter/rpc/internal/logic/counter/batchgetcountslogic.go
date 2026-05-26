package counterlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/sds"
)

type BatchGetCountsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetCountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetCountsLogic {
	return &BatchGetCountsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchGetCounts 用 Pipeline 批量 GET，单次往返成本与单条接近。
func (l *BatchGetCountsLogic) BatchGetCounts(in *counter.BatchGetCountsReq) (*counter.BatchGetCountsResp, error) {
	if len(in.EntityIds) == 0 {
		return &counter.BatchGetCountsResp{Result: map[string]*counter.GetCountsResp{}}, nil
	}
	pipe := l.svcCtx.Redis.Pipeline()
	cmds := make(map[string]any, len(in.EntityIds))
	for _, eid := range in.EntityIds {
		cmds[eid] = pipe.Get(l.ctx, schema.SdsKey(in.EntityType, eid))
	}
	_, _ = pipe.Exec(l.ctx) // 即使部分 key 不存在也不阻断；逐个判断 err。

	out := make(map[string]*counter.GetCountsResp, len(in.EntityIds))
	for eid, c := range cmds {
		fields := [schema.SchemaLen]int64{}
		if cmd, ok := c.(interface{ Bytes() ([]byte, error) }); ok {
			if raw, err := cmd.Bytes(); err == nil {
				fields = sds.Decode(raw)
			}
		}
		m := map[string]int64{}
		for _, metric := range schema.Supported {
			if idx := schema.IdxOf(metric); idx >= 0 {
				m[metric] = fields[idx]
			}
		}
		out[eid] = &counter.GetCountsResp{Counts: m}
	}
	return &counter.BatchGetCountsResp{Result: out}, nil
}
