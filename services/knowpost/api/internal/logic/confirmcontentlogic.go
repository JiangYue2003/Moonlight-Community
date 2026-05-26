package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type ConfirmContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewConfirmContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmContentLogic {
	return &ConfirmContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// WithId 由 handler 传入 path 解析出的 id。
func (l *ConfirmContentLogic) WithId(id int64) *ConfirmContentLogic { l.id = id; return l }

func (l *ConfirmContentLogic) ConfirmContent(req *types.ConfirmContentReq) (*types.Empty, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	_, err := l.svcCtx.KnowPostRpc.ConfirmContent(l.ctx, &pb.ConfirmContentReq{
		Id:        l.id,
		CreatorId: uid,
		ObjectKey: req.ObjectKey,
		Etag:      req.Etag,
		Size:      req.Size,
		Sha256:    req.Sha256,
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
