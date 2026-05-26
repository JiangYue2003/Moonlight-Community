package knowpostlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type ReindexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReindexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReindexLogic {
	return &ReindexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Reindex 阶段2 stub：不做任何动作，仅校验归属。
// 阶段3 接 RAG indexer 时把 ev 写到 Kafka 让 worker 消费。
func (l *ReindexLogic) Reindex(in *knowpost.ReindexReq) (*knowpost.Empty, error) {
	if _, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId); err != nil {
		return nil, err
	}
	l.Logger.Infof("[stub] reindex requested for post %d", in.Id)
	return &knowpost.Empty{}, nil
}
