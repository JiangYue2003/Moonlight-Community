package logic

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	avatarMaxBytes int64 = 5 * 1024 * 1024 // 5MB
)

// UploadAvatar 后端中转：multipart → ossx.PutObject → user-rpc.UpdateProfile(avatar=url)
func (l *UploadAvatarLogic) UploadAvatar(r *http.Request) (*types.UploadAvatarResp, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	// 限制 multipart 解析时的最大缓冲，超过部分会直接拒绝
	if err := r.ParseMultipartForm(avatarMaxBytes); err != nil {
		return nil, errorx.Wrap(errorx.CodeBadRequest, "parse multipart failed", err)
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, errorx.Wrap(errorx.CodeBadRequest, "missing form field 'file'", err)
	}
	defer file.Close()
	if header.Size > avatarMaxBytes {
		return nil, errorx.New(errorx.CodeBadRequest, "avatar exceeds 5MB")
	}
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return nil, errorx.New(errorx.CodeBadRequest, "avatar must be an image")
	}

	ext := ossx.ExtFromContentType(contentType, ossx.SceneAvatar, "")
	key := ossx.ObjectKeyFor(ossx.SceneAvatar, "", uid, l.svcCtx.Oss.AvatarFolder(), ext)

	if _, err := l.svcCtx.Oss.PutObject(l.ctx, key, file, contentType); err != nil {
		return nil, err
	}
	url := l.svcCtx.Oss.BuildContentUrl(key)

	if _, err := l.svcCtx.UserRpc.UpdateProfile(l.ctx, &userpb.UpdateProfileReq{
		Id:        uid,
		Avatar:    url,
		AvatarSet: true,
	}); err != nil {
		return nil, err
	}
	return &types.UploadAvatarResp{Url: url, Avatar: url, ObjectKey: key}, nil
}
