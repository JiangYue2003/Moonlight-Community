package server

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/search/internal/application"
	searchlogic "github.com/zhiguang/zhiguang-go/services/search/internal/application/search"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
)

type SearchServer struct {
	svcCtx *application.ServiceContext
	searchpb.UnimplementedSearchServer
}

func NewSearchServer(svcCtx *application.ServiceContext) *SearchServer {
	return &SearchServer{svcCtx: svcCtx}
}

func (s *SearchServer) Search(ctx context.Context, in *searchpb.SearchReq) (*searchpb.SearchResp, error) {
	return searchlogic.NewSearchLogic(ctx, s.svcCtx).Search(in)
}

func (s *SearchServer) Suggest(ctx context.Context, in *searchpb.SuggestReq) (*searchpb.SuggestResp, error) {
	return searchlogic.NewSuggestLogic(ctx, s.svcCtx).Suggest(in)
}
