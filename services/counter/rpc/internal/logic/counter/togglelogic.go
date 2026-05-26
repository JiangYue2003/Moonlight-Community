package counterlogic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/event"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

type ToggleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewToggleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToggleLogic {
	return &ToggleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Toggle 原子切换位图位；位图实际状态变化时才发 Kafka 事件，保证幂等。
//   - add=true 且 bit=0 → 置 1，发 +1 事件
//   - add=true 且 bit=1 → 幂等，不发事件
//   - add=false 且 bit=1 → 置 0，发 -1 事件
//   - add=false 且 bit=0 → 幂等，不发事件
func (l *ToggleLogic) Toggle(in *counter.ToggleReq) (*counter.ToggleResp, error) {
	if in.UserId <= 0 || in.EntityType == "" || in.EntityId == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "missing required fields")
	}
	idx := schema.IdxOf(in.Metric)
	if idx < 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "unsupported metric")
	}

	chunk := schema.ChunkOf(in.UserId)
	off := schema.BitOf(in.UserId)
	bmKey := schema.BitmapKey(in.Metric, in.EntityType, in.EntityId, chunk)

	op := "remove"
	delta := -1
	if in.Add {
		op = "add"
		delta = 1
	}

	// Lua 返回 1 = 状态变更，0 = 幂等。
	res, err := l.svcCtx.ToggleScript.Run(l.ctx, l.svcCtx.Redis,
		[]string{bmKey}, off, op).Int64()
	if err != nil {
		return nil, err
	}
	if res == 0 {
		return &counter.ToggleResp{Changed: false}, nil
	}

	ev := event.CounterEvent{
		EntityType: in.EntityType,
		EntityId:   in.EntityId,
		Metric:     in.Metric,
		Idx:        idx,
		UserId:     in.UserId,
		Delta:      int32(delta),
	}
	body, _ := json.Marshal(ev)
	// 用 entityId 做 key 保证同一实体顺序消费。
	if err := l.svcCtx.Kafka.Publish(l.ctx, l.svcCtx.Config.Kafka.Topic, in.EntityId, body); err != nil {
		// 写位图已成功，发 Kafka 失败不能回滚——记日志，由 rebuild consumer 兜底。
		l.Logger.Errorf("kafka publish failed: %v event=%+v", err, ev)
	}
	return &counter.ToggleResp{Changed: true}, nil
}
