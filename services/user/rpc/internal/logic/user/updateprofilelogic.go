package userlogic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProfile 增量更新：仅当 *_set=true 的字段才覆盖；其它字段沿用 DB 现值。
// 实现策略：先 FindOne，再按 set 标记修改字段，最后 Update（go-zero model 会自动失效缓存）。
func (l *UpdateProfileLogic) UpdateProfile(in *user.UpdateProfileReq) (*user.UpdateProfileResp, error) {
	if in.Id <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid id")
	}
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, "user not found")
		}
		return nil, err
	}
	if in.NicknameSet {
		u.Nickname = in.Nickname
	}
	if in.BioSet {
		u.Bio = setOrEmpty(in.Bio)
	}
	if in.GenderSet {
		u.Gender = setOrEmpty(in.Gender)
	}
	if in.BirthdaySet {
		if in.Birthday == "" {
			u.Birthday = sql.NullTime{}
		} else {
			t, err := time.Parse("2006-01-02", in.Birthday)
			if err != nil {
				return nil, errorx.Wrap(errorx.CodeBadRequest, "invalid birthday format", err)
			}
			u.Birthday = sql.NullTime{Time: t, Valid: true}
		}
	}
	if in.ZgIdSet {
		u.ZgId = setOrEmpty(in.ZgId)
	}
	if in.SchoolSet {
		u.School = setOrEmpty(in.School)
	}
	if in.AvatarSet {
		u.Avatar = setOrEmpty(in.Avatar)
	}
	if in.TagsJsonSet {
		u.TagsJson = setOrEmpty(in.TagsJson)
	}
	if err := l.svcCtx.UsersModel.Update(l.ctx, u); err != nil {
		return nil, err
	}
	// 回读最新（model 缓存层会用刚刚写入的值兜住）
	fresh, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	return &user.UpdateProfileResp{User: toUserInfo(fresh)}, nil
}

// setOrEmpty 若设置但值为空字符串，则把列置 NULL；非空则写入 NotNull。
func setOrEmpty(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
