package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	authclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/auth"
)

func OptionalAuth(auth authclient.Auth) gin.HandlerFunc {
	return authMiddleware(auth, false)
}

func RequiredAuth(auth authclient.Auth) gin.HandlerFunc {
	return authMiddleware(auth, true)
}

func authMiddleware(auth authclient.Auth, required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := bearerToken(c.GetHeader("Authorization"))
		if token == "" {
			if required {
				gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing access token"))
				c.Abort()
				return
			}
			c.Next()
			return
		}
		resp, err := auth.VerifyToken(c.Request.Context(), &authclient.VerifyTokenReq{AccessToken: token})
		if err != nil || resp == nil || !resp.Valid {
			if required {
				gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "invalid access token"))
				c.Abort()
				return
			}
			c.Next()
			return
		}
		ctx := ctxdata.WithUserId(c.Request.Context(), resp.UserId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func bearerToken(h string) string {
	const prefix = "Bearer "
	if !strings.HasPrefix(h, prefix) {
		return ""
	}
	return strings.TrimSpace(h[len(prefix):])
}
