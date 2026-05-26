package knowpostlogic

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type CreateDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDraftLogic {
	return &CreateDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateDraft 用 snowflake 自分配 id（与 Java 一致），插入 status=draft 行。
func (l *CreateDraftLogic) CreateDraft(in *knowpost.CreateDraftReq) (*knowpost.CreateDraftResp, error) {
	if in.CreatorId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "creator_id required")
	}
	id, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}
	row := &model.KnowPosts{
		Id:        uint64(id),
		CreatorId: uint64(in.CreatorId),
		Type:      "image_text",
		Visible:   "public",
		Status:    "draft",
		IsTop:     0,
		// JSON 字段允许 NULL；DB 默认 NULL 即可。
		Tags:    sql.NullString{},
		ImgUrls: sql.NullString{},
	}
	if _, err := l.svcCtx.KnowPostsModel.Insert(l.ctx, row); err != nil {
		return nil, err
	}
	return &knowpost.CreateDraftResp{Id: formatId(uint64(id))}, nil
}
