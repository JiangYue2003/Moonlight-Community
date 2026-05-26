// 由人工调整：接入 AuthRpcMiddleware（强制鉴权）。

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zhiguang/zhiguang-go/common/middleware"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	authMw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, true)

	// 按主目录 docs 约定：relation 全部接口要求携带 Authorization。
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMw.Handle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "/follow", Handler: FollowHandler(serverCtx)},
				{Method: http.MethodPost, Path: "/unfollow", Handler: UnfollowHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/status", Handler: StatusHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/following", Handler: ListFollowingHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/followers", Handler: ListFollowersHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/counter", Handler: CounterHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/api/v1/relation"),
	)
}
