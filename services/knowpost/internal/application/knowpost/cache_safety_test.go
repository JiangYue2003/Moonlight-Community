package knowpostlogic

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc"

	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	svc "github.com/zhiguang/zhiguang-go/services/knowpost/internal/application"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/feed"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	outboxmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

type mutationKnowPostsModel struct {
	*stubKnowPostsModel
	updateErr       error
	invalidateErr   error
	updateCalls     int
	invalidateCalls int
}

func (m *mutationKnowPostsModel) UpdateInTx(context.Context, sqlx.Session, *model.KnowPosts) error {
	m.updateCalls++
	return m.updateErr
}

func (m *mutationKnowPostsModel) InvalidateCache(context.Context, int64) error {
	m.invalidateCalls++
	return m.invalidateErr
}

type mutationOutboxModel struct {
	insertErr   error
	insertCalls int
}

func (*mutationOutboxModel) Insert(context.Context, *outboxmodel.Outbox) (sql.Result, error) {
	panic("not implemented")
}

func (*mutationOutboxModel) FindOne(context.Context, uint64) (*outboxmodel.Outbox, error) {
	panic("not implemented")
}

func (*mutationOutboxModel) Update(context.Context, *outboxmodel.Outbox) error {
	panic("not implemented")
}

func (*mutationOutboxModel) Delete(context.Context, uint64) error {
	panic("not implemented")
}

func (m *mutationOutboxModel) InsertInTx(
	context.Context,
	sqlx.Session,
	int64,
	string,
	int64,
	string,
	string,
) error {
	m.insertCalls++
	return m.insertErr
}

type mutationEpochStore struct {
	safety     uint64
	bumpErr    error
	flushErr   error
	pending    bool
	bumpCalls  int
	markCalls  int
	flushCalls int
}

func (*mutationEpochStore) Relation(context.Context, int64) (uint64, error)     { return 0, nil }
func (*mutationEpochStore) BumpRelation(context.Context, int64) (uint64, error) { return 0, nil }
func (s *mutationEpochStore) Safety(context.Context) (uint64, error)            { return s.safety, nil }

func (s *mutationEpochStore) BumpSafety(context.Context) (uint64, error) {
	s.bumpCalls++
	if s.bumpErr != nil {
		return 0, s.bumpErr
	}
	s.safety++
	s.pending = false
	return s.safety, nil
}

func (s *mutationEpochStore) MarkSafetyPending() {
	s.markCalls++
	s.pending = true
}

func (s *mutationEpochStore) SafetyPending() bool { return s.pending }

func (s *mutationEpochStore) FlushPendingSafety(context.Context) error {
	s.flushCalls++
	if s.flushErr != nil {
		return s.flushErr
	}
	if s.pending {
		s.safety++
		s.pending = false
	}
	return nil
}

type sensitiveMutation func(*svc.ServiceContext) error

type publishUserCounter struct{}

func (publishUserCounter) UserIncrement(
	context.Context,
	*counterpb.UserIncrementReq,
	...grpc.CallOption,
) (*counterpb.UserIncrementResp, error) {
	return &counterpb.UserIncrementResp{}, nil
}

func (publishUserCounter) GetUserSnapshot(
	context.Context,
	*counterpb.GetUserSnapshotReq,
	...grpc.CallOption,
) (*counterpb.GetUserSnapshotResp, error) {
	return &counterpb.GetUserSnapshotResp{Snapshot: &counterpb.UserSnapshot{Followers: feed.BIGV_THRESHOLD + 1}}, nil
}

func (publishUserCounter) BatchGetUserSnapshot(
	context.Context,
	*counterpb.BatchGetUserSnapshotReq,
	...grpc.CallOption,
) (*counterpb.BatchGetUserSnapshotResp, error) {
	return &counterpb.BatchGetUserSnapshotResp{}, nil
}

