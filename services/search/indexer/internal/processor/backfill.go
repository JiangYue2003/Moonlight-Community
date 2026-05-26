package processor

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/svc"
)

// Backfiller 启动时若 ES 索引为空 → 全量回填一次。
type Backfiller struct{ sc *svc.ServiceContext }

func NewBackfiller(sc *svc.ServiceContext) *Backfiller { return &Backfiller{sc: sc} }

// Run 阻塞直到回填完成或上下文取消。
//
// 简化策略：分页拉 GetPublicFeed（已 published+public，按 publish_time DESC），逐条 Index。
// limit 50，最多扫 200 页防御无限循环（10000 条对 dev 足够）。
func (b *Backfiller) Run(ctx context.Context) error {
	count, err := b.sc.Es.Count(ctx, b.sc.Config.ContentIndex)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	logx.Info("search-indexer: index empty, starting backfill")
	const pageSize int32 = 50
	const maxPages = 200
	indexed := 0
	for page := int32(1); page <= maxPages; page++ {
		resp, err := b.sc.KnowPostRpc.GetPublicFeed(ctx, &knowpostpb.GetPublicFeedReq{
			Page: page,
			Size: pageSize,
		})
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
			detail, err := b.sc.KnowPostRpc.GetDetail(ctx, &knowpostpb.GetDetailReq{Id: pid})
			if err != nil {
				logx.Errorf("backfill GetDetail %d: %v", pid, err)
				continue
			}
			body, _ := fetchContent(ctx, b.sc.HttpClient, detail.ContentUrl, b.sc.Config.ContentMaxRunes)
			doc := buildDoc(detail, body)
			if err := b.sc.Es.Index(ctx, b.sc.Config.ContentIndex, it.Id, doc); err != nil {
				logx.Errorf("backfill index %d: %v", pid, err)
				continue
			}
			indexed++
		}
		offset := page * pageSize
		_ = offset
		if int32(len(resp.Items)) < pageSize {
			break
		}
	}
	logx.Infof("search-indexer: backfill done, indexed=%d", indexed)
	return nil
}
