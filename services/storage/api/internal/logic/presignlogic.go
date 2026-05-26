package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	"github.com/zhiguang/zhiguang-go/services/storage/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/storage/api/internal/types"
)

type PresignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPresignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PresignLogic {
	return &PresignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Presign 流程：
//  1. 校验 scene 与 contentType
//  2. 校验 postId 归属（创建者必须是当前用户，否则 Forbidden）
//  3. 推导扩展名与 objectKey
//  4. 调用 ossx 生成预签名 PUT URL
//  5. 返回 url + headers + 公开访问 ContentUrl
func (l *PresignLogic) Presign(req *types.PresignReq) (*types.PresignResp, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
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
		postId, err := strconv.ParseInt(req.PostId, 10, 64)
		if err != nil {
			return nil, errorx.Wrap(errorx.CodeBadRequest, "postId not a number", err)
		}
		post, err := l.svcCtx.KnowPostsModel.FindOne(l.ctx, uint64(postId))
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, errorx.New(errorx.CodeForbidden, "draft not found or no permission")
			}
			return nil, err
		}
		if int64(post.CreatorId) != uid {
			return nil, errorx.New(errorx.CodeForbidden, "draft not owned by you")
		}
	default:
		return nil, errorx.New(errorx.CodeBadRequest, "unknown scene: "+req.Scene)
	}

	ext := ossx.ExtFromContentType(req.ContentType, req.Scene, req.Ext)
	objectKey := ossx.ObjectKeyFor(req.Scene, req.PostId, uid, l.svcCtx.Oss.AvatarFolder(), ext)

	signed, err := l.svcCtx.Oss.Presign(l.ctx, ossx.PresignReq{
		ObjectKey:   objectKey,
		ContentType: req.ContentType,
	})
	if err != nil {
		return nil, err
	}
	return &types.PresignResp{
		Url:        signed.Url,
		ObjectKey:  signed.ObjectKey,
		ExpiresIn:  signed.ExpiresIn,
		Headers:    signed.Headers,
		ContentUrl: signed.ContentUrl,
	}, nil
}
