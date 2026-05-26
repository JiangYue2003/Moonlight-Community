package knowpostlogic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/txx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delete 软删除：status='deleted'，事务内同步写 outbox（KnowPostDeleted）。
func (l *DeleteLogic) Delete(in *knowpost.DeleteReq) (*knowpost.Empty, error) {
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	row.Status = "deleted"

	outboxId, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(event.KnowPostEvent{
		Type:   event.TypeKnowPostDeleted,
		PostId: int64(row.Id),
		Author: int64(row.CreatorId),
	})
	if err := txx.WithTx(l.ctx, l.svcCtx.Db, func(ctx context.Context, sess sqlx.Session) error {
		if err := l.svcCtx.KnowPostsModel.UpdateInTx(ctx, sess, row); err != nil {
			return err
		}
		return l.svcCtx.OutboxModel.InsertInTx(ctx, sess, outboxId,
			event.AggregateType, int64(row.Id), event.TypeKnowPostDeleted, string(payload))
	}); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	return &knowpost.Empty{}, nil
}
