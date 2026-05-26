package storagelogic

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	knowmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/svc"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type fakeKnowPostsModel struct {
	knowmodel.KnowPostsModel
	row *knowmodel.KnowPosts
	err error
}

func (f *fakeKnowPostsModel) FindOne(ctx context.Context, id uint64) (*knowmodel.KnowPosts, error) {
	return f.row, f.err
}

func TestPresign_RejectsAnonymous(t *testing.T) {
	sc := &svc.ServiceContext{KnowPostsModel: &fakeKnowPostsModel{}}
	_, err := NewPresignLogic(context.Background(), sc).Presign(&storagepb.PresignReq{
		Scene: "knowpost_content", PostId: "1", ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeUnauthorized {
		t.Fatalf("want unauthorized, got %v", err)
	}
}

func TestPresign_RejectsNonOwner(t *testing.T) {
	sc := &svc.ServiceContext{
		KnowPostsModel: &fakeKnowPostsModel{row: &knowmodel.KnowPosts{
			Id:         7,
			CreatorId:  9,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}},
	}
	_, err := NewPresignLogic(context.Background(), sc).Presign(&storagepb.PresignReq{
		UserId:      1,
		Scene:       "knowpost_content",
		PostId:      "7",
		ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("want forbidden, got %v", err)
	}
}

func TestPresign_NotFoundMappedToForbidden(t *testing.T) {
	sc := &svc.ServiceContext{KnowPostsModel: &fakeKnowPostsModel{err: sqlx.ErrNotFound}}
	_, err := NewPresignLogic(context.Background(), sc).Presign(&storagepb.PresignReq{
		UserId:      1,
		Scene:       "knowpost_content",
		PostId:      "7",
		ContentType: "text/markdown",
	})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeForbidden {
		t.Fatalf("want forbidden, got %v", err)
	}
}

var _ = errors.New
var _ = sql.NullTime{}
