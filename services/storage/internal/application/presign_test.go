package application

import (
	"context"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	knowmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type fakeKnowPostsModel struct {
	knowmodel.KnowPostsModel
	row *knowmodel.KnowPosts
	err error
}

func (f *fakeKnowPostsModel) FindOne(context.Context, uint64) (*knowmodel.KnowPosts, error) {
	return f.row, f.err
}

func TestPresignRejectsAnonymous(t *testing.T) {
	service := NewService(&fakeKnowPostsModel{}, nil)
	_, err := service.Presign(context.Background(), &storagepb.PresignReq{
		Scene: "knowpost_content", PostId: "1", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeUnauthorized {
		t.Fatalf("want unauthorized, got %v", err)
	}
}

func TestPresignRejectsNonOwner(t *testing.T) {
	service := NewService(&fakeKnowPostsModel{row: &knowmodel.KnowPosts{
		Id: 7, CreatorId: 9, CreateTime: time.Now(), UpdateTime: time.Now(),
	}}, nil)
	_, err := service.Presign(context.Background(), &storagepb.PresignReq{
		UserId: 1, Scene: "knowpost_content", PostId: "7", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("want forbidden, got %v", err)
	}
}

func TestPresignMapsNotFoundToForbidden(t *testing.T) {
	service := NewService(&fakeKnowPostsModel{err: sqlx.ErrNotFound}, nil)
	_, err := service.Presign(context.Background(), &storagepb.PresignReq{
		UserId: 1, Scene: "knowpost_content", PostId: "7", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("want forbidden, got %v", err)
	}
}
