package authlogic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/token"
)

func newRevokeSvc(t *testing.T) (*svc.ServiceContext, *jwtx.Signer, goredis.UniversalClient) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis: %v", err)
	}
	t.Cleanup(mr.Close)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer, _ := jwtx.NewSigner(jwtx.Config{
		PrivateKey: priv, PublicKey: &priv.PublicKey,
		Issuer: "zhiguang", AccessTtl: time.Minute, RefreshTtl: time.Hour,
	})
	sc := &svc.ServiceContext{
		Redis:     rdb,
		JwtSigner: signer,
		Tokens:    token.NewStore(rdb),
	}
	return sc, signer, rdb
}

func TestRevokeRefresh_RemovesWhitelistEntry(t *testing.T) {
	sc, signer, _ := newRevokeSvc(t)
	pair, err := signer.IssuePair(42, "alice")
	if err != nil {
		t.Fatal(err)
	}
	if err := sc.Tokens.Save(context.Background(), 42, pair.RefreshTokenId, time.Hour); err != nil {
		t.Fatal(err)
	}

	resp, err := NewRevokeRefreshLogic(context.Background(), sc).RevokeRefresh(
		&user.RevokeRefreshReq{RefreshToken: pair.RefreshToken})
	if err != nil {
		t.Fatalf("RevokeRefresh: %v", err)
	}
	if !resp.Revoked {
		t.Fatal("expected revoked=true")
	}
	ok, _ := sc.Tokens.Valid(context.Background(), 42, pair.RefreshTokenId)
	if ok {
		t.Fatal("token should have been removed from whitelist")
	}
}

func TestRevokeRefresh_InvalidTokenIsIdempotentSuccess(t *testing.T) {
	sc, _, _ := newRevokeSvc(t)
	resp, err := NewRevokeRefreshLogic(context.Background(), sc).RevokeRefresh(
		&user.RevokeRefreshReq{RefreshToken: "garbled.junk.string"})
	if err != nil {
		t.Fatalf("invalid token should not error, got %v", err)
	}
	if resp.Revoked {
		t.Fatal("invalid token should yield revoked=false")
	}
}

func TestRevokeRefresh_RejectsEmpty(t *testing.T) {
	sc, _, _ := newRevokeSvc(t)
	if _, err := NewRevokeRefreshLogic(context.Background(), sc).RevokeRefresh(
		&user.RevokeRefreshReq{RefreshToken: ""}); err == nil {
		t.Fatal("empty token must error")
	}
}

func TestRevokeRefresh_AccessTokenIsRejectedAsInvalid(t *testing.T) {
	// 用 access token 调用 RevokeRefresh，ParseRefresh 会拒绝（token_type 不对）。
	sc, signer, _ := newRevokeSvc(t)
	pair, _ := signer.IssuePair(1, "n")
	resp, err := NewRevokeRefreshLogic(context.Background(), sc).RevokeRefresh(
		&user.RevokeRefreshReq{RefreshToken: pair.AccessToken})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Revoked {
		t.Fatal("access token must not be revoked through RevokeRefresh")
	}
}
