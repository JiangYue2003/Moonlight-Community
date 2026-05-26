// Package processor 负责把 OutboxRow 分发到 follow_handler 等具体处理器。
package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/event"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/svc"
)

// Processor 入口：把 FlatMessage 拆为 OutboxRow，去重后分发。
type Processor struct {
	sc       *svc.ServiceContext
	follow   *FollowHandler
	dedup    *Dedup
	dedupTtl time.Duration
}

func New(sc *svc.ServiceContext) *Processor {
	return &Processor{
		sc:       sc,
		follow:   NewFollowHandler(sc),
		dedup:    NewDedup(sc.Redis),
		dedupTtl: time.Duration(sc.Config.Dedup.TtlSeconds) * time.Second,
	}
}

// Handle 处理一条 Kafka 消息。返回非 nil error 时调用方不应 commit offset，让 Kafka 重投。
func (p *Processor) Handle(ctx context.Context, value []byte) error {
	flat, err := canalx.ParseFlat(value)
	if err != nil {
		// 坏消息：记日志后丢弃（不阻塞 group 进度）
		logx.WithContext(ctx).Errorf("canalx ParseFlat: %v", err)
		return nil
	}
	rows := canalx.ExtractOutboxRows(flat)
	for _, row := range rows {
		if err := p.processRow(ctx, row); err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) processRow(ctx context.Context, row canalx.OutboxRow) error {
	if row.AggregateType != "following" {
		return nil
	}

	// SETNX 去重：以 outbox.id 为单位
	dedupKey := fmt.Sprintf("dedup:rel:%s:%d", row.Type, row.Id)
	ok, err := p.dedup.Acquire(ctx, dedupKey, p.dedupTtl)
	if err != nil {
		return err
	}
	if !ok {
		return nil // 已处理过
	}

	var ev event.RelationEvent
	if err := json.Unmarshal([]byte(row.Payload), &ev); err != nil {
		// payload 损坏：跳过（已经 dedup，不会重试导致死循环）
		logx.WithContext(ctx).Errorf("payload unmarshal: %v body=%s", err, row.Payload)
		return nil
	}

	switch ev.Type {
	case event.TypeFollowCreated:
		return p.follow.HandleCreated(ctx, ev)
	case event.TypeFollowCanceled:
		return p.follow.HandleCanceled(ctx, ev)
	default:
		logx.WithContext(ctx).Infof("unknown relation event type: %s", ev.Type)
		return nil
	}
}
