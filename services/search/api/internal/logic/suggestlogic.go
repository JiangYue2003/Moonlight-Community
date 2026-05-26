package logic

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/types"
)

type SuggestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuggestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuggestLogic {
	return &SuggestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuggestLogic) Suggest(req *types.SuggestReq) (*types.SuggestResp, error) {
	prefix := strings.TrimSpace(req.Prefix)
	if prefix == "" {
		return &types.SuggestResp{Items: []string{}}, nil
	}
	size := req.Size
	if size <= 0 || size > 50 {
		size = 10
	}
	items, err := l.svcCtx.Es.Suggest(l.ctx, l.svcCtx.Config.ContentIndex, "title_suggest", prefix, size)
	if err != nil {
		return nil, err
	}
	if items == nil {
		items = []string{}
	}
	return &types.SuggestResp{Items: items}, nil
}
