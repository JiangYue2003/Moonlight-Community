package knowpostlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
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

// PatchMetadata 增量更新元数据；只覆盖 *_set=true 的字段。
func (l *PatchMetadataLogic) PatchMetadata(in *knowpost.PatchMetadataReq) (*knowpost.KnowPostDetail, error) {
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	if in.TitleSet {
		row.Title = setNS(in.Title)
	}
	if in.DescriptionSet {
		if len([]rune(in.Description)) > 50 {
			return nil, errorx.New(errorx.CodeBadRequest, "description max 50 runes")
		}
		row.Description = setNS(in.Description)
	}
	if in.TagIdSet {
		if in.TagId > 0 {
			row.TagId.Int64 = in.TagId
			row.TagId.Valid = true
		} else {
			row.TagId.Valid = false
		}
	}
	if in.TagsSet {
		s, err := encodeStringList(in.Tags)
		if err != nil {
			return nil, errorx.Wrap(errorx.CodeBadRequest, "encode tags", err)
		}
		row.Tags = setNS(s)
	}
	if in.ImgUrlsSet {
		s, err := encodeStringList(in.ImgUrls)
		if err != nil {
			return nil, errorx.Wrap(errorx.CodeBadRequest, "encode img_urls", err)
		}
		row.ImgUrls = setNS(s)
	}
	if in.VisibleSet {
		if !validVisible(in.Visible) {
			return nil, errorx.New(errorx.CodeBadRequest, "invalid visible value")
		}
		row.Visible = in.Visible
	}
	if in.IsTopSet {
		if in.IsTop {
			row.IsTop = 1
		} else {
			row.IsTop = 0
		}
	}
	if err := updateAndEmitOutbox(l.ctx, l.svcCtx, row); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	return rowToDetail(row), nil
}

func validVisible(v string) bool {
	switch v {
	case "public", "followers", "school", "private", "unlisted":
		return true
	}
	return false
}
