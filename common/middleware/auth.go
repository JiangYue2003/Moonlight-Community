// Package middleware 提供 HTTP 中间件。
//
// AuthMiddleware：从 Authorization: Bearer <jwt> 解析 access token，
// 通过 jwtx.Signer 验签，将 userId 写入 context.Context。
package middleware

import (
	"net/http"
	"strings"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	"github.com/zhiguang/zhiguang-go/pkg/responsex"
)

// AuthMiddleware 强制校验 access token；缺失或无效返回 401。
type AuthMiddleware struct {
	signer *jwtx.Signer
}

// NewAuthMiddleware 构造中间件，依赖 jwtx.Signer 进行 RS256 验签。
func NewAuthMiddleware(signer *jwtx.Signer) *AuthMiddleware {
	return &AuthMiddleware{signer: signer}
}

// Handle 实现 http.Handler 包装：解析→ctx 注入 userId→放行。
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractBearer(r)
		if token == "" {
			responsex.Fail(w, errorx.New(errorx.CodeUnauthorized, "missing access token"))
			return
		}
		claims, err := m.signer.ParseAccess(token)
		if err != nil {
			responsex.Fail(w, errorx.Wrap(errorx.CodeUnauthorized, "invalid access token", err))
			return
		}
		ctx := ctxdata.WithUserId(r.Context(), claims.Uid)
		next(w, r.WithContext(ctx))
	}
}

// HandleOptional 解析 token 但不强制；无 token 也放行（用于公开 + 个性化接口）。
func (m *AuthMiddleware) HandleOptional(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractBearer(r)
		if token == "" {
			next(w, r)
			return
		}
		claims, err := m.signer.ParseAccess(token)
		if err == nil {
			ctx := ctxdata.WithUserId(r.Context(), claims.Uid)
			r = r.WithContext(ctx)
		}
		next(w, r)
	}
}

func extractBearer(r *http.Request) string {
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
