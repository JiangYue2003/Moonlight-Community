package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

// AuthRpcMiddleware 通过 auth-rpc.VerifyToken 校验 access token；
// 与 jwtx.Signer 的本地校验等价，但避免 HTTP 网关持有 RSA 私钥/公钥。
type AuthRpcMiddleware struct {
	auth     userpb.AuthClient
	required bool
}

// NewAuthRpcMiddleware 强制鉴权（required=true）或可选鉴权（false）。
func NewAuthRpcMiddleware(auth userpb.AuthClient, required bool) *AuthRpcMiddleware {
	return &AuthRpcMiddleware{auth: auth, required: required}
}

func (m *AuthRpcMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := bearerToken(r)
		if token == "" {
			if m.required {
				httpx.ErrorCtx(r.Context(), w, errorx.New(errorx.CodeUnauthorized, "missing access token"))
				return
			}
			next(w, r)
			return
		}
		v, err := m.auth.VerifyToken(r.Context(), &userpb.VerifyTokenReq{AccessToken: token})
		if err != nil || !v.Valid {
			if m.required {
				httpx.ErrorCtx(r.Context(), w, errorx.New(errorx.CodeUnauthorized, "invalid access token"))
				return
			}
			next(w, r)
			return
		}
		ctx := ctxdata.WithUserId(r.Context(), v.UserId)
		next(w, r.WithContext(ctx))
	}
}

// HelperContext 补丁：包一层让 r.WithContext 调用方便。
func _() context.Context { return nil }

func bearerToken(r *http.Request) string {
	h := r.Header.Get("Authorization")
	if h == "" {
		return ""
	}
	const prefix = "Bearer "
	if strings.HasPrefix(h, prefix) {
		return strings.TrimSpace(h[len(prefix):])
	}
	return ""
}
