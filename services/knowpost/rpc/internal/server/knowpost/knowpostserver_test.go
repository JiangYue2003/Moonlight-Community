package server

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type noFollowingsRelation struct{}

func (noFollowingsRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func (noFollowingsRelation) GetFollowings(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func TestGetUserFeedIsServedByRegisteredServer(t *testing.T) {
	serviceCtx := &svc.ServiceContext{
		FeedReader: feed.NewFeedReader(
			nil,
			noFollowingsRelation{},
			nil,
			logx.WithContext(context.Background()),
		),
	}

	listener := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	knowpost.RegisterKnowPostServer(grpcServer, NewKnowPostServer(serviceCtx))

	go func() {
		_ = grpcServer.Serve(listener)
	}()
	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("create in-memory gRPC client: %v", err)
	}
	t.Cleanup(func() {
		_ = conn.Close()
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := knowpost.NewKnowPostClient(conn).GetUserFeed(ctx, &knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   1,
		Size:   20,
	})
	if status.Code(err) == codes.Unimplemented {
		t.Fatalf("GetUserFeed reached embedded unimplemented handler: %v", err)
	}
	if err != nil {
		t.Fatalf("GetUserFeed RPC failed: %v", err)
	}
	if resp == nil {
		t.Fatal("GetUserFeed returned a nil response")
	}
}
