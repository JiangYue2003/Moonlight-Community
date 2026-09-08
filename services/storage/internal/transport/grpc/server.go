package grpcserver

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/storage/internal/application"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type StorageServer struct {
	service *application.Service
	storagepb.UnimplementedStorageServer
}

func NewStorageServer(service *application.Service) *StorageServer {
	return &StorageServer{service: service}
}

func (s *StorageServer) Presign(ctx context.Context, req *storagepb.PresignReq) (*storagepb.PresignResp, error) {
	return s.service.Presign(ctx, req)
}
