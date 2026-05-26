package storage

import (
	"context"

	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PresignReq  = storagepb.PresignReq
	PresignResp = storagepb.PresignResp

	Storage interface {
		Presign(ctx context.Context, in *PresignReq, opts ...grpc.CallOption) (*PresignResp, error)
	}

	defaultStorage struct {
		cli zrpc.Client
	}
)

func NewStorage(cli zrpc.Client) Storage {
	return &defaultStorage{cli: cli}
}

func (m *defaultStorage) Presign(ctx context.Context, in *PresignReq, opts ...grpc.CallOption) (*PresignResp, error) {
	client := storagepb.NewStorageClient(m.cli.Conn())
	return client.Presign(ctx, in, opts...)
}
