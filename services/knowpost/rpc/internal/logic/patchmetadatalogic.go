package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchMetadataLogic {
	return &PatchMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchMetadataLogic) PatchMetadata(in *knowpost.PatchMetadataReq) (*knowpost.KnowPostDetail, error) {
	// todo: add your logic here and delete this line

	return &knowpost.KnowPostDetail{}, nil
}
