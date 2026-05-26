package knowpostlogic

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

type ConfirmContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfirmContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmContentLogic {
	return &ConfirmContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ConfirmContent 校验归属 → 写入 OSS 元数据 → 双删详情缓存。
func (l *ConfirmContentLogic) ConfirmContent(in *knowpost.ConfirmContentReq) (*knowpost.Empty, error) {
	if strings.TrimSpace(in.ObjectKey) == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "object_key required")
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	row.ContentObjectKey = setNS(in.ObjectKey)
	row.ContentEtag = setNS(in.Etag)
	row.ContentSize = sql.NullInt64{Int64: in.Size, Valid: in.Size > 0}
	row.ContentSha256 = setNS(in.Sha256)
	if err := l.svcCtx.KnowPostsModel.Update(l.ctx, row); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	return &knowpost.Empty{}, nil
}

// findOwnedRow 查 + 校验归属：未找到/已删除/不属于当前 creator → Forbidden（隐藏存在性）。
func findOwnedRow(ctx context.Context, sc *svc.ServiceContext, id, creatorId int64) (*model.KnowPosts, error) {
	row, err := sc.KnowPostsModel.FindOne(ctx, uint64(id))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.New(errorx.CodeForbidden, "post not found or no permission")
		}
		return nil, err
	}
	if int64(row.CreatorId) != creatorId || row.Status == "deleted" {
		return nil, errorx.New(errorx.CodeForbidden, "post not owned by you")
	}
	return row, nil
}
