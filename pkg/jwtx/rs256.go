// Package jwtx 提供 RS256 JWT 的签发与验证。
//
// 与原 Java Nimbus 版本保持兼容：
//   - 签名算法 RS256（PKCS#8 PEM 私钥 / X.509 PEM 公钥）
//   - Claims：token_type ∈ {access, refresh}, uid, nickname (access only),
//     jti (refresh only), iss, sub, exp, iat
package jwtx

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Signer 持有私钥（签发用）与公钥（验签用）。
type Signer struct {
	priv       *rsa.PrivateKey
	pub        *rsa.PublicKey
	issuer     string
	accessTtl  time.Duration
	refreshTtl time.Duration
}

// Config 创建 Signer 所需参数。
type Config struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Issuer     string
	AccessTtl  time.Duration
	RefreshTtl time.Duration
}

// NewSigner 构造 Signer；priv 用于签发，pub 用于验签。
func NewSigner(c Config) (*Signer, error) {
	if c.PrivateKey == nil || c.PublicKey == nil {
		return nil, errors.New("jwtx: priv/pub key required")
	}
	if c.Issuer == "" {
		c.Issuer = "zhiguang"
	}
	if c.AccessTtl <= 0 {
		c.AccessTtl = 15 * time.Minute
	}
	if c.RefreshTtl <= 0 {
		c.RefreshTtl = 7 * 24 * time.Hour
	}
	return &Signer{
		priv:       c.PrivateKey,
		pub:        c.PublicKey,
		issuer:     c.Issuer,
		accessTtl:  c.AccessTtl,
		refreshTtl: c.RefreshTtl,
	}, nil
}

// TokenPair 与 proto 中 TokenPair 字段保持一致。
type TokenPair struct {
	AccessToken      string
	AccessExpiresAt  time.Time
	RefreshToken     string
	RefreshExpiresAt time.Time
	RefreshTokenId   string
}

// IssuePair 一次性签发 access + refresh 双令牌。
func (s *Signer) IssuePair(uid int64, nickname string) (*TokenPair, error) {
	now := time.Now()
	accessExp := now.Add(s.accessTtl)
	refreshExp := now.Add(s.refreshTtl)
	jti := uuid.NewString()

	access, err := s.signAccess(uid, nickname, now, accessExp)
	if err != nil {
		return nil, err
	}
	refresh, err := s.signRefresh(uid, jti, now, refreshExp)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:      access,
		AccessExpiresAt:  accessExp,
		RefreshToken:     refresh,
		RefreshExpiresAt: refreshExp,
		RefreshTokenId:   jti,
	}, nil
}

func (s *Signer) signAccess(uid int64, nickname string, iat, exp time.Time) (string, error) {
	c := AccessClaims{
		TokenType: TokenTypeAccess,
		Uid:       uid,
		Nickname:  nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   fmt.Sprintf("%d", uid),
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return t.SignedString(s.priv)
}

func (s *Signer) signRefresh(uid int64, jti string, iat, exp time.Time) (string, error) {
	c := RefreshClaims{
		TokenType: TokenTypeRefresh,
		Uid:       uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   fmt.Sprintf("%d", uid),
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return t.SignedString(s.priv)
}

// ParseAccess 解析并验证 access token，强校验 token_type=access。
func (s *Signer) ParseAccess(tokenStr string) (*AccessClaims, error) {
	out := &AccessClaims{}
	if _, err := s.parse(tokenStr, out); err != nil {
		return nil, err
	}
	if out.TokenType != TokenTypeAccess {
		return nil, errors.New("jwtx: not an access token")
	}
	return out, nil
}

// ParseRefresh 解析并验证 refresh token，强校验 token_type=refresh 与 jti。
func (s *Signer) ParseRefresh(tokenStr string) (*RefreshClaims, error) {
	out := &RefreshClaims{}
	if _, err := s.parse(tokenStr, out); err != nil {
		return nil, err
	}
	if out.TokenType != TokenTypeRefresh {
		return nil, errors.New("jwtx: not a refresh token")
	}
	if out.ID == "" {
		return nil, errors.New("jwtx: refresh token missing jti")
	}
	return out, nil
}

func (s *Signer) parse(tokenStr string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("jwtx: unexpected signing method: %v", t.Header["alg"])
		}
		return s.pub, nil
	}, jwt.WithIssuer(s.issuer))
}

// AccessTtl / RefreshTtl 暴露给上层（如 Redis 白名单 TTL）。
func (s *Signer) AccessTtl() time.Duration  { return s.accessTtl }
func (s *Signer) RefreshTtl() time.Duration { return s.refreshTtl }
