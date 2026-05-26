package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type PatchProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchProfileLogic {
	return &PatchProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PatchProfile 流程：
//  1. 校验入参（与 Java ProfilePatchRequest 一致）
//  2. 若提交了 ZgId，调 user-rpc.ExistsByZgIdExceptId 排他校验
//  3. 拼装 UpdateProfileReq（仅 *_set=true 的字段）调 user-rpc
//  4. 返回最新 profile
func (l *PatchProfileLogic) PatchProfile(req *types.PatchProfileReq) (*types.ProfileResp, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	if err := validatePatchReq(req); err != nil {
		return nil, err
	}

	if req.ZgId != nil && *req.ZgId != "" {
		ex, err := l.svcCtx.UserRpc.ExistsByZgIdExceptId(l.ctx,
			&userpb.ExistsByZgIdExceptIdReq{ZgId: *req.ZgId, ExceptId: uid})
		if err != nil {
			return nil, err
		}
		if ex.Exists {
			return nil, errorx.New(errorx.CodeZgIdExists, "zgId already taken")
		}
	}

	in := &userpb.UpdateProfileReq{Id: uid}
	if req.Nickname != nil {
		in.Nickname, in.NicknameSet = *req.Nickname, true
	}
	if req.Bio != nil {
		in.Bio, in.BioSet = *req.Bio, true
	}
	if req.Gender != nil {
		in.Gender, in.GenderSet = *req.Gender, true
	}
	if req.Birthday != nil {
		in.Birthday, in.BirthdaySet = *req.Birthday, true
	}
	if req.ZgId != nil {
		in.ZgId, in.ZgIdSet = *req.ZgId, true
	}
	if req.School != nil {
		in.School, in.SchoolSet = *req.School, true
	}
	if req.TagsJson != nil {
		in.TagsJson, in.TagsJsonSet = *req.TagsJson, true
	} else if req.TagJson != nil {
		in.TagsJson, in.TagsJsonSet = *req.TagJson, true
	}
	// 前端当前可能会带 phone/email，这里先兼容接收但不更新（避免误报 no field）
	if req.Phone != nil || req.Email != nil {
		// no-op: 用户手机/邮箱改绑当前不在 profile API 范围
	}

	resp, err := l.svcCtx.UserRpc.UpdateProfile(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return toProfileResp(resp.User), nil
}