func TestSensitiveMutationsBumpSafetyExactlyOnceAfterCommit(t *testing.T) {
	for _, tc := range []struct {
		name string
		run  sensitiveMutation
	}{
		{name: "delete", run: func(sc *svc.ServiceContext) error {
			_, err := NewDeleteLogic(context.Background(), sc).Delete(&pb.DeleteReq{Id: 9, CreatorId: 7})
			return err
		}},
		{name: "visibility", run: func(sc *svc.ServiceContext) error {
			_, err := NewUpdateVisibilityLogic(context.Background(), sc).UpdateVisibility(&pb.UpdateVisibilityReq{Id: 9, CreatorId: 7, Visible: "private"})
			return err
		}},
		{name: "top", run: func(sc *svc.ServiceContext) error {
			_, err := NewUpdateTopLogic(context.Background(), sc).UpdateTop(&pb.UpdateTopReq{Id: 9, CreatorId: 7, IsTop: true})
			return err
		}},
		{name: "metadata", run: func(sc *svc.ServiceContext) error {
			_, err := NewPatchMetadataLogic(context.Background(), sc).PatchMetadata(&pb.PatchMetadataReq{
				Id: 9, CreatorId: 7, TitleSet: true, Title: "changed",
			})
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sc, sqlMock, epochs, model := newSensitiveMutationFixture(t, nil)
			sqlMock.ExpectBegin()
			sqlMock.ExpectCommit()

			err := tc.run(sc)
			if err != nil {
				t.Fatalf("mutation failed: %v", err)
			}
			if epochs.bumpCalls != 1 || epochs.safety != 1 {
				t.Fatalf("safety bump calls=%d epoch=%d, want 1/1", epochs.bumpCalls, epochs.safety)
			}
			if model.updateCalls != 1 || model.invalidateCalls != 1 {
				t.Fatalf("model calls update=%d invalidate=%d, want 1/1", model.updateCalls, model.invalidateCalls)
			}
			if err := sqlMock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSensitiveMutationsDoNotBumpSafetyWhenTransactionFails(t *testing.T) {
	updateErr := errors.New("update failed")
	for _, tc := range []struct {
		name string
		run  sensitiveMutation
	}{
		{name: "delete", run: func(sc *svc.ServiceContext) error {
			_, err := NewDeleteLogic(context.Background(), sc).Delete(&pb.DeleteReq{Id: 9, CreatorId: 7})
			return err
		}},
		{name: "visibility", run: func(sc *svc.ServiceContext) error {
			_, err := NewUpdateVisibilityLogic(context.Background(), sc).UpdateVisibility(&pb.UpdateVisibilityReq{Id: 9, CreatorId: 7, Visible: "private"})
			return err
		}},
		{name: "top", run: func(sc *svc.ServiceContext) error {
			_, err := NewUpdateTopLogic(context.Background(), sc).UpdateTop(&pb.UpdateTopReq{Id: 9, CreatorId: 7, IsTop: true})
			return err
		}},
		{name: "metadata", run: func(sc *svc.ServiceContext) error {
			_, err := NewPatchMetadataLogic(context.Background(), sc).PatchMetadata(&pb.PatchMetadataReq{Id: 9, CreatorId: 7, TitleSet: true, Title: "changed"})
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sc, sqlMock, epochs, _ := newSensitiveMutationFixture(t, updateErr)
			sqlMock.ExpectBegin()
			sqlMock.ExpectRollback()

			err := tc.run(sc)
			if !errors.Is(err, updateErr) {
				t.Fatalf("mutation error=%v, want %v", err, updateErr)
			}
			if epochs.bumpCalls != 0 || epochs.markCalls != 0 {
				t.Fatalf("failed transaction touched safety epoch: bump=%d mark=%d", epochs.bumpCalls, epochs.markCalls)
			}
			if err := sqlMock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSafetyBumpFailureKeepsCommittedMutationSuccessfulAndMarksPending(t *testing.T) {
	sc, sqlMock, epochs, _ := newSensitiveMutationFixture(t, nil)
	epochs.bumpErr = errors.New("redis unavailable")
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	_, err := NewUpdateTopLogic(context.Background(), sc).UpdateTop(&pb.UpdateTopReq{Id: 9, CreatorId: 7, IsTop: true})
	if err != nil {
		t.Fatalf("committed mutation returned safety error: %v", err)
	}
	if epochs.bumpCalls != 1 || epochs.markCalls != 1 || !epochs.pending {
		t.Fatalf("pending state bump=%d mark=%d pending=%t, want 1/1/true", epochs.bumpCalls, epochs.markCalls, epochs.pending)
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestCommittedMutationStillBumpsSafetyWhenModelCacheInvalidationFails(t *testing.T) {
	sc, sqlMock, epochs, postModel := newSensitiveMutationFixture(t, nil)
	cacheErr := errors.New("model cache invalidation failed")
	postModel.invalidateErr = cacheErr
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	_, err := NewUpdateTopLogic(context.Background(), sc).UpdateTop(&pb.UpdateTopReq{Id: 9, CreatorId: 7, IsTop: true})
	if !errors.Is(err, cacheErr) {
		t.Fatalf("mutation error=%v, want original cache error %v", err, cacheErr)
	}
	if epochs.bumpCalls != 1 || epochs.safety != 1 {
		t.Fatalf("committed mutation bump calls=%d epoch=%d, want 1/1", epochs.bumpCalls, epochs.safety)
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestRepeatedSafetyBumpFailuresCoalescePendingCompensation(t *testing.T) {
	sc, sqlMock, epochs, _ := newSensitiveMutationFixture(t, nil)
	epochs.bumpErr = errors.New("redis unavailable")
	for range 2 {
		sqlMock.ExpectBegin()
		sqlMock.ExpectCommit()
		if _, err := NewUpdateTopLogic(context.Background(), sc).UpdateTop(&pb.UpdateTopReq{Id: 9, CreatorId: 7, IsTop: true}); err != nil {
			t.Fatalf("committed mutation returned safety error: %v", err)
		}
	}
	if epochs.bumpCalls != 2 || epochs.markCalls != 1 || !epochs.pending {
		t.Fatalf("pending state bump=%d mark=%d pending=%t, want 2/1/true", epochs.bumpCalls, epochs.markCalls, epochs.pending)
	}
	epochs.bumpErr = nil
	if err := epochs.FlushPendingSafety(context.Background()); err != nil {
		t.Fatalf("flush pending safety: %v", err)
	}
	if epochs.safety != 1 || epochs.pending {
		t.Fatalf("compensated epoch=%d pending=%t, want 1/false", epochs.safety, epochs.pending)
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestPublishDoesNotBumpGlobalSafetyEpoch(t *testing.T) {
	sc, sqlMock, epochs, _ := newSensitiveMutationFixture(t, nil)
	row := sc.KnowPostsModel.(*mutationKnowPostsModel).row
	row.Status = "draft"
	row.ContentObjectKey = sql.NullString{String: "knowpost/9.md", Valid: true}

	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	sc.Redis = rdb
	sc.UserCounterRpc = publishUserCounter{}
	sc.FeedWriter = feed.NewFeedWriter(feed.NewRedisAdapter(rdb), nil, logx.WithContext(context.Background()))
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	if _, err := NewPublishLogic(context.Background(), sc).Publish(&pb.PublishReq{Id: 9, CreatorId: 7}); err != nil {
		t.Fatalf("publish failed: %v", err)
	}
	if epochs.bumpCalls != 0 || epochs.markCalls != 0 {
		t.Fatalf("publish touched global safety epoch: bump=%d mark=%d", epochs.bumpCalls, epochs.markCalls)
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func newSensitiveMutationFixture(
	t *testing.T,
	updateErr error,
) (*svc.ServiceContext, sqlmock.Sqlmock, *mutationEpochStore, *mutationKnowPostsModel) {
	t.Helper()
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	row := newKnowpostRow("public")
	postModel := &mutationKnowPostsModel{
		stubKnowPostsModel: &stubKnowPostsModel{row: row},
		updateErr:          updateErr,
	}
	epochs := &mutationEpochStore{}
	return &svc.ServiceContext{
		Db:             sqlx.NewSqlConnFromDB(db),
		KnowPostsModel: postModel,
		OutboxModel:    &mutationOutboxModel{},
		Snowflake:      snowflakex.MustNew(1, 1),
		FeedEpochs:     epochs,
	}, sqlMock, epochs, postModel
}
