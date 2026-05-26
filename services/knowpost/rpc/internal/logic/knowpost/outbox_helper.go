package knowpostlogic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/txx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

// updateAndEmitOutbox 在同一事务里：UPDATE knowposts + INSERT outbox。
//
// 无论 status 是否为 published，都写 outbox（KnowPostUpdated）。
// search-indexer 收到事件后会检查 status/visible，对非 published 帖子执行 SoftDelete，
// 确保撤稿、改可见性等操作能及时从 ES 中移除文档。
// 删除事件由 deletelogic 自行处理（事件类型 KnowPostDeleted），不走此 helper。
func updateAndEmitOutbox(ctx context.Context, sc *svc.ServiceContext, row *model.KnowPosts) error {
	outboxId, err := sc.Snowflake.NextId()
	if err != nil {
		return err
	}
	payload, _ := json.Marshal(event.KnowPostEvent{
		Type:   event.TypeKnowPostUpdated,
		PostId: int64(row.Id),
		Author: int64(row.CreatorId),
	})
	return txx.WithTx(ctx, sc.Db, func(ctx context.Context, sess sqlx.Session) error {
		if err := sc.KnowPostsModel.UpdateInTx(ctx, sess, row); err != nil {
			return err
		}
		return sc.OutboxModel.InsertInTx(ctx, sess, outboxId,
			event.AggregateType, int64(row.Id), event.TypeKnowPostUpdated, string(payload))
	})
}
