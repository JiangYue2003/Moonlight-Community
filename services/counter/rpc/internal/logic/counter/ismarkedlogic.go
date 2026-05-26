package counterlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

type IsMarkedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsMarkedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsMarkedLogic {
	return &IsMarkedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IsMarked 直接 GETBIT 位图，不依赖 SDS。
func (l *IsMarkedLogic) IsMarked(in *counter.IsMarkedReq) (*counter.IsMarkedResp, error) {
	if schema.IdxOf(in.Metric) < 0 || in.UserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid request")
	}
	chunk := schema.ChunkOf(in.UserId)
	off := schema.BitOf(in.UserId)
	bmKey := schema.BitmapKey(in.Metric, in.EntityType, in.EntityId, chunk)
	v, err := l.svcCtx.Redis.GetBit(l.ctx, bmKey, off).Result()
	if err != nil {
		return nil, err
	}
	return &counter.IsMarkedResp{Marked: v == 1}, nil
}
