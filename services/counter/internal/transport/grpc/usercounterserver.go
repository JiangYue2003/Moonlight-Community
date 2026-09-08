package grpcserver

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/counter/internal/application"
	usercounterlogic "github.com/zhiguang/zhiguang-go/services/counter/internal/application/usercounter"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type UserCounterServer struct {
	svcCtx *application.ServiceContext
	counterpb.UnimplementedUserCounterServer
}

func NewUserCounterServer(svcCtx *application.ServiceContext) *UserCounterServer {
	return &UserCounterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCounterServer) UserIncrement(ctx context.Context, in *counterpb.UserIncrementReq) (*counterpb.UserIncrementResp, error) {
	l := usercounterlogic.NewUserIncrementLogic(ctx, s.svcCtx)
	return l.UserIncrement(in)
}

func (s *UserCounterServer) GetUserSnapshot(ctx context.Context, in *counterpb.GetUserSnapshotReq) (*counterpb.GetUserSnapshotResp, error) {
	l := usercounterlogic.NewGetUserSnapshotLogic(ctx, s.svcCtx)
	return l.GetUserSnapshot(in)
}

func (s *UserCounterServer) BatchGetUserSnapshot(ctx context.Context, in *counterpb.BatchGetUserSnapshotReq) (*counterpb.BatchGetUserSnapshotResp, error) {
	l := usercounterlogic.NewBatchGetUserSnapshotLogic(ctx, s.svcCtx)
	return l.BatchGetUserSnapshot(in)
}
