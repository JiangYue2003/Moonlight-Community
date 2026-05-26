package jwtx

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"
	"time"
)

func newSignerForTest(t *testing.T, accessTtl, refreshTtl time.Duration) *Signer {
	t.Helper()
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("genkey: %v", err)
	}
	s, err := NewSigner(Config{
		PrivateKey: priv,
		PublicKey:  &priv.PublicKey,
		Issuer:     "zhiguang-test",
		AccessTtl:  accessTtl,
		RefreshTtl: refreshTtl,
	})
	if err != nil {
		t.Fatalf("NewSigner: %v", err)
	}
	return s
}

func TestNewSigner_ValidatesKeys(t *testing.T) {
	if _, err := NewSigner(Config{}); err == nil {
		t.Fatal("NewSigner with no keys should fail")
	}
}

func TestNewSigner_DefaultsTtlAndIssuer(t *testing.T) {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	s, err := NewSigner(Config{PrivateKey: priv, PublicKey: &priv.PublicKey})
	if err != nil {
		t.Fatal(err)
	}
	if s.AccessTtl() != 15*time.Minute || s.RefreshTtl() != 7*24*time.Hour {
		t.Fatalf("default ttl drift: access=%v refresh=%v", s.AccessTtl(), s.RefreshTtl())
	}
}

func TestIssuePair_RoundTripAccessAndRefresh(t *testing.T) {
	s := newSignerForTest(t, 5*time.Minute, time.Hour)
	pair, err := s.IssuePair(42, "alice")
	if err != nil {
		t.Fatalf("IssuePair: %v", err)
	}
	if pair.AccessToken == "" || pair.RefreshToken == "" {
		t.Fatal("tokens must be non-empty")
	}
	if pair.RefreshTokenId == "" {
		t.Fatal("refresh jti must be set")
	}
	if pair.AccessExpiresAt.Before(time.Now()) {
		t.Fatal("access exp should be in the future")
	}
	if !pair.RefreshExpiresAt.After(pair.AccessExpiresAt) {
		t.Fatal("refresh exp must be later than access exp")
	}

	// access 验签
	ac, err := s.ParseAccess(pair.AccessToken)
	if err != nil {
		t.Fatalf("ParseAccess: %v", err)
	}
	if ac.Uid != 42 || ac.Nickname != "alice" || ac.TokenType != TokenTypeAccess {
		t.Fatalf("access claims drift: %+v", ac)
	}
	if ac.Issuer != "zhiguang-test" {
		t.Fatalf("issuer drift: %q", ac.Issuer)
	}

	// refresh 验签
	rc, err := s.ParseRefresh(pair.RefreshToken)
	if err != nil {
		t.Fatalf("ParseRefresh: %v", err)
	}
	if rc.Uid != 42 || rc.TokenType != TokenTypeRefresh || rc.ID != pair.RefreshTokenId {
		t.Fatalf("refresh claims drift: %+v", rc)
	}
}

func TestParseAccess_RejectsRefreshToken(t *testing.T) {
	s := newSignerForTest(t, time.Minute, time.Hour)
	pair, _ := s.IssuePair(1, "n")
	if _, err := s.ParseAccess(pair.RefreshToken); err == nil {
		t.Fatal("ParseAccess must reject refresh token")
	}
}

func TestParseRefresh_RejectsAccessToken(t *testing.T) {
	s := newSignerForTest(t, time.Minute, time.Hour)
	pair, _ := s.IssuePair(1, "n")
	if _, err := s.ParseRefresh(pair.AccessToken); err == nil {
		t.Fatal("ParseRefresh must reject access token")
	}
}

func TestParseAccess_RejectsExpired(t *testing.T) {
	// 极短 TTL：签发后等待一个窗口必失效。
	s := newSignerForTest(t, time.Millisecond, time.Hour)
	pair, _ := s.IssuePair(1, "n")
	time.Sleep(20 * time.Millisecond)
	if _, err := s.ParseAccess(pair.AccessToken); err == nil {
		t.Fatal("expired access token must be rejected")
	}
}

func TestParseAccess_RejectsTampered(t *testing.T) {
	s := newSignerForTest(t, time.Minute, time.Hour)
	pair, _ := s.IssuePair(1, "n")
	// 翻转 payload 中间一个字符（base64 字母替换），破坏签名。
	parts := strings.Split(pair.AccessToken, ".")
	if len(parts) != 3 {
		t.Fatal("token must be 3 segments")
	}
	if parts[1][0] == 'a' {
		parts[1] = "b" + parts[1][1:]
	} else {
		parts[1] = "a" + parts[1][1:]
	}
	tampered := strings.Join(parts, ".")
	if _, err := s.ParseAccess(tampered); err == nil {
		t.Fatal("tampered access token must be rejected")
	}
}

func TestParseAccess_RejectsCrossKey(t *testing.T) {
	// 用另一对密钥签发的 token，本 signer 验签必须失败。
	other := newSignerForTest(t, time.Minute, time.Hour)
	mine := newSignerForTest(t, time.Minute, time.Hour)
	pair, _ := other.IssuePair(1, "n")
	if _, err := mine.ParseAccess(pair.AccessToken); err == nil {
		t.Fatal("cross-key validation must fail")
	}
}

func TestParseAccess_RejectsWrongIssuer(t *testing.T) {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	other, _ := NewSigner(Config{
		PrivateKey: priv, PublicKey: &priv.PublicKey,
		Issuer: "evil", AccessTtl: time.Minute, RefreshTtl: time.Hour,
	})
	mine, _ := NewSigner(Config{
		PrivateKey: priv, PublicKey: &priv.PublicKey,
		Issuer: "zhiguang-test", AccessTtl: time.Minute, RefreshTtl: time.Hour,
	})
	pair, _ := other.IssuePair(1, "n")
	if _, err := mine.ParseAccess(pair.AccessToken); err == nil {
		t.Fatal("wrong-issuer token must be rejected")
	}
}
