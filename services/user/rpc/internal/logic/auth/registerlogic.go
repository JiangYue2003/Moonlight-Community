package authlogic

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/model"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.AuthResp, error) {
	if !in.AgreeTerms {
		return nil, errorx.New(errorx.CodeTermsNotAccepted, "must accept terms")
	}
	id := normalizeIdentifier(in.Identifier)
	if err := validateIdentifier(id); err != nil {
		return nil, err
	}
	if err := validatePassword(in.Password, l.svcCtx.Config.Password.MinLength); err != nil {
		return nil, err
	}
	if err := l.svcCtx.Verifier.Verify(l.ctx, "REGISTER", id, in.Code); err != nil {
		return nil, err
	}

	// 检查 identifier 是否已存在
	existing, _ := l.svcCtx.UsersModel.FindOneByIdentifier(l.ctx, id)
	if existing != nil {
		return nil, errorx.New(errorx.CodeIdentifierExists, "identifier already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), l.svcCtx.Config.Password.BcryptCost)
	if err != nil {
		return nil, err
	}

	row := &model.Users{
		Nickname:     in.Nickname,
		PasswordHash: sql.NullString{String: string(hash), Valid: true},
	}
	if phoneRe.MatchString(id) {
		row.Phone = sql.NullString{String: id, Valid: true}
	} else {
		row.Email = sql.NullString{String: id, Valid: true}
	}
	result, err := l.svcCtx.UsersModel.Insert(l.ctx, row)
	if err != nil {
		return nil, err
	}
	newId, _ := result.LastInsertId()

	pair, err := issueAndPersist(l.ctx, l.svcCtx, newId, in.Nickname)
	if err != nil {
		return nil, err
	}
	recordLoginLog(l.ctx, l.svcCtx, newId, id, "REGISTER", in.Ip, in.UserAgent, "SUCCESS")

	created, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(newId))
	if err != nil {
		return nil, err
	}
	return &user.AuthResp{User: modelToAuthUser(created), Token: pair}, nil
}
