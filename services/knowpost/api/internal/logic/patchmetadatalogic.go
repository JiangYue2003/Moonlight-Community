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

type PatchMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewPatchMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchMetadataLogic {
	return &PatchMetadataLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}
func (l *PatchMetadataLogic) WithId(id int64) *PatchMetadataLogic { l.id = id; return l }

func (l *PatchMetadataLogic) PatchMetadata(req *types.PatchMetadataReq) (*types.KnowPostDetail, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	in := &pb.PatchMetadataReq{Id: l.id, CreatorId: uid}
	if req.Title != nil {
		in.Title, in.TitleSet = *req.Title, true
	}
	if req.Description != nil {
		in.Description, in.DescriptionSet = *req.Description, true
	}
	if req.TagId != nil {
		in.TagId, in.TagIdSet = *req.TagId, true
	}
	if req.TagsSet || req.Tags != nil {
		in.Tags, in.TagsSet = req.Tags, true
	}
	if req.ImgUrlsSet || req.ImgUrls != nil {
		in.ImgUrls, in.ImgUrlsSet = req.ImgUrls, true
	}
	if req.Visible != nil {
		in.Visible, in.VisibleSet = *req.Visible, true
	}
	if req.IsTop != nil {
		in.IsTop, in.IsTopSet = *req.IsTop, true
	}
	r, err := l.svcCtx.KnowPostRpc.PatchMetadata(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return detailFromPb(l.ctx, l.svcCtx, r), nil
}
