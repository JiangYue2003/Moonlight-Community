package searchlogic

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/search/rpc/internal/svc"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
)

type SuggestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSuggestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuggestLogic {
	return &SuggestLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *SuggestLogic) Suggest(req *searchpb.SuggestReq) (*searchpb.SuggestResp, error) {
	prefix := strings.TrimSpace(req.Prefix)
	if prefix == "" {
		return &searchpb.SuggestResp{Items: []string{}}, nil
	}
	size := int(req.Size)
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
	return &searchpb.SuggestResp{Items: items}, nil
}
