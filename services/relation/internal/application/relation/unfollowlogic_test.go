package relationlogic

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

func TestUnfollow_DuplicateUnfollowDoesNotWriteOutbox(t *testing.T) {
	sc, mock, _ := newFollowFixture(t)
	mock.ExpectBegin()
	mock.ExpectExec("update `following` set rel_status=0").
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	resp, err := NewUnfollowLogic(context.Background(), sc).Unfollow(&relation.UnfollowReq{FromUserId: 21, ToUserId: 22})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Changed {
		t.Fatalf("duplicate unfollow should not report changed=true")
	}
}
