package processor

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/svc"
)

type Backfiller struct{ sc *svc.ServiceContext }

func NewBackfiller(sc *svc.ServiceContext) *Backfiller { return &Backfiller{sc: sc} }

// Run 仅当向量索引为空时全量回填一次。
func (b *Backfiller) Run(ctx context.Context, p *Processor) error {
	count, err := b.sc.Es.Count(ctx, b.sc.Config.RagIndex)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	logx.Info("rag-indexer: vector index empty, starting backfill")
	const pageSize int32 = 50
	indexed := 0
	for page := int32(1); page <= 200; page++ {
		resp, err := b.sc.KnowPostRpc.GetPublicFeed(ctx, &knowpostpb.GetPublicFeedReq{Page: page, Size: pageSize})
		if err != nil {
			return err
		}
		if len(resp.Items) == 0 {
			break
		}
		for _, it := range resp.Items {
			pid, _ := strconv.ParseInt(it.Id, 10, 64)
			if pid <= 0 {
				continue
			}
			if err := p.upsert(ctx, pid); err != nil {
				logx.Errorf("rag backfill upsert %d: %v", pid, err)
				continue
			}
			indexed++
		}
		if int32(len(resp.Items)) < pageSize {
			break
		}
	}
	logx.Infof("rag-indexer: backfill done, processed=%d", indexed)
	return nil
}
