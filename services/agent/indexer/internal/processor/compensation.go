package processor

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// RunCompensation 定时执行失败重试与反向对账。
func RunCompensation(ctx context.Context, p *Processor, interval time.Duration) {
	if interval <= 0 {
		interval = 10 * time.Minute
	}
	tk := time.NewTicker(interval)
	defer tk.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tk.C:
			if err := p.ReconcileAndRetry(ctx, 200); err != nil {
				logx.Errorf("agent-indexer compensation err: %v", err)
			} else {
				logx.Infof("agent-indexer compensation tick done")
			}
		}
	}
}
