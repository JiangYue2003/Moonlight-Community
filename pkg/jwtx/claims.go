package jwtx

import "github.com/golang-jwt/jwt/v5"

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

// AccessClaims 与 Java Nimbus 实现保持字段一致。
type AccessClaims struct {
	TokenType string `json:"token_type"`
	Uid       int64  `json:"uid"`
	Nickname  string `json:"nickname,omitempty"`
	jwt.RegisteredClaims
}

// RefreshClaims 含 jti 用于白名单存储与单次使用旋转。
type RefreshClaims struct {
	TokenType string `json:"token_type"`
	Uid       int64  `json:"uid"`
	jwt.RegisteredClaims
}
