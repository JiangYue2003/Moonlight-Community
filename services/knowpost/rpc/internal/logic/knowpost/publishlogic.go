package knowpostlogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/txx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Publish 校验内容已上传 → 设置 status=published、publish_time=now
//
//	→ 双删详情缓存 → usercounter posts +1
func (l *PublishLogic) Publish(in *knowpost.PublishReq) (*knowpost.KnowPostDetail, error) {
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	if !row.ContentObjectKey.Valid || row.ContentObjectKey.String == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "content not uploaded yet")
	}
	now := time.Now()
	row.Status = "published"
	row.PublishTime = sql.NullTime{Time: now, Valid: true}

	outboxId, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(event.KnowPostEvent{
		Type:   event.TypeKnowPostPublished,
		PostId: int64(row.Id),
		Author: int64(row.CreatorId),
	})
	if err := txx.WithTx(l.ctx, l.svcCtx.Db, func(ctx context.Context, sess sqlx.Session) error {
		if err := l.svcCtx.KnowPostsModel.UpdateInTx(ctx, sess, row); err != nil {
			return err
		}
		return l.svcCtx.OutboxModel.InsertInTx(ctx, sess, outboxId,
			event.AggregateType, int64(row.Id), event.TypeKnowPostPublished, string(payload))
	}); err != nil {
		return nil, err
	}
	if err := l.svcCtx.KnowPostsModel.InvalidateCache(l.ctx, int64(row.Id)); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)

	// usercounter posts +1（异步即可，但阶段2先同步以便端到端验证）
	if _, err := l.svcCtx.UserCounterRpc.UserIncrement(l.ctx, &counterpb.UserIncrementReq{
		UserId: int64(row.CreatorId),
		Field:  "posts",
		Delta:  1,
	}); err != nil {
		l.Logger.Errorf("usercounter increment posts: %v", err)
	}
	return rowToDetail(row), nil
}
