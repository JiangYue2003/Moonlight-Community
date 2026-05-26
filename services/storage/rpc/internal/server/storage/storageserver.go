package server

import (
	"context"

	storagelogic "github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/logic/storage"
	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/svc"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type StorageServer struct {
	svcCtx *svc.ServiceContext
	storagepb.UnimplementedStorageServer
}

func NewStorageServer(svcCtx *svc.ServiceContext) *StorageServer {
	return &StorageServer{svcCtx: svcCtx}
}

func (s *StorageServer) Presign(ctx context.Context, in *storagepb.PresignReq) (*storagepb.PresignResp, error) {
	return storagelogic.NewPresignLogic(ctx, s.svcCtx).Presign(in)
}
