package search

import (
	"context"

	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SearchReq   = searchpb.SearchReq
	SearchResp  = searchpb.SearchResp
	SuggestReq  = searchpb.SuggestReq
	SuggestResp = searchpb.SuggestResp

	Search interface {
		Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error)
		Suggest(ctx context.Context, in *SuggestReq, opts ...grpc.CallOption) (*SuggestResp, error)
	}

	defaultSearch struct {
		cli zrpc.Client
	}
)

func NewSearch(cli zrpc.Client) Search {
	return &defaultSearch{cli: cli}
}

func (m *defaultSearch) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error) {
	client := searchpb.NewSearchClient(m.cli.Conn())
	return client.Search(ctx, in, opts...)
}

func (m *defaultSearch) Suggest(ctx context.Context, in *SuggestReq, opts ...grpc.CallOption) (*SuggestResp, error) {
	client := searchpb.NewSearchClient(m.cli.Conn())
	return client.Suggest(ctx, in, opts...)
}
