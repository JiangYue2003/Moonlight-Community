package usercounterlogic

import (
	"context"
	"errors"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/sds"
)

type UserIncrementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIncrementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIncrementLogic {
	return &UserIncrementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIncrementLogic) UserIncrement(in *counter.UserIncrementReq) (*counter.UserIncrementResp, error) {
	if in.UserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid user_id")
	}
	idx := schema.UserIdxOf(in.Field)
	if idx <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid field: "+in.Field)
	}
	val, err := l.svcCtx.IncrFieldScript.Run(l.ctx, l.svcCtx.Redis,
		[]string{schema.UserSdsKey(in.UserId)},
		idx, in.Delta, schema.UserSchemaLen, schema.UserFieldSize,
	).Int64()
	if err != nil {
		return nil, err
	}
	return &counter.UserIncrementResp{Value: val}, nil
}

type GetUserSnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSnapshotLogic {
	return &GetUserSnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSnapshotLogic) GetUserSnapshot(in *counter.GetUserSnapshotReq) (*counter.GetUserSnapshotResp, error) {
	if in.UserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid user_id")
	}
	raw, err := l.svcCtx.Redis.Get(l.ctx, schema.UserSdsKey(in.UserId)).Bytes()
	if err != nil && !errors.Is(err, goredis.Nil) {
		return nil, err
	}
	fields := sds.DecodeN(raw, schema.UserSchemaLen)
	return &counter.GetUserSnapshotResp{Snapshot: snapshotFromFields(in.UserId, fields)}, nil
}

type BatchGetUserSnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetUserSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetUserSnapshotLogic {
	return &BatchGetUserSnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchGetUserSnapshotLogic) BatchGetUserSnapshot(in *counter.BatchGetUserSnapshotReq) (*counter.BatchGetUserSnapshotResp, error) {
	out := &counter.BatchGetUserSnapshotResp{Result: make(map[int64]*counter.UserSnapshot, len(in.UserIds))}
	if len(in.UserIds) == 0 {
		return out, nil
	}
	pipe := l.svcCtx.Redis.Pipeline()
	cmds := make(map[int64]interface{}, len(in.UserIds))
	for _, uid := range in.UserIds {
		if uid <= 0 {
			continue
		}
		cmds[uid] = pipe.Get(l.ctx, schema.UserSdsKey(uid))
	}
	_, _ = pipe.Exec(l.ctx)

	for uid, c := range cmds {
		var raw []byte
		if cmd, ok := c.(interface{ Bytes() ([]byte, error) }); ok {
			if b, err := cmd.Bytes(); err == nil {
				raw = b
			}
		}
		out.Result[uid] = snapshotFromFields(uid, sds.DecodeN(raw, schema.UserSchemaLen))
	}
	return out, nil
}

func snapshotFromFields(uid int64, f []int64) *counter.UserSnapshot {
	get := func(idx int) int64 {
		if idx < 0 || idx >= len(f) {
			return 0
		}
		return f[idx]
	}
	return &counter.UserSnapshot{
		UserId:        uid,
		Read:          get(schema.UserIdxRead),
		Followings:    get(schema.UserIdxFollowings),
		Followers:     get(schema.UserIdxFollowers),
		Posts:         get(schema.UserIdxPosts),
		LikesReceived: get(schema.UserIdxLikesReceived),
	}
}
