package storagelogic

import (
	"context"
	"errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/svc"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type PresignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPresignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PresignLogic {
	return &PresignLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *PresignLogic) Presign(req *storagepb.PresignReq) (*storagepb.PresignResp, error) {
	if req.UserId <= 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	if req.ContentType == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "contentType required")
	}
	switch req.Scene {
	case ossx.SceneKnowPostContent, ossx.SceneKnowPostImage:
		if req.PostId == "" {
			return nil, errorx.New(errorx.CodeBadRequest, "postId required for knowpost scenes")
		}
		postID, err := strconv.ParseInt(req.PostId, 10, 64)
		if err != nil {
			return nil, errorx.Wrap(errorx.CodeBadRequest, "postId not a number", err)
		}
		post, err := l.svcCtx.KnowPostsModel.FindOne(l.ctx, uint64(postID))
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, errorx.New(errorx.CodeForbidden, "draft not found or no permission")
			}
			return nil, err
		}
		if int64(post.CreatorId) != req.UserId {
			return nil, errorx.New(errorx.CodeForbidden, "draft not owned by you")
		}
	default:
		return nil, errorx.New(errorx.CodeBadRequest, "unknown scene: "+req.Scene)
	}

	ext := ossx.ExtFromContentType(req.ContentType, req.Scene, req.Ext)
	objectKey := ossx.ObjectKeyFor(req.Scene, req.PostId, req.UserId, l.svcCtx.Oss.AvatarFolder(), ext)
	signed, err := l.svcCtx.Oss.Presign(l.ctx, ossx.PresignReq{
		ObjectKey:   objectKey,
		ContentType: req.ContentType,
	})
	if err != nil {
		return nil, err
	}
	return &storagepb.PresignResp{
		Url:        signed.Url,
		ObjectKey:  signed.ObjectKey,
		ExpiresIn:  signed.ExpiresIn,
		Headers:    signed.Headers,
		ContentUrl: signed.ContentUrl,
	}, nil
}
