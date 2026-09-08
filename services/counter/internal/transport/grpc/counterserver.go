package grpcserver

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/counter/internal/application"
	counterlogic "github.com/zhiguang/zhiguang-go/services/counter/internal/application/counter"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type CounterServer struct {
	svcCtx *application.ServiceContext
	counterpb.UnimplementedCounterServer
}

func NewCounterServer(svcCtx *application.ServiceContext) *CounterServer {
	return &CounterServer{
		svcCtx: svcCtx,
	}
}

func (s *CounterServer) Toggle(ctx context.Context, in *counterpb.ToggleReq) (*counterpb.ToggleResp, error) {
	l := counterlogic.NewToggleLogic(ctx, s.svcCtx)
	return l.Toggle(in)
}

func (s *CounterServer) GetCounts(ctx context.Context, in *counterpb.GetCountsReq) (*counterpb.GetCountsResp, error) {
	l := counterlogic.NewGetCountsLogic(ctx, s.svcCtx)
	return l.GetCounts(in)
}

func (s *CounterServer) IsMarked(ctx context.Context, in *counterpb.IsMarkedReq) (*counterpb.IsMarkedResp, error) {
	l := counterlogic.NewIsMarkedLogic(ctx, s.svcCtx)
	return l.IsMarked(in)
}

func (s *CounterServer) BatchGetCounts(ctx context.Context, in *counterpb.BatchGetCountsReq) (*counterpb.BatchGetCountsResp, error) {
	l := counterlogic.NewBatchGetCountsLogic(ctx, s.svcCtx)
	return l.BatchGetCounts(in)
}
