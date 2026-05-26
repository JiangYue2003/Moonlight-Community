package logic

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"github.com/zhiguang/zhiguang-go/services/storage/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/storage/api/internal/types"
)

// fakeModel 用于把 KnowPostsModel 的 FindOne 行为替换成测试可控版本。
// 仅实现 FindOne；其余方法签名以 panic 提示未覆盖。
type fakeModel struct {
	model.KnowPostsModel
	row *model.KnowPosts
	err error
}

func (f *fakeModel) FindOne(ctx context.Context, id uint64) (*model.KnowPosts, error) {
	return f.row, f.err
}

func newSvcWithModel(t *testing.T, m model.KnowPostsModel) *svc.ServiceContext {
	t.Helper()
	return &svc.ServiceContext{KnowPostsModel: m}
}

func newCtx(uid int64) context.Context {
	return ctxdata.WithUserId(context.Background(), uid)
}

func TestPresign_RejectsAnonymous(t *testing.T) {
	sc := newSvcWithModel(t, &fakeModel{})
	_, err := NewPresignLogic(context.Background(), sc).Presign(&types.PresignReq{
		Scene: "knowpost_content", PostId: "1", ContentType: "text/markdown",
	})
	if err == nil {
		t.Fatal("anonymous request must be rejected")
	}
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeUnauthorized {
		t.Fatalf("want CodeUnauthorized, got %v", err)
	}
}

func TestPresign_RejectsMissingContentType(t *testing.T) {
	sc := newSvcWithModel(t, &fakeModel{})
	_, err := NewPresignLogic(newCtx(1), sc).Presign(&types.PresignReq{
		Scene: "knowpost_content", PostId: "1",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeBadRequest {
		t.Fatalf("want CodeBadRequest, got %v", err)
	}
}

func TestPresign_RejectsUnknownScene(t *testing.T) {
	sc := newSvcWithModel(t, &fakeModel{})
	_, err := NewPresignLogic(newCtx(1), sc).Presign(&types.PresignReq{
		Scene: "weird_scene", PostId: "1", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeBadRequest {
		t.Fatalf("want CodeBadRequest, got %v", err)
	}
}

func TestPresign_RejectsNonOwnerPost(t *testing.T) {
	owner := &model.KnowPosts{
		Id:         100,
		CreatorId:  999, // 帖子归 999
		CreateTime: time.Now(), UpdateTime: time.Now(),
	}
	sc := newSvcWithModel(t, &fakeModel{row: owner})
	_, err := NewPresignLogic(newCtx(1), sc).Presign(&types.PresignReq{
		Scene: "knowpost_content", PostId: "100", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("want CodeForbidden, got %v", err)
	}
}

func TestPresign_PostNotFound(t *testing.T) {
	sc := newSvcWithModel(t, &fakeModel{err: sqlx.ErrNotFound})
	_, err := NewPresignLogic(newCtx(1), sc).Presign(&types.PresignReq{
		Scene: "knowpost_content", PostId: "404", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("not-found should map to Forbidden (隐藏存在性), got %v", err)
	}
}

func TestPresign_ParsePostIdFailure(t *testing.T) {
	sc := newSvcWithModel(t, &fakeModel{})
	_, err := NewPresignLogic(newCtx(1), sc).Presign(&types.PresignReq{
		Scene: "knowpost_content", PostId: "abc", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeBadRequest {
		t.Fatalf("non-numeric postId should error: %v", err)
	}
}

// 占位防 unused import 警告（sql 仅用于让 NullTime 可访问）
var _ = sql.NullTime{}
var _ = errors.New
